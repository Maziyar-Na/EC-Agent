// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: agent.proto

package agentcomm

import (
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

// The request message containing the user's name.
type ConnectContainerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GcmIP    string `protobuf:"bytes,1,opt,name=gcmIP,proto3" json:"gcmIP,omitempty"`
	PodName  string `protobuf:"bytes,2,opt,name=podName,proto3" json:"podName,omitempty"`
	DockerId string `protobuf:"bytes,3,opt,name=dockerId,proto3" json:"dockerId,omitempty"`
	AppNum   int32  `protobuf:"varint,4,opt,name=appNum,proto3" json:"appNum,omitempty"`
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

func (x *ConnectContainerRequest) GetAppNum() int32 {
	if x != nil {
		return x.AppNum
	}
	return 0
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

type TriggerPodDeploymentWatcherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GcmIP     string `protobuf:"bytes,1,opt,name=gcmIP,proto3" json:"gcmIP,omitempty"`
	AgentIP   string `protobuf:"bytes,2,opt,name=agentIP,proto3" json:"agentIP,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	AppCount  int32  `protobuf:"varint,4,opt,name=appCount,proto3" json:"appCount,omitempty"`
}

func (x *TriggerPodDeploymentWatcherRequest) Reset() {
	*x = TriggerPodDeploymentWatcherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerPodDeploymentWatcherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerPodDeploymentWatcherRequest) ProtoMessage() {}

func (x *TriggerPodDeploymentWatcherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerPodDeploymentWatcherRequest.ProtoReflect.Descriptor instead.
func (*TriggerPodDeploymentWatcherRequest) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{2}
}

func (x *TriggerPodDeploymentWatcherRequest) GetGcmIP() string {
	if x != nil {
		return x.GcmIP
	}
	return ""
}

func (x *TriggerPodDeploymentWatcherRequest) GetAgentIP() string {
	if x != nil {
		return x.AgentIP
	}
	return ""
}

func (x *TriggerPodDeploymentWatcherRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *TriggerPodDeploymentWatcherRequest) GetAppCount() int32 {
	if x != nil {
		return x.AppCount
	}
	return 0
}

type TriggerPodDeploymentWatcherReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReturnStatus int32 `protobuf:"varint,1,opt,name=returnStatus,proto3" json:"returnStatus,omitempty"`
}

func (x *TriggerPodDeploymentWatcherReply) Reset() {
	*x = TriggerPodDeploymentWatcherReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerPodDeploymentWatcherReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerPodDeploymentWatcherReply) ProtoMessage() {}

func (x *TriggerPodDeploymentWatcherReply) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerPodDeploymentWatcherReply.ProtoReflect.Descriptor instead.
func (*TriggerPodDeploymentWatcherReply) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{3}
}

func (x *TriggerPodDeploymentWatcherReply) GetReturnStatus() int32 {
	if x != nil {
		return x.ReturnStatus
	}
	return 0
}

var File_agent_proto protoreflect.FileDescriptor

var file_agent_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x63, 0x6f, 0x6d, 0x6d, 0x22, 0x7d, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x63, 0x6d, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x67, 0x63, 0x6d, 0x49, 0x50, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6f, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x64, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x4e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x61, 0x70, 0x70, 0x4e, 0x75, 0x6d, 0x22, 0x69, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x6f,
	0x63, 0x6b, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f,
	0x63, 0x6b, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x44, 0x22, 0x8e, 0x01, 0x0a, 0x22, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x50, 0x6f,
	0x64, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x63, 0x6d,
	0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x63, 0x6d, 0x49, 0x50, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x50, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x70, 0x70, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x46, 0x0a, 0x20, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x50, 0x6f,
	0x64, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xe0, 0x01, 0x0a, 0x07,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x5d, 0x0a, 0x13, 0x52, 0x65, 0x71, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x22,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x76, 0x0a, 0x16, 0x52, 0x65, 0x71, 0x54, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x2d, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x54, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x50, 0x6f, 0x64, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x54, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x50, 0x6f, 0x64, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0d,
	0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x63, 0x6f, 0x6d, 0x6d, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_agent_proto_goTypes = []interface{}{
	(*ConnectContainerRequest)(nil),            // 0: agentcomm.ConnectContainerRequest
	(*ConnectContainerReply)(nil),              // 1: agentcomm.ConnectContainerReply
	(*TriggerPodDeploymentWatcherRequest)(nil), // 2: agentcomm.TriggerPodDeploymentWatcherRequest
	(*TriggerPodDeploymentWatcherReply)(nil),   // 3: agentcomm.TriggerPodDeploymentWatcherReply
}
var file_agent_proto_depIdxs = []int32{
	0, // 0: agentcomm.Handler.ReqConnectContainer:input_type -> agentcomm.ConnectContainerRequest
	2, // 1: agentcomm.Handler.ReqTriggerAgentWatcher:input_type -> agentcomm.TriggerPodDeploymentWatcherRequest
	1, // 2: agentcomm.Handler.ReqConnectContainer:output_type -> agentcomm.ConnectContainerReply
	3, // 3: agentcomm.Handler.ReqTriggerAgentWatcher:output_type -> agentcomm.TriggerPodDeploymentWatcherReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
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
		file_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerPodDeploymentWatcherRequest); i {
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
		file_agent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerPodDeploymentWatcherReply); i {
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
			NumMessages:   4,
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
