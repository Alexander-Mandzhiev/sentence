package client_factory

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func (p *ClientProvider) getClient(
	ctx context.Context,
	serviceType ServiceType,
	addr string,
	factory func(*grpc.ClientConn) interface{},
) (interface{}, error) {
	p.clientsMutex.RLock()
	if client, exists := p.clients[serviceType]; exists {
		p.clientsMutex.RUnlock()
		return client, nil
	}
	p.clientsMutex.RUnlock()

	conn, err := p.manager.GetClientConn(ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s: %w", serviceType.String(), err)
	}

	p.clientsMutex.Lock()
	defer p.clientsMutex.Unlock()

	if client, exists := p.clients[serviceType]; exists {
		return client, nil
	}

	client := factory(conn)
	p.clients[serviceType] = client
	p.logger.Info("gRPC client created", "service", serviceType.String(), "address", addr)
	return client, nil
}

func (p *ClientProvider) Close() error {
	p.clientsMutex.Lock()
	defer p.clientsMutex.Unlock()

	var errs []error
	for _, client := range p.clients {
		if closer, ok := client.(interface{ Close() error }); ok {
			if err := closer.Close(); err != nil {
				errs = append(errs, err)
			}
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors while closing clients: %v", errs)
	}
	return nil
}
