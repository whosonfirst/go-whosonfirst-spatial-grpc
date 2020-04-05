// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spatial.proto

package spatial

import (
	context "context"
	fmt "fmt"
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

type ExistentialFlag int32

const (
	ExistentialFlag_FALSE   ExistentialFlag = 0
	ExistentialFlag_TRUE    ExistentialFlag = 1
	ExistentialFlag_UNKNOWN ExistentialFlag = -1
)

var ExistentialFlag_name = map[int32]string{
	0:  "FALSE",
	1:  "TRUE",
	-1: "UNKNOWN",
}

var ExistentialFlag_value = map[string]int32{
	"FALSE":   0,
	"TRUE":    1,
	"UNKNOWN": -1,
}

func (x ExistentialFlag) String() string {
	return proto.EnumName(ExistentialFlag_name, int32(x))
}

func (ExistentialFlag) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ab81574adde31065, []int{0}
}

type Coordinate struct {
	Latitude             float32  `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float32  `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Coordinate) Reset()         { *m = Coordinate{} }
func (m *Coordinate) String() string { return proto.CompactTextString(m) }
func (*Coordinate) ProtoMessage()    {}
func (*Coordinate) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab81574adde31065, []int{0}
}

func (m *Coordinate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coordinate.Unmarshal(m, b)
}
func (m *Coordinate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coordinate.Marshal(b, m, deterministic)
}
func (m *Coordinate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coordinate.Merge(m, src)
}
func (m *Coordinate) XXX_Size() int {
	return xxx_messageInfo_Coordinate.Size(m)
}
func (m *Coordinate) XXX_DiscardUnknown() {
	xxx_messageInfo_Coordinate.DiscardUnknown(m)
}

var xxx_messageInfo_Coordinate proto.InternalMessageInfo

func (m *Coordinate) GetLatitude() float32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Coordinate) GetLongitude() float32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

type StandardPlaceResponse struct {
	// int64 wofid = 1;
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId             string          `protobuf:"bytes,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Placetype            string          `protobuf:"bytes,3,opt,name=placetype,proto3" json:"placetype,omitempty"`
	Country              string          `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Repo                 string          `protobuf:"bytes,5,opt,name=repo,proto3" json:"repo,omitempty"`
	Path                 string          `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
	Uri                  string          `protobuf:"bytes,7,opt,name=uri,proto3" json:"uri,omitempty"`
	Latitude             float32         `protobuf:"fixed32,8,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float32         `protobuf:"fixed32,9,opt,name=longitude,proto3" json:"longitude,omitempty"`
	IsCurrent            ExistentialFlag `protobuf:"varint,10,opt,name=is_current,json=isCurrent,proto3,enum=ExistentialFlag" json:"is_current,omitempty"`
	IsCeased             ExistentialFlag `protobuf:"varint,11,opt,name=is_ceased,json=isCeased,proto3,enum=ExistentialFlag" json:"is_ceased,omitempty"`
	IsDeprecated         ExistentialFlag `protobuf:"varint,12,opt,name=is_deprecated,json=isDeprecated,proto3,enum=ExistentialFlag" json:"is_deprecated,omitempty"`
	IsSuperseded         ExistentialFlag `protobuf:"varint,13,opt,name=is_superseded,json=isSuperseded,proto3,enum=ExistentialFlag" json:"is_superseded,omitempty"`
	IsSuperseding        ExistentialFlag `protobuf:"varint,14,opt,name=is_superseding,json=isSuperseding,proto3,enum=ExistentialFlag" json:"is_superseding,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StandardPlaceResponse) Reset()         { *m = StandardPlaceResponse{} }
func (m *StandardPlaceResponse) String() string { return proto.CompactTextString(m) }
func (*StandardPlaceResponse) ProtoMessage()    {}
func (*StandardPlaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab81574adde31065, []int{1}
}

func (m *StandardPlaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StandardPlaceResponse.Unmarshal(m, b)
}
func (m *StandardPlaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StandardPlaceResponse.Marshal(b, m, deterministic)
}
func (m *StandardPlaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StandardPlaceResponse.Merge(m, src)
}
func (m *StandardPlaceResponse) XXX_Size() int {
	return xxx_messageInfo_StandardPlaceResponse.Size(m)
}
func (m *StandardPlaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StandardPlaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StandardPlaceResponse proto.InternalMessageInfo

func (m *StandardPlaceResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StandardPlaceResponse) GetParentId() string {
	if m != nil {
		return m.ParentId
	}
	return ""
}

func (m *StandardPlaceResponse) GetPlacetype() string {
	if m != nil {
		return m.Placetype
	}
	return ""
}

func (m *StandardPlaceResponse) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *StandardPlaceResponse) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *StandardPlaceResponse) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *StandardPlaceResponse) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *StandardPlaceResponse) GetLatitude() float32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *StandardPlaceResponse) GetLongitude() float32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *StandardPlaceResponse) GetIsCurrent() ExistentialFlag {
	if m != nil {
		return m.IsCurrent
	}
	return ExistentialFlag_FALSE
}

func (m *StandardPlaceResponse) GetIsCeased() ExistentialFlag {
	if m != nil {
		return m.IsCeased
	}
	return ExistentialFlag_FALSE
}

func (m *StandardPlaceResponse) GetIsDeprecated() ExistentialFlag {
	if m != nil {
		return m.IsDeprecated
	}
	return ExistentialFlag_FALSE
}

func (m *StandardPlaceResponse) GetIsSuperseded() ExistentialFlag {
	if m != nil {
		return m.IsSuperseded
	}
	return ExistentialFlag_FALSE
}

func (m *StandardPlaceResponse) GetIsSuperseding() ExistentialFlag {
	if m != nil {
		return m.IsSuperseding
	}
	return ExistentialFlag_FALSE
}

type StandardPlacesResults struct {
	Results              []*StandardPlaceResponse `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *StandardPlacesResults) Reset()         { *m = StandardPlacesResults{} }
func (m *StandardPlacesResults) String() string { return proto.CompactTextString(m) }
func (*StandardPlacesResults) ProtoMessage()    {}
func (*StandardPlacesResults) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab81574adde31065, []int{2}
}

func (m *StandardPlacesResults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StandardPlacesResults.Unmarshal(m, b)
}
func (m *StandardPlacesResults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StandardPlacesResults.Marshal(b, m, deterministic)
}
func (m *StandardPlacesResults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StandardPlacesResults.Merge(m, src)
}
func (m *StandardPlacesResults) XXX_Size() int {
	return xxx_messageInfo_StandardPlacesResults.Size(m)
}
func (m *StandardPlacesResults) XXX_DiscardUnknown() {
	xxx_messageInfo_StandardPlacesResults.DiscardUnknown(m)
}

var xxx_messageInfo_StandardPlacesResults proto.InternalMessageInfo

func (m *StandardPlacesResults) GetResults() []*StandardPlaceResponse {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterEnum("ExistentialFlag", ExistentialFlag_name, ExistentialFlag_value)
	proto.RegisterType((*Coordinate)(nil), "Coordinate")
	proto.RegisterType((*StandardPlaceResponse)(nil), "StandardPlaceResponse")
	proto.RegisterType((*StandardPlacesResults)(nil), "StandardPlacesResults")
}

func init() {
	proto.RegisterFile("spatial.proto", fileDescriptor_ab81574adde31065)
}

var fileDescriptor_ab81574adde31065 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xc1, 0xaf, 0xd2, 0x40,
	0x10, 0xc6, 0x5f, 0x81, 0xf7, 0xda, 0x0e, 0x0f, 0x6c, 0x26, 0x6a, 0x36, 0x4f, 0x0f, 0x84, 0x13,
	0x31, 0x11, 0x0d, 0xc6, 0x78, 0xf1, 0xa2, 0x08, 0x09, 0xd1, 0x20, 0x29, 0x12, 0x8f, 0x64, 0x65,
	0x27, 0x75, 0x93, 0x66, 0x77, 0xb3, 0xbb, 0x4d, 0xe4, 0x2f, 0xf7, 0xa8, 0xe9, 0x56, 0x40, 0x89,
	0xbc, 0x9e, 0x66, 0xbf, 0xf9, 0x7e, 0x3b, 0x33, 0xed, 0x14, 0x7a, 0xce, 0x70, 0x2f, 0x79, 0x39,
	0x36, 0x56, 0x7b, 0x3d, 0x9c, 0x03, 0x4c, 0xb5, 0xb6, 0x42, 0x2a, 0xee, 0x09, 0xef, 0x20, 0x29,
	0xb9, 0x97, 0xbe, 0x12, 0xc4, 0xa2, 0x41, 0x34, 0x6a, 0xe5, 0xc7, 0x33, 0x3e, 0x85, 0xb4, 0xd4,
	0xaa, 0x68, 0x92, 0xad, 0x90, 0x3c, 0x09, 0xc3, 0x9f, 0x6d, 0x78, 0xb4, 0xf6, 0x5c, 0x09, 0x6e,
	0xc5, 0xaa, 0xe4, 0x3b, 0xca, 0xc9, 0x19, 0xad, 0x1c, 0x61, 0x1f, 0x5a, 0x52, 0x84, 0xdb, 0xd2,
	0xbc, 0x25, 0x05, 0x3e, 0x81, 0xd4, 0x70, 0x4b, 0xca, 0x6f, 0xa5, 0x08, 0xf7, 0xa4, 0x79, 0xd2,
	0x08, 0x0b, 0x51, 0x17, 0x31, 0x35, 0xed, 0xf7, 0x86, 0x58, 0x3b, 0x24, 0x4f, 0x02, 0x32, 0x88,
	0x77, 0xba, 0x52, 0xde, 0xee, 0x59, 0x27, 0xe4, 0x0e, 0x47, 0x44, 0xe8, 0x58, 0x32, 0x9a, 0x5d,
	0x07, 0x39, 0xc4, 0xb5, 0x66, 0xb8, 0xff, 0xce, 0x6e, 0x1a, 0xad, 0x8e, 0x31, 0x83, 0x76, 0x65,
	0x25, 0x8b, 0x83, 0x54, 0x87, 0xff, 0x8c, 0x9c, 0xdc, 0x37, 0x72, 0x7a, 0x36, 0x32, 0xbe, 0x00,
	0x90, 0x6e, 0xbb, 0xab, 0x6c, 0xdd, 0x3b, 0x83, 0x41, 0x34, 0xea, 0x4f, 0xb2, 0xf1, 0xec, 0x87,
	0x74, 0x9e, 0x54, 0xfd, 0x8a, 0xe7, 0x25, 0x2f, 0xf2, 0x54, 0xba, 0x69, 0x63, 0xc1, 0xe7, 0x90,
	0xd6, 0x00, 0x71, 0x47, 0x82, 0x75, 0x2f, 0xf8, 0x13, 0xe9, 0xa6, 0xc1, 0x81, 0xaf, 0xa1, 0x27,
	0xdd, 0x56, 0x90, 0xb1, 0xb4, 0xe3, 0x9e, 0x04, 0xbb, 0xbd, 0x80, 0xdc, 0x4a, 0xf7, 0xe1, 0xe8,
	0xfa, 0x83, 0xb9, 0xca, 0x90, 0x75, 0x24, 0x48, 0xb0, 0xde, 0x65, 0x6c, 0x7d, 0x74, 0xe1, 0x1b,
	0xe8, 0xff, 0x85, 0x49, 0x55, 0xb0, 0xfe, 0x05, 0xae, 0x77, 0xe2, 0xa4, 0x2a, 0x86, 0x8b, 0xb3,
	0x0f, 0xef, 0x72, 0x72, 0x55, 0xe9, 0x1d, 0xbe, 0x84, 0xd8, 0x36, 0x21, 0x8b, 0x06, 0xed, 0x51,
	0x77, 0xf2, 0x78, 0xfc, 0xdf, 0x0d, 0xc9, 0x0f, 0xb6, 0x67, 0x6f, 0xe1, 0xc1, 0x59, 0x31, 0x4c,
	0xe1, 0x7a, 0xfe, 0xee, 0xd3, 0x7a, 0x96, 0x5d, 0x61, 0x02, 0x9d, 0x2f, 0xf9, 0x66, 0x96, 0x45,
	0xf8, 0x10, 0xe2, 0xcd, 0xf2, 0xe3, 0xf2, 0xf3, 0xd7, 0x65, 0xf6, 0xeb, 0xf0, 0x44, 0x93, 0xf7,
	0x10, 0xaf, 0x9b, 0xdd, 0xae, 0x87, 0x59, 0x69, 0xa9, 0xfc, 0x42, 0xad, 0x74, 0xb9, 0x2f, 0xb4,
	0xc2, 0xee, 0xf8, 0xb4, 0xe6, 0x77, 0x67, 0x8d, 0x1c, 0x3a, 0x1e, 0x5e, 0x7d, 0xbb, 0x09, 0x7f,
	0xc5, 0xab, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xe8, 0x15, 0x9b, 0x26, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SpatialClient is the client API for Spatial service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SpatialClient interface {
	PointInPolygon(ctx context.Context, in *Coordinate, opts ...grpc.CallOption) (*StandardPlacesResults, error)
}

type spatialClient struct {
	cc grpc.ClientConnInterface
}

func NewSpatialClient(cc grpc.ClientConnInterface) SpatialClient {
	return &spatialClient{cc}
}

func (c *spatialClient) PointInPolygon(ctx context.Context, in *Coordinate, opts ...grpc.CallOption) (*StandardPlacesResults, error) {
	out := new(StandardPlacesResults)
	err := c.cc.Invoke(ctx, "/Spatial/PointInPolygon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpatialServer is the server API for Spatial service.
type SpatialServer interface {
	PointInPolygon(context.Context, *Coordinate) (*StandardPlacesResults, error)
}

// UnimplementedSpatialServer can be embedded to have forward compatible implementations.
type UnimplementedSpatialServer struct {
}

func (*UnimplementedSpatialServer) PointInPolygon(ctx context.Context, req *Coordinate) (*StandardPlacesResults, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PointInPolygon not implemented")
}

func RegisterSpatialServer(s *grpc.Server, srv SpatialServer) {
	s.RegisterService(&_Spatial_serviceDesc, srv)
}

func _Spatial_PointInPolygon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Coordinate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpatialServer).PointInPolygon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Spatial/PointInPolygon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpatialServer).PointInPolygon(ctx, req.(*Coordinate))
	}
	return interceptor(ctx, in, info, handler)
}

var _Spatial_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Spatial",
	HandlerType: (*SpatialServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PointInPolygon",
			Handler:    _Spatial_PointInPolygon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spatial.proto",
}
