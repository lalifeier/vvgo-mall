// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: sms/interface/v1/sms_interface.proto

package v1

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

// SmsClient is the client API for Sms service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmsClient interface {
	SendSmsCode(ctx context.Context, in *SendSmsCodeRequest, opts ...grpc.CallOption) (*SendSmsCodeReply, error)
	VerifySmsCode(ctx context.Context, in *VerifySmsCodeRequest, opts ...grpc.CallOption) (*VerifySmsCodeReply, error)
}

type smsClient struct {
	cc grpc.ClientConnInterface
}

func NewSmsClient(cc grpc.ClientConnInterface) SmsClient {
	return &smsClient{cc}
}

func (c *smsClient) SendSmsCode(ctx context.Context, in *SendSmsCodeRequest, opts ...grpc.CallOption) (*SendSmsCodeReply, error) {
	out := new(SendSmsCodeReply)
	err := c.cc.Invoke(ctx, "/api.sms.interface.v1.Sms/SendSmsCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) VerifySmsCode(ctx context.Context, in *VerifySmsCodeRequest, opts ...grpc.CallOption) (*VerifySmsCodeReply, error) {
	out := new(VerifySmsCodeReply)
	err := c.cc.Invoke(ctx, "/api.sms.interface.v1.Sms/VerifySmsCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmsServer is the server API for Sms service.
// All implementations must embed UnimplementedSmsServer
// for forward compatibility
type SmsServer interface {
	SendSmsCode(context.Context, *SendSmsCodeRequest) (*SendSmsCodeReply, error)
	VerifySmsCode(context.Context, *VerifySmsCodeRequest) (*VerifySmsCodeReply, error)
	mustEmbedUnimplementedSmsServer()
}

// UnimplementedSmsServer must be embedded to have forward compatible implementations.
type UnimplementedSmsServer struct {
}

func (UnimplementedSmsServer) SendSmsCode(context.Context, *SendSmsCodeRequest) (*SendSmsCodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSmsCode not implemented")
}
func (UnimplementedSmsServer) VerifySmsCode(context.Context, *VerifySmsCodeRequest) (*VerifySmsCodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifySmsCode not implemented")
}
func (UnimplementedSmsServer) mustEmbedUnimplementedSmsServer() {}

// UnsafeSmsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmsServer will
// result in compilation errors.
type UnsafeSmsServer interface {
	mustEmbedUnimplementedSmsServer()
}

func RegisterSmsServer(s grpc.ServiceRegistrar, srv SmsServer) {
	s.RegisterService(&Sms_ServiceDesc, srv)
}

func _Sms_SendSmsCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSmsCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).SendSmsCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sms.interface.v1.Sms/SendSmsCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).SendSmsCode(ctx, req.(*SendSmsCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_VerifySmsCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifySmsCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).VerifySmsCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sms.interface.v1.Sms/VerifySmsCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).VerifySmsCode(ctx, req.(*VerifySmsCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sms_ServiceDesc is the grpc.ServiceDesc for Sms service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sms_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.sms.interface.v1.Sms",
	HandlerType: (*SmsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSmsCode",
			Handler:    _Sms_SendSmsCode_Handler,
		},
		{
			MethodName: "VerifySmsCode",
			Handler:    _Sms_VerifySmsCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sms/interface/v1/sms_interface.proto",
}
