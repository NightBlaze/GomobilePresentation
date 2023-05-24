// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: feedwidget_viewmodels.proto

package protos

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

type FeedItemViewModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Caption  string `protobuf:"bytes,2,opt,name=caption,proto3" json:"caption,omitempty"`
	Title    string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	ImageUrl string `protobuf:"bytes,4,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *FeedItemViewModel) Reset() {
	*x = FeedItemViewModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedwidget_viewmodels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedItemViewModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedItemViewModel) ProtoMessage() {}

func (x *FeedItemViewModel) ProtoReflect() protoreflect.Message {
	mi := &file_feedwidget_viewmodels_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedItemViewModel.ProtoReflect.Descriptor instead.
func (*FeedItemViewModel) Descriptor() ([]byte, []int) {
	return file_feedwidget_viewmodels_proto_rawDescGZIP(), []int{0}
}

func (x *FeedItemViewModel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FeedItemViewModel) GetCaption() string {
	if x != nil {
		return x.Caption
	}
	return ""
}

func (x *FeedItemViewModel) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *FeedItemViewModel) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

type FeedItemLocalizableViewModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Caption string `protobuf:"bytes,2,opt,name=caption,proto3" json:"caption,omitempty"`
}

func (x *FeedItemLocalizableViewModel) Reset() {
	*x = FeedItemLocalizableViewModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedwidget_viewmodels_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedItemLocalizableViewModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedItemLocalizableViewModel) ProtoMessage() {}

func (x *FeedItemLocalizableViewModel) ProtoReflect() protoreflect.Message {
	mi := &file_feedwidget_viewmodels_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedItemLocalizableViewModel.ProtoReflect.Descriptor instead.
func (*FeedItemLocalizableViewModel) Descriptor() ([]byte, []int) {
	return file_feedwidget_viewmodels_proto_rawDescGZIP(), []int{1}
}

func (x *FeedItemLocalizableViewModel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FeedItemLocalizableViewModel) GetCaption() string {
	if x != nil {
		return x.Caption
	}
	return ""
}

type FeedLocalizationDidChangeViewModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeedItems []*FeedItemLocalizableViewModel `protobuf:"bytes,3,rep,name=feed_items,json=feedItems,proto3" json:"feed_items,omitempty"`
}

func (x *FeedLocalizationDidChangeViewModel) Reset() {
	*x = FeedLocalizationDidChangeViewModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedwidget_viewmodels_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedLocalizationDidChangeViewModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedLocalizationDidChangeViewModel) ProtoMessage() {}

func (x *FeedLocalizationDidChangeViewModel) ProtoReflect() protoreflect.Message {
	mi := &file_feedwidget_viewmodels_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedLocalizationDidChangeViewModel.ProtoReflect.Descriptor instead.
func (*FeedLocalizationDidChangeViewModel) Descriptor() ([]byte, []int) {
	return file_feedwidget_viewmodels_proto_rawDescGZIP(), []int{2}
}

func (x *FeedLocalizationDidChangeViewModel) GetFeedItems() []*FeedItemLocalizableViewModel {
	if x != nil {
		return x.FeedItems
	}
	return nil
}

var File_feedwidget_viewmodels_proto protoreflect.FileDescriptor

var file_feedwidget_viewmodels_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x66, 0x65, 0x65, 0x64, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x76, 0x69, 0x65,
	0x77, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a,
	0x11, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x56, 0x69, 0x65, 0x77, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x22,
	0x48, 0x0a, 0x1c, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x69, 0x7a, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x69, 0x65, 0x77, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x62, 0x0a, 0x22, 0x46, 0x65, 0x65,
	0x64, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x64,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x56, 0x69, 0x65, 0x77, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x3c, 0x0a, 0x0a, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x69, 0x65, 0x77, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x52, 0x09, 0x66, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x09, 0x5a,
	0x07, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feedwidget_viewmodels_proto_rawDescOnce sync.Once
	file_feedwidget_viewmodels_proto_rawDescData = file_feedwidget_viewmodels_proto_rawDesc
)

func file_feedwidget_viewmodels_proto_rawDescGZIP() []byte {
	file_feedwidget_viewmodels_proto_rawDescOnce.Do(func() {
		file_feedwidget_viewmodels_proto_rawDescData = protoimpl.X.CompressGZIP(file_feedwidget_viewmodels_proto_rawDescData)
	})
	return file_feedwidget_viewmodels_proto_rawDescData
}

var file_feedwidget_viewmodels_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_feedwidget_viewmodels_proto_goTypes = []interface{}{
	(*FeedItemViewModel)(nil),                  // 0: FeedItemViewModel
	(*FeedItemLocalizableViewModel)(nil),       // 1: FeedItemLocalizableViewModel
	(*FeedLocalizationDidChangeViewModel)(nil), // 2: FeedLocalizationDidChangeViewModel
}
var file_feedwidget_viewmodels_proto_depIdxs = []int32{
	1, // 0: FeedLocalizationDidChangeViewModel.feed_items:type_name -> FeedItemLocalizableViewModel
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_feedwidget_viewmodels_proto_init() }
func file_feedwidget_viewmodels_proto_init() {
	if File_feedwidget_viewmodels_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_feedwidget_viewmodels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedItemViewModel); i {
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
		file_feedwidget_viewmodels_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedItemLocalizableViewModel); i {
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
		file_feedwidget_viewmodels_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedLocalizationDidChangeViewModel); i {
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
			RawDescriptor: file_feedwidget_viewmodels_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_feedwidget_viewmodels_proto_goTypes,
		DependencyIndexes: file_feedwidget_viewmodels_proto_depIdxs,
		MessageInfos:      file_feedwidget_viewmodels_proto_msgTypes,
	}.Build()
	File_feedwidget_viewmodels_proto = out.File
	file_feedwidget_viewmodels_proto_rawDesc = nil
	file_feedwidget_viewmodels_proto_goTypes = nil
	file_feedwidget_viewmodels_proto_depIdxs = nil
}
