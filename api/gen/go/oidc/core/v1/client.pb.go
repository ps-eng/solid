// Licensed to SolID under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. SolID licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.12.0
// source: oidc/core/v1/client.proto

package corev1

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

// ClientType describes OIDC Client type enumeration.
// https://tools.ietf.org/html/rfc6749#section-2.1
type ClientType int32

const (
	// Default value
	ClientType_CLIENT_TYPE_INVALID ClientType = 0
	// Explicit unknown
	ClientType_CLIENT_TYPE_UNKNOWN ClientType = 1
	// Clients capable of maintaining the confidentiality of their
	// credentials (e.g., client implemented on a secure server with
	// restricted access to the client credentials), or capable of secure
	// client authentication using other means.
	ClientType_CLIENT_TYPE_CONFIDENTIAL ClientType = 2
	// Clients incapable of maintaining the confidentiality of their
	// credentials (e.g., clients executing on the device used by the
	// resource owner, such as an installed native application or a web
	// browser-based application), and incapable of secure client
	// authentication via any other means.
	ClientType_CLIENT_TYPE_PUBLIC ClientType = 3
)

// Enum value maps for ClientType.
var (
	ClientType_name = map[int32]string{
		0: "CLIENT_TYPE_INVALID",
		1: "CLIENT_TYPE_UNKNOWN",
		2: "CLIENT_TYPE_CONFIDENTIAL",
		3: "CLIENT_TYPE_PUBLIC",
	}
	ClientType_value = map[string]int32{
		"CLIENT_TYPE_INVALID":      0,
		"CLIENT_TYPE_UNKNOWN":      1,
		"CLIENT_TYPE_CONFIDENTIAL": 2,
		"CLIENT_TYPE_PUBLIC":       3,
	}
)

func (x ClientType) Enum() *ClientType {
	p := new(ClientType)
	*p = x
	return p
}

func (x ClientType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientType) Descriptor() protoreflect.EnumDescriptor {
	return file_oidc_core_v1_client_proto_enumTypes[0].Descriptor()
}

func (ClientType) Type() protoreflect.EnumType {
	return &file_oidc_core_v1_client_proto_enumTypes[0]
}

func (x ClientType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientType.Descriptor instead.
func (ClientType) EnumDescriptor() ([]byte, []int) {
	return file_oidc_core_v1_client_proto_rawDescGZIP(), []int{0}
}

// Client defines internal OIDC client properties.
type Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId        string     `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientType      ClientType `protobuf:"varint,2,opt,name=client_type,json=clientType,proto3,enum=oidc.core.v1.ClientType" json:"client_type,omitempty"`
	RedirectUris    []string   `protobuf:"bytes,3,rep,name=redirect_uris,json=redirectUris,proto3" json:"redirect_uris,omitempty"`
	ResponseTypes   []string   `protobuf:"bytes,4,rep,name=response_types,json=responseTypes,proto3" json:"response_types,omitempty"`
	GrantTypes      []string   `protobuf:"bytes,5,rep,name=grant_types,json=grantTypes,proto3" json:"grant_types,omitempty"`
	ApplicationType string     `protobuf:"bytes,6,opt,name=application_type,json=applicationType,proto3" json:"application_type,omitempty"`
	Contacts        []string   `protobuf:"bytes,7,rep,name=contacts,proto3" json:"contacts,omitempty"`
	ClientName      string     `protobuf:"bytes,8,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
	LogoUri         string     `protobuf:"bytes,9,opt,name=logo_uri,json=logoUri,proto3" json:"logo_uri,omitempty"`
	ClientUri       string     `protobuf:"bytes,10,opt,name=client_uri,json=clientUri,proto3" json:"client_uri,omitempty"`
	PolicyUri       string     `protobuf:"bytes,11,opt,name=policy_uri,json=policyUri,proto3" json:"policy_uri,omitempty"`
	TosUri          string     `protobuf:"bytes,12,opt,name=tos_uri,json=tosUri,proto3" json:"tos_uri,omitempty"`
	JwksUri         string     `protobuf:"bytes,13,opt,name=jwks_uri,json=jwksUri,proto3" json:"jwks_uri,omitempty"`
	Jwks            []byte     `protobuf:"bytes,14,opt,name=jwks,proto3" json:"jwks,omitempty"`
	ClientSecret    []byte     `protobuf:"bytes,15,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
}

func (x *Client) Reset() {
	*x = Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidc_core_v1_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
	mi := &file_oidc_core_v1_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client.ProtoReflect.Descriptor instead.
func (*Client) Descriptor() ([]byte, []int) {
	return file_oidc_core_v1_client_proto_rawDescGZIP(), []int{0}
}

func (x *Client) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Client) GetClientType() ClientType {
	if x != nil {
		return x.ClientType
	}
	return ClientType_CLIENT_TYPE_INVALID
}

func (x *Client) GetRedirectUris() []string {
	if x != nil {
		return x.RedirectUris
	}
	return nil
}

func (x *Client) GetResponseTypes() []string {
	if x != nil {
		return x.ResponseTypes
	}
	return nil
}

func (x *Client) GetGrantTypes() []string {
	if x != nil {
		return x.GrantTypes
	}
	return nil
}

func (x *Client) GetApplicationType() string {
	if x != nil {
		return x.ApplicationType
	}
	return ""
}

func (x *Client) GetContacts() []string {
	if x != nil {
		return x.Contacts
	}
	return nil
}

func (x *Client) GetClientName() string {
	if x != nil {
		return x.ClientName
	}
	return ""
}

func (x *Client) GetLogoUri() string {
	if x != nil {
		return x.LogoUri
	}
	return ""
}

func (x *Client) GetClientUri() string {
	if x != nil {
		return x.ClientUri
	}
	return ""
}

func (x *Client) GetPolicyUri() string {
	if x != nil {
		return x.PolicyUri
	}
	return ""
}

func (x *Client) GetTosUri() string {
	if x != nil {
		return x.TosUri
	}
	return ""
}

func (x *Client) GetJwksUri() string {
	if x != nil {
		return x.JwksUri
	}
	return ""
}

func (x *Client) GetJwks() []byte {
	if x != nil {
		return x.Jwks
	}
	return nil
}

func (x *Client) GetClientSecret() []byte {
	if x != nil {
		return x.ClientSecret
	}
	return nil
}

var File_oidc_core_v1_client_proto protoreflect.FileDescriptor

var file_oidc_core_v1_client_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6f, 0x69, 0x64,
	0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x22, 0xfb, 0x03, 0x0a, 0x06, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x39, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x69,
	0x73, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x72, 0x61, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x67,
	0x72, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x69, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x72, 0x69, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x55, 0x72, 0x69, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f,
	0x73, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x73,
	0x55, 0x72, 0x69, 0x12, 0x19, 0x0a, 0x08, 0x6a, 0x77, 0x6b, 0x73, 0x5f, 0x75, 0x72, 0x69, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x77, 0x6b, 0x73, 0x55, 0x72, 0x69, 0x12, 0x12,
	0x0a, 0x04, 0x6a, 0x77, 0x6b, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6a, 0x77,
	0x6b, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2a, 0x74, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x17,
	0x0a, 0x13, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4c, 0x49, 0x45, 0x4e,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x44, 0x45, 0x4e, 0x54,
	0x49, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x03, 0x42, 0x15, 0x5a,
	0x13, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f,
	0x72, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oidc_core_v1_client_proto_rawDescOnce sync.Once
	file_oidc_core_v1_client_proto_rawDescData = file_oidc_core_v1_client_proto_rawDesc
)

func file_oidc_core_v1_client_proto_rawDescGZIP() []byte {
	file_oidc_core_v1_client_proto_rawDescOnce.Do(func() {
		file_oidc_core_v1_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_oidc_core_v1_client_proto_rawDescData)
	})
	return file_oidc_core_v1_client_proto_rawDescData
}

var file_oidc_core_v1_client_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_oidc_core_v1_client_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_oidc_core_v1_client_proto_goTypes = []interface{}{
	(ClientType)(0), // 0: oidc.core.v1.ClientType
	(*Client)(nil),  // 1: oidc.core.v1.Client
}
var file_oidc_core_v1_client_proto_depIdxs = []int32{
	0, // 0: oidc.core.v1.Client.client_type:type_name -> oidc.core.v1.ClientType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_oidc_core_v1_client_proto_init() }
func file_oidc_core_v1_client_proto_init() {
	if File_oidc_core_v1_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oidc_core_v1_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client); i {
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
			RawDescriptor: file_oidc_core_v1_client_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oidc_core_v1_client_proto_goTypes,
		DependencyIndexes: file_oidc_core_v1_client_proto_depIdxs,
		EnumInfos:         file_oidc_core_v1_client_proto_enumTypes,
		MessageInfos:      file_oidc_core_v1_client_proto_msgTypes,
	}.Build()
	File_oidc_core_v1_client_proto = out.File
	file_oidc_core_v1_client_proto_rawDesc = nil
	file_oidc_core_v1_client_proto_goTypes = nil
	file_oidc_core_v1_client_proto_depIdxs = nil
}
