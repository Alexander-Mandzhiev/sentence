package client_factory

import (
	"backend/protos/gen/go/history"
	"context"
	"google.golang.org/grpc"
)

func (p *ClientProvider) GetHistoryClient(ctx context.Context, addr string) (HistoryClient, error) {
	client, err := p.getClient(ctx, ServiceTypeHistory, addr, func(conn *grpc.ClientConn) interface{} {
		return &historyClient{
			HistoryProviderClient: history.NewHistoryProviderClient(conn),
			conn:                  conn,
		}
	})
	if err != nil {
		return nil, err
	}
	return client.(HistoryClient), nil
}

type historyClient struct {
	history.HistoryProviderClient
	conn *grpc.ClientConn
}

func (c *historyClient) Close() error {
	return c.conn.Close()
}
