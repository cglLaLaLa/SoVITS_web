// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: slice_pb/start_slice.proto

package slicepb

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

// 请求结构体
type SliceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InputPath   string  `protobuf:"bytes,1,opt,name=input_path,json=inputPath,proto3" json:"input_path,omitempty"`
	OutputRoot  string  `protobuf:"bytes,2,opt,name=output_root,json=outputRoot,proto3" json:"output_root,omitempty"`
	Threshold   float32 `protobuf:"fixed32,3,opt,name=threshold,proto3" json:"threshold,omitempty"`
	MinLength   int32   `protobuf:"varint,4,opt,name=min_length,json=minLength,proto3" json:"min_length,omitempty"`
	MinInterval int32   `protobuf:"varint,5,opt,name=min_interval,json=minInterval,proto3" json:"min_interval,omitempty"`
	HopSize     int32   `protobuf:"varint,6,opt,name=hop_size,json=hopSize,proto3" json:"hop_size,omitempty"`
	MaxSilKept  int32   `protobuf:"varint,7,opt,name=max_sil_kept,json=maxSilKept,proto3" json:"max_sil_kept,omitempty"`
	XMax        int32   `protobuf:"varint,8,opt,name=_max,json=Max,proto3" json:"_max,omitempty"`
	Alpha       float32 `protobuf:"fixed32,9,opt,name=alpha,proto3" json:"alpha,omitempty"`
	NParts      int32   `protobuf:"varint,10,opt,name=n_parts,json=nParts,proto3" json:"n_parts,omitempty"` // 分割部分数
}

func (x *SliceRequest) Reset() {
	*x = SliceRequest{}
	mi := &file_slice_pb_start_slice_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SliceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SliceRequest) ProtoMessage() {}

func (x *SliceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_slice_pb_start_slice_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SliceRequest.ProtoReflect.Descriptor instead.
func (*SliceRequest) Descriptor() ([]byte, []int) {
	return file_slice_pb_start_slice_proto_rawDescGZIP(), []int{0}
}

func (x *SliceRequest) GetInputPath() string {
	if x != nil {
		return x.InputPath
	}
	return ""
}

func (x *SliceRequest) GetOutputRoot() string {
	if x != nil {
		return x.OutputRoot
	}
	return ""
}

func (x *SliceRequest) GetThreshold() float32 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

func (x *SliceRequest) GetMinLength() int32 {
	if x != nil {
		return x.MinLength
	}
	return 0
}

func (x *SliceRequest) GetMinInterval() int32 {
	if x != nil {
		return x.MinInterval
	}
	return 0
}

func (x *SliceRequest) GetHopSize() int32 {
	if x != nil {
		return x.HopSize
	}
	return 0
}

func (x *SliceRequest) GetMaxSilKept() int32 {
	if x != nil {
		return x.MaxSilKept
	}
	return 0
}

func (x *SliceRequest) GetXMax() int32 {
	if x != nil {
		return x.XMax
	}
	return 0
}

func (x *SliceRequest) GetAlpha() float32 {
	if x != nil {
		return x.Alpha
	}
	return 0
}

func (x *SliceRequest) GetNParts() int32 {
	if x != nil {
		return x.NParts
	}
	return 0
}

// 响应结构体
type StartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string          `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` //成功还是失败
	Data   []*KeyValuePair `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`     //返回键值对格式的数据信息
}

func (x *StartResponse) Reset() {
	*x = StartResponse{}
	mi := &file_slice_pb_start_slice_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartResponse) ProtoMessage() {}

func (x *StartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_slice_pb_start_slice_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartResponse.ProtoReflect.Descriptor instead.
func (*StartResponse) Descriptor() ([]byte, []int) {
	return file_slice_pb_start_slice_proto_rawDescGZIP(), []int{1}
}

func (x *StartResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *StartResponse) GetData() []*KeyValuePair {
	if x != nil {
		return x.Data
	}
	return nil
}

type KeyValuePair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KeyValuePair) Reset() {
	*x = KeyValuePair{}
	mi := &file_slice_pb_start_slice_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KeyValuePair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValuePair) ProtoMessage() {}

func (x *KeyValuePair) ProtoReflect() protoreflect.Message {
	mi := &file_slice_pb_start_slice_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValuePair.ProtoReflect.Descriptor instead.
func (*KeyValuePair) Descriptor() ([]byte, []int) {
	return file_slice_pb_start_slice_proto_rawDescGZIP(), []int{2}
}

func (x *KeyValuePair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValuePair) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_slice_pb_start_slice_proto protoreflect.FileDescriptor

var file_slice_pb_start_slice_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x6c,
	0x69, 0x63, 0x65, 0x22, 0xad, 0x02, 0x0a, 0x0c, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x72, 0x6f,
	0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x69, 0x6e, 0x4c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x70, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x68, 0x6f, 0x70, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x20, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x69, 0x6c, 0x5f, 0x6b, 0x65, 0x70, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x53, 0x69, 0x6c, 0x4b, 0x65, 0x70,
	0x74, 0x12, 0x11, 0x0a, 0x04, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x4d, 0x61, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x5f,
	0x70, 0x61, 0x72, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x50, 0x61,
	0x72, 0x74, 0x73, 0x22, 0x50, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x27, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x6c, 0x69,
	0x63, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x36, 0x0a, 0x0c, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x50, 0x61, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x47, 0x0a,
	0x0c, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a,
	0x0a, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x6c,
	0x69, 0x63, 0x65, 0x2e, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x73, 0x6c, 0x69, 0x63,
	0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_slice_pb_start_slice_proto_rawDescOnce sync.Once
	file_slice_pb_start_slice_proto_rawDescData = file_slice_pb_start_slice_proto_rawDesc
)

func file_slice_pb_start_slice_proto_rawDescGZIP() []byte {
	file_slice_pb_start_slice_proto_rawDescOnce.Do(func() {
		file_slice_pb_start_slice_proto_rawDescData = protoimpl.X.CompressGZIP(file_slice_pb_start_slice_proto_rawDescData)
	})
	return file_slice_pb_start_slice_proto_rawDescData
}

var file_slice_pb_start_slice_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_slice_pb_start_slice_proto_goTypes = []any{
	(*SliceRequest)(nil),  // 0: slice.SliceRequest
	(*StartResponse)(nil), // 1: slice.StartResponse
	(*KeyValuePair)(nil),  // 2: slice.KeyValuePair
}
var file_slice_pb_start_slice_proto_depIdxs = []int32{
	2, // 0: slice.StartResponse.data:type_name -> slice.KeyValuePair
	0, // 1: slice.SliceService.StartSlice:input_type -> slice.SliceRequest
	1, // 2: slice.SliceService.StartSlice:output_type -> slice.StartResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_slice_pb_start_slice_proto_init() }
func file_slice_pb_start_slice_proto_init() {
	if File_slice_pb_start_slice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_slice_pb_start_slice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_slice_pb_start_slice_proto_goTypes,
		DependencyIndexes: file_slice_pb_start_slice_proto_depIdxs,
		MessageInfos:      file_slice_pb_start_slice_proto_msgTypes,
	}.Build()
	File_slice_pb_start_slice_proto = out.File
	file_slice_pb_start_slice_proto_rawDesc = nil
	file_slice_pb_start_slice_proto_goTypes = nil
	file_slice_pb_start_slice_proto_depIdxs = nil
}
