// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: agent.proto

package agentcomm

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

// The request message containing the user's name.
type ConnectContainerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GcmIP    string `protobuf:"bytes,1,opt,name=gcmIP,proto3" json:"gcmIP,omitempty"`
	PodName  string `protobuf:"bytes,2,opt,name=podName,proto3" json:"podName,omitempty"`
	DockerId string `protobuf:"bytes,3,opt,name=dockerId,proto3" json:"dockerId,omitempty"`
}

func (x *ConnectContainerRequest) Reset() {
	*x = ConnectContainerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectContainerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectContainerRequest) ProtoMessage() {}

func (x *ConnectContainerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectContainerRequest.ProtoReflect.Descriptor instead.
func (*ConnectContainerRequest) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectContainerRequest) GetGcmIP() string {
	if x != nil {
		return x.GcmIP
	}
	return ""
}

func (x *ConnectContainerRequest) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

func (x *ConnectContainerRequest) GetDockerId() string {
	if x != nil {
		return x.DockerId
	}
	return ""
}

// The response message containing the greetings
type ConnectContainerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodName  string `protobuf:"bytes,1,opt,name=podName,proto3" json:"podName,omitempty"`
	DockerID string `protobuf:"bytes,2,opt,name=dockerID,proto3" json:"dockerID,omitempty"`
	CgroupID int32  `protobuf:"varint,3,opt,name=cgroupID,proto3" json:"cgroupID,omitempty"`
}

func (x *ConnectContainerReply) Reset() {
	*x = ConnectContainerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectContainerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectContainerReply) ProtoMessage() {}

func (x *ConnectContainerReply) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectContainerReply.ProtoReflect.Descriptor instead.
func (*ConnectContainerReply) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectContainerReply) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

func (x *ConnectContainerReply) GetDockerID() string {
	if x != nil {
		return x.DockerID
	}
	return ""
}

func (x *ConnectContainerReply) GetCgroupID() int32 {
	if x != nil {
		return x.CgroupID
	}
	return 0
}

var File_agent_proto protoreflect.FileDescriptor

var file_agent_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x63, 0x6f, 0x6d, 0x6d, 0x22, 0x65, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x63, 0x6d, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x67, 0x63, 0x6d, 0x49, 0x50, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6f, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x64, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x69, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6f, 0x64, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x63, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x32, 0x68, 0x0a, 0x07, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x5d, 0x0a, 0x13, 0x52, 0x65, 0x71, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x63,
	0x6f, 0x6d, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agent_proto_rawDescOnce sync.Once
	file_agent_proto_rawDescData = file_agent_proto_rawDesc
)

func file_agent_proto_rawDescGZIP() []byte {
	file_agent_proto_rawDescOnce.Do(func() {
		file_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_agent_proto_rawDescData)
	})
	return file_agent_proto_rawDescData
}

var file_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_agent_proto_goTypes = []interface{}{
	(*ConnectContainerRequest)(nil), // 0: agentcomm.ConnectContainerRequest
	(*ConnectContainerReply)(nil),   // 1: agentcomm.ConnectContainerReply
}
var file_agent_proto_depIdxs = []int32{
	0, // 0: agentcomm.Handler.ReqConnectContainer:input_type -> agentcomm.ConnectContainerRequest
	1, // 1: agentcomm.Handler.ReqConnectContainer:output_type -> agentcomm.ConnectContainerReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_agent_proto_init() }
func file_agent_proto_init() {
	if File_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectContainerRequest); i {
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
		file_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectContainerReply); i {
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
			RawDescriptor: file_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agent_proto_goTypes,
		DependencyIndexes: file_agent_proto_depIdxs,
		MessageInfos:      file_agent_proto_msgTypes,
	}.Build()
	File_agent_proto = out.File
	file_agent_proto_rawDesc = nil
	file_agent_proto_goTypes = nil
	file_agent_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HandlerClient is the client API for Handler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HandlerClient interface {
	// Sends a greeting
	ReqConnectContainer(ctx context.Context, in *ConnectContainerRequest, opts ...grpc.CallOption) (*ConnectContainerReply, error)
}

type handlerClient struct {
	cc grpc.ClientConnInterface
}

func NewHandlerClient(cc grpc.ClientConnInterface) HandlerClient {
	return &handlerClient{cc}
}

func (c *handlerClient) ReqConnectContainer(ctx context.Context, in *ConnectContainerRequest, opts ...grpc.CallOption) (*ConnectContainerReply, error) {
	out := new(ConnectContainerReply)
	err := c.cc.Invoke(ctx, "/agentcomm.Handler/ReqConnectContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HandlerServer is the server API for Handler service.
type HandlerServer interface {
	// Sends a greeting
	ReqConnectContainer(context.Context, *ConnectContainerRequest) (*ConnectContainerReply, error)
}

// UnimplementedHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedHandlerServer struct {
}

func (*UnimplementedHandlerServer) ReqConnectContainer(context.Context, *ConnectContainerRequest) (*ConnectContainerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReqConnectContainer not implemented")
}

func RegisterHandlerServer(s *grpc.Server, srv HandlerServer) {
	s.RegisterService(&_Handler_serviceDesc, srv)
}

func _Handler_ReqConnectContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerServer).ReqConnectContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentcomm.Handler/ReqConnectContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerServer).ReqConnectContainer(ctx, req.(*ConnectContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Handler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "agentcomm.Handler",
	HandlerType: (*HandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReqConnectContainer",
			Handler:    _Handler_ReqConnectContainer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}
