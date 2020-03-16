// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/network/tcp_proxy/v3/tcp_proxy.proto

package envoy_extensions_filters_network_tcp_proxy_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v31 "github.com/envoyproxy/go-control-plane/envoy/config/accesslog/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v32 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type TcpProxy struct {
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Types that are valid to be assigned to ClusterSpecifier:
	//	*TcpProxy_Cluster
	//	*TcpProxy_WeightedClusters
	ClusterSpecifier                  isTcpProxy_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
	MetadataMatch                     *v3.Metadata                `protobuf:"bytes,9,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	IdleTimeout                       *duration.Duration          `protobuf:"bytes,8,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	DownstreamIdleTimeout             *duration.Duration          `protobuf:"bytes,3,opt,name=downstream_idle_timeout,json=downstreamIdleTimeout,proto3" json:"downstream_idle_timeout,omitempty"`
	UpstreamIdleTimeout               *duration.Duration          `protobuf:"bytes,4,opt,name=upstream_idle_timeout,json=upstreamIdleTimeout,proto3" json:"upstream_idle_timeout,omitempty"`
	AccessLog                         []*v31.AccessLog            `protobuf:"bytes,5,rep,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	HiddenEnvoyDeprecatedDeprecatedV1 *TcpProxy_DeprecatedV1      `protobuf:"bytes,6,opt,name=hidden_envoy_deprecated_deprecated_v1,json=hiddenEnvoyDeprecatedDeprecatedV1,proto3" json:"hidden_envoy_deprecated_deprecated_v1,omitempty"` // Deprecated: Do not use.
	MaxConnectAttempts                *wrappers.UInt32Value       `protobuf:"bytes,7,opt,name=max_connect_attempts,json=maxConnectAttempts,proto3" json:"max_connect_attempts,omitempty"`
	HashPolicy                        []*v32.HashPolicy           `protobuf:"bytes,11,rep,name=hash_policy,json=hashPolicy,proto3" json:"hash_policy,omitempty"`
	TunnelingConfig                   *TcpProxy_TunnelingConfig   `protobuf:"bytes,12,opt,name=tunneling_config,json=tunnelingConfig,proto3" json:"tunneling_config,omitempty"`
	XXX_NoUnkeyedLiteral              struct{}                    `json:"-"`
	XXX_unrecognized                  []byte                      `json:"-"`
	XXX_sizecache                     int32                       `json:"-"`
}

func (m *TcpProxy) Reset()         { *m = TcpProxy{} }
func (m *TcpProxy) String() string { return proto.CompactTextString(m) }
func (*TcpProxy) ProtoMessage()    {}
func (*TcpProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_926036763459e38f, []int{0}
}

func (m *TcpProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy.Unmarshal(m, b)
}
func (m *TcpProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy.Marshal(b, m, deterministic)
}
func (m *TcpProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy.Merge(m, src)
}
func (m *TcpProxy) XXX_Size() int {
	return xxx_messageInfo_TcpProxy.Size(m)
}
func (m *TcpProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy proto.InternalMessageInfo

func (m *TcpProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

type isTcpProxy_ClusterSpecifier interface {
	isTcpProxy_ClusterSpecifier()
}

type TcpProxy_Cluster struct {
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3,oneof"`
}

type TcpProxy_WeightedClusters struct {
	WeightedClusters *TcpProxy_WeightedCluster `protobuf:"bytes,10,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

func (*TcpProxy_Cluster) isTcpProxy_ClusterSpecifier() {}

func (*TcpProxy_WeightedClusters) isTcpProxy_ClusterSpecifier() {}

func (m *TcpProxy) GetClusterSpecifier() isTcpProxy_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (m *TcpProxy) GetCluster() string {
	if x, ok := m.GetClusterSpecifier().(*TcpProxy_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *TcpProxy) GetWeightedClusters() *TcpProxy_WeightedCluster {
	if x, ok := m.GetClusterSpecifier().(*TcpProxy_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

func (m *TcpProxy) GetMetadataMatch() *v3.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func (m *TcpProxy) GetIdleTimeout() *duration.Duration {
	if m != nil {
		return m.IdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetDownstreamIdleTimeout() *duration.Duration {
	if m != nil {
		return m.DownstreamIdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetUpstreamIdleTimeout() *duration.Duration {
	if m != nil {
		return m.UpstreamIdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetAccessLog() []*v31.AccessLog {
	if m != nil {
		return m.AccessLog
	}
	return nil
}

// Deprecated: Do not use.
func (m *TcpProxy) GetHiddenEnvoyDeprecatedDeprecatedV1() *TcpProxy_DeprecatedV1 {
	if m != nil {
		return m.HiddenEnvoyDeprecatedDeprecatedV1
	}
	return nil
}

func (m *TcpProxy) GetMaxConnectAttempts() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConnectAttempts
	}
	return nil
}

func (m *TcpProxy) GetHashPolicy() []*v32.HashPolicy {
	if m != nil {
		return m.HashPolicy
	}
	return nil
}

func (m *TcpProxy) GetTunnelingConfig() *TcpProxy_TunnelingConfig {
	if m != nil {
		return m.TunnelingConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TcpProxy) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TcpProxy_Cluster)(nil),
		(*TcpProxy_WeightedClusters)(nil),
	}
}

// Deprecated: Do not use.
type TcpProxy_DeprecatedV1 struct {
	Routes               []*TcpProxy_DeprecatedV1_TCPRoute `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *TcpProxy_DeprecatedV1) Reset()         { *m = TcpProxy_DeprecatedV1{} }
func (m *TcpProxy_DeprecatedV1) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_DeprecatedV1) ProtoMessage()    {}
func (*TcpProxy_DeprecatedV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_926036763459e38f, []int{0, 0}
}

func (m *TcpProxy_DeprecatedV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_DeprecatedV1.Unmarshal(m, b)
}
func (m *TcpProxy_DeprecatedV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_DeprecatedV1.Marshal(b, m, deterministic)
}
func (m *TcpProxy_DeprecatedV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_DeprecatedV1.Merge(m, src)
}
func (m *TcpProxy_DeprecatedV1) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_DeprecatedV1.Size(m)
}
func (m *TcpProxy_DeprecatedV1) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_DeprecatedV1.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_DeprecatedV1 proto.InternalMessageInfo

func (m *TcpProxy_DeprecatedV1) GetRoutes() []*TcpProxy_DeprecatedV1_TCPRoute {
	if m != nil {
		return m.Routes
	}
	return nil
}

type TcpProxy_DeprecatedV1_TCPRoute struct {
	Cluster              string          `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	DestinationIpList    []*v3.CidrRange `protobuf:"bytes,2,rep,name=destination_ip_list,json=destinationIpList,proto3" json:"destination_ip_list,omitempty"`
	DestinationPorts     string          `protobuf:"bytes,3,opt,name=destination_ports,json=destinationPorts,proto3" json:"destination_ports,omitempty"`
	SourceIpList         []*v3.CidrRange `protobuf:"bytes,4,rep,name=source_ip_list,json=sourceIpList,proto3" json:"source_ip_list,omitempty"`
	SourcePorts          string          `protobuf:"bytes,5,opt,name=source_ports,json=sourcePorts,proto3" json:"source_ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) Reset()         { *m = TcpProxy_DeprecatedV1_TCPRoute{} }
func (m *TcpProxy_DeprecatedV1_TCPRoute) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_DeprecatedV1_TCPRoute) ProtoMessage()    {}
func (*TcpProxy_DeprecatedV1_TCPRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_926036763459e38f, []int{0, 0, 0}
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.Unmarshal(m, b)
}
func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.Marshal(b, m, deterministic)
}
func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.Merge(m, src)
}
func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.Size(m)
}
func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute proto.InternalMessageInfo

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetDestinationIpList() []*v3.CidrRange {
	if m != nil {
		return m.DestinationIpList
	}
	return nil
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetDestinationPorts() string {
	if m != nil {
		return m.DestinationPorts
	}
	return ""
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetSourceIpList() []*v3.CidrRange {
	if m != nil {
		return m.SourceIpList
	}
	return nil
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetSourcePorts() string {
	if m != nil {
		return m.SourcePorts
	}
	return ""
}

type TcpProxy_WeightedCluster struct {
	Clusters             []*TcpProxy_WeightedCluster_ClusterWeight `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *TcpProxy_WeightedCluster) Reset()         { *m = TcpProxy_WeightedCluster{} }
func (m *TcpProxy_WeightedCluster) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_WeightedCluster) ProtoMessage()    {}
func (*TcpProxy_WeightedCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_926036763459e38f, []int{0, 1}
}

func (m *TcpProxy_WeightedCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_WeightedCluster.Unmarshal(m, b)
}
func (m *TcpProxy_WeightedCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_WeightedCluster.Marshal(b, m, deterministic)
}
func (m *TcpProxy_WeightedCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_WeightedCluster.Merge(m, src)
}
func (m *TcpProxy_WeightedCluster) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_WeightedCluster.Size(m)
}
func (m *TcpProxy_WeightedCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_WeightedCluster.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_WeightedCluster proto.InternalMessageInfo

func (m *TcpProxy_WeightedCluster) GetClusters() []*TcpProxy_WeightedCluster_ClusterWeight {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type TcpProxy_WeightedCluster_ClusterWeight struct {
	Name                 string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Weight               uint32       `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	MetadataMatch        *v3.Metadata `protobuf:"bytes,3,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) Reset() {
	*m = TcpProxy_WeightedCluster_ClusterWeight{}
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_WeightedCluster_ClusterWeight) ProtoMessage()    {}
func (*TcpProxy_WeightedCluster_ClusterWeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_926036763459e38f, []int{0, 1, 0}
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Unmarshal(m, b)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Marshal(b, m, deterministic)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Merge(m, src)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Size(m)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight proto.InternalMessageInfo

func (m *TcpProxy_WeightedCluster_ClusterWeight) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) GetMetadataMatch() *v3.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

type TcpProxy_TunnelingConfig struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcpProxy_TunnelingConfig) Reset()         { *m = TcpProxy_TunnelingConfig{} }
func (m *TcpProxy_TunnelingConfig) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_TunnelingConfig) ProtoMessage()    {}
func (*TcpProxy_TunnelingConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_926036763459e38f, []int{0, 2}
}

func (m *TcpProxy_TunnelingConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_TunnelingConfig.Unmarshal(m, b)
}
func (m *TcpProxy_TunnelingConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_TunnelingConfig.Marshal(b, m, deterministic)
}
func (m *TcpProxy_TunnelingConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_TunnelingConfig.Merge(m, src)
}
func (m *TcpProxy_TunnelingConfig) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_TunnelingConfig.Size(m)
}
func (m *TcpProxy_TunnelingConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_TunnelingConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_TunnelingConfig proto.InternalMessageInfo

func (m *TcpProxy_TunnelingConfig) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func init() {
	proto.RegisterType((*TcpProxy)(nil), "envoy.extensions.filters.network.tcp_proxy.v3.TcpProxy")
	proto.RegisterType((*TcpProxy_DeprecatedV1)(nil), "envoy.extensions.filters.network.tcp_proxy.v3.TcpProxy.DeprecatedV1")
	proto.RegisterType((*TcpProxy_DeprecatedV1_TCPRoute)(nil), "envoy.extensions.filters.network.tcp_proxy.v3.TcpProxy.DeprecatedV1.TCPRoute")
	proto.RegisterType((*TcpProxy_WeightedCluster)(nil), "envoy.extensions.filters.network.tcp_proxy.v3.TcpProxy.WeightedCluster")
	proto.RegisterType((*TcpProxy_WeightedCluster_ClusterWeight)(nil), "envoy.extensions.filters.network.tcp_proxy.v3.TcpProxy.WeightedCluster.ClusterWeight")
	proto.RegisterType((*TcpProxy_TunnelingConfig)(nil), "envoy.extensions.filters.network.tcp_proxy.v3.TcpProxy.TunnelingConfig")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/tcp_proxy/v3/tcp_proxy.proto", fileDescriptor_926036763459e38f)
}

var fileDescriptor_926036763459e38f = []byte{
	// 977 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcb, 0x6e, 0x1b, 0x37,
	0x14, 0xcd, 0xc8, 0x2f, 0xf9, 0xca, 0xae, 0x6d, 0xa6, 0x41, 0x26, 0x6a, 0xe1, 0x47, 0xda, 0x02,
	0x6e, 0x8b, 0xce, 0xc0, 0xd6, 0xa6, 0x50, 0x1f, 0xa8, 0x25, 0x07, 0xb6, 0x81, 0x38, 0x51, 0x06,
	0x8e, 0xb3, 0x1c, 0xd0, 0x33, 0x94, 0x44, 0x74, 0x44, 0x0e, 0x48, 0x8e, 0x24, 0xa3, 0x9b, 0x2e,
	0xbb, 0xee, 0x32, 0x5f, 0xd1, 0x65, 0x7f, 0xa0, 0xe8, 0x97, 0x74, 0xd7, 0x0f, 0x28, 0xb2, 0x2a,
	0x48, 0xce, 0xe8, 0x55, 0xb5, 0xae, 0xe3, 0xac, 0x66, 0x78, 0x1f, 0xe7, 0x90, 0xe7, 0x5e, 0x5e,
	0xc2, 0x37, 0x84, 0xf5, 0xf9, 0xb5, 0x4f, 0x86, 0x8a, 0x30, 0x49, 0x39, 0x93, 0x7e, 0x9b, 0x26,
	0x8a, 0x08, 0xe9, 0x33, 0xa2, 0x06, 0x5c, 0x7c, 0xef, 0xab, 0x28, 0x0d, 0x53, 0xc1, 0x87, 0xd7,
	0x7e, 0xbf, 0x36, 0x5e, 0x78, 0xa9, 0xe0, 0x8a, 0xa3, 0x2f, 0x4c, 0xba, 0x37, 0x4e, 0xf7, 0xf2,
	0x74, 0x2f, 0x4f, 0xf7, 0xc6, 0x19, 0xfd, 0x5a, 0xf5, 0x53, 0xcb, 0x16, 0x71, 0xd6, 0xa6, 0x1d,
	0x1f, 0x47, 0x11, 0x91, 0x32, 0xe1, 0x1d, 0x8d, 0x3c, 0x5a, 0x58, 0xe4, 0xea, 0xe3, 0xa9, 0xd0,
	0x88, 0x0b, 0x62, 0xa2, 0xe2, 0x58, 0x10, 0x29, 0xf3, 0x98, 0x9d, 0xb9, 0x31, 0x57, 0x58, 0x92,
	0xe9, 0x00, 0x75, 0x9d, 0x1a, 0x4f, 0x17, 0xcb, 0x6e, 0x98, 0xf2, 0x84, 0x46, 0xf9, 0xfe, 0xab,
	0xdb, 0x1d, 0xce, 0x3b, 0x09, 0xf1, 0xcd, 0xea, 0x2a, 0x6b, 0xfb, 0x71, 0x26, 0xb0, 0xa2, 0x9c,
	0xfd, 0x9b, 0x7f, 0x20, 0x70, 0x9a, 0xea, 0xf3, 0x59, 0xff, 0x5e, 0x16, 0xa7, 0xd8, 0xc7, 0x8c,
	0x71, 0x65, 0xd2, 0xa4, 0xdf, 0x27, 0x42, 0x0b, 0x41, 0x59, 0x71, 0x90, 0x87, 0x7d, 0x9c, 0xd0,
	0x18, 0x2b, 0xe2, 0x17, 0x3f, 0xd6, 0xf1, 0xf8, 0xd7, 0x2d, 0x28, 0x5f, 0x44, 0x69, 0x4b, 0x8b,
	0x83, 0xf6, 0xa1, 0x22, 0x15, 0x56, 0x61, 0x2a, 0x48, 0x9b, 0x0e, 0x5d, 0x67, 0xd7, 0xd9, 0x5f,
	0x6d, 0xac, 0xbc, 0x69, 0x2c, 0x8a, 0xd2, 0xae, 0x13, 0x80, 0xf6, 0xb5, 0x8c, 0x0b, 0x55, 0x61,
	0x25, 0x4a, 0x32, 0xa9, 0x88, 0x70, 0x4b, 0x3a, 0xea, 0xf4, 0x5e, 0x50, 0x18, 0x50, 0x1f, 0xb6,
	0x06, 0x84, 0x76, 0xba, 0x8a, 0xc4, 0x61, 0x6e, 0x93, 0x2e, 0xec, 0x3a, 0xfb, 0x95, 0xc3, 0x13,
	0xef, 0x56, 0xa5, 0xf2, 0x8a, 0x9d, 0x79, 0xaf, 0x72, 0xc0, 0xa6, 0xc5, 0x3b, 0xbd, 0x17, 0x6c,
	0x0e, 0xa6, 0x4d, 0x12, 0x3d, 0x81, 0xf7, 0x7a, 0x44, 0xe1, 0x18, 0x2b, 0x1c, 0xf6, 0xb0, 0x8a,
	0xba, 0xee, 0xaa, 0x21, 0xdd, 0xce, 0x49, 0x6d, 0x85, 0x3c, 0x5d, 0x21, 0x8d, 0x7d, 0x9e, 0xc7,
	0x06, 0xeb, 0x45, 0xd6, 0xb9, 0x4e, 0x42, 0x5f, 0xc3, 0x1a, 0x8d, 0x13, 0x12, 0x2a, 0xda, 0x23,
	0x3c, 0x53, 0x6e, 0xd9, 0x80, 0x3c, 0xf2, 0x6c, 0x11, 0xbc, 0xa2, 0x08, 0xde, 0x71, 0x5e, 0xa4,
	0xa0, 0xa2, 0xc3, 0x2f, 0x6c, 0x34, 0x7a, 0x01, 0x0f, 0x63, 0x3e, 0x60, 0x52, 0x09, 0x82, 0x7b,
	0xe1, 0x14, 0xd0, 0xc2, 0x4d, 0x40, 0x0f, 0xc6, 0x99, 0x67, 0x13, 0x90, 0xe7, 0xf0, 0x20, 0x4b,
	0xe7, 0x01, 0x2e, 0xde, 0x04, 0x78, 0xbf, 0xc8, 0x9b, 0x84, 0x6b, 0x02, 0xd8, 0x36, 0x0f, 0x13,
	0xde, 0x71, 0x97, 0x76, 0x17, 0xf6, 0x2b, 0x87, 0x1f, 0x4f, 0x4b, 0x34, 0xbe, 0x06, 0xfd, 0x9a,
	0x77, 0x64, 0x16, 0x4f, 0x79, 0x27, 0x58, 0xc5, 0xc5, 0x2f, 0x7a, 0xed, 0xc0, 0x27, 0x5d, 0x1a,
	0xc7, 0x84, 0x85, 0x26, 0x33, 0x8c, 0x49, 0x2a, 0x48, 0x84, 0x75, 0xcd, 0x27, 0x7e, 0xfb, 0x07,
	0xee, 0xb2, 0xd9, 0xe4, 0xf1, 0xdb, 0x16, 0xfe, 0x78, 0x04, 0x76, 0x79, 0xd0, 0x28, 0xb9, 0x4e,
	0xb0, 0x67, 0x69, 0x9f, 0x68, 0xb8, 0xb1, 0x73, 0x32, 0x0c, 0xbd, 0x82, 0xf7, 0x7b, 0x78, 0x18,
	0x46, 0x9c, 0x31, 0x12, 0xa9, 0x10, 0x2b, 0x45, 0x7a, 0xa9, 0x92, 0xee, 0x8a, 0xd9, 0xca, 0x87,
	0xff, 0xd0, 0xeb, 0xe5, 0x19, 0x53, 0xb5, 0xc3, 0x4b, 0x9c, 0x64, 0xc4, 0x74, 0xfb, 0x67, 0xa5,
	0x7d, 0x27, 0x40, 0x3d, 0x3c, 0x6c, 0x5a, 0x84, 0xa3, 0x1c, 0x00, 0x1d, 0x43, 0x65, 0xe2, 0xf6,
	0xba, 0x15, 0xa3, 0xdd, 0xa3, 0xfc, 0x68, 0xfa, 0x7e, 0xeb, 0xad, 0x9f, 0x62, 0xd9, 0x6d, 0x99,
	0x80, 0x46, 0xf9, 0x4d, 0x63, 0xe9, 0x67, 0xa7, 0xb4, 0xe9, 0x04, 0xd0, 0x1d, 0x59, 0x91, 0x80,
	0x4d, 0x95, 0x31, 0x46, 0x12, 0xca, 0x3a, 0xa1, 0x55, 0xdc, 0x5d, 0xbb, 0xdb, 0xf5, 0xb8, 0x28,
	0xf0, 0x9a, 0x06, 0x2e, 0xd8, 0x50, 0xd3, 0x86, 0xea, 0x2f, 0x8b, 0xb0, 0x36, 0xa5, 0x11, 0x87,
	0x65, 0xc1, 0x33, 0x45, 0xa4, 0xeb, 0x98, 0x53, 0x9c, 0xbf, 0x8b, 0x02, 0x79, 0x17, 0xcd, 0x56,
	0xa0, 0x51, 0x8b, 0x93, 0x97, 0x9d, 0x20, 0xa7, 0xa9, 0xfe, 0x51, 0x82, 0x72, 0xe1, 0x46, 0x7b,
	0xe3, 0xf1, 0x31, 0x33, 0x64, 0x46, 0x53, 0xe4, 0x39, 0xdc, 0x8f, 0x89, 0x54, 0x94, 0x99, 0x56,
	0x0e, 0x69, 0x1a, 0x26, 0x54, 0x2a, 0xb7, 0x64, 0x76, 0xbb, 0x33, 0xff, 0x4a, 0x37, 0x69, 0x2c,
	0x02, 0xcc, 0x3a, 0x24, 0xd8, 0x9a, 0xc8, 0x3d, 0x4b, 0x9f, 0x52, 0xa9, 0xd0, 0xe7, 0x30, 0x69,
	0x0c, 0x53, 0x2e, 0x94, 0x34, 0x77, 0x72, 0x35, 0xd8, 0x9c, 0x70, 0xb4, 0xb4, 0x5d, 0xcf, 0x12,
	0xc9, 0x33, 0x11, 0x91, 0x11, 0xf1, 0xe2, 0xff, 0x23, 0x5e, 0xb3, 0x69, 0x39, 0xe7, 0x1e, 0xe4,
	0xeb, 0x9c, 0x6e, 0xc9, 0xd0, 0x55, 0xac, 0xcd, 0x30, 0xd5, 0x9f, 0xbd, 0xfe, 0xed, 0xa7, 0xed,
	0x33, 0x38, 0x99, 0xc2, 0xb5, 0xd2, 0xcf, 0x53, 0xfe, 0xf0, 0x06, 0xe5, 0xeb, 0x27, 0x1a, 0xef,
	0x3b, 0xf8, 0xf6, 0x6e, 0x78, 0xae, 0x53, 0xfd, 0x7d, 0x01, 0x36, 0x66, 0xc6, 0x2e, 0xfa, 0x01,
	0xca, 0xa3, 0x89, 0x6e, 0xfb, 0xe6, 0xe5, 0x3b, 0x9a, 0xe8, 0x5e, 0xfe, 0xb5, 0xe6, 0x89, 0xfe,
	0x19, 0x11, 0x56, 0xff, 0x74, 0x60, 0x7d, 0x2a, 0x0a, 0x7d, 0x00, 0x8b, 0x0c, 0xf7, 0xc8, 0x6c,
	0x0f, 0x19, 0x23, 0xda, 0x81, 0x65, 0xfb, 0x44, 0x98, 0x17, 0x6a, 0x7d, 0x7c, 0xb3, 0x73, 0xf3,
	0x9c, 0xf7, 0x62, 0xe1, 0x2d, 0xde, 0x8b, 0xfa, 0xa5, 0x16, 0xfc, 0x05, 0x3c, 0xbf, 0xbd, 0xe0,
	0xff, 0x29, 0x41, 0xfd, 0x54, 0xe3, 0x36, 0xe1, 0xe8, 0xce, 0xb8, 0xd5, 0x1f, 0x1d, 0xd8, 0x98,
	0x99, 0x10, 0xe8, 0x23, 0x28, 0x77, 0xb9, 0x54, 0xf3, 0xe4, 0x1b, 0x39, 0xee, 0xb0, 0x85, 0x19,
	0xba, 0xfa, 0x97, 0x1a, 0xa9, 0x06, 0x07, 0xb7, 0x46, 0x6a, 0xb8, 0xb0, 0x95, 0x77, 0x40, 0x28,
	0x53, 0x12, 0xd1, 0x36, 0x25, 0x02, 0x2d, 0xfc, 0xd5, 0x70, 0x1a, 0xcf, 0xe0, 0x2b, 0xca, 0x6d,
	0xad, 0x6c, 0xda, 0xad, 0x3a, 0xb1, 0xb1, 0x5e, 0x50, 0xb4, 0xf4, 0x3b, 0xd0, 0x72, 0xae, 0x96,
	0xcd, 0x83, 0x50, 0xfb, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xac, 0x77, 0x64, 0x3f, 0x8e, 0x0a, 0x00,
	0x00,
}
