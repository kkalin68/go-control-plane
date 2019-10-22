// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/tap/v3alpha/tapds.proto

package envoy_service_tap_v3alpha

import (
	context "context"
	fmt "fmt"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type TapResource struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Config               *TapConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TapResource) Reset()         { *m = TapResource{} }
func (m *TapResource) String() string { return proto.CompactTextString(m) }
func (*TapResource) ProtoMessage()    {}
func (*TapResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba445dccde093552, []int{0}
}

func (m *TapResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TapResource.Unmarshal(m, b)
}
func (m *TapResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TapResource.Marshal(b, m, deterministic)
}
func (m *TapResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TapResource.Merge(m, src)
}
func (m *TapResource) XXX_Size() int {
	return xxx_messageInfo_TapResource.Size(m)
}
func (m *TapResource) XXX_DiscardUnknown() {
	xxx_messageInfo_TapResource.DiscardUnknown(m)
}

var xxx_messageInfo_TapResource proto.InternalMessageInfo

func (m *TapResource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TapResource) GetConfig() *TapConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*TapResource)(nil), "envoy.service.tap.v3alpha.TapResource")
}

func init() {
	proto.RegisterFile("envoy/service/tap/v3alpha/tapds.proto", fileDescriptor_ba445dccde093552)
}

var fileDescriptor_ba445dccde093552 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x3f, 0x4f, 0xeb, 0x30,
	0x14, 0xc5, 0x9f, 0xa3, 0xaa, 0x4f, 0xcf, 0x1d, 0xfa, 0x94, 0x37, 0xbc, 0x12, 0x10, 0x2a, 0x6d,
	0x81, 0x94, 0x21, 0x41, 0xed, 0x56, 0x31, 0x85, 0x8a, 0xb9, 0x4a, 0xb3, 0xa3, 0x4b, 0x6a, 0x5a,
	0x4b, 0x89, 0xaf, 0x89, 0xdd, 0x88, 0xae, 0x2c, 0xb0, 0xf3, 0xd1, 0xf8, 0x0a, 0x7c, 0x02, 0x46,
	0x26, 0x94, 0x3f, 0x8d, 0x40, 0xa5, 0x88, 0x85, 0xed, 0x2a, 0xf9, 0x9d, 0x73, 0x7c, 0xed, 0x43,
	0x0f, 0x99, 0x48, 0x71, 0xe5, 0x2a, 0x96, 0xa4, 0x3c, 0x64, 0xae, 0x06, 0xe9, 0xa6, 0x43, 0x88,
	0xe4, 0x02, 0xb2, 0x79, 0xa6, 0x1c, 0x99, 0xa0, 0x46, 0x73, 0x27, 0xc7, 0x9c, 0x12, 0x73, 0x34,
	0x48, 0xa7, 0xc4, 0xac, 0x83, 0xc2, 0x01, 0x24, 0xaf, 0x94, 0x33, 0xae, 0x42, 0x4c, 0x59, 0xb2,
	0x2a, 0xd4, 0xd6, 0xd1, 0xf6, 0x90, 0x10, 0xe3, 0x18, 0x45, 0xc9, 0xed, 0xcd, 0x11, 0xe7, 0x11,
	0xcb, 0xbd, 0x40, 0x08, 0xd4, 0xa0, 0x39, 0x8a, 0xf2, 0x0c, 0xd6, 0xff, 0x14, 0x22, 0x3e, 0x03,
	0xcd, 0xdc, 0xf5, 0x50, 0xfc, 0xe8, 0x2c, 0x68, 0x23, 0x00, 0xe9, 0x33, 0x85, 0xcb, 0x24, 0x64,
	0xe6, 0x2e, 0xad, 0x09, 0x88, 0x59, 0x8b, 0xb4, 0x89, 0xfd, 0xc7, 0xfb, 0xfd, 0xea, 0xd5, 0x12,
	0xa3, 0x4d, 0xfc, 0xfc, 0xa3, 0x79, 0x46, 0xeb, 0x21, 0x8a, 0x6b, 0x3e, 0x6f, 0x19, 0x6d, 0x62,
	0x37, 0x06, 0x3d, 0x67, 0xeb, 0x66, 0x4e, 0x00, 0xf2, 0x3c, 0x67, 0xfd, 0x52, 0x33, 0x78, 0x31,
	0xe8, 0xbf, 0x00, 0xe4, 0x78, 0xbd, 0xdf, 0xb4, 0x50, 0x99, 0x21, 0xfd, 0x3b, 0xd5, 0x09, 0x83,
	0xb8, 0x92, 0x28, 0xb3, 0x5b, 0x3a, 0x83, 0xe4, 0x95, 0x63, 0x25, 0xf4, 0xd9, 0xcd, 0x92, 0x29,
	0x6d, 0xf5, 0xbe, 0x86, 0x94, 0x44, 0xa1, 0x58, 0xe7, 0x97, 0x4d, 0x4e, 0x89, 0x19, 0xd1, 0xe6,
	0x98, 0x45, 0x1a, 0xde, 0x65, 0xd8, 0x9f, 0xc9, 0x33, 0x66, 0x23, 0xa8, 0xff, 0x0d, 0xf2, 0x43,
	0xda, 0x3d, 0xa1, 0xcd, 0x0b, 0xa6, 0xc3, 0xc5, 0xcf, 0xac, 0xd4, 0xbf, 0x7b, 0x7a, 0x7e, 0x34,
	0xba, 0x9d, 0xfd, 0xcd, 0xd2, 0x8c, 0x34, 0xc8, 0xcb, 0xe2, 0xc2, 0xd5, 0x88, 0x9c, 0x78, 0x23,
	0x7a, 0xcc, 0xb1, 0x30, 0x95, 0x09, 0xde, 0xae, 0xb6, 0xbf, 0x98, 0x47, 0x83, 0xac, 0xb3, 0x93,
	0xac, 0x15, 0x13, 0xf2, 0x40, 0xc8, 0x55, 0x3d, 0x6f, 0xc8, 0xf0, 0x2d, 0x00, 0x00, 0xff, 0xff,
	0x64, 0xb8, 0xf5, 0x18, 0xe7, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TapDiscoveryServiceClient is the client API for TapDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TapDiscoveryServiceClient interface {
	StreamTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_StreamTapConfigsClient, error)
	DeltaTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_DeltaTapConfigsClient, error)
	FetchTapConfigs(ctx context.Context, in *v3alpha.DiscoveryRequest, opts ...grpc.CallOption) (*v3alpha.DiscoveryResponse, error)
}

type tapDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewTapDiscoveryServiceClient(cc *grpc.ClientConn) TapDiscoveryServiceClient {
	return &tapDiscoveryServiceClient{cc}
}

func (c *tapDiscoveryServiceClient) StreamTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_StreamTapConfigsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TapDiscoveryService_serviceDesc.Streams[0], "/envoy.service.tap.v3alpha.TapDiscoveryService/StreamTapConfigs", opts...)
	if err != nil {
		return nil, err
	}
	x := &tapDiscoveryServiceStreamTapConfigsClient{stream}
	return x, nil
}

type TapDiscoveryService_StreamTapConfigsClient interface {
	Send(*v3alpha.DiscoveryRequest) error
	Recv() (*v3alpha.DiscoveryResponse, error)
	grpc.ClientStream
}

type tapDiscoveryServiceStreamTapConfigsClient struct {
	grpc.ClientStream
}

func (x *tapDiscoveryServiceStreamTapConfigsClient) Send(m *v3alpha.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tapDiscoveryServiceStreamTapConfigsClient) Recv() (*v3alpha.DiscoveryResponse, error) {
	m := new(v3alpha.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tapDiscoveryServiceClient) DeltaTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_DeltaTapConfigsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TapDiscoveryService_serviceDesc.Streams[1], "/envoy.service.tap.v3alpha.TapDiscoveryService/DeltaTapConfigs", opts...)
	if err != nil {
		return nil, err
	}
	x := &tapDiscoveryServiceDeltaTapConfigsClient{stream}
	return x, nil
}

type TapDiscoveryService_DeltaTapConfigsClient interface {
	Send(*v3alpha.DeltaDiscoveryRequest) error
	Recv() (*v3alpha.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type tapDiscoveryServiceDeltaTapConfigsClient struct {
	grpc.ClientStream
}

func (x *tapDiscoveryServiceDeltaTapConfigsClient) Send(m *v3alpha.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tapDiscoveryServiceDeltaTapConfigsClient) Recv() (*v3alpha.DeltaDiscoveryResponse, error) {
	m := new(v3alpha.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tapDiscoveryServiceClient) FetchTapConfigs(ctx context.Context, in *v3alpha.DiscoveryRequest, opts ...grpc.CallOption) (*v3alpha.DiscoveryResponse, error) {
	out := new(v3alpha.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.tap.v3alpha.TapDiscoveryService/FetchTapConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TapDiscoveryServiceServer is the server API for TapDiscoveryService service.
type TapDiscoveryServiceServer interface {
	StreamTapConfigs(TapDiscoveryService_StreamTapConfigsServer) error
	DeltaTapConfigs(TapDiscoveryService_DeltaTapConfigsServer) error
	FetchTapConfigs(context.Context, *v3alpha.DiscoveryRequest) (*v3alpha.DiscoveryResponse, error)
}

func RegisterTapDiscoveryServiceServer(s *grpc.Server, srv TapDiscoveryServiceServer) {
	s.RegisterService(&_TapDiscoveryService_serviceDesc, srv)
}

func _TapDiscoveryService_StreamTapConfigs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TapDiscoveryServiceServer).StreamTapConfigs(&tapDiscoveryServiceStreamTapConfigsServer{stream})
}

type TapDiscoveryService_StreamTapConfigsServer interface {
	Send(*v3alpha.DiscoveryResponse) error
	Recv() (*v3alpha.DiscoveryRequest, error)
	grpc.ServerStream
}

type tapDiscoveryServiceStreamTapConfigsServer struct {
	grpc.ServerStream
}

func (x *tapDiscoveryServiceStreamTapConfigsServer) Send(m *v3alpha.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tapDiscoveryServiceStreamTapConfigsServer) Recv() (*v3alpha.DiscoveryRequest, error) {
	m := new(v3alpha.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TapDiscoveryService_DeltaTapConfigs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TapDiscoveryServiceServer).DeltaTapConfigs(&tapDiscoveryServiceDeltaTapConfigsServer{stream})
}

type TapDiscoveryService_DeltaTapConfigsServer interface {
	Send(*v3alpha.DeltaDiscoveryResponse) error
	Recv() (*v3alpha.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type tapDiscoveryServiceDeltaTapConfigsServer struct {
	grpc.ServerStream
}

func (x *tapDiscoveryServiceDeltaTapConfigsServer) Send(m *v3alpha.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tapDiscoveryServiceDeltaTapConfigsServer) Recv() (*v3alpha.DeltaDiscoveryRequest, error) {
	m := new(v3alpha.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TapDiscoveryService_FetchTapConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3alpha.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TapDiscoveryServiceServer).FetchTapConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.tap.v3alpha.TapDiscoveryService/FetchTapConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TapDiscoveryServiceServer).FetchTapConfigs(ctx, req.(*v3alpha.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TapDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.tap.v3alpha.TapDiscoveryService",
	HandlerType: (*TapDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchTapConfigs",
			Handler:    _TapDiscoveryService_FetchTapConfigs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTapConfigs",
			Handler:       _TapDiscoveryService_StreamTapConfigs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaTapConfigs",
			Handler:       _TapDiscoveryService_DeltaTapConfigs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/tap/v3alpha/tapds.proto",
}