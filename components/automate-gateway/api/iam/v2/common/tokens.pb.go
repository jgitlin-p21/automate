// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2/common/tokens.proto

package common

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

type Token struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Active               bool     `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Projects             []string `protobuf:"bytes,7,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f15397d34553417, []int{0}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Token) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Token) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Token) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Token) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Token) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Token) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

func init() {
	proto.RegisterType((*Token)(nil), "chef.automate.api.iam.v2.Token")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2/common/tokens.proto", fileDescriptor_4f15397d34553417)
}

var fileDescriptor_4f15397d34553417 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0xb4, 0x09, 0x8d, 0x07, 0x06, 0x0b, 0x21, 0x0b, 0x09, 0x29, 0x62, 0xca, 0x82,
	0x2d, 0x95, 0x8d, 0xad, 0x0c, 0x88, 0x39, 0x62, 0x62, 0x41, 0x57, 0xe7, 0x68, 0x0d, 0x38, 0x67,
	0x25, 0x97, 0x20, 0x7e, 0x17, 0x7f, 0x10, 0xc5, 0x4d, 0xbb, 0x77, 0xf3, 0xfb, 0x9e, 0xdf, 0x70,
	0x9f, 0x78, 0xb4, 0xe4, 0x03, 0xb5, 0xd8, 0x72, 0x6f, 0x60, 0x60, 0xf2, 0xc0, 0x78, 0xbf, 0x03,
	0xc6, 0x1f, 0xf8, 0x35, 0x10, 0x9c, 0x71, 0xe0, 0xcd, 0xb8, 0x36, 0x96, 0xbc, 0xa7, 0xd6, 0x30,
	0x7d, 0x61, 0xdb, 0xeb, 0xd0, 0x11, 0x93, 0x54, 0x76, 0x8f, 0x1f, 0xfa, 0xb8, 0xd2, 0x10, 0x9c,
	0x76, 0xe0, 0xf5, 0xb8, 0xbe, 0xfb, 0x4b, 0x44, 0xf6, 0x3a, 0x7d, 0x95, 0x97, 0x22, 0x75, 0x8d,
	0x4a, 0xca, 0xa4, 0x2a, 0xea, 0xd4, 0x35, 0x52, 0x8a, 0x65, 0x0b, 0x1e, 0x55, 0x1a, 0x49, 0x7c,
	0xcb, 0x2b, 0x91, 0x8d, 0xf0, 0x3d, 0xa0, 0x5a, 0x44, 0x78, 0x08, 0xf2, 0x5a, 0xe4, 0x60, 0xd9,
	0x8d, 0xa8, 0x96, 0x65, 0x52, 0xad, 0xea, 0x39, 0xc9, 0x5b, 0x21, 0x6c, 0x87, 0xc0, 0xd8, 0xbc,
	0x03, 0xab, 0x2c, 0x4e, 0x8a, 0x99, 0x6c, 0x78, 0xaa, 0x87, 0xd0, 0x1c, 0xeb, 0xfc, 0x50, 0xcf,
	0x64, 0xc3, 0xf2, 0x46, 0xac, 0x42, 0x47, 0x9f, 0x68, 0xb9, 0x57, 0x17, 0xe5, 0xa2, 0x2a, 0xea,
	0x53, 0x7e, 0x7a, 0x79, 0x7b, 0xde, 0x39, 0xde, 0x0f, 0x5b, 0x6d, 0xc9, 0x9b, 0xe9, 0xb8, 0x93,
	0x12, 0x73, 0x96, 0xa6, 0x6d, 0x1e, 0x05, 0x3d, 0xfc, 0x07, 0x00, 0x00, 0xff, 0xff, 0x17, 0x26,
	0xa4, 0x8a, 0x5e, 0x01, 0x00, 0x00,
}
