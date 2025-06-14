// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: proto/kafka.proto

package kafka_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type KafkaEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TypeEvent     []byte                 `protobuf:"bytes,1,opt,name=TypeEvent,proto3" json:"TypeEvent,omitempty"`
	Payload       []byte                 `protobuf:"bytes,2,opt,name=Payload,proto3" json:"Payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KafkaEvent) Reset() {
	*x = KafkaEvent{}
	mi := &file_proto_kafka_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KafkaEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KafkaEvent) ProtoMessage() {}

func (x *KafkaEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kafka_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KafkaEvent.ProtoReflect.Descriptor instead.
func (*KafkaEvent) Descriptor() ([]byte, []int) {
	return file_proto_kafka_proto_rawDescGZIP(), []int{0}
}

func (x *KafkaEvent) GetTypeEvent() []byte {
	if x != nil {
		return x.TypeEvent
	}
	return nil
}

func (x *KafkaEvent) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type PublishRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topic         string                 `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Event         *KafkaEvent            `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	mi := &file_proto_kafka_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kafka_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_proto_kafka_proto_rawDescGZIP(), []int{1}
}

func (x *PublishRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *PublishRequest) GetEvent() *KafkaEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

type PublishResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PublishResponse) Reset() {
	*x = PublishResponse{}
	mi := &file_proto_kafka_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishResponse) ProtoMessage() {}

func (x *PublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kafka_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishResponse.ProtoReflect.Descriptor instead.
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return file_proto_kafka_proto_rawDescGZIP(), []int{2}
}

func (x *PublishResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_kafka_proto protoreflect.FileDescriptor

const file_proto_kafka_proto_rawDesc = "" +
	"\n" +
	"\x11proto/kafka.proto\x12\x05kafka\"D\n" +
	"\n" +
	"KafkaEvent\x12\x1c\n" +
	"\tTypeEvent\x18\x01 \x01(\fR\tTypeEvent\x12\x18\n" +
	"\aPayload\x18\x02 \x01(\fR\aPayload\"O\n" +
	"\x0ePublishRequest\x12\x14\n" +
	"\x05topic\x18\x01 \x01(\tR\x05topic\x12'\n" +
	"\x05event\x18\x02 \x01(\v2\x11.kafka.KafkaEventR\x05event\"+\n" +
	"\x0fPublishResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage2H\n" +
	"\fKafkaService\x128\n" +
	"\aPublish\x12\x15.kafka.PublishRequest\x1a\x16.kafka.PublishResponseB\x10Z\x0eproto/kafka_pbb\x06proto3"

var (
	file_proto_kafka_proto_rawDescOnce sync.Once
	file_proto_kafka_proto_rawDescData []byte
)

func file_proto_kafka_proto_rawDescGZIP() []byte {
	file_proto_kafka_proto_rawDescOnce.Do(func() {
		file_proto_kafka_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_kafka_proto_rawDesc), len(file_proto_kafka_proto_rawDesc)))
	})
	return file_proto_kafka_proto_rawDescData
}

var file_proto_kafka_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_kafka_proto_goTypes = []any{
	(*KafkaEvent)(nil),      // 0: kafka.KafkaEvent
	(*PublishRequest)(nil),  // 1: kafka.PublishRequest
	(*PublishResponse)(nil), // 2: kafka.PublishResponse
}
var file_proto_kafka_proto_depIdxs = []int32{
	0, // 0: kafka.PublishRequest.event:type_name -> kafka.KafkaEvent
	1, // 1: kafka.KafkaService.Publish:input_type -> kafka.PublishRequest
	2, // 2: kafka.KafkaService.Publish:output_type -> kafka.PublishResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_kafka_proto_init() }
func file_proto_kafka_proto_init() {
	if File_proto_kafka_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_kafka_proto_rawDesc), len(file_proto_kafka_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_kafka_proto_goTypes,
		DependencyIndexes: file_proto_kafka_proto_depIdxs,
		MessageInfos:      file_proto_kafka_proto_msgTypes,
	}.Build()
	File_proto_kafka_proto = out.File
	file_proto_kafka_proto_goTypes = nil
	file_proto_kafka_proto_depIdxs = nil
}
