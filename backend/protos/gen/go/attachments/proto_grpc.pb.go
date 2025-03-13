// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: attachments/proto.proto

package attachments

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AttachmentsProvider_Create_FullMethodName = "/attachments_provider.AttachmentsProvider/Create"
	AttachmentsProvider_Get_FullMethodName    = "/attachments_provider.AttachmentsProvider/Get"
	AttachmentsProvider_Update_FullMethodName = "/attachments_provider.AttachmentsProvider/Update"
	AttachmentsProvider_Delete_FullMethodName = "/attachments_provider.AttachmentsProvider/Delete"
	AttachmentsProvider_List_FullMethodName   = "/attachments_provider.AttachmentsProvider/List"
)

// AttachmentsProviderClient is the client API for AttachmentsProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AttachmentsProviderClient interface {
	Create(ctx context.Context, in *CreateAttachmentRequest, opts ...grpc.CallOption) (*AttachmentResponse, error)
	Get(ctx context.Context, in *GetAttachmentRequest, opts ...grpc.CallOption) (*AttachmentResponse, error)
	Update(ctx context.Context, in *UpdateAttachmentRequest, opts ...grpc.CallOption) (*AttachmentResponse, error)
	Delete(ctx context.Context, in *DeleteAttachmentRequest, opts ...grpc.CallOption) (*DeleteAttachmentResponse, error)
	List(ctx context.Context, in *ListAttachmentsRequest, opts ...grpc.CallOption) (*AttachmentsListResponse, error)
}

type attachmentsProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewAttachmentsProviderClient(cc grpc.ClientConnInterface) AttachmentsProviderClient {
	return &attachmentsProviderClient{cc}
}

func (c *attachmentsProviderClient) Create(ctx context.Context, in *CreateAttachmentRequest, opts ...grpc.CallOption) (*AttachmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AttachmentResponse)
	err := c.cc.Invoke(ctx, AttachmentsProvider_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentsProviderClient) Get(ctx context.Context, in *GetAttachmentRequest, opts ...grpc.CallOption) (*AttachmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AttachmentResponse)
	err := c.cc.Invoke(ctx, AttachmentsProvider_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentsProviderClient) Update(ctx context.Context, in *UpdateAttachmentRequest, opts ...grpc.CallOption) (*AttachmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AttachmentResponse)
	err := c.cc.Invoke(ctx, AttachmentsProvider_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentsProviderClient) Delete(ctx context.Context, in *DeleteAttachmentRequest, opts ...grpc.CallOption) (*DeleteAttachmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteAttachmentResponse)
	err := c.cc.Invoke(ctx, AttachmentsProvider_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentsProviderClient) List(ctx context.Context, in *ListAttachmentsRequest, opts ...grpc.CallOption) (*AttachmentsListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AttachmentsListResponse)
	err := c.cc.Invoke(ctx, AttachmentsProvider_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttachmentsProviderServer is the server API for AttachmentsProvider service.
// All implementations must embed UnimplementedAttachmentsProviderServer
// for forward compatibility.
type AttachmentsProviderServer interface {
	Create(context.Context, *CreateAttachmentRequest) (*AttachmentResponse, error)
	Get(context.Context, *GetAttachmentRequest) (*AttachmentResponse, error)
	Update(context.Context, *UpdateAttachmentRequest) (*AttachmentResponse, error)
	Delete(context.Context, *DeleteAttachmentRequest) (*DeleteAttachmentResponse, error)
	List(context.Context, *ListAttachmentsRequest) (*AttachmentsListResponse, error)
	mustEmbedUnimplementedAttachmentsProviderServer()
}

// UnimplementedAttachmentsProviderServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAttachmentsProviderServer struct{}

func (UnimplementedAttachmentsProviderServer) Create(context.Context, *CreateAttachmentRequest) (*AttachmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAttachmentsProviderServer) Get(context.Context, *GetAttachmentRequest) (*AttachmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAttachmentsProviderServer) Update(context.Context, *UpdateAttachmentRequest) (*AttachmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAttachmentsProviderServer) Delete(context.Context, *DeleteAttachmentRequest) (*DeleteAttachmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAttachmentsProviderServer) List(context.Context, *ListAttachmentsRequest) (*AttachmentsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAttachmentsProviderServer) mustEmbedUnimplementedAttachmentsProviderServer() {}
func (UnimplementedAttachmentsProviderServer) testEmbeddedByValue()                             {}

// UnsafeAttachmentsProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AttachmentsProviderServer will
// result in compilation errors.
type UnsafeAttachmentsProviderServer interface {
	mustEmbedUnimplementedAttachmentsProviderServer()
}

func RegisterAttachmentsProviderServer(s grpc.ServiceRegistrar, srv AttachmentsProviderServer) {
	// If the following call pancis, it indicates UnimplementedAttachmentsProviderServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AttachmentsProvider_ServiceDesc, srv)
}

func _AttachmentsProvider_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAttachmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentsProviderServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttachmentsProvider_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentsProviderServer).Create(ctx, req.(*CreateAttachmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttachmentsProvider_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttachmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentsProviderServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttachmentsProvider_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentsProviderServer).Get(ctx, req.(*GetAttachmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttachmentsProvider_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAttachmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentsProviderServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttachmentsProvider_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentsProviderServer).Update(ctx, req.(*UpdateAttachmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttachmentsProvider_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAttachmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentsProviderServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttachmentsProvider_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentsProviderServer).Delete(ctx, req.(*DeleteAttachmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttachmentsProvider_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAttachmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentsProviderServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttachmentsProvider_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentsProviderServer).List(ctx, req.(*ListAttachmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AttachmentsProvider_ServiceDesc is the grpc.ServiceDesc for AttachmentsProvider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AttachmentsProvider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "attachments_provider.AttachmentsProvider",
	HandlerType: (*AttachmentsProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AttachmentsProvider_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AttachmentsProvider_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AttachmentsProvider_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AttachmentsProvider_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _AttachmentsProvider_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "attachments/proto.proto",
}
