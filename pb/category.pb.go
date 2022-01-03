// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/category.proto

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

type Category struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Slug                 string   `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Created              int64    `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bfb247fa9b1cc73, []int{0}
}

func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (m *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(m, src)
}
func (m *Category) XXX_Size() int {
	return xxx_messageInfo_Category.Size(m)
}
func (m *Category) XXX_DiscardUnknown() {
	xxx_messageInfo_Category.DiscardUnknown(m)
}

var xxx_messageInfo_Category proto.InternalMessageInfo

func (m *Category) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Category) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Category) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Category) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Category) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

type CreateCategoryRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCategoryRequest) Reset()         { *m = CreateCategoryRequest{} }
func (m *CreateCategoryRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCategoryRequest) ProtoMessage()    {}
func (*CreateCategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bfb247fa9b1cc73, []int{1}
}

func (m *CreateCategoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCategoryRequest.Unmarshal(m, b)
}
func (m *CreateCategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCategoryRequest.Marshal(b, m, deterministic)
}
func (m *CreateCategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCategoryRequest.Merge(m, src)
}
func (m *CreateCategoryRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCategoryRequest.Size(m)
}
func (m *CreateCategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCategoryRequest proto.InternalMessageInfo

func (m *CreateCategoryRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateCategoryReponse struct {
	Category             *Category `protobuf:"bytes,1,opt,name=Category,proto3" json:"Category,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateCategoryReponse) Reset()         { *m = CreateCategoryReponse{} }
func (m *CreateCategoryReponse) String() string { return proto.CompactTextString(m) }
func (*CreateCategoryReponse) ProtoMessage()    {}
func (*CreateCategoryReponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bfb247fa9b1cc73, []int{2}
}

func (m *CreateCategoryReponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCategoryReponse.Unmarshal(m, b)
}
func (m *CreateCategoryReponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCategoryReponse.Marshal(b, m, deterministic)
}
func (m *CreateCategoryReponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCategoryReponse.Merge(m, src)
}
func (m *CreateCategoryReponse) XXX_Size() int {
	return xxx_messageInfo_CreateCategoryReponse.Size(m)
}
func (m *CreateCategoryReponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCategoryReponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCategoryReponse proto.InternalMessageInfo

func (m *CreateCategoryReponse) GetCategory() *Category {
	if m != nil {
		return m.Category
	}
	return nil
}

type GetCategoriesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCategoriesRequest) Reset()         { *m = GetCategoriesRequest{} }
func (m *GetCategoriesRequest) String() string { return proto.CompactTextString(m) }
func (*GetCategoriesRequest) ProtoMessage()    {}
func (*GetCategoriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bfb247fa9b1cc73, []int{3}
}

func (m *GetCategoriesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCategoriesRequest.Unmarshal(m, b)
}
func (m *GetCategoriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCategoriesRequest.Marshal(b, m, deterministic)
}
func (m *GetCategoriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCategoriesRequest.Merge(m, src)
}
func (m *GetCategoriesRequest) XXX_Size() int {
	return xxx_messageInfo_GetCategoriesRequest.Size(m)
}
func (m *GetCategoriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCategoriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCategoriesRequest proto.InternalMessageInfo

type GetCategoriesReponse struct {
	Category             []*Category `protobuf:"bytes,1,rep,name=Category,proto3" json:"Category,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetCategoriesReponse) Reset()         { *m = GetCategoriesReponse{} }
func (m *GetCategoriesReponse) String() string { return proto.CompactTextString(m) }
func (*GetCategoriesReponse) ProtoMessage()    {}
func (*GetCategoriesReponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bfb247fa9b1cc73, []int{4}
}

func (m *GetCategoriesReponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCategoriesReponse.Unmarshal(m, b)
}
func (m *GetCategoriesReponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCategoriesReponse.Marshal(b, m, deterministic)
}
func (m *GetCategoriesReponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCategoriesReponse.Merge(m, src)
}
func (m *GetCategoriesReponse) XXX_Size() int {
	return xxx_messageInfo_GetCategoriesReponse.Size(m)
}
func (m *GetCategoriesReponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCategoriesReponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCategoriesReponse proto.InternalMessageInfo

func (m *GetCategoriesReponse) GetCategory() []*Category {
	if m != nil {
		return m.Category
	}
	return nil
}

type GetCategoryRequest struct {
	CategoryId           string   `protobuf:"bytes,1,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCategoryRequest) Reset()         { *m = GetCategoryRequest{} }
func (m *GetCategoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetCategoryRequest) ProtoMessage()    {}
func (*GetCategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bfb247fa9b1cc73, []int{5}
}

func (m *GetCategoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCategoryRequest.Unmarshal(m, b)
}
func (m *GetCategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCategoryRequest.Marshal(b, m, deterministic)
}
func (m *GetCategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCategoryRequest.Merge(m, src)
}
func (m *GetCategoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetCategoryRequest.Size(m)
}
func (m *GetCategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCategoryRequest proto.InternalMessageInfo

func (m *GetCategoryRequest) GetCategoryId() string {
	if m != nil {
		return m.CategoryId
	}
	return ""
}

type GetCategoryReponse struct {
	Category             *Category `protobuf:"bytes,1,opt,name=Category,proto3" json:"Category,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetCategoryReponse) Reset()         { *m = GetCategoryReponse{} }
func (m *GetCategoryReponse) String() string { return proto.CompactTextString(m) }
func (*GetCategoryReponse) ProtoMessage()    {}
func (*GetCategoryReponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bfb247fa9b1cc73, []int{6}
}

func (m *GetCategoryReponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCategoryReponse.Unmarshal(m, b)
}
func (m *GetCategoryReponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCategoryReponse.Marshal(b, m, deterministic)
}
func (m *GetCategoryReponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCategoryReponse.Merge(m, src)
}
func (m *GetCategoryReponse) XXX_Size() int {
	return xxx_messageInfo_GetCategoryReponse.Size(m)
}
func (m *GetCategoryReponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCategoryReponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCategoryReponse proto.InternalMessageInfo

func (m *GetCategoryReponse) GetCategory() *Category {
	if m != nil {
		return m.Category
	}
	return nil
}

type UpdateCategoryRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CategoryId           string   `protobuf:"bytes,3,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCategoryRequest) Reset()         { *m = UpdateCategoryRequest{} }
func (m *UpdateCategoryRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCategoryRequest) ProtoMessage()    {}
func (*UpdateCategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bfb247fa9b1cc73, []int{7}
}

func (m *UpdateCategoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCategoryRequest.Unmarshal(m, b)
}
func (m *UpdateCategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCategoryRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCategoryRequest.Merge(m, src)
}
func (m *UpdateCategoryRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCategoryRequest.Size(m)
}
func (m *UpdateCategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCategoryRequest proto.InternalMessageInfo

func (m *UpdateCategoryRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateCategoryRequest) GetCategoryId() string {
	if m != nil {
		return m.CategoryId
	}
	return ""
}

type UpdateCategoryResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCategoryResponse) Reset()         { *m = UpdateCategoryResponse{} }
func (m *UpdateCategoryResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateCategoryResponse) ProtoMessage()    {}
func (*UpdateCategoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bfb247fa9b1cc73, []int{8}
}

func (m *UpdateCategoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCategoryResponse.Unmarshal(m, b)
}
func (m *UpdateCategoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCategoryResponse.Marshal(b, m, deterministic)
}
func (m *UpdateCategoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCategoryResponse.Merge(m, src)
}
func (m *UpdateCategoryResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateCategoryResponse.Size(m)
}
func (m *UpdateCategoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCategoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCategoryResponse proto.InternalMessageInfo

func (m *UpdateCategoryResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type DeleteCategoryReponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCategoryReponse) Reset()         { *m = DeleteCategoryReponse{} }
func (m *DeleteCategoryReponse) String() string { return proto.CompactTextString(m) }
func (*DeleteCategoryReponse) ProtoMessage()    {}
func (*DeleteCategoryReponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bfb247fa9b1cc73, []int{9}
}

func (m *DeleteCategoryReponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCategoryReponse.Unmarshal(m, b)
}
func (m *DeleteCategoryReponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCategoryReponse.Marshal(b, m, deterministic)
}
func (m *DeleteCategoryReponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCategoryReponse.Merge(m, src)
}
func (m *DeleteCategoryReponse) XXX_Size() int {
	return xxx_messageInfo_DeleteCategoryReponse.Size(m)
}
func (m *DeleteCategoryReponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCategoryReponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCategoryReponse proto.InternalMessageInfo

func (m *DeleteCategoryReponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Category)(nil), "pb.Category")
	proto.RegisterType((*CreateCategoryRequest)(nil), "pb.CreateCategoryRequest")
	proto.RegisterType((*CreateCategoryReponse)(nil), "pb.CreateCategoryReponse")
	proto.RegisterType((*GetCategoriesRequest)(nil), "pb.GetCategoriesRequest")
	proto.RegisterType((*GetCategoriesReponse)(nil), "pb.GetCategoriesReponse")
	proto.RegisterType((*GetCategoryRequest)(nil), "pb.GetCategoryRequest")
	proto.RegisterType((*GetCategoryReponse)(nil), "pb.GetCategoryReponse")
	proto.RegisterType((*UpdateCategoryRequest)(nil), "pb.UpdateCategoryRequest")
	proto.RegisterType((*UpdateCategoryResponse)(nil), "pb.UpdateCategoryResponse")
	proto.RegisterType((*DeleteCategoryReponse)(nil), "pb.DeleteCategoryReponse")
}

func init() { proto.RegisterFile("proto/category.proto", fileDescriptor_1bfb247fa9b1cc73) }

var fileDescriptor_1bfb247fa9b1cc73 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcd, 0x8e, 0x94, 0x40,
	0x10, 0xc7, 0x03, 0x33, 0x8e, 0x5a, 0xab, 0x68, 0x2a, 0x3b, 0x88, 0xa3, 0x6b, 0x36, 0xad, 0x89,
	0x93, 0x31, 0x0e, 0xba, 0x7a, 0xf2, 0x60, 0xd4, 0x35, 0x31, 0xc6, 0xdb, 0x18, 0x2f, 0x7a, 0x82,
	0xa1, 0x42, 0x48, 0x90, 0x46, 0xba, 0x99, 0xc4, 0x18, 0x2f, 0xbe, 0x82, 0x4f, 0xe0, 0x33, 0xf9,
	0x0a, 0x3e, 0x88, 0xa1, 0x9b, 0x8f, 0x6d, 0x60, 0x27, 0x7a, 0xa3, 0x3e, 0xfa, 0xff, 0xab, 0xae,
	0x3f, 0x00, 0x87, 0x79, 0xc1, 0x25, 0xf7, 0xb7, 0x81, 0xa4, 0x98, 0x17, 0x5f, 0xd7, 0x2a, 0x44,
	0x3b, 0x0f, 0x17, 0xb7, 0x63, 0xce, 0xe3, 0x94, 0xfc, 0x20, 0x4f, 0xfc, 0x20, 0xcb, 0xb8, 0x0c,
	0x64, 0xc2, 0x33, 0xa1, 0x3b, 0xd8, 0x0e, 0x2e, 0x9d, 0xd6, 0x67, 0xd0, 0x01, 0x3b, 0x89, 0x3c,
	0xeb, 0xd8, 0x5a, 0x5e, 0xde, 0xd8, 0x49, 0x84, 0x08, 0xd3, 0x2c, 0xf8, 0x4c, 0x9e, 0xad, 0x32,
	0xea, 0xb9, 0xca, 0x89, 0xb4, 0x8c, 0xbd, 0x89, 0xce, 0x55, 0xcf, 0xe8, 0xc1, 0xc5, 0x6d, 0x41,
	0x81, 0xa4, 0xc8, 0x9b, 0x1e, 0x5b, 0xcb, 0xc9, 0xa6, 0x09, 0xab, 0x4a, 0x99, 0x47, 0xaa, 0x72,
	0x41, 0x57, 0xea, 0x90, 0x3d, 0x80, 0xf9, 0xa9, 0x6a, 0x6a, 0xe8, 0x1b, 0xfa, 0x52, 0x92, 0x90,
	0x2d, 0xd4, 0xea, 0xa0, 0xec, 0xe5, 0xb0, 0x39, 0xe7, 0x99, 0x20, 0x5c, 0x76, 0xd3, 0xab, 0x03,
	0x07, 0x27, 0x57, 0xd6, 0x79, 0xb8, 0x6e, 0xdb, 0xda, 0x2a, 0x73, 0xe1, 0xf0, 0x0d, 0xc9, 0x3a,
	0x4c, 0x48, 0xd4, 0x38, 0xf6, 0x62, 0x90, 0x1f, 0x53, 0x9e, 0xec, 0x51, 0x7e, 0x0a, 0xd8, 0x29,
	0xb4, 0xd7, 0xb8, 0x03, 0xd0, 0x78, 0xf1, 0xb6, 0xd9, 0xe9, 0x99, 0x0c, 0x7b, 0xde, 0x3b, 0xf5,
	0xbf, 0xf7, 0x79, 0x07, 0xf3, 0x0f, 0x6a, 0x95, 0xff, 0xb0, 0xbf, 0xde, 0x30, 0x93, 0xc1, 0x30,
	0x8f, 0xc0, 0xed, 0x8b, 0x09, 0x3d, 0x90, 0x0b, 0xb3, 0x82, 0x44, 0x99, 0xca, 0x5a, 0xaf, 0x8e,
	0x98, 0x0f, 0xf3, 0xd7, 0x94, 0xd2, 0xd0, 0x91, 0x73, 0x0e, 0x9c, 0xfc, 0x9a, 0xc2, 0xb5, 0xa6,
	0xf7, 0x3d, 0x15, 0xbb, 0x64, 0x4b, 0x48, 0xe0, 0x98, 0xb6, 0xe2, 0x4d, 0x75, 0xdb, 0xb1, 0xf7,
	0x62, 0x31, 0x5a, 0x52, 0x4c, 0x76, 0xf4, 0xe3, 0xf7, 0x9f, 0x9f, 0xf6, 0x0d, 0x86, 0xfe, 0xee,
	0x71, 0xf3, 0x05, 0x24, 0x24, 0xfc, 0x20, 0x8a, 0x9e, 0x59, 0x2b, 0xfc, 0x04, 0x57, 0x0d, 0x8b,
	0xd1, 0xab, 0xa4, 0xc6, 0xde, 0x86, 0xc5, 0x58, 0x45, 0x33, 0x5c, 0xc5, 0xb8, 0x8e, 0x8e, 0xc9,
	0xc0, 0x10, 0x0e, 0xce, 0xf8, 0x88, 0xae, 0x29, 0xd0, 0x4e, 0x3f, 0xcc, 0x6b, 0xd9, 0xbb, 0x4a,
	0xf6, 0x08, 0x6f, 0xf5, 0x46, 0xff, 0xd6, 0xb9, 0xf3, 0x1d, 0x4b, 0x70, 0x4c, 0x7b, 0xf4, 0x9e,
	0x46, 0xfd, 0xd7, 0x7b, 0x1a, 0xf5, 0x86, 0x3d, 0x54, 0xb0, 0xfb, 0x0b, 0xd6, 0x83, 0xe9, 0x6f,
	0xd2, 0x60, 0x56, 0x7b, 0xe3, 0xe0, 0x98, 0x3a, 0xe7, 0xde, 0x6e, 0x0f, 0x73, 0xa5, 0x98, 0xf7,
	0x56, 0x7d, 0x66, 0xa4, 0xba, 0x0d, 0xe6, 0xab, 0xd9, 0xc7, 0xe9, 0xda, 0xcf, 0xc3, 0x70, 0xa6,
	0x7e, 0x4d, 0x4f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xb1, 0x81, 0xfd, 0xd4, 0x04, 0x00,
	0x00,
}