// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: kubearmor.proto

package protobuf

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

const (
	LogService_HealthCheck_FullMethodName   = "/feeder.LogService/HealthCheck"
	LogService_WatchMessages_FullMethodName = "/feeder.LogService/WatchMessages"
	LogService_WatchAlerts_FullMethodName   = "/feeder.LogService/WatchAlerts"
	LogService_WatchLogs_FullMethodName     = "/feeder.LogService/WatchLogs"
)

// LogServiceClient is the client API for LogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogServiceClient interface {
	// DEPRECATED: use "google.golang.org/grpc/health/grpc_health_v1"
	HealthCheck(ctx context.Context, in *NonceMessage, opts ...grpc.CallOption) (*ReplyMessage, error)
	WatchMessages(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (LogService_WatchMessagesClient, error)
	WatchAlerts(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (LogService_WatchAlertsClient, error)
	WatchLogs(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (LogService_WatchLogsClient, error)
}

type logServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogServiceClient(cc grpc.ClientConnInterface) LogServiceClient {
	return &logServiceClient{cc}
}

func (c *logServiceClient) HealthCheck(ctx context.Context, in *NonceMessage, opts ...grpc.CallOption) (*ReplyMessage, error) {
	out := new(ReplyMessage)
	err := c.cc.Invoke(ctx, LogService_HealthCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) WatchMessages(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (LogService_WatchMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &LogService_ServiceDesc.Streams[0], LogService_WatchMessages_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &logServiceWatchMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogService_WatchMessagesClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type logServiceWatchMessagesClient struct {
	grpc.ClientStream
}

func (x *logServiceWatchMessagesClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logServiceClient) WatchAlerts(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (LogService_WatchAlertsClient, error) {
	stream, err := c.cc.NewStream(ctx, &LogService_ServiceDesc.Streams[1], LogService_WatchAlerts_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &logServiceWatchAlertsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogService_WatchAlertsClient interface {
	Recv() (*Alert, error)
	grpc.ClientStream
}

type logServiceWatchAlertsClient struct {
	grpc.ClientStream
}

func (x *logServiceWatchAlertsClient) Recv() (*Alert, error) {
	m := new(Alert)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logServiceClient) WatchLogs(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (LogService_WatchLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &LogService_ServiceDesc.Streams[2], LogService_WatchLogs_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &logServiceWatchLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogService_WatchLogsClient interface {
	Recv() (*Log, error)
	grpc.ClientStream
}

type logServiceWatchLogsClient struct {
	grpc.ClientStream
}

func (x *logServiceWatchLogsClient) Recv() (*Log, error) {
	m := new(Log)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogServiceServer is the server API for LogService service.
// All implementations should embed UnimplementedLogServiceServer
// for forward compatibility
type LogServiceServer interface {
	// DEPRECATED: use "google.golang.org/grpc/health/grpc_health_v1"
	HealthCheck(context.Context, *NonceMessage) (*ReplyMessage, error)
	WatchMessages(*RequestMessage, LogService_WatchMessagesServer) error
	WatchAlerts(*RequestMessage, LogService_WatchAlertsServer) error
	WatchLogs(*RequestMessage, LogService_WatchLogsServer) error
}

// UnimplementedLogServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLogServiceServer struct {
}

func (UnimplementedLogServiceServer) HealthCheck(context.Context, *NonceMessage) (*ReplyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedLogServiceServer) WatchMessages(*RequestMessage, LogService_WatchMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchMessages not implemented")
}
func (UnimplementedLogServiceServer) WatchAlerts(*RequestMessage, LogService_WatchAlertsServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchAlerts not implemented")
}
func (UnimplementedLogServiceServer) WatchLogs(*RequestMessage, LogService_WatchLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchLogs not implemented")
}

// UnsafeLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogServiceServer will
// result in compilation errors.
type UnsafeLogServiceServer interface {
	mustEmbedUnimplementedLogServiceServer()
}

func RegisterLogServiceServer(s grpc.ServiceRegistrar, srv LogServiceServer) {
	s.RegisterService(&LogService_ServiceDesc, srv)
}

func _LogService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NonceMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_HealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).HealthCheck(ctx, req.(*NonceMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_WatchMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogServiceServer).WatchMessages(m, &logServiceWatchMessagesServer{stream})
}

type LogService_WatchMessagesServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type logServiceWatchMessagesServer struct {
	grpc.ServerStream
}

func (x *logServiceWatchMessagesServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _LogService_WatchAlerts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogServiceServer).WatchAlerts(m, &logServiceWatchAlertsServer{stream})
}

type LogService_WatchAlertsServer interface {
	Send(*Alert) error
	grpc.ServerStream
}

type logServiceWatchAlertsServer struct {
	grpc.ServerStream
}

func (x *logServiceWatchAlertsServer) Send(m *Alert) error {
	return x.ServerStream.SendMsg(m)
}

func _LogService_WatchLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogServiceServer).WatchLogs(m, &logServiceWatchLogsServer{stream})
}

type LogService_WatchLogsServer interface {
	Send(*Log) error
	grpc.ServerStream
}

type logServiceWatchLogsServer struct {
	grpc.ServerStream
}

func (x *logServiceWatchLogsServer) Send(m *Log) error {
	return x.ServerStream.SendMsg(m)
}

// LogService_ServiceDesc is the grpc.ServiceDesc for LogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "feeder.LogService",
	HandlerType: (*LogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _LogService_HealthCheck_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchMessages",
			Handler:       _LogService_WatchMessages_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WatchAlerts",
			Handler:       _LogService_WatchAlerts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WatchLogs",
			Handler:       _LogService_WatchLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "kubearmor.proto",
}
