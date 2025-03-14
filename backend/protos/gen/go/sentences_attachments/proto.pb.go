// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.1
// source: sentences_attachments/proto.proto

package sentences_attachments

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateSentenceAttachmentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SentenceId    int64                  `protobuf:"varint,1,opt,name=sentence_id,json=sentenceId,proto3" json:"sentence_id,omitempty"`
	AttachmentId  int32                  `protobuf:"varint,2,opt,name=attachment_id,json=attachmentId,proto3" json:"attachment_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateSentenceAttachmentRequest) Reset() {
	*x = CreateSentenceAttachmentRequest{}
	mi := &file_sentences_attachments_proto_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSentenceAttachmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSentenceAttachmentRequest) ProtoMessage() {}

func (x *CreateSentenceAttachmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sentences_attachments_proto_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSentenceAttachmentRequest.ProtoReflect.Descriptor instead.
func (*CreateSentenceAttachmentRequest) Descriptor() ([]byte, []int) {
	return file_sentences_attachments_proto_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSentenceAttachmentRequest) GetSentenceId() int64 {
	if x != nil {
		return x.SentenceId
	}
	return 0
}

func (x *CreateSentenceAttachmentRequest) GetAttachmentId() int32 {
	if x != nil {
		return x.AttachmentId
	}
	return 0
}

type SentenceAttachmentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SentenceId    int64                  `protobuf:"varint,1,opt,name=sentence_id,json=sentenceId,proto3" json:"sentence_id,omitempty"`
	AttachmentId  int32                  `protobuf:"varint,2,opt,name=attachment_id,json=attachmentId,proto3" json:"attachment_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SentenceAttachmentResponse) Reset() {
	*x = SentenceAttachmentResponse{}
	mi := &file_sentences_attachments_proto_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SentenceAttachmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SentenceAttachmentResponse) ProtoMessage() {}

func (x *SentenceAttachmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sentences_attachments_proto_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SentenceAttachmentResponse.ProtoReflect.Descriptor instead.
func (*SentenceAttachmentResponse) Descriptor() ([]byte, []int) {
	return file_sentences_attachments_proto_proto_rawDescGZIP(), []int{1}
}

func (x *SentenceAttachmentResponse) GetSentenceId() int64 {
	if x != nil {
		return x.SentenceId
	}
	return 0
}

func (x *SentenceAttachmentResponse) GetAttachmentId() int32 {
	if x != nil {
		return x.AttachmentId
	}
	return 0
}

type DeleteSentenceAttachmentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SentenceId    int64                  `protobuf:"varint,1,opt,name=sentence_id,json=sentenceId,proto3" json:"sentence_id,omitempty"`
	AttachmentId  int32                  `protobuf:"varint,2,opt,name=attachment_id,json=attachmentId,proto3" json:"attachment_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteSentenceAttachmentRequest) Reset() {
	*x = DeleteSentenceAttachmentRequest{}
	mi := &file_sentences_attachments_proto_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSentenceAttachmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSentenceAttachmentRequest) ProtoMessage() {}

func (x *DeleteSentenceAttachmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sentences_attachments_proto_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSentenceAttachmentRequest.ProtoReflect.Descriptor instead.
func (*DeleteSentenceAttachmentRequest) Descriptor() ([]byte, []int) {
	return file_sentences_attachments_proto_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteSentenceAttachmentRequest) GetSentenceId() int64 {
	if x != nil {
		return x.SentenceId
	}
	return 0
}

func (x *DeleteSentenceAttachmentRequest) GetAttachmentId() int32 {
	if x != nil {
		return x.AttachmentId
	}
	return 0
}

type DeleteSentenceAttachmentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteSentenceAttachmentResponse) Reset() {
	*x = DeleteSentenceAttachmentResponse{}
	mi := &file_sentences_attachments_proto_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSentenceAttachmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSentenceAttachmentResponse) ProtoMessage() {}

func (x *DeleteSentenceAttachmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sentences_attachments_proto_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSentenceAttachmentResponse.ProtoReflect.Descriptor instead.
func (*DeleteSentenceAttachmentResponse) Descriptor() ([]byte, []int) {
	return file_sentences_attachments_proto_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteSentenceAttachmentResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeleteSentenceAttachmentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListBySentenceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SentenceId    int64                  `protobuf:"varint,1,opt,name=sentence_id,json=sentenceId,proto3" json:"sentence_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListBySentenceRequest) Reset() {
	*x = ListBySentenceRequest{}
	mi := &file_sentences_attachments_proto_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBySentenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBySentenceRequest) ProtoMessage() {}

func (x *ListBySentenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sentences_attachments_proto_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBySentenceRequest.ProtoReflect.Descriptor instead.
func (*ListBySentenceRequest) Descriptor() ([]byte, []int) {
	return file_sentences_attachments_proto_proto_rawDescGZIP(), []int{4}
}

func (x *ListBySentenceRequest) GetSentenceId() int64 {
	if x != nil {
		return x.SentenceId
	}
	return 0
}

type ListByAttachmentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AttachmentId  int32                  `protobuf:"varint,1,opt,name=attachment_id,json=attachmentId,proto3" json:"attachment_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListByAttachmentRequest) Reset() {
	*x = ListByAttachmentRequest{}
	mi := &file_sentences_attachments_proto_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListByAttachmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListByAttachmentRequest) ProtoMessage() {}

func (x *ListByAttachmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sentences_attachments_proto_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListByAttachmentRequest.ProtoReflect.Descriptor instead.
func (*ListByAttachmentRequest) Descriptor() ([]byte, []int) {
	return file_sentences_attachments_proto_proto_rawDescGZIP(), []int{5}
}

func (x *ListByAttachmentRequest) GetAttachmentId() int32 {
	if x != nil {
		return x.AttachmentId
	}
	return 0
}

type SentenceAttachmentsListResponse struct {
	state         protoimpl.MessageState        `protogen:"open.v1"`
	Data          []*SentenceAttachmentResponse `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SentenceAttachmentsListResponse) Reset() {
	*x = SentenceAttachmentsListResponse{}
	mi := &file_sentences_attachments_proto_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SentenceAttachmentsListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SentenceAttachmentsListResponse) ProtoMessage() {}

func (x *SentenceAttachmentsListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sentences_attachments_proto_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SentenceAttachmentsListResponse.ProtoReflect.Descriptor instead.
func (*SentenceAttachmentsListResponse) Descriptor() ([]byte, []int) {
	return file_sentences_attachments_proto_proto_rawDescGZIP(), []int{6}
}

func (x *SentenceAttachmentsListResponse) GetData() []*SentenceAttachmentResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_sentences_attachments_proto_proto protoreflect.FileDescriptor

var file_sentences_attachments_proto_proto_rawDesc = string([]byte{
	0x0a, 0x21, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x5f, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x22, 0x67, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e,
	0x74, 0x65, 0x6e, 0x63, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x65, 0x6e,
	0x74, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x62, 0x0a, 0x1a,
	0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65,
	0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x22, 0x67, 0x0a, 0x1f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e,
	0x63, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x20, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x38, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x53, 0x65, 0x6e, 0x74, 0x65,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65,
	0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x17, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x79, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x71, 0x0a, 0x1f, 0x53,
	0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x73,
	0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x65,
	0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xce,
	0x04, 0x0a, 0x1c, 0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12,
	0x85, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x3f, 0x2e, 0x73, 0x65, 0x6e,
	0x74, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x73, 0x65,
	0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x6e,
	0x74, 0x65, 0x6e, 0x63, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x8b, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x3f, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x5f, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e,
	0x63, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x40, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x5f,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x74, 0x65,
	0x6e, 0x63, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79,
	0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x35, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79,
	0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x3f, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x8c, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x37, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f,
	0x2e, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x34, 0x5a, 0x32, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x70, 0x70, 0x75, 0x2e, 0x76, 0x31, 0x3b,
	0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_sentences_attachments_proto_proto_rawDescOnce sync.Once
	file_sentences_attachments_proto_proto_rawDescData []byte
)

func file_sentences_attachments_proto_proto_rawDescGZIP() []byte {
	file_sentences_attachments_proto_proto_rawDescOnce.Do(func() {
		file_sentences_attachments_proto_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_sentences_attachments_proto_proto_rawDesc), len(file_sentences_attachments_proto_proto_rawDesc)))
	})
	return file_sentences_attachments_proto_proto_rawDescData
}

var file_sentences_attachments_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_sentences_attachments_proto_proto_goTypes = []any{
	(*CreateSentenceAttachmentRequest)(nil),  // 0: sentences_attachments_provider.CreateSentenceAttachmentRequest
	(*SentenceAttachmentResponse)(nil),       // 1: sentences_attachments_provider.SentenceAttachmentResponse
	(*DeleteSentenceAttachmentRequest)(nil),  // 2: sentences_attachments_provider.DeleteSentenceAttachmentRequest
	(*DeleteSentenceAttachmentResponse)(nil), // 3: sentences_attachments_provider.DeleteSentenceAttachmentResponse
	(*ListBySentenceRequest)(nil),            // 4: sentences_attachments_provider.ListBySentenceRequest
	(*ListByAttachmentRequest)(nil),          // 5: sentences_attachments_provider.ListByAttachmentRequest
	(*SentenceAttachmentsListResponse)(nil),  // 6: sentences_attachments_provider.SentenceAttachmentsListResponse
}
var file_sentences_attachments_proto_proto_depIdxs = []int32{
	1, // 0: sentences_attachments_provider.SentenceAttachmentsListResponse.data:type_name -> sentences_attachments_provider.SentenceAttachmentResponse
	0, // 1: sentences_attachments_provider.SentencesAttachmentsProvider.Create:input_type -> sentences_attachments_provider.CreateSentenceAttachmentRequest
	2, // 2: sentences_attachments_provider.SentencesAttachmentsProvider.Delete:input_type -> sentences_attachments_provider.DeleteSentenceAttachmentRequest
	4, // 3: sentences_attachments_provider.SentencesAttachmentsProvider.ListBySentence:input_type -> sentences_attachments_provider.ListBySentenceRequest
	5, // 4: sentences_attachments_provider.SentencesAttachmentsProvider.ListByAttachment:input_type -> sentences_attachments_provider.ListByAttachmentRequest
	1, // 5: sentences_attachments_provider.SentencesAttachmentsProvider.Create:output_type -> sentences_attachments_provider.SentenceAttachmentResponse
	3, // 6: sentences_attachments_provider.SentencesAttachmentsProvider.Delete:output_type -> sentences_attachments_provider.DeleteSentenceAttachmentResponse
	6, // 7: sentences_attachments_provider.SentencesAttachmentsProvider.ListBySentence:output_type -> sentences_attachments_provider.SentenceAttachmentsListResponse
	6, // 8: sentences_attachments_provider.SentencesAttachmentsProvider.ListByAttachment:output_type -> sentences_attachments_provider.SentenceAttachmentsListResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sentences_attachments_proto_proto_init() }
func file_sentences_attachments_proto_proto_init() {
	if File_sentences_attachments_proto_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_sentences_attachments_proto_proto_rawDesc), len(file_sentences_attachments_proto_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sentences_attachments_proto_proto_goTypes,
		DependencyIndexes: file_sentences_attachments_proto_proto_depIdxs,
		MessageInfos:      file_sentences_attachments_proto_proto_msgTypes,
	}.Build()
	File_sentences_attachments_proto_proto = out.File
	file_sentences_attachments_proto_proto_goTypes = nil
	file_sentences_attachments_proto_proto_depIdxs = nil
}
