// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: filetransfer/filetransfer.proto

package filetransfer

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

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Content []byte `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filetransfer_filetransfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_filetransfer_filetransfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_filetransfer_filetransfer_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type FileName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *FileName) Reset() {
	*x = FileName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filetransfer_filetransfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileName) ProtoMessage() {}

func (x *FileName) ProtoReflect() protoreflect.Message {
	mi := &file_filetransfer_filetransfer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileName.ProtoReflect.Descriptor instead.
func (*FileName) Descriptor() ([]byte, []int) {
	return file_filetransfer_filetransfer_proto_rawDescGZIP(), []int{1}
}

func (x *FileName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FileContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *FileContent) Reset() {
	*x = FileContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filetransfer_filetransfer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileContent) ProtoMessage() {}

func (x *FileContent) ProtoReflect() protoreflect.Message {
	mi := &file_filetransfer_filetransfer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileContent.ProtoReflect.Descriptor instead.
func (*FileContent) Descriptor() ([]byte, []int) {
	return file_filetransfer_filetransfer_proto_rawDescGZIP(), []int{2}
}

func (x *FileContent) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type FileList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files map[string]uint64 `protobuf:"bytes,1,rep,name=Files,proto3" json:"Files,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *FileList) Reset() {
	*x = FileList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filetransfer_filetransfer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileList) ProtoMessage() {}

func (x *FileList) ProtoReflect() protoreflect.Message {
	mi := &file_filetransfer_filetransfer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileList.ProtoReflect.Descriptor instead.
func (*FileList) Descriptor() ([]byte, []int) {
	return file_filetransfer_filetransfer_proto_rawDescGZIP(), []int{3}
}

func (x *FileList) GetFiles() map[string]uint64 {
	if x != nil {
		return x.Files
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filetransfer_filetransfer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_filetransfer_filetransfer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_filetransfer_filetransfer_proto_rawDescGZIP(), []int{4}
}

var File_filetransfer_filetransfer_proto protoreflect.FileDescriptor

var file_filetransfer_filetransfer_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x66,
	0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x22,
	0x34, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x1e, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x21, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x7d, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x05, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x1a, 0x38, 0x0a,
	0x0a, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x32, 0xf7, 0x01, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x12, 0x34, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x1a, 0x13,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12, 0x3e, 0x0a, 0x05, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x12, 0x16, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x19, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x37, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x16, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x13, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x38, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x13, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x16, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x68, 0x73, 0x75, 0x6e, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x66, 0x69, 0x6c, 0x65, 0x2d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_filetransfer_filetransfer_proto_rawDescOnce sync.Once
	file_filetransfer_filetransfer_proto_rawDescData = file_filetransfer_filetransfer_proto_rawDesc
)

func file_filetransfer_filetransfer_proto_rawDescGZIP() []byte {
	file_filetransfer_filetransfer_proto_rawDescOnce.Do(func() {
		file_filetransfer_filetransfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_filetransfer_filetransfer_proto_rawDescData)
	})
	return file_filetransfer_filetransfer_proto_rawDescData
}

var file_filetransfer_filetransfer_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_filetransfer_filetransfer_proto_goTypes = []interface{}{
	(*File)(nil),        // 0: filetransfer.File
	(*FileName)(nil),    // 1: filetransfer.FileName
	(*FileContent)(nil), // 2: filetransfer.FileContent
	(*FileList)(nil),    // 3: filetransfer.FileList
	(*Empty)(nil),       // 4: filetransfer.Empty
	nil,                 // 5: filetransfer.FileList.FilesEntry
}
var file_filetransfer_filetransfer_proto_depIdxs = []int32{
	5, // 0: filetransfer.FileList.Files:type_name -> filetransfer.FileList.FilesEntry
	0, // 1: filetransfer.FileTransfer.Store:input_type -> filetransfer.File
	1, // 2: filetransfer.FileTransfer.Fetch:input_type -> filetransfer.FileName
	1, // 3: filetransfer.FileTransfer.Delete:input_type -> filetransfer.FileName
	4, // 4: filetransfer.FileTransfer.ListAll:input_type -> filetransfer.Empty
	4, // 5: filetransfer.FileTransfer.Store:output_type -> filetransfer.Empty
	2, // 6: filetransfer.FileTransfer.Fetch:output_type -> filetransfer.FileContent
	4, // 7: filetransfer.FileTransfer.Delete:output_type -> filetransfer.Empty
	3, // 8: filetransfer.FileTransfer.ListAll:output_type -> filetransfer.FileList
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_filetransfer_filetransfer_proto_init() }
func file_filetransfer_filetransfer_proto_init() {
	if File_filetransfer_filetransfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_filetransfer_filetransfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_filetransfer_filetransfer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileName); i {
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
		file_filetransfer_filetransfer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileContent); i {
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
		file_filetransfer_filetransfer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileList); i {
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
		file_filetransfer_filetransfer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_filetransfer_filetransfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_filetransfer_filetransfer_proto_goTypes,
		DependencyIndexes: file_filetransfer_filetransfer_proto_depIdxs,
		MessageInfos:      file_filetransfer_filetransfer_proto_msgTypes,
	}.Build()
	File_filetransfer_filetransfer_proto = out.File
	file_filetransfer_filetransfer_proto_rawDesc = nil
	file_filetransfer_filetransfer_proto_goTypes = nil
	file_filetransfer_filetransfer_proto_depIdxs = nil
}
