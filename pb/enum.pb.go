// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/enum.proto

package pb

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

type UserRole int32

const (
	UserRole_USER_UNKNOWN UserRole = 0
	UserRole_SUPER_ADMIN  UserRole = 1
	UserRole_ADMIN        UserRole = 2
	UserRole_CUSTOMER     UserRole = 3
	UserRole_USER         UserRole = 4
)

var UserRole_name = map[int32]string{
	0: "USER_UNKNOWN",
	1: "SUPER_ADMIN",
	2: "ADMIN",
	3: "CUSTOMER",
	4: "USER",
}

var UserRole_value = map[string]int32{
	"USER_UNKNOWN": 0,
	"SUPER_ADMIN":  1,
	"ADMIN":        2,
	"CUSTOMER":     3,
	"USER":         4,
}

func (x UserRole) String() string {
	return proto.EnumName(UserRole_name, int32(x))
}

func (UserRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_abeb21b8cd2b7d75, []int{0}
}

type UserSex int32

const (
	UserSex_SEX_UNKNOWN UserSex = 0
	UserSex_FEMALE      UserSex = 1
	UserSex_MALE        UserSex = 2
)

var UserSex_name = map[int32]string{
	0: "SEX_UNKNOWN",
	1: "FEMALE",
	2: "MALE",
}

var UserSex_value = map[string]int32{
	"SEX_UNKNOWN": 0,
	"FEMALE":      1,
	"MALE":        2,
}

func (x UserSex) String() string {
	return proto.EnumName(UserSex_name, int32(x))
}

func (UserSex) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_abeb21b8cd2b7d75, []int{1}
}

type FileType int32

const (
	FileType_NONE  FileType = 0
	FileType_PDF   FileType = 1
	FileType_EXCEL FileType = 2
	FileType_IMAGE FileType = 3
)

var FileType_name = map[int32]string{
	0: "NONE",
	1: "PDF",
	2: "EXCEL",
	3: "IMAGE",
}

var FileType_value = map[string]int32{
	"NONE":  0,
	"PDF":   1,
	"EXCEL": 2,
	"IMAGE": 3,
}

func (x FileType) String() string {
	return proto.EnumName(FileType_name, int32(x))
}

func (FileType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_abeb21b8cd2b7d75, []int{2}
}

type CharityType int32

const (
	CharityType_CHARITY_UNKNOWN CharityType = 0
	CharityType_CHARITY_GIVE    CharityType = 1
	CharityType_CHARITY_TAKE    CharityType = 2
)

var CharityType_name = map[int32]string{
	0: "CHARITY_UNKNOWN",
	1: "CHARITY_GIVE",
	2: "CHARITY_TAKE",
}

var CharityType_value = map[string]int32{
	"CHARITY_UNKNOWN": 0,
	"CHARITY_GIVE":    1,
	"CHARITY_TAKE":    2,
}

func (x CharityType) String() string {
	return proto.EnumName(CharityType_name, int32(x))
}

func (CharityType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_abeb21b8cd2b7d75, []int{3}
}

type ExchangeType int32

const (
	ExchangeType_EXCHANGE_UNKNOWN ExchangeType = 0
	ExchangeType_EXCHANGE_MAKE    ExchangeType = 1
	ExchangeType_EXCHANGE_TAKE    ExchangeType = 2
)

var ExchangeType_name = map[int32]string{
	0: "EXCHANGE_UNKNOWN",
	1: "EXCHANGE_MAKE",
	2: "EXCHANGE_TAKE",
}

var ExchangeType_value = map[string]int32{
	"EXCHANGE_UNKNOWN": 0,
	"EXCHANGE_MAKE":    1,
	"EXCHANGE_TAKE":    2,
}

func (x ExchangeType) String() string {
	return proto.EnumName(ExchangeType_name, int32(x))
}

func (ExchangeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_abeb21b8cd2b7d75, []int{4}
}

type CovidType int32

const (
	CovidType_COVID_UNKNOWN CovidType = 0
	CovidType_COVID_GIVE    CovidType = 1
	CovidType_COVID_TAKE    CovidType = 2
)

var CovidType_name = map[int32]string{
	0: "COVID_UNKNOWN",
	1: "COVID_GIVE",
	2: "COVID_TAKE",
}

var CovidType_value = map[string]int32{
	"COVID_UNKNOWN": 0,
	"COVID_GIVE":    1,
	"COVID_TAKE":    2,
}

func (x CovidType) String() string {
	return proto.EnumName(CovidType_name, int32(x))
}

func (CovidType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_abeb21b8cd2b7d75, []int{5}
}

func init() {
	proto.RegisterEnum("enum_pb.UserRole", UserRole_name, UserRole_value)
	proto.RegisterEnum("enum_pb.UserSex", UserSex_name, UserSex_value)
	proto.RegisterEnum("enum_pb.FileType", FileType_name, FileType_value)
	proto.RegisterEnum("enum_pb.CharityType", CharityType_name, CharityType_value)
	proto.RegisterEnum("enum_pb.ExchangeType", ExchangeType_name, ExchangeType_value)
	proto.RegisterEnum("enum_pb.CovidType", CovidType_name, CovidType_value)
}

func init() { proto.RegisterFile("proto/enum.proto", fileDescriptor_abeb21b8cd2b7d75) }

var fileDescriptor_abeb21b8cd2b7d75 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x4d, 0x4f, 0xf2, 0x40,
	0x14, 0x85, 0xf9, 0x7a, 0xa1, 0x5c, 0xca, 0xcb, 0x65, 0xf4, 0x57, 0x74, 0x01, 0x26, 0xec, 0x4d,
	0xc6, 0xe1, 0x16, 0x2a, 0x74, 0x4a, 0xfa, 0x81, 0xe8, 0x86, 0x80, 0x4e, 0x84, 0x04, 0x69, 0x83,
	0x68, 0xe0, 0xdf, 0x9b, 0x99, 0x22, 0xa9, 0xbb, 0x33, 0x27, 0x79, 0x9e, 0x73, 0x93, 0x01, 0xcc,
	0x0e, 0xe9, 0x31, 0xed, 0xab, 0xfd, 0xd7, 0x47, 0xcf, 0x44, 0xd6, 0xd0, 0x79, 0x99, 0xad, 0x9d,
	0x19, 0x58, 0xc9, 0xa7, 0x3a, 0x84, 0xe9, 0x4e, 0x31, 0x04, 0x3b, 0x89, 0x28, 0x5c, 0x26, 0x72,
	0x22, 0x83, 0x27, 0x89, 0x25, 0xd6, 0x81, 0x56, 0x94, 0xcc, 0x28, 0x5c, 0xf2, 0xa1, 0xef, 0x49,
	0x2c, 0xb3, 0x26, 0xfc, 0xcb, 0x63, 0x85, 0xd9, 0x60, 0x89, 0x24, 0x8a, 0x03, 0x9f, 0x42, 0xac,
	0x32, 0x0b, 0x6a, 0x9a, 0xc5, 0x9a, 0x73, 0x07, 0x0d, 0x6d, 0x8c, 0xd4, 0xc9, 0xe0, 0xb4, 0x28,
	0xf8, 0x00, 0xea, 0x2e, 0xf9, 0x7c, 0x4a, 0x58, 0xd6, 0x84, 0x49, 0x15, 0x67, 0x00, 0x96, 0xbb,
	0xdd, 0xa9, 0xf8, 0x9c, 0x29, 0xdd, 0xca, 0x40, 0x12, 0x96, 0x58, 0x03, 0xaa, 0xb3, 0xa1, 0x9b,
	0x6f, 0xd2, 0x42, 0xd0, 0x14, 0x2b, 0x3a, 0x7a, 0x3e, 0x1f, 0x11, 0x56, 0x1d, 0x17, 0x5a, 0x62,
	0xb3, 0x3a, 0x6c, 0x8f, 0x67, 0xc3, 0xdd, 0x40, 0x47, 0x8c, 0x79, 0xe8, 0xc5, 0xcf, 0x85, 0x39,
	0x04, 0xfb, 0xb7, 0x1c, 0x79, 0x73, 0x3d, 0x5a, 0x68, 0x62, 0x3e, 0xd1, 0xe3, 0x8f, 0x60, 0xd3,
	0xe9, 0x75, 0xb3, 0xda, 0xbf, 0xe7, 0x07, 0xdc, 0x02, 0xd2, 0x42, 0x8c, 0xb9, 0x1c, 0x51, 0xc1,
	0xd4, 0x85, 0xf6, 0xb5, 0xf5, 0x35, 0x58, 0xfe, 0x53, 0x5d, 0x5c, 0xf7, 0xd0, 0x14, 0xe9, 0xf7,
	0xf6, 0xcd, 0x88, 0xba, 0xd0, 0x16, 0xc1, 0xdc, 0x1b, 0x16, 0x2c, 0xff, 0x01, 0xf2, 0xea, 0x72,
	0xcd, 0xf5, 0x9d, 0xf3, 0x0f, 0xf5, 0x97, 0x5a, 0xaf, 0x9f, 0xad, 0xd7, 0x75, 0xf3, 0x49, 0x83,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xde, 0xd2, 0x2f, 0x48, 0xb8, 0x01, 0x00, 0x00,
}
