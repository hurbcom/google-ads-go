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
// source: google/ads/googleads/v8/enums/reach_plan_age_range.proto

package enums

import (
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

// Possible plannable age range values.
type ReachPlanAgeRangeEnum_ReachPlanAgeRange int32

const (
	// Not specified.
	ReachPlanAgeRangeEnum_UNSPECIFIED ReachPlanAgeRangeEnum_ReachPlanAgeRange = 0
	// The value is unknown in this version.
	ReachPlanAgeRangeEnum_UNKNOWN ReachPlanAgeRangeEnum_ReachPlanAgeRange = 1
	// Between 18 and 24 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_18_24 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 503001
	// Between 18 and 34 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_18_34 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 2
	// Between 18 and 44 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_18_44 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 3
	// Between 18 and 49 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_18_49 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 4
	// Between 18 and 54 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_18_54 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 5
	// Between 18 and 64 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_18_64 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 6
	// Between 18 and 65+ years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_18_65_UP ReachPlanAgeRangeEnum_ReachPlanAgeRange = 7
	// Between 21 and 34 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_21_34 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 8
	// Between 25 and 34 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_25_34 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 503002
	// Between 25 and 44 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_25_44 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 9
	// Between 25 and 49 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_25_49 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 10
	// Between 25 and 54 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_25_54 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 11
	// Between 25 and 64 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_25_64 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 12
	// Between 25 and 65+ years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_25_65_UP ReachPlanAgeRangeEnum_ReachPlanAgeRange = 13
	// Between 35 and 44 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_35_44 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 503003
	// Between 35 and 49 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_35_49 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 14
	// Between 35 and 54 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_35_54 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 15
	// Between 35 and 64 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_35_64 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 16
	// Between 35 and 65+ years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_35_65_UP ReachPlanAgeRangeEnum_ReachPlanAgeRange = 17
	// Between 45 and 54 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_45_54 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 503004
	// Between 45 and 64 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_45_64 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 18
	// Between 45 and 65+ years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_45_65_UP ReachPlanAgeRangeEnum_ReachPlanAgeRange = 19
	// Between 50 and 65+ years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_50_65_UP ReachPlanAgeRangeEnum_ReachPlanAgeRange = 20
	// Between 55 and 64 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_55_64 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 503005
	// Between 55 and 65+ years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_55_65_UP ReachPlanAgeRangeEnum_ReachPlanAgeRange = 21
	// 65 years old and beyond.
	ReachPlanAgeRangeEnum_AGE_RANGE_65_UP ReachPlanAgeRangeEnum_ReachPlanAgeRange = 503006
)

// Enum value maps for ReachPlanAgeRangeEnum_ReachPlanAgeRange.
var (
	ReachPlanAgeRangeEnum_ReachPlanAgeRange_name = map[int32]string{
		0:      "UNSPECIFIED",
		1:      "UNKNOWN",
		503001: "AGE_RANGE_18_24",
		2:      "AGE_RANGE_18_34",
		3:      "AGE_RANGE_18_44",
		4:      "AGE_RANGE_18_49",
		5:      "AGE_RANGE_18_54",
		6:      "AGE_RANGE_18_64",
		7:      "AGE_RANGE_18_65_UP",
		8:      "AGE_RANGE_21_34",
		503002: "AGE_RANGE_25_34",
		9:      "AGE_RANGE_25_44",
		10:     "AGE_RANGE_25_49",
		11:     "AGE_RANGE_25_54",
		12:     "AGE_RANGE_25_64",
		13:     "AGE_RANGE_25_65_UP",
		503003: "AGE_RANGE_35_44",
		14:     "AGE_RANGE_35_49",
		15:     "AGE_RANGE_35_54",
		16:     "AGE_RANGE_35_64",
		17:     "AGE_RANGE_35_65_UP",
		503004: "AGE_RANGE_45_54",
		18:     "AGE_RANGE_45_64",
		19:     "AGE_RANGE_45_65_UP",
		20:     "AGE_RANGE_50_65_UP",
		503005: "AGE_RANGE_55_64",
		21:     "AGE_RANGE_55_65_UP",
		503006: "AGE_RANGE_65_UP",
	}
	ReachPlanAgeRangeEnum_ReachPlanAgeRange_value = map[string]int32{
		"UNSPECIFIED":        0,
		"UNKNOWN":            1,
		"AGE_RANGE_18_24":    503001,
		"AGE_RANGE_18_34":    2,
		"AGE_RANGE_18_44":    3,
		"AGE_RANGE_18_49":    4,
		"AGE_RANGE_18_54":    5,
		"AGE_RANGE_18_64":    6,
		"AGE_RANGE_18_65_UP": 7,
		"AGE_RANGE_21_34":    8,
		"AGE_RANGE_25_34":    503002,
		"AGE_RANGE_25_44":    9,
		"AGE_RANGE_25_49":    10,
		"AGE_RANGE_25_54":    11,
		"AGE_RANGE_25_64":    12,
		"AGE_RANGE_25_65_UP": 13,
		"AGE_RANGE_35_44":    503003,
		"AGE_RANGE_35_49":    14,
		"AGE_RANGE_35_54":    15,
		"AGE_RANGE_35_64":    16,
		"AGE_RANGE_35_65_UP": 17,
		"AGE_RANGE_45_54":    503004,
		"AGE_RANGE_45_64":    18,
		"AGE_RANGE_45_65_UP": 19,
		"AGE_RANGE_50_65_UP": 20,
		"AGE_RANGE_55_64":    503005,
		"AGE_RANGE_55_65_UP": 21,
		"AGE_RANGE_65_UP":    503006,
	}
)

func (x ReachPlanAgeRangeEnum_ReachPlanAgeRange) Enum() *ReachPlanAgeRangeEnum_ReachPlanAgeRange {
	p := new(ReachPlanAgeRangeEnum_ReachPlanAgeRange)
	*p = x
	return p
}

func (x ReachPlanAgeRangeEnum_ReachPlanAgeRange) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReachPlanAgeRangeEnum_ReachPlanAgeRange) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_enumTypes[0].Descriptor()
}

func (ReachPlanAgeRangeEnum_ReachPlanAgeRange) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_enumTypes[0]
}

func (x ReachPlanAgeRangeEnum_ReachPlanAgeRange) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReachPlanAgeRangeEnum_ReachPlanAgeRange.Descriptor instead.
func (ReachPlanAgeRangeEnum_ReachPlanAgeRange) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_rawDescGZIP(), []int{0, 0}
}

// Message describing plannable age ranges.
type ReachPlanAgeRangeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReachPlanAgeRangeEnum) Reset() {
	*x = ReachPlanAgeRangeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReachPlanAgeRangeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReachPlanAgeRangeEnum) ProtoMessage() {}

func (x *ReachPlanAgeRangeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReachPlanAgeRangeEnum.ProtoReflect.Descriptor instead.
func (*ReachPlanAgeRangeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v8_enums_reach_plan_age_range_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x72, 0x65, 0x61, 0x63, 0x68, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x61, 0x67, 0x65, 0x5f, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x05, 0x0a, 0x15, 0x52, 0x65, 0x61, 0x63,
	0x68, 0x50, 0x6c, 0x61, 0x6e, 0x41, 0x67, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x6e, 0x75,
	0x6d, 0x22, 0xf1, 0x04, 0x0a, 0x11, 0x52, 0x65, 0x61, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x6e, 0x41,
	0x67, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x0f, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x41, 0x4e,
	0x47, 0x45, 0x5f, 0x31, 0x38, 0x5f, 0x32, 0x34, 0x10, 0xd9, 0xd9, 0x1e, 0x12, 0x13, 0x0a, 0x0f,
	0x41, 0x47, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x31, 0x38, 0x5f, 0x33, 0x34, 0x10,
	0x02, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x31,
	0x38, 0x5f, 0x34, 0x34, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x41,
	0x4e, 0x47, 0x45, 0x5f, 0x31, 0x38, 0x5f, 0x34, 0x39, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x41,
	0x47, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x31, 0x38, 0x5f, 0x35, 0x34, 0x10, 0x05,
	0x12, 0x13, 0x0a, 0x0f, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x31, 0x38,
	0x5f, 0x36, 0x34, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x41, 0x4e,
	0x47, 0x45, 0x5f, 0x31, 0x38, 0x5f, 0x36, 0x35, 0x5f, 0x55, 0x50, 0x10, 0x07, 0x12, 0x13, 0x0a,
	0x0f, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x32, 0x31, 0x5f, 0x33, 0x34,
	0x10, 0x08, 0x12, 0x15, 0x0a, 0x0f, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f,
	0x32, 0x35, 0x5f, 0x33, 0x34, 0x10, 0xda, 0xd9, 0x1e, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x47, 0x45,
	0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x32, 0x35, 0x5f, 0x34, 0x34, 0x10, 0x09, 0x12, 0x13,
	0x0a, 0x0f, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x32, 0x35, 0x5f, 0x34,
	0x39, 0x10, 0x0a, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45,
	0x5f, 0x32, 0x35, 0x5f, 0x35, 0x34, 0x10, 0x0b, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x47, 0x45, 0x5f,
	0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x32, 0x35, 0x5f, 0x36, 0x34, 0x10, 0x0c, 0x12, 0x16, 0x0a,
	0x12, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x32, 0x35, 0x5f, 0x36, 0x35,
	0x5f, 0x55, 0x50, 0x10, 0x0d, 0x12, 0x15, 0x0a, 0x0f, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x41, 0x4e,
	0x47, 0x45, 0x5f, 0x33, 0x35, 0x5f, 0x34, 0x34, 0x10, 0xdb, 0xd9, 0x1e, 0x12, 0x13, 0x0a, 0x0f,
	0x41, 0x47, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x33, 0x35, 0x5f, 0x34, 0x39, 0x10,
	0x0e, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x33,
	0x35, 0x5f, 0x35, 0x34, 0x10, 0x0f, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x41,
	0x4e, 0x47, 0x45, 0x5f, 0x33, 0x35, 0x5f, 0x36, 0x34, 0x10, 0x10, 0x12, 0x16, 0x0a, 0x12, 0x41,
	0x47, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x33, 0x35, 0x5f, 0x36, 0x35, 0x5f, 0x55,
	0x50, 0x10, 0x11, 0x12, 0x15, 0x0a, 0x0f, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45,
	0x5f, 0x34, 0x35, 0x5f, 0x35, 0x34, 0x10, 0xdc, 0xd9, 0x1e, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x47,
	0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x34, 0x35, 0x5f, 0x36, 0x34, 0x10, 0x12, 0x12,
	0x16, 0x0a, 0x12, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x34, 0x35, 0x5f,
	0x36, 0x35, 0x5f, 0x55, 0x50, 0x10, 0x13, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x47, 0x45, 0x5f, 0x52,
	0x41, 0x4e, 0x47, 0x45, 0x5f, 0x35, 0x30, 0x5f, 0x36, 0x35, 0x5f, 0x55, 0x50, 0x10, 0x14, 0x12,
	0x15, 0x0a, 0x0f, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x35, 0x35, 0x5f,
	0x36, 0x34, 0x10, 0xdd, 0xd9, 0x1e, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x41,
	0x4e, 0x47, 0x45, 0x5f, 0x35, 0x35, 0x5f, 0x36, 0x35, 0x5f, 0x55, 0x50, 0x10, 0x15, 0x12, 0x15,
	0x0a, 0x0f, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x36, 0x35, 0x5f, 0x55,
	0x50, 0x10, 0xde, 0xd9, 0x1e, 0x42, 0xeb, 0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x16, 0x52, 0x65, 0x61,
	0x63, 0x68, 0x50, 0x6c, 0x61, 0x6e, 0x41, 0x67, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa,
	0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x38, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca,
	0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x38, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea,
	0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x38, 0x3a, 0x3a, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_rawDescData = file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_rawDesc
)

func file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_rawDescData)
	})
	return file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_rawDescData
}

var file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_goTypes = []interface{}{
	(ReachPlanAgeRangeEnum_ReachPlanAgeRange)(0), // 0: google.ads.googleads.v8.enums.ReachPlanAgeRangeEnum.ReachPlanAgeRange
	(*ReachPlanAgeRangeEnum)(nil),                // 1: google.ads.googleads.v8.enums.ReachPlanAgeRangeEnum
}
var file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_init() }
func file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_init() {
	if File_google_ads_googleads_v8_enums_reach_plan_age_range_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReachPlanAgeRangeEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v8_enums_reach_plan_age_range_proto = out.File
	file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_rawDesc = nil
	file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_goTypes = nil
	file_google_ads_googleads_v8_enums_reach_plan_age_range_proto_depIdxs = nil
}
