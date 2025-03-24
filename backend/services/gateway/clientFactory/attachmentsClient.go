package client_factory

import (
	"backend/protos/gen/go/attachments"
	"context"
	"google.golang.org/grpc"
)

func (p *ClientProvider) GetAttachmentsClient(ctx context.Context, addr string) (AttachmentsClient, error) {
	client, err := p.getClient(ctx, ServiceTypeAttachments, addr, func(conn *grpc.ClientConn) interface{} {
		return &attachmentsClient{
			AttachmentsProviderClient: attachments.NewAttachmentsProviderClient(conn),
			conn:                      conn,
		}
	})
	if err != nil {
		return nil, err
	}
	return client.(AttachmentsClient), nil
}

type attachmentsClient struct {
	attachments.AttachmentsProviderClient
	conn *grpc.ClientConn
}

func (c *attachmentsClient) Close() error {
	return c.conn.Close()
}
