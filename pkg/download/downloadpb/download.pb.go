// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: pkg/download/downloadpb/download.proto

package downloadpb

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

type FileDownloadInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FileDownloadInfo) Reset() {
	*x = FileDownloadInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_download_downloadpb_download_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileDownloadInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileDownloadInfo) ProtoMessage() {}

func (x *FileDownloadInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_download_downloadpb_download_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileDownloadInfo.ProtoReflect.Descriptor instead.
func (*FileDownloadInfo) Descriptor() ([]byte, []int) {
	return file_pkg_download_downloadpb_download_proto_rawDescGZIP(), []int{0}
}

func (x *FileDownloadInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FileDownloadBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bytes []byte `protobuf:"bytes,4,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (x *FileDownloadBody) Reset() {
	*x = FileDownloadBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_download_downloadpb_download_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileDownloadBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileDownloadBody) ProtoMessage() {}

func (x *FileDownloadBody) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_download_downloadpb_download_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileDownloadBody.ProtoReflect.Descriptor instead.
func (*FileDownloadBody) Descriptor() ([]byte, []int) {
	return file_pkg_download_downloadpb_download_proto_rawDescGZIP(), []int{1}
}

func (x *FileDownloadBody) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

type FileDownloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*FileDownloadResponse_Info
	//	*FileDownloadResponse_Body
	Data isFileDownloadResponse_Data `protobuf_oneof:"data"`
}

func (x *FileDownloadResponse) Reset() {
	*x = FileDownloadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_download_downloadpb_download_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileDownloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileDownloadResponse) ProtoMessage() {}

func (x *FileDownloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_download_downloadpb_download_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileDownloadResponse.ProtoReflect.Descriptor instead.
func (*FileDownloadResponse) Descriptor() ([]byte, []int) {
	return file_pkg_download_downloadpb_download_proto_rawDescGZIP(), []int{2}
}

func (m *FileDownloadResponse) GetData() isFileDownloadResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *FileDownloadResponse) GetInfo() *FileDownloadInfo {
	if x, ok := x.GetData().(*FileDownloadResponse_Info); ok {
		return x.Info
	}
	return nil
}

func (x *FileDownloadResponse) GetBody() *FileDownloadBody {
	if x, ok := x.GetData().(*FileDownloadResponse_Body); ok {
		return x.Body
	}
	return nil
}

type isFileDownloadResponse_Data interface {
	isFileDownloadResponse_Data()
}

type FileDownloadResponse_Info struct {
	Info *FileDownloadInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type FileDownloadResponse_Body struct {
	Body *FileDownloadBody `protobuf:"bytes,2,opt,name=body,proto3,oneof"`
}

func (*FileDownloadResponse_Info) isFileDownloadResponse_Data() {}

func (*FileDownloadResponse_Body) isFileDownloadResponse_Data() {}

type FileDownloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FileDownloadRequest) Reset() {
	*x = FileDownloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_download_downloadpb_download_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileDownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileDownloadRequest) ProtoMessage() {}

func (x *FileDownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_download_downloadpb_download_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileDownloadRequest.ProtoReflect.Descriptor instead.
func (*FileDownloadRequest) Descriptor() ([]byte, []int) {
	return file_pkg_download_downloadpb_download_proto_rawDescGZIP(), []int{3}
}

func (x *FileDownloadRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FileDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FileDeleteRequest) Reset() {
	*x = FileDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_download_downloadpb_download_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileDeleteRequest) ProtoMessage() {}

func (x *FileDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_download_downloadpb_download_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileDeleteRequest.ProtoReflect.Descriptor instead.
func (*FileDeleteRequest) Descriptor() ([]byte, []int) {
	return file_pkg_download_downloadpb_download_proto_rawDescGZIP(), []int{4}
}

func (x *FileDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FileDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok  bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *FileDeleteResponse) Reset() {
	*x = FileDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_download_downloadpb_download_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileDeleteResponse) ProtoMessage() {}

func (x *FileDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_download_downloadpb_download_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileDeleteResponse.ProtoReflect.Descriptor instead.
func (*FileDeleteResponse) Descriptor() ([]byte, []int) {
	return file_pkg_download_downloadpb_download_proto_rawDescGZIP(), []int{5}
}

func (x *FileDeleteResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *FileDeleteResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_pkg_download_downloadpb_download_proto protoreflect.FileDescriptor

var file_pkg_download_downloadpb_download_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x64,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x70, 0x62, 0x2f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x22, 0x26, 0x0a, 0x10, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x10, 0x46, 0x69,
	0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x14, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12,
	0x30, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6f, 0x64, 0x79, 0x48, 0x00, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x25, 0x0a, 0x13, 0x46, 0x69, 0x6c,
	0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x23, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x12, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xb3, 0x01,
	0x0a, 0x13, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x49, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_download_downloadpb_download_proto_rawDescOnce sync.Once
	file_pkg_download_downloadpb_download_proto_rawDescData = file_pkg_download_downloadpb_download_proto_rawDesc
)

func file_pkg_download_downloadpb_download_proto_rawDescGZIP() []byte {
	file_pkg_download_downloadpb_download_proto_rawDescOnce.Do(func() {
		file_pkg_download_downloadpb_download_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_download_downloadpb_download_proto_rawDescData)
	})
	return file_pkg_download_downloadpb_download_proto_rawDescData
}

var file_pkg_download_downloadpb_download_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_download_downloadpb_download_proto_goTypes = []interface{}{
	(*FileDownloadInfo)(nil),     // 0: download.FileDownloadInfo
	(*FileDownloadBody)(nil),     // 1: download.FileDownloadBody
	(*FileDownloadResponse)(nil), // 2: download.FileDownloadResponse
	(*FileDownloadRequest)(nil),  // 3: download.FileDownloadRequest
	(*FileDeleteRequest)(nil),    // 4: download.FileDeleteRequest
	(*FileDeleteResponse)(nil),   // 5: download.FileDeleteResponse
}
var file_pkg_download_downloadpb_download_proto_depIdxs = []int32{
	0, // 0: download.FileDownloadResponse.info:type_name -> download.FileDownloadInfo
	1, // 1: download.FileDownloadResponse.body:type_name -> download.FileDownloadBody
	3, // 2: download.FileDownloadService.DownloadFile:input_type -> download.FileDownloadRequest
	4, // 3: download.FileDownloadService.DeleteFile:input_type -> download.FileDeleteRequest
	2, // 4: download.FileDownloadService.DownloadFile:output_type -> download.FileDownloadResponse
	5, // 5: download.FileDownloadService.DeleteFile:output_type -> download.FileDeleteResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_download_downloadpb_download_proto_init() }
func file_pkg_download_downloadpb_download_proto_init() {
	if File_pkg_download_downloadpb_download_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_download_downloadpb_download_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileDownloadInfo); i {
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
		file_pkg_download_downloadpb_download_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileDownloadBody); i {
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
		file_pkg_download_downloadpb_download_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileDownloadResponse); i {
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
		file_pkg_download_downloadpb_download_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileDownloadRequest); i {
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
		file_pkg_download_downloadpb_download_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileDeleteRequest); i {
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
		file_pkg_download_downloadpb_download_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileDeleteResponse); i {
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
	file_pkg_download_downloadpb_download_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*FileDownloadResponse_Info)(nil),
		(*FileDownloadResponse_Body)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_download_downloadpb_download_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_download_downloadpb_download_proto_goTypes,
		DependencyIndexes: file_pkg_download_downloadpb_download_proto_depIdxs,
		MessageInfos:      file_pkg_download_downloadpb_download_proto_msgTypes,
	}.Build()
	File_pkg_download_downloadpb_download_proto = out.File
	file_pkg_download_downloadpb_download_proto_rawDesc = nil
	file_pkg_download_downloadpb_download_proto_goTypes = nil
	file_pkg_download_downloadpb_download_proto_depIdxs = nil
}
