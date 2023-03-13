// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: pb/unit/device/device.proto

package device

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

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           *ID           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Spec         *Spec         `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	BatteryLevel *BatteryLevel `protobuf:"bytes,3,opt,name=batteryLevel,proto3" json:"batteryLevel,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_unit_device_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_pb_unit_device_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_pb_unit_device_device_proto_rawDescGZIP(), []int{0}
}

func (x *Device) GetId() *ID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Device) GetSpec() *Spec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Device) GetBatteryLevel() *BatteryLevel {
	if x != nil {
		return x.BatteryLevel
	}
	return nil
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_unit_device_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_pb_unit_device_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_pb_unit_device_device_proto_rawDescGZIP(), []int{1}
}

func (x *ID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Spec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type       string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	OS         string `protobuf:"bytes,4,opt,name=OS,proto3" json:"OS,omitempty"`
	OsVersion  string `protobuf:"bytes,5,opt,name=osVersion,proto3" json:"osVersion,omitempty"`
	AppVersion string `protobuf:"bytes,6,opt,name=appVersion,proto3" json:"appVersion,omitempty"`
}

func (x *Spec) Reset() {
	*x = Spec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_unit_device_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spec) ProtoMessage() {}

func (x *Spec) ProtoReflect() protoreflect.Message {
	mi := &file_pb_unit_device_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spec.ProtoReflect.Descriptor instead.
func (*Spec) Descriptor() ([]byte, []int) {
	return file_pb_unit_device_device_proto_rawDescGZIP(), []int{2}
}

func (x *Spec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Spec) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Spec) GetOS() string {
	if x != nil {
		return x.OS
	}
	return ""
}

func (x *Spec) GetOsVersion() string {
	if x != nil {
		return x.OsVersion
	}
	return ""
}

func (x *Spec) GetAppVersion() string {
	if x != nil {
		return x.AppVersion
	}
	return ""
}

type BatteryLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time          string `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	BatteryLevel  int64  `protobuf:"varint,2,opt,name=batteryLevel,proto3" json:"batteryLevel,omitempty"`
	BatteryStatus string `protobuf:"bytes,3,opt,name=batteryStatus,proto3" json:"batteryStatus,omitempty"`
}

func (x *BatteryLevel) Reset() {
	*x = BatteryLevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_unit_device_device_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatteryLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatteryLevel) ProtoMessage() {}

func (x *BatteryLevel) ProtoReflect() protoreflect.Message {
	mi := &file_pb_unit_device_device_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatteryLevel.ProtoReflect.Descriptor instead.
func (*BatteryLevel) Descriptor() ([]byte, []int) {
	return file_pb_unit_device_device_proto_rawDescGZIP(), []int{3}
}

func (x *BatteryLevel) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *BatteryLevel) GetBatteryLevel() int64 {
	if x != nil {
		return x.BatteryLevel
	}
	return 0
}

func (x *BatteryLevel) GetBatteryStatus() string {
	if x != nil {
		return x.BatteryStatus
	}
	return ""
}

var File_pb_unit_device_device_proto protoreflect.FileDescriptor

var file_pb_unit_device_device_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x62, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70,
	0x62, 0x2e, 0x75, 0x6e, 0x69, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x98, 0x01,
	0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x75, 0x6e, 0x69, 0x74, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e,
	0x75, 0x6e, 0x69, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x40, 0x0a, 0x0c, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70,
	0x62, 0x2e, 0x75, 0x6e, 0x69, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x0c, 0x62, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7c,
	0x0a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x4f, 0x53, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x4f, 0x53, 0x12, 0x1c,
	0x0a, 0x09, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a,
	0x61, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x61, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x6c, 0x0a, 0x0c,
	0x42, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x61, 0x74,
	0x74, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x79, 0x65, 0x6f, 0x6c, 0x2d, 0x69,
	0x2f, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x2d, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x2d, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x2f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_unit_device_device_proto_rawDescOnce sync.Once
	file_pb_unit_device_device_proto_rawDescData = file_pb_unit_device_device_proto_rawDesc
)

func file_pb_unit_device_device_proto_rawDescGZIP() []byte {
	file_pb_unit_device_device_proto_rawDescOnce.Do(func() {
		file_pb_unit_device_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_unit_device_device_proto_rawDescData)
	})
	return file_pb_unit_device_device_proto_rawDescData
}

var file_pb_unit_device_device_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_unit_device_device_proto_goTypes = []interface{}{
	(*Device)(nil),       // 0: pb.unit.device.Device
	(*ID)(nil),           // 1: pb.unit.device.ID
	(*Spec)(nil),         // 2: pb.unit.device.Spec
	(*BatteryLevel)(nil), // 3: pb.unit.device.BatteryLevel
}
var file_pb_unit_device_device_proto_depIdxs = []int32{
	1, // 0: pb.unit.device.Device.id:type_name -> pb.unit.device.ID
	2, // 1: pb.unit.device.Device.spec:type_name -> pb.unit.device.Spec
	3, // 2: pb.unit.device.Device.batteryLevel:type_name -> pb.unit.device.BatteryLevel
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_unit_device_device_proto_init() }
func file_pb_unit_device_device_proto_init() {
	if File_pb_unit_device_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_unit_device_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_pb_unit_device_device_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
		file_pb_unit_device_device_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spec); i {
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
		file_pb_unit_device_device_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatteryLevel); i {
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
			RawDescriptor: file_pb_unit_device_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_unit_device_device_proto_goTypes,
		DependencyIndexes: file_pb_unit_device_device_proto_depIdxs,
		MessageInfos:      file_pb_unit_device_device_proto_msgTypes,
	}.Build()
	File_pb_unit_device_device_proto = out.File
	file_pb_unit_device_device_proto_rawDesc = nil
	file_pb_unit_device_device_proto_goTypes = nil
	file_pb_unit_device_device_proto_depIdxs = nil
}
