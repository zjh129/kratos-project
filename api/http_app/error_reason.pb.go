// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.19.4
// source: http_app/error_reason.proto

package http_app

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ErrorReason int32

const (
	ErrorReason_GREETER_UNSPECIFIED ErrorReason = 0
	ErrorReason_USER_LOGIN_FAILED   ErrorReason = 1
	ErrorReason_USER_NOT_FOUND      ErrorReason = 2
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "GREETER_UNSPECIFIED",
		1: "USER_LOGIN_FAILED",
		2: "USER_NOT_FOUND",
	}
	ErrorReason_value = map[string]int32{
		"GREETER_UNSPECIFIED": 0,
		"USER_LOGIN_FAILED":   1,
		"USER_NOT_FOUND":      2,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_http_app_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_http_app_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_http_app_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_http_app_error_reason_proto protoreflect.FileDescriptor

var file_http_app_error_reason_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x61, 0x70, 0x70, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x68,
	0x74, 0x74, 0x70, 0x5f, 0x61, 0x70, 0x70, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x63, 0x0a, 0x0b,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x13, 0x47,
	0x52, 0x45, 0x45, 0x54, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x11, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4c, 0x4f, 0x47,
	0x49, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x91,
	0x03, 0x12, 0x18, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4,
	0x03, 0x42, 0x26, 0x5a, 0x24, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x61, 0x70, 0x70,
	0x3b, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x61, 0x70, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_http_app_error_reason_proto_rawDescOnce sync.Once
	file_http_app_error_reason_proto_rawDescData = file_http_app_error_reason_proto_rawDesc
)

func file_http_app_error_reason_proto_rawDescGZIP() []byte {
	file_http_app_error_reason_proto_rawDescOnce.Do(func() {
		file_http_app_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_http_app_error_reason_proto_rawDescData)
	})
	return file_http_app_error_reason_proto_rawDescData
}

var file_http_app_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_http_app_error_reason_proto_goTypes = []any{
	(ErrorReason)(0), // 0: http_app.ErrorReason
}
var file_http_app_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_http_app_error_reason_proto_init() }
func file_http_app_error_reason_proto_init() {
	if File_http_app_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_http_app_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_http_app_error_reason_proto_goTypes,
		DependencyIndexes: file_http_app_error_reason_proto_depIdxs,
		EnumInfos:         file_http_app_error_reason_proto_enumTypes,
	}.Build()
	File_http_app_error_reason_proto = out.File
	file_http_app_error_reason_proto_rawDesc = nil
	file_http_app_error_reason_proto_goTypes = nil
	file_http_app_error_reason_proto_depIdxs = nil
}
