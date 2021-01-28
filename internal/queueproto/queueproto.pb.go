// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        (unknown)
// source: queueproto.proto

package queueproto

import (
	context "context"
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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queueproto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_queueproto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_queueproto_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type ResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors int32 `protobuf:"varint,1,opt,name=errors,proto3" json:"errors,omitempty"`
}

func (x *ResponseMessage) Reset() {
	*x = ResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queueproto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMessage) ProtoMessage() {}

func (x *ResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_queueproto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMessage.ProtoReflect.Descriptor instead.
func (*ResponseMessage) Descriptor() ([]byte, []int) {
	return file_queueproto_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseMessage) GetErrors() int32 {
	if x != nil {
		return x.Errors
	}
	return 0
}

type CreateStreamMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamName      string `protobuf:"bytes,1,opt,name=streamName,proto3" json:"streamName,omitempty"`
	PartitionsCount int32  `protobuf:"varint,2,opt,name=partitionsCount,proto3" json:"partitionsCount,omitempty"`
}

func (x *CreateStreamMessage) Reset() {
	*x = CreateStreamMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queueproto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStreamMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStreamMessage) ProtoMessage() {}

func (x *CreateStreamMessage) ProtoReflect() protoreflect.Message {
	mi := &file_queueproto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStreamMessage.ProtoReflect.Descriptor instead.
func (*CreateStreamMessage) Descriptor() ([]byte, []int) {
	return file_queueproto_proto_rawDescGZIP(), []int{2}
}

func (x *CreateStreamMessage) GetStreamName() string {
	if x != nil {
		return x.StreamName
	}
	return ""
}

func (x *CreateStreamMessage) GetPartitionsCount() int32 {
	if x != nil {
		return x.PartitionsCount
	}
	return 0
}

type DetailStreamMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamName string `protobuf:"bytes,1,opt,name=streamName,proto3" json:"streamName,omitempty"`
}

func (x *DetailStreamMessage) Reset() {
	*x = DetailStreamMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queueproto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailStreamMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailStreamMessage) ProtoMessage() {}

func (x *DetailStreamMessage) ProtoReflect() protoreflect.Message {
	mi := &file_queueproto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailStreamMessage.ProtoReflect.Descriptor instead.
func (*DetailStreamMessage) Descriptor() ([]byte, []int) {
	return file_queueproto_proto_rawDescGZIP(), []int{3}
}

func (x *DetailStreamMessage) GetStreamName() string {
	if x != nil {
		return x.StreamName
	}
	return ""
}

type DetailStreamResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamName      string `protobuf:"bytes,1,opt,name=streamName,proto3" json:"streamName,omitempty"`
	PartitionsCount string `protobuf:"bytes,2,opt,name=partitionsCount,proto3" json:"partitionsCount,omitempty"`
}

func (x *DetailStreamResponseMessage) Reset() {
	*x = DetailStreamResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queueproto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailStreamResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailStreamResponseMessage) ProtoMessage() {}

func (x *DetailStreamResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_queueproto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailStreamResponseMessage.ProtoReflect.Descriptor instead.
func (*DetailStreamResponseMessage) Descriptor() ([]byte, []int) {
	return file_queueproto_proto_rawDescGZIP(), []int{4}
}

func (x *DetailStreamResponseMessage) GetStreamName() string {
	if x != nil {
		return x.StreamName
	}
	return ""
}

func (x *DetailStreamResponseMessage) GetPartitionsCount() string {
	if x != nil {
		return x.PartitionsCount
	}
	return ""
}

type EnqueueMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamName string `protobuf:"bytes,1,opt,name=streamName,proto3" json:"streamName,omitempty"`
	Patition   int32  `protobuf:"varint,2,opt,name=patition,proto3" json:"patition,omitempty"`
	Payload    string `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *EnqueueMessage) Reset() {
	*x = EnqueueMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queueproto_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnqueueMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnqueueMessage) ProtoMessage() {}

func (x *EnqueueMessage) ProtoReflect() protoreflect.Message {
	mi := &file_queueproto_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnqueueMessage.ProtoReflect.Descriptor instead.
func (*EnqueueMessage) Descriptor() ([]byte, []int) {
	return file_queueproto_proto_rawDescGZIP(), []int{5}
}

func (x *EnqueueMessage) GetStreamName() string {
	if x != nil {
		return x.StreamName
	}
	return ""
}

func (x *EnqueueMessage) GetPatition() int32 {
	if x != nil {
		return x.Patition
	}
	return 0
}

func (x *EnqueueMessage) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type DequeueMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamName string `protobuf:"bytes,1,opt,name=streamName,proto3" json:"streamName,omitempty"`
	Patition   int32  `protobuf:"varint,2,opt,name=patition,proto3" json:"patition,omitempty"`
}

func (x *DequeueMessage) Reset() {
	*x = DequeueMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queueproto_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DequeueMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DequeueMessage) ProtoMessage() {}

func (x *DequeueMessage) ProtoReflect() protoreflect.Message {
	mi := &file_queueproto_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DequeueMessage.ProtoReflect.Descriptor instead.
func (*DequeueMessage) Descriptor() ([]byte, []int) {
	return file_queueproto_proto_rawDescGZIP(), []int{6}
}

func (x *DequeueMessage) GetStreamName() string {
	if x != nil {
		return x.StreamName
	}
	return ""
}

func (x *DequeueMessage) GetPatition() int32 {
	if x != nil {
		return x.Patition
	}
	return 0
}

type DequeueResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload string `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *DequeueResponseMessage) Reset() {
	*x = DequeueResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queueproto_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DequeueResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DequeueResponseMessage) ProtoMessage() {}

func (x *DequeueResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_queueproto_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DequeueResponseMessage.ProtoReflect.Descriptor instead.
func (*DequeueResponseMessage) Descriptor() ([]byte, []int) {
	return file_queueproto_proto_rawDescGZIP(), []int{7}
}

func (x *DequeueResponseMessage) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

var File_queueproto_proto protoreflect.FileDescriptor

var file_queueproto_proto_rawDesc = []byte{
	0x0a, 0x10, 0x71, 0x75, 0x65, 0x75, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x71, 0x75, 0x65, 0x75, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x29, 0x0a,
	0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x5f, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x28, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x35, 0x0a, 0x13, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x67, 0x0a, 0x1b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x28, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x66, 0x0a, 0x0e, 0x45, 0x6e, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x61, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0x4c, 0x0a, 0x0e, 0x44, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x32, 0x0a, 0x16, 0x44, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x32, 0x85, 0x03, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x12, 0x13, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x13, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1f, 0x2e, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1b, 0x2e,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0c,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1f, 0x2e, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x27, 0x2e,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x07, 0x45, 0x6e, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x12, 0x1a, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x1b, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x4b,
	0x0a, 0x07, 0x44, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x1a, 0x2e, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x22, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_queueproto_proto_rawDescOnce sync.Once
	file_queueproto_proto_rawDescData = file_queueproto_proto_rawDesc
)

func file_queueproto_proto_rawDescGZIP() []byte {
	file_queueproto_proto_rawDescOnce.Do(func() {
		file_queueproto_proto_rawDescData = protoimpl.X.CompressGZIP(file_queueproto_proto_rawDescData)
	})
	return file_queueproto_proto_rawDescData
}

var file_queueproto_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_queueproto_proto_goTypes = []interface{}{
	(*Message)(nil),                     // 0: queueproto.Message
	(*ResponseMessage)(nil),             // 1: queueproto.ResponseMessage
	(*CreateStreamMessage)(nil),         // 2: queueproto.CreateStreamMessage
	(*DetailStreamMessage)(nil),         // 3: queueproto.DetailStreamMessage
	(*DetailStreamResponseMessage)(nil), // 4: queueproto.DetailStreamResponseMessage
	(*EnqueueMessage)(nil),              // 5: queueproto.EnqueueMessage
	(*DequeueMessage)(nil),              // 6: queueproto.DequeueMessage
	(*DequeueResponseMessage)(nil),      // 7: queueproto.DequeueResponseMessage
}
var file_queueproto_proto_depIdxs = []int32{
	0, // 0: queueproto.QueueService.SayHello:input_type -> queueproto.Message
	2, // 1: queueproto.QueueService.CreateStream:input_type -> queueproto.CreateStreamMessage
	3, // 2: queueproto.QueueService.DetailStream:input_type -> queueproto.DetailStreamMessage
	5, // 3: queueproto.QueueService.Enqueue:input_type -> queueproto.EnqueueMessage
	6, // 4: queueproto.QueueService.Dequeue:input_type -> queueproto.DequeueMessage
	0, // 5: queueproto.QueueService.SayHello:output_type -> queueproto.Message
	1, // 6: queueproto.QueueService.CreateStream:output_type -> queueproto.ResponseMessage
	4, // 7: queueproto.QueueService.DetailStream:output_type -> queueproto.DetailStreamResponseMessage
	1, // 8: queueproto.QueueService.Enqueue:output_type -> queueproto.ResponseMessage
	7, // 9: queueproto.QueueService.Dequeue:output_type -> queueproto.DequeueResponseMessage
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_queueproto_proto_init() }
func file_queueproto_proto_init() {
	if File_queueproto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_queueproto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_queueproto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMessage); i {
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
		file_queueproto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStreamMessage); i {
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
		file_queueproto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailStreamMessage); i {
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
		file_queueproto_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailStreamResponseMessage); i {
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
		file_queueproto_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnqueueMessage); i {
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
		file_queueproto_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DequeueMessage); i {
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
		file_queueproto_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DequeueResponseMessage); i {
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
			RawDescriptor: file_queueproto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_queueproto_proto_goTypes,
		DependencyIndexes: file_queueproto_proto_depIdxs,
		MessageInfos:      file_queueproto_proto_msgTypes,
	}.Build()
	File_queueproto_proto = out.File
	file_queueproto_proto_rawDesc = nil
	file_queueproto_proto_goTypes = nil
	file_queueproto_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// QueueServiceClient is the client API for QueueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueueServiceClient interface {
	SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	CreateStream(ctx context.Context, in *CreateStreamMessage, opts ...grpc.CallOption) (*ResponseMessage, error)
	DetailStream(ctx context.Context, in *DetailStreamMessage, opts ...grpc.CallOption) (*DetailStreamResponseMessage, error)
	Enqueue(ctx context.Context, in *EnqueueMessage, opts ...grpc.CallOption) (*ResponseMessage, error)
	Dequeue(ctx context.Context, in *DequeueMessage, opts ...grpc.CallOption) (*DequeueResponseMessage, error)
}

type queueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQueueServiceClient(cc grpc.ClientConnInterface) QueueServiceClient {
	return &queueServiceClient{cc}
}

func (c *queueServiceClient) SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/queueproto.QueueService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueServiceClient) CreateStream(ctx context.Context, in *CreateStreamMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
	out := new(ResponseMessage)
	err := c.cc.Invoke(ctx, "/queueproto.QueueService/CreateStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueServiceClient) DetailStream(ctx context.Context, in *DetailStreamMessage, opts ...grpc.CallOption) (*DetailStreamResponseMessage, error) {
	out := new(DetailStreamResponseMessage)
	err := c.cc.Invoke(ctx, "/queueproto.QueueService/DetailStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueServiceClient) Enqueue(ctx context.Context, in *EnqueueMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
	out := new(ResponseMessage)
	err := c.cc.Invoke(ctx, "/queueproto.QueueService/Enqueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueServiceClient) Dequeue(ctx context.Context, in *DequeueMessage, opts ...grpc.CallOption) (*DequeueResponseMessage, error) {
	out := new(DequeueResponseMessage)
	err := c.cc.Invoke(ctx, "/queueproto.QueueService/Dequeue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueueServiceServer is the server API for QueueService service.
type QueueServiceServer interface {
	SayHello(context.Context, *Message) (*Message, error)
	CreateStream(context.Context, *CreateStreamMessage) (*ResponseMessage, error)
	DetailStream(context.Context, *DetailStreamMessage) (*DetailStreamResponseMessage, error)
	Enqueue(context.Context, *EnqueueMessage) (*ResponseMessage, error)
	Dequeue(context.Context, *DequeueMessage) (*DequeueResponseMessage, error)
}

// UnimplementedQueueServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueueServiceServer struct {
}

func (*UnimplementedQueueServiceServer) SayHello(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedQueueServiceServer) CreateStream(context.Context, *CreateStreamMessage) (*ResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStream not implemented")
}
func (*UnimplementedQueueServiceServer) DetailStream(context.Context, *DetailStreamMessage) (*DetailStreamResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetailStream not implemented")
}
func (*UnimplementedQueueServiceServer) Enqueue(context.Context, *EnqueueMessage) (*ResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enqueue not implemented")
}
func (*UnimplementedQueueServiceServer) Dequeue(context.Context, *DequeueMessage) (*DequeueResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dequeue not implemented")
}

func RegisterQueueServiceServer(s *grpc.Server, srv QueueServiceServer) {
	s.RegisterService(&_QueueService_serviceDesc, srv)
}

func _QueueService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queueproto.QueueService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServiceServer).SayHello(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueService_CreateStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStreamMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServiceServer).CreateStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queueproto.QueueService/CreateStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServiceServer).CreateStream(ctx, req.(*CreateStreamMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueService_DetailStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailStreamMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServiceServer).DetailStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queueproto.QueueService/DetailStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServiceServer).DetailStream(ctx, req.(*DetailStreamMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueService_Enqueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnqueueMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServiceServer).Enqueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queueproto.QueueService/Enqueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServiceServer).Enqueue(ctx, req.(*EnqueueMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueService_Dequeue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DequeueMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServiceServer).Dequeue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queueproto.QueueService/Dequeue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServiceServer).Dequeue(ctx, req.(*DequeueMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueueService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "queueproto.QueueService",
	HandlerType: (*QueueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _QueueService_SayHello_Handler,
		},
		{
			MethodName: "CreateStream",
			Handler:    _QueueService_CreateStream_Handler,
		},
		{
			MethodName: "DetailStream",
			Handler:    _QueueService_DetailStream_Handler,
		},
		{
			MethodName: "Enqueue",
			Handler:    _QueueService_Enqueue_Handler,
		},
		{
			MethodName: "Dequeue",
			Handler:    _QueueService_Dequeue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "queueproto.proto",
}
