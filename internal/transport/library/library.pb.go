// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: library.proto

package library

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

// Logic models
type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Country string `protobuf:"bytes,3,opt,name=Country,proto3" json:"Country,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_library_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_library_proto_rawDescGZIP(), []int{0}
}

func (x *Author) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Author) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	AuthorId    string `protobuf:"bytes,3,opt,name=AuthorId,proto3" json:"AuthorId,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_library_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_library_proto_rawDescGZIP(), []int{1}
}

func (x *Book) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *Book) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// Requests
type CUDAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation string  `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	Author    *Author `protobuf:"bytes,2,opt,name=Author,proto3" json:"Author,omitempty"`
}

func (x *CUDAuthorRequest) Reset() {
	*x = CUDAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CUDAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CUDAuthorRequest) ProtoMessage() {}

func (x *CUDAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_library_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CUDAuthorRequest.ProtoReflect.Descriptor instead.
func (*CUDAuthorRequest) Descriptor() ([]byte, []int) {
	return file_library_proto_rawDescGZIP(), []int{2}
}

func (x *CUDAuthorRequest) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *CUDAuthorRequest) GetAuthor() *Author {
	if x != nil {
		return x.Author
	}
	return nil
}

type CUDBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation string `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	Book      *Book  `protobuf:"bytes,2,opt,name=Book,proto3" json:"Book,omitempty"`
}

func (x *CUDBookRequest) Reset() {
	*x = CUDBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CUDBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CUDBookRequest) ProtoMessage() {}

func (x *CUDBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_library_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CUDBookRequest.ProtoReflect.Descriptor instead.
func (*CUDBookRequest) Descriptor() ([]byte, []int) {
	return file_library_proto_rawDescGZIP(), []int{3}
}

func (x *CUDBookRequest) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *CUDBookRequest) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

// Responses
type CUDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CUDResponse) Reset() {
	*x = CUDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CUDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CUDResponse) ProtoMessage() {}

func (x *CUDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_library_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CUDResponse.ProtoReflect.Descriptor instead.
func (*CUDResponse) Descriptor() ([]byte, []int) {
	return file_library_proto_rawDescGZIP(), []int{4}
}

func (x *CUDResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_library_proto protoreflect.FileDescriptor

var file_library_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x46, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x6a, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x51, 0x0a, 0x10, 0x43, 0x55, 0x44, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x06,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x49, 0x0a, 0x0e, 0x43, 0x55, 0x44, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x42, 0x6f, 0x6f,
	0x6b, 0x22, 0x25, 0x0a, 0x0b, 0x43, 0x55, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x71, 0x0a, 0x07, 0x4c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x12, 0x33, 0x0a, 0x10, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x43, 0x55, 0x44, 0x12, 0x11, 0x2e, 0x43, 0x55, 0x44, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x43, 0x55, 0x44,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0e, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x55, 0x44, 0x12, 0x11, 0x2e, 0x43, 0x55, 0x44,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e,
	0x43, 0x55, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_library_proto_rawDescOnce sync.Once
	file_library_proto_rawDescData = file_library_proto_rawDesc
)

func file_library_proto_rawDescGZIP() []byte {
	file_library_proto_rawDescOnce.Do(func() {
		file_library_proto_rawDescData = protoimpl.X.CompressGZIP(file_library_proto_rawDescData)
	})
	return file_library_proto_rawDescData
}

var file_library_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_library_proto_goTypes = []interface{}{
	(*Author)(nil),           // 0: Author
	(*Book)(nil),             // 1: Book
	(*CUDAuthorRequest)(nil), // 2: CUDAuthorRequest
	(*CUDBookRequest)(nil),   // 3: CUDBookRequest
	(*CUDResponse)(nil),      // 4: CUDResponse
}
var file_library_proto_depIdxs = []int32{
	0, // 0: CUDAuthorRequest.Author:type_name -> Author
	1, // 1: CUDBookRequest.Book:type_name -> Book
	2, // 2: Library.HandlerAuthorCUD:input_type -> CUDAuthorRequest
	2, // 3: Library.HandlerBookCUD:input_type -> CUDAuthorRequest
	4, // 4: Library.HandlerAuthorCUD:output_type -> CUDResponse
	4, // 5: Library.HandlerBookCUD:output_type -> CUDResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_library_proto_init() }
func file_library_proto_init() {
	if File_library_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_library_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
		file_library_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_library_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CUDAuthorRequest); i {
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
		file_library_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CUDBookRequest); i {
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
		file_library_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CUDResponse); i {
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
			RawDescriptor: file_library_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_library_proto_goTypes,
		DependencyIndexes: file_library_proto_depIdxs,
		MessageInfos:      file_library_proto_msgTypes,
	}.Build()
	File_library_proto = out.File
	file_library_proto_rawDesc = nil
	file_library_proto_goTypes = nil
	file_library_proto_depIdxs = nil
}