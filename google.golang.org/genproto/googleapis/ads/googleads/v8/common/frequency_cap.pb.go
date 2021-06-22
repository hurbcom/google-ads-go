// Copyright 2021 Google LLC
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
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.2
// source: google/ads/googleads/v8/common/frequency_cap.proto

package common

import (
	enums "google.golang.org/genproto/googleapis/ads/googleads/v8/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A rule specifying the maximum number of times an ad (or some set of ads) can
// be shown to a user over a particular time period.
type FrequencyCapEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The key of a particular frequency cap. There can be no more
	// than one frequency cap with the same key.
	Key *FrequencyCapKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Maximum number of events allowed during the time range by this cap.
	Cap *int32 `protobuf:"varint,3,opt,name=cap,proto3,oneof" json:"cap,omitempty"`
}

func (x *FrequencyCapEntry) Reset() {
	*x = FrequencyCapEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_common_frequency_cap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrequencyCapEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrequencyCapEntry) ProtoMessage() {}

func (x *FrequencyCapEntry) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_common_frequency_cap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrequencyCapEntry.ProtoReflect.Descriptor instead.
func (*FrequencyCapEntry) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_common_frequency_cap_proto_rawDescGZIP(), []int{0}
}

func (x *FrequencyCapEntry) GetKey() *FrequencyCapKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *FrequencyCapEntry) GetCap() int32 {
	if x != nil && x.Cap != nil {
		return *x.Cap
	}
	return 0
}

// A group of fields used as keys for a frequency cap.
// There can be no more than one frequency cap with the same key.
type FrequencyCapKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The level on which the cap is to be applied (e.g. ad group ad, ad group).
	// The cap is applied to all the entities of this level.
	Level enums.FrequencyCapLevelEnum_FrequencyCapLevel `protobuf:"varint,1,opt,name=level,proto3,enum=google.ads.googleads.v8.enums.FrequencyCapLevelEnum_FrequencyCapLevel" json:"level,omitempty"`
	// The type of event that the cap applies to (e.g. impression).
	EventType enums.FrequencyCapEventTypeEnum_FrequencyCapEventType `protobuf:"varint,3,opt,name=event_type,json=eventType,proto3,enum=google.ads.googleads.v8.enums.FrequencyCapEventTypeEnum_FrequencyCapEventType" json:"event_type,omitempty"`
	// Unit of time the cap is defined at (e.g. day, week).
	TimeUnit enums.FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit `protobuf:"varint,2,opt,name=time_unit,json=timeUnit,proto3,enum=google.ads.googleads.v8.enums.FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit" json:"time_unit,omitempty"`
	// Number of time units the cap lasts.
	TimeLength *int32 `protobuf:"varint,5,opt,name=time_length,json=timeLength,proto3,oneof" json:"time_length,omitempty"`
}

func (x *FrequencyCapKey) Reset() {
	*x = FrequencyCapKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_common_frequency_cap_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrequencyCapKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrequencyCapKey) ProtoMessage() {}

func (x *FrequencyCapKey) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_common_frequency_cap_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrequencyCapKey.ProtoReflect.Descriptor instead.
func (*FrequencyCapKey) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_common_frequency_cap_proto_rawDescGZIP(), []int{1}
}

func (x *FrequencyCapKey) GetLevel() enums.FrequencyCapLevelEnum_FrequencyCapLevel {
	if x != nil {
		return x.Level
	}
	return enums.FrequencyCapLevelEnum_UNSPECIFIED
}

func (x *FrequencyCapKey) GetEventType() enums.FrequencyCapEventTypeEnum_FrequencyCapEventType {
	if x != nil {
		return x.EventType
	}
	return enums.FrequencyCapEventTypeEnum_UNSPECIFIED
}

func (x *FrequencyCapKey) GetTimeUnit() enums.FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit {
	if x != nil {
		return x.TimeUnit
	}
	return enums.FrequencyCapTimeUnitEnum_UNSPECIFIED
}

func (x *FrequencyCapKey) GetTimeLength() int32 {
	if x != nil && x.TimeLength != nil {
		return *x.TimeLength
	}
	return 0
}

var File_google_ads_googleads_v8_common_frequency_cap_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v8_common_frequency_cap_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x61, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x61,
	0x70, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2f, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x61, 0x70, 0x5f,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x66, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x61, 0x70, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x6e,
	0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x11, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x79, 0x43, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x41, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x79, 0x43, 0x61, 0x70, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x15,
	0x0a, 0x03, 0x63, 0x61, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x03, 0x63,
	0x61, 0x70, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x63, 0x61, 0x70, 0x22, 0xff, 0x02,
	0x0a, 0x0f, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x61, 0x70, 0x4b, 0x65,
	0x79, 0x12, 0x5c, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x46, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x61, 0x70, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79,
	0x43, 0x61, 0x70, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x6d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x4e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x61, 0x70,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x46, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x61, 0x70, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x69,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x4c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x61, 0x70, 0x54, 0x69,
	0x6d, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x46, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x79, 0x43, 0x61, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x52,
	0x08, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00,
	0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x88, 0x01, 0x01, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x42,
	0xec, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x11, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x79, 0x43, 0x61, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x38, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x38, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x38, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x56, 0x38, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v8_common_frequency_cap_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v8_common_frequency_cap_proto_rawDescData = file_google_ads_googleads_v8_common_frequency_cap_proto_rawDesc
)

func file_google_ads_googleads_v8_common_frequency_cap_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v8_common_frequency_cap_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v8_common_frequency_cap_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v8_common_frequency_cap_proto_rawDescData)
	})
	return file_google_ads_googleads_v8_common_frequency_cap_proto_rawDescData
}

var file_google_ads_googleads_v8_common_frequency_cap_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_ads_googleads_v8_common_frequency_cap_proto_goTypes = []interface{}{
	(*FrequencyCapEntry)(nil),                                  // 0: google.ads.googleads.v8.common.FrequencyCapEntry
	(*FrequencyCapKey)(nil),                                    // 1: google.ads.googleads.v8.common.FrequencyCapKey
	(enums.FrequencyCapLevelEnum_FrequencyCapLevel)(0),         // 2: google.ads.googleads.v8.enums.FrequencyCapLevelEnum.FrequencyCapLevel
	(enums.FrequencyCapEventTypeEnum_FrequencyCapEventType)(0), // 3: google.ads.googleads.v8.enums.FrequencyCapEventTypeEnum.FrequencyCapEventType
	(enums.FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit)(0),   // 4: google.ads.googleads.v8.enums.FrequencyCapTimeUnitEnum.FrequencyCapTimeUnit
}
var file_google_ads_googleads_v8_common_frequency_cap_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v8.common.FrequencyCapEntry.key:type_name -> google.ads.googleads.v8.common.FrequencyCapKey
	2, // 1: google.ads.googleads.v8.common.FrequencyCapKey.level:type_name -> google.ads.googleads.v8.enums.FrequencyCapLevelEnum.FrequencyCapLevel
	3, // 2: google.ads.googleads.v8.common.FrequencyCapKey.event_type:type_name -> google.ads.googleads.v8.enums.FrequencyCapEventTypeEnum.FrequencyCapEventType
	4, // 3: google.ads.googleads.v8.common.FrequencyCapKey.time_unit:type_name -> google.ads.googleads.v8.enums.FrequencyCapTimeUnitEnum.FrequencyCapTimeUnit
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v8_common_frequency_cap_proto_init() }
func file_google_ads_googleads_v8_common_frequency_cap_proto_init() {
	if File_google_ads_googleads_v8_common_frequency_cap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v8_common_frequency_cap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrequencyCapEntry); i {
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
		file_google_ads_googleads_v8_common_frequency_cap_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrequencyCapKey); i {
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
	file_google_ads_googleads_v8_common_frequency_cap_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_google_ads_googleads_v8_common_frequency_cap_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v8_common_frequency_cap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v8_common_frequency_cap_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v8_common_frequency_cap_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v8_common_frequency_cap_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v8_common_frequency_cap_proto = out.File
	file_google_ads_googleads_v8_common_frequency_cap_proto_rawDesc = nil
	file_google_ads_googleads_v8_common_frequency_cap_proto_goTypes = nil
	file_google_ads_googleads_v8_common_frequency_cap_proto_depIdxs = nil
}
