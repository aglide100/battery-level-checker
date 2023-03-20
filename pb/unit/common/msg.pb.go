// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: pb/unit/common/msg.proto

package common

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

type ReturnMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ReturnMsg) Reset() {
	*x = ReturnMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_unit_common_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReturnMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnMsg) ProtoMessage() {}

func (x *ReturnMsg) ProtoReflect() protoreflect.Message {
	mi := &file_pb_unit_common_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnMsg.ProtoReflect.Descriptor instead.
func (*ReturnMsg) Descriptor() ([]byte, []int) {
	return file_pb_unit_common_msg_proto_rawDescGZIP(), []int{0}
}

func (x *ReturnMsg) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *ReturnMsg) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pb_unit_common_msg_proto protoreflect.FileDescriptor

var file_pb_unit_common_msg_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x62, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x62, 0x2e, 0x75,
	0x6e, 0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x39, 0x0a, 0x09, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x79, 0x65, 0x6f, 0x6c, 0x2d, 0x69, 0x2f, 0x62, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x79, 0x2d, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x2d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65,
	0x72, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_unit_common_msg_proto_rawDescOnce sync.Once
	file_pb_unit_common_msg_proto_rawDescData = file_pb_unit_common_msg_proto_rawDesc
)

func file_pb_unit_common_msg_proto_rawDescGZIP() []byte {
	file_pb_unit_common_msg_proto_rawDescOnce.Do(func() {
		file_pb_unit_common_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_unit_common_msg_proto_rawDescData)
	})
	return file_pb_unit_common_msg_proto_rawDescData
}

var file_pb_unit_common_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pb_unit_common_msg_proto_goTypes = []interface{}{
	(*ReturnMsg)(nil), // 0: pb.unit.common.ReturnMsg
}
var file_pb_unit_common_msg_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_unit_common_msg_proto_init() }
func file_pb_unit_common_msg_proto_init() {
	if File_pb_unit_common_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_unit_common_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReturnMsg); i {
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
			RawDescriptor: file_pb_unit_common_msg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_unit_common_msg_proto_goTypes,
		DependencyIndexes: file_pb_unit_common_msg_proto_depIdxs,
		MessageInfos:      file_pb_unit_common_msg_proto_msgTypes,
	}.Build()
	File_pb_unit_common_msg_proto = out.File
	file_pb_unit_common_msg_proto_rawDesc = nil
	file_pb_unit_common_msg_proto_goTypes = nil
	file_pb_unit_common_msg_proto_depIdxs = nil
}
