package grpc_client

import (
	"fmt"
	"log"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClientManager struct {
	connections sync.Map
}

func NewGRPCClientManager() *GRPCClientManager {
	return &GRPCClientManager{}
}

func (m *GRPCClientManager) GetClientConn(addr string) (*grpc.ClientConn, error) {
	if conn, exists := m.connections.Load(addr); exists {
		if conn.(*grpc.ClientConn).GetState() != connectivity.Shutdown {
			return conn.(*grpc.ClientConn), nil
		}
		m.connections.Delete(addr)
		if err := conn.(*grpc.ClientConn).Close(); err != nil {
			log.Printf("Failed to close connection to %s: %v\n", addr, err)
		}
	}

	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(16<<20)))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to service at %s: %w", addr, err)
	}

	defer func() {
		if err != nil {
			if closeErr := conn.Close(); closeErr != nil {
				log.Printf("Failed to close connection to %s: %v\n", addr, closeErr)
			}
		}
	}()

	m.connections.Store(addr, conn)
	return conn, nil
}

func (m *GRPCClientManager) CloseAll() error {
	var errs []error

	m.connections.Range(func(key, value interface{}) bool {
		if value.(*grpc.ClientConn) != nil {
			if err := value.(*grpc.ClientConn).Close(); err != nil {
				errs = append(errs, fmt.Errorf("failed to close connection to %s: %w", key.(string), err))
			}
		}
		m.connections.Delete(key)
		return true
	})

	if len(errs) > 0 {
		return fmt.Errorf("errors occurred while closing connections: %v", errs)
	}
	return nil
}
