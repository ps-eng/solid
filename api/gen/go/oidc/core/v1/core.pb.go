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
// 	protoc        v3.11.4
// source: oidc/core/v1/core.proto

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

// Display defines values to set how to displays the authentication and consent
// user interface pages to the End-User.
type Display int32

const (
	// Default value when nothing specificied.
	Display_DISPLAY_INVALID Display = 0
	// Value to set as unknown.
	Display_DISPLAY_UNKNOWN Display = 1
	// The Authorization Server SHOULD display the authentication and consent UI
	// consistent with a full User Agent page view. If the display parameter is
	// not specified, this is the default display mode.
	Display_DISPLAY_PAGE Display = 2
	// The Authorization Server SHOULD display the authentication and consent UI
	// consistent with a popup User Agent window. The popup User Agent window
	// should be of an appropriate size for a login-focused dialog and should not
	// obscure the entire window that it is popping up over.
	Display_DISPLAY_POPUP Display = 3
	// The Authorization Server SHOULD display the authentication and consent UI
	// consistent with a device that leverages a touch interface.
	Display_DISPLAY_TOUCH Display = 4
	// The Authorization Server SHOULD display the authentication and consent UI
	// consistent with a "feature phone" type display.
	Display_DISPLAY_WAP Display = 5
)

// Enum value maps for Display.
var (
	Display_name = map[int32]string{
		0: "DISPLAY_INVALID",
		1: "DISPLAY_UNKNOWN",
		2: "DISPLAY_PAGE",
		3: "DISPLAY_POPUP",
		4: "DISPLAY_TOUCH",
		5: "DISPLAY_WAP",
	}
	Display_value = map[string]int32{
		"DISPLAY_INVALID": 0,
		"DISPLAY_UNKNOWN": 1,
		"DISPLAY_PAGE":    2,
		"DISPLAY_POPUP":   3,
		"DISPLAY_TOUCH":   4,
		"DISPLAY_WAP":     5,
	}
)

func (x Display) Enum() *Display {
	p := new(Display)
	*p = x
	return p
}

func (x Display) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Display) Descriptor() protoreflect.EnumDescriptor {
	return file_oidc_core_v1_core_proto_enumTypes[0].Descriptor()
}

func (Display) Type() protoreflect.EnumType {
	return &file_oidc_core_v1_core_proto_enumTypes[0]
}

func (x Display) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Display.Descriptor instead.
func (Display) EnumDescriptor() ([]byte, []int) {
	return file_oidc_core_v1_core_proto_rawDescGZIP(), []int{0}
}

// Prompt defines values for required prompt actions.
type Prompt int32

const (
	// Default value when nothing specificied.
	Prompt_PROMPT_INVALID Prompt = 0
	// Value to set as unknown.
	Prompt_PROMPT_UNKNOWN Prompt = 1
	// The Authorization Server MUST NOT display any authentication or consent
	// user interface pages. An error is returned if an End-User is not already
	// authenticated or the Client does not have pre-configured consent for the
	// requested Claims or does not fulfill other conditions for processing the
	// request. The error code will typically be login_required,
	// interaction_required, or another code defined in Section 3.1.2.6. This
	// can be used as a method to check for existing authentication and/or
	// consent.
	Prompt_PROMPT_NONE Prompt = 2
	// The Authorization Server SHOULD prompt the End-User for reauthentication.
	// If it cannot reauthenticate the End-User, it MUST return an error,
	// typically login_required.
	Prompt_PROMPT_LOGIN Prompt = 3
	// The Authorization Server SHOULD prompt the End-User for consent before
	// returning information to the Client. If it cannot obtain consent, it MUST
	// return an error, typically consent_required.
	Prompt_PROMPT_CONSENT Prompt = 4
	// The Authorization Server SHOULD prompt the End-User to select a user
	// account. This enables an End-User who has multiple accounts at the
	// Authorization Server to select amongst the multiple accounts that they
	// might have current sessions for. If it cannot obtain an account selection
	// choice made by the End-User, it MUST return an error, typically
	// account_selection_required.
	Prompt_PROMPT_SELECT_ACCOUNT Prompt = 5
)

// Enum value maps for Prompt.
var (
	Prompt_name = map[int32]string{
		0: "PROMPT_INVALID",
		1: "PROMPT_UNKNOWN",
		2: "PROMPT_NONE",
		3: "PROMPT_LOGIN",
		4: "PROMPT_CONSENT",
		5: "PROMPT_SELECT_ACCOUNT",
	}
	Prompt_value = map[string]int32{
		"PROMPT_INVALID":        0,
		"PROMPT_UNKNOWN":        1,
		"PROMPT_NONE":           2,
		"PROMPT_LOGIN":          3,
		"PROMPT_CONSENT":        4,
		"PROMPT_SELECT_ACCOUNT": 5,
	}
)

func (x Prompt) Enum() *Prompt {
	p := new(Prompt)
	*p = x
	return p
}

func (x Prompt) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Prompt) Descriptor() protoreflect.EnumDescriptor {
	return file_oidc_core_v1_core_proto_enumTypes[1].Descriptor()
}

func (Prompt) Type() protoreflect.EnumType {
	return &file_oidc_core_v1_core_proto_enumTypes[1]
}

func (x Prompt) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Prompt.Descriptor instead.
func (Prompt) EnumDescriptor() ([]byte, []int) {
	return file_oidc_core_v1_core_proto_rawDescGZIP(), []int{1}
}

type GrantAuthorizationCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code         string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	RedirectUri  string `protobuf:"bytes,2,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	CodeVerifier string `protobuf:"bytes,3,opt,name=code_verifier,json=codeVerifier,proto3" json:"code_verifier,omitempty"`
}

func (x *GrantAuthorizationCode) Reset() {
	*x = GrantAuthorizationCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidc_core_v1_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrantAuthorizationCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantAuthorizationCode) ProtoMessage() {}

func (x *GrantAuthorizationCode) ProtoReflect() protoreflect.Message {
	mi := &file_oidc_core_v1_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantAuthorizationCode.ProtoReflect.Descriptor instead.
func (*GrantAuthorizationCode) Descriptor() ([]byte, []int) {
	return file_oidc_core_v1_core_proto_rawDescGZIP(), []int{0}
}

func (x *GrantAuthorizationCode) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GrantAuthorizationCode) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

func (x *GrantAuthorizationCode) GetCodeVerifier() string {
	if x != nil {
		return x.CodeVerifier
	}
	return ""
}

type GrantRefreshToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *GrantRefreshToken) Reset() {
	*x = GrantRefreshToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidc_core_v1_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrantRefreshToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantRefreshToken) ProtoMessage() {}

func (x *GrantRefreshToken) ProtoReflect() protoreflect.Message {
	mi := &file_oidc_core_v1_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantRefreshToken.ProtoReflect.Descriptor instead.
func (*GrantRefreshToken) Descriptor() ([]byte, []int) {
	return file_oidc_core_v1_core_proto_rawDescGZIP(), []int{1}
}

func (x *GrantRefreshToken) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type GrantDeviceCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GrantDeviceCode) Reset() {
	*x = GrantDeviceCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidc_core_v1_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrantDeviceCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantDeviceCode) ProtoMessage() {}

func (x *GrantDeviceCode) ProtoReflect() protoreflect.Message {
	mi := &file_oidc_core_v1_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantDeviceCode.ProtoReflect.Descriptor instead.
func (*GrantDeviceCode) Descriptor() ([]byte, []int) {
	return file_oidc_core_v1_core_proto_rawDescGZIP(), []int{2}
}

type GrantClientCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GrantClientCredentials) Reset() {
	*x = GrantClientCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidc_core_v1_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrantClientCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantClientCredentials) ProtoMessage() {}

func (x *GrantClientCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_oidc_core_v1_core_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantClientCredentials.ProtoReflect.Descriptor instead.
func (*GrantClientCredentials) Descriptor() ([]byte, []int) {
	return file_oidc_core_v1_core_proto_rawDescGZIP(), []int{3}
}

var File_oidc_core_v1_core_proto protoreflect.FileDescriptor

var file_oidc_core_v1_core_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6f, 0x69, 0x64, 0x63, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x74, 0x0a, 0x16, 0x47, 0x72, 0x61, 0x6e, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x69, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x64, 0x65,
	0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x6f, 0x64, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x38, 0x0a,
	0x11, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x72, 0x61, 0x6e, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x72,
	0x61, 0x6e, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x2a, 0x7c, 0x0a, 0x07, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12,
	0x13, 0x0a, 0x0f, 0x44, 0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x49, 0x53,
	0x50, 0x4c, 0x41, 0x59, 0x5f, 0x50, 0x41, 0x47, 0x45, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x44,
	0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x50, 0x4f, 0x50, 0x55, 0x50, 0x10, 0x03, 0x12, 0x11,
	0x0a, 0x0d, 0x44, 0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x54, 0x4f, 0x55, 0x43, 0x48, 0x10,
	0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x57, 0x41, 0x50,
	0x10, 0x05, 0x2a, 0x82, 0x01, 0x0a, 0x06, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x12, 0x0a,
	0x0e, 0x50, 0x52, 0x4f, 0x4d, 0x50, 0x54, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x52, 0x4f, 0x4d, 0x50, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x52, 0x4f, 0x4d, 0x50, 0x54, 0x5f,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x4d, 0x50, 0x54,
	0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x52, 0x4f, 0x4d,
	0x50, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x45, 0x4e, 0x54, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15,
	0x50, 0x52, 0x4f, 0x4d, 0x50, 0x54, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x5f, 0x41, 0x43,
	0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x05, 0x42, 0x15, 0x5a, 0x13, 0x6f, 0x69, 0x64, 0x63, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oidc_core_v1_core_proto_rawDescOnce sync.Once
	file_oidc_core_v1_core_proto_rawDescData = file_oidc_core_v1_core_proto_rawDesc
)

func file_oidc_core_v1_core_proto_rawDescGZIP() []byte {
	file_oidc_core_v1_core_proto_rawDescOnce.Do(func() {
		file_oidc_core_v1_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_oidc_core_v1_core_proto_rawDescData)
	})
	return file_oidc_core_v1_core_proto_rawDescData
}

var file_oidc_core_v1_core_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_oidc_core_v1_core_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_oidc_core_v1_core_proto_goTypes = []interface{}{
	(Display)(0),                   // 0: oidc.core.v1.Display
	(Prompt)(0),                    // 1: oidc.core.v1.Prompt
	(*GrantAuthorizationCode)(nil), // 2: oidc.core.v1.GrantAuthorizationCode
	(*GrantRefreshToken)(nil),      // 3: oidc.core.v1.GrantRefreshToken
	(*GrantDeviceCode)(nil),        // 4: oidc.core.v1.GrantDeviceCode
	(*GrantClientCredentials)(nil), // 5: oidc.core.v1.GrantClientCredentials
}
var file_oidc_core_v1_core_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_oidc_core_v1_core_proto_init() }
func file_oidc_core_v1_core_proto_init() {
	if File_oidc_core_v1_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oidc_core_v1_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrantAuthorizationCode); i {
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
		file_oidc_core_v1_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrantRefreshToken); i {
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
		file_oidc_core_v1_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrantDeviceCode); i {
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
		file_oidc_core_v1_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrantClientCredentials); i {
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
			RawDescriptor: file_oidc_core_v1_core_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oidc_core_v1_core_proto_goTypes,
		DependencyIndexes: file_oidc_core_v1_core_proto_depIdxs,
		EnumInfos:         file_oidc_core_v1_core_proto_enumTypes,
		MessageInfos:      file_oidc_core_v1_core_proto_msgTypes,
	}.Build()
	File_oidc_core_v1_core_proto = out.File
	file_oidc_core_v1_core_proto_rawDesc = nil
	file_oidc_core_v1_core_proto_goTypes = nil
	file_oidc_core_v1_core_proto_depIdxs = nil
}
