package client_factory

import (
	"backend/protos/gen/go/attachment_types"
	"context"
	"google.golang.org/grpc"
)

func (p *ClientProvider) GetAttachmentTypesClient(ctx context.Context, addr string) (AttachmentTypesClient, error) {
	client, err := p.getClient(ctx, ServiceTypeAttachmentTypes, addr, func(conn *grpc.ClientConn) interface{} {
		return &attachmentTypesClient{
			AttachmentTypesProviderClient: attachment_types.NewAttachmentTypesProviderClient(conn),
			conn:                          conn,
		}
	})
	if err != nil {
		return nil, err
	}
	return client.(AttachmentTypesClient), nil
}

type attachmentTypesClient struct {
	attachment_types.AttachmentTypesProviderClient
	conn *grpc.ClientConn
}

func (c *attachmentTypesClient) Close() error {
	return c.conn.Close()
}
