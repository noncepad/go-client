// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: job.proto

package job

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WorkClient is the client API for Work service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkClient interface {
	SubmitJob(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*Job, error)
	UpdateJob(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Job, error)
}

type workClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkClient(cc grpc.ClientConnInterface) WorkClient {
	return &workClient{cc}
}

func (c *workClient) SubmitJob(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := c.cc.Invoke(ctx, "/job.Work/SubmitJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workClient) UpdateJob(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := c.cc.Invoke(ctx, "/job.Work/UpdateJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkServer is the server API for Work service.
// All implementations must embed UnimplementedWorkServer
// for forward compatibility
type WorkServer interface {
	SubmitJob(context.Context, *SubmitRequest) (*Job, error)
	UpdateJob(context.Context, *UpdateRequest) (*Job, error)
	mustEmbedUnimplementedWorkServer()
}

// UnimplementedWorkServer must be embedded to have forward compatible implementations.
type UnimplementedWorkServer struct {
}

func (UnimplementedWorkServer) SubmitJob(context.Context, *SubmitRequest) (*Job, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitJob not implemented")
}
func (UnimplementedWorkServer) UpdateJob(context.Context, *UpdateRequest) (*Job, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateJob not implemented")
}
func (UnimplementedWorkServer) mustEmbedUnimplementedWorkServer() {}

// UnsafeWorkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkServer will
// result in compilation errors.
type UnsafeWorkServer interface {
	mustEmbedUnimplementedWorkServer()
}

func RegisterWorkServer(s grpc.ServiceRegistrar, srv WorkServer) {
	s.RegisterService(&Work_ServiceDesc, srv)
}

func _Work_SubmitJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkServer).SubmitJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.Work/SubmitJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkServer).SubmitJob(ctx, req.(*SubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Work_UpdateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkServer).UpdateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.Work/UpdateJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkServer).UpdateJob(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Work_ServiceDesc is the grpc.ServiceDesc for Work service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Work_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "job.Work",
	HandlerType: (*WorkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitJob",
			Handler:    _Work_SubmitJob_Handler,
		},
		{
			MethodName: "UpdateJob",
			Handler:    _Work_UpdateJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job.proto",
}

// OwnerClient is the client API for Owner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OwnerClient interface {
	GetStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*Job, error)
	GetStatusStream(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (Owner_GetStatusStreamClient, error)
}

type ownerClient struct {
	cc grpc.ClientConnInterface
}

func NewOwnerClient(cc grpc.ClientConnInterface) OwnerClient {
	return &ownerClient{cc}
}

func (c *ownerClient) GetStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := c.cc.Invoke(ctx, "/job.Owner/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ownerClient) GetStatusStream(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (Owner_GetStatusStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Owner_ServiceDesc.Streams[0], "/job.Owner/GetStatusStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &ownerGetStatusStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Owner_GetStatusStreamClient interface {
	Recv() (*Job, error)
	grpc.ClientStream
}

type ownerGetStatusStreamClient struct {
	grpc.ClientStream
}

func (x *ownerGetStatusStreamClient) Recv() (*Job, error) {
	m := new(Job)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OwnerServer is the server API for Owner service.
// All implementations must embed UnimplementedOwnerServer
// for forward compatibility
type OwnerServer interface {
	GetStatus(context.Context, *StatusRequest) (*Job, error)
	GetStatusStream(*StatusRequest, Owner_GetStatusStreamServer) error
	mustEmbedUnimplementedOwnerServer()
}

// UnimplementedOwnerServer must be embedded to have forward compatible implementations.
type UnimplementedOwnerServer struct {
}

func (UnimplementedOwnerServer) GetStatus(context.Context, *StatusRequest) (*Job, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedOwnerServer) GetStatusStream(*StatusRequest, Owner_GetStatusStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStatusStream not implemented")
}
func (UnimplementedOwnerServer) mustEmbedUnimplementedOwnerServer() {}

// UnsafeOwnerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OwnerServer will
// result in compilation errors.
type UnsafeOwnerServer interface {
	mustEmbedUnimplementedOwnerServer()
}

func RegisterOwnerServer(s grpc.ServiceRegistrar, srv OwnerServer) {
	s.RegisterService(&Owner_ServiceDesc, srv)
}

func _Owner_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnerServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.Owner/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnerServer).GetStatus(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Owner_GetStatusStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StatusRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OwnerServer).GetStatusStream(m, &ownerGetStatusStreamServer{stream})
}

type Owner_GetStatusStreamServer interface {
	Send(*Job) error
	grpc.ServerStream
}

type ownerGetStatusStreamServer struct {
	grpc.ServerStream
}

func (x *ownerGetStatusStreamServer) Send(m *Job) error {
	return x.ServerStream.SendMsg(m)
}

// Owner_ServiceDesc is the grpc.ServiceDesc for Owner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Owner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "job.Owner",
	HandlerType: (*OwnerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _Owner_GetStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStatusStream",
			Handler:       _Owner_GetStatusStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "job.proto",
}