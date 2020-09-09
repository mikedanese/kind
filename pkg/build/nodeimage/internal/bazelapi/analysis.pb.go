// Copyright 2020 The Kubernetes Authors.
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

// We depend on these types in the bazel build. These are the insteresting bits
// of analysis extracted from:
//
// https://cs.opensource.google/bazel/bazel/+/master:src/main/protobuf/analysis.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: pkg/build/nodeimage/internal/bazelapi/analysis.proto

package bazelapi

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ActionGraphContainer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Artifacts []*Artifact `protobuf:"bytes,1,rep,name=artifacts,proto3" json:"artifacts,omitempty"`
	Actions   []*Action   `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *ActionGraphContainer) Reset() {
	*x = ActionGraphContainer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionGraphContainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionGraphContainer) ProtoMessage() {}

func (x *ActionGraphContainer) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionGraphContainer.ProtoReflect.Descriptor instead.
func (*ActionGraphContainer) Descriptor() ([]byte, []int) {
	return file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_rawDescGZIP(), []int{0}
}

func (x *ActionGraphContainer) GetArtifacts() []*Artifact {
	if x != nil {
		return x.Artifacts
	}
	return nil
}

func (x *ActionGraphContainer) GetActions() []*Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

type Artifact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ExecPath string `protobuf:"bytes,2,opt,name=exec_path,json=execPath,proto3" json:"exec_path,omitempty"`
}

func (x *Artifact) Reset() {
	*x = Artifact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Artifact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Artifact) ProtoMessage() {}

func (x *Artifact) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Artifact.ProtoReflect.Descriptor instead.
func (*Artifact) Descriptor() ([]byte, []int) {
	return file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_rawDescGZIP(), []int{1}
}

func (x *Artifact) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Artifact) GetExecPath() string {
	if x != nil {
		return x.ExecPath
	}
	return ""
}

type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mnemonic  string   `protobuf:"bytes,4,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"`
	OutputIds []string `protobuf:"bytes,9,rep,name=output_ids,json=outputIds,proto3" json:"output_ids,omitempty"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_rawDescGZIP(), []int{2}
}

func (x *Action) GetMnemonic() string {
	if x != nil {
		return x.Mnemonic
	}
	return ""
}

func (x *Action) GetOutputIds() []string {
	if x != nil {
		return x.OutputIds
	}
	return nil
}

var File_pkg_build_nodeimage_internal_bazelapi_analysis_proto protoreflect.FileDescriptor

var file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_rawDesc = []byte{
	0x0a, 0x34, 0x70, 0x6b, 0x67, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x6e, 0x6f, 0x64, 0x65,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x62,
	0x61, 0x7a, 0x65, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x61, 0x70, 0x69,
	0x22, 0x74, 0x0a, 0x14, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x09, 0x61, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x61,
	0x7a, 0x65, 0x6c, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52,
	0x09, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x07, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x61,
	0x7a, 0x65, 0x6c, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x37, 0x0a, 0x08, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x65, 0x63, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x65, 0x63, 0x50, 0x61, 0x74, 0x68, 0x22,
	0x43, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6e, 0x65,
	0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6e, 0x65,
	0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x49, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_rawDescOnce sync.Once
	file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_rawDescData = file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_rawDesc
)

func file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_rawDescGZIP() []byte {
	file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_rawDescOnce.Do(func() {
		file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_rawDescData)
	})
	return file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_rawDescData
}

var file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_goTypes = []interface{}{
	(*ActionGraphContainer)(nil), // 0: bazelapi.ActionGraphContainer
	(*Artifact)(nil),             // 1: bazelapi.Artifact
	(*Action)(nil),               // 2: bazelapi.Action
}
var file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_depIdxs = []int32{
	1, // 0: bazelapi.ActionGraphContainer.artifacts:type_name -> bazelapi.Artifact
	2, // 1: bazelapi.ActionGraphContainer.actions:type_name -> bazelapi.Action
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_init() }
func file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_init() {
	if File_pkg_build_nodeimage_internal_bazelapi_analysis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionGraphContainer); i {
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
		file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Artifact); i {
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
		file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
			RawDescriptor: file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_goTypes,
		DependencyIndexes: file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_depIdxs,
		MessageInfos:      file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_msgTypes,
	}.Build()
	File_pkg_build_nodeimage_internal_bazelapi_analysis_proto = out.File
	file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_rawDesc = nil
	file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_goTypes = nil
	file_pkg_build_nodeimage_internal_bazelapi_analysis_proto_depIdxs = nil
}
