// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/artifacts_store.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Artifacts store access specification
type ArtifactsStore struct {
	// UUID
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Optional description
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Optional Tags of this entity
	Tags []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	// Optional time when the entityt was created
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Optional last time the entity was updated
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Optional a flag to freeze the store
	Frozen bool `protobuf:"varint,7,opt,name=frozen,proto3" json:"frozen,omitempty"`
	// Optional a flag to disable the store
	Disabled bool `protobuf:"varint,8,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Optional if the entity has been deleted
	Deleted bool `protobuf:"varint,9,opt,name=deleted,proto3" json:"deleted,omitempty"`
	// Optional the k8s secret to use
	K8SSecret string `protobuf:"bytes,10,opt,name=k8s_secret,json=k8sSecret,proto3" json:"k8s_secret,omitempty"`
	// Optional type of the store
	Type string `protobuf:"bytes,11,opt,name=type,proto3" json:"type,omitempty"`
	// Optional mounth path
	MountPath string `protobuf:"bytes,12,opt,name=mount_path,json=mountPath,proto3" json:"mount_path,omitempty"`
	// Optional host path
	HostPath string `protobuf:"bytes,13,opt,name=host_path,json=hostPath,proto3" json:"host_path,omitempty"`
	// Optional volume claim
	VolumeClaim string `protobuf:"bytes,14,opt,name=volume_claim,json=volumeClaim,proto3" json:"volume_claim,omitempty"`
	// Optional bucket
	Bucket string `protobuf:"bytes,15,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// Optional flag to set this store to read only mode
	ReadOnly             bool     `protobuf:"varint,16,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArtifactsStore) Reset()         { *m = ArtifactsStore{} }
func (m *ArtifactsStore) String() string { return proto.CompactTextString(m) }
func (*ArtifactsStore) ProtoMessage()    {}
func (*ArtifactsStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_53cb0be1952660cd, []int{0}
}

func (m *ArtifactsStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactsStore.Unmarshal(m, b)
}
func (m *ArtifactsStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactsStore.Marshal(b, m, deterministic)
}
func (m *ArtifactsStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactsStore.Merge(m, src)
}
func (m *ArtifactsStore) XXX_Size() int {
	return xxx_messageInfo_ArtifactsStore.Size(m)
}
func (m *ArtifactsStore) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactsStore.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactsStore proto.InternalMessageInfo

func (m *ArtifactsStore) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ArtifactsStore) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArtifactsStore) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ArtifactsStore) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ArtifactsStore) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *ArtifactsStore) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *ArtifactsStore) GetFrozen() bool {
	if m != nil {
		return m.Frozen
	}
	return false
}

func (m *ArtifactsStore) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *ArtifactsStore) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *ArtifactsStore) GetK8SSecret() string {
	if m != nil {
		return m.K8SSecret
	}
	return ""
}

func (m *ArtifactsStore) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ArtifactsStore) GetMountPath() string {
	if m != nil {
		return m.MountPath
	}
	return ""
}

func (m *ArtifactsStore) GetHostPath() string {
	if m != nil {
		return m.HostPath
	}
	return ""
}

func (m *ArtifactsStore) GetVolumeClaim() string {
	if m != nil {
		return m.VolumeClaim
	}
	return ""
}

func (m *ArtifactsStore) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *ArtifactsStore) GetReadOnly() bool {
	if m != nil {
		return m.ReadOnly
	}
	return false
}

// Request data to create/update artifacts store
type ArtifactsStoreBodyRequest struct {
	// Owner of the namespace
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// Artifact store body
	ArtifactStore        *ArtifactsStore `protobuf:"bytes,2,opt,name=artifact_store,json=artifactStore,proto3" json:"artifact_store,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ArtifactsStoreBodyRequest) Reset()         { *m = ArtifactsStoreBodyRequest{} }
func (m *ArtifactsStoreBodyRequest) String() string { return proto.CompactTextString(m) }
func (*ArtifactsStoreBodyRequest) ProtoMessage()    {}
func (*ArtifactsStoreBodyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_53cb0be1952660cd, []int{1}
}

func (m *ArtifactsStoreBodyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactsStoreBodyRequest.Unmarshal(m, b)
}
func (m *ArtifactsStoreBodyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactsStoreBodyRequest.Marshal(b, m, deterministic)
}
func (m *ArtifactsStoreBodyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactsStoreBodyRequest.Merge(m, src)
}
func (m *ArtifactsStoreBodyRequest) XXX_Size() int {
	return xxx_messageInfo_ArtifactsStoreBodyRequest.Size(m)
}
func (m *ArtifactsStoreBodyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactsStoreBodyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactsStoreBodyRequest proto.InternalMessageInfo

func (m *ArtifactsStoreBodyRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ArtifactsStoreBodyRequest) GetArtifactStore() *ArtifactsStore {
	if m != nil {
		return m.ArtifactStore
	}
	return nil
}

// Contains list artifacts stores
type ListArtifactsStoresResponse struct {
	// Count of the entities
	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// List of all entities
	Results []*ArtifactsStore `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	// Previous page
	Previous string `protobuf:"bytes,3,opt,name=previous,proto3" json:"previous,omitempty"`
	// Next page
	Next                 string   `protobuf:"bytes,4,opt,name=next,proto3" json:"next,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListArtifactsStoresResponse) Reset()         { *m = ListArtifactsStoresResponse{} }
func (m *ListArtifactsStoresResponse) String() string { return proto.CompactTextString(m) }
func (*ListArtifactsStoresResponse) ProtoMessage()    {}
func (*ListArtifactsStoresResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_53cb0be1952660cd, []int{2}
}

func (m *ListArtifactsStoresResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListArtifactsStoresResponse.Unmarshal(m, b)
}
func (m *ListArtifactsStoresResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListArtifactsStoresResponse.Marshal(b, m, deterministic)
}
func (m *ListArtifactsStoresResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListArtifactsStoresResponse.Merge(m, src)
}
func (m *ListArtifactsStoresResponse) XXX_Size() int {
	return xxx_messageInfo_ListArtifactsStoresResponse.Size(m)
}
func (m *ListArtifactsStoresResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListArtifactsStoresResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListArtifactsStoresResponse proto.InternalMessageInfo

func (m *ListArtifactsStoresResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListArtifactsStoresResponse) GetResults() []*ArtifactsStore {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *ListArtifactsStoresResponse) GetPrevious() string {
	if m != nil {
		return m.Previous
	}
	return ""
}

func (m *ListArtifactsStoresResponse) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

// Artifact tree response
type ArtifactTreeResponse struct {
	Files                map[string]int64 `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Dirs                 []string         `protobuf:"bytes,2,rep,name=dirs,proto3" json:"dirs,omitempty"`
	IsDone               bool             `protobuf:"varint,3,opt,name=is_done,json=isDone,proto3" json:"is_done,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ArtifactTreeResponse) Reset()         { *m = ArtifactTreeResponse{} }
func (m *ArtifactTreeResponse) String() string { return proto.CompactTextString(m) }
func (*ArtifactTreeResponse) ProtoMessage()    {}
func (*ArtifactTreeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_53cb0be1952660cd, []int{3}
}

func (m *ArtifactTreeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactTreeResponse.Unmarshal(m, b)
}
func (m *ArtifactTreeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactTreeResponse.Marshal(b, m, deterministic)
}
func (m *ArtifactTreeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactTreeResponse.Merge(m, src)
}
func (m *ArtifactTreeResponse) XXX_Size() int {
	return xxx_messageInfo_ArtifactTreeResponse.Size(m)
}
func (m *ArtifactTreeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactTreeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactTreeResponse proto.InternalMessageInfo

func (m *ArtifactTreeResponse) GetFiles() map[string]int64 {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *ArtifactTreeResponse) GetDirs() []string {
	if m != nil {
		return m.Dirs
	}
	return nil
}

func (m *ArtifactTreeResponse) GetIsDone() bool {
	if m != nil {
		return m.IsDone
	}
	return false
}

func init() {
	proto.RegisterType((*ArtifactsStore)(nil), "v1.ArtifactsStore")
	proto.RegisterType((*ArtifactsStoreBodyRequest)(nil), "v1.ArtifactsStoreBodyRequest")
	proto.RegisterType((*ListArtifactsStoresResponse)(nil), "v1.ListArtifactsStoresResponse")
	proto.RegisterType((*ArtifactTreeResponse)(nil), "v1.ArtifactTreeResponse")
	proto.RegisterMapType((map[string]int64)(nil), "v1.ArtifactTreeResponse.FilesEntry")
}

func init() { proto.RegisterFile("v1/artifacts_store.proto", fileDescriptor_53cb0be1952660cd) }

var fileDescriptor_53cb0be1952660cd = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcb, 0x6e, 0xdb, 0x3a,
	0x10, 0x85, 0xec, 0xf8, 0x35, 0x4e, 0x72, 0x03, 0x22, 0xb8, 0x65, 0x1d, 0x14, 0x75, 0xdd, 0x4d,
	0x16, 0x85, 0x82, 0xa4, 0x1b, 0xa7, 0xbb, 0xf4, 0xb5, 0x2a, 0xd0, 0x42, 0xc9, 0x5e, 0xa0, 0xc5,
	0x71, 0x42, 0x58, 0x16, 0x55, 0x72, 0xa4, 0x46, 0xfd, 0x8b, 0xfe, 0x4d, 0x7f, 0xad, 0xbb, 0x82,
	0xa4, 0x94, 0xc6, 0x8b, 0xa2, 0xbb, 0x39, 0x67, 0xe6, 0x70, 0x1e, 0x3c, 0xc0, 0xeb, 0xf3, 0x33,
	0x61, 0x48, 0xad, 0x45, 0x46, 0x36, 0xb5, 0xa4, 0x0d, 0xc6, 0xa5, 0xd1, 0xa4, 0x59, 0xaf, 0x3e,
	0x9f, 0x3d, 0xbf, 0xd5, 0xfa, 0x36, 0xc7, 0x33, 0xcf, 0xac, 0xaa, 0xf5, 0x19, 0xa9, 0x2d, 0x5a,
	0x12, 0xdb, 0x32, 0x14, 0x2d, 0x7e, 0xf5, 0xe1, 0xf0, 0xaa, 0x93, 0x5f, 0x3b, 0x35, 0x63, 0xb0,
	0x57, 0x55, 0x4a, 0xf2, 0x68, 0x1e, 0x9d, 0x4e, 0x12, 0x1f, 0x3b, 0xae, 0x10, 0x5b, 0xe4, 0xbd,
	0xc0, 0xb9, 0x98, 0xcd, 0x61, 0x2a, 0xd1, 0x66, 0x46, 0x95, 0xa4, 0x74, 0xc1, 0xfb, 0x3e, 0xf5,
	0x98, 0x72, 0x2a, 0x12, 0xb7, 0x96, 0xef, 0xcd, 0xfb, 0x4e, 0xe5, 0x62, 0x76, 0x09, 0x90, 0x19,
	0x14, 0x84, 0x32, 0x15, 0xc4, 0x07, 0xf3, 0xe8, 0x74, 0x7a, 0x31, 0x8b, 0xc3, 0x98, 0x71, 0x37,
	0x66, 0x7c, 0xd3, 0x8d, 0x99, 0x4c, 0xda, 0xea, 0x2b, 0x72, 0xd2, 0xaa, 0x94, 0x9d, 0x74, 0xf8,
	0x6f, 0x69, 0x5b, 0x7d, 0x45, 0xec, 0x7f, 0x18, 0xae, 0x8d, 0xfe, 0x8e, 0x05, 0x1f, 0xcd, 0xa3,
	0xd3, 0x71, 0xd2, 0x22, 0x36, 0x83, 0xb1, 0x54, 0x56, 0xac, 0x72, 0x94, 0x7c, 0xec, 0x33, 0x0f,
	0x98, 0x71, 0x18, 0x49, 0xcc, 0x91, 0x50, 0xf2, 0x89, 0x4f, 0x75, 0x90, 0x3d, 0x03, 0xd8, 0x2c,
	0x6d, 0x6a, 0x31, 0x33, 0x48, 0x1c, 0xfc, 0xe2, 0x93, 0xcd, 0xd2, 0x5e, 0x7b, 0xc2, 0xaf, 0xdd,
	0x94, 0xc8, 0xa7, 0xe1, 0x58, 0x2e, 0x76, 0x92, 0xad, 0xae, 0x0a, 0x4a, 0x4b, 0x41, 0x77, 0x7c,
	0x3f, 0x48, 0x3c, 0xf3, 0x45, 0xd0, 0x1d, 0x3b, 0x81, 0xc9, 0x9d, 0xb6, 0x6d, 0xf6, 0xc0, 0x67,
	0xc7, 0x8e, 0xf0, 0xc9, 0x17, 0xb0, 0x5f, 0xeb, 0xbc, 0xda, 0x62, 0x9a, 0xe5, 0x42, 0x6d, 0xf9,
	0x61, 0xb8, 0x74, 0xe0, 0xde, 0x39, 0xca, 0xed, 0xb7, 0xaa, 0xb2, 0x0d, 0x12, 0xff, 0xcf, 0x27,
	0x5b, 0xe4, 0xde, 0x35, 0x28, 0x64, 0xaa, 0x8b, 0xbc, 0xe1, 0x47, 0x61, 0x41, 0x47, 0x7c, 0x2e,
	0xf2, 0x66, 0x91, 0xc3, 0xd3, 0xdd, 0xaf, 0x7f, 0xab, 0x65, 0x93, 0xe0, 0xd7, 0x0a, 0x2d, 0xb1,
	0x63, 0x18, 0xe8, 0x6f, 0x05, 0x9a, 0xd6, 0x06, 0x01, 0xb0, 0x4b, 0x38, 0xec, 0xcc, 0x16, 0xbc,
	0xe6, 0x1d, 0x31, 0xbd, 0x60, 0x71, 0x7d, 0x1e, 0xef, 0x3e, 0x96, 0x1c, 0x74, 0x95, 0x1e, 0x2e,
	0x7e, 0x44, 0x70, 0xf2, 0x49, 0x59, 0xda, 0xad, 0xb2, 0x09, 0xda, 0x52, 0x17, 0x16, 0x5d, 0xc3,
	0xcc, 0xdd, 0xc3, 0x37, 0x1c, 0x24, 0x01, 0xb0, 0x57, 0x30, 0x32, 0x68, 0xab, 0x9c, 0x2c, 0xef,
	0xcd, 0xfb, 0x7f, 0xe9, 0xd4, 0x95, 0xb8, 0xef, 0x2c, 0x0d, 0xd6, 0x4a, 0x57, 0xb6, 0xf5, 0xe3,
	0x03, 0xf6, 0x16, 0xc6, 0x7b, 0xe2, 0x7b, 0xad, 0x85, 0xf1, 0x9e, 0x16, 0x3f, 0x23, 0x38, 0xee,
	0xde, 0xba, 0x31, 0x88, 0x0f, 0xc3, 0x5c, 0xc2, 0x60, 0xad, 0x72, 0xb4, 0x3c, 0xf2, 0x4d, 0x5f,
	0x3e, 0x6e, 0xfa, 0xb8, 0x30, 0xfe, 0xe8, 0xaa, 0x3e, 0x14, 0x64, 0x9a, 0x24, 0x28, 0x5c, 0x1f,
	0xa9, 0x4c, 0x18, 0x77, 0x92, 0xf8, 0x98, 0x3d, 0x81, 0x91, 0xb2, 0xa9, 0xd4, 0x05, 0xfa, 0xb1,
	0xc6, 0xc9, 0x50, 0xd9, 0xf7, 0xba, 0xc0, 0xd9, 0x12, 0xe0, 0xcf, 0x0b, 0xec, 0x08, 0xfa, 0x1b,
	0x6c, 0xda, 0x8b, 0xbb, 0xd0, 0x1d, 0xa5, 0x16, 0x79, 0x15, 0xce, 0xdc, 0x4f, 0x02, 0x78, 0xd3,
	0x5b, 0x46, 0xab, 0xa1, 0x37, 0xfc, 0xeb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x00, 0x88,
	0x72, 0x00, 0x04, 0x00, 0x00,
}