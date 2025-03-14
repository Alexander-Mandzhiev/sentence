package client_factory

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
	return [...]string{
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
	}[s]
}
