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
// source: google/ads/googleads/v8/resources/batch_job.proto

package resources

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

// A list of mutates being processed asynchronously. The mutates are uploaded
// by the user. The mutates themselves aren't readable and the results of the
// job can only be read using BatchJobService.ListBatchJobResults.
type BatchJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the batch job.
	// Batch job resource names have the form:
	//
	// `customers/{customer_id}/batchJobs/{batch_job_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. ID of this batch job.
	Id *int64 `protobuf:"varint,7,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// Output only. The next sequence token to use when adding operations. Only set when the
	// batch job status is PENDING.
	NextAddSequenceToken *string `protobuf:"bytes,8,opt,name=next_add_sequence_token,json=nextAddSequenceToken,proto3,oneof" json:"next_add_sequence_token,omitempty"`
	// Output only. Contains additional information about this batch job.
	Metadata *BatchJob_BatchJobMetadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Output only. Status of this batch job.
	Status enums.BatchJobStatusEnum_BatchJobStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v8.enums.BatchJobStatusEnum_BatchJobStatus" json:"status,omitempty"`
	// Output only. The resource name of the long-running operation that can be used to poll
	// for completion. Only set when the batch job status is RUNNING or DONE.
	LongRunningOperation *string `protobuf:"bytes,9,opt,name=long_running_operation,json=longRunningOperation,proto3,oneof" json:"long_running_operation,omitempty"`
}

func (x *BatchJob) Reset() {
	*x = BatchJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_resources_batch_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchJob) ProtoMessage() {}

func (x *BatchJob) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_resources_batch_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchJob.ProtoReflect.Descriptor instead.
func (*BatchJob) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_resources_batch_job_proto_rawDescGZIP(), []int{0}
}

func (x *BatchJob) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *BatchJob) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *BatchJob) GetNextAddSequenceToken() string {
	if x != nil && x.NextAddSequenceToken != nil {
		return *x.NextAddSequenceToken
	}
	return ""
}

func (x *BatchJob) GetMetadata() *BatchJob_BatchJobMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *BatchJob) GetStatus() enums.BatchJobStatusEnum_BatchJobStatus {
	if x != nil {
		return x.Status
	}
	return enums.BatchJobStatusEnum_UNSPECIFIED
}

func (x *BatchJob) GetLongRunningOperation() string {
	if x != nil && x.LongRunningOperation != nil {
		return *x.LongRunningOperation
	}
	return ""
}

// Additional information about the batch job. This message is also used as
// metadata returned in batch job Long Running Operations.
type BatchJob_BatchJobMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The time when this batch job was created.
	// Formatted as yyyy-mm-dd hh:mm:ss. Example: "2018-03-05 09:15:00"
	CreationDateTime *string `protobuf:"bytes,8,opt,name=creation_date_time,json=creationDateTime,proto3,oneof" json:"creation_date_time,omitempty"`
	// Output only. The time when this batch job started running.
	// Formatted as yyyy-mm-dd hh:mm:ss. Example: "2018-03-05 09:15:30"
	StartDateTime *string `protobuf:"bytes,7,opt,name=start_date_time,json=startDateTime,proto3,oneof" json:"start_date_time,omitempty"`
	// Output only. The time when this batch job was completed.
	// Formatted as yyyy-MM-dd HH:mm:ss. Example: "2018-03-05 09:16:00"
	CompletionDateTime *string `protobuf:"bytes,9,opt,name=completion_date_time,json=completionDateTime,proto3,oneof" json:"completion_date_time,omitempty"`
	// Output only. The fraction (between 0.0 and 1.0) of mutates that have been processed.
	// This is empty if the job hasn't started running yet.
	EstimatedCompletionRatio *float64 `protobuf:"fixed64,10,opt,name=estimated_completion_ratio,json=estimatedCompletionRatio,proto3,oneof" json:"estimated_completion_ratio,omitempty"`
	// Output only. The number of mutate operations in the batch job.
	OperationCount *int64 `protobuf:"varint,11,opt,name=operation_count,json=operationCount,proto3,oneof" json:"operation_count,omitempty"`
	// Output only. The number of mutate operations executed by the batch job.
	// Present only if the job has started running.
	ExecutedOperationCount *int64 `protobuf:"varint,12,opt,name=executed_operation_count,json=executedOperationCount,proto3,oneof" json:"executed_operation_count,omitempty"`
}

func (x *BatchJob_BatchJobMetadata) Reset() {
	*x = BatchJob_BatchJobMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_resources_batch_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchJob_BatchJobMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchJob_BatchJobMetadata) ProtoMessage() {}

func (x *BatchJob_BatchJobMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_resources_batch_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchJob_BatchJobMetadata.ProtoReflect.Descriptor instead.
func (*BatchJob_BatchJobMetadata) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_resources_batch_job_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BatchJob_BatchJobMetadata) GetCreationDateTime() string {
	if x != nil && x.CreationDateTime != nil {
		return *x.CreationDateTime
	}
	return ""
}

func (x *BatchJob_BatchJobMetadata) GetStartDateTime() string {
	if x != nil && x.StartDateTime != nil {
		return *x.StartDateTime
	}
	return ""
}

func (x *BatchJob_BatchJobMetadata) GetCompletionDateTime() string {
	if x != nil && x.CompletionDateTime != nil {
		return *x.CompletionDateTime
	}
	return ""
}

func (x *BatchJob_BatchJobMetadata) GetEstimatedCompletionRatio() float64 {
	if x != nil && x.EstimatedCompletionRatio != nil {
		return *x.EstimatedCompletionRatio
	}
	return 0
}

func (x *BatchJob_BatchJobMetadata) GetOperationCount() int64 {
	if x != nil && x.OperationCount != nil {
		return *x.OperationCount
	}
	return 0
}

func (x *BatchJob_BatchJobMetadata) GetExecutedOperationCount() int64 {
	if x != nil && x.ExecutedOperationCount != nil {
		return *x.ExecutedOperationCount
	}
	return 0
}

var File_google_ads_googleads_v8_resources_batch_job_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v8_resources_batch_job_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6a, 0x6f, 0x62, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x08, 0x0a, 0x08, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x4a, 0x6f, 0x62, 0x12, 0x4e, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x05, 0xfa,
	0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a,
	0x17, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x48, 0x01, 0x52, 0x14, 0x6e, 0x65, 0x78, 0x74, 0x41, 0x64, 0x64, 0x53, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x5d,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x2e, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x5d, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x40, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d,
	0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3e, 0x0a, 0x16,
	0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x48, 0x02, 0x52, 0x14, 0x6c, 0x6f, 0x6e, 0x67, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x1a, 0x8b, 0x04, 0x0a,
	0x10, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x36, 0x0a, 0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x48, 0x00, 0x52, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x0f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x01, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x14, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x02,
	0x52, 0x12, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x46, 0x0a, 0x1a, 0x65, 0x73, 0x74, 0x69, 0x6d,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x48, 0x03, 0x52, 0x18, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x88, 0x01, 0x01, 0x12,
	0x31, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x04, 0x52,
	0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x42, 0x0a, 0x18, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x05, 0x52, 0x16, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x12, 0x0a,
	0x10, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x1d, 0x0a, 0x1b, 0x5f, 0x65,
	0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x1b, 0x0a,
	0x19, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x58, 0xea, 0x41, 0x55, 0x0a,
	0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a,
	0x6f, 0x62, 0x12, 0x30, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x4a, 0x6f, 0x62, 0x73, 0x2f, 0x7b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6a, 0x6f, 0x62,
	0x5f, 0x69, 0x64, 0x7d, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x1a, 0x0a, 0x18, 0x5f,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x6c, 0x6f, 0x6e, 0x67,
	0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0xfa, 0x01, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x0d, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa,
	0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x38, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73,
	0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x38, 0x5c, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x56, 0x38, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v8_resources_batch_job_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v8_resources_batch_job_proto_rawDescData = file_google_ads_googleads_v8_resources_batch_job_proto_rawDesc
)

func file_google_ads_googleads_v8_resources_batch_job_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v8_resources_batch_job_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v8_resources_batch_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v8_resources_batch_job_proto_rawDescData)
	})
	return file_google_ads_googleads_v8_resources_batch_job_proto_rawDescData
}

var file_google_ads_googleads_v8_resources_batch_job_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_ads_googleads_v8_resources_batch_job_proto_goTypes = []interface{}{
	(*BatchJob)(nil),                             // 0: google.ads.googleads.v8.resources.BatchJob
	(*BatchJob_BatchJobMetadata)(nil),            // 1: google.ads.googleads.v8.resources.BatchJob.BatchJobMetadata
	(enums.BatchJobStatusEnum_BatchJobStatus)(0), // 2: google.ads.googleads.v8.enums.BatchJobStatusEnum.BatchJobStatus
}
var file_google_ads_googleads_v8_resources_batch_job_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v8.resources.BatchJob.metadata:type_name -> google.ads.googleads.v8.resources.BatchJob.BatchJobMetadata
	2, // 1: google.ads.googleads.v8.resources.BatchJob.status:type_name -> google.ads.googleads.v8.enums.BatchJobStatusEnum.BatchJobStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v8_resources_batch_job_proto_init() }
func file_google_ads_googleads_v8_resources_batch_job_proto_init() {
	if File_google_ads_googleads_v8_resources_batch_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v8_resources_batch_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchJob); i {
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
		file_google_ads_googleads_v8_resources_batch_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchJob_BatchJobMetadata); i {
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
	file_google_ads_googleads_v8_resources_batch_job_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_google_ads_googleads_v8_resources_batch_job_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v8_resources_batch_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v8_resources_batch_job_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v8_resources_batch_job_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v8_resources_batch_job_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v8_resources_batch_job_proto = out.File
	file_google_ads_googleads_v8_resources_batch_job_proto_rawDesc = nil
	file_google_ads_googleads_v8_resources_batch_job_proto_goTypes = nil
	file_google_ads_googleads_v8_resources_batch_job_proto_depIdxs = nil
}
