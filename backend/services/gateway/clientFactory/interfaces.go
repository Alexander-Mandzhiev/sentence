package client_factory

import (
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
)

type (
	AttachmentTypesClient interface {
		attachment_types.AttachmentTypesProviderClient
		Close() error
	}

	DirectionsClient interface {
		directions.DirectionsProviderClient
		Close() error
	}

	HistoryClient interface {
		history.HistoryProviderClient
		Close() error
	}

	AttachmentsClient interface {
		attachments.AttachmentsProviderClient
		Close() error
	}

	DepartmentsClient interface {
		departments.DepartmentsProviderClient
		Close() error
	}

	ImplementorsClient interface {
		implementors.ImplementorsProviderClient
		Close() error
	}

	PrioritiesClient interface {
		priorities.PrioritiesProviderClient
		Close() error
	}

	StatusesClient interface {
		statuses.StatusProviderClient
		Close() error
	}

	SentencesClient interface {
		sentences.SentencesProviderClient
		Close() error
	}

	SentencesAttachmentsClient interface {
		sentences_attachments.SentencesAttachmentsProviderClient
		Close() error
	}
)
