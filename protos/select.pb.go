// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/select.proto

package _select

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

// The request message containing the sql statement.
type SelectRequest struct {
	Sql                  string   `protobuf:"bytes,1,opt,name=sql,proto3" json:"sql,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SelectRequest) Reset()         { *m = SelectRequest{} }
func (m *SelectRequest) String() string { return proto.CompactTextString(m) }
func (*SelectRequest) ProtoMessage()    {}
func (*SelectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_094bf54abdad424b, []int{0}
}

func (m *SelectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelectRequest.Unmarshal(m, b)
}
func (m *SelectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelectRequest.Marshal(b, m, deterministic)
}
func (m *SelectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelectRequest.Merge(m, src)
}
func (m *SelectRequest) XXX_Size() int {
	return xxx_messageInfo_SelectRequest.Size(m)
}
func (m *SelectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SelectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SelectRequest proto.InternalMessageInfo

func (m *SelectRequest) GetSql() string {
	if m != nil {
		return m.Sql
	}
	return ""
}

// The response message containing the greetings
type SelectReply struct {
	Results              []*SelectReply_Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SelectReply) Reset()         { *m = SelectReply{} }
func (m *SelectReply) String() string { return proto.CompactTextString(m) }
func (*SelectReply) ProtoMessage()    {}
func (*SelectReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_094bf54abdad424b, []int{1}
}

func (m *SelectReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelectReply.Unmarshal(m, b)
}
func (m *SelectReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelectReply.Marshal(b, m, deterministic)
}
func (m *SelectReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelectReply.Merge(m, src)
}
func (m *SelectReply) XXX_Size() int {
	return xxx_messageInfo_SelectReply.Size(m)
}
func (m *SelectReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SelectReply.DiscardUnknown(m)
}

var xxx_messageInfo_SelectReply proto.InternalMessageInfo

func (m *SelectReply) GetResults() []*SelectReply_Result {
	if m != nil {
		return m.Results
	}
	return nil
}

type SelectReply_Result struct {
	Fields               []string `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SelectReply_Result) Reset()         { *m = SelectReply_Result{} }
func (m *SelectReply_Result) String() string { return proto.CompactTextString(m) }
func (*SelectReply_Result) ProtoMessage()    {}
func (*SelectReply_Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_094bf54abdad424b, []int{1, 0}
}

func (m *SelectReply_Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelectReply_Result.Unmarshal(m, b)
}
func (m *SelectReply_Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelectReply_Result.Marshal(b, m, deterministic)
}
func (m *SelectReply_Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelectReply_Result.Merge(m, src)
}
func (m *SelectReply_Result) XXX_Size() int {
	return xxx_messageInfo_SelectReply_Result.Size(m)
}
func (m *SelectReply_Result) XXX_DiscardUnknown() {
	xxx_messageInfo_SelectReply_Result.DiscardUnknown(m)
}

var xxx_messageInfo_SelectReply_Result proto.InternalMessageInfo

func (m *SelectReply_Result) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func init() {
	proto.RegisterType((*SelectRequest)(nil), "select.SelectRequest")
	proto.RegisterType((*SelectReply)(nil), "select.SelectReply")
	proto.RegisterType((*SelectReply_Result)(nil), "select.SelectReply.Result")
}

func init() { proto.RegisterFile("protos/select.proto", fileDescriptor_094bf54abdad424b) }

var fileDescriptor_094bf54abdad424b = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x2f, 0x4e, 0xcd, 0x49, 0x4d, 0x2e, 0xd1, 0x03, 0xf3, 0x84, 0xd8, 0x20, 0x3c,
	0x25, 0x45, 0x2e, 0xde, 0x60, 0x30, 0x2b, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x80,
	0x8b, 0xb9, 0xb8, 0x30, 0x47, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc4, 0x54, 0x4a, 0xe5,
	0xe2, 0x86, 0x29, 0x29, 0xc8, 0xa9, 0x14, 0x32, 0xe1, 0x62, 0x2f, 0x4a, 0x2d, 0x2e, 0xcd, 0x29,
	0x29, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x92, 0xd2, 0x83, 0x9a, 0x8c, 0xa4, 0x4a, 0x2f,
	0x08, 0xac, 0x24, 0x08, 0xa6, 0x54, 0x4a, 0x81, 0x8b, 0x0d, 0x22, 0x24, 0x24, 0xc6, 0xc5, 0x96,
	0x96, 0x99, 0x9a, 0x93, 0x02, 0xd1, 0xce, 0x19, 0x04, 0xe5, 0x19, 0x79, 0x72, 0xb1, 0xbb, 0x17,
	0xa5, 0xa6, 0x96, 0xa4, 0x16, 0x09, 0xd9, 0x71, 0xf1, 0x41, 0xcc, 0x72, 0x49, 0x2c, 0x49, 0x4c,
	0x4a, 0x2c, 0x4e, 0x15, 0x12, 0x45, 0xb7, 0x03, 0xec, 0x58, 0x29, 0x61, 0x2c, 0x56, 0x2b, 0x31,
	0x38, 0xe9, 0x71, 0x89, 0x67, 0xe6, 0xeb, 0xa5, 0x17, 0x15, 0x24, 0xeb, 0xa5, 0x56, 0x24, 0xe6,
	0x16, 0xe4, 0xa4, 0x16, 0x43, 0x15, 0x3a, 0x41, 0xbd, 0x12, 0x00, 0x0a, 0x84, 0x00, 0xc6, 0x45,
	0x4c, 0xcc, 0x1e, 0x3e, 0xe1, 0x49, 0x6c, 0xe0, 0x30, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xd5, 0xb1, 0xf0, 0xfc, 0x2a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SelectDatabase(ctx context.Context, in *SelectRequest, opts ...grpc.CallOption) (*SelectReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SelectDatabase(ctx context.Context, in *SelectRequest, opts ...grpc.CallOption) (*SelectReply, error) {
	out := new(SelectReply)
	err := c.cc.Invoke(ctx, "/select.Greeter/SelectDatabase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	// Sends a greeting
	SelectDatabase(context.Context, *SelectRequest) (*SelectReply, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) SelectDatabase(ctx context.Context, req *SelectRequest) (*SelectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectDatabase not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SelectDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SelectDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/select.Greeter/SelectDatabase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SelectDatabase(ctx, req.(*SelectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "select.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SelectDatabase",
			Handler:    _Greeter_SelectDatabase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/select.proto",
}
