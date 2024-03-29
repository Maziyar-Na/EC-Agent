// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: deploy.proto

package ec_rpc

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
type ExportPodSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DockerId string `protobuf:"bytes,1,opt,name=docker_id,json=dockerId,proto3" json:"docker_id,omitempty"`
	CgroupId int32  `protobuf:"varint,2,opt,name=cgroup_id,json=cgroupId,proto3" json:"cgroup_id,omitempty"`
	NodeIp   string `protobuf:"bytes,3,opt,name=node_ip,json=nodeIp,proto3" json:"node_ip,omitempty"`
}

func (x *ExportPodSpec) Reset() {
	*x = ExportPodSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deploy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportPodSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportPodSpec) ProtoMessage() {}

func (x *ExportPodSpec) ProtoReflect() protoreflect.Message {
	mi := &file_deploy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportPodSpec.ProtoReflect.Descriptor instead.
func (*ExportPodSpec) Descriptor() ([]byte, []int) {
	return file_deploy_proto_rawDescGZIP(), []int{0}
}

func (x *ExportPodSpec) GetDockerId() string {
	if x != nil {
		return x.DockerId
	}
	return ""
}

func (x *ExportPodSpec) GetCgroupId() int32 {
	if x != nil {
		return x.CgroupId
	}
	return 0
}

func (x *ExportPodSpec) GetNodeIp() string {
	if x != nil {
		return x.NodeIp
	}
	return ""
}

// The response message containing the greetings
type PodSpecReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DockerId string `protobuf:"bytes,1,opt,name=docker_id,json=dockerId,proto3" json:"docker_id,omitempty"`
	CgroupId int32  `protobuf:"varint,2,opt,name=cgroup_id,json=cgroupId,proto3" json:"cgroup_id,omitempty"`
	NodeIp   string `protobuf:"bytes,3,opt,name=node_ip,json=nodeIp,proto3" json:"node_ip,omitempty"`
	Thanks   string `protobuf:"bytes,4,opt,name=thanks,proto3" json:"thanks,omitempty"`
}

func (x *PodSpecReply) Reset() {
	*x = PodSpecReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deploy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodSpecReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodSpecReply) ProtoMessage() {}

func (x *PodSpecReply) ProtoReflect() protoreflect.Message {
	mi := &file_deploy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodSpecReply.ProtoReflect.Descriptor instead.
func (*PodSpecReply) Descriptor() ([]byte, []int) {
	return file_deploy_proto_rawDescGZIP(), []int{1}
}

func (x *PodSpecReply) GetDockerId() string {
	if x != nil {
		return x.DockerId
	}
	return ""
}

func (x *PodSpecReply) GetCgroupId() int32 {
	if x != nil {
		return x.CgroupId
	}
	return 0
}

func (x *PodSpecReply) GetNodeIp() string {
	if x != nil {
		return x.NodeIp
	}
	return ""
}

func (x *PodSpecReply) GetThanks() string {
	if x != nil {
		return x.Thanks
	}
	return ""
}

type ExportDeletePod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DockerId string `protobuf:"bytes,1,opt,name=docker_id,json=dockerId,proto3" json:"docker_id,omitempty"`
}

func (x *ExportDeletePod) Reset() {
	*x = ExportDeletePod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deploy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportDeletePod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportDeletePod) ProtoMessage() {}

func (x *ExportDeletePod) ProtoReflect() protoreflect.Message {
	mi := &file_deploy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportDeletePod.ProtoReflect.Descriptor instead.
func (*ExportDeletePod) Descriptor() ([]byte, []int) {
	return file_deploy_proto_rawDescGZIP(), []int{2}
}

func (x *ExportDeletePod) GetDockerId() string {
	if x != nil {
		return x.DockerId
	}
	return ""
}

type DeletePodReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DockerId string `protobuf:"bytes,1,opt,name=docker_id,json=dockerId,proto3" json:"docker_id,omitempty"`
	Thanks   string `protobuf:"bytes,4,opt,name=thanks,proto3" json:"thanks,omitempty"`
}

func (x *DeletePodReply) Reset() {
	*x = DeletePodReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deploy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePodReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePodReply) ProtoMessage() {}

func (x *DeletePodReply) ProtoReflect() protoreflect.Message {
	mi := &file_deploy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePodReply.ProtoReflect.Descriptor instead.
func (*DeletePodReply) Descriptor() ([]byte, []int) {
	return file_deploy_proto_rawDescGZIP(), []int{3}
}

func (x *DeletePodReply) GetDockerId() string {
	if x != nil {
		return x.DockerId
	}
	return ""
}

func (x *DeletePodReply) GetThanks() string {
	if x != nil {
		return x.Thanks
	}
	return ""
}

// The request message containing App Specs
type ExportAppSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppName  string `protobuf:"bytes,1,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	CpuLimit uint64 `protobuf:"varint,2,opt,name=cpu_limit,json=cpuLimit,proto3" json:"cpu_limit,omitempty"`
	MemLimit uint64 `protobuf:"varint,3,opt,name=mem_limit,json=memLimit,proto3" json:"mem_limit,omitempty"`
}

func (x *ExportAppSpec) Reset() {
	*x = ExportAppSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deploy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportAppSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportAppSpec) ProtoMessage() {}

func (x *ExportAppSpec) ProtoReflect() protoreflect.Message {
	mi := &file_deploy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportAppSpec.ProtoReflect.Descriptor instead.
func (*ExportAppSpec) Descriptor() ([]byte, []int) {
	return file_deploy_proto_rawDescGZIP(), []int{4}
}

func (x *ExportAppSpec) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *ExportAppSpec) GetCpuLimit() uint64 {
	if x != nil {
		return x.CpuLimit
	}
	return 0
}

func (x *ExportAppSpec) GetMemLimit() uint64 {
	if x != nil {
		return x.MemLimit
	}
	return 0
}

// The AppSpecReply message containing confirmation of App Specs
type AppSpecReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppName   string `protobuf:"bytes,1,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	Cpu_Limit uint64 `protobuf:"varint,2,opt,name=cpu_Limit,json=cpuLimit,proto3" json:"cpu_Limit,omitempty"`
	MemLimit  uint64 `protobuf:"varint,3,opt,name=mem_limit,json=memLimit,proto3" json:"mem_limit,omitempty"`
	Thanks    string `protobuf:"bytes,4,opt,name=thanks,proto3" json:"thanks,omitempty"`
}

func (x *AppSpecReply) Reset() {
	*x = AppSpecReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deploy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppSpecReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppSpecReply) ProtoMessage() {}

func (x *AppSpecReply) ProtoReflect() protoreflect.Message {
	mi := &file_deploy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppSpecReply.ProtoReflect.Descriptor instead.
func (*AppSpecReply) Descriptor() ([]byte, []int) {
	return file_deploy_proto_rawDescGZIP(), []int{5}
}

func (x *AppSpecReply) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *AppSpecReply) GetCpu_Limit() uint64 {
	if x != nil {
		return x.Cpu_Limit
	}
	return 0
}

func (x *AppSpecReply) GetMemLimit() uint64 {
	if x != nil {
		return x.MemLimit
	}
	return 0
}

func (x *AppSpecReply) GetThanks() string {
	if x != nil {
		return x.Thanks
	}
	return ""
}

var File_deploy_proto protoreflect.FileDescriptor

var file_deploy_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x65, 0x63, 0x2e, 0x72, 0x70, 0x63, 0x22, 0x62, 0x0a, 0x0d, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x50, 0x6f, 0x64, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x70, 0x22, 0x79, 0x0a, 0x0c, 0x50, 0x6f,
	0x64, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f,
	0x63, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x70, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x68, 0x61, 0x6e, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x68, 0x61, 0x6e, 0x6b, 0x73, 0x22, 0x2e, 0x0a, 0x0f, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x6f, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x68, 0x61, 0x6e, 0x6b, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x68, 0x61, 0x6e, 0x6b, 0x73, 0x22, 0x64, 0x0a, 0x0d,
	0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x70, 0x70, 0x53, 0x70, 0x65, 0x63, 0x12, 0x19, 0x0a,
	0x08, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x70, 0x75, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63, 0x70, 0x75,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x7b, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x70, 0x75, 0x5f, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x63, 0x70, 0x75, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65,
	0x6d, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6d,
	0x65, 0x6d, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x68, 0x61, 0x6e, 0x6b,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x68, 0x61, 0x6e, 0x6b, 0x73, 0x32,
	0xd0, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x45, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x6f, 0x64, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x15, 0x2e, 0x65, 0x63, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x50, 0x6f, 0x64, 0x53, 0x70, 0x65, 0x63, 0x1a, 0x14, 0x2e, 0x65, 0x63, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x50, 0x6f, 0x64, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x3e, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x12,
	0x17, 0x2e, 0x65, 0x63, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x1a, 0x16, 0x2e, 0x65, 0x63, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x70, 0x70, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x15, 0x2e, 0x65, 0x63, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x41, 0x70, 0x70, 0x53, 0x70, 0x65, 0x63, 0x1a, 0x14, 0x2e, 0x65, 0x63, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x41, 0x70, 0x70, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x65, 0x63, 0x2e, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deploy_proto_rawDescOnce sync.Once
	file_deploy_proto_rawDescData = file_deploy_proto_rawDesc
)

func file_deploy_proto_rawDescGZIP() []byte {
	file_deploy_proto_rawDescOnce.Do(func() {
		file_deploy_proto_rawDescData = protoimpl.X.CompressGZIP(file_deploy_proto_rawDescData)
	})
	return file_deploy_proto_rawDescData
}

var file_deploy_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_deploy_proto_goTypes = []interface{}{
	(*ExportPodSpec)(nil),   // 0: ec.rpc.ExportPodSpec
	(*PodSpecReply)(nil),    // 1: ec.rpc.PodSpecReply
	(*ExportDeletePod)(nil), // 2: ec.rpc.ExportDeletePod
	(*DeletePodReply)(nil),  // 3: ec.rpc.DeletePodReply
	(*ExportAppSpec)(nil),   // 4: ec.rpc.ExportAppSpec
	(*AppSpecReply)(nil),    // 5: ec.rpc.AppSpecReply
}
var file_deploy_proto_depIdxs = []int32{
	0, // 0: ec.rpc.DeployerExport.ReportPodSpec:input_type -> ec.rpc.ExportPodSpec
	2, // 1: ec.rpc.DeployerExport.DeletePod:input_type -> ec.rpc.ExportDeletePod
	4, // 2: ec.rpc.DeployerExport.ReportAppSpec:input_type -> ec.rpc.ExportAppSpec
	1, // 3: ec.rpc.DeployerExport.ReportPodSpec:output_type -> ec.rpc.PodSpecReply
	3, // 4: ec.rpc.DeployerExport.DeletePod:output_type -> ec.rpc.DeletePodReply
	5, // 5: ec.rpc.DeployerExport.ReportAppSpec:output_type -> ec.rpc.AppSpecReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_deploy_proto_init() }
func file_deploy_proto_init() {
	if File_deploy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deploy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportPodSpec); i {
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
		file_deploy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodSpecReply); i {
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
		file_deploy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportDeletePod); i {
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
		file_deploy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePodReply); i {
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
		file_deploy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportAppSpec); i {
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
		file_deploy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppSpecReply); i {
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
			RawDescriptor: file_deploy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_deploy_proto_goTypes,
		DependencyIndexes: file_deploy_proto_depIdxs,
		MessageInfos:      file_deploy_proto_msgTypes,
	}.Build()
	File_deploy_proto = out.File
	file_deploy_proto_rawDesc = nil
	file_deploy_proto_goTypes = nil
	file_deploy_proto_depIdxs = nil
}
