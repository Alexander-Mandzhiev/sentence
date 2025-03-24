package client_factory

import (
	"backend/protos/gen/go/sentences"
	"context"
	"google.golang.org/grpc"
)

func (p *ClientProvider) GetSentencesClient(ctx context.Context, addr string) (SentencesClient, error) {
	client, err := p.getClient(ctx, ServiceTypeSentences, addr, func(conn *grpc.ClientConn) interface{} {
		return &sentencesClient{
			SentencesProviderClient: sentences.NewSentencesProviderClient(conn),
			conn:                    conn,
		}
	})
	if err != nil {
		return nil, err
	}
	return client.(SentencesClient), nil
}

type sentencesClient struct {
	sentences.SentencesProviderClient
	conn *grpc.ClientConn
}

func (c *sentencesClient) Close() error {
	return c.conn.Close()
}
