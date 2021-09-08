// Code generated by protoc-gen-go. DO NOT EDIT.
// source: replica.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e84aa831fb48ea1, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Response struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Prev                 string   `protobuf:"bytes,3,opt,name=prev,proto3" json:"prev,omitempty"`
	Clock                int32    `protobuf:"varint,4,opt,name=clock,proto3" json:"clock,omitempty"`
	ID                   int32    `protobuf:"varint,5,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e84aa831fb48ea1, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Response) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Response) GetPrev() string {
	if m != nil {
		return m.Prev
	}
	return ""
}

func (m *Response) GetClock() int32 {
	if m != nil {
		return m.Clock
	}
	return 0
}

func (m *Response) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type Verify struct {
	Vertices             []string `protobuf:"bytes,1,rep,name=vertices,proto3" json:"vertices,omitempty"`
	ID                   int32    `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Verify) Reset()         { *m = Verify{} }
func (m *Verify) String() string { return proto.CompactTextString(m) }
func (*Verify) ProtoMessage()    {}
func (*Verify) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e84aa831fb48ea1, []int{2}
}

func (m *Verify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Verify.Unmarshal(m, b)
}
func (m *Verify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Verify.Marshal(b, m, deterministic)
}
func (m *Verify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Verify.Merge(m, src)
}
func (m *Verify) XXX_Size() int {
	return xxx_messageInfo_Verify.Size(m)
}
func (m *Verify) XXX_DiscardUnknown() {
	xxx_messageInfo_Verify.DiscardUnknown(m)
}

var xxx_messageInfo_Verify proto.InternalMessageInfo

func (m *Verify) GetVertices() []string {
	if m != nil {
		return m.Vertices
	}
	return nil
}

func (m *Verify) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e84aa831fb48ea1, []int{3}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Request)(nil), "protos.Request")
	proto.RegisterType((*Response)(nil), "protos.Response")
	proto.RegisterType((*Verify)(nil), "protos.Verify")
	proto.RegisterType((*Void)(nil), "protos.Void")
}

func init() { proto.RegisterFile("replica.proto", fileDescriptor_1e84aa831fb48ea1) }

var fileDescriptor_1e84aa831fb48ea1 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x9b, 0xbf, 0xb4, 0x07, 0x29, 0x95, 0xc5, 0x60, 0x32, 0x55, 0x9e, 0xba, 0x50, 0xa1,
	0x82, 0x98, 0x58, 0x10, 0x15, 0xa8, 0xab, 0x2b, 0x75, 0x0f, 0xce, 0xa1, 0x5a, 0x09, 0x75, 0xb0,
	0xdd, 0xa0, 0x7e, 0x7b, 0x14, 0xbb, 0x49, 0x27, 0xdf, 0xef, 0xf9, 0xde, 0x9d, 0x9f, 0x21, 0xd3,
	0xd8, 0xd4, 0x52, 0x14, 0xcb, 0x46, 0x2b, 0xab, 0x48, 0xea, 0x0e, 0xc3, 0xee, 0xe1, 0x8a, 0xe3,
	0xef, 0x11, 0x8d, 0x25, 0x53, 0x08, 0x65, 0x49, 0x83, 0x79, 0xb0, 0x48, 0x78, 0x28, 0x4b, 0x56,
	0xc3, 0x98, 0xa3, 0x69, 0xd4, 0xc1, 0x20, 0x99, 0x41, 0x54, 0xe1, 0xc9, 0x5d, 0x4e, 0x78, 0x57,
	0x92, 0x3b, 0x48, 0xda, 0xa2, 0x3e, 0x22, 0x0d, 0x9d, 0xe6, 0x81, 0x10, 0x88, 0x1b, 0x8d, 0x2d,
	0x8d, 0x9c, 0xe8, 0xea, 0xae, 0x53, 0xd4, 0x4a, 0x54, 0x34, 0x76, 0xa3, 0x3d, 0x74, 0xdb, 0x36,
	0x6b, 0x9a, 0xf8, 0x6d, 0x9b, 0x35, 0x7b, 0x86, 0x74, 0x87, 0x5a, 0x7e, 0x9f, 0x48, 0x0e, 0xe3,
	0x16, 0xb5, 0x95, 0x02, 0x0d, 0x0d, 0xe6, 0xd1, 0x62, 0xc2, 0x07, 0x3e, 0xbb, 0xc2, 0xc1, 0x95,
	0x42, 0xbc, 0x53, 0xb2, 0x5c, 0x7d, 0x42, 0xb6, 0xb5, 0x1a, 0x8b, 0x9f, 0x2d, 0xea, 0x56, 0x0a,
	0x24, 0x2f, 0x90, 0x7d, 0xa0, 0x15, 0xfb, 0x21, 0xc1, 0xad, 0x0f, 0x6e, 0x96, 0xe7, 0xb8, 0xf9,
	0xec, 0x22, 0xf8, 0x16, 0x36, 0x7a, 0x0c, 0x56, 0xaf, 0x70, 0xfd, 0xbe, 0x2f, 0x6c, 0x3f, 0xe6,
	0xa1, 0x43, 0x14, 0xd5, 0xdb, 0xc1, 0xfc, 0xa1, 0x26, 0xd3, 0xde, 0xe3, 0x9f, 0x9a, 0xdf, 0x0c,
	0xac, 0x64, 0xc9, 0x46, 0x5f, 0xfe, 0x57, 0x9f, 0xfe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x19,
	0x05, 0xa3, 0x6d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamServiceClient interface {
	FetchResponse(ctx context.Context, in *Request, opts ...grpc.CallOption) (StreamService_FetchResponseClient, error)
}

type streamServiceClient struct {
	cc *grpc.ClientConn
}

func NewStreamServiceClient(cc *grpc.ClientConn) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) FetchResponse(ctx context.Context, in *Request, opts ...grpc.CallOption) (StreamService_FetchResponseClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[0], "/protos.StreamService/FetchResponse", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceFetchResponseClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamService_FetchResponseClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type streamServiceFetchResponseClient struct {
	grpc.ClientStream
}

func (x *streamServiceFetchResponseClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServiceServer is the server API for StreamService service.
type StreamServiceServer interface {
	FetchResponse(*Request, StreamService_FetchResponseServer) error
}

// UnimplementedStreamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStreamServiceServer struct {
}

func (*UnimplementedStreamServiceServer) FetchResponse(req *Request, srv StreamService_FetchResponseServer) error {
	return status.Errorf(codes.Unimplemented, "method FetchResponse not implemented")
}

func RegisterStreamServiceServer(s *grpc.Server, srv StreamServiceServer) {
	s.RegisterService(&_StreamService_serviceDesc, srv)
}

func _StreamService_FetchResponse_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServiceServer).FetchResponse(m, &streamServiceFetchResponseServer{stream})
}

type StreamService_FetchResponseServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type streamServiceFetchResponseServer struct {
	grpc.ServerStream
}

func (x *streamServiceFetchResponseServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

var _StreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FetchResponse",
			Handler:       _StreamService_FetchResponse_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "replica.proto",
}

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	CheckAnswer(ctx context.Context, in *Verify, opts ...grpc.CallOption) (*Void, error)
}

type chatServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatServiceClient(cc *grpc.ClientConn) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) CheckAnswer(ctx context.Context, in *Verify, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.ChatService/CheckAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	CheckAnswer(context.Context, *Verify) (*Void, error)
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) CheckAnswer(ctx context.Context, req *Verify) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAnswer not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_CheckAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Verify)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CheckAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ChatService/CheckAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CheckAnswer(ctx, req.(*Verify))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckAnswer",
			Handler:    _ChatService_CheckAnswer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "replica.proto",
}
