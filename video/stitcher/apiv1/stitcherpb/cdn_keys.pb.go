// Copyright 2022 Google LLC
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/cloud/video/stitcher/v1/cdn_keys.proto

package stitcherpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Configuration for a CDN key. Used by the Video Stitcher
// to sign URIs for fetching video manifests and signing
// media segments for playback.
type CdnKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration associated with the CDN key.
	//
	// Types that are assignable to CdnKeyConfig:
	//
	//	*CdnKey_GoogleCdnKey
	//	*CdnKey_AkamaiCdnKey
	CdnKeyConfig isCdnKey_CdnKeyConfig `protobuf_oneof:"cdn_key_config"`
	// The resource name of the CDN key, in the form of
	// `projects/{project}/locations/{location}/cdnKeys/{id}`.
	// The name is ignored when creating a CDN key.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The hostname this key applies to.
	Hostname string `protobuf:"bytes,4,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *CdnKey) Reset() {
	*x = CdnKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_video_stitcher_v1_cdn_keys_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CdnKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CdnKey) ProtoMessage() {}

func (x *CdnKey) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_video_stitcher_v1_cdn_keys_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CdnKey.ProtoReflect.Descriptor instead.
func (*CdnKey) Descriptor() ([]byte, []int) {
	return file_google_cloud_video_stitcher_v1_cdn_keys_proto_rawDescGZIP(), []int{0}
}

func (m *CdnKey) GetCdnKeyConfig() isCdnKey_CdnKeyConfig {
	if m != nil {
		return m.CdnKeyConfig
	}
	return nil
}

func (x *CdnKey) GetGoogleCdnKey() *GoogleCdnKey {
	if x, ok := x.GetCdnKeyConfig().(*CdnKey_GoogleCdnKey); ok {
		return x.GoogleCdnKey
	}
	return nil
}

func (x *CdnKey) GetAkamaiCdnKey() *AkamaiCdnKey {
	if x, ok := x.GetCdnKeyConfig().(*CdnKey_AkamaiCdnKey); ok {
		return x.AkamaiCdnKey
	}
	return nil
}

func (x *CdnKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CdnKey) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

type isCdnKey_CdnKeyConfig interface {
	isCdnKey_CdnKeyConfig()
}

type CdnKey_GoogleCdnKey struct {
	// The configuration for a Google Cloud CDN key.
	GoogleCdnKey *GoogleCdnKey `protobuf:"bytes,5,opt,name=google_cdn_key,json=googleCdnKey,proto3,oneof"`
}

type CdnKey_AkamaiCdnKey struct {
	// The configuration for an Akamai CDN key.
	AkamaiCdnKey *AkamaiCdnKey `protobuf:"bytes,6,opt,name=akamai_cdn_key,json=akamaiCdnKey,proto3,oneof"`
}

func (*CdnKey_GoogleCdnKey) isCdnKey_CdnKeyConfig() {}

func (*CdnKey_AkamaiCdnKey) isCdnKey_CdnKeyConfig() {}

// Configuration for a Google Cloud CDN key.
type GoogleCdnKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Input only. Secret for this Google Cloud CDN key.
	PrivateKey []byte `protobuf:"bytes,1,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// The public name of the Google Cloud CDN key.
	KeyName string `protobuf:"bytes,2,opt,name=key_name,json=keyName,proto3" json:"key_name,omitempty"`
}

func (x *GoogleCdnKey) Reset() {
	*x = GoogleCdnKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_video_stitcher_v1_cdn_keys_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleCdnKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleCdnKey) ProtoMessage() {}

func (x *GoogleCdnKey) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_video_stitcher_v1_cdn_keys_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleCdnKey.ProtoReflect.Descriptor instead.
func (*GoogleCdnKey) Descriptor() ([]byte, []int) {
	return file_google_cloud_video_stitcher_v1_cdn_keys_proto_rawDescGZIP(), []int{1}
}

func (x *GoogleCdnKey) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (x *GoogleCdnKey) GetKeyName() string {
	if x != nil {
		return x.KeyName
	}
	return ""
}

// Configuration for an Akamai CDN key.
type AkamaiCdnKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Input only. Token key for the Akamai CDN edge configuration.
	TokenKey []byte `protobuf:"bytes,1,opt,name=token_key,json=tokenKey,proto3" json:"token_key,omitempty"`
}

func (x *AkamaiCdnKey) Reset() {
	*x = AkamaiCdnKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_video_stitcher_v1_cdn_keys_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AkamaiCdnKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AkamaiCdnKey) ProtoMessage() {}

func (x *AkamaiCdnKey) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_video_stitcher_v1_cdn_keys_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AkamaiCdnKey.ProtoReflect.Descriptor instead.
func (*AkamaiCdnKey) Descriptor() ([]byte, []int) {
	return file_google_cloud_video_stitcher_v1_cdn_keys_proto_rawDescGZIP(), []int{2}
}

func (x *AkamaiCdnKey) GetTokenKey() []byte {
	if x != nil {
		return x.TokenKey
	}
	return nil
}

var File_google_cloud_video_stitcher_v1_cdn_keys_proto protoreflect.FileDescriptor

var file_google_cloud_video_stitcher_v1_cdn_keys_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2f, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x64, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x02, 0x0a, 0x06,
	0x43, 0x64, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x54, 0x0a, 0x0e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5f, 0x63, 0x64, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x64, 0x6e, 0x4b, 0x65, 0x79, 0x48, 0x00, 0x52, 0x0c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x64, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x54, 0x0a, 0x0e,
	0x61, 0x6b, 0x61, 0x6d, 0x61, 0x69, 0x5f, 0x63, 0x64, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6b, 0x61, 0x6d, 0x61, 0x69, 0x43, 0x64, 0x6e, 0x4b,
	0x65, 0x79, 0x48, 0x00, 0x52, 0x0c, 0x61, 0x6b, 0x61, 0x6d, 0x61, 0x69, 0x43, 0x64, 0x6e, 0x4b,
	0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x3a, 0x63, 0xea, 0x41, 0x60, 0x0a, 0x23, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x74,
	0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x64, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x39, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x64, 0x6e, 0x4b, 0x65, 0x79, 0x73, 0x2f, 0x7b, 0x63,
	0x64, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x7d, 0x42, 0x10, 0x0a, 0x0e, 0x63, 0x64, 0x6e, 0x5f, 0x6b,
	0x65, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x4f, 0x0a, 0x0c, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x43, 0x64, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x0b, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x03,
	0xe0, 0x41, 0x04, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12,
	0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x30, 0x0a, 0x0c, 0x41, 0x6b,
	0x61, 0x6d, 0x61, 0x69, 0x43, 0x64, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x09, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x03, 0xe0,
	0x41, 0x04, 0x52, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4b, 0x65, 0x79, 0x42, 0x7c, 0x0a, 0x22,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x42, 0x0c, 0x43, 0x64, 0x6e, 0x4b, 0x65, 0x79, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x46, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x3b, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_cloud_video_stitcher_v1_cdn_keys_proto_rawDescOnce sync.Once
	file_google_cloud_video_stitcher_v1_cdn_keys_proto_rawDescData = file_google_cloud_video_stitcher_v1_cdn_keys_proto_rawDesc
)

func file_google_cloud_video_stitcher_v1_cdn_keys_proto_rawDescGZIP() []byte {
	file_google_cloud_video_stitcher_v1_cdn_keys_proto_rawDescOnce.Do(func() {
		file_google_cloud_video_stitcher_v1_cdn_keys_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_video_stitcher_v1_cdn_keys_proto_rawDescData)
	})
	return file_google_cloud_video_stitcher_v1_cdn_keys_proto_rawDescData
}

var file_google_cloud_video_stitcher_v1_cdn_keys_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_video_stitcher_v1_cdn_keys_proto_goTypes = []interface{}{
	(*CdnKey)(nil),       // 0: google.cloud.video.stitcher.v1.CdnKey
	(*GoogleCdnKey)(nil), // 1: google.cloud.video.stitcher.v1.GoogleCdnKey
	(*AkamaiCdnKey)(nil), // 2: google.cloud.video.stitcher.v1.AkamaiCdnKey
}
var file_google_cloud_video_stitcher_v1_cdn_keys_proto_depIdxs = []int32{
	1, // 0: google.cloud.video.stitcher.v1.CdnKey.google_cdn_key:type_name -> google.cloud.video.stitcher.v1.GoogleCdnKey
	2, // 1: google.cloud.video.stitcher.v1.CdnKey.akamai_cdn_key:type_name -> google.cloud.video.stitcher.v1.AkamaiCdnKey
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_video_stitcher_v1_cdn_keys_proto_init() }
func file_google_cloud_video_stitcher_v1_cdn_keys_proto_init() {
	if File_google_cloud_video_stitcher_v1_cdn_keys_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_video_stitcher_v1_cdn_keys_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CdnKey); i {
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
		file_google_cloud_video_stitcher_v1_cdn_keys_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleCdnKey); i {
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
		file_google_cloud_video_stitcher_v1_cdn_keys_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AkamaiCdnKey); i {
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
	file_google_cloud_video_stitcher_v1_cdn_keys_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CdnKey_GoogleCdnKey)(nil),
		(*CdnKey_AkamaiCdnKey)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_video_stitcher_v1_cdn_keys_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_video_stitcher_v1_cdn_keys_proto_goTypes,
		DependencyIndexes: file_google_cloud_video_stitcher_v1_cdn_keys_proto_depIdxs,
		MessageInfos:      file_google_cloud_video_stitcher_v1_cdn_keys_proto_msgTypes,
	}.Build()
	File_google_cloud_video_stitcher_v1_cdn_keys_proto = out.File
	file_google_cloud_video_stitcher_v1_cdn_keys_proto_rawDesc = nil
	file_google_cloud_video_stitcher_v1_cdn_keys_proto_goTypes = nil
	file_google_cloud_video_stitcher_v1_cdn_keys_proto_depIdxs = nil
}
