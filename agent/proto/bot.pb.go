// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/focalsolution/micro-go-micro/agent/proto/bot.proto

package go_micro_bot

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type HelpRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelpRequest) Reset()         { *m = HelpRequest{} }
func (m *HelpRequest) String() string { return proto.CompactTextString(m) }
func (*HelpRequest) ProtoMessage()    {}
func (*HelpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_018e8d5b14a89d12, []int{0}
}

func (m *HelpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelpRequest.Unmarshal(m, b)
}
func (m *HelpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelpRequest.Marshal(b, m, deterministic)
}
func (m *HelpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelpRequest.Merge(m, src)
}
func (m *HelpRequest) XXX_Size() int {
	return xxx_messageInfo_HelpRequest.Size(m)
}
func (m *HelpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelpRequest proto.InternalMessageInfo

type HelpResponse struct {
	Usage                string   `protobuf:"bytes,1,opt,name=usage,proto3" json:"usage,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelpResponse) Reset()         { *m = HelpResponse{} }
func (m *HelpResponse) String() string { return proto.CompactTextString(m) }
func (*HelpResponse) ProtoMessage()    {}
func (*HelpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_018e8d5b14a89d12, []int{1}
}

func (m *HelpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelpResponse.Unmarshal(m, b)
}
func (m *HelpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelpResponse.Marshal(b, m, deterministic)
}
func (m *HelpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelpResponse.Merge(m, src)
}
func (m *HelpResponse) XXX_Size() int {
	return xxx_messageInfo_HelpResponse.Size(m)
}
func (m *HelpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelpResponse proto.InternalMessageInfo

func (m *HelpResponse) GetUsage() string {
	if m != nil {
		return m.Usage
	}
	return ""
}

func (m *HelpResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ExecRequest struct {
	Args                 []string `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecRequest) Reset()         { *m = ExecRequest{} }
func (m *ExecRequest) String() string { return proto.CompactTextString(m) }
func (*ExecRequest) ProtoMessage()    {}
func (*ExecRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_018e8d5b14a89d12, []int{2}
}

func (m *ExecRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecRequest.Unmarshal(m, b)
}
func (m *ExecRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecRequest.Marshal(b, m, deterministic)
}
func (m *ExecRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecRequest.Merge(m, src)
}
func (m *ExecRequest) XXX_Size() int {
	return xxx_messageInfo_ExecRequest.Size(m)
}
func (m *ExecRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecRequest proto.InternalMessageInfo

func (m *ExecRequest) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

type ExecResponse struct {
	Result               []byte   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecResponse) Reset()         { *m = ExecResponse{} }
func (m *ExecResponse) String() string { return proto.CompactTextString(m) }
func (*ExecResponse) ProtoMessage()    {}
func (*ExecResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_018e8d5b14a89d12, []int{3}
}

func (m *ExecResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecResponse.Unmarshal(m, b)
}
func (m *ExecResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecResponse.Marshal(b, m, deterministic)
}
func (m *ExecResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecResponse.Merge(m, src)
}
func (m *ExecResponse) XXX_Size() int {
	return xxx_messageInfo_ExecResponse.Size(m)
}
func (m *ExecResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecResponse proto.InternalMessageInfo

func (m *ExecResponse) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *ExecResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*HelpRequest)(nil), "go.micro.bot.HelpRequest")
	proto.RegisterType((*HelpResponse)(nil), "go.micro.bot.HelpResponse")
	proto.RegisterType((*ExecRequest)(nil), "go.micro.bot.ExecRequest")
	proto.RegisterType((*ExecResponse)(nil), "go.micro.bot.ExecResponse")
}

func init() {
	proto.RegisterFile("github.com/focalsolution/micro-go-micro/agent/proto/bot.proto", fileDescriptor_018e8d5b14a89d12)
}

var fileDescriptor_018e8d5b14a89d12 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0xdd, 0xea, 0xba, 0xb2, 0xd3, 0x7a, 0x09, 0x22, 0x75, 0x4f, 0x35, 0xa7, 0xbd, 0x98, 0x80,
	0x5e, 0x05, 0x0f, 0xa2, 0x78, 0xee, 0x3f, 0x68, 0xbb, 0x43, 0x2c, 0x6c, 0x3b, 0x35, 0x99, 0x82,
	0xff, 0xc1, 0x3f, 0x2d, 0x4d, 0x72, 0x08, 0xc5, 0xdb, 0x7b, 0x79, 0xe1, 0x7d, 0x0c, 0x68, 0xd3,
	0xf3, 0xd7, 0xdc, 0xaa, 0x8e, 0x06, 0x3d, 0xf4, 0x9d, 0x25, 0x6d, 0xe8, 0x31, 0x80, 0xc6, 0xe0,
	0xc8, 0x7a, 0xb2, 0xc4, 0xa4, 0x5b, 0x62, 0xe5, 0x91, 0x28, 0x0c, 0x29, 0xaf, 0xab, 0x96, 0x58,
	0xde, 0x40, 0xfe, 0x89, 0xe7, 0xa9, 0xc6, 0xef, 0x19, 0x1d, 0xcb, 0x0f, 0x28, 0x02, 0x75, 0x13,
	0x8d, 0x0e, 0xc5, 0x2d, 0x5c, 0xcd, 0xae, 0x31, 0x58, 0x66, 0x55, 0x76, 0xdc, 0xd7, 0x81, 0x88,
	0x0a, 0xf2, 0x13, 0xba, 0xce, 0xf6, 0x13, 0xf7, 0x34, 0x96, 0x17, 0x5e, 0x4b, 0x9f, 0xe4, 0x03,
	0xe4, 0xef, 0x3f, 0xd8, 0x45, 0x5b, 0x21, 0x60, 0xdb, 0x58, 0xe3, 0xca, 0xac, 0xba, 0x3c, 0xee,
	0x6b, 0x8f, 0xe5, 0x0b, 0x14, 0xe1, 0x4b, 0x8c, 0xba, 0x83, 0x9d, 0x45, 0x37, 0x9f, 0xd9, 0x67,
	0x15, 0x75, 0x64, 0x4b, 0x05, 0xb4, 0x96, 0x6c, 0x8c, 0x09, 0xe4, 0xe9, 0x37, 0x83, 0xeb, 0x37,
	0x1a, 0x86, 0x66, 0x3c, 0x89, 0x57, 0xd8, 0x2e, 0xa5, 0xc5, 0xbd, 0x4a, 0xa7, 0xa9, 0x64, 0xd7,
	0xe1, 0xf0, 0x9f, 0x14, 0x82, 0xe5, 0x66, 0x31, 0x58, 0xaa, 0xac, 0x0d, 0x92, 0x05, 0x6b, 0x83,
	0xb4, 0xb9, 0xdc, 0xb4, 0x3b, 0x7f, 0xda, 0xe7, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x18, 0xbd,
	0x39, 0x29, 0x8d, 0x01, 0x00, 0x00,
}
