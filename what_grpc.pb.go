// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: what.proto

package trade

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

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloClient interface {
	Ping(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Nothing, error)
}

type helloClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloClient(cc grpc.ClientConnInterface) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) Ping(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/trade.Hello/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServer is the server API for Hello service.
// All implementations must embed UnimplementedHelloServer
// for forward compatibility
type HelloServer interface {
	Ping(context.Context, *Msg) (*Nothing, error)
	mustEmbedUnimplementedHelloServer()
}

// UnimplementedHelloServer must be embedded to have forward compatible implementations.
type UnimplementedHelloServer struct {
}

func (UnimplementedHelloServer) Ping(context.Context, *Msg) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedHelloServer) mustEmbedUnimplementedHelloServer() {}

// UnsafeHelloServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServer will
// result in compilation errors.
type UnsafeHelloServer interface {
	mustEmbedUnimplementedHelloServer()
}

func RegisterHelloServer(s grpc.ServiceRegistrar, srv HelloServer) {
	s.RegisterService(&Hello_ServiceDesc, srv)
}

func _Hello_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trade.Hello/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).Ping(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

// Hello_ServiceDesc is the grpc.ServiceDesc for Hello service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hello_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "trade.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Hello_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "what.proto",
}

// TextmeClient is the client API for Textme service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TextmeClient interface {
	Ping(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Nothing, error)
	Send(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	ListMessageSenders(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*MessageSenders, error)
	React(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Nothing, error)
	ReadTextTemplate(ctx context.Context, in *Key, opts ...grpc.CallOption) (*TextTemplate, error)
	ListTextTemplates(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*TextTemplates, error)
	CreateTextTemplate(ctx context.Context, in *TextTemplate, opts ...grpc.CallOption) (*TextTemplate, error)
	DeleteTextTemplate(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Nothing, error)
	UpdateTextTemplate(ctx context.Context, in *TextTemplate, opts ...grpc.CallOption) (*TextTemplate, error)
}

type textmeClient struct {
	cc grpc.ClientConnInterface
}

func NewTextmeClient(cc grpc.ClientConnInterface) TextmeClient {
	return &textmeClient{cc}
}

func (c *textmeClient) Ping(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/trade.Textme/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textmeClient) Send(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/trade.Textme/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textmeClient) ListMessageSenders(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*MessageSenders, error) {
	out := new(MessageSenders)
	err := c.cc.Invoke(ctx, "/trade.Textme/ListMessageSenders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textmeClient) React(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/trade.Textme/React", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textmeClient) ReadTextTemplate(ctx context.Context, in *Key, opts ...grpc.CallOption) (*TextTemplate, error) {
	out := new(TextTemplate)
	err := c.cc.Invoke(ctx, "/trade.Textme/ReadTextTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textmeClient) ListTextTemplates(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*TextTemplates, error) {
	out := new(TextTemplates)
	err := c.cc.Invoke(ctx, "/trade.Textme/ListTextTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textmeClient) CreateTextTemplate(ctx context.Context, in *TextTemplate, opts ...grpc.CallOption) (*TextTemplate, error) {
	out := new(TextTemplate)
	err := c.cc.Invoke(ctx, "/trade.Textme/CreateTextTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textmeClient) DeleteTextTemplate(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/trade.Textme/DeleteTextTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textmeClient) UpdateTextTemplate(ctx context.Context, in *TextTemplate, opts ...grpc.CallOption) (*TextTemplate, error) {
	out := new(TextTemplate)
	err := c.cc.Invoke(ctx, "/trade.Textme/UpdateTextTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TextmeServer is the server API for Textme service.
// All implementations must embed UnimplementedTextmeServer
// for forward compatibility
type TextmeServer interface {
	Ping(context.Context, *Msg) (*Nothing, error)
	Send(context.Context, *Message) (*Message, error)
	ListMessageSenders(context.Context, *ListRequest) (*MessageSenders, error)
	React(context.Context, *Message) (*Nothing, error)
	ReadTextTemplate(context.Context, *Key) (*TextTemplate, error)
	ListTextTemplates(context.Context, *ListRequest) (*TextTemplates, error)
	CreateTextTemplate(context.Context, *TextTemplate) (*TextTemplate, error)
	DeleteTextTemplate(context.Context, *Key) (*Nothing, error)
	UpdateTextTemplate(context.Context, *TextTemplate) (*TextTemplate, error)
	mustEmbedUnimplementedTextmeServer()
}

// UnimplementedTextmeServer must be embedded to have forward compatible implementations.
type UnimplementedTextmeServer struct {
}

func (UnimplementedTextmeServer) Ping(context.Context, *Msg) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedTextmeServer) Send(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedTextmeServer) ListMessageSenders(context.Context, *ListRequest) (*MessageSenders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMessageSenders not implemented")
}
func (UnimplementedTextmeServer) React(context.Context, *Message) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method React not implemented")
}
func (UnimplementedTextmeServer) ReadTextTemplate(context.Context, *Key) (*TextTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadTextTemplate not implemented")
}
func (UnimplementedTextmeServer) ListTextTemplates(context.Context, *ListRequest) (*TextTemplates, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTextTemplates not implemented")
}
func (UnimplementedTextmeServer) CreateTextTemplate(context.Context, *TextTemplate) (*TextTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTextTemplate not implemented")
}
func (UnimplementedTextmeServer) DeleteTextTemplate(context.Context, *Key) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTextTemplate not implemented")
}
func (UnimplementedTextmeServer) UpdateTextTemplate(context.Context, *TextTemplate) (*TextTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTextTemplate not implemented")
}
func (UnimplementedTextmeServer) mustEmbedUnimplementedTextmeServer() {}

// UnsafeTextmeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TextmeServer will
// result in compilation errors.
type UnsafeTextmeServer interface {
	mustEmbedUnimplementedTextmeServer()
}

func RegisterTextmeServer(s grpc.ServiceRegistrar, srv TextmeServer) {
	s.RegisterService(&Textme_ServiceDesc, srv)
}

func _Textme_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextmeServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trade.Textme/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextmeServer).Ping(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Textme_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextmeServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trade.Textme/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextmeServer).Send(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Textme_ListMessageSenders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextmeServer).ListMessageSenders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trade.Textme/ListMessageSenders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextmeServer).ListMessageSenders(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Textme_React_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextmeServer).React(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trade.Textme/React",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextmeServer).React(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Textme_ReadTextTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextmeServer).ReadTextTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trade.Textme/ReadTextTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextmeServer).ReadTextTemplate(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _Textme_ListTextTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextmeServer).ListTextTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trade.Textme/ListTextTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextmeServer).ListTextTemplates(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Textme_CreateTextTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextTemplate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextmeServer).CreateTextTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trade.Textme/CreateTextTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextmeServer).CreateTextTemplate(ctx, req.(*TextTemplate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Textme_DeleteTextTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextmeServer).DeleteTextTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trade.Textme/DeleteTextTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextmeServer).DeleteTextTemplate(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _Textme_UpdateTextTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextTemplate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextmeServer).UpdateTextTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trade.Textme/UpdateTextTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextmeServer).UpdateTextTemplate(ctx, req.(*TextTemplate))
	}
	return interceptor(ctx, in, info, handler)
}

// Textme_ServiceDesc is the grpc.ServiceDesc for Textme service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Textme_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "trade.Textme",
	HandlerType: (*TextmeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Textme_Ping_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Textme_Send_Handler,
		},
		{
			MethodName: "ListMessageSenders",
			Handler:    _Textme_ListMessageSenders_Handler,
		},
		{
			MethodName: "React",
			Handler:    _Textme_React_Handler,
		},
		{
			MethodName: "ReadTextTemplate",
			Handler:    _Textme_ReadTextTemplate_Handler,
		},
		{
			MethodName: "ListTextTemplates",
			Handler:    _Textme_ListTextTemplates_Handler,
		},
		{
			MethodName: "CreateTextTemplate",
			Handler:    _Textme_CreateTextTemplate_Handler,
		},
		{
			MethodName: "DeleteTextTemplate",
			Handler:    _Textme_DeleteTextTemplate_Handler,
		},
		{
			MethodName: "UpdateTextTemplate",
			Handler:    _Textme_UpdateTextTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "what.proto",
}

// SidefileClient is the client API for Sidefile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SidefileClient interface {
	ToThumbnails(ctx context.Context, in *ImageFile, opts ...grpc.CallOption) (*ImageFile, error)
}

type sidefileClient struct {
	cc grpc.ClientConnInterface
}

func NewSidefileClient(cc grpc.ClientConnInterface) SidefileClient {
	return &sidefileClient{cc}
}

func (c *sidefileClient) ToThumbnails(ctx context.Context, in *ImageFile, opts ...grpc.CallOption) (*ImageFile, error) {
	out := new(ImageFile)
	err := c.cc.Invoke(ctx, "/trade.Sidefile/ToThumbnails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SidefileServer is the server API for Sidefile service.
// All implementations must embed UnimplementedSidefileServer
// for forward compatibility
type SidefileServer interface {
	ToThumbnails(context.Context, *ImageFile) (*ImageFile, error)
	mustEmbedUnimplementedSidefileServer()
}

// UnimplementedSidefileServer must be embedded to have forward compatible implementations.
type UnimplementedSidefileServer struct {
}

func (UnimplementedSidefileServer) ToThumbnails(context.Context, *ImageFile) (*ImageFile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToThumbnails not implemented")
}
func (UnimplementedSidefileServer) mustEmbedUnimplementedSidefileServer() {}

// UnsafeSidefileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SidefileServer will
// result in compilation errors.
type UnsafeSidefileServer interface {
	mustEmbedUnimplementedSidefileServer()
}

func RegisterSidefileServer(s grpc.ServiceRegistrar, srv SidefileServer) {
	s.RegisterService(&Sidefile_ServiceDesc, srv)
}

func _Sidefile_ToThumbnails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageFile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SidefileServer).ToThumbnails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trade.Sidefile/ToThumbnails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SidefileServer).ToThumbnails(ctx, req.(*ImageFile))
	}
	return interceptor(ctx, in, info, handler)
}

// Sidefile_ServiceDesc is the grpc.ServiceDesc for Sidefile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sidefile_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "trade.Sidefile",
	HandlerType: (*SidefileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ToThumbnails",
			Handler:    _Sidefile_ToThumbnails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "what.proto",
}
