package client_factory

import (
	"backend/pkg/server/grpc_client"
	"log/slog"
	"sync"
)

type ServiceType int

const (
	ServiceTypeAttachmentTypes ServiceType = iota
	ServiceTypeDirections
	ServiceTypeHistory
	ServiceTypeAttachments
	ServiceTypeDepartments
	ServiceTypeImplementors
	ServiceTypePriorities
	ServiceTypeStatuses
	ServiceTypeSentences
	ServiceTypeSentencesAttachments
)

func (s ServiceType) String() string {
	names := [...]string{
		"attachment_types",
		"directions",
		"history",
		"attachments",
		"departments",
		"implementors",
		"priorities",
		"statuses",
		"sentences",
		"sentences_attachments",
	}
	if int(s) < len(names) {
		return names[s]
	}
	return "unknown"
}

type ClientProvider struct {
	manager      *grpc_client.GRPCClientManager
	clients      map[ServiceType]interface{}
	clientsMutex sync.RWMutex
	logger       *slog.Logger
}

func New(manager *grpc_client.GRPCClientManager, logger *slog.Logger) *ClientProvider {
	return &ClientProvider{
		manager: manager,
		clients: make(map[ServiceType]interface{}),
		logger:  logger,
	}
}
