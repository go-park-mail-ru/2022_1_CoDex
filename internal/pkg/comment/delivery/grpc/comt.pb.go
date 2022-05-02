// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: comt.proto

package grpc

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

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Imgsrc   string `protobuf:"bytes,1,opt,name=imgsrc,proto3" json:"imgsrc,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	UserId   string `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
	Rating   string `protobuf:"bytes,4,opt,name=rating,proto3" json:"rating,omitempty"`
	Date     string `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	Content  string `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Type     string `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_comt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_comt_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetImgsrc() string {
	if x != nil {
		return x.Imgsrc
	}
	return ""
}

func (x *Comment) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Comment) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Comment) GetRating() string {
	if x != nil {
		return x.Rating
	}
	return ""
}

func (x *Comment) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieId     uint64 `protobuf:"varint,1,opt,name=movieId,proto3" json:"movieId,omitempty"`
	UserId      uint64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Content     string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	CommentType int32  `protobuf:"varint,4,opt,name=commentType,proto3" json:"commentType,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_comt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_comt_proto_rawDescGZIP(), []int{1}
}

func (x *Data) GetMovieId() uint64 {
	if x != nil {
		return x.MovieId
	}
	return 0
}

func (x *Data) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Data) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Data) GetCommentType() int32 {
	if x != nil {
		return x.CommentType
	}
	return 0
}

var File_comt_proto protoreflect.FileDescriptor

var file_comt_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x43, 0x6f,
	0x6d, 0x6d, 0x22, 0xaf, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x69, 0x6d, 0x67, 0x73, 0x72, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x69, 0x6d, 0x67, 0x73, 0x72, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x74, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x32, 0x32, 0x0a, 0x06, 0x50, 0x6f,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x0a, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a,
	0x0d, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x27,
	0x5a, 0x25, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comt_proto_rawDescOnce sync.Once
	file_comt_proto_rawDescData = file_comt_proto_rawDesc
)

func file_comt_proto_rawDescGZIP() []byte {
	file_comt_proto_rawDescOnce.Do(func() {
		file_comt_proto_rawDescData = protoimpl.X.CompressGZIP(file_comt_proto_rawDescData)
	})
	return file_comt_proto_rawDescData
}

var file_comt_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_comt_proto_goTypes = []interface{}{
	(*Comment)(nil), // 0: Comm.Comment
	(*Data)(nil),    // 1: Comm.Data
}
var file_comt_proto_depIdxs = []int32{
	1, // 0: Comm.Poster.PostComment:input_type -> Comm.Data
	0, // 1: Comm.Poster.PostComment:output_type -> Comm.Comment
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_comt_proto_init() }
func file_comt_proto_init() {
	if File_comt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_comt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
			RawDescriptor: file_comt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comt_proto_goTypes,
		DependencyIndexes: file_comt_proto_depIdxs,
		MessageInfos:      file_comt_proto_msgTypes,
	}.Build()
	File_comt_proto = out.File
	file_comt_proto_rawDesc = nil
	file_comt_proto_goTypes = nil
	file_comt_proto_depIdxs = nil
}
