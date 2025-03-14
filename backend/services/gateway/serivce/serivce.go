package serivce

import client "backend/services/gateway/clientFactory"

type Service struct {
	AttachmentTypesClient      client.AttachmentTypesClient
	DirectionsClient           client.DirectionsClient
	HistoryClient              client.HistoryClient
	AttachmentsClient          client.AttachmentsClient
	DepartmentsClient          client.DepartmentsClient
	ImplementorsClient         client.ImplementorsClient
	PrioritiesClient           client.PrioritiesClient
	StatusesClient             client.StatusesClient
	SentencesClient            client.SentencesClient
	SentencesAttachmentsClient client.SentencesAttachmentsClient
}

func New(
	attachmentTypesClient client.AttachmentTypesClient,
	directionsClient client.DirectionsClient,
	historyClient client.HistoryClient,
	attachmentsClient client.AttachmentsClient,
	departmentsClient client.DepartmentsClient,
	implementorsClient client.ImplementorsClient,
	prioritiesClient client.PrioritiesClient,
	statusesClient client.StatusesClient,
	sentencesClient client.SentencesClient,
	sentencesAttachmentsClient client.SentencesAttachmentsClient,
) *Service {
	return &Service{
		AttachmentTypesClient:      attachmentTypesClient,
		DirectionsClient:           directionsClient,
		HistoryClient:              historyClient,
		AttachmentsClient:          attachmentsClient,
		DepartmentsClient:          departmentsClient,
		ImplementorsClient:         implementorsClient,
		PrioritiesClient:           prioritiesClient,
		StatusesClient:             statusesClient,
		SentencesClient:            sentencesClient,
		SentencesAttachmentsClient: sentencesAttachmentsClient,
	}
}
