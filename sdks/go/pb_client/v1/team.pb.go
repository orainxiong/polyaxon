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
// source: v1/team.proto

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

// Team specification
type Team struct {
	// Uuid
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Name
	Projects []string `protobuf:"bytes,3,rep,name=projects,proto3" json:"projects,omitempty"`
	// Optional time when the entityt was created
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Optional last time the entity was updated
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Team) Reset()         { *m = Team{} }
func (m *Team) String() string { return proto.CompactTextString(m) }
func (*Team) ProtoMessage()    {}
func (*Team) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc1b2b6d2e6c21ef, []int{0}
}

func (m *Team) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Team.Unmarshal(m, b)
}
func (m *Team) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Team.Marshal(b, m, deterministic)
}
func (m *Team) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Team.Merge(m, src)
}
func (m *Team) XXX_Size() int {
	return xxx_messageInfo_Team.Size(m)
}
func (m *Team) XXX_DiscardUnknown() {
	xxx_messageInfo_Team.DiscardUnknown(m)
}

var xxx_messageInfo_Team proto.InternalMessageInfo

func (m *Team) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Team) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Team) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

func (m *Team) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Team) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

// Request data to update teams
type TeamBodyRequest struct {
	// Owner of the namespace
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// Team body
	Team                 *Team    `protobuf:"bytes,2,opt,name=team,proto3" json:"team,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeamBodyRequest) Reset()         { *m = TeamBodyRequest{} }
func (m *TeamBodyRequest) String() string { return proto.CompactTextString(m) }
func (*TeamBodyRequest) ProtoMessage()    {}
func (*TeamBodyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc1b2b6d2e6c21ef, []int{1}
}

func (m *TeamBodyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamBodyRequest.Unmarshal(m, b)
}
func (m *TeamBodyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamBodyRequest.Marshal(b, m, deterministic)
}
func (m *TeamBodyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamBodyRequest.Merge(m, src)
}
func (m *TeamBodyRequest) XXX_Size() int {
	return xxx_messageInfo_TeamBodyRequest.Size(m)
}
func (m *TeamBodyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamBodyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TeamBodyRequest proto.InternalMessageInfo

func (m *TeamBodyRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *TeamBodyRequest) GetTeam() *Team {
	if m != nil {
		return m.Team
	}
	return nil
}

// Contains list teams
type ListTeamsResponse struct {
	// Count of the entities
	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// List of all entities
	Results []*Team `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	// Previous page
	Previous string `protobuf:"bytes,3,opt,name=previous,proto3" json:"previous,omitempty"`
	// Next page
	Next                 string   `protobuf:"bytes,4,opt,name=next,proto3" json:"next,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTeamsResponse) Reset()         { *m = ListTeamsResponse{} }
func (m *ListTeamsResponse) String() string { return proto.CompactTextString(m) }
func (*ListTeamsResponse) ProtoMessage()    {}
func (*ListTeamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc1b2b6d2e6c21ef, []int{2}
}

func (m *ListTeamsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTeamsResponse.Unmarshal(m, b)
}
func (m *ListTeamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTeamsResponse.Marshal(b, m, deterministic)
}
func (m *ListTeamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTeamsResponse.Merge(m, src)
}
func (m *ListTeamsResponse) XXX_Size() int {
	return xxx_messageInfo_ListTeamsResponse.Size(m)
}
func (m *ListTeamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTeamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTeamsResponse proto.InternalMessageInfo

func (m *ListTeamsResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListTeamsResponse) GetResults() []*Team {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *ListTeamsResponse) GetPrevious() string {
	if m != nil {
		return m.Previous
	}
	return ""
}

func (m *ListTeamsResponse) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

// Team member specification
type TeamMember struct {
	// User
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// User email
	UserEmail string `protobuf:"bytes,2,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	// Role
	Role string `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	// Organization Role
	OrgRole string `protobuf:"bytes,4,opt,name=org_role,json=orgRole,proto3" json:"org_role,omitempty"`
	// Optional time when the entityt was created
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Optional last time the entity was updated
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TeamMember) Reset()         { *m = TeamMember{} }
func (m *TeamMember) String() string { return proto.CompactTextString(m) }
func (*TeamMember) ProtoMessage()    {}
func (*TeamMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc1b2b6d2e6c21ef, []int{3}
}

func (m *TeamMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamMember.Unmarshal(m, b)
}
func (m *TeamMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamMember.Marshal(b, m, deterministic)
}
func (m *TeamMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamMember.Merge(m, src)
}
func (m *TeamMember) XXX_Size() int {
	return xxx_messageInfo_TeamMember.Size(m)
}
func (m *TeamMember) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamMember.DiscardUnknown(m)
}

var xxx_messageInfo_TeamMember proto.InternalMessageInfo

func (m *TeamMember) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *TeamMember) GetUserEmail() string {
	if m != nil {
		return m.UserEmail
	}
	return ""
}

func (m *TeamMember) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *TeamMember) GetOrgRole() string {
	if m != nil {
		return m.OrgRole
	}
	return ""
}

func (m *TeamMember) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *TeamMember) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

// Request data to create/update organization members
type TeamMemberBodyRequest struct {
	// Owner of the namespace
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// Team
	Team string `protobuf:"bytes,2,opt,name=team,proto3" json:"team,omitempty"`
	// Team body
	Member               *TeamMember `protobuf:"bytes,3,opt,name=member,proto3" json:"member,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TeamMemberBodyRequest) Reset()         { *m = TeamMemberBodyRequest{} }
func (m *TeamMemberBodyRequest) String() string { return proto.CompactTextString(m) }
func (*TeamMemberBodyRequest) ProtoMessage()    {}
func (*TeamMemberBodyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc1b2b6d2e6c21ef, []int{4}
}

func (m *TeamMemberBodyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamMemberBodyRequest.Unmarshal(m, b)
}
func (m *TeamMemberBodyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamMemberBodyRequest.Marshal(b, m, deterministic)
}
func (m *TeamMemberBodyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamMemberBodyRequest.Merge(m, src)
}
func (m *TeamMemberBodyRequest) XXX_Size() int {
	return xxx_messageInfo_TeamMemberBodyRequest.Size(m)
}
func (m *TeamMemberBodyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamMemberBodyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TeamMemberBodyRequest proto.InternalMessageInfo

func (m *TeamMemberBodyRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *TeamMemberBodyRequest) GetTeam() string {
	if m != nil {
		return m.Team
	}
	return ""
}

func (m *TeamMemberBodyRequest) GetMember() *TeamMember {
	if m != nil {
		return m.Member
	}
	return nil
}

// Contains list organization members
type ListTeamMembersResponse struct {
	// Count of the entities
	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// List of all entities
	Results []*TeamMember `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	// Previous page
	Previous string `protobuf:"bytes,3,opt,name=previous,proto3" json:"previous,omitempty"`
	// Next page
	Next                 string   `protobuf:"bytes,4,opt,name=next,proto3" json:"next,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTeamMembersResponse) Reset()         { *m = ListTeamMembersResponse{} }
func (m *ListTeamMembersResponse) String() string { return proto.CompactTextString(m) }
func (*ListTeamMembersResponse) ProtoMessage()    {}
func (*ListTeamMembersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc1b2b6d2e6c21ef, []int{5}
}

func (m *ListTeamMembersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTeamMembersResponse.Unmarshal(m, b)
}
func (m *ListTeamMembersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTeamMembersResponse.Marshal(b, m, deterministic)
}
func (m *ListTeamMembersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTeamMembersResponse.Merge(m, src)
}
func (m *ListTeamMembersResponse) XXX_Size() int {
	return xxx_messageInfo_ListTeamMembersResponse.Size(m)
}
func (m *ListTeamMembersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTeamMembersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTeamMembersResponse proto.InternalMessageInfo

func (m *ListTeamMembersResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListTeamMembersResponse) GetResults() []*TeamMember {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *ListTeamMembersResponse) GetPrevious() string {
	if m != nil {
		return m.Previous
	}
	return ""
}

func (m *ListTeamMembersResponse) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

func init() {
	proto.RegisterType((*Team)(nil), "v1.Team")
	proto.RegisterType((*TeamBodyRequest)(nil), "v1.TeamBodyRequest")
	proto.RegisterType((*ListTeamsResponse)(nil), "v1.ListTeamsResponse")
	proto.RegisterType((*TeamMember)(nil), "v1.TeamMember")
	proto.RegisterType((*TeamMemberBodyRequest)(nil), "v1.TeamMemberBodyRequest")
	proto.RegisterType((*ListTeamMembersResponse)(nil), "v1.ListTeamMembersResponse")
}

func init() { proto.RegisterFile("v1/team.proto", fileDescriptor_bc1b2b6d2e6c21ef) }

var fileDescriptor_bc1b2b6d2e6c21ef = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xbf, 0x8e, 0xd3, 0x40,
	0x10, 0xc6, 0xe5, 0xc4, 0xce, 0xc5, 0x13, 0x01, 0x62, 0x05, 0xc2, 0x44, 0x20, 0x2c, 0x17, 0xc8,
	0x95, 0xa3, 0x84, 0x8a, 0xf2, 0x90, 0xae, 0x83, 0x66, 0x75, 0x7d, 0xe4, 0xc4, 0x83, 0x65, 0xe4,
	0xf5, 0x9a, 0xfd, 0x63, 0x8e, 0x82, 0x07, 0xe0, 0xad, 0x78, 0x18, 0x1e, 0x04, 0xcd, 0xae, 0x7d,
	0x70, 0x69, 0x22, 0x52, 0x79, 0xe6, 0x9b, 0x9d, 0xf1, 0xcc, 0xf7, 0x83, 0x47, 0xc3, 0x76, 0x63,
	0xb0, 0x14, 0x45, 0xaf, 0xa4, 0x91, 0x6c, 0x36, 0x6c, 0xd7, 0x6f, 0x6a, 0x29, 0xeb, 0x16, 0x37,
	0x4e, 0x39, 0xd8, 0xcf, 0x1b, 0xd3, 0x08, 0xd4, 0xa6, 0x14, 0xbd, 0x7f, 0x94, 0xfd, 0x0a, 0x20,
	0xbc, 0xc5, 0x52, 0x30, 0x06, 0xa1, 0xb5, 0x4d, 0x95, 0x04, 0x69, 0x90, 0xc7, 0xdc, 0xc5, 0xa4,
	0x75, 0xa5, 0xc0, 0x64, 0xe6, 0x35, 0x8a, 0xd9, 0x1a, 0x96, 0xbd, 0x92, 0x5f, 0xf0, 0x68, 0x74,
	0x32, 0x4f, 0xe7, 0x79, 0xcc, 0xef, 0x73, 0xf6, 0x1e, 0xe0, 0xa8, 0xb0, 0x34, 0x58, 0xed, 0x4b,
	0x93, 0x84, 0x69, 0x90, 0xaf, 0x76, 0xeb, 0xc2, 0xaf, 0x50, 0x4c, 0x2b, 0x14, 0xb7, 0xd3, 0x0a,
	0x3c, 0x1e, 0x5f, 0x5f, 0x1b, 0x6a, 0xb5, 0x7d, 0x35, 0xb5, 0x46, 0xe7, 0x5b, 0xc7, 0xd7, 0xd7,
	0x26, 0xbb, 0x81, 0x27, 0x74, 0xc1, 0x07, 0x59, 0x7d, 0xe7, 0xf8, 0xd5, 0xa2, 0x36, 0xec, 0x19,
	0x44, 0xf2, 0x5b, 0x87, 0x6a, 0xbc, 0xc6, 0x27, 0xec, 0x15, 0x84, 0x64, 0x8f, 0x3b, 0x67, 0xb5,
	0x5b, 0x16, 0xc3, 0xb6, 0xa0, 0x46, 0xee, 0xd4, 0xec, 0x07, 0x3c, 0xfd, 0xd8, 0x68, 0x43, 0x8a,
	0xe6, 0xa8, 0x7b, 0xd9, 0x69, 0xa4, 0x41, 0x47, 0x69, 0x3b, 0xe3, 0x06, 0x45, 0xdc, 0x27, 0x2c,
	0x83, 0x2b, 0x85, 0xda, 0xb6, 0x46, 0x27, 0xb3, 0x74, 0xfe, 0x60, 0xd6, 0x54, 0xf0, 0x3e, 0xe1,
	0xd0, 0x48, 0x4b, 0x3e, 0x05, 0xde, 0x27, 0x9f, 0x3b, 0x5f, 0xf1, 0xce, 0x3b, 0x44, 0xbe, 0xe2,
	0x9d, 0xc9, 0x7e, 0x07, 0x00, 0x34, 0xe1, 0x13, 0x8a, 0x03, 0x2a, 0x87, 0x43, 0xdf, 0x1f, 0xe0,
	0x62, 0xf6, 0x1a, 0x80, 0xbe, 0x7b, 0x14, 0x65, 0xd3, 0x8e, 0x50, 0x62, 0x52, 0x6e, 0x48, 0xa0,
	0x16, 0x25, 0x5b, 0x1c, 0xff, 0xe6, 0x62, 0xf6, 0x12, 0x96, 0x52, 0xd5, 0x7b, 0xa7, 0xfb, 0xbf,
	0x5d, 0x49, 0x55, 0x73, 0x2a, 0x3d, 0x84, 0x15, 0x5d, 0x0e, 0x6b, 0xf1, 0x3f, 0xb0, 0x1a, 0x78,
	0xfe, 0xf7, 0xca, 0xf3, 0xc8, 0xd8, 0x3f, 0xc8, 0x62, 0x0f, 0x8a, 0xbd, 0x85, 0x85, 0x70, 0xed,
	0xee, 0xd2, 0xd5, 0xee, 0xf1, 0x64, 0xbe, 0x1f, 0xca, 0xc7, 0x6a, 0xf6, 0x33, 0x80, 0x17, 0x13,
	0x51, 0x5f, 0x3a, 0xc7, 0x35, 0x3f, 0xe5, 0x7a, 0x3a, 0xfa, 0x52, 0xba, 0x87, 0x85, 0x73, 0xe5,
	0xdd, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0xfa, 0x6d, 0xbf, 0xa3, 0x03, 0x00, 0x00,
}
