package sentences_service

import (
	"backend/protos/gen/go/sentences"
	attachment_service "backend/services/gateway/serivce/attachmentService"
	attachment_types_service "backend/services/gateway/serivce/attachmentTypesService"
	departments_service "backend/services/gateway/serivce/departmentsService"
	direction_service "backend/services/gateway/serivce/directionService"
	historyes_service "backend/services/gateway/serivce/historyesService"
	implementors_service "backend/services/gateway/serivce/implementorsService"
	priorities_service "backend/services/gateway/serivce/prioritiesService"
	sentences_attachments_service "backend/services/gateway/serivce/sentencesAttachmentsService"
	statuses_service "backend/services/gateway/serivce/statusesService"
	"errors"
	"log/slog"
)

var (
	ErrSentenceCreation   = errors.New("invalid sentence creation")
	ErrInvalidStatus      = errors.New("invalid status ID")
	ErrInvalidDirection   = errors.New("invalid direction ID")
	ErrInvalidPriority    = errors.New("invalid priority ID")
	ErrInvalidDepartment  = errors.New("invalid department ID")
	ErrInvalidImplementor = errors.New("invalid implementor ID")
	ErrSentenceNotFound   = errors.New("sentence not found")
)

type Service struct {
	sentenceClient        sentences.SentencesProviderClient
	statusService         *statuses_service.Service
	directionService      *direction_service.Service
	priorityService       *priorities_service.Service
	departmentService     *departments_service.Service
	implementorService    *implementors_service.Service
	attachmentService     *attachment_service.Service
	attachmentTypeService *attachment_types_service.Service
	historyService        *historyes_service.Service
	sentenceAttachService *sentences_attachments_service.Service
	logger                *slog.Logger
}

func New(
	sentenceClient sentences.SentencesProviderClient,
	statusService *statuses_service.Service,
	directionService *direction_service.Service,
	priorityService *priorities_service.Service,
	departmentService *departments_service.Service,
	implementorService *implementors_service.Service,
	attachmentService *attachment_service.Service,
	attachmentTypeService *attachment_types_service.Service,
	historyService *historyes_service.Service,
	sentenceAttachService *sentences_attachments_service.Service,
	logger *slog.Logger,
) *Service {
	return &Service{
		sentenceClient:        sentenceClient,
		statusService:         statusService,
		directionService:      directionService,
		priorityService:       priorityService,
		departmentService:     departmentService,
		implementorService:    implementorService,
		attachmentService:     attachmentService,
		attachmentTypeService: attachmentTypeService,
		historyService:        historyService,
		sentenceAttachService: sentenceAttachService,
		logger:                logger,
	}
}
