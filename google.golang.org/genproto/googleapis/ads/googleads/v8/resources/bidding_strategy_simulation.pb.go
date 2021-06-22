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
// source: google/ads/googleads/v8/resources/bidding_strategy_simulation.proto

package resources

import (
	common "google.golang.org/genproto/googleapis/ads/googleads/v8/common"
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

// A bidding strategy simulation. Supported combinations of simulation type
// and simulation modification method are detailed below respectively.
//
// 1. TARGET_CPA - UNIFORM
// 2. TARGET_ROAS - UNIFORM
type BiddingStrategySimulation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the bidding strategy simulation.
	// Bidding strategy simulation resource names have the form:
	//
	// `customers/{customer_id}/biddingStrategySimulations/{bidding_strategy_id}~{type}~{modification_method}~{start_date}~{end_date}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Bidding strategy shared set id of the simulation.
	BiddingStrategyId int64 `protobuf:"varint,2,opt,name=bidding_strategy_id,json=biddingStrategyId,proto3" json:"bidding_strategy_id,omitempty"`
	// Output only. The field that the simulation modifies.
	Type enums.SimulationTypeEnum_SimulationType `protobuf:"varint,3,opt,name=type,proto3,enum=google.ads.googleads.v8.enums.SimulationTypeEnum_SimulationType" json:"type,omitempty"`
	// Output only. How the simulation modifies the field.
	ModificationMethod enums.SimulationModificationMethodEnum_SimulationModificationMethod `protobuf:"varint,4,opt,name=modification_method,json=modificationMethod,proto3,enum=google.ads.googleads.v8.enums.SimulationModificationMethodEnum_SimulationModificationMethod" json:"modification_method,omitempty"`
	// Output only. First day on which the simulation is based, in YYYY-MM-DD format.
	StartDate string `protobuf:"bytes,5,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// Output only. Last day on which the simulation is based, in YYYY-MM-DD format
	EndDate string `protobuf:"bytes,6,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	// List of simulation points.
	//
	// Types that are assignable to PointList:
	//	*BiddingStrategySimulation_TargetCpaPointList
	//	*BiddingStrategySimulation_TargetRoasPointList
	PointList isBiddingStrategySimulation_PointList `protobuf_oneof:"point_list"`
}

func (x *BiddingStrategySimulation) Reset() {
	*x = BiddingStrategySimulation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BiddingStrategySimulation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BiddingStrategySimulation) ProtoMessage() {}

func (x *BiddingStrategySimulation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BiddingStrategySimulation.ProtoReflect.Descriptor instead.
func (*BiddingStrategySimulation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_rawDescGZIP(), []int{0}
}

func (x *BiddingStrategySimulation) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *BiddingStrategySimulation) GetBiddingStrategyId() int64 {
	if x != nil {
		return x.BiddingStrategyId
	}
	return 0
}

func (x *BiddingStrategySimulation) GetType() enums.SimulationTypeEnum_SimulationType {
	if x != nil {
		return x.Type
	}
	return enums.SimulationTypeEnum_UNSPECIFIED
}

func (x *BiddingStrategySimulation) GetModificationMethod() enums.SimulationModificationMethodEnum_SimulationModificationMethod {
	if x != nil {
		return x.ModificationMethod
	}
	return enums.SimulationModificationMethodEnum_UNSPECIFIED
}

func (x *BiddingStrategySimulation) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *BiddingStrategySimulation) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (m *BiddingStrategySimulation) GetPointList() isBiddingStrategySimulation_PointList {
	if m != nil {
		return m.PointList
	}
	return nil
}

func (x *BiddingStrategySimulation) GetTargetCpaPointList() *common.TargetCpaSimulationPointList {
	if x, ok := x.GetPointList().(*BiddingStrategySimulation_TargetCpaPointList); ok {
		return x.TargetCpaPointList
	}
	return nil
}

func (x *BiddingStrategySimulation) GetTargetRoasPointList() *common.TargetRoasSimulationPointList {
	if x, ok := x.GetPointList().(*BiddingStrategySimulation_TargetRoasPointList); ok {
		return x.TargetRoasPointList
	}
	return nil
}

type isBiddingStrategySimulation_PointList interface {
	isBiddingStrategySimulation_PointList()
}

type BiddingStrategySimulation_TargetCpaPointList struct {
	// Output only. Simulation points if the simulation type is TARGET_CPA.
	TargetCpaPointList *common.TargetCpaSimulationPointList `protobuf:"bytes,7,opt,name=target_cpa_point_list,json=targetCpaPointList,proto3,oneof"`
}

type BiddingStrategySimulation_TargetRoasPointList struct {
	// Output only. Simulation points if the simulation type is TARGET_ROAS.
	TargetRoasPointList *common.TargetRoasSimulationPointList `protobuf:"bytes,8,opt,name=target_roas_point_list,json=targetRoasPointList,proto3,oneof"`
}

func (*BiddingStrategySimulation_TargetCpaPointList) isBiddingStrategySimulation_PointList() {}

func (*BiddingStrategySimulation_TargetRoasPointList) isBiddingStrategySimulation_PointList() {}

var File_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x38, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x73, 0x69, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x07, 0x0a,
	0x19, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5f, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x3a, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x34, 0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x13, 0x62,
	0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x11, 0x62,
	0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x64,
	0x12, 0x59, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x40,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x53,
	0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75,
	0x6d, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x92, 0x01, 0x0a, 0x13,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x12, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x22, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x76, 0x0a, 0x15, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x63,
	0x70, 0x61, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x70, 0x61, 0x53, 0x69,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x12, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x43, 0x70, 0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x79, 0x0a, 0x16,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x6f, 0x61, 0x73, 0x5f, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x52, 0x6f, 0x61, 0x73, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x48, 0x00, 0x52, 0x13, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x6f, 0x61, 0x73, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x3a, 0xb7, 0x01, 0xea, 0x41, 0xb3, 0x01, 0x0a, 0x32,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x7d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x62, 0x69, 0x64, 0x64,
	0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x53, 0x69, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x7e, 0x7b, 0x74, 0x79,
	0x70, 0x65, 0x7d, 0x7e, 0x7b, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x7d, 0x7e, 0x7b, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x7d, 0x7e, 0x7b, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x7d, 0x42, 0x0c, 0x0a, 0x0a, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x42,
	0x8b, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x1e, 0x42, 0x69, 0x64, 0x64, 0x69,
	0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x38, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x38, 0x5c, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x56, 0x38, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_rawDescData = file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_rawDesc
)

func file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_rawDescData)
	})
	return file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_rawDescData
}

var file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_goTypes = []interface{}{
	(*BiddingStrategySimulation)(nil),                                        // 0: google.ads.googleads.v8.resources.BiddingStrategySimulation
	(enums.SimulationTypeEnum_SimulationType)(0),                             // 1: google.ads.googleads.v8.enums.SimulationTypeEnum.SimulationType
	(enums.SimulationModificationMethodEnum_SimulationModificationMethod)(0), // 2: google.ads.googleads.v8.enums.SimulationModificationMethodEnum.SimulationModificationMethod
	(*common.TargetCpaSimulationPointList)(nil),                              // 3: google.ads.googleads.v8.common.TargetCpaSimulationPointList
	(*common.TargetRoasSimulationPointList)(nil),                             // 4: google.ads.googleads.v8.common.TargetRoasSimulationPointList
}
var file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v8.resources.BiddingStrategySimulation.type:type_name -> google.ads.googleads.v8.enums.SimulationTypeEnum.SimulationType
	2, // 1: google.ads.googleads.v8.resources.BiddingStrategySimulation.modification_method:type_name -> google.ads.googleads.v8.enums.SimulationModificationMethodEnum.SimulationModificationMethod
	3, // 2: google.ads.googleads.v8.resources.BiddingStrategySimulation.target_cpa_point_list:type_name -> google.ads.googleads.v8.common.TargetCpaSimulationPointList
	4, // 3: google.ads.googleads.v8.resources.BiddingStrategySimulation.target_roas_point_list:type_name -> google.ads.googleads.v8.common.TargetRoasSimulationPointList
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_init() }
func file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_init() {
	if File_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BiddingStrategySimulation); i {
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
	file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*BiddingStrategySimulation_TargetCpaPointList)(nil),
		(*BiddingStrategySimulation_TargetRoasPointList)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto = out.File
	file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_rawDesc = nil
	file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_goTypes = nil
	file_google_ads_googleads_v8_resources_bidding_strategy_simulation_proto_depIdxs = nil
}
