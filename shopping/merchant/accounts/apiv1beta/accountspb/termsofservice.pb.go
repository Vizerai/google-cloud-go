// Copyright 2024 Google LLC
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
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.3
// source: google/shopping/merchant/accounts/v1beta/termsofservice.proto

package accountspb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A `TermsOfService`.
type TermsOfService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. The resource name of the terms of service version.
	// Format: `termsOfService/{version}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Region code as defined by [CLDR](https://cldr.unicode.org/). This is either
	// a country where the ToS applies specifically to that country or `001` when
	// the same `TermsOfService` can be signed in any country. However note that
	// when signing a ToS that applies globally we still expect that a specific
	// country is provided  (this should be merchant business country or program
	// country of participation).
	RegionCode string `protobuf:"bytes,2,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	// The Kind this terms of service version applies to.
	Kind TermsOfServiceKind `protobuf:"varint,3,opt,name=kind,proto3,enum=google.shopping.merchant.accounts.v1beta.TermsOfServiceKind" json:"kind,omitempty"`
	// URI for terms of service file that needs to be displayed to signing users.
	FileUri *string `protobuf:"bytes,4,opt,name=file_uri,json=fileUri,proto3,oneof" json:"file_uri,omitempty"`
	// Whether this terms of service version is external. External terms of
	// service versions can only be agreed through external processes and not
	// directly by the merchant through UI or API.
	External bool `protobuf:"varint,5,opt,name=external,proto3" json:"external,omitempty"`
}

func (x *TermsOfService) Reset() {
	*x = TermsOfService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TermsOfService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TermsOfService) ProtoMessage() {}

func (x *TermsOfService) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TermsOfService.ProtoReflect.Descriptor instead.
func (*TermsOfService) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_rawDescGZIP(), []int{0}
}

func (x *TermsOfService) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TermsOfService) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *TermsOfService) GetKind() TermsOfServiceKind {
	if x != nil {
		return x.Kind
	}
	return TermsOfServiceKind_TERMS_OF_SERVICE_KIND_UNSPECIFIED
}

func (x *TermsOfService) GetFileUri() string {
	if x != nil && x.FileUri != nil {
		return *x.FileUri
	}
	return ""
}

func (x *TermsOfService) GetExternal() bool {
	if x != nil {
		return x.External
	}
	return false
}

// Request message for the `GetTermsOfService` method.
type GetTermsOfServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the terms of service version.
	// Format: `termsOfService/{version}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetTermsOfServiceRequest) Reset() {
	*x = GetTermsOfServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTermsOfServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTermsOfServiceRequest) ProtoMessage() {}

func (x *GetTermsOfServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTermsOfServiceRequest.ProtoReflect.Descriptor instead.
func (*GetTermsOfServiceRequest) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_rawDescGZIP(), []int{1}
}

func (x *GetTermsOfServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request message for the `RetrieveLatestTermsOfService` method.
type RetrieveLatestTermsOfServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Region code as defined by [CLDR](https://cldr.unicode.org/). This is either
	// a country when the ToS applies specifically to that country or 001 when it
	// applies globally.
	RegionCode string `protobuf:"bytes,1,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	// The Kind this terms of service version applies to.
	Kind TermsOfServiceKind `protobuf:"varint,2,opt,name=kind,proto3,enum=google.shopping.merchant.accounts.v1beta.TermsOfServiceKind" json:"kind,omitempty"`
}

func (x *RetrieveLatestTermsOfServiceRequest) Reset() {
	*x = RetrieveLatestTermsOfServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveLatestTermsOfServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveLatestTermsOfServiceRequest) ProtoMessage() {}

func (x *RetrieveLatestTermsOfServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveLatestTermsOfServiceRequest.ProtoReflect.Descriptor instead.
func (*RetrieveLatestTermsOfServiceRequest) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_rawDescGZIP(), []int{2}
}

func (x *RetrieveLatestTermsOfServiceRequest) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *RetrieveLatestTermsOfServiceRequest) GetKind() TermsOfServiceKind {
	if x != nil {
		return x.Kind
	}
	return TermsOfServiceKind_TERMS_OF_SERVICE_KIND_UNSPECIFIED
}

// Request message for the `AcceptTermsOfService` method.
type AcceptTermsOfServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the terms of service version.
	// Format: `termsOfService/{version}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The account for which to accept the ToS.
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	// Required. Region code as defined by [CLDR](https://cldr.unicode.org/). This
	// is either a country when the ToS applies specifically to that country or
	// 001 when it applies globally.
	RegionCode string `protobuf:"bytes,3,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
}

func (x *AcceptTermsOfServiceRequest) Reset() {
	*x = AcceptTermsOfServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptTermsOfServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptTermsOfServiceRequest) ProtoMessage() {}

func (x *AcceptTermsOfServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptTermsOfServiceRequest.ProtoReflect.Descriptor instead.
func (*AcceptTermsOfServiceRequest) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_rawDescGZIP(), []int{3}
}

func (x *AcceptTermsOfServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AcceptTermsOfServiceRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *AcceptTermsOfServiceRequest) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

var File_google_shopping_merchant_accounts_v1beta_termsofservice_proto protoreflect.FileDescriptor

var file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x74, 0x65, 0x72, 0x6d, 0x73,
	0x6f, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x2f, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x6f, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x6b, 0x69, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x02, 0x0a,
	0x0e, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x45, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x31, 0xe0,
	0x41, 0x08, 0xfa, 0x41, 0x2b, 0x0a, 0x29, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x61,
	0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x50, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x2e, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4b,
	0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x66,
	0x69, 0x6c, 0x65, 0x55, 0x72, 0x69, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x3a, 0x48, 0xea, 0x41, 0x45, 0x0a, 0x29, 0x6d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x18, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x7b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x69, 0x22, 0x61, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x31, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2b, 0x0a, 0x29,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x72, 0x6d, 0x73,
	0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x98, 0x01, 0x0a, 0x23, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x50, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0xd0, 0x01, 0x0a, 0x1b, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x31, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2b,
	0x0a, 0x29, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x72,
	0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x44, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x2a, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x24, 0x0a, 0x22, 0x6d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0xc7, 0x05,
	0x0a, 0x15, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xca, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54,
	0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x72, 0x6d,
	0x73, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x54, 0x65, 0x72,
	0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x37, 0xda, 0x41, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x12, 0x28, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x2a, 0x7d, 0x12, 0xdf, 0x01, 0x0a, 0x1c, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x65, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x54,
	0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68,
	0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e,
	0x54, 0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x36,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x12, 0x2e, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x12, 0xb5, 0x01, 0x0a, 0x14, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x45, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3e,
	0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x12, 0x2f, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f,
	0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x2a, 0x7d, 0x3a, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x1a, 0x47,
	0xca, 0x41, 0x1a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x95, 0x01, 0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x42, 0x13, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x4f,
	0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x4e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x70, 0x62, 0x3b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_rawDescOnce sync.Once
	file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_rawDescData = file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_rawDesc
)

func file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_rawDescGZIP() []byte {
	file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_rawDescOnce.Do(func() {
		file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_rawDescData)
	})
	return file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_rawDescData
}

var file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_goTypes = []interface{}{
	(*TermsOfService)(nil),                      // 0: google.shopping.merchant.accounts.v1beta.TermsOfService
	(*GetTermsOfServiceRequest)(nil),            // 1: google.shopping.merchant.accounts.v1beta.GetTermsOfServiceRequest
	(*RetrieveLatestTermsOfServiceRequest)(nil), // 2: google.shopping.merchant.accounts.v1beta.RetrieveLatestTermsOfServiceRequest
	(*AcceptTermsOfServiceRequest)(nil),         // 3: google.shopping.merchant.accounts.v1beta.AcceptTermsOfServiceRequest
	(TermsOfServiceKind)(0),                     // 4: google.shopping.merchant.accounts.v1beta.TermsOfServiceKind
	(*emptypb.Empty)(nil),                       // 5: google.protobuf.Empty
}
var file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_depIdxs = []int32{
	4, // 0: google.shopping.merchant.accounts.v1beta.TermsOfService.kind:type_name -> google.shopping.merchant.accounts.v1beta.TermsOfServiceKind
	4, // 1: google.shopping.merchant.accounts.v1beta.RetrieveLatestTermsOfServiceRequest.kind:type_name -> google.shopping.merchant.accounts.v1beta.TermsOfServiceKind
	1, // 2: google.shopping.merchant.accounts.v1beta.TermsOfServiceService.GetTermsOfService:input_type -> google.shopping.merchant.accounts.v1beta.GetTermsOfServiceRequest
	2, // 3: google.shopping.merchant.accounts.v1beta.TermsOfServiceService.RetrieveLatestTermsOfService:input_type -> google.shopping.merchant.accounts.v1beta.RetrieveLatestTermsOfServiceRequest
	3, // 4: google.shopping.merchant.accounts.v1beta.TermsOfServiceService.AcceptTermsOfService:input_type -> google.shopping.merchant.accounts.v1beta.AcceptTermsOfServiceRequest
	0, // 5: google.shopping.merchant.accounts.v1beta.TermsOfServiceService.GetTermsOfService:output_type -> google.shopping.merchant.accounts.v1beta.TermsOfService
	0, // 6: google.shopping.merchant.accounts.v1beta.TermsOfServiceService.RetrieveLatestTermsOfService:output_type -> google.shopping.merchant.accounts.v1beta.TermsOfService
	5, // 7: google.shopping.merchant.accounts.v1beta.TermsOfServiceService.AcceptTermsOfService:output_type -> google.protobuf.Empty
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_init() }
func file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_init() {
	if File_google_shopping_merchant_accounts_v1beta_termsofservice_proto != nil {
		return
	}
	file_google_shopping_merchant_accounts_v1beta_termsofservicekind_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TermsOfService); i {
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
		file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTermsOfServiceRequest); i {
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
		file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveLatestTermsOfServiceRequest); i {
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
		file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptTermsOfServiceRequest); i {
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
	file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_goTypes,
		DependencyIndexes: file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_depIdxs,
		MessageInfos:      file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_msgTypes,
	}.Build()
	File_google_shopping_merchant_accounts_v1beta_termsofservice_proto = out.File
	file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_rawDesc = nil
	file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_goTypes = nil
	file_google_shopping_merchant_accounts_v1beta_termsofservice_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TermsOfServiceServiceClient is the client API for TermsOfServiceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TermsOfServiceServiceClient interface {
	// Retrieves the `TermsOfService` associated with the provided version.
	GetTermsOfService(ctx context.Context, in *GetTermsOfServiceRequest, opts ...grpc.CallOption) (*TermsOfService, error)
	// Retrieves the latest version of the `TermsOfService` for a given `kind` and
	// `region_code`.
	RetrieveLatestTermsOfService(ctx context.Context, in *RetrieveLatestTermsOfServiceRequest, opts ...grpc.CallOption) (*TermsOfService, error)
	// Accepts a `TermsOfService`. Executing this method requires admin access.
	AcceptTermsOfService(ctx context.Context, in *AcceptTermsOfServiceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type termsOfServiceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTermsOfServiceServiceClient(cc grpc.ClientConnInterface) TermsOfServiceServiceClient {
	return &termsOfServiceServiceClient{cc}
}

func (c *termsOfServiceServiceClient) GetTermsOfService(ctx context.Context, in *GetTermsOfServiceRequest, opts ...grpc.CallOption) (*TermsOfService, error) {
	out := new(TermsOfService)
	err := c.cc.Invoke(ctx, "/google.shopping.merchant.accounts.v1beta.TermsOfServiceService/GetTermsOfService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *termsOfServiceServiceClient) RetrieveLatestTermsOfService(ctx context.Context, in *RetrieveLatestTermsOfServiceRequest, opts ...grpc.CallOption) (*TermsOfService, error) {
	out := new(TermsOfService)
	err := c.cc.Invoke(ctx, "/google.shopping.merchant.accounts.v1beta.TermsOfServiceService/RetrieveLatestTermsOfService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *termsOfServiceServiceClient) AcceptTermsOfService(ctx context.Context, in *AcceptTermsOfServiceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/google.shopping.merchant.accounts.v1beta.TermsOfServiceService/AcceptTermsOfService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TermsOfServiceServiceServer is the server API for TermsOfServiceService service.
type TermsOfServiceServiceServer interface {
	// Retrieves the `TermsOfService` associated with the provided version.
	GetTermsOfService(context.Context, *GetTermsOfServiceRequest) (*TermsOfService, error)
	// Retrieves the latest version of the `TermsOfService` for a given `kind` and
	// `region_code`.
	RetrieveLatestTermsOfService(context.Context, *RetrieveLatestTermsOfServiceRequest) (*TermsOfService, error)
	// Accepts a `TermsOfService`. Executing this method requires admin access.
	AcceptTermsOfService(context.Context, *AcceptTermsOfServiceRequest) (*emptypb.Empty, error)
}

// UnimplementedTermsOfServiceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTermsOfServiceServiceServer struct {
}

func (*UnimplementedTermsOfServiceServiceServer) GetTermsOfService(context.Context, *GetTermsOfServiceRequest) (*TermsOfService, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTermsOfService not implemented")
}
func (*UnimplementedTermsOfServiceServiceServer) RetrieveLatestTermsOfService(context.Context, *RetrieveLatestTermsOfServiceRequest) (*TermsOfService, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveLatestTermsOfService not implemented")
}
func (*UnimplementedTermsOfServiceServiceServer) AcceptTermsOfService(context.Context, *AcceptTermsOfServiceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptTermsOfService not implemented")
}

func RegisterTermsOfServiceServiceServer(s *grpc.Server, srv TermsOfServiceServiceServer) {
	s.RegisterService(&_TermsOfServiceService_serviceDesc, srv)
}

func _TermsOfServiceService_GetTermsOfService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTermsOfServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TermsOfServiceServiceServer).GetTermsOfService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.shopping.merchant.accounts.v1beta.TermsOfServiceService/GetTermsOfService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TermsOfServiceServiceServer).GetTermsOfService(ctx, req.(*GetTermsOfServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TermsOfServiceService_RetrieveLatestTermsOfService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveLatestTermsOfServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TermsOfServiceServiceServer).RetrieveLatestTermsOfService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.shopping.merchant.accounts.v1beta.TermsOfServiceService/RetrieveLatestTermsOfService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TermsOfServiceServiceServer).RetrieveLatestTermsOfService(ctx, req.(*RetrieveLatestTermsOfServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TermsOfServiceService_AcceptTermsOfService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptTermsOfServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TermsOfServiceServiceServer).AcceptTermsOfService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.shopping.merchant.accounts.v1beta.TermsOfServiceService/AcceptTermsOfService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TermsOfServiceServiceServer).AcceptTermsOfService(ctx, req.(*AcceptTermsOfServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TermsOfServiceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.shopping.merchant.accounts.v1beta.TermsOfServiceService",
	HandlerType: (*TermsOfServiceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTermsOfService",
			Handler:    _TermsOfServiceService_GetTermsOfService_Handler,
		},
		{
			MethodName: "RetrieveLatestTermsOfService",
			Handler:    _TermsOfServiceService_RetrieveLatestTermsOfService_Handler,
		},
		{
			MethodName: "AcceptTermsOfService",
			Handler:    _TermsOfServiceService_AcceptTermsOfService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/shopping/merchant/accounts/v1beta/termsofservice.proto",
}
