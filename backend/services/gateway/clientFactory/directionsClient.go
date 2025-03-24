package client_factory

import (
	"backend/protos/gen/go/directions"
	"context"
	"google.golang.org/grpc"
)

func (p *ClientProvider) GetDirectionsClient(ctx context.Context, addr string) (DirectionsClient, error) {
	client, err := p.getClient(ctx, ServiceTypeDirections, addr, func(conn *grpc.ClientConn) interface{} {
		return &directionsClient{
			DirectionsProviderClient: directions.NewDirectionsProviderClient(conn),
			conn:                     conn,
		}
	})
	if err != nil {
		return nil, err
	}
	return client.(DirectionsClient), nil
}

type directionsClient struct {
	directions.DirectionsProviderClient
	conn *grpc.ClientConn
}

func (c *directionsClient) Close() error {
	return c.conn.Close()
}
