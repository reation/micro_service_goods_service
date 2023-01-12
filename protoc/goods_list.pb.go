// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: goods_list.proto

package protoc

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

type GoodsListInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title        string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Picture      string  `protobuf:"bytes,3,opt,name=picture,proto3" json:"picture,omitempty"`
	Cover        float32 `protobuf:"fixed32,4,opt,name=cover,proto3" json:"cover,omitempty"`
	BusinessId   int64   `protobuf:"varint,5,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty"`
	BusinessName string  `protobuf:"bytes,6,opt,name=business_name,json=businessName,proto3" json:"business_name,omitempty"`
}

func (x *GoodsListInfo) Reset() {
	*x = GoodsListInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsListInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsListInfo) ProtoMessage() {}

func (x *GoodsListInfo) ProtoReflect() protoreflect.Message {
	mi := &file_goods_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsListInfo.ProtoReflect.Descriptor instead.
func (*GoodsListInfo) Descriptor() ([]byte, []int) {
	return file_goods_list_proto_rawDescGZIP(), []int{0}
}

func (x *GoodsListInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GoodsListInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GoodsListInfo) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *GoodsListInfo) GetCover() float32 {
	if x != nil {
		return x.Cover
	}
	return 0
}

func (x *GoodsListInfo) GetBusinessId() int64 {
	if x != nil {
		return x.BusinessId
	}
	return 0
}

func (x *GoodsListInfo) GetBusinessName() string {
	if x != nil {
		return x.BusinessName
	}
	return ""
}

type GoodsListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	States   int64            `protobuf:"varint,1,opt,name=states,proto3" json:"states,omitempty"`
	GoodList []*GoodsListInfo `protobuf:"bytes,2,rep,name=good_list,json=goodList,proto3" json:"good_list,omitempty"`
}

func (x *GoodsListResponse) Reset() {
	*x = GoodsListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsListResponse) ProtoMessage() {}

func (x *GoodsListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goods_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsListResponse.ProtoReflect.Descriptor instead.
func (*GoodsListResponse) Descriptor() ([]byte, []int) {
	return file_goods_list_proto_rawDescGZIP(), []int{1}
}

func (x *GoodsListResponse) GetStates() int64 {
	if x != nil {
		return x.States
	}
	return 0
}

func (x *GoodsListResponse) GetGoodList() []*GoodsListInfo {
	if x != nil {
		return x.GoodList
	}
	return nil
}

type GoodsListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  int64 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Page  int64 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit int64 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GoodsListRequest) Reset() {
	*x = GoodsListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsListRequest) ProtoMessage() {}

func (x *GoodsListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goods_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsListRequest.ProtoReflect.Descriptor instead.
func (*GoodsListRequest) Descriptor() ([]byte, []int) {
	return file_goods_list_proto_rawDescGZIP(), []int{2}
}

func (x *GoodsListRequest) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *GoodsListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GoodsListRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

var File_goods_list_proto protoreflect.FileDescriptor

var file_goods_list_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xab,
	0x01, 0x0a, 0x0d, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x63, 0x0a, 0x11,
	0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x09, 0x67, 0x6f, 0x6f,
	0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x50, 0x0a, 0x10, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x32, 0x55, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x48, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f,
	0x6f, 0x64, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goods_list_proto_rawDescOnce sync.Once
	file_goods_list_proto_rawDescData = file_goods_list_proto_rawDesc
)

func file_goods_list_proto_rawDescGZIP() []byte {
	file_goods_list_proto_rawDescOnce.Do(func() {
		file_goods_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_goods_list_proto_rawDescData)
	})
	return file_goods_list_proto_rawDescData
}

var file_goods_list_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_goods_list_proto_goTypes = []interface{}{
	(*GoodsListInfo)(nil),     // 0: goods_list.goodsListInfo
	(*GoodsListResponse)(nil), // 1: goods_list.goodsListResponse
	(*GoodsListRequest)(nil),  // 2: goods_list.goodsListRequest
}
var file_goods_list_proto_depIdxs = []int32{
	0, // 0: goods_list.goodsListResponse.good_list:type_name -> goods_list.goodsListInfo
	2, // 1: goods_list.GoodsList.GoodsList:input_type -> goods_list.goodsListRequest
	1, // 2: goods_list.GoodsList.GoodsList:output_type -> goods_list.goodsListResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_goods_list_proto_init() }
func file_goods_list_proto_init() {
	if File_goods_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goods_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsListInfo); i {
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
		file_goods_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsListResponse); i {
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
		file_goods_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsListRequest); i {
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
			RawDescriptor: file_goods_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goods_list_proto_goTypes,
		DependencyIndexes: file_goods_list_proto_depIdxs,
		MessageInfos:      file_goods_list_proto_msgTypes,
	}.Build()
	File_goods_list_proto = out.File
	file_goods_list_proto_rawDesc = nil
	file_goods_list_proto_goTypes = nil
	file_goods_list_proto_depIdxs = nil
}