// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0-devel
// 	protoc        v3.6.0
// source: server.proto

package path_api_v1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var File_server_proto protoreflect.FileDescriptor

var file_server_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x70, 0x61, 0x74, 0x68, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x59, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x4f, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x12, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_server_proto_goTypes = []interface{}{
	(*empty.Empty)(nil), // 0: google.protobuf.Empty
}
var file_server_proto_depIdxs = []int32{
	0, // 0: path_api.v1.Server.GetHealth:input_type -> google.protobuf.Empty
	0, // 1: path_api.v1.Server.GetHealth:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_proto_init() }
func file_server_proto_init() {
	if File_server_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_proto_goTypes,
		DependencyIndexes: file_server_proto_depIdxs,
	}.Build()
	File_server_proto = out.File
	file_server_proto_rawDesc = nil
	file_server_proto_goTypes = nil
	file_server_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServerClient is the client API for Server service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerClient interface {
	// Liveness/readiness probe, used by ESP to ensure the backend is up and ready.
	GetHealth(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type serverClient struct {
	cc grpc.ClientConnInterface
}

func NewServerClient(cc grpc.ClientConnInterface) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) GetHealth(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/path_api.v1.Server/GetHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerServer is the server API for Server service.
type ServerServer interface {
	// Liveness/readiness probe, used by ESP to ensure the backend is up and ready.
	GetHealth(context.Context, *empty.Empty) (*empty.Empty, error)
}

// UnimplementedServerServer can be embedded to have forward compatible implementations.
type UnimplementedServerServer struct {
}

func (*UnimplementedServerServer) GetHealth(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealth not implemented")
}

func RegisterServerServer(s *grpc.Server, srv ServerServer) {
	s.RegisterService(&_Server_serviceDesc, srv)
}

func _Server_GetHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).GetHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/path_api.v1.Server/GetHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).GetHealth(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Server_serviceDesc = grpc.ServiceDesc{
	ServiceName: "path_api.v1.Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHealth",
			Handler:    _Server_GetHealth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
