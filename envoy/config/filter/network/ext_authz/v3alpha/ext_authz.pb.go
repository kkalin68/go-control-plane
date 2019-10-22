// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/ext_authz/v3alpha/ext_authz.proto

package envoy_config_filter_network_ext_authz_v3alpha

import (
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type ExtAuthz struct {
	StatPrefix           string            `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	GrpcService          *core.GrpcService `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	FailureModeAllow     bool              `protobuf:"varint,3,opt,name=failure_mode_allow,json=failureModeAllow,proto3" json:"failure_mode_allow,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ExtAuthz) Reset()         { *m = ExtAuthz{} }
func (m *ExtAuthz) String() string { return proto.CompactTextString(m) }
func (*ExtAuthz) ProtoMessage()    {}
func (*ExtAuthz) Descriptor() ([]byte, []int) {
	return fileDescriptor_c588f640e17fc55c, []int{0}
}

func (m *ExtAuthz) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtAuthz.Unmarshal(m, b)
}
func (m *ExtAuthz) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtAuthz.Marshal(b, m, deterministic)
}
func (m *ExtAuthz) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtAuthz.Merge(m, src)
}
func (m *ExtAuthz) XXX_Size() int {
	return xxx_messageInfo_ExtAuthz.Size(m)
}
func (m *ExtAuthz) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtAuthz.DiscardUnknown(m)
}

var xxx_messageInfo_ExtAuthz proto.InternalMessageInfo

func (m *ExtAuthz) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *ExtAuthz) GetGrpcService() *core.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func (m *ExtAuthz) GetFailureModeAllow() bool {
	if m != nil {
		return m.FailureModeAllow
	}
	return false
}

func init() {
	proto.RegisterType((*ExtAuthz)(nil), "envoy.config.filter.network.ext_authz.v3alpha.ExtAuthz")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/ext_authz/v3alpha/ext_authz.proto", fileDescriptor_c588f640e17fc55c)
}

var fileDescriptor_c588f640e17fc55c = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xc1, 0x4e, 0xbc, 0x30,
	0x10, 0xc6, 0xd3, 0xfd, 0xff, 0xa3, 0x6b, 0xd1, 0xc4, 0x70, 0x91, 0xec, 0x89, 0xe8, 0x05, 0x13,
	0x6d, 0x13, 0xf7, 0x68, 0x3c, 0x2c, 0x89, 0x7a, 0xd2, 0x10, 0x7c, 0x00, 0x52, 0x61, 0x60, 0x1b,
	0x2b, 0x6d, 0x66, 0xbb, 0x2c, 0xeb, 0x23, 0xf9, 0x88, 0x9e, 0x0c, 0x14, 0xc4, 0xab, 0xb7, 0x76,
	0xbe, 0xf9, 0xcd, 0x37, 0xdf, 0xd0, 0x3b, 0xa8, 0x1b, 0xbd, 0xe7, 0xb9, 0xae, 0x4b, 0x59, 0xf1,
	0x52, 0x2a, 0x0b, 0xc8, 0x6b, 0xb0, 0x3b, 0x8d, 0x6f, 0x1c, 0x5a, 0x9b, 0x89, 0xad, 0x5d, 0x7f,
	0xf0, 0x66, 0x29, 0x94, 0x59, 0x8b, 0xa9, 0xc2, 0x0c, 0x6a, 0xab, 0xfd, 0xeb, 0x1e, 0x67, 0x0e,
	0x67, 0x0e, 0x67, 0x03, 0xce, 0xa6, 0xe6, 0x01, 0x5f, 0x5c, 0x3a, 0x37, 0x61, 0xe4, 0xcf, 0xc4,
	0x5c, 0x23, 0xf0, 0x0a, 0x4d, 0x9e, 0x6d, 0x00, 0x1b, 0x99, 0x83, 0x9b, 0xbc, 0x38, 0x6b, 0x84,
	0x92, 0x85, 0xb0, 0xc0, 0xc7, 0x87, 0x13, 0xce, 0x3f, 0x09, 0x9d, 0xdf, 0xb7, 0x76, 0xd5, 0x0d,
	0xf6, 0x23, 0xea, 0x6d, 0xac, 0xb0, 0x99, 0x41, 0x28, 0x65, 0x1b, 0x90, 0x90, 0x44, 0x47, 0xf1,
	0xe1, 0x57, 0xfc, 0x1f, 0x67, 0x21, 0x49, 0x69, 0xa7, 0x25, 0xbd, 0xe4, 0x3f, 0xd0, 0xe3, 0xdf,
	0x2e, 0xc1, 0x2c, 0x24, 0x91, 0x77, 0x73, 0xc1, 0x5c, 0x00, 0x61, 0xe4, 0xb8, 0x24, 0xeb, 0x36,
	0x62, 0x8f, 0x68, 0xf2, 0x17, 0xd7, 0x9a, 0x7a, 0xd5, 0xf4, 0xf1, 0xaf, 0xa8, 0x5f, 0x0a, 0xa9,
	0xb6, 0x08, 0xd9, 0xbb, 0x2e, 0x20, 0x13, 0x4a, 0xe9, 0x5d, 0xf0, 0x2f, 0x24, 0xd1, 0x3c, 0x3d,
	0x1d, 0x94, 0x27, 0x5d, 0xc0, 0xaa, 0xab, 0xc7, 0xcf, 0xf4, 0x56, 0x6a, 0xe7, 0x61, 0x50, 0xb7,
	0x7b, 0xf6, 0xa7, 0x7b, 0xc5, 0x27, 0x63, 0xd0, 0xa4, 0x8b, 0x9e, 0x90, 0xd7, 0x83, 0xfe, 0x06,
	0xcb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x77, 0x45, 0x7b, 0xb7, 0x01, 0x00, 0x00,
}