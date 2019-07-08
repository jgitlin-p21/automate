// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/config/ingest/config_request.proto

package ingest

import (
	fmt "fmt"
	shared "github.com/chef/automate/api/config/shared"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-a2-config/api/a2conf"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type ConfigRequest struct {
	V1                   *ConfigRequest_V1 `protobuf:"bytes,3,opt,name=v1,proto3" json:"v1,omitempty" toml:"v1,omitempty" mapstructure:"v1,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte            `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32             `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest) Reset()         { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()    {}
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8c399c32386e790, []int{0}
}

func (m *ConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest.Unmarshal(m, b)
}
func (m *ConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest.Marshal(b, m, deterministic)
}
func (m *ConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest.Merge(m, src)
}
func (m *ConfigRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest.Size(m)
}
func (m *ConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest proto.InternalMessageInfo

func (m *ConfigRequest) GetV1() *ConfigRequest_V1 {
	if m != nil {
		return m.V1
	}
	return nil
}

type ConfigRequest_V1 struct {
	Sys                  *ConfigRequest_V1_System  `protobuf:"bytes,1,opt,name=sys,proto3" json:"sys,omitempty" toml:"sys,omitempty" mapstructure:"sys,omitempty"`
	Svc                  *ConfigRequest_V1_Service `protobuf:"bytes,2,opt,name=svc,proto3" json:"svc,omitempty" toml:"svc,omitempty" mapstructure:"svc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                    `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                     `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1) Reset()         { *m = ConfigRequest_V1{} }
func (m *ConfigRequest_V1) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1) ProtoMessage()    {}
func (*ConfigRequest_V1) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8c399c32386e790, []int{0, 0}
}

func (m *ConfigRequest_V1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1.Unmarshal(m, b)
}
func (m *ConfigRequest_V1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1.Merge(m, src)
}
func (m *ConfigRequest_V1) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1.Size(m)
}
func (m *ConfigRequest_V1) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1 proto.InternalMessageInfo

func (m *ConfigRequest_V1) GetSys() *ConfigRequest_V1_System {
	if m != nil {
		return m.Sys
	}
	return nil
}

func (m *ConfigRequest_V1) GetSvc() *ConfigRequest_V1_Service {
	if m != nil {
		return m.Svc
	}
	return nil
}

type ConfigRequest_V1_System struct {
	Mlsa                 *shared.Mlsa                     `protobuf:"bytes,1,opt,name=mlsa,proto3" json:"mlsa,omitempty" toml:"mlsa,omitempty" mapstructure:"mlsa,omitempty"`
	Service              *ConfigRequest_V1_System_Service `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty" toml:"service,omitempty" mapstructure:"service,omitempty"`
	Tls                  *shared.TLSCredentials           `protobuf:"bytes,3,opt,name=tls,proto3" json:"tls,omitempty" toml:"tls,omitempty" mapstructure:"tls,omitempty"`
	Log                  *ConfigRequest_V1_System_Log     `protobuf:"bytes,4,opt,name=log,proto3" json:"log,omitempty" toml:"log,omitempty" mapstructure:"log,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                           `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                            `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System) Reset()         { *m = ConfigRequest_V1_System{} }
func (m *ConfigRequest_V1_System) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System) ProtoMessage()    {}
func (*ConfigRequest_V1_System) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8c399c32386e790, []int{0, 0, 0}
}

func (m *ConfigRequest_V1_System) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System.Merge(m, src)
}
func (m *ConfigRequest_V1_System) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System.Size(m)
}
func (m *ConfigRequest_V1_System) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System proto.InternalMessageInfo

func (m *ConfigRequest_V1_System) GetMlsa() *shared.Mlsa {
	if m != nil {
		return m.Mlsa
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetService() *ConfigRequest_V1_System_Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetTls() *shared.TLSCredentials {
	if m != nil {
		return m.Tls
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetLog() *ConfigRequest_V1_System_Log {
	if m != nil {
		return m.Log
	}
	return nil
}

type ConfigRequest_V1_System_Service struct {
	Host                          *wrappers.StringValue `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty" toml:"host,omitempty" mapstructure:"host,omitempty"`
	Port                          *wrappers.Int32Value  `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty" toml:"port,omitempty" mapstructure:"port,omitempty"`
	PurgeConvergeHistoryAfterDays *wrappers.Int32Value  `protobuf:"bytes,3,opt,name=purge_converge_history_after_days,json=purgeConvergeHistoryAfterDays,proto3" json:"purge_converge_history_after_days,omitempty" toml:"purge_converge_history_after_days,omitempty" mapstructure:"purge_converge_history_after_days,omitempty"`
	PurgeActionsAfterDays         *wrappers.Int32Value  `protobuf:"bytes,4,opt,name=purge_actions_after_days,json=purgeActionsAfterDays,proto3" json:"purge_actions_after_days,omitempty" toml:"purge_actions_after_days,omitempty" mapstructure:"purge_actions_after_days,omitempty"`
	MaxNumberOfBundledRunMsgs     *wrappers.Int32Value  `protobuf:"bytes,6,opt,name=max_number_of_bundled_run_msgs,json=maxNumberOfBundledRunMsgs,proto3" json:"max_number_of_bundled_run_msgs,omitempty" toml:"max_number_of_bundled_run_msgs,omitempty" mapstructure:"max_number_of_bundled_run_msgs,omitempty"`
	MaxNumberOfBundledActionMsgs  *wrappers.Int32Value  `protobuf:"bytes,7,opt,name=max_number_of_bundled_action_msgs,json=maxNumberOfBundledActionMsgs,proto3" json:"max_number_of_bundled_action_msgs,omitempty" toml:"max_number_of_bundled_action_msgs,omitempty" mapstructure:"max_number_of_bundled_action_msgs,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized              []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache                 int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Service) Reset()         { *m = ConfigRequest_V1_System_Service{} }
func (m *ConfigRequest_V1_System_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8c399c32386e790, []int{0, 0, 0, 0}
}

func (m *ConfigRequest_V1_System_Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Service.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Service) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Size(m)
}
func (m *ConfigRequest_V1_System_Service) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Service.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Service proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Service) GetHost() *wrappers.StringValue {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetPort() *wrappers.Int32Value {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetPurgeConvergeHistoryAfterDays() *wrappers.Int32Value {
	if m != nil {
		return m.PurgeConvergeHistoryAfterDays
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetPurgeActionsAfterDays() *wrappers.Int32Value {
	if m != nil {
		return m.PurgeActionsAfterDays
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetMaxNumberOfBundledRunMsgs() *wrappers.Int32Value {
	if m != nil {
		return m.MaxNumberOfBundledRunMsgs
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetMaxNumberOfBundledActionMsgs() *wrappers.Int32Value {
	if m != nil {
		return m.MaxNumberOfBundledActionMsgs
	}
	return nil
}

type ConfigRequest_V1_System_Log struct {
	Format               *wrappers.StringValue `protobuf:"bytes,1,opt,name=format,proto3" json:"format,omitempty" toml:"format,omitempty" mapstructure:"format,omitempty"`
	Level                *wrappers.StringValue `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty" toml:"level,omitempty" mapstructure:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Log) Reset()         { *m = ConfigRequest_V1_System_Log{} }
func (m *ConfigRequest_V1_System_Log) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Log) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8c399c32386e790, []int{0, 0, 0, 1}
}

func (m *ConfigRequest_V1_System_Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Log.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Log.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Log.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Log) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Log.Size(m)
}
func (m *ConfigRequest_V1_System_Log) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Log.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Log proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Log) GetFormat() *wrappers.StringValue {
	if m != nil {
		return m.Format
	}
	return nil
}

func (m *ConfigRequest_V1_System_Log) GetLevel() *wrappers.StringValue {
	if m != nil {
		return m.Level
	}
	return nil
}

type ConfigRequest_V1_Service struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_Service) Reset()         { *m = ConfigRequest_V1_Service{} }
func (m *ConfigRequest_V1_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8c399c32386e790, []int{0, 0, 1}
}

func (m *ConfigRequest_V1_Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_Service.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_Service.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_Service.Merge(m, src)
}
func (m *ConfigRequest_V1_Service) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_Service.Size(m)
}
func (m *ConfigRequest_V1_Service) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_Service.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_Service proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ConfigRequest)(nil), "chef.automate.domain.ingest.ConfigRequest")
	proto.RegisterType((*ConfigRequest_V1)(nil), "chef.automate.domain.ingest.ConfigRequest.V1")
	proto.RegisterType((*ConfigRequest_V1_System)(nil), "chef.automate.domain.ingest.ConfigRequest.V1.System")
	proto.RegisterType((*ConfigRequest_V1_System_Service)(nil), "chef.automate.domain.ingest.ConfigRequest.V1.System.Service")
	proto.RegisterType((*ConfigRequest_V1_System_Log)(nil), "chef.automate.domain.ingest.ConfigRequest.V1.System.Log")
	proto.RegisterType((*ConfigRequest_V1_Service)(nil), "chef.automate.domain.ingest.ConfigRequest.V1.Service")
}

func init() {
	proto.RegisterFile("api/config/ingest/config_request.proto", fileDescriptor_a8c399c32386e790)
}

var fileDescriptor_a8c399c32386e790 = []byte{
	// 654 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xd1, 0x6e, 0xd3, 0x3a,
	0x18, 0xc7, 0xb5, 0x26, 0x6b, 0x37, 0x1f, 0x9d, 0xa3, 0xca, 0x3a, 0xe7, 0x28, 0x64, 0x63, 0xda,
	0xb8, 0x40, 0x68, 0xa2, 0x0e, 0xed, 0x86, 0x84, 0xd0, 0xb8, 0xd8, 0x86, 0x18, 0x54, 0x1b, 0x93,
	0xb2, 0xa9, 0x17, 0x5c, 0x10, 0xb9, 0x89, 0xe3, 0x46, 0x4a, 0xec, 0x60, 0x3b, 0x65, 0x7d, 0x05,
	0x24, 0x1e, 0x82, 0x47, 0xe0, 0x15, 0x26, 0xde, 0x68, 0x2f, 0x80, 0x62, 0xbb, 0x65, 0x50, 0xd4,
	0xb1, 0xdd, 0x35, 0xfa, 0xbe, 0xff, 0xef, 0xff, 0x7d, 0x9f, 0xfe, 0x35, 0x78, 0x88, 0xcb, 0x2c,
	0x88, 0x39, 0x4b, 0x33, 0x1a, 0x64, 0x8c, 0x12, 0xa9, 0xec, 0x57, 0x24, 0xc8, 0x87, 0x8a, 0x48,
	0x85, 0x4a, 0xc1, 0x15, 0x87, 0x6b, 0xf1, 0x88, 0xa4, 0x08, 0x57, 0x8a, 0x17, 0x58, 0x11, 0x94,
	0xf0, 0x02, 0x67, 0x0c, 0x19, 0x85, 0xbf, 0x71, 0x0d, 0x22, 0x47, 0x58, 0x90, 0x24, 0xa0, 0x39,
	0x1f, 0xe2, 0xdc, 0x88, 0xfd, 0xb5, 0xf9, 0xba, 0xca, 0xa5, 0x2d, 0xf6, 0x63, 0x5e, 0x94, 0x9c,
	0x11, 0xa6, 0x64, 0x30, 0xe5, 0x77, 0xa8, 0x28, 0xe3, 0x40, 0xd7, 0xe3, 0x0e, 0x25, 0xac, 0x83,
	0x7b, 0x1d, 0xab, 0xaf, 0x51, 0xb8, 0x57, 0x7f, 0x04, 0x98, 0x31, 0xae, 0xb0, 0xca, 0x38, 0x9b,
	0xb2, 0x36, 0x28, 0xe7, 0x34, 0x27, 0x46, 0x39, 0xac, 0xd2, 0xe0, 0xa3, 0xc0, 0x65, 0x49, 0x84,
	0xad, 0x3f, 0xf8, 0xba, 0x0a, 0xfe, 0x3e, 0xd4, 0x9c, 0xd0, 0x6c, 0x07, 0x5f, 0x80, 0xc6, 0xb8,
	0xeb, 0x39, 0x9b, 0x4b, 0x8f, 0xfe, 0xea, 0x75, 0xd0, 0x82, 0x25, 0xd1, 0x4f, 0x3a, 0x34, 0xe8,
	0x86, 0x8d, 0x71, 0xd7, 0xff, 0xb6, 0x02, 0x1a, 0x83, 0x2e, 0x7c, 0x05, 0x1c, 0x39, 0x91, 0xde,
	0x92, 0xc6, 0xec, 0xde, 0x0a, 0x83, 0xce, 0x26, 0x52, 0x91, 0x22, 0xac, 0x01, 0xf0, 0x08, 0x38,
	0x72, 0x1c, 0x7b, 0x0d, 0xcd, 0x79, 0x7a, 0x4b, 0x0e, 0x11, 0xe3, 0x2c, 0x26, 0x61, 0x4d, 0xf0,
	0xbf, 0xb4, 0x40, 0xd3, 0x80, 0xe1, 0x2e, 0x70, 0x8b, 0x5c, 0x62, 0x3b, 0xdc, 0xe6, 0x2f, 0xd0,
	0x8c, 0xa5, 0x02, 0x23, 0x73, 0x5b, 0x74, 0x92, 0x4b, 0x1c, 0xea, 0x6e, 0x38, 0x00, 0x2d, 0x69,
	0x80, 0x76, 0x9a, 0xbd, 0xbb, 0x6c, 0x35, 0x1b, 0x6a, 0x0a, 0x83, 0x7b, 0xc0, 0x51, 0xb9, 0xb4,
	0x07, 0xdf, 0x5e, 0x34, 0xcc, 0xf9, 0xf1, 0xd9, 0xa1, 0x20, 0x09, 0x61, 0x2a, 0xc3, 0xb9, 0x0c,
	0x6b, 0x19, 0xec, 0x03, 0x27, 0xe7, 0xd4, 0x73, 0xb5, 0xfa, 0xd9, 0x9d, 0x26, 0x3a, 0xe6, 0x34,
	0xac, 0x21, 0xfe, 0x67, 0x17, 0xb4, 0xec, 0x78, 0xf0, 0x09, 0x70, 0x47, 0x5c, 0x2a, 0x7b, 0xa3,
	0x75, 0x64, 0x62, 0x84, 0xa6, 0x31, 0x42, 0x67, 0x4a, 0x64, 0x8c, 0x0e, 0x70, 0x5e, 0x91, 0x50,
	0x77, 0xc2, 0x23, 0xe0, 0x96, 0x5c, 0x28, 0x7b, 0x9c, 0xb5, 0x39, 0xc5, 0x1b, 0xa6, 0x76, 0x7a,
	0x5a, 0x70, 0xf0, 0xff, 0xe5, 0x95, 0x07, 0x67, 0xe7, 0x6c, 0x7f, 0x3a, 0xf5, 0xdd, 0x3a, 0xde,
	0xa1, 0x06, 0x40, 0x02, 0xb6, 0xca, 0x4a, 0x50, 0x12, 0xc5, 0x9c, 0x8d, 0x49, 0xfd, 0x63, 0x94,
	0x49, 0xc5, 0xc5, 0x24, 0xc2, 0xa9, 0x22, 0x22, 0x4a, 0xf0, 0x64, 0x7a, 0xae, 0x45, 0x2e, 0xe1,
	0x7d, 0x4d, 0x39, 0xb4, 0x90, 0xd7, 0x86, 0xb1, 0x5f, 0x23, 0x5e, 0xe2, 0x89, 0x84, 0xe7, 0xc0,
	0x33, 0x36, 0x38, 0xd6, 0x7f, 0x98, 0xeb, 0x74, 0xf7, 0x66, 0xfa, 0x7f, 0x5a, 0xbc, 0x6f, 0xb4,
	0x3f, 0xa8, 0xef, 0xc1, 0x46, 0x81, 0x2f, 0x22, 0x56, 0x15, 0x43, 0x22, 0x22, 0x9e, 0x46, 0xc3,
	0x8a, 0x25, 0x39, 0x49, 0x22, 0x51, 0xb1, 0xa8, 0x90, 0x54, 0x7a, 0xcd, 0x9b, 0xd9, 0xf7, 0x0a,
	0x7c, 0xf1, 0x56, 0x13, 0x4e, 0xd3, 0x03, 0xa3, 0x0f, 0x2b, 0x76, 0x22, 0xa9, 0x84, 0x09, 0xd8,
	0xfa, 0x3d, 0xdf, 0x6c, 0x61, 0x2c, 0x5a, 0x37, 0x5b, 0xac, 0xcf, 0x5b, 0x98, 0x5d, 0x6a, 0x97,
	0xbe, 0xbb, 0xb2, 0xdc, 0x6e, 0xfa, 0x1c, 0x38, 0xc7, 0x9c, 0xc2, 0x5d, 0xd0, 0x4c, 0xb9, 0x28,
	0xf0, 0x9f, 0x85, 0xc1, 0xf6, 0xc2, 0x1e, 0x58, 0xce, 0xc9, 0x98, 0xe4, 0x36, 0x0f, 0x8b, 0x45,
	0xa6, 0xd5, 0x5f, 0x9d, 0xe5, 0xef, 0xf9, 0xbf, 0x97, 0x57, 0x5e, 0x1b, 0xfc, 0x63, 0xa2, 0xdb,
	0xb1, 0x49, 0xe9, 0xbb, 0x2b, 0x4b, 0x6d, 0xe7, 0xe0, 0xf1, 0xbb, 0x6d, 0x9a, 0xa9, 0x51, 0x35,
	0x44, 0x31, 0x2f, 0x82, 0x3a, 0xf2, 0xb3, 0x67, 0x32, 0x98, 0x7b, 0xbc, 0x87, 0x4d, 0xed, 0xb8,
	0xf3, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x77, 0x9e, 0x82, 0x2f, 0xd8, 0x05, 0x00, 0x00,
}
