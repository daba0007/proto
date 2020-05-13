// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Retcode int32

const (
	Retcode_UNKNOWN_RETCODE       Retcode = 0
	Retcode_OK                    Retcode = 200
	Retcode_BadRequest            Retcode = 400
	Retcode_UNAUTHORIZED          Retcode = 401
	Retcode_SERVER_ERROR          Retcode = 500
	Retcode_NOT_SUCH_OBJECT       Retcode = 80001
	Retcode_PARAM_CHECK_FAILED    Retcode = 80002
	Retcode_DATABASE_OP_FAILED    Retcode = 80003
	Retcode_SYNTAX_ERROR          Retcode = 80004
	Retcode_PASSWORD_CHECK_FAILED Retcode = 80005
	Retcode_FILE_OP_FAILED        Retcode = 80006
	Retcode_OBJECT_CONFLICT       Retcode = 80007
)

var Retcode_name = map[int32]string{
	0:     "UNKNOWN_RETCODE",
	200:   "OK",
	400:   "BadRequest",
	401:   "UNAUTHORIZED",
	500:   "SERVER_ERROR",
	80001: "NOT_SUCH_OBJECT",
	80002: "PARAM_CHECK_FAILED",
	80003: "DATABASE_OP_FAILED",
	80004: "SYNTAX_ERROR",
	80005: "PASSWORD_CHECK_FAILED",
	80006: "FILE_OP_FAILED",
	80007: "OBJECT_CONFLICT",
}

var Retcode_value = map[string]int32{
	"UNKNOWN_RETCODE":       0,
	"OK":                    200,
	"BadRequest":            400,
	"UNAUTHORIZED":          401,
	"SERVER_ERROR":          500,
	"NOT_SUCH_OBJECT":       80001,
	"PARAM_CHECK_FAILED":    80002,
	"DATABASE_OP_FAILED":    80003,
	"SYNTAX_ERROR":          80004,
	"PASSWORD_CHECK_FAILED": 80005,
	"FILE_OP_FAILED":        80006,
	"OBJECT_CONFLICT":       80007,
}

func (x Retcode) String() string {
	return proto.EnumName(Retcode_name, int32(x))
}

func (Retcode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

type UnitState int32

const (
	UnitState_UNKNOWN_WORKFLOW_STATE UnitState = 0
	UnitState_CREATED                UnitState = 1
	UnitState_RUNNING                UnitState = 2
	UnitState_SUCCEEDED              UnitState = 3
	UnitState_FAILED                 UnitState = 4
	UnitState_SUSPEND                UnitState = 5
	UnitState_PENDING                UnitState = 6
)

var UnitState_name = map[int32]string{
	0: "UNKNOWN_WORKFLOW_STATE",
	1: "CREATED",
	2: "RUNNING",
	3: "SUCCEEDED",
	4: "FAILED",
	5: "SUSPEND",
	6: "PENDING",
}

var UnitState_value = map[string]int32{
	"UNKNOWN_WORKFLOW_STATE": 0,
	"CREATED":                1,
	"RUNNING":                2,
	"SUCCEEDED":              3,
	"FAILED":                 4,
	"SUSPEND":                5,
	"PENDING":                6,
}

func (x UnitState) String() string {
	return proto.EnumName(UnitState_name, int32(x))
}

func (UnitState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

type BaseRet struct {
	Retcode              Retcode  `protobuf:"varint,10,opt,name=retcode,proto3,enum=common.Retcode" json:"retcode,omitempty"`
	Message              string   `protobuf:"bytes,20,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseRet) Reset()         { *m = BaseRet{} }
func (m *BaseRet) String() string { return proto.CompactTextString(m) }
func (*BaseRet) ProtoMessage()    {}
func (*BaseRet) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

func (m *BaseRet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseRet.Unmarshal(m, b)
}
func (m *BaseRet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseRet.Marshal(b, m, deterministic)
}
func (m *BaseRet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseRet.Merge(m, src)
}
func (m *BaseRet) XXX_Size() int {
	return xxx_messageInfo_BaseRet.Size(m)
}
func (m *BaseRet) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseRet.DiscardUnknown(m)
}

var xxx_messageInfo_BaseRet proto.InternalMessageInfo

func (m *BaseRet) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_UNKNOWN_RETCODE
}

func (m *BaseRet) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Page struct {
	Index                int32      `protobuf:"varint,10,opt,name=index,proto3" json:"index,omitempty"`
	Num                  int32      `protobuf:"varint,20,opt,name=num,proto3" json:"num,omitempty"`
	Size                 int32      `protobuf:"varint,30,opt,name=size,proto3" json:"size,omitempty"`
	All                  int32      `protobuf:"varint,50,opt,name=all,proto3" json:"all,omitempty"`
	Total                int32      `protobuf:"varint,60,opt,name=total,proto3" json:"total,omitempty"`
	Data                 []*any.Any `protobuf:"bytes,70,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Page) Reset()         { *m = Page{} }
func (m *Page) String() string { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()    {}
func (*Page) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

func (m *Page) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Page.Unmarshal(m, b)
}
func (m *Page) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Page.Marshal(b, m, deterministic)
}
func (m *Page) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Page.Merge(m, src)
}
func (m *Page) XXX_Size() int {
	return xxx_messageInfo_Page.Size(m)
}
func (m *Page) XXX_DiscardUnknown() {
	xxx_messageInfo_Page.DiscardUnknown(m)
}

var xxx_messageInfo_Page proto.InternalMessageInfo

func (m *Page) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Page) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *Page) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Page) GetAll() int32 {
	if m != nil {
		return m.All
	}
	return 0
}

func (m *Page) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Page) GetData() []*any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("common.Retcode", Retcode_name, Retcode_value)
	proto.RegisterEnum("common.UnitState", UnitState_name, UnitState_value)
	proto.RegisterType((*BaseRet)(nil), "common.BaseRet")
	proto.RegisterType((*Page)(nil), "common.Page")
}

func init() {
	proto.RegisterFile("common/common.proto", fileDescriptor_8f954d82c0b891f6)
}

var fileDescriptor_8f954d82c0b891f6 = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x52, 0x5d, 0x8f, 0xd2, 0x40,
	0x14, 0x15, 0x28, 0x34, 0xdc, 0x5d, 0x97, 0x71, 0x96, 0x35, 0x75, 0x4d, 0x0c, 0xd9, 0x27, 0xdc,
	0x64, 0x21, 0xe2, 0xab, 0x2f, 0x43, 0x3b, 0x08, 0x82, 0x33, 0x64, 0xda, 0x8a, 0xee, 0x4b, 0x53,
	0x60, 0x44, 0x62, 0x69, 0x75, 0x3b, 0x6c, 0xc4, 0x37, 0x3f, 0xf7, 0x55, 0x13, 0x7f, 0x94, 0x7f,
	0xa7, 0x89, 0x3f, 0xc0, 0xf4, 0x83, 0x44, 0x9f, 0x7a, 0xcf, 0xb9, 0x67, 0x4e, 0xef, 0xbd, 0x39,
	0x70, 0xbc, 0x88, 0x36, 0x9b, 0x28, 0xec, 0xe6, 0x9f, 0xce, 0xbb, 0xab, 0x48, 0x45, 0xb8, 0x96,
	0xa3, 0xd3, 0x7b, 0xab, 0x28, 0x5a, 0x05, 0xb2, 0x9b, 0xb1, 0xf3, 0xed, 0xeb, 0xae, 0x1f, 0xee,
	0x72, 0xc9, 0x19, 0x03, 0xbd, 0xef, 0xc7, 0x52, 0x48, 0x85, 0x1f, 0x82, 0x7e, 0x25, 0xd5, 0x22,
	0x5a, 0x4a, 0x03, 0x5a, 0xa5, 0xf6, 0x51, 0xaf, 0xd1, 0x29, 0xdc, 0x44, 0x4e, 0x8b, 0x7d, 0x1f,
	0x1b, 0xa0, 0x6f, 0x64, 0x1c, 0xfb, 0x2b, 0x69, 0x34, 0x5b, 0xa5, 0x76, 0x5d, 0xec, 0xe1, 0xd9,
	0xaf, 0x12, 0x68, 0x53, 0x7f, 0x25, 0x71, 0x13, 0xaa, 0xeb, 0x70, 0x29, 0x3f, 0x64, 0x5e, 0x55,
	0x91, 0x03, 0x8c, 0xa0, 0x12, 0x6e, 0x37, 0xd9, 0xa3, 0xaa, 0x48, 0x4b, 0x8c, 0x41, 0x8b, 0xd7,
	0x1f, 0xa5, 0xf1, 0x20, 0xa3, 0xb2, 0x3a, 0x55, 0xf9, 0x41, 0x60, 0xf4, 0x72, 0x95, 0x1f, 0x04,
	0xa9, 0x9b, 0x8a, 0x94, 0x1f, 0x18, 0x4f, 0x72, 0xb7, 0x0c, 0xe0, 0x36, 0x68, 0x4b, 0x5f, 0xf9,
	0xc6, 0xa0, 0x55, 0x69, 0x1f, 0xf4, 0x9a, 0x9d, 0x7c, 0xcd, 0xce, 0x7e, 0xcd, 0x0e, 0x09, 0x77,
	0x22, 0x53, 0x9c, 0xdf, 0x94, 0x41, 0x2f, 0xb6, 0xc0, 0xc7, 0xd0, 0x70, 0xd9, 0x98, 0xf1, 0x19,
	0xf3, 0x04, 0x75, 0x4c, 0x6e, 0x51, 0x74, 0x0b, 0xeb, 0x50, 0xe6, 0x63, 0xf4, 0xbb, 0x84, 0x1b,
	0x00, 0x7d, 0x7f, 0x29, 0xe4, 0xfb, 0xad, 0x8c, 0x15, 0xfa, 0x51, 0xc1, 0x77, 0xe0, 0xd0, 0x65,
	0xc4, 0x75, 0x86, 0x5c, 0x8c, 0x2e, 0xa9, 0x85, 0x7e, 0x66, 0x94, 0x4d, 0xc5, 0x0b, 0x2a, 0x3c,
	0x2a, 0x04, 0x17, 0xe8, 0x4f, 0x05, 0x9f, 0x40, 0x83, 0x71, 0xc7, 0xb3, 0x5d, 0x73, 0xe8, 0xf1,
	0xfe, 0x33, 0x6a, 0x3a, 0xe8, 0x53, 0xa2, 0x61, 0x03, 0xf0, 0x94, 0x08, 0xf2, 0xdc, 0x33, 0x87,
	0xd4, 0x1c, 0x7b, 0x03, 0x32, 0x9a, 0x50, 0x0b, 0x7d, 0xce, 0x3b, 0x16, 0x71, 0x48, 0x9f, 0xd8,
	0xd4, 0xe3, 0xd3, 0x7d, 0xe7, 0x4b, 0xa2, 0x61, 0x0c, 0x87, 0xf6, 0x2b, 0xe6, 0x90, 0x97, 0x85,
	0xfb, 0xd7, 0x44, 0xc3, 0xf7, 0xe1, 0x64, 0x4a, 0x6c, 0x7b, 0xc6, 0x85, 0xf5, 0xbf, 0xd5, 0xb7,
	0x44, 0xc3, 0x4d, 0x38, 0x1a, 0x8c, 0x26, 0xff, 0xda, 0x7c, 0x4f, 0xb4, 0x74, 0xa2, 0x7c, 0x10,
	0xcf, 0xe4, 0x6c, 0x30, 0x19, 0x99, 0x0e, 0xba, 0x49, 0xb4, 0xf3, 0x6b, 0xa8, 0xbb, 0xe1, 0x5a,
	0xd9, 0xca, 0x57, 0x12, 0x9f, 0xc2, 0xdd, 0xfd, 0x29, 0x66, 0x5c, 0x8c, 0x07, 0x13, 0x3e, 0xf3,
	0x6c, 0x87, 0x38, 0xe9, 0x45, 0x0e, 0x40, 0x37, 0x05, 0x25, 0x0e, 0xb5, 0x50, 0x29, 0x05, 0xc2,
	0x65, 0x6c, 0xc4, 0x9e, 0xa2, 0x32, 0xbe, 0x0d, 0x75, 0xdb, 0x35, 0x4d, 0x4a, 0x2d, 0x6a, 0xa1,
	0x0a, 0x06, 0xa8, 0x15, 0xbf, 0xd5, 0x52, 0x9d, 0xed, 0xda, 0x53, 0xca, 0x2c, 0x54, 0x4d, 0x41,
	0x5a, 0xa5, 0x8f, 0x6a, 0xfd, 0x1e, 0x14, 0x69, 0xbc, 0x6c, 0xaf, 0xd6, 0xea, 0xcd, 0x76, 0x9e,
	0x86, 0xab, 0x1b, 0xbf, 0xdd, 0x5d, 0x2c, 0x82, 0x68, 0xbb, 0xbc, 0x50, 0x72, 0x91, 0xe7, 0xb3,
	0x7b, 0xfd, 0xa8, 0x48, 0xf1, 0xbc, 0x96, 0x11, 0x8f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xd0,
	0x6a, 0xa6, 0x14, 0xdd, 0x02, 0x00, 0x00,
}
