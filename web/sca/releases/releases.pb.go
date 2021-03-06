// Code generated by protoc-gen-go. DO NOT EDIT.
// source: releases.proto

package releases

import (
	context "context"
	fmt "fmt"
	pkgs "github.com/IBM/cloudland/web/sca/pkgs"
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

type ListRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b86db821e9daf711, []int{0}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeleteRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b86db821e9daf711, []int{1}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeleteRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type GetRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b86db821e9daf711, []int{2}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type RefreshRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshRequest) Reset()         { *m = RefreshRequest{} }
func (m *RefreshRequest) String() string { return proto.CompactTextString(m) }
func (*RefreshRequest) ProtoMessage()    {}
func (*RefreshRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b86db821e9daf711, []int{3}
}

func (m *RefreshRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshRequest.Unmarshal(m, b)
}
func (m *RefreshRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshRequest.Marshal(b, m, deterministic)
}
func (m *RefreshRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshRequest.Merge(m, src)
}
func (m *RefreshRequest) XXX_Size() int {
	return xxx_messageInfo_RefreshRequest.Size(m)
}
func (m *RefreshRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshRequest proto.InternalMessageInfo

func (m *RefreshRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b86db821e9daf711, []int{4}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AddRequest struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Pkg                  *pkgs.Pkg `protobuf:"bytes,2,opt,name=pkg,proto3" json:"pkg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b86db821e9daf711, []int{5}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddRequest) GetPkg() *pkgs.Pkg {
	if m != nil {
		return m.Pkg
	}
	return nil
}

type AddReply struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Pkg                  *pkgs.Pkg `protobuf:"bytes,2,opt,name=pkg,proto3" json:"pkg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddReply) Reset()         { *m = AddReply{} }
func (m *AddReply) String() string { return proto.CompactTextString(m) }
func (*AddReply) ProtoMessage()    {}
func (*AddReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b86db821e9daf711, []int{6}
}

func (m *AddReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddReply.Unmarshal(m, b)
}
func (m *AddReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddReply.Marshal(b, m, deterministic)
}
func (m *AddReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddReply.Merge(m, src)
}
func (m *AddReply) XXX_Size() int {
	return xxx_messageInfo_AddReply.Size(m)
}
func (m *AddReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AddReply.DiscardUnknown(m)
}

var xxx_messageInfo_AddReply proto.InternalMessageInfo

func (m *AddReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddReply) GetPkg() *pkgs.Pkg {
	if m != nil {
		return m.Pkg
	}
	return nil
}

type RemoveRequest struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Pkg                  *pkgs.Pkg `protobuf:"bytes,2,opt,name=pkg,proto3" json:"pkg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RemoveRequest) Reset()         { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()    {}
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b86db821e9daf711, []int{7}
}

func (m *RemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRequest.Unmarshal(m, b)
}
func (m *RemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRequest.Marshal(b, m, deterministic)
}
func (m *RemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRequest.Merge(m, src)
}
func (m *RemoveRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveRequest.Size(m)
}
func (m *RemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRequest proto.InternalMessageInfo

func (m *RemoveRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RemoveRequest) GetPkg() *pkgs.Pkg {
	if m != nil {
		return m.Pkg
	}
	return nil
}

type RemoveReply struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Pkg                  *pkgs.Pkg `protobuf:"bytes,2,opt,name=pkg,proto3" json:"pkg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RemoveReply) Reset()         { *m = RemoveReply{} }
func (m *RemoveReply) String() string { return proto.CompactTextString(m) }
func (*RemoveReply) ProtoMessage()    {}
func (*RemoveReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b86db821e9daf711, []int{8}
}

func (m *RemoveReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveReply.Unmarshal(m, b)
}
func (m *RemoveReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveReply.Marshal(b, m, deterministic)
}
func (m *RemoveReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveReply.Merge(m, src)
}
func (m *RemoveReply) XXX_Size() int {
	return xxx_messageInfo_RemoveReply.Size(m)
}
func (m *RemoveReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveReply.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveReply proto.InternalMessageInfo

func (m *RemoveReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RemoveReply) GetPkg() *pkgs.Pkg {
	if m != nil {
		return m.Pkg
	}
	return nil
}

type Release struct {
	Name                 string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string      `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Pkgs                 []*pkgs.Pkg `protobuf:"bytes,3,rep,name=pkgs,proto3" json:"pkgs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Release) Reset()         { *m = Release{} }
func (m *Release) String() string { return proto.CompactTextString(m) }
func (*Release) ProtoMessage()    {}
func (*Release) Descriptor() ([]byte, []int) {
	return fileDescriptor_b86db821e9daf711, []int{9}
}

func (m *Release) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Release.Unmarshal(m, b)
}
func (m *Release) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Release.Marshal(b, m, deterministic)
}
func (m *Release) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Release.Merge(m, src)
}
func (m *Release) XXX_Size() int {
	return xxx_messageInfo_Release.Size(m)
}
func (m *Release) XXX_DiscardUnknown() {
	xxx_messageInfo_Release.DiscardUnknown(m)
}

var xxx_messageInfo_Release proto.InternalMessageInfo

func (m *Release) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Release) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Release) GetPkgs() []*pkgs.Pkg {
	if m != nil {
		return m.Pkgs
	}
	return nil
}

type PublishRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishRequest) Reset()         { *m = PublishRequest{} }
func (m *PublishRequest) String() string { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()    {}
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b86db821e9daf711, []int{10}
}

func (m *PublishRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRequest.Unmarshal(m, b)
}
func (m *PublishRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRequest.Marshal(b, m, deterministic)
}
func (m *PublishRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRequest.Merge(m, src)
}
func (m *PublishRequest) XXX_Size() int {
	return xxx_messageInfo_PublishRequest.Size(m)
}
func (m *PublishRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRequest proto.InternalMessageInfo

func (m *PublishRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "com.github.ibm.cloudland.sca.releases.ListRequest")
	proto.RegisterType((*DeleteRequest)(nil), "com.github.ibm.cloudland.sca.releases.DeleteRequest")
	proto.RegisterType((*GetRequest)(nil), "com.github.ibm.cloudland.sca.releases.GetRequest")
	proto.RegisterType((*RefreshRequest)(nil), "com.github.ibm.cloudland.sca.releases.RefreshRequest")
	proto.RegisterType((*CreateRequest)(nil), "com.github.ibm.cloudland.sca.releases.CreateRequest")
	proto.RegisterType((*AddRequest)(nil), "com.github.ibm.cloudland.sca.releases.AddRequest")
	proto.RegisterType((*AddReply)(nil), "com.github.ibm.cloudland.sca.releases.AddReply")
	proto.RegisterType((*RemoveRequest)(nil), "com.github.ibm.cloudland.sca.releases.RemoveRequest")
	proto.RegisterType((*RemoveReply)(nil), "com.github.ibm.cloudland.sca.releases.RemoveReply")
	proto.RegisterType((*Release)(nil), "com.github.ibm.cloudland.sca.releases.Release")
	proto.RegisterType((*PublishRequest)(nil), "com.github.ibm.cloudland.sca.releases.PublishRequest")
}

func init() { proto.RegisterFile("releases.proto", fileDescriptor_b86db821e9daf711) }

var fileDescriptor_b86db821e9daf711 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x5d, 0x8b, 0x9b, 0x40,
	0x14, 0x25, 0x35, 0x68, 0x7b, 0xd3, 0xa4, 0x30, 0x4f, 0xe2, 0x53, 0x6a, 0x3f, 0x68, 0x29, 0x8c,
	0xad, 0x6d, 0xa1, 0x04, 0xfa, 0x90, 0xb4, 0x10, 0x0a, 0x2d, 0x04, 0x9f, 0x4a, 0x4a, 0x1f, 0xfc,
	0xb8, 0x6b, 0x24, 0x7e, 0xad, 0xa3, 0x59, 0xf6, 0x27, 0xec, 0xbf, 0x5e, 0x66, 0x8c, 0x9b, 0xf5,
	0x45, 0x66, 0x96, 0xec, 0x8b, 0xa8, 0xdc, 0x73, 0xcf, 0xf1, 0xce, 0x39, 0x57, 0x98, 0x55, 0x98,
	0xa2, 0xcf, 0x90, 0xd1, 0xb2, 0x2a, 0xea, 0x82, 0xbc, 0x09, 0x8b, 0x8c, 0xc6, 0x49, 0xbd, 0x6b,
	0x02, 0x9a, 0x04, 0x19, 0x0d, 0xd3, 0xa2, 0x89, 0x52, 0x3f, 0x8f, 0x28, 0x0b, 0x7d, 0xda, 0x15,
	0x5b, 0x2f, 0xca, 0x7d, 0xcc, 0x1c, 0x7e, 0x69, 0x71, 0xf6, 0x4b, 0x98, 0xfc, 0x4e, 0x58, 0xed,
	0xe1, 0x65, 0x83, 0xac, 0x26, 0x04, 0xc6, 0xb9, 0x9f, 0xa1, 0x39, 0x9a, 0x8f, 0xde, 0x3d, 0xf3,
	0xc4, 0xbd, 0xfd, 0x1d, 0xa6, 0x3f, 0x31, 0xc5, 0x1a, 0x07, 0x8a, 0x88, 0x09, 0xc6, 0x01, 0x2b,
	0x96, 0x14, 0xb9, 0xf9, 0x44, 0xbc, 0xee, 0x1e, 0xed, 0x05, 0xc0, 0x1a, 0xeb, 0x87, 0x61, 0x5f,
	0xc3, 0xcc, 0xc3, 0x8b, 0x0a, 0xd9, 0x6e, 0x48, 0xe0, 0x2b, 0x98, 0xfe, 0xa8, 0xd0, 0x1f, 0x14,
	0x68, 0x6f, 0x01, 0x96, 0x51, 0x34, 0x24, 0xe3, 0x1b, 0x68, 0xe5, 0x3e, 0x16, 0x12, 0x26, 0xee,
	0x5b, 0x3a, 0x38, 0x50, 0x31, 0xc1, 0xcd, 0x3e, 0xf6, 0x38, 0xc4, 0xfe, 0x0b, 0x4f, 0x45, 0xef,
	0x32, 0xbd, 0x3e, 0x73, 0xe7, 0xff, 0x30, 0xf5, 0x30, 0x2b, 0x0e, 0xf8, 0x38, 0xc2, 0xff, 0xc1,
	0xa4, 0x6b, 0x7f, 0x7e, 0xed, 0x0c, 0x0c, 0xaf, 0xf5, 0x9d, 0xda, 0xa9, 0x93, 0x05, 0x8c, 0x79,
	0x27, 0x53, 0x9b, 0x6b, 0x0a, 0x9c, 0x02, 0xc3, 0x1d, 0xb3, 0x69, 0x82, 0x34, 0x19, 0x74, 0x8c,
	0x7b, 0x63, 0xc0, 0xf3, 0xa3, 0xb6, 0x65, 0x94, 0x25, 0x39, 0xc9, 0x41, 0x6f, 0x2d, 0x44, 0xbe,
	0x50, 0xa9, 0x24, 0xd1, 0x9e, 0xe3, 0x2c, 0x2a, 0x89, 0xea, 0x06, 0x92, 0x83, 0xde, 0x66, 0x4a,
	0x9a, 0xaf, 0x17, 0x41, 0x65, 0xbe, 0x14, 0xc6, 0x3c, 0xe6, 0xc4, 0x95, 0xc4, 0xdd, 0xdb, 0x09,
	0xaa, 0x5c, 0x1f, 0x47, 0x64, 0x07, 0xda, 0x1a, 0x6b, 0xf2, 0x49, 0x12, 0x78, 0x5a, 0x0f, 0xca,
	0xdf, 0x95, 0x80, 0xb6, 0x8c, 0x22, 0x69, 0xa6, 0xd3, 0x06, 0xb0, 0x1c, 0x15, 0x08, 0x0f, 0x47,
	0x05, 0x7a, 0x9b, 0x15, 0xe9, 0x23, 0xeb, 0x25, 0xd7, 0x72, 0x15, 0x51, 0x9c, 0xb3, 0x04, 0xe3,
	0xe8, 0x66, 0xf2, 0x55, 0x12, 0xde, 0x77, 0xbf, 0xf2, 0x40, 0x4b, 0x1e, 0x5a, 0xb1, 0x71, 0xa5,
	0x19, 0xfb, 0x1b, 0x5a, 0x95, 0x71, 0xf5, 0x61, 0xfb, 0xfe, 0x58, 0x1c, 0x16, 0x99, 0xf3, 0x6b,
	0xf5, 0xc7, 0xb9, 0x03, 0x38, 0x57, 0x18, 0x38, 0x2c, 0xf4, 0x9d, 0x0e, 0x14, 0xe8, 0xe2, 0xaf,
	0xf5, 0xf9, 0x36, 0x00, 0x00, 0xff, 0xff, 0xae, 0xd9, 0xd1, 0x74, 0xff, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReleaseAdminClient is the client API for ReleaseAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReleaseAdminClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Release, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Release, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (ReleaseAdmin_ListClient, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Release, error)
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveReply, error)
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*Release, error)
	Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*Release, error)
}

type releaseAdminClient struct {
	cc *grpc.ClientConn
}

func NewReleaseAdminClient(cc *grpc.ClientConn) ReleaseAdminClient {
	return &releaseAdminClient{cc}
}

func (c *releaseAdminClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Release, error) {
	out := new(Release)
	err := c.cc.Invoke(ctx, "/com.github.ibm.cloudland.sca.releases.ReleaseAdmin/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseAdminClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Release, error) {
	out := new(Release)
	err := c.cc.Invoke(ctx, "/com.github.ibm.cloudland.sca.releases.ReleaseAdmin/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseAdminClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (ReleaseAdmin_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ReleaseAdmin_serviceDesc.Streams[0], "/com.github.ibm.cloudland.sca.releases.ReleaseAdmin/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &releaseAdminListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ReleaseAdmin_ListClient interface {
	Recv() (*Release, error)
	grpc.ClientStream
}

type releaseAdminListClient struct {
	grpc.ClientStream
}

func (x *releaseAdminListClient) Recv() (*Release, error) {
	m := new(Release)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *releaseAdminClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Release, error) {
	out := new(Release)
	err := c.cc.Invoke(ctx, "/com.github.ibm.cloudland.sca.releases.ReleaseAdmin/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseAdminClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error) {
	out := new(AddReply)
	err := c.cc.Invoke(ctx, "/com.github.ibm.cloudland.sca.releases.ReleaseAdmin/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseAdminClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveReply, error) {
	out := new(RemoveReply)
	err := c.cc.Invoke(ctx, "/com.github.ibm.cloudland.sca.releases.ReleaseAdmin/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseAdminClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*Release, error) {
	out := new(Release)
	err := c.cc.Invoke(ctx, "/com.github.ibm.cloudland.sca.releases.ReleaseAdmin/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseAdminClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*Release, error) {
	out := new(Release)
	err := c.cc.Invoke(ctx, "/com.github.ibm.cloudland.sca.releases.ReleaseAdmin/Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReleaseAdminServer is the server API for ReleaseAdmin service.
type ReleaseAdminServer interface {
	Create(context.Context, *CreateRequest) (*Release, error)
	Delete(context.Context, *DeleteRequest) (*Release, error)
	List(*ListRequest, ReleaseAdmin_ListServer) error
	Get(context.Context, *GetRequest) (*Release, error)
	Add(context.Context, *AddRequest) (*AddReply, error)
	Remove(context.Context, *RemoveRequest) (*RemoveReply, error)
	Publish(context.Context, *PublishRequest) (*Release, error)
	Refresh(context.Context, *RefreshRequest) (*Release, error)
}

// UnimplementedReleaseAdminServer can be embedded to have forward compatible implementations.
type UnimplementedReleaseAdminServer struct {
}

func (*UnimplementedReleaseAdminServer) Create(ctx context.Context, req *CreateRequest) (*Release, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedReleaseAdminServer) Delete(ctx context.Context, req *DeleteRequest) (*Release, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedReleaseAdminServer) List(req *ListRequest, srv ReleaseAdmin_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedReleaseAdminServer) Get(ctx context.Context, req *GetRequest) (*Release, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedReleaseAdminServer) Add(ctx context.Context, req *AddRequest) (*AddReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedReleaseAdminServer) Remove(ctx context.Context, req *RemoveRequest) (*RemoveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (*UnimplementedReleaseAdminServer) Publish(ctx context.Context, req *PublishRequest) (*Release, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (*UnimplementedReleaseAdminServer) Refresh(ctx context.Context, req *RefreshRequest) (*Release, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

func RegisterReleaseAdminServer(s *grpc.Server, srv ReleaseAdminServer) {
	s.RegisterService(&_ReleaseAdmin_serviceDesc, srv)
}

func _ReleaseAdmin_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseAdminServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.ibm.cloudland.sca.releases.ReleaseAdmin/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseAdminServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseAdmin_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseAdminServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.ibm.cloudland.sca.releases.ReleaseAdmin/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseAdminServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseAdmin_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReleaseAdminServer).List(m, &releaseAdminListServer{stream})
}

type ReleaseAdmin_ListServer interface {
	Send(*Release) error
	grpc.ServerStream
}

type releaseAdminListServer struct {
	grpc.ServerStream
}

func (x *releaseAdminListServer) Send(m *Release) error {
	return x.ServerStream.SendMsg(m)
}

func _ReleaseAdmin_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseAdminServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.ibm.cloudland.sca.releases.ReleaseAdmin/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseAdminServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseAdmin_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseAdminServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.ibm.cloudland.sca.releases.ReleaseAdmin/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseAdminServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseAdmin_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseAdminServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.ibm.cloudland.sca.releases.ReleaseAdmin/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseAdminServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseAdmin_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseAdminServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.ibm.cloudland.sca.releases.ReleaseAdmin/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseAdminServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseAdmin_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseAdminServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.ibm.cloudland.sca.releases.ReleaseAdmin/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseAdminServer).Refresh(ctx, req.(*RefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReleaseAdmin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.github.ibm.cloudland.sca.releases.ReleaseAdmin",
	HandlerType: (*ReleaseAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ReleaseAdmin_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ReleaseAdmin_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ReleaseAdmin_Get_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _ReleaseAdmin_Add_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _ReleaseAdmin_Remove_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _ReleaseAdmin_Publish_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _ReleaseAdmin_Refresh_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _ReleaseAdmin_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "releases.proto",
}
