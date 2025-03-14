package client_factory

import (
	"backend/pkg/server/grpc_client"
	"backend/protos/gen/go/attachment_types"
	"backend/protos/gen/go/attachments"
	"backend/protos/gen/go/departments"
	"backend/protos/gen/go/directions"
	"backend/protos/gen/go/history"
	"backend/protos/gen/go/implementors"
	"backend/protos/gen/go/priorities"
	"backend/protos/gen/go/sentences"
	"backend/protos/gen/go/sentences_attachments"
	"backend/protos/gen/go/statuses"
	"fmt"
)

type ClientFactory interface {
	NewClient(addr string, serviceType ServiceType) (interface{}, error)
}

type clientFactory struct {
	manager *grpc_client.GRPCClientManager
}

func NewClientFactory(manager *grpc_client.GRPCClientManager) ClientFactory {
	return &clientFactory{manager: manager}
}

func (f *clientFactory) NewClient(addr string, serviceType ServiceType) (interface{}, error) {
	conn, err := f.manager.GetClientConn(addr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s: %w", serviceType.String(), err)
	}

	switch serviceType {
	case ServiceTypeAttachmentTypes:
		return attachment_types.NewAttachmentTypesProviderClient(conn), nil
	case ServiceTypeDirections:
		return directions.NewDirectionsProviderClient(conn), nil
	case ServiceTypeHistory:
		return history.NewHistoryProviderClient(conn), nil
	case ServiceTypeAttachments:
		return attachments.NewAttachmentsProviderClient(conn), nil
	case ServiceTypeDepartments:
		return departments.NewDepartmentsProviderClient(conn), nil
	case ServiceTypeImplementors:
		return implementors.NewImplementorsProviderClient(conn), nil
	case ServiceTypePriorities:
		return priorities.NewPrioritiesProviderClient(conn), nil
	case ServiceTypeStatuses:
		return statuses.NewStatusProviderClient(conn), nil
	case ServiceTypeSentences:
		return sentences.NewSentencesProviderClient(conn), nil
	case ServiceTypeSentencesAttachments:
		return sentences_attachments.NewSentencesAttachmentsProviderClient(conn), nil
	default:
		return nil, fmt.Errorf("unknown service type: %v", serviceType)
	}
}
