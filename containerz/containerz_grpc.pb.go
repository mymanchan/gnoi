// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.10
// source: containerz/containerz.proto

package containerz

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

// ContainerzClient is the client API for Containerz service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContainerzClient interface {
	Deploy(ctx context.Context, opts ...grpc.CallOption) (Containerz_DeployClient, error)
	RemoveContainer(ctx context.Context, in *RemoveContainerRequest, opts ...grpc.CallOption) (*RemoveContainerResponse, error)
	ListContainer(ctx context.Context, in *ListContainerRequest, opts ...grpc.CallOption) (Containerz_ListContainerClient, error)
	StartContainer(ctx context.Context, in *StartContainerRequest, opts ...grpc.CallOption) (*StartContainerResponse, error)
	StopContainer(ctx context.Context, in *StopContainerRequest, opts ...grpc.CallOption) (*StopContainerResponse, error)
	Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (Containerz_LogClient, error)
	CreateVolume(ctx context.Context, in *CreateVolumeRequest, opts ...grpc.CallOption) (*CreateVolumeResponse, error)
	RemoveVolume(ctx context.Context, in *RemoveVolumeRequest, opts ...grpc.CallOption) (*RemoveVolumeResponse, error)
	ListVolume(ctx context.Context, in *ListVolumeRequest, opts ...grpc.CallOption) (Containerz_ListVolumeClient, error)
}

type containerzClient struct {
	cc grpc.ClientConnInterface
}

func NewContainerzClient(cc grpc.ClientConnInterface) ContainerzClient {
	return &containerzClient{cc}
}

func (c *containerzClient) Deploy(ctx context.Context, opts ...grpc.CallOption) (Containerz_DeployClient, error) {
	stream, err := c.cc.NewStream(ctx, &Containerz_ServiceDesc.Streams[0], "/gnoi.containerz.Containerz/Deploy", opts...)
	if err != nil {
		return nil, err
	}
	x := &containerzDeployClient{stream}
	return x, nil
}

type Containerz_DeployClient interface {
	Send(*DeployRequest) error
	Recv() (*DeployResponse, error)
	grpc.ClientStream
}

type containerzDeployClient struct {
	grpc.ClientStream
}

func (x *containerzDeployClient) Send(m *DeployRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *containerzDeployClient) Recv() (*DeployResponse, error) {
	m := new(DeployResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *containerzClient) RemoveContainer(ctx context.Context, in *RemoveContainerRequest, opts ...grpc.CallOption) (*RemoveContainerResponse, error) {
	out := new(RemoveContainerResponse)
	err := c.cc.Invoke(ctx, "/gnoi.containerz.Containerz/RemoveContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerzClient) ListContainer(ctx context.Context, in *ListContainerRequest, opts ...grpc.CallOption) (Containerz_ListContainerClient, error) {
	stream, err := c.cc.NewStream(ctx, &Containerz_ServiceDesc.Streams[1], "/gnoi.containerz.Containerz/ListContainer", opts...)
	if err != nil {
		return nil, err
	}
	x := &containerzListContainerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Containerz_ListContainerClient interface {
	Recv() (*ListContainerResponse, error)
	grpc.ClientStream
}

type containerzListContainerClient struct {
	grpc.ClientStream
}

func (x *containerzListContainerClient) Recv() (*ListContainerResponse, error) {
	m := new(ListContainerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *containerzClient) StartContainer(ctx context.Context, in *StartContainerRequest, opts ...grpc.CallOption) (*StartContainerResponse, error) {
	out := new(StartContainerResponse)
	err := c.cc.Invoke(ctx, "/gnoi.containerz.Containerz/StartContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerzClient) StopContainer(ctx context.Context, in *StopContainerRequest, opts ...grpc.CallOption) (*StopContainerResponse, error) {
	out := new(StopContainerResponse)
	err := c.cc.Invoke(ctx, "/gnoi.containerz.Containerz/StopContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerzClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (Containerz_LogClient, error) {
	stream, err := c.cc.NewStream(ctx, &Containerz_ServiceDesc.Streams[2], "/gnoi.containerz.Containerz/Log", opts...)
	if err != nil {
		return nil, err
	}
	x := &containerzLogClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Containerz_LogClient interface {
	Recv() (*LogResponse, error)
	grpc.ClientStream
}

type containerzLogClient struct {
	grpc.ClientStream
}

func (x *containerzLogClient) Recv() (*LogResponse, error) {
	m := new(LogResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *containerzClient) CreateVolume(ctx context.Context, in *CreateVolumeRequest, opts ...grpc.CallOption) (*CreateVolumeResponse, error) {
	out := new(CreateVolumeResponse)
	err := c.cc.Invoke(ctx, "/gnoi.containerz.Containerz/CreateVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerzClient) RemoveVolume(ctx context.Context, in *RemoveVolumeRequest, opts ...grpc.CallOption) (*RemoveVolumeResponse, error) {
	out := new(RemoveVolumeResponse)
	err := c.cc.Invoke(ctx, "/gnoi.containerz.Containerz/RemoveVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerzClient) ListVolume(ctx context.Context, in *ListVolumeRequest, opts ...grpc.CallOption) (Containerz_ListVolumeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Containerz_ServiceDesc.Streams[3], "/gnoi.containerz.Containerz/ListVolume", opts...)
	if err != nil {
		return nil, err
	}
	x := &containerzListVolumeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Containerz_ListVolumeClient interface {
	Recv() (*ListVolumeResponse, error)
	grpc.ClientStream
}

type containerzListVolumeClient struct {
	grpc.ClientStream
}

func (x *containerzListVolumeClient) Recv() (*ListVolumeResponse, error) {
	m := new(ListVolumeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ContainerzServer is the server API for Containerz service.
// All implementations must embed UnimplementedContainerzServer
// for forward compatibility
type ContainerzServer interface {
	Deploy(Containerz_DeployServer) error
	RemoveContainer(context.Context, *RemoveContainerRequest) (*RemoveContainerResponse, error)
	ListContainer(*ListContainerRequest, Containerz_ListContainerServer) error
	StartContainer(context.Context, *StartContainerRequest) (*StartContainerResponse, error)
	StopContainer(context.Context, *StopContainerRequest) (*StopContainerResponse, error)
	Log(*LogRequest, Containerz_LogServer) error
	CreateVolume(context.Context, *CreateVolumeRequest) (*CreateVolumeResponse, error)
	RemoveVolume(context.Context, *RemoveVolumeRequest) (*RemoveVolumeResponse, error)
	ListVolume(*ListVolumeRequest, Containerz_ListVolumeServer) error
	mustEmbedUnimplementedContainerzServer()
}

// UnimplementedContainerzServer must be embedded to have forward compatible implementations.
type UnimplementedContainerzServer struct {
}

func (UnimplementedContainerzServer) Deploy(Containerz_DeployServer) error {
	return status.Errorf(codes.Unimplemented, "method Deploy not implemented")
}
func (UnimplementedContainerzServer) RemoveContainer(context.Context, *RemoveContainerRequest) (*RemoveContainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveContainer not implemented")
}
func (UnimplementedContainerzServer) ListContainer(*ListContainerRequest, Containerz_ListContainerServer) error {
	return status.Errorf(codes.Unimplemented, "method ListContainer not implemented")
}
func (UnimplementedContainerzServer) StartContainer(context.Context, *StartContainerRequest) (*StartContainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartContainer not implemented")
}
func (UnimplementedContainerzServer) StopContainer(context.Context, *StopContainerRequest) (*StopContainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopContainer not implemented")
}
func (UnimplementedContainerzServer) Log(*LogRequest, Containerz_LogServer) error {
	return status.Errorf(codes.Unimplemented, "method Log not implemented")
}
func (UnimplementedContainerzServer) CreateVolume(context.Context, *CreateVolumeRequest) (*CreateVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVolume not implemented")
}
func (UnimplementedContainerzServer) RemoveVolume(context.Context, *RemoveVolumeRequest) (*RemoveVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveVolume not implemented")
}
func (UnimplementedContainerzServer) ListVolume(*ListVolumeRequest, Containerz_ListVolumeServer) error {
	return status.Errorf(codes.Unimplemented, "method ListVolume not implemented")
}
func (UnimplementedContainerzServer) mustEmbedUnimplementedContainerzServer() {}

// UnsafeContainerzServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContainerzServer will
// result in compilation errors.
type UnsafeContainerzServer interface {
	mustEmbedUnimplementedContainerzServer()
}

func RegisterContainerzServer(s grpc.ServiceRegistrar, srv ContainerzServer) {
	s.RegisterService(&Containerz_ServiceDesc, srv)
}

func _Containerz_Deploy_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ContainerzServer).Deploy(&containerzDeployServer{stream})
}

type Containerz_DeployServer interface {
	Send(*DeployResponse) error
	Recv() (*DeployRequest, error)
	grpc.ServerStream
}

type containerzDeployServer struct {
	grpc.ServerStream
}

func (x *containerzDeployServer) Send(m *DeployResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *containerzDeployServer) Recv() (*DeployRequest, error) {
	m := new(DeployRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Containerz_RemoveContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerzServer).RemoveContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.containerz.Containerz/RemoveContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerzServer).RemoveContainer(ctx, req.(*RemoveContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Containerz_ListContainer_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListContainerRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ContainerzServer).ListContainer(m, &containerzListContainerServer{stream})
}

type Containerz_ListContainerServer interface {
	Send(*ListContainerResponse) error
	grpc.ServerStream
}

type containerzListContainerServer struct {
	grpc.ServerStream
}

func (x *containerzListContainerServer) Send(m *ListContainerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Containerz_StartContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerzServer).StartContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.containerz.Containerz/StartContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerzServer).StartContainer(ctx, req.(*StartContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Containerz_StopContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerzServer).StopContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.containerz.Containerz/StopContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerzServer).StopContainer(ctx, req.(*StopContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Containerz_Log_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ContainerzServer).Log(m, &containerzLogServer{stream})
}

type Containerz_LogServer interface {
	Send(*LogResponse) error
	grpc.ServerStream
}

type containerzLogServer struct {
	grpc.ServerStream
}

func (x *containerzLogServer) Send(m *LogResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Containerz_CreateVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerzServer).CreateVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.containerz.Containerz/CreateVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerzServer).CreateVolume(ctx, req.(*CreateVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Containerz_RemoveVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerzServer).RemoveVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.containerz.Containerz/RemoveVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerzServer).RemoveVolume(ctx, req.(*RemoveVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Containerz_ListVolume_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListVolumeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ContainerzServer).ListVolume(m, &containerzListVolumeServer{stream})
}

type Containerz_ListVolumeServer interface {
	Send(*ListVolumeResponse) error
	grpc.ServerStream
}

type containerzListVolumeServer struct {
	grpc.ServerStream
}

func (x *containerzListVolumeServer) Send(m *ListVolumeResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Containerz_ServiceDesc is the grpc.ServiceDesc for Containerz service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Containerz_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnoi.containerz.Containerz",
	HandlerType: (*ContainerzServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RemoveContainer",
			Handler:    _Containerz_RemoveContainer_Handler,
		},
		{
			MethodName: "StartContainer",
			Handler:    _Containerz_StartContainer_Handler,
		},
		{
			MethodName: "StopContainer",
			Handler:    _Containerz_StopContainer_Handler,
		},
		{
			MethodName: "CreateVolume",
			Handler:    _Containerz_CreateVolume_Handler,
		},
		{
			MethodName: "RemoveVolume",
			Handler:    _Containerz_RemoveVolume_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Deploy",
			Handler:       _Containerz_Deploy_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ListContainer",
			Handler:       _Containerz_ListContainer_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Log",
			Handler:       _Containerz_Log_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListVolume",
			Handler:       _Containerz_ListVolume_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "containerz/containerz.proto",
}
