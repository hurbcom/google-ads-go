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
// source: google/ads/googleads/v8/enums/flight_placeholder_field.proto

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

// Possible values for Flight placeholder fields.
type FlightPlaceholderFieldEnum_FlightPlaceholderField int32

const (
	// Not specified.
	FlightPlaceholderFieldEnum_UNSPECIFIED FlightPlaceholderFieldEnum_FlightPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	FlightPlaceholderFieldEnum_UNKNOWN FlightPlaceholderFieldEnum_FlightPlaceholderField = 1
	// Data Type: STRING. Required. Destination id. Example: PAR, LON.
	// For feed items that only have destination id, destination id must be a
	// unique key. For feed items that have both destination id and origin id,
	// then the combination must be a unique key.
	FlightPlaceholderFieldEnum_DESTINATION_ID FlightPlaceholderFieldEnum_FlightPlaceholderField = 2
	// Data Type: STRING. Origin id. Example: PAR, LON.
	// Optional. Combination of destination id and origin id must be unique per
	// offer.
	FlightPlaceholderFieldEnum_ORIGIN_ID FlightPlaceholderFieldEnum_FlightPlaceholderField = 3
	// Data Type: STRING. Required. Main headline with product name to be shown
	// in dynamic ad.
	FlightPlaceholderFieldEnum_FLIGHT_DESCRIPTION FlightPlaceholderFieldEnum_FlightPlaceholderField = 4
	// Data Type: STRING. Shorter names are recommended.
	FlightPlaceholderFieldEnum_ORIGIN_NAME FlightPlaceholderFieldEnum_FlightPlaceholderField = 5
	// Data Type: STRING. Shorter names are recommended.
	FlightPlaceholderFieldEnum_DESTINATION_NAME FlightPlaceholderFieldEnum_FlightPlaceholderField = 6
	// Data Type: STRING. Price to be shown in the ad.
	// Example: "100.00 USD"
	FlightPlaceholderFieldEnum_FLIGHT_PRICE FlightPlaceholderFieldEnum_FlightPlaceholderField = 7
	// Data Type: STRING. Formatted price to be shown in the ad.
	// Example: "Starting at $100.00 USD", "$80 - $100"
	FlightPlaceholderFieldEnum_FORMATTED_PRICE FlightPlaceholderFieldEnum_FlightPlaceholderField = 8
	// Data Type: STRING. Sale price to be shown in the ad.
	// Example: "80.00 USD"
	FlightPlaceholderFieldEnum_FLIGHT_SALE_PRICE FlightPlaceholderFieldEnum_FlightPlaceholderField = 9
	// Data Type: STRING. Formatted sale price to be shown in the ad.
	// Example: "On sale for $80.00", "$60 - $80"
	FlightPlaceholderFieldEnum_FORMATTED_SALE_PRICE FlightPlaceholderFieldEnum_FlightPlaceholderField = 10
	// Data Type: URL. Image to be displayed in the ad.
	FlightPlaceholderFieldEnum_IMAGE_URL FlightPlaceholderFieldEnum_FlightPlaceholderField = 11
	// Data Type: URL_LIST. Required. Final URLs for the ad when using Upgraded
	// URLs. User will be redirected to these URLs when they click on an ad, or
	// when they click on a specific flight for ads that show multiple
	// flights.
	FlightPlaceholderFieldEnum_FINAL_URLS FlightPlaceholderFieldEnum_FlightPlaceholderField = 12
	// Data Type: URL_LIST. Final mobile URLs for the ad when using Upgraded
	// URLs.
	FlightPlaceholderFieldEnum_FINAL_MOBILE_URLS FlightPlaceholderFieldEnum_FlightPlaceholderField = 13
	// Data Type: URL. Tracking template for the ad when using Upgraded URLs.
	FlightPlaceholderFieldEnum_TRACKING_URL FlightPlaceholderFieldEnum_FlightPlaceholderField = 14
	// Data Type: STRING. Android app link. Must be formatted as:
	// android-app://{package_id}/{scheme}/{host_path}.
	// The components are defined as follows:
	// package_id: app ID as specified in Google Play.
	// scheme: the scheme to pass to the application. Can be HTTP, or a custom
	//   scheme.
	// host_path: identifies the specific content within your application.
	FlightPlaceholderFieldEnum_ANDROID_APP_LINK FlightPlaceholderFieldEnum_FlightPlaceholderField = 15
	// Data Type: STRING_LIST. List of recommended destination IDs to show
	// together with this item.
	FlightPlaceholderFieldEnum_SIMILAR_DESTINATION_IDS FlightPlaceholderFieldEnum_FlightPlaceholderField = 16
	// Data Type: STRING. iOS app link.
	FlightPlaceholderFieldEnum_IOS_APP_LINK FlightPlaceholderFieldEnum_FlightPlaceholderField = 17
	// Data Type: INT64. iOS app store ID.
	FlightPlaceholderFieldEnum_IOS_APP_STORE_ID FlightPlaceholderFieldEnum_FlightPlaceholderField = 18
)

// Enum value maps for FlightPlaceholderFieldEnum_FlightPlaceholderField.
var (
	FlightPlaceholderFieldEnum_FlightPlaceholderField_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "DESTINATION_ID",
		3:  "ORIGIN_ID",
		4:  "FLIGHT_DESCRIPTION",
		5:  "ORIGIN_NAME",
		6:  "DESTINATION_NAME",
		7:  "FLIGHT_PRICE",
		8:  "FORMATTED_PRICE",
		9:  "FLIGHT_SALE_PRICE",
		10: "FORMATTED_SALE_PRICE",
		11: "IMAGE_URL",
		12: "FINAL_URLS",
		13: "FINAL_MOBILE_URLS",
		14: "TRACKING_URL",
		15: "ANDROID_APP_LINK",
		16: "SIMILAR_DESTINATION_IDS",
		17: "IOS_APP_LINK",
		18: "IOS_APP_STORE_ID",
	}
	FlightPlaceholderFieldEnum_FlightPlaceholderField_value = map[string]int32{
		"UNSPECIFIED":             0,
		"UNKNOWN":                 1,
		"DESTINATION_ID":          2,
		"ORIGIN_ID":               3,
		"FLIGHT_DESCRIPTION":      4,
		"ORIGIN_NAME":             5,
		"DESTINATION_NAME":        6,
		"FLIGHT_PRICE":            7,
		"FORMATTED_PRICE":         8,
		"FLIGHT_SALE_PRICE":       9,
		"FORMATTED_SALE_PRICE":    10,
		"IMAGE_URL":               11,
		"FINAL_URLS":              12,
		"FINAL_MOBILE_URLS":       13,
		"TRACKING_URL":            14,
		"ANDROID_APP_LINK":        15,
		"SIMILAR_DESTINATION_IDS": 16,
		"IOS_APP_LINK":            17,
		"IOS_APP_STORE_ID":        18,
	}
)

func (x FlightPlaceholderFieldEnum_FlightPlaceholderField) Enum() *FlightPlaceholderFieldEnum_FlightPlaceholderField {
	p := new(FlightPlaceholderFieldEnum_FlightPlaceholderField)
	*p = x
	return p
}

func (x FlightPlaceholderFieldEnum_FlightPlaceholderField) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FlightPlaceholderFieldEnum_FlightPlaceholderField) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_enumTypes[0].Descriptor()
}

func (FlightPlaceholderFieldEnum_FlightPlaceholderField) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_enumTypes[0]
}

func (x FlightPlaceholderFieldEnum_FlightPlaceholderField) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FlightPlaceholderFieldEnum_FlightPlaceholderField.Descriptor instead.
func (FlightPlaceholderFieldEnum_FlightPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_rawDescGZIP(), []int{0, 0}
}

// Values for Flight placeholder fields.
// For more information about dynamic remarketing feeds, see
// https://support.google.com/google-ads/answer/6053288.
type FlightPlaceholderFieldEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FlightPlaceholderFieldEnum) Reset() {
	*x = FlightPlaceholderFieldEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlightPlaceholderFieldEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlightPlaceholderFieldEnum) ProtoMessage() {}

func (x *FlightPlaceholderFieldEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlightPlaceholderFieldEnum.ProtoReflect.Descriptor instead.
func (*FlightPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v8_enums_flight_placeholder_field_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x03, 0x0a, 0x1a,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x93, 0x03, 0x0a, 0x16, 0x46,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x4f, 0x52, 0x49, 0x47, 0x49,
	0x4e, 0x5f, 0x49, 0x44, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x4c, 0x49, 0x47, 0x48, 0x54,
	0x5f, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x0f,
	0x0a, 0x0b, 0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x05, 0x12,
	0x14, 0x0a, 0x10, 0x44, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e,
	0x41, 0x4d, 0x45, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x46, 0x4c, 0x49, 0x47, 0x48, 0x54, 0x5f,
	0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x4f, 0x52, 0x4d, 0x41,
	0x54, 0x54, 0x45, 0x44, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x08, 0x12, 0x15, 0x0a, 0x11,
	0x46, 0x4c, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x53, 0x41, 0x4c, 0x45, 0x5f, 0x50, 0x52, 0x49, 0x43,
	0x45, 0x10, 0x09, 0x12, 0x18, 0x0a, 0x14, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x54, 0x45, 0x44,
	0x5f, 0x53, 0x41, 0x4c, 0x45, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x0a, 0x12, 0x0d, 0x0a,
	0x09, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x0b, 0x12, 0x0e, 0x0a, 0x0a,
	0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0x0c, 0x12, 0x15, 0x0a, 0x11,
	0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x55, 0x52, 0x4c,
	0x53, 0x10, 0x0d, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f,
	0x55, 0x52, 0x4c, 0x10, 0x0e, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44,
	0x5f, 0x41, 0x50, 0x50, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x0f, 0x12, 0x1b, 0x0a, 0x17, 0x53,
	0x49, 0x4d, 0x49, 0x4c, 0x41, 0x52, 0x5f, 0x44, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x53, 0x10, 0x10, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4f, 0x53, 0x5f,
	0x41, 0x50, 0x50, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x11, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x4f,
	0x53, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x12,
	0x42, 0xf1, 0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x1c, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41,
	0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x38, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0xca, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x38, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x38, 0x3a, 0x3a, 0x45,
	0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_rawDescData = file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_rawDesc
)

func file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_rawDescData)
	})
	return file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_rawDescData
}

var file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_goTypes = []interface{}{
	(FlightPlaceholderFieldEnum_FlightPlaceholderField)(0), // 0: google.ads.googleads.v8.enums.FlightPlaceholderFieldEnum.FlightPlaceholderField
	(*FlightPlaceholderFieldEnum)(nil),                     // 1: google.ads.googleads.v8.enums.FlightPlaceholderFieldEnum
}
var file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_init() }
func file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_init() {
	if File_google_ads_googleads_v8_enums_flight_placeholder_field_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlightPlaceholderFieldEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v8_enums_flight_placeholder_field_proto = out.File
	file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_rawDesc = nil
	file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_goTypes = nil
	file_google_ads_googleads_v8_enums_flight_placeholder_field_proto_depIdxs = nil
}
