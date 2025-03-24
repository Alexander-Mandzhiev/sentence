package client_factory

import (
	"backend/protos/gen/go/priorities"
	"context"
	"google.golang.org/grpc"
)

func (p *ClientProvider) GetPrioritiesClient(ctx context.Context, addr string) (PrioritiesClient, error) {
	client, err := p.getClient(ctx, ServiceTypePriorities, addr, func(conn *grpc.ClientConn) interface{} {
		return &prioritiesClient{
			PrioritiesProviderClient: priorities.NewPrioritiesProviderClient(conn),
			conn:                     conn,
		}
	})
	if err != nil {
		return nil, err
	}
	return client.(PrioritiesClient), nil
}

type prioritiesClient struct {
	priorities.PrioritiesProviderClient
	conn *grpc.ClientConn
}

func (c *prioritiesClient) Close() error {
	return c.conn.Close()
}
