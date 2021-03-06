// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/target.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Target struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Slug                 string   `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Created              int64    `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	StartDate            string   `protobuf:"bytes,6,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate              string   `protobuf:"bytes,7,opt,name=endDate,proto3" json:"endDate,omitempty"`
	Description          string   `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	UserId               int64    `protobuf:"varint,9,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Target) Reset()         { *m = Target{} }
func (m *Target) String() string { return proto.CompactTextString(m) }
func (*Target) ProtoMessage()    {}
func (*Target) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c555c36abe6851, []int{0}
}

func (m *Target) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Target.Unmarshal(m, b)
}
func (m *Target) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Target.Marshal(b, m, deterministic)
}
func (m *Target) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Target.Merge(m, src)
}
func (m *Target) XXX_Size() int {
	return xxx_messageInfo_Target.Size(m)
}
func (m *Target) XXX_DiscardUnknown() {
	xxx_messageInfo_Target.DiscardUnknown(m)
}

var xxx_messageInfo_Target proto.InternalMessageInfo

func (m *Target) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Target) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Target) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Target) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Target) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *Target) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *Target) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

func (m *Target) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Target) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CreateTargetRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	StartDate            string   `protobuf:"bytes,2,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate              string   `protobuf:"bytes,3,opt,name=endDate,proto3" json:"endDate,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	UserId               int64    `protobuf:"varint,5,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTargetRequest) Reset()         { *m = CreateTargetRequest{} }
func (m *CreateTargetRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTargetRequest) ProtoMessage()    {}
func (*CreateTargetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c555c36abe6851, []int{1}
}

func (m *CreateTargetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTargetRequest.Unmarshal(m, b)
}
func (m *CreateTargetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTargetRequest.Marshal(b, m, deterministic)
}
func (m *CreateTargetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTargetRequest.Merge(m, src)
}
func (m *CreateTargetRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTargetRequest.Size(m)
}
func (m *CreateTargetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTargetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTargetRequest proto.InternalMessageInfo

func (m *CreateTargetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateTargetRequest) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *CreateTargetRequest) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

func (m *CreateTargetRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateTargetRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CreateTargetReponse struct {
	Target               *Target  `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTargetReponse) Reset()         { *m = CreateTargetReponse{} }
func (m *CreateTargetReponse) String() string { return proto.CompactTextString(m) }
func (*CreateTargetReponse) ProtoMessage()    {}
func (*CreateTargetReponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c555c36abe6851, []int{2}
}

func (m *CreateTargetReponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTargetReponse.Unmarshal(m, b)
}
func (m *CreateTargetReponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTargetReponse.Marshal(b, m, deterministic)
}
func (m *CreateTargetReponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTargetReponse.Merge(m, src)
}
func (m *CreateTargetReponse) XXX_Size() int {
	return xxx_messageInfo_CreateTargetReponse.Size(m)
}
func (m *CreateTargetReponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTargetReponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTargetReponse proto.InternalMessageInfo

func (m *CreateTargetReponse) GetTarget() *Target {
	if m != nil {
		return m.Target
	}
	return nil
}

type GetTargetsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTargetsRequest) Reset()         { *m = GetTargetsRequest{} }
func (m *GetTargetsRequest) String() string { return proto.CompactTextString(m) }
func (*GetTargetsRequest) ProtoMessage()    {}
func (*GetTargetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c555c36abe6851, []int{3}
}

func (m *GetTargetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTargetsRequest.Unmarshal(m, b)
}
func (m *GetTargetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTargetsRequest.Marshal(b, m, deterministic)
}
func (m *GetTargetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTargetsRequest.Merge(m, src)
}
func (m *GetTargetsRequest) XXX_Size() int {
	return xxx_messageInfo_GetTargetsRequest.Size(m)
}
func (m *GetTargetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTargetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTargetsRequest proto.InternalMessageInfo

type GetTargetsReponse struct {
	Target               []*Target `protobuf:"bytes,1,rep,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetTargetsReponse) Reset()         { *m = GetTargetsReponse{} }
func (m *GetTargetsReponse) String() string { return proto.CompactTextString(m) }
func (*GetTargetsReponse) ProtoMessage()    {}
func (*GetTargetsReponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c555c36abe6851, []int{4}
}

func (m *GetTargetsReponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTargetsReponse.Unmarshal(m, b)
}
func (m *GetTargetsReponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTargetsReponse.Marshal(b, m, deterministic)
}
func (m *GetTargetsReponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTargetsReponse.Merge(m, src)
}
func (m *GetTargetsReponse) XXX_Size() int {
	return xxx_messageInfo_GetTargetsReponse.Size(m)
}
func (m *GetTargetsReponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTargetsReponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTargetsReponse proto.InternalMessageInfo

func (m *GetTargetsReponse) GetTarget() []*Target {
	if m != nil {
		return m.Target
	}
	return nil
}

type GetTargetRequest struct {
	TargetId             string   `protobuf:"bytes,1,opt,name=targetId,proto3" json:"targetId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTargetRequest) Reset()         { *m = GetTargetRequest{} }
func (m *GetTargetRequest) String() string { return proto.CompactTextString(m) }
func (*GetTargetRequest) ProtoMessage()    {}
func (*GetTargetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c555c36abe6851, []int{5}
}

func (m *GetTargetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTargetRequest.Unmarshal(m, b)
}
func (m *GetTargetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTargetRequest.Marshal(b, m, deterministic)
}
func (m *GetTargetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTargetRequest.Merge(m, src)
}
func (m *GetTargetRequest) XXX_Size() int {
	return xxx_messageInfo_GetTargetRequest.Size(m)
}
func (m *GetTargetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTargetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTargetRequest proto.InternalMessageInfo

func (m *GetTargetRequest) GetTargetId() string {
	if m != nil {
		return m.TargetId
	}
	return ""
}

type GetTargetReponse struct {
	Target               *Target  `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTargetReponse) Reset()         { *m = GetTargetReponse{} }
func (m *GetTargetReponse) String() string { return proto.CompactTextString(m) }
func (*GetTargetReponse) ProtoMessage()    {}
func (*GetTargetReponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c555c36abe6851, []int{6}
}

func (m *GetTargetReponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTargetReponse.Unmarshal(m, b)
}
func (m *GetTargetReponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTargetReponse.Marshal(b, m, deterministic)
}
func (m *GetTargetReponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTargetReponse.Merge(m, src)
}
func (m *GetTargetReponse) XXX_Size() int {
	return xxx_messageInfo_GetTargetReponse.Size(m)
}
func (m *GetTargetReponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTargetReponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTargetReponse proto.InternalMessageInfo

func (m *GetTargetReponse) GetTarget() *Target {
	if m != nil {
		return m.Target
	}
	return nil
}

type UpdateTargetRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	StartDate            string   `protobuf:"bytes,2,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate              string   `protobuf:"bytes,3,opt,name=endDate,proto3" json:"endDate,omitempty"`
	TargetId             string   `protobuf:"bytes,4,opt,name=targetId,proto3" json:"targetId,omitempty"`
	Description          string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTargetRequest) Reset()         { *m = UpdateTargetRequest{} }
func (m *UpdateTargetRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTargetRequest) ProtoMessage()    {}
func (*UpdateTargetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c555c36abe6851, []int{7}
}

func (m *UpdateTargetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTargetRequest.Unmarshal(m, b)
}
func (m *UpdateTargetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTargetRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTargetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTargetRequest.Merge(m, src)
}
func (m *UpdateTargetRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTargetRequest.Size(m)
}
func (m *UpdateTargetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTargetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTargetRequest proto.InternalMessageInfo

func (m *UpdateTargetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateTargetRequest) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *UpdateTargetRequest) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

func (m *UpdateTargetRequest) GetTargetId() string {
	if m != nil {
		return m.TargetId
	}
	return ""
}

func (m *UpdateTargetRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type UpdateTargetResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTargetResponse) Reset()         { *m = UpdateTargetResponse{} }
func (m *UpdateTargetResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateTargetResponse) ProtoMessage()    {}
func (*UpdateTargetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c555c36abe6851, []int{8}
}

func (m *UpdateTargetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTargetResponse.Unmarshal(m, b)
}
func (m *UpdateTargetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTargetResponse.Marshal(b, m, deterministic)
}
func (m *UpdateTargetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTargetResponse.Merge(m, src)
}
func (m *UpdateTargetResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateTargetResponse.Size(m)
}
func (m *UpdateTargetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTargetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTargetResponse proto.InternalMessageInfo

func (m *UpdateTargetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type DeleteTargetReponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTargetReponse) Reset()         { *m = DeleteTargetReponse{} }
func (m *DeleteTargetReponse) String() string { return proto.CompactTextString(m) }
func (*DeleteTargetReponse) ProtoMessage()    {}
func (*DeleteTargetReponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c555c36abe6851, []int{9}
}

func (m *DeleteTargetReponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTargetReponse.Unmarshal(m, b)
}
func (m *DeleteTargetReponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTargetReponse.Marshal(b, m, deterministic)
}
func (m *DeleteTargetReponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTargetReponse.Merge(m, src)
}
func (m *DeleteTargetReponse) XXX_Size() int {
	return xxx_messageInfo_DeleteTargetReponse.Size(m)
}
func (m *DeleteTargetReponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTargetReponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTargetReponse proto.InternalMessageInfo

func (m *DeleteTargetReponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Target)(nil), "pb.Target")
	proto.RegisterType((*CreateTargetRequest)(nil), "pb.CreateTargetRequest")
	proto.RegisterType((*CreateTargetReponse)(nil), "pb.CreateTargetReponse")
	proto.RegisterType((*GetTargetsRequest)(nil), "pb.GetTargetsRequest")
	proto.RegisterType((*GetTargetsReponse)(nil), "pb.GetTargetsReponse")
	proto.RegisterType((*GetTargetRequest)(nil), "pb.GetTargetRequest")
	proto.RegisterType((*GetTargetReponse)(nil), "pb.GetTargetReponse")
	proto.RegisterType((*UpdateTargetRequest)(nil), "pb.UpdateTargetRequest")
	proto.RegisterType((*UpdateTargetResponse)(nil), "pb.UpdateTargetResponse")
	proto.RegisterType((*DeleteTargetReponse)(nil), "pb.DeleteTargetReponse")
}

func init() { proto.RegisterFile("proto/target.proto", fileDescriptor_b2c555c36abe6851) }

var fileDescriptor_b2c555c36abe6851 = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x8a, 0xd4, 0x40,
	0x10, 0x26, 0x3f, 0x93, 0xdd, 0xa9, 0x99, 0x95, 0xb5, 0xb2, 0x3b, 0x1b, 0xc2, 0x0a, 0x43, 0x9f,
	0x96, 0x05, 0x13, 0x5c, 0x41, 0xd1, 0xa3, 0x2e, 0xc8, 0x9e, 0x84, 0xf1, 0x07, 0x11, 0x2f, 0xc9,
	0xa4, 0x19, 0x02, 0x63, 0x12, 0xd3, 0x9d, 0xbd, 0x88, 0x17, 0x5f, 0xc1, 0xa3, 0x07, 0x1f, 0x4a,
	0x1f, 0x41, 0x7c, 0x0e, 0x49, 0x75, 0x67, 0x26, 0x99, 0xc9, 0x0c, 0x5e, 0xbc, 0x75, 0xd5, 0x57,
	0xf5, 0xd5, 0x57, 0x3f, 0x34, 0x60, 0x51, 0xe6, 0x32, 0x0f, 0x65, 0x54, 0x2e, 0xb8, 0x0c, 0xc8,
	0x40, 0xb3, 0x88, 0xfd, 0xf3, 0x45, 0x9e, 0x2f, 0x96, 0x3c, 0x8c, 0x8a, 0x34, 0x8c, 0xb2, 0x2c,
	0x97, 0x91, 0x4c, 0xf3, 0x4c, 0xa8, 0x08, 0xf6, 0xc7, 0x00, 0xe7, 0x35, 0xa5, 0xe0, 0x1d, 0x30,
	0xd3, 0xc4, 0x33, 0xa6, 0xc6, 0xc5, 0x70, 0x66, 0xa6, 0x09, 0x22, 0xd8, 0x59, 0xf4, 0x91, 0x7b,
	0x26, 0x79, 0xe8, 0x5d, 0xfb, 0xc4, 0xb2, 0x5a, 0x78, 0x96, 0xf2, 0xd5, 0x6f, 0xf4, 0xe0, 0x60,
	0x5e, 0xf2, 0x48, 0xf2, 0xc4, 0xb3, 0xa7, 0xc6, 0x85, 0x35, 0x6b, 0xcc, 0x1a, 0xa9, 0x8a, 0x84,
	0x90, 0x81, 0x42, 0xb4, 0x89, 0xe7, 0x30, 0x14, 0x32, 0x2a, 0xe5, 0x75, 0x24, 0xb9, 0xe7, 0x10,
	0xd9, 0xda, 0x51, 0xe7, 0xf1, 0x2c, 0x21, 0xec, 0x80, 0xb0, 0xc6, 0xc4, 0x29, 0x8c, 0x12, 0x2e,
	0xe6, 0x65, 0x5a, 0xd4, 0x4d, 0x78, 0x87, 0x84, 0xb6, 0x5d, 0x38, 0x01, 0xa7, 0x12, 0xbc, 0xbc,
	0x49, 0xbc, 0x21, 0x95, 0xd4, 0x16, 0xfb, 0x6e, 0x80, 0xfb, 0x9c, 0x74, 0xa9, 0x76, 0x67, 0xfc,
	0x53, 0xc5, 0x85, 0x5c, 0x75, 0x69, 0xb4, 0xba, 0xec, 0xa8, 0x33, 0xf7, 0xa8, 0xb3, 0xf6, 0xaa,
	0xb3, 0xf7, 0xa9, 0x1b, 0x74, 0xd4, 0x3d, 0xd9, 0x14, 0x57, 0xe4, 0x99, 0xe0, 0xc8, 0xc0, 0x51,
	0xfb, 0x24, 0x79, 0xa3, 0x2b, 0x08, 0x8a, 0x38, 0xd0, 0x21, 0x1a, 0x61, 0x2e, 0xdc, 0x7d, 0xc1,
	0xa5, 0x72, 0x0a, 0xdd, 0x15, 0x7b, 0xdc, 0x75, 0x6e, 0xb3, 0x59, 0x3b, 0xd8, 0x02, 0x38, 0x5e,
	0x25, 0x36, 0x23, 0xf2, 0xe1, 0x50, 0xa1, 0x37, 0xcd, 0x79, 0xac, 0x6c, 0xf6, 0xa8, 0x13, 0xff,
	0xef, 0xaa, 0x7f, 0x18, 0xe0, 0xbe, 0xa1, 0x63, 0xf8, 0x7f, 0xeb, 0x68, 0xeb, 0xb6, 0xbb, 0xba,
	0x37, 0x57, 0x35, 0xd8, 0x5a, 0x15, 0x0b, 0xe0, 0xa4, 0x2b, 0x50, 0xa8, 0xee, 0x26, 0xe0, 0x94,
	0x5c, 0x54, 0x4b, 0xa9, 0x35, 0x6a, 0x8b, 0xdd, 0x07, 0xf7, 0x9a, 0x2f, 0xf9, 0xe6, 0x0a, 0x77,
	0x84, 0x5f, 0xfd, 0xb2, 0xe0, 0x48, 0x45, 0xbe, 0xe2, 0xe5, 0x6d, 0x3a, 0xe7, 0xf8, 0x0e, 0xc6,
	0xed, 0x1b, 0xc0, 0xb3, 0x7a, 0x6c, 0x3d, 0x27, 0xeb, 0xf7, 0x00, 0x54, 0x8b, 0x4d, 0xbe, 0xfe,
	0xfc, 0xfd, 0xcd, 0x3c, 0x66, 0xa3, 0xf0, 0xf6, 0x81, 0xfe, 0x08, 0xc4, 0x53, 0xe3, 0x12, 0x5f,
	0x02, 0xac, 0xaf, 0x01, 0x4f, 0xeb, 0xf4, 0xad, 0x93, 0xf1, 0xb7, 0xdc, 0x8a, 0xd3, 0x25, 0xce,
	0x23, 0x6c, 0x73, 0xe2, 0x5b, 0x18, 0xae, 0x22, 0xf1, 0xa4, 0x93, 0xd8, 0xd0, 0x6d, 0x7a, 0x15,
	0xdb, 0x3d, 0x62, 0x3b, 0xc3, 0xd3, 0x35, 0x5b, 0xf8, 0xb9, 0x59, 0xca, 0x17, 0x8c, 0x61, 0xdc,
	0x9e, 0xb9, 0x1a, 0x41, 0xcf, 0x99, 0xa8, 0x11, 0xf4, 0x8c, 0x9b, 0x4d, 0xa9, 0x80, 0xef, 0xf7,
	0x17, 0xa8, 0x87, 0xf1, 0x01, 0xc6, 0xed, 0xc4, 0x1d, 0xf2, 0x77, 0x16, 0xd0, 0x1d, 0x5c, 0x52,
	0x81, 0x84, 0x02, 0x5a, 0x05, 0x9e, 0x39, 0xef, 0xed, 0x20, 0x2c, 0xe2, 0xd8, 0xa1, 0xef, 0xf5,
	0xe1, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x8a, 0x81, 0x20, 0x96, 0x05, 0x00, 0x00,
}
