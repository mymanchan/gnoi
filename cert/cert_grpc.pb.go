// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.10
// source: cert/cert.proto

package cert

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

// CertificateManagementClient is the client API for CertificateManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CertificateManagementClient interface {
	Rotate(ctx context.Context, opts ...grpc.CallOption) (CertificateManagement_RotateClient, error)
	Install(ctx context.Context, opts ...grpc.CallOption) (CertificateManagement_InstallClient, error)
	GenerateCSR(ctx context.Context, in *GenerateCSRRequest, opts ...grpc.CallOption) (*GenerateCSRResponse, error)
	LoadCertificate(ctx context.Context, in *LoadCertificateRequest, opts ...grpc.CallOption) (*LoadCertificateResponse, error)
	LoadCertificateAuthorityBundle(ctx context.Context, in *LoadCertificateAuthorityBundleRequest, opts ...grpc.CallOption) (*LoadCertificateAuthorityBundleResponse, error)
	GetCertificates(ctx context.Context, in *GetCertificatesRequest, opts ...grpc.CallOption) (*GetCertificatesResponse, error)
	RevokeCertificates(ctx context.Context, in *RevokeCertificatesRequest, opts ...grpc.CallOption) (*RevokeCertificatesResponse, error)
	CanGenerateCSR(ctx context.Context, in *CanGenerateCSRRequest, opts ...grpc.CallOption) (*CanGenerateCSRResponse, error)
}

type certificateManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewCertificateManagementClient(cc grpc.ClientConnInterface) CertificateManagementClient {
	return &certificateManagementClient{cc}
}

func (c *certificateManagementClient) Rotate(ctx context.Context, opts ...grpc.CallOption) (CertificateManagement_RotateClient, error) {
	stream, err := c.cc.NewStream(ctx, &CertificateManagement_ServiceDesc.Streams[0], "/gnoi.certificate.CertificateManagement/Rotate", opts...)
	if err != nil {
		return nil, err
	}
	x := &certificateManagementRotateClient{stream}
	return x, nil
}

type CertificateManagement_RotateClient interface {
	Send(*RotateCertificateRequest) error
	Recv() (*RotateCertificateResponse, error)
	grpc.ClientStream
}

type certificateManagementRotateClient struct {
	grpc.ClientStream
}

func (x *certificateManagementRotateClient) Send(m *RotateCertificateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *certificateManagementRotateClient) Recv() (*RotateCertificateResponse, error) {
	m := new(RotateCertificateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *certificateManagementClient) Install(ctx context.Context, opts ...grpc.CallOption) (CertificateManagement_InstallClient, error) {
	stream, err := c.cc.NewStream(ctx, &CertificateManagement_ServiceDesc.Streams[1], "/gnoi.certificate.CertificateManagement/Install", opts...)
	if err != nil {
		return nil, err
	}
	x := &certificateManagementInstallClient{stream}
	return x, nil
}

type CertificateManagement_InstallClient interface {
	Send(*InstallCertificateRequest) error
	Recv() (*InstallCertificateResponse, error)
	grpc.ClientStream
}

type certificateManagementInstallClient struct {
	grpc.ClientStream
}

func (x *certificateManagementInstallClient) Send(m *InstallCertificateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *certificateManagementInstallClient) Recv() (*InstallCertificateResponse, error) {
	m := new(InstallCertificateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *certificateManagementClient) GenerateCSR(ctx context.Context, in *GenerateCSRRequest, opts ...grpc.CallOption) (*GenerateCSRResponse, error) {
	out := new(GenerateCSRResponse)
	err := c.cc.Invoke(ctx, "/gnoi.certificate.CertificateManagement/GenerateCSR", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateManagementClient) LoadCertificate(ctx context.Context, in *LoadCertificateRequest, opts ...grpc.CallOption) (*LoadCertificateResponse, error) {
	out := new(LoadCertificateResponse)
	err := c.cc.Invoke(ctx, "/gnoi.certificate.CertificateManagement/LoadCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateManagementClient) LoadCertificateAuthorityBundle(ctx context.Context, in *LoadCertificateAuthorityBundleRequest, opts ...grpc.CallOption) (*LoadCertificateAuthorityBundleResponse, error) {
	out := new(LoadCertificateAuthorityBundleResponse)
	err := c.cc.Invoke(ctx, "/gnoi.certificate.CertificateManagement/LoadCertificateAuthorityBundle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateManagementClient) GetCertificates(ctx context.Context, in *GetCertificatesRequest, opts ...grpc.CallOption) (*GetCertificatesResponse, error) {
	out := new(GetCertificatesResponse)
	err := c.cc.Invoke(ctx, "/gnoi.certificate.CertificateManagement/GetCertificates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateManagementClient) RevokeCertificates(ctx context.Context, in *RevokeCertificatesRequest, opts ...grpc.CallOption) (*RevokeCertificatesResponse, error) {
	out := new(RevokeCertificatesResponse)
	err := c.cc.Invoke(ctx, "/gnoi.certificate.CertificateManagement/RevokeCertificates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateManagementClient) CanGenerateCSR(ctx context.Context, in *CanGenerateCSRRequest, opts ...grpc.CallOption) (*CanGenerateCSRResponse, error) {
	out := new(CanGenerateCSRResponse)
	err := c.cc.Invoke(ctx, "/gnoi.certificate.CertificateManagement/CanGenerateCSR", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateManagementServer is the server API for CertificateManagement service.
// All implementations must embed UnimplementedCertificateManagementServer
// for forward compatibility
type CertificateManagementServer interface {
	Rotate(CertificateManagement_RotateServer) error
	Install(CertificateManagement_InstallServer) error
	GenerateCSR(context.Context, *GenerateCSRRequest) (*GenerateCSRResponse, error)
	LoadCertificate(context.Context, *LoadCertificateRequest) (*LoadCertificateResponse, error)
	LoadCertificateAuthorityBundle(context.Context, *LoadCertificateAuthorityBundleRequest) (*LoadCertificateAuthorityBundleResponse, error)
	GetCertificates(context.Context, *GetCertificatesRequest) (*GetCertificatesResponse, error)
	RevokeCertificates(context.Context, *RevokeCertificatesRequest) (*RevokeCertificatesResponse, error)
	CanGenerateCSR(context.Context, *CanGenerateCSRRequest) (*CanGenerateCSRResponse, error)
	mustEmbedUnimplementedCertificateManagementServer()
}

// UnimplementedCertificateManagementServer must be embedded to have forward compatible implementations.
type UnimplementedCertificateManagementServer struct {
}

func (UnimplementedCertificateManagementServer) Rotate(CertificateManagement_RotateServer) error {
	return status.Errorf(codes.Unimplemented, "method Rotate not implemented")
}
func (UnimplementedCertificateManagementServer) Install(CertificateManagement_InstallServer) error {
	return status.Errorf(codes.Unimplemented, "method Install not implemented")
}
func (UnimplementedCertificateManagementServer) GenerateCSR(context.Context, *GenerateCSRRequest) (*GenerateCSRResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateCSR not implemented")
}
func (UnimplementedCertificateManagementServer) LoadCertificate(context.Context, *LoadCertificateRequest) (*LoadCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadCertificate not implemented")
}
func (UnimplementedCertificateManagementServer) LoadCertificateAuthorityBundle(context.Context, *LoadCertificateAuthorityBundleRequest) (*LoadCertificateAuthorityBundleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadCertificateAuthorityBundle not implemented")
}
func (UnimplementedCertificateManagementServer) GetCertificates(context.Context, *GetCertificatesRequest) (*GetCertificatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCertificates not implemented")
}
func (UnimplementedCertificateManagementServer) RevokeCertificates(context.Context, *RevokeCertificatesRequest) (*RevokeCertificatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeCertificates not implemented")
}
func (UnimplementedCertificateManagementServer) CanGenerateCSR(context.Context, *CanGenerateCSRRequest) (*CanGenerateCSRResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CanGenerateCSR not implemented")
}
func (UnimplementedCertificateManagementServer) mustEmbedUnimplementedCertificateManagementServer() {}

// UnsafeCertificateManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CertificateManagementServer will
// result in compilation errors.
type UnsafeCertificateManagementServer interface {
	mustEmbedUnimplementedCertificateManagementServer()
}

func RegisterCertificateManagementServer(s grpc.ServiceRegistrar, srv CertificateManagementServer) {
	s.RegisterService(&CertificateManagement_ServiceDesc, srv)
}

func _CertificateManagement_Rotate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CertificateManagementServer).Rotate(&certificateManagementRotateServer{stream})
}

type CertificateManagement_RotateServer interface {
	Send(*RotateCertificateResponse) error
	Recv() (*RotateCertificateRequest, error)
	grpc.ServerStream
}

type certificateManagementRotateServer struct {
	grpc.ServerStream
}

func (x *certificateManagementRotateServer) Send(m *RotateCertificateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *certificateManagementRotateServer) Recv() (*RotateCertificateRequest, error) {
	m := new(RotateCertificateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CertificateManagement_Install_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CertificateManagementServer).Install(&certificateManagementInstallServer{stream})
}

type CertificateManagement_InstallServer interface {
	Send(*InstallCertificateResponse) error
	Recv() (*InstallCertificateRequest, error)
	grpc.ServerStream
}

type certificateManagementInstallServer struct {
	grpc.ServerStream
}

func (x *certificateManagementInstallServer) Send(m *InstallCertificateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *certificateManagementInstallServer) Recv() (*InstallCertificateRequest, error) {
	m := new(InstallCertificateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CertificateManagement_GenerateCSR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateCSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateManagementServer).GenerateCSR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.certificate.CertificateManagement/GenerateCSR",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateManagementServer).GenerateCSR(ctx, req.(*GenerateCSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateManagement_LoadCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateManagementServer).LoadCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.certificate.CertificateManagement/LoadCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateManagementServer).LoadCertificate(ctx, req.(*LoadCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateManagement_LoadCertificateAuthorityBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadCertificateAuthorityBundleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateManagementServer).LoadCertificateAuthorityBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.certificate.CertificateManagement/LoadCertificateAuthorityBundle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateManagementServer).LoadCertificateAuthorityBundle(ctx, req.(*LoadCertificateAuthorityBundleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateManagement_GetCertificates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCertificatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateManagementServer).GetCertificates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.certificate.CertificateManagement/GetCertificates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateManagementServer).GetCertificates(ctx, req.(*GetCertificatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateManagement_RevokeCertificates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeCertificatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateManagementServer).RevokeCertificates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.certificate.CertificateManagement/RevokeCertificates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateManagementServer).RevokeCertificates(ctx, req.(*RevokeCertificatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateManagement_CanGenerateCSR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CanGenerateCSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateManagementServer).CanGenerateCSR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.certificate.CertificateManagement/CanGenerateCSR",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateManagementServer).CanGenerateCSR(ctx, req.(*CanGenerateCSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CertificateManagement_ServiceDesc is the grpc.ServiceDesc for CertificateManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CertificateManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnoi.certificate.CertificateManagement",
	HandlerType: (*CertificateManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateCSR",
			Handler:    _CertificateManagement_GenerateCSR_Handler,
		},
		{
			MethodName: "LoadCertificate",
			Handler:    _CertificateManagement_LoadCertificate_Handler,
		},
		{
			MethodName: "LoadCertificateAuthorityBundle",
			Handler:    _CertificateManagement_LoadCertificateAuthorityBundle_Handler,
		},
		{
			MethodName: "GetCertificates",
			Handler:    _CertificateManagement_GetCertificates_Handler,
		},
		{
			MethodName: "RevokeCertificates",
			Handler:    _CertificateManagement_RevokeCertificates_Handler,
		},
		{
			MethodName: "CanGenerateCSR",
			Handler:    _CertificateManagement_CanGenerateCSR_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Rotate",
			Handler:       _CertificateManagement_Rotate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Install",
			Handler:       _CertificateManagement_Install_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cert/cert.proto",
}