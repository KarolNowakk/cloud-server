// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: pkg/upload/uploadpb/upload.proto

package uploadpb

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

type FileUploadInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SearchTags string `protobuf:"bytes,3,opt,name=searchTags,proto3" json:"searchTags,omitempty"`
}

func (x *FileUploadInfo) Reset() {
	*x = FileUploadInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_upload_uploadpb_upload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUploadInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadInfo) ProtoMessage() {}

func (x *FileUploadInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_upload_uploadpb_upload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadInfo.ProtoReflect.Descriptor instead.
func (*FileUploadInfo) Descriptor() ([]byte, []int) {
	return file_pkg_upload_uploadpb_upload_proto_rawDescGZIP(), []int{0}
}

func (x *FileUploadInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileUploadInfo) GetSearchTags() string {
	if x != nil {
		return x.SearchTags
	}
	return ""
}

type FileUploadBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bytes []byte `protobuf:"bytes,4,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (x *FileUploadBody) Reset() {
	*x = FileUploadBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_upload_uploadpb_upload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUploadBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadBody) ProtoMessage() {}

func (x *FileUploadBody) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_upload_uploadpb_upload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadBody.ProtoReflect.Descriptor instead.
func (*FileUploadBody) Descriptor() ([]byte, []int) {
	return file_pkg_upload_uploadpb_upload_proto_rawDescGZIP(), []int{1}
}

func (x *FileUploadBody) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

type FileUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*FileUploadRequest_Info
	//	*FileUploadRequest_Body
	Data isFileUploadRequest_Data `protobuf_oneof:"data"`
}

func (x *FileUploadRequest) Reset() {
	*x = FileUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_upload_uploadpb_upload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadRequest) ProtoMessage() {}

func (x *FileUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_upload_uploadpb_upload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadRequest.ProtoReflect.Descriptor instead.
func (*FileUploadRequest) Descriptor() ([]byte, []int) {
	return file_pkg_upload_uploadpb_upload_proto_rawDescGZIP(), []int{2}
}

func (m *FileUploadRequest) GetData() isFileUploadRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *FileUploadRequest) GetInfo() *FileUploadInfo {
	if x, ok := x.GetData().(*FileUploadRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (x *FileUploadRequest) GetBody() *FileUploadBody {
	if x, ok := x.GetData().(*FileUploadRequest_Body); ok {
		return x.Body
	}
	return nil
}

type isFileUploadRequest_Data interface {
	isFileUploadRequest_Data()
}

type FileUploadRequest_Info struct {
	Info *FileUploadInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type FileUploadRequest_Body struct {
	Body *FileUploadBody `protobuf:"bytes,2,opt,name=body,proto3,oneof"`
}

func (*FileUploadRequest_Info) isFileUploadRequest_Data() {}

func (*FileUploadRequest_Body) isFileUploadRequest_Data() {}

type FileUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok  bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *FileUploadResponse) Reset() {
	*x = FileUploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_upload_uploadpb_upload_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadResponse) ProtoMessage() {}

func (x *FileUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_upload_uploadpb_upload_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadResponse.ProtoReflect.Descriptor instead.
func (*FileUploadResponse) Descriptor() ([]byte, []int) {
	return file_pkg_upload_uploadpb_upload_proto_rawDescGZIP(), []int{3}
}

func (x *FileUploadResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *FileUploadResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_pkg_upload_uploadpb_upload_proto protoreflect.FileDescriptor

var file_pkg_upload_uploadpb_upload_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x6b, 0x67, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x70, 0x62, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x44, 0x0a, 0x0e, 0x46, 0x69,
	0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x61, 0x67, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x61, 0x67, 0x73,
	0x22, 0x26, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6f,
	0x64, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x22, 0x77, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x65,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6f, 0x64,
	0x79, 0x48, 0x00, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x36, 0x0a, 0x12, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x5c, 0x0a, 0x11, 0x46, 0x69, 0x6c,
	0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47,
	0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x19, 0x2e, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_upload_uploadpb_upload_proto_rawDescOnce sync.Once
	file_pkg_upload_uploadpb_upload_proto_rawDescData = file_pkg_upload_uploadpb_upload_proto_rawDesc
)

func file_pkg_upload_uploadpb_upload_proto_rawDescGZIP() []byte {
	file_pkg_upload_uploadpb_upload_proto_rawDescOnce.Do(func() {
		file_pkg_upload_uploadpb_upload_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_upload_uploadpb_upload_proto_rawDescData)
	})
	return file_pkg_upload_uploadpb_upload_proto_rawDescData
}

var file_pkg_upload_uploadpb_upload_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_upload_uploadpb_upload_proto_goTypes = []interface{}{
	(*FileUploadInfo)(nil),     // 0: upload.FileUploadInfo
	(*FileUploadBody)(nil),     // 1: upload.FileUploadBody
	(*FileUploadRequest)(nil),  // 2: upload.FileUploadRequest
	(*FileUploadResponse)(nil), // 3: upload.FileUploadResponse
}
var file_pkg_upload_uploadpb_upload_proto_depIdxs = []int32{
	0, // 0: upload.FileUploadRequest.info:type_name -> upload.FileUploadInfo
	1, // 1: upload.FileUploadRequest.body:type_name -> upload.FileUploadBody
	2, // 2: upload.FileUploadService.UploadFile:input_type -> upload.FileUploadRequest
	3, // 3: upload.FileUploadService.UploadFile:output_type -> upload.FileUploadResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_upload_uploadpb_upload_proto_init() }
func file_pkg_upload_uploadpb_upload_proto_init() {
	if File_pkg_upload_uploadpb_upload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_upload_uploadpb_upload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUploadInfo); i {
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
		file_pkg_upload_uploadpb_upload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUploadBody); i {
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
		file_pkg_upload_uploadpb_upload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUploadRequest); i {
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
		file_pkg_upload_uploadpb_upload_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUploadResponse); i {
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
	file_pkg_upload_uploadpb_upload_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*FileUploadRequest_Info)(nil),
		(*FileUploadRequest_Body)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_upload_uploadpb_upload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_upload_uploadpb_upload_proto_goTypes,
		DependencyIndexes: file_pkg_upload_uploadpb_upload_proto_depIdxs,
		MessageInfos:      file_pkg_upload_uploadpb_upload_proto_msgTypes,
	}.Build()
	File_pkg_upload_uploadpb_upload_proto = out.File
	file_pkg_upload_uploadpb_upload_proto_rawDesc = nil
	file_pkg_upload_uploadpb_upload_proto_goTypes = nil
	file_pkg_upload_uploadpb_upload_proto_depIdxs = nil
}