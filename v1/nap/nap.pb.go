// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nap/nap.proto

package nap

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

type PolicyTicketStage int32

const (
	PolicyTicketStage_UNKNOWN_POLICY_TICKET_STAGE PolicyTicketStage = 0
	PolicyTicketStage_CREATED                     PolicyTicketStage = 1
	PolicyTicketStage_ROUTE_ANALYZED              PolicyTicketStage = 2
	PolicyTicketStage_OBJECT_ANALYZED             PolicyTicketStage = 3
	PolicyTicketStage_AUDITED                     PolicyTicketStage = 4
	PolicyTicketStage_DECLINED                    PolicyTicketStage = 5
	PolicyTicketStage_DISPATCH_PENDING            PolicyTicketStage = 6
	PolicyTicketStage_DISPATCHING                 PolicyTicketStage = 7
	PolicyTicketStage_DISPATCHED                  PolicyTicketStage = 8
	PolicyTicketStage_ROLLBACK_PENDING            PolicyTicketStage = 9
	PolicyTicketStage_ROLLING_BACK                PolicyTicketStage = 10
	PolicyTicketStage_ROLLED_BACK                 PolicyTicketStage = 11
	PolicyTicketStage_ROUTE_ANALYZED_ERROR        PolicyTicketStage = 12
	PolicyTicketStage_OBJECT_ANALYZED_ERROR       PolicyTicketStage = 13
	PolicyTicketStage_DISPATCHING_ERROR           PolicyTicketStage = 14
	PolicyTicketStage_ROLLING_BACK_ERROR          PolicyTicketStage = 15
	PolicyTicketStage_ERROR                       PolicyTicketStage = 16
)

var PolicyTicketStage_name = map[int32]string{
	0:  "UNKNOWN_POLICY_TICKET_STAGE",
	1:  "CREATED",
	2:  "ROUTE_ANALYZED",
	3:  "OBJECT_ANALYZED",
	4:  "AUDITED",
	5:  "DECLINED",
	6:  "DISPATCH_PENDING",
	7:  "DISPATCHING",
	8:  "DISPATCHED",
	9:  "ROLLBACK_PENDING",
	10: "ROLLING_BACK",
	11: "ROLLED_BACK",
	12: "ROUTE_ANALYZED_ERROR",
	13: "OBJECT_ANALYZED_ERROR",
	14: "DISPATCHING_ERROR",
	15: "ROLLING_BACK_ERROR",
	16: "ERROR",
}

var PolicyTicketStage_value = map[string]int32{
	"UNKNOWN_POLICY_TICKET_STAGE": 0,
	"CREATED":                     1,
	"ROUTE_ANALYZED":              2,
	"OBJECT_ANALYZED":             3,
	"AUDITED":                     4,
	"DECLINED":                    5,
	"DISPATCH_PENDING":            6,
	"DISPATCHING":                 7,
	"DISPATCHED":                  8,
	"ROLLBACK_PENDING":            9,
	"ROLLING_BACK":                10,
	"ROLLED_BACK":                 11,
	"ROUTE_ANALYZED_ERROR":        12,
	"OBJECT_ANALYZED_ERROR":       13,
	"DISPATCHING_ERROR":           14,
	"ROLLING_BACK_ERROR":          15,
	"ERROR":                       16,
}

func (x PolicyTicketStage) String() string {
	return proto.EnumName(PolicyTicketStage_name, int32(x))
}

func (PolicyTicketStage) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dde4d8ab16868518, []int{0}
}

type CapType int32

const (
	CapType_UNKNOWN_CAP_TYPE         CapType = 0
	CapType_POLICY_SEARCH            CapType = 1
	CapType_POLICY_TICKET            CapType = 2
	CapType_NSXT_POLICY_TICKET       CapType = 3
	CapType_OBJECTS_MANAGER          CapType = 4
	CapType_NAT_MANAGER              CapType = 5
	CapType_DEVICE_DISCOVERER        CapType = 6
	CapType_NAT_AND_POLICY           CapType = 7
	CapType_POLICY_TRACE             CapType = 8
	CapType_SWITCH_INTERFACE_CONFIG  CapType = 9
	CapType_POLICY_CHANGE_REPORT     CapType = 10
	CapType_POLICY_COMPLIANCE_RULE   CapType = 11
	CapType_DEVICE_QUOTA_MAX_DEVICES CapType = 12
)

var CapType_name = map[int32]string{
	0:  "UNKNOWN_CAP_TYPE",
	1:  "POLICY_SEARCH",
	2:  "POLICY_TICKET",
	3:  "NSXT_POLICY_TICKET",
	4:  "OBJECTS_MANAGER",
	5:  "NAT_MANAGER",
	6:  "DEVICE_DISCOVERER",
	7:  "NAT_AND_POLICY",
	8:  "POLICY_TRACE",
	9:  "SWITCH_INTERFACE_CONFIG",
	10: "POLICY_CHANGE_REPORT",
	11: "POLICY_COMPLIANCE_RULE",
	12: "DEVICE_QUOTA_MAX_DEVICES",
}

var CapType_value = map[string]int32{
	"UNKNOWN_CAP_TYPE":         0,
	"POLICY_SEARCH":            1,
	"POLICY_TICKET":            2,
	"NSXT_POLICY_TICKET":       3,
	"OBJECTS_MANAGER":          4,
	"NAT_MANAGER":              5,
	"DEVICE_DISCOVERER":        6,
	"NAT_AND_POLICY":           7,
	"POLICY_TRACE":             8,
	"SWITCH_INTERFACE_CONFIG":  9,
	"POLICY_CHANGE_REPORT":     10,
	"POLICY_COMPLIANCE_RULE":   11,
	"DEVICE_QUOTA_MAX_DEVICES": 12,
}

func (x CapType) String() string {
	return proto.EnumName(CapType_name, int32(x))
}

func (CapType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dde4d8ab16868518, []int{1}
}

type DeviceState int32

const (
	DeviceState_UNKNOWN_DEVICE_STATE DeviceState = 0
	DeviceState_RUNNING              DeviceState = 1
	DeviceState_DOWN                 DeviceState = 2
)

var DeviceState_name = map[int32]string{
	0: "UNKNOWN_DEVICE_STATE",
	1: "RUNNING",
	2: "DOWN",
}

var DeviceState_value = map[string]int32{
	"UNKNOWN_DEVICE_STATE": 0,
	"RUNNING":              1,
	"DOWN":                 2,
}

func (x DeviceState) String() string {
	return proto.EnumName(DeviceState_name, int32(x))
}

func (DeviceState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dde4d8ab16868518, []int{2}
}

func init() {
	proto.RegisterEnum("nap.PolicyTicketStage", PolicyTicketStage_name, PolicyTicketStage_value)
	proto.RegisterEnum("nap.CapType", CapType_name, CapType_value)
	proto.RegisterEnum("nap.DeviceState", DeviceState_name, DeviceState_value)
}

func init() {
	proto.RegisterFile("nap/nap.proto", fileDescriptor_dde4d8ab16868518)
}

var fileDescriptor_dde4d8ab16868518 = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0xdf, 0x4e, 0xdb, 0x30,
	0x14, 0xc6, 0x47, 0xf9, 0xd3, 0x72, 0x5a, 0xc0, 0x78, 0xc0, 0xd8, 0x98, 0xb4, 0xab, 0xdd, 0x20,
	0xd1, 0x6a, 0xda, 0xfd, 0x26, 0x63, 0x1f, 0x5a, 0xaf, 0xa9, 0x9d, 0x39, 0x0e, 0xa5, 0xdc, 0x58,
	0x21, 0x8b, 0x58, 0x55, 0xd6, 0x46, 0x23, 0x20, 0xf5, 0xb1, 0xf6, 0x4e, 0x7b, 0x90, 0xc9, 0x6d,
	0x02, 0x65, 0x77, 0xc9, 0xef, 0x9c, 0xcf, 0xe7, 0xcb, 0x77, 0x1c, 0xd8, 0x99, 0x26, 0x79, 0x67,
	0x9a, 0xe4, 0xed, 0xfc, 0xf7, 0xac, 0x98, 0xd1, 0xf5, 0x69, 0x92, 0x9f, 0xfe, 0xad, 0xc1, 0x7e,
	0x38, 0xbb, 0x1b, 0xa7, 0x73, 0x3b, 0x4e, 0x27, 0x59, 0x11, 0x15, 0xc9, 0x6d, 0x46, 0x3f, 0xc0,
	0x49, 0xac, 0xfa, 0x4a, 0x0f, 0x95, 0x0b, 0x75, 0x20, 0xf9, 0xc8, 0x59, 0xc9, 0xfb, 0x68, 0x5d,
	0x64, 0x59, 0x17, 0xc9, 0x2b, 0xda, 0x84, 0x3a, 0x37, 0xc8, 0x2c, 0x0a, 0xb2, 0x46, 0x29, 0xec,
	0x1a, 0x1d, 0x5b, 0x74, 0x4c, 0xb1, 0x60, 0x74, 0x8d, 0x82, 0xd4, 0xe8, 0x6b, 0xd8, 0xd3, 0xe7,
	0xdf, 0x90, 0xdb, 0x67, 0xb8, 0xee, 0x55, 0x2c, 0x16, 0xd2, 0xab, 0x36, 0x68, 0x0b, 0x1a, 0x02,
	0x79, 0x20, 0x15, 0x0a, 0xb2, 0x49, 0x0f, 0x80, 0x08, 0x19, 0x85, 0xcc, 0xf2, 0x9e, 0x0b, 0x51,
	0x09, 0xa9, 0xba, 0x64, 0x8b, 0xee, 0x41, 0xb3, 0xa2, 0x1e, 0xd4, 0xe9, 0x2e, 0x40, 0x05, 0x50,
	0x90, 0x86, 0x97, 0x19, 0x1d, 0x04, 0xe7, 0x8c, 0xf7, 0x9f, 0x64, 0xdb, 0x94, 0x40, 0xcb, 0x53,
	0xa9, 0xba, 0xce, 0x57, 0x08, 0xf8, 0x83, 0x3c, 0x41, 0xb1, 0x04, 0x4d, 0x7a, 0x0c, 0x07, 0x2f,
	0x3d, 0x3b, 0x34, 0x46, 0x1b, 0xd2, 0xa2, 0x6f, 0xe1, 0xf0, 0x3f, 0xe7, 0x65, 0x69, 0x87, 0x1e,
	0xc2, 0xfe, 0x8a, 0x9d, 0x12, 0xef, 0xd2, 0x23, 0xa0, 0xab, 0xe3, 0x4a, 0xbe, 0x47, 0xb7, 0x61,
	0x73, 0xf9, 0x48, 0x4e, 0xff, 0xd4, 0xa0, 0xce, 0x93, 0xdc, 0xce, 0xf3, 0xcc, 0x7b, 0xae, 0xc2,
	0xe5, 0x2c, 0x74, 0x76, 0x14, 0xfa, 0x44, 0xf7, 0x61, 0xa7, 0x8c, 0x3a, 0x42, 0x66, 0x78, 0x8f,
	0xac, 0xad, 0xa0, 0x65, 0xfa, 0xa4, 0xe6, 0x47, 0xa9, 0xe8, 0xca, 0xbe, 0xdc, 0x0a, 0x59, 0x7f,
	0x8e, 0x3b, 0x72, 0x03, 0xa6, 0x58, 0x17, 0x0d, 0xd9, 0xf0, 0x1f, 0xad, 0x98, 0x7d, 0x02, 0x9b,
	0x0b, 0xff, 0x78, 0x29, 0x39, 0x3a, 0x21, 0x23, 0xae, 0x2f, 0xd1, 0xa0, 0x21, 0x5b, 0x7e, 0x7f,
	0xbe, 0x8f, 0x29, 0x51, 0x9e, 0x4b, 0xea, 0x3e, 0xc2, 0x6a, 0x86, 0x61, 0x1c, 0x49, 0x83, 0x9e,
	0xc0, 0x9b, 0x68, 0x28, 0xfd, 0x7e, 0xa4, 0xb2, 0x68, 0x2e, 0x18, 0x47, 0xc7, 0xb5, 0xba, 0x90,
	0x3e, 0xf1, 0x63, 0x38, 0x28, 0xdb, 0x79, 0x8f, 0xa9, 0x2e, 0x3a, 0x83, 0xa1, 0x36, 0x96, 0x00,
	0x7d, 0x07, 0x47, 0x55, 0x45, 0x0f, 0xc2, 0x40, 0x32, 0xc5, 0xd1, 0x99, 0x38, 0x40, 0xd2, 0xa4,
	0xef, 0xe1, 0xb8, 0xf4, 0xf3, 0x3d, 0xd6, 0x96, 0xb9, 0x01, 0xbb, 0x72, 0x4b, 0x10, 0x91, 0xd6,
	0xe9, 0x17, 0x68, 0x8a, 0xec, 0x71, 0x9c, 0x66, 0x51, 0x91, 0x14, 0x99, 0x1f, 0x51, 0xc5, 0x56,
	0x8a, 0x22, 0xcb, 0x6c, 0x79, 0x19, 0x4d, 0xac, 0x94, 0xdf, 0xfd, 0x1a, 0x6d, 0xc0, 0x86, 0xd0,
	0x43, 0x45, 0x6a, 0xe7, 0x5f, 0xe1, 0x68, 0x9a, 0x15, 0xed, 0xfb, 0xc9, 0x3c, 0xbd, 0x9b, 0x3d,
	0xfc, 0x58, 0xde, 0xfa, 0xf6, 0x34, 0xc9, 0xaf, 0x3f, 0xde, 0x8e, 0x8b, 0x9f, 0x0f, 0x37, 0xed,
	0x74, 0xf6, 0xab, 0x73, 0x3f, 0x99, 0x9f, 0x2d, 0xea, 0x67, 0x45, 0x96, 0x76, 0x16, 0x3d, 0x9d,
	0xc7, 0x4f, 0xfe, 0x37, 0xb9, 0xd9, 0x5a, 0xbc, 0x7d, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0xd3,
	0xc3, 0x2a, 0x65, 0x38, 0x03, 0x00, 0x00,
}
