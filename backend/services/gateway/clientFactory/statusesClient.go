package client_factory

import (
	"backend/protos/gen/go/statuses"
	"context"
	"google.golang.org/grpc"
)

func (p *ClientProvider) GetStatusesClient(ctx context.Context, addr string) (StatusesClient, error) {
	client, err := p.getClient(ctx, ServiceTypeStatuses, addr, func(conn *grpc.ClientConn) interface{} {
		return &statusesClient{
			StatusProviderClient: statuses.NewStatusProviderClient(conn),
			conn:                 conn,
		}
	})
	if err != nil {
		return nil, err
	}
	return client.(StatusesClient), nil
}

type statusesClient struct {
	statuses.StatusProviderClient
	conn *grpc.ClientConn
}

func (c *statusesClient) Close() error {
	return c.conn.Close()
}
