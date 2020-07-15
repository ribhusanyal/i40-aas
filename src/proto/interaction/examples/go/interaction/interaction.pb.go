// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.11.2
// source: interaction.proto

package interaction

import (
	proto "github.com/golang/protobuf/proto"
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

type InteractionStatusCode int32

const (
	InteractionStatusCode_UNKNOWN InteractionStatusCode = 0
	InteractionStatusCode_OK      InteractionStatusCode = 200
	InteractionStatusCode_FAILED  InteractionStatusCode = 500
)

// Enum value maps for InteractionStatusCode.
var (
	InteractionStatusCode_name = map[int32]string{
		0:   "UNKNOWN",
		200: "OK",
		500: "FAILED",
	}
	InteractionStatusCode_value = map[string]int32{
		"UNKNOWN": 0,
		"OK":      200,
		"FAILED":  500,
	}
)

func (x InteractionStatusCode) Enum() *InteractionStatusCode {
	p := new(InteractionStatusCode)
	*p = x
	return p
}

func (x InteractionStatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InteractionStatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_interaction_proto_enumTypes[0].Descriptor()
}

func (InteractionStatusCode) Type() protoreflect.EnumType {
	return &file_interaction_proto_enumTypes[0]
}

func (x InteractionStatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InteractionStatusCode.Descriptor instead.
func (InteractionStatusCode) EnumDescriptor() ([]byte, []int) {
	return file_interaction_proto_rawDescGZIP(), []int{0}
}

type InteractionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code InteractionStatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=InteractionStatusCode" json:"code,omitempty"`
}

func (x *InteractionStatus) Reset() {
	*x = InteractionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InteractionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InteractionStatus) ProtoMessage() {}

func (x *InteractionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_interaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InteractionStatus.ProtoReflect.Descriptor instead.
func (*InteractionStatus) Descriptor() ([]byte, []int) {
	return file_interaction_proto_rawDescGZIP(), []int{0}
}

func (x *InteractionStatus) GetCode() InteractionStatusCode {
	if x != nil {
		return x.Code
	}
	return InteractionStatusCode_UNKNOWN
}

type InteractionMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Frame               *Frame `protobuf:"bytes,1,opt,name=frame,proto3" json:"frame,omitempty"`
	InteractionElements []byte `protobuf:"bytes,2,opt,name=interactionElements,proto3" json:"interactionElements,omitempty"`
}

func (x *InteractionMessage) Reset() {
	*x = InteractionMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InteractionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InteractionMessage) ProtoMessage() {}

func (x *InteractionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_interaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InteractionMessage.ProtoReflect.Descriptor instead.
func (*InteractionMessage) Descriptor() ([]byte, []int) {
	return file_interaction_proto_rawDescGZIP(), []int{1}
}

func (x *InteractionMessage) GetFrame() *Frame {
	if x != nil {
		return x.Frame
	}
	return nil
}

func (x *InteractionMessage) GetInteractionElements() []byte {
	if x != nil {
		return x.InteractionElements
	}
	return nil
}

type Frame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SemanticProtocol string              `protobuf:"bytes,1,opt,name=semanticProtocol,proto3" json:"semanticProtocol,omitempty"`
	Type             string              `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	MessageId        string              `protobuf:"bytes,3,opt,name=messageId,proto3" json:"messageId,omitempty"`
	ReplyBy          uint32              `protobuf:"varint,4,opt,name=replyBy,proto3" json:"replyBy,omitempty"`
	Receiver         *ConversationMember `protobuf:"bytes,5,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Sender           *ConversationMember `protobuf:"bytes,6,opt,name=sender,proto3" json:"sender,omitempty"`
	ConversationId   string              `protobuf:"bytes,7,opt,name=conversationId,proto3" json:"conversationId,omitempty"`
}

func (x *Frame) Reset() {
	*x = Frame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Frame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Frame) ProtoMessage() {}

func (x *Frame) ProtoReflect() protoreflect.Message {
	mi := &file_interaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Frame.ProtoReflect.Descriptor instead.
func (*Frame) Descriptor() ([]byte, []int) {
	return file_interaction_proto_rawDescGZIP(), []int{2}
}

func (x *Frame) GetSemanticProtocol() string {
	if x != nil {
		return x.SemanticProtocol
	}
	return ""
}

func (x *Frame) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Frame) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *Frame) GetReplyBy() uint32 {
	if x != nil {
		return x.ReplyBy
	}
	return 0
}

func (x *Frame) GetReceiver() *ConversationMember {
	if x != nil {
		return x.Receiver
	}
	return nil
}

func (x *Frame) GetSender() *ConversationMember {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *Frame) GetConversationId() string {
	if x != nil {
		return x.ConversationId
	}
	return ""
}

type ConversationMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identification *Identification `protobuf:"bytes,1,opt,name=identification,proto3" json:"identification,omitempty"`
	Role           *Role           `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *ConversationMember) Reset() {
	*x = ConversationMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interaction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversationMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversationMember) ProtoMessage() {}

func (x *ConversationMember) ProtoReflect() protoreflect.Message {
	mi := &file_interaction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversationMember.ProtoReflect.Descriptor instead.
func (*ConversationMember) Descriptor() ([]byte, []int) {
	return file_interaction_proto_rawDescGZIP(), []int{3}
}

func (x *ConversationMember) GetIdentification() *Identification {
	if x != nil {
		return x.Identification
	}
	return nil
}

func (x *ConversationMember) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interaction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_interaction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_interaction_proto_rawDescGZIP(), []int{4}
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Identification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IdType string `protobuf:"bytes,2,opt,name=idType,proto3" json:"idType,omitempty"`
}

func (x *Identification) Reset() {
	*x = Identification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interaction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identification) ProtoMessage() {}

func (x *Identification) ProtoReflect() protoreflect.Message {
	mi := &file_interaction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identification.ProtoReflect.Descriptor instead.
func (*Identification) Descriptor() ([]byte, []int) {
	return file_interaction_proto_rawDescGZIP(), []int{5}
}

func (x *Identification) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Identification) GetIdType() string {
	if x != nil {
		return x.IdType
	}
	return ""
}

var File_interaction_proto protoreflect.FileDescriptor

var file_interaction_proto_rawDesc = []byte{
	0x0a, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x11, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x22, 0x64, 0x0a, 0x12, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x46, 0x72, 0x61, 0x6d,
	0x65, 0x52, 0x05, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x13, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x85, 0x02, 0x0a, 0x05, 0x46,
	0x72, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x79, 0x12, 0x2f, 0x0a, 0x08,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x2b, 0x0a,
	0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0x68, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x0e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x19, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x1a, 0x0a, 0x04,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x0e, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x2a, 0x3a, 0x0a, 0x15, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0xc8,
	0x01, 0x12, 0x0b, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xf4, 0x03, 0x32, 0x59,
	0x0a, 0x12, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x43, 0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x13,
	0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x12, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x70, 0x2f, 0x69, 0x34, 0x30, 0x2d,
	0x61, 0x61, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interaction_proto_rawDescOnce sync.Once
	file_interaction_proto_rawDescData = file_interaction_proto_rawDesc
)

func file_interaction_proto_rawDescGZIP() []byte {
	file_interaction_proto_rawDescOnce.Do(func() {
		file_interaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_interaction_proto_rawDescData)
	})
	return file_interaction_proto_rawDescData
}

var file_interaction_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_interaction_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_interaction_proto_goTypes = []interface{}{
	(InteractionStatusCode)(0), // 0: InteractionStatusCode
	(*InteractionStatus)(nil),  // 1: InteractionStatus
	(*InteractionMessage)(nil), // 2: InteractionMessage
	(*Frame)(nil),              // 3: Frame
	(*ConversationMember)(nil), // 4: ConversationMember
	(*Role)(nil),               // 5: Role
	(*Identification)(nil),     // 6: Identification
}
var file_interaction_proto_depIdxs = []int32{
	0, // 0: InteractionStatus.code:type_name -> InteractionStatusCode
	3, // 1: InteractionMessage.frame:type_name -> Frame
	4, // 2: Frame.receiver:type_name -> ConversationMember
	4, // 3: Frame.sender:type_name -> ConversationMember
	6, // 4: ConversationMember.identification:type_name -> Identification
	5, // 5: ConversationMember.role:type_name -> Role
	2, // 6: InteractionIngress.SendInteractionMessage:input_type -> InteractionMessage
	1, // 7: InteractionIngress.SendInteractionMessage:output_type -> InteractionStatus
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_interaction_proto_init() }
func file_interaction_proto_init() {
	if File_interaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InteractionStatus); i {
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
		file_interaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InteractionMessage); i {
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
		file_interaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Frame); i {
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
		file_interaction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversationMember); i {
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
		file_interaction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_interaction_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identification); i {
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
			RawDescriptor: file_interaction_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_interaction_proto_goTypes,
		DependencyIndexes: file_interaction_proto_depIdxs,
		EnumInfos:         file_interaction_proto_enumTypes,
		MessageInfos:      file_interaction_proto_msgTypes,
	}.Build()
	File_interaction_proto = out.File
	file_interaction_proto_rawDesc = nil
	file_interaction_proto_goTypes = nil
	file_interaction_proto_depIdxs = nil
}
