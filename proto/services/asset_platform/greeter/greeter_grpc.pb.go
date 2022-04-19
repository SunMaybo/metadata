// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: greeter.proto

package greeter

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

// GreeterServiceClient is the client API for GreeterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterServiceClient interface {
	SayHelloWord(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SayStream(ctx context.Context, opts ...grpc.CallOption) (GreeterService_SayStreamClient, error)
	SayStream1(ctx context.Context, opts ...grpc.CallOption) (GreeterService_SayStream1Client, error)
	SayStream2(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (GreeterService_SayStream2Client, error)
}

type greeterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterServiceClient(cc grpc.ClientConnInterface) GreeterServiceClient {
	return &greeterServiceClient{cc}
}

func (c *greeterServiceClient) SayHelloWord(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/greeter.GreeterService/SayHelloWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterServiceClient) SayStream(ctx context.Context, opts ...grpc.CallOption) (GreeterService_SayStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreeterService_ServiceDesc.Streams[0], "/greeter.GreeterService/SayStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterServiceSayStreamClient{stream}
	return x, nil
}

type GreeterService_SayStreamClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterServiceSayStreamClient struct {
	grpc.ClientStream
}

func (x *greeterServiceSayStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterServiceSayStreamClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterServiceClient) SayStream1(ctx context.Context, opts ...grpc.CallOption) (GreeterService_SayStream1Client, error) {
	stream, err := c.cc.NewStream(ctx, &GreeterService_ServiceDesc.Streams[1], "/greeter.GreeterService/SayStream1", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterServiceSayStream1Client{stream}
	return x, nil
}

type GreeterService_SayStream1Client interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterServiceSayStream1Client struct {
	grpc.ClientStream
}

func (x *greeterServiceSayStream1Client) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterServiceSayStream1Client) CloseAndRecv() (*HelloReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterServiceClient) SayStream2(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (GreeterService_SayStream2Client, error) {
	stream, err := c.cc.NewStream(ctx, &GreeterService_ServiceDesc.Streams[2], "/greeter.GreeterService/SayStream2", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterServiceSayStream2Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreeterService_SayStream2Client interface {
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterServiceSayStream2Client struct {
	grpc.ClientStream
}

func (x *greeterServiceSayStream2Client) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServiceServer is the server API for GreeterService service.
// All implementations must embed UnimplementedGreeterServiceServer
// for forward compatibility
type GreeterServiceServer interface {
	SayHelloWord(context.Context, *HelloRequest) (*HelloReply, error)
	SayStream(GreeterService_SayStreamServer) error
	SayStream1(GreeterService_SayStream1Server) error
	SayStream2(*HelloRequest, GreeterService_SayStream2Server) error
	mustEmbedUnimplementedGreeterServiceServer()
}

// UnimplementedGreeterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServiceServer struct {
}

func (UnimplementedGreeterServiceServer) SayHelloWord(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHelloWord not implemented")
}
func (UnimplementedGreeterServiceServer) SayStream(GreeterService_SayStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayStream not implemented")
}
func (UnimplementedGreeterServiceServer) SayStream1(GreeterService_SayStream1Server) error {
	return status.Errorf(codes.Unimplemented, "method SayStream1 not implemented")
}
func (UnimplementedGreeterServiceServer) SayStream2(*HelloRequest, GreeterService_SayStream2Server) error {
	return status.Errorf(codes.Unimplemented, "method SayStream2 not implemented")
}
func (UnimplementedGreeterServiceServer) mustEmbedUnimplementedGreeterServiceServer() {}

// UnsafeGreeterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServiceServer will
// result in compilation errors.
type UnsafeGreeterServiceServer interface {
	mustEmbedUnimplementedGreeterServiceServer()
}

func RegisterGreeterServiceServer(s grpc.ServiceRegistrar, srv GreeterServiceServer) {
	s.RegisterService(&GreeterService_ServiceDesc, srv)
}

func _GreeterService_SayHelloWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServiceServer).SayHelloWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greeter.GreeterService/SayHelloWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServiceServer).SayHelloWord(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreeterService_SayStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServiceServer).SayStream(&greeterServiceSayStreamServer{stream})
}

type GreeterService_SayStreamServer interface {
	Send(*HelloReply) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterServiceSayStreamServer struct {
	grpc.ServerStream
}

func (x *greeterServiceSayStreamServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterServiceSayStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreeterService_SayStream1_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServiceServer).SayStream1(&greeterServiceSayStream1Server{stream})
}

type GreeterService_SayStream1Server interface {
	SendAndClose(*HelloReply) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterServiceSayStream1Server struct {
	grpc.ServerStream
}

func (x *greeterServiceSayStream1Server) SendAndClose(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterServiceSayStream1Server) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreeterService_SayStream2_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServiceServer).SayStream2(m, &greeterServiceSayStream2Server{stream})
}

type GreeterService_SayStream2Server interface {
	Send(*HelloReply) error
	grpc.ServerStream
}

type greeterServiceSayStream2Server struct {
	grpc.ServerStream
}

func (x *greeterServiceSayStream2Server) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

// GreeterService_ServiceDesc is the grpc.ServiceDesc for GreeterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreeterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greeter.GreeterService",
	HandlerType: (*GreeterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHelloWord",
			Handler:    _GreeterService_SayHelloWord_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayStream",
			Handler:       _GreeterService_SayStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SayStream1",
			Handler:       _GreeterService_SayStream1_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SayStream2",
			Handler:       _GreeterService_SayStream2_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "greeter.proto",
}