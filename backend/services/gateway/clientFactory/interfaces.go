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
	"context"
)

type AttachmentTypesClient interface {
	Create(ctx context.Context, req *attachment_types.CreateAttachmentTypeRequest) (*attachment_types.AttachmentTypeResponse, error)
	Get(ctx context.Context, req *attachment_types.GetAttachmentTypeRequest) (*attachment_types.AttachmentTypeResponse, error)
	Update(ctx context.Context, req *attachment_types.UpdateAttachmentTypeRequest) (*attachment_types.AttachmentTypeResponse, error)
	Delete(ctx context.Context, req *attachment_types.DeleteAttachmentTypeRequest) (*attachment_types.DeleteAttachmentTypeResponse, error)
	List(ctx context.Context, req *attachment_types.ListAttachmentTypesRequest) (*attachment_types.AttachmentTypesListResponse, error)
}

type DirectionsClient interface {
	Create(ctx context.Context, req *directions.CreateDirectionRequest) (*directions.DirectionResponse, error)
	Get(ctx context.Context, req *directions.GetDirectionRequest) (*directions.DirectionResponse, error)
	Update(ctx context.Context, req *directions.UpdateDirectionRequest) (*directions.DirectionResponse, error)
	Delete(ctx context.Context, req *directions.DeleteDirectionRequest) (*directions.DeleteDirectionResponse, error)
	List(ctx context.Context, req *directions.ListDirectionsRequest) (*directions.DirectionsListResponse, error)
}

type HistoryClient interface {
	Create(ctx context.Context, req *history.CreateHistoryRequest) (*history.HistoryResponse, error)
	Get(ctx context.Context, req *history.GetHistoryRequest) (*history.HistoryResponse, error)
	ListBySentence(ctx context.Context, req *history.ListBySentenceRequest) (*history.HistoryListResponse, error)
}

type AttachmentsClient interface {
	Create(ctx context.Context, req *attachments.CreateAttachmentRequest) (*attachments.AttachmentResponse, error)
	Get(ctx context.Context, req *attachments.GetAttachmentRequest) (*attachments.AttachmentResponse, error)
	Update(ctx context.Context, req *attachments.UpdateAttachmentRequest) (*attachments.AttachmentResponse, error)
	Delete(ctx context.Context, req *attachments.DeleteAttachmentRequest) (*attachments.DeleteAttachmentResponse, error)
	List(ctx context.Context, req *attachments.ListAttachmentsRequest) (*attachments.AttachmentsListResponse, error)
}

type DepartmentsClient interface {
	Create(ctx context.Context, req *departments.CreateDepartmentRequest) (*departments.DepartmentResponse, error)
	Get(ctx context.Context, req *departments.GetDepartmentRequest) (*departments.DepartmentResponse, error)
	Update(ctx context.Context, req *departments.UpdateDepartmentRequest) (*departments.DepartmentResponse, error)
	Delete(ctx context.Context, req *departments.DeleteDepartmentRequest) (*departments.DeleteDepartmentResponse, error)
	List(ctx context.Context, req *departments.ListDepartmentsRequest) (*departments.DepartmentsListResponse, error)
}

type ImplementorsClient interface {
	Create(ctx context.Context, req *implementors.CreateImplementorRequest) (*implementors.ImplementorResponse, error)
	Get(ctx context.Context, req *implementors.GetImplementorRequest) (*implementors.ImplementorResponse, error)
	Update(ctx context.Context, req *implementors.UpdateImplementorRequest) (*implementors.ImplementorResponse, error)
	Delete(ctx context.Context, req *implementors.DeleteImplementorRequest) (*implementors.DeleteImplementorResponse, error)
	List(ctx context.Context, req *implementors.ListImplementorsRequest) (*implementors.ImplementorsListResponse, error)
}

type PrioritiesClient interface {
	Create(ctx context.Context, req *priorities.CreatePrioritiesRequest) (*priorities.PrioritiesResponse, error)
	Get(ctx context.Context, req *priorities.GetPrioritiesRequest) (*priorities.PrioritiesResponse, error)
	Update(ctx context.Context, req *priorities.UpdatePrioritiesRequest) (*priorities.PrioritiesResponse, error)
	Delete(ctx context.Context, req *priorities.DeletePrioritiesRequest) (*priorities.DeletePrioritiesResponse, error)
	List(ctx context.Context, req *priorities.ListPrioritiesRequest) (*priorities.PrioritiesListResponse, error)
}

type StatusesClient interface {
	Create(ctx context.Context, req *statuses.CreateStatusRequest) (*statuses.StatusResponse, error)
	Get(ctx context.Context, req *statuses.GetStatusRequest) (*statuses.StatusResponse, error)
	Update(ctx context.Context, req *statuses.UpdateStatusRequest) (*statuses.StatusResponse, error)
	Delete(ctx context.Context, req *statuses.DeleteStatusRequest) (*statuses.DeleteStatusResponse, error)
	List(ctx context.Context, req *statuses.ListStatusesRequest) (*statuses.StatusListResponse, error)
}

type SentencesClient interface {
	Create(ctx context.Context, req *sentences.CreateSentenceRequest) (*sentences.SentenceResponse, error)
	Get(ctx context.Context, req *sentences.GetSentenceRequest) (*sentences.SentenceResponse, error)
	Update(ctx context.Context, req *sentences.UpdateSentenceRequest) (*sentences.SentenceResponse, error)
	Delete(ctx context.Context, req *sentences.DeleteSentenceRequest) (*sentences.DeleteSentenceResponse, error)
	List(ctx context.Context, req *sentences.ListSentencesRequest) (*sentences.SentencesListResponse, error)
}

type SentencesAttachmentsClient interface {
	Create(ctx context.Context, req *sentences_attachments.CreateSentenceAttachmentRequest) (*sentences_attachments.SentenceAttachmentResponse, error)
	Delete(ctx context.Context, req *sentences_attachments.DeleteSentenceAttachmentRequest) (*sentences_attachments.DeleteSentenceAttachmentResponse, error)
	ListBySentence(ctx context.Context, req *sentences_attachments.ListBySentenceRequest) (*sentences_attachments.SentenceAttachmentsListResponse, error)
	ListByAttachment(ctx context.Context, req *sentences_attachments.ListByAttachmentRequest) (*sentences_attachments.SentenceAttachmentsListResponse, error)
}
