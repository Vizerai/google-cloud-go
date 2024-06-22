// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: google/cloud/aiplatform/v1beta1/notebook_idle_shutdown_config.proto

package aiplatformpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The idle shutdown configuration of NotebookRuntimeTemplate, which contains
// the idle_timeout as required field.
type NotebookIdleShutdownConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Duration is accurate to the second. In Notebook, Idle Timeout is
	// accurate to minute so the range of idle_timeout (second) is: 10 * 60 ~ 1440
	// * 60.
	IdleTimeout *durationpb.Duration `protobuf:"bytes,1,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	// Whether Idle Shutdown is disabled in this NotebookRuntimeTemplate.
	IdleShutdownDisabled bool `protobuf:"varint,2,opt,name=idle_shutdown_disabled,json=idleShutdownDisabled,proto3" json:"idle_shutdown_disabled,omitempty"`
}

func (x *NotebookIdleShutdownConfig) Reset() {
	*x = NotebookIdleShutdownConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotebookIdleShutdownConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotebookIdleShutdownConfig) ProtoMessage() {}

func (x *NotebookIdleShutdownConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotebookIdleShutdownConfig.ProtoReflect.Descriptor instead.
func (*NotebookIdleShutdownConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_rawDescGZIP(), []int{0}
}

func (x *NotebookIdleShutdownConfig) GetIdleTimeout() *durationpb.Duration {
	if x != nil {
		return x.IdleTimeout
	}
	return nil
}

func (x *NotebookIdleShutdownConfig) GetIdleShutdownDisabled() bool {
	if x != nil {
		return x.IdleShutdownDisabled
	}
	return false
}

var File_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x5f,
	0x73, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x1a, 0x4e, 0x6f, 0x74, 0x65,
	0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x6c, 0x65, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x41, 0x0a, 0x0c, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x69, 0x64,
	0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x69, 0x64, 0x6c,
	0x65, 0x5f, 0x73, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x64, 0x6c, 0x65, 0x53,
	0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42,
	0xf6, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x1f, 0x4e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f,
	0x6b, 0x49, 0x64, 0x6c, 0x65, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x70, 0x62, 0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa,
	0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41,
	0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61,
	0x31, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_goTypes = []any{
	(*NotebookIdleShutdownConfig)(nil), // 0: google.cloud.aiplatform.v1beta1.NotebookIdleShutdownConfig
	(*durationpb.Duration)(nil),        // 1: google.protobuf.Duration
}
var file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_depIdxs = []int32{
	1, // 0: google.cloud.aiplatform.v1beta1.NotebookIdleShutdownConfig.idle_timeout:type_name -> google.protobuf.Duration
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_init() }
func file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NotebookIdleShutdownConfig); i {
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
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto = out.File
	file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_notebook_idle_shutdown_config_proto_depIdxs = nil
}
