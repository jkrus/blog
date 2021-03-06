// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.0
// source: models/grpc/note.proto

package proto

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

type Note struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"id" gorm:"primaryKey"`                  // @gotags: json:"id" gorm:"primaryKey"
	Header    string `protobuf:"bytes,2,opt,name=Header,proto3" json:"header" gorm:"header"`               // @gotags: json:"header" gorm:"header"
	Body      string `protobuf:"bytes,3,opt,name=Body,proto3" json:"body" gorm:"body"`                     // @gotags: json:"body" gorm:"body"
	Tags      string `protobuf:"bytes,4,opt,name=Tags,proto3" json:"tags" gorm:"tags"`                     // @gotags: json:"tags" gorm:"tags"
	CreatedAt string `protobuf:"bytes,5,opt,name=CreatedAt,proto3" json:"createdAt" gorm:"autoCreateTime"` // @gotags: json:"createdAt" gorm:"autoCreateTime"
}

func (x *Note) Reset() {
	*x = Note{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_grpc_note_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Note) ProtoMessage() {}

func (x *Note) ProtoReflect() protoreflect.Message {
	mi := &file_models_grpc_note_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Note.ProtoReflect.Descriptor instead.
func (*Note) Descriptor() ([]byte, []int) {
	return file_models_grpc_note_proto_rawDescGZIP(), []int{0}
}

func (x *Note) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Note) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

func (x *Note) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Note) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *Note) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type NoteDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header string   `protobuf:"bytes,2,opt,name=Header,proto3" json:"header"` // @gotags: json:"header"
	Body   string   `protobuf:"bytes,3,opt,name=Body,proto3" json:"body"`     // @gotags: json:"body"
	Tags   []string `protobuf:"bytes,4,rep,name=Tags,proto3" json:"tags"`     // @gotags: json:"tags"
}

func (x *NoteDTO) Reset() {
	*x = NoteDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_grpc_note_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteDTO) ProtoMessage() {}

func (x *NoteDTO) ProtoReflect() protoreflect.Message {
	mi := &file_models_grpc_note_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteDTO.ProtoReflect.Descriptor instead.
func (*NoteDTO) Descriptor() ([]byte, []int) {
	return file_models_grpc_note_proto_rawDescGZIP(), []int{1}
}

func (x *NoteDTO) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

func (x *NoteDTO) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *NoteDTO) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type Notes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32   `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Notes  []*Note `protobuf:"bytes,2,rep,name=Notes,proto3" json:"Notes,omitempty"`
}

func (x *Notes) Reset() {
	*x = Notes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_grpc_note_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notes) ProtoMessage() {}

func (x *Notes) ProtoReflect() protoreflect.Message {
	mi := &file_models_grpc_note_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notes.ProtoReflect.Descriptor instead.
func (*Notes) Descriptor() ([]byte, []int) {
	return file_models_grpc_note_proto_rawDescGZIP(), []int{2}
}

func (x *Notes) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Notes) GetNotes() []*Note {
	if x != nil {
		return x.Notes
	}
	return nil
}

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_grpc_note_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_models_grpc_note_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_models_grpc_note_proto_rawDescGZIP(), []int{3}
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_grpc_note_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_grpc_note_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_models_grpc_note_proto_rawDescGZIP(), []int{4}
}

func (x *EmptyResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_models_grpc_note_proto protoreflect.FileDescriptor

var file_models_grpc_note_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6e, 0x6f,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x04, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x61, 0x67, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x61, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x49, 0x0a, 0x07, 0x4e, 0x6f,
	0x74, 0x65, 0x44, 0x54, 0x4f, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x42, 0x6f, 0x64,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x54, 0x61, 0x67, 0x73, 0x22, 0x47, 0x0a, 0x05, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x05, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x0e,
	0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x27,
	0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x18, 0x5a, 0x16, 0x62, 0x6c, 0x6f, 0x67, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_grpc_note_proto_rawDescOnce sync.Once
	file_models_grpc_note_proto_rawDescData = file_models_grpc_note_proto_rawDesc
)

func file_models_grpc_note_proto_rawDescGZIP() []byte {
	file_models_grpc_note_proto_rawDescOnce.Do(func() {
		file_models_grpc_note_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_grpc_note_proto_rawDescData)
	})
	return file_models_grpc_note_proto_rawDescData
}

var file_models_grpc_note_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_models_grpc_note_proto_goTypes = []interface{}{
	(*Note)(nil),          // 0: blog.proto.Note
	(*NoteDTO)(nil),       // 1: blog.proto.NoteDTO
	(*Notes)(nil),         // 2: blog.proto.Notes
	(*EmptyRequest)(nil),  // 3: blog.proto.EmptyRequest
	(*EmptyResponse)(nil), // 4: blog.proto.EmptyResponse
}
var file_models_grpc_note_proto_depIdxs = []int32{
	0, // 0: blog.proto.Notes.Notes:type_name -> blog.proto.Note
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_grpc_note_proto_init() }
func file_models_grpc_note_proto_init() {
	if File_models_grpc_note_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_grpc_note_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Note); i {
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
		file_models_grpc_note_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteDTO); i {
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
		file_models_grpc_note_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notes); i {
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
		file_models_grpc_note_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
		file_models_grpc_note_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
			RawDescriptor: file_models_grpc_note_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_grpc_note_proto_goTypes,
		DependencyIndexes: file_models_grpc_note_proto_depIdxs,
		MessageInfos:      file_models_grpc_note_proto_msgTypes,
	}.Build()
	File_models_grpc_note_proto = out.File
	file_models_grpc_note_proto_rawDesc = nil
	file_models_grpc_note_proto_goTypes = nil
	file_models_grpc_note_proto_depIdxs = nil
}
