package client_factory

import (
	"backend/protos/gen/go/implementors"
	"context"
	"google.golang.org/grpc"
)

func (p *ClientProvider) GetImplementorsClient(ctx context.Context, addr string) (ImplementorsClient, error) {
	client, err := p.getClient(ctx, ServiceTypeImplementors, addr, func(conn *grpc.ClientConn) interface{} {
		return &implementorsClient{
			ImplementorsProviderClient: implementors.NewImplementorsProviderClient(conn),
			conn:                       conn,
		}
	})
	if err != nil {
		return nil, err
	}
	return client.(ImplementorsClient), nil
}

type implementorsClient struct {
	implementors.ImplementorsProviderClient
	conn *grpc.ClientConn
}

func (c *implementorsClient) Close() error {
	return c.conn.Close()
}
