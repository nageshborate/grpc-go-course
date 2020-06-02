// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.1
// source: pnd/pndpb/pnd.proto

package pndpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PrimeNumberDecompositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InputNumber int64 `protobuf:"varint,1,opt,name=inputNumber,proto3" json:"inputNumber,omitempty"`
}

func (x *PrimeNumberDecompositionRequest) Reset() {
	*x = PrimeNumberDecompositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pnd_pndpb_pnd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumberDecompositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumberDecompositionRequest) ProtoMessage() {}

func (x *PrimeNumberDecompositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pnd_pndpb_pnd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeNumberDecompositionRequest.ProtoReflect.Descriptor instead.
func (*PrimeNumberDecompositionRequest) Descriptor() ([]byte, []int) {
	return file_pnd_pndpb_pnd_proto_rawDescGZIP(), []int{0}
}

func (x *PrimeNumberDecompositionRequest) GetInputNumber() int64 {
	if x != nil {
		return x.InputNumber
	}
	return 0
}

type PrimeNumberDecompositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PrimeNumberDecompositionResponse) Reset() {
	*x = PrimeNumberDecompositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pnd_pndpb_pnd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumberDecompositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumberDecompositionResponse) ProtoMessage() {}

func (x *PrimeNumberDecompositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pnd_pndpb_pnd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeNumberDecompositionResponse.ProtoReflect.Descriptor instead.
func (*PrimeNumberDecompositionResponse) Descriptor() ([]byte, []int) {
	return file_pnd_pndpb_pnd_proto_rawDescGZIP(), []int{1}
}

func (x *PrimeNumberDecompositionResponse) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_pnd_pndpb_pnd_proto protoreflect.FileDescriptor

var file_pnd_pndpb_pnd_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6e, 0x64, 0x2f, 0x70, 0x6e, 0x64, 0x70, 0x62, 0x2f, 0x70, 0x6e, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x70, 0x6e, 0x64, 0x22, 0x43, 0x0a, 0x1f, 0x50, 0x72,
	0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22,
	0x3a, 0x0a, 0x20, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x7f, 0x0a, 0x1f, 0x50,
	0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c,
	0x0a, 0x09, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x24, 0x2e, 0x70, 0x6e,
	0x64, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x70, 0x6e, 0x64, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09,
	0x70, 0x6e, 0x64, 0x2f, 0x70, 0x6e, 0x64, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pnd_pndpb_pnd_proto_rawDescOnce sync.Once
	file_pnd_pndpb_pnd_proto_rawDescData = file_pnd_pndpb_pnd_proto_rawDesc
)

func file_pnd_pndpb_pnd_proto_rawDescGZIP() []byte {
	file_pnd_pndpb_pnd_proto_rawDescOnce.Do(func() {
		file_pnd_pndpb_pnd_proto_rawDescData = protoimpl.X.CompressGZIP(file_pnd_pndpb_pnd_proto_rawDescData)
	})
	return file_pnd_pndpb_pnd_proto_rawDescData
}

var file_pnd_pndpb_pnd_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pnd_pndpb_pnd_proto_goTypes = []interface{}{
	(*PrimeNumberDecompositionRequest)(nil),  // 0: pnd.PrimeNumberDecompositionRequest
	(*PrimeNumberDecompositionResponse)(nil), // 1: pnd.PrimeNumberDecompositionResponse
}
var file_pnd_pndpb_pnd_proto_depIdxs = []int32{
	0, // 0: pnd.PrimeNumberDecompositionService.Decompose:input_type -> pnd.PrimeNumberDecompositionRequest
	1, // 1: pnd.PrimeNumberDecompositionService.Decompose:output_type -> pnd.PrimeNumberDecompositionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pnd_pndpb_pnd_proto_init() }
func file_pnd_pndpb_pnd_proto_init() {
	if File_pnd_pndpb_pnd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pnd_pndpb_pnd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumberDecompositionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pnd_pndpb_pnd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumberDecompositionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pnd_pndpb_pnd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pnd_pndpb_pnd_proto_goTypes,
		DependencyIndexes: file_pnd_pndpb_pnd_proto_depIdxs,
		MessageInfos:      file_pnd_pndpb_pnd_proto_msgTypes,
	}.Build()
	File_pnd_pndpb_pnd_proto = out.File
	file_pnd_pndpb_pnd_proto_rawDesc = nil
	file_pnd_pndpb_pnd_proto_goTypes = nil
	file_pnd_pndpb_pnd_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PrimeNumberDecompositionServiceClient is the client API for PrimeNumberDecompositionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PrimeNumberDecompositionServiceClient interface {
	Decompose(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (PrimeNumberDecompositionService_DecomposeClient, error)
}

type primeNumberDecompositionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrimeNumberDecompositionServiceClient(cc grpc.ClientConnInterface) PrimeNumberDecompositionServiceClient {
	return &primeNumberDecompositionServiceClient{cc}
}

func (c *primeNumberDecompositionServiceClient) Decompose(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (PrimeNumberDecompositionService_DecomposeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PrimeNumberDecompositionService_serviceDesc.Streams[0], "/pnd.PrimeNumberDecompositionService/Decompose", opts...)
	if err != nil {
		return nil, err
	}
	x := &primeNumberDecompositionServiceDecomposeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PrimeNumberDecompositionService_DecomposeClient interface {
	Recv() (*PrimeNumberDecompositionResponse, error)
	grpc.ClientStream
}

type primeNumberDecompositionServiceDecomposeClient struct {
	grpc.ClientStream
}

func (x *primeNumberDecompositionServiceDecomposeClient) Recv() (*PrimeNumberDecompositionResponse, error) {
	m := new(PrimeNumberDecompositionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PrimeNumberDecompositionServiceServer is the server API for PrimeNumberDecompositionService service.
type PrimeNumberDecompositionServiceServer interface {
	Decompose(*PrimeNumberDecompositionRequest, PrimeNumberDecompositionService_DecomposeServer) error
}

// UnimplementedPrimeNumberDecompositionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPrimeNumberDecompositionServiceServer struct {
}

func (*UnimplementedPrimeNumberDecompositionServiceServer) Decompose(*PrimeNumberDecompositionRequest, PrimeNumberDecompositionService_DecomposeServer) error {
	return status.Errorf(codes.Unimplemented, "method Decompose not implemented")
}

func RegisterPrimeNumberDecompositionServiceServer(s *grpc.Server, srv PrimeNumberDecompositionServiceServer) {
	s.RegisterService(&_PrimeNumberDecompositionService_serviceDesc, srv)
}

func _PrimeNumberDecompositionService_Decompose_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberDecompositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PrimeNumberDecompositionServiceServer).Decompose(m, &primeNumberDecompositionServiceDecomposeServer{stream})
}

type PrimeNumberDecompositionService_DecomposeServer interface {
	Send(*PrimeNumberDecompositionResponse) error
	grpc.ServerStream
}

type primeNumberDecompositionServiceDecomposeServer struct {
	grpc.ServerStream
}

func (x *primeNumberDecompositionServiceDecomposeServer) Send(m *PrimeNumberDecompositionResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _PrimeNumberDecompositionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pnd.PrimeNumberDecompositionService",
	HandlerType: (*PrimeNumberDecompositionServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Decompose",
			Handler:       _PrimeNumberDecompositionService_Decompose_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pnd/pndpb/pnd.proto",
}
