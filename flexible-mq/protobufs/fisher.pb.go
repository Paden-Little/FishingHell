// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: proto-files/fisher.proto

package protobufs

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

type StartFish struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Pool   string `protobuf:"bytes,2,opt,name=pool,proto3" json:"pool,omitempty"`
	PoleID int32  `protobuf:"varint,3,opt,name=poleID,proto3" json:"poleID,omitempty"`
	BaitID int32  `protobuf:"varint,4,opt,name=baitID,proto3" json:"baitID,omitempty"`
}

func (x *StartFish) Reset() {
	*x = StartFish{}
	mi := &file_proto_files_fisher_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartFish) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartFish) ProtoMessage() {}

func (x *StartFish) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_fisher_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartFish.ProtoReflect.Descriptor instead.
func (*StartFish) Descriptor() ([]byte, []int) {
	return file_proto_files_fisher_proto_rawDescGZIP(), []int{0}
}

func (x *StartFish) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *StartFish) GetPool() string {
	if x != nil {
		return x.Pool
	}
	return ""
}

func (x *StartFish) GetPoleID() int32 {
	if x != nil {
		return x.PoleID
	}
	return 0
}

func (x *StartFish) GetBaitID() int32 {
	if x != nil {
		return x.BaitID
	}
	return 0
}

type Fish struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Taxonomy   string  `protobuf:"bytes,3,opt,name=taxonomy,proto3" json:"taxonomy,omitempty"`
	Pool       string  `protobuf:"bytes,4,opt,name=pool,proto3" json:"pool,omitempty"`
	Weight     float32 `protobuf:"fixed32,5,opt,name=weight,proto3" json:"weight,omitempty"`
	TimeCaught string  `protobuf:"bytes,6,opt,name=timeCaught,proto3" json:"timeCaught,omitempty"`
	BaitUsed   int32   `protobuf:"varint,7,opt,name=baitUsed,proto3" json:"baitUsed,omitempty"`
}

func (x *Fish) Reset() {
	*x = Fish{}
	mi := &file_proto_files_fisher_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Fish) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fish) ProtoMessage() {}

func (x *Fish) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_fisher_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fish.ProtoReflect.Descriptor instead.
func (*Fish) Descriptor() ([]byte, []int) {
	return file_proto_files_fisher_proto_rawDescGZIP(), []int{1}
}

func (x *Fish) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Fish) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Fish) GetTaxonomy() string {
	if x != nil {
		return x.Taxonomy
	}
	return ""
}

func (x *Fish) GetPool() string {
	if x != nil {
		return x.Pool
	}
	return ""
}

func (x *Fish) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Fish) GetTimeCaught() string {
	if x != nil {
		return x.TimeCaught
	}
	return ""
}

func (x *Fish) GetBaitUsed() int32 {
	if x != nil {
		return x.BaitUsed
	}
	return 0
}

var File_proto_files_fisher_proto protoreflect.FileDescriptor

var file_proto_files_fisher_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x66, 0x69,
	0x73, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x73, 0x22, 0x67, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x46, 0x69,
	0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6f, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x70, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x61, 0x69, 0x74, 0x49, 0x44,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x61, 0x69, 0x74, 0x49, 0x44, 0x22, 0xae,
	0x01, 0x0a, 0x04, 0x46, 0x69, 0x73, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74,
	0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6f, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x61, 0x75, 0x67, 0x68,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x61, 0x75,
	0x67, 0x68, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x61, 0x69, 0x74, 0x55, 0x73, 0x65, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x62, 0x61, 0x69, 0x74, 0x55, 0x73, 0x65, 0x64, 0x42,
	0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_files_fisher_proto_rawDescOnce sync.Once
	file_proto_files_fisher_proto_rawDescData = file_proto_files_fisher_proto_rawDesc
)

func file_proto_files_fisher_proto_rawDescGZIP() []byte {
	file_proto_files_fisher_proto_rawDescOnce.Do(func() {
		file_proto_files_fisher_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_files_fisher_proto_rawDescData)
	})
	return file_proto_files_fisher_proto_rawDescData
}

var file_proto_files_fisher_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_files_fisher_proto_goTypes = []any{
	(*StartFish)(nil), // 0: protobufs.StartFish
	(*Fish)(nil),      // 1: protobufs.Fish
}
var file_proto_files_fisher_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_files_fisher_proto_init() }
func file_proto_files_fisher_proto_init() {
	if File_proto_files_fisher_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_files_fisher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_files_fisher_proto_goTypes,
		DependencyIndexes: file_proto_files_fisher_proto_depIdxs,
		MessageInfos:      file_proto_files_fisher_proto_msgTypes,
	}.Build()
	File_proto_files_fisher_proto = out.File
	file_proto_files_fisher_proto_rawDesc = nil
	file_proto_files_fisher_proto_goTypes = nil
	file_proto_files_fisher_proto_depIdxs = nil
}