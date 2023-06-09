// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.2
// source: snake.proto

package game_system

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

type StartGameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AId    int32 `protobuf:"varint,1,opt,name=aId,proto3" json:"aId,omitempty"`
	ABotId int32 `protobuf:"varint,2,opt,name=aBotId,proto3" json:"aBotId,omitempty"`
	BId    int32 `protobuf:"varint,3,opt,name=bId,proto3" json:"bId,omitempty"`
	BBotId int32 `protobuf:"varint,4,opt,name=bBotId,proto3" json:"bBotId,omitempty"`
}

func (x *StartGameReq) Reset() {
	*x = StartGameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snake_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartGameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartGameReq) ProtoMessage() {}

func (x *StartGameReq) ProtoReflect() protoreflect.Message {
	mi := &file_snake_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartGameReq.ProtoReflect.Descriptor instead.
func (*StartGameReq) Descriptor() ([]byte, []int) {
	return file_snake_proto_rawDescGZIP(), []int{0}
}

func (x *StartGameReq) GetAId() int32 {
	if x != nil {
		return x.AId
	}
	return 0
}

func (x *StartGameReq) GetABotId() int32 {
	if x != nil {
		return x.ABotId
	}
	return 0
}

func (x *StartGameReq) GetBId() int32 {
	if x != nil {
		return x.BId
	}
	return 0
}

func (x *StartGameReq) GetBBotId() int32 {
	if x != nil {
		return x.BBotId
	}
	return 0
}

type StartGameResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StartGameResp) Reset() {
	*x = StartGameResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snake_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartGameResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartGameResp) ProtoMessage() {}

func (x *StartGameResp) ProtoReflect() protoreflect.Message {
	mi := &file_snake_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartGameResp.ProtoReflect.Descriptor instead.
func (*StartGameResp) Descriptor() ([]byte, []int) {
	return file_snake_proto_rawDescGZIP(), []int{1}
}

func (x *StartGameResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SetNextStepReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Direction string `protobuf:"bytes,1,opt,name=direction,proto3" json:"direction,omitempty"`
	PlayerId  string `protobuf:"bytes,2,opt,name=PlayerId,proto3" json:"PlayerId,omitempty"`
	IsCode    bool   `protobuf:"varint,3,opt,name=IsCode,proto3" json:"IsCode,omitempty"`
}

func (x *SetNextStepReq) Reset() {
	*x = SetNextStepReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snake_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetNextStepReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetNextStepReq) ProtoMessage() {}

func (x *SetNextStepReq) ProtoReflect() protoreflect.Message {
	mi := &file_snake_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetNextStepReq.ProtoReflect.Descriptor instead.
func (*SetNextStepReq) Descriptor() ([]byte, []int) {
	return file_snake_proto_rawDescGZIP(), []int{2}
}

func (x *SetNextStepReq) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

func (x *SetNextStepReq) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *SetNextStepReq) GetIsCode() bool {
	if x != nil {
		return x.IsCode
	}
	return false
}

type SetNextStepResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event      string `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	ADirection string `protobuf:"bytes,2,opt,name=aDirection,proto3" json:"aDirection,omitempty"`
	BDirection string `protobuf:"bytes,3,opt,name=bDirection,proto3" json:"bDirection,omitempty"`
}

func (x *SetNextStepResp) Reset() {
	*x = SetNextStepResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snake_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetNextStepResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetNextStepResp) ProtoMessage() {}

func (x *SetNextStepResp) ProtoReflect() protoreflect.Message {
	mi := &file_snake_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetNextStepResp.ProtoReflect.Descriptor instead.
func (*SetNextStepResp) Descriptor() ([]byte, []int) {
	return file_snake_proto_rawDescGZIP(), []int{3}
}

func (x *SetNextStepResp) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *SetNextStepResp) GetADirection() string {
	if x != nil {
		return x.ADirection
	}
	return ""
}

func (x *SetNextStepResp) GetBDirection() string {
	if x != nil {
		return x.BDirection
	}
	return ""
}

var File_snake_proto protoreflect.FileDescriptor

var file_snake_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x6e, 0x61, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67,
	0x61, 0x6d, 0x65, 0x2e, 0x70, 0x62, 0x22, 0x62, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x42, 0x6f, 0x74,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x42, 0x6f, 0x74, 0x49, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x62, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x62,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x42, 0x6f, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x62, 0x42, 0x6f, 0x74, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x0d, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x62, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x4e, 0x65, 0x78, 0x74,
	0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x49, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x67, 0x0a, 0x0f, 0x53, 0x65, 0x74,
	0x4e, 0x65, 0x78, 0x74, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x32, 0x8f, 0x01, 0x0a, 0x0b, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x12, 0x3c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x12,
	0x15, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x42, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x4e, 0x65, 0x78, 0x74, 0x53, 0x74, 0x65, 0x70, 0x12,
	0x17, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x74, 0x4e, 0x65, 0x78,
	0x74, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x65, 0x74, 0x4e, 0x65, 0x78, 0x74, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x3b, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_snake_proto_rawDescOnce sync.Once
	file_snake_proto_rawDescData = file_snake_proto_rawDesc
)

func file_snake_proto_rawDescGZIP() []byte {
	file_snake_proto_rawDescOnce.Do(func() {
		file_snake_proto_rawDescData = protoimpl.X.CompressGZIP(file_snake_proto_rawDescData)
	})
	return file_snake_proto_rawDescData
}

var file_snake_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_snake_proto_goTypes = []interface{}{
	(*StartGameReq)(nil),    // 0: game.pb.StartGameReq
	(*StartGameResp)(nil),   // 1: game.pb.StartGameResp
	(*SetNextStepReq)(nil),  // 2: game.pb.SetNextStepReq
	(*SetNextStepResp)(nil), // 3: game.pb.SetNextStepResp
}
var file_snake_proto_depIdxs = []int32{
	0, // 0: game.pb.game_system.StartGame:input_type -> game.pb.StartGameReq
	2, // 1: game.pb.game_system.SetNextStep:input_type -> game.pb.SetNextStepReq
	1, // 2: game.pb.game_system.StartGame:output_type -> game.pb.StartGameResp
	3, // 3: game.pb.game_system.SetNextStep:output_type -> game.pb.SetNextStepResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_snake_proto_init() }
func file_snake_proto_init() {
	if File_snake_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_snake_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartGameReq); i {
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
		file_snake_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartGameResp); i {
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
		file_snake_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetNextStepReq); i {
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
		file_snake_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetNextStepResp); i {
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
			RawDescriptor: file_snake_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_snake_proto_goTypes,
		DependencyIndexes: file_snake_proto_depIdxs,
		MessageInfos:      file_snake_proto_msgTypes,
	}.Build()
	File_snake_proto = out.File
	file_snake_proto_rawDesc = nil
	file_snake_proto_goTypes = nil
	file_snake_proto_depIdxs = nil
}
