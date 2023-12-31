// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: post.proto

package social_media_proto

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat float64 `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng float64 `protobuf:"fixed64,2,opt,name=lng,proto3" json:"lng,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Location) GetLng() float64 {
	if x != nil {
		return x.Lng
	}
	return 0
}

type PostDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid string    `protobuf:"bytes,1,opt,name=userUuid,proto3" json:"userUuid,omitempty"`
	Caption  string    `protobuf:"bytes,2,opt,name=caption,proto3" json:"caption,omitempty"`
	ImageUrl string    `protobuf:"bytes,3,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	Location *Location `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *PostDetail) Reset() {
	*x = PostDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostDetail) ProtoMessage() {}

func (x *PostDetail) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostDetail.ProtoReflect.Descriptor instead.
func (*PostDetail) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{1}
}

func (x *PostDetail) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *PostDetail) GetCaption() string {
	if x != nil {
		return x.Caption
	}
	return ""
}

func (x *PostDetail) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *PostDetail) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type PostList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post []*PostDetail `protobuf:"bytes,1,rep,name=post,proto3" json:"post,omitempty"`
}

func (x *PostList) Reset() {
	*x = PostList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostList) ProtoMessage() {}

func (x *PostList) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostList.ProtoReflect.Descriptor instead.
func (*PostList) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{2}
}

func (x *PostList) GetPost() []*PostDetail {
	if x != nil {
		return x.Post
	}
	return nil
}

var File_post_proto protoreflect.FileDescriptor

var file_post_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x6f,
	0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03,
	0x6c, 0x6e, 0x67, 0x22, 0x92, 0x01, 0x0a, 0x0a, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x32, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x38, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x04, 0x70, 0x6f,
	0x73, 0x74, 0x32, 0xbf, 0x01, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a,
	0x18, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0f, 0x52,
	0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x55, 0x75,
	0x69, 0x64, 0x1a, 0x16, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x50, 0x4e, 0x59, 0x77, 0x69, 0x73, 0x65, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2d, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_post_proto_rawDescOnce sync.Once
	file_post_proto_rawDescData = file_post_proto_rawDesc
)

func file_post_proto_rawDescGZIP() []byte {
	file_post_proto_rawDescOnce.Do(func() {
		file_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_post_proto_rawDescData)
	})
	return file_post_proto_rawDescData
}

var file_post_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_post_proto_goTypes = []interface{}{
	(*Location)(nil),    // 0: social_media.Location
	(*PostDetail)(nil),  // 1: social_media.PostDetail
	(*PostList)(nil),    // 2: social_media.PostList
	(*Uuid)(nil),        // 3: social_media.Uuid
	(*empty.Empty)(nil), // 4: google.protobuf.Empty
}
var file_post_proto_depIdxs = []int32{
	0, // 0: social_media.PostDetail.location:type_name -> social_media.Location
	1, // 1: social_media.PostList.post:type_name -> social_media.PostDetail
	1, // 2: social_media.Post.Create:input_type -> social_media.PostDetail
	3, // 3: social_media.Post.ReadAllByUserId:input_type -> social_media.Uuid
	3, // 4: social_media.Post.Delete:input_type -> social_media.Uuid
	1, // 5: social_media.Post.Create:output_type -> social_media.PostDetail
	2, // 6: social_media.Post.ReadAllByUserId:output_type -> social_media.PostList
	4, // 7: social_media.Post.Delete:output_type -> google.protobuf.Empty
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_post_proto_init() }
func file_post_proto_init() {
	if File_post_proto != nil {
		return
	}
	file_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostDetail); i {
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
		file_post_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostList); i {
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
			RawDescriptor: file_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_post_proto_goTypes,
		DependencyIndexes: file_post_proto_depIdxs,
		MessageInfos:      file_post_proto_msgTypes,
	}.Build()
	File_post_proto = out.File
	file_post_proto_rawDesc = nil
	file_post_proto_goTypes = nil
	file_post_proto_depIdxs = nil
}
