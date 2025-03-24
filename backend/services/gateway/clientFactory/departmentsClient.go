package client_factory

import (
	"backend/protos/gen/go/departments"
	"context"
	"google.golang.org/grpc"
)

func (p *ClientProvider) GetDepartmentsClient(ctx context.Context, addr string) (DepartmentsClient, error) {
	client, err := p.getClient(ctx, ServiceTypeDepartments, addr, func(conn *grpc.ClientConn) interface{} {
		return &departmentsClient{
			DepartmentsProviderClient: departments.NewDepartmentsProviderClient(conn),
			conn:                      conn,
		}
	})
	if err != nil {
		return nil, err
	}
	return client.(DepartmentsClient), nil
}

type departmentsClient struct {
	departments.DepartmentsProviderClient
	conn *grpc.ClientConn
}

func (c *departmentsClient) Close() error {
	return c.conn.Close()
}
