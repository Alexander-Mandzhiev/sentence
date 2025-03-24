package client_factory

import (
	"backend/protos/gen/go/sentences_attachments"
	"context"
	"google.golang.org/grpc"
)

func (p *ClientProvider) GetSentencesAttachmentsClient(ctx context.Context, addr string) (SentencesAttachmentsClient, error) {
	client, err := p.getClient(ctx, ServiceTypeSentencesAttachments, addr, func(conn *grpc.ClientConn) interface{} {
		return &sentencesAttachmentsClient{
			SentencesAttachmentsProviderClient: sentences_attachments.NewSentencesAttachmentsProviderClient(conn),
			conn:                               conn,
		}
	})
	if err != nil {
		return nil, err
	}
	return client.(SentencesAttachmentsClient), nil
}

type sentencesAttachmentsClient struct {
	sentences_attachments.SentencesAttachmentsProviderClient
	conn *grpc.ClientConn
}

func (c *sentencesAttachmentsClient) Close() error {
	return c.conn.Close()
}
