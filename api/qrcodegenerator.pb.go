// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: lorawan-stack/api/qrcodegenerator.proto

package api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

type QRCodeFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The entity fields required to generate the QR code.
	FieldMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
}

func (x *QRCodeFormat) Reset() {
	*x = QRCodeFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lorawan_stack_api_qrcodegenerator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QRCodeFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QRCodeFormat) ProtoMessage() {}

func (x *QRCodeFormat) ProtoReflect() protoreflect.Message {
	mi := &file_lorawan_stack_api_qrcodegenerator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QRCodeFormat.ProtoReflect.Descriptor instead.
func (*QRCodeFormat) Descriptor() ([]byte, []int) {
	return file_lorawan_stack_api_qrcodegenerator_proto_rawDescGZIP(), []int{0}
}

func (x *QRCodeFormat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *QRCodeFormat) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *QRCodeFormat) GetFieldMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FieldMask
	}
	return nil
}

type QRCodeFormats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Formats map[string]*QRCodeFormat `protobuf:"bytes,1,rep,name=formats,proto3" json:"formats,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *QRCodeFormats) Reset() {
	*x = QRCodeFormats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lorawan_stack_api_qrcodegenerator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QRCodeFormats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QRCodeFormats) ProtoMessage() {}

func (x *QRCodeFormats) ProtoReflect() protoreflect.Message {
	mi := &file_lorawan_stack_api_qrcodegenerator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QRCodeFormats.ProtoReflect.Descriptor instead.
func (*QRCodeFormats) Descriptor() ([]byte, []int) {
	return file_lorawan_stack_api_qrcodegenerator_proto_rawDescGZIP(), []int{1}
}

func (x *QRCodeFormats) GetFormats() map[string]*QRCodeFormat {
	if x != nil {
		return x.Formats
	}
	return nil
}

type GetQRCodeFormatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FormatId string `protobuf:"bytes,1,opt,name=format_id,json=formatId,proto3" json:"format_id,omitempty"`
}

func (x *GetQRCodeFormatRequest) Reset() {
	*x = GetQRCodeFormatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lorawan_stack_api_qrcodegenerator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQRCodeFormatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQRCodeFormatRequest) ProtoMessage() {}

func (x *GetQRCodeFormatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lorawan_stack_api_qrcodegenerator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQRCodeFormatRequest.ProtoReflect.Descriptor instead.
func (*GetQRCodeFormatRequest) Descriptor() ([]byte, []int) {
	return file_lorawan_stack_api_qrcodegenerator_proto_rawDescGZIP(), []int{2}
}

func (x *GetQRCodeFormatRequest) GetFormatId() string {
	if x != nil {
		return x.FormatId
	}
	return ""
}

type GenerateEndDeviceQRCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FormatId  string                                `protobuf:"bytes,1,opt,name=format_id,json=formatId,proto3" json:"format_id,omitempty"`
	EndDevice *EndDevice                            `protobuf:"bytes,2,opt,name=end_device,json=endDevice,proto3" json:"end_device,omitempty"`
	Image     *GenerateEndDeviceQRCodeRequest_Image `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *GenerateEndDeviceQRCodeRequest) Reset() {
	*x = GenerateEndDeviceQRCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lorawan_stack_api_qrcodegenerator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateEndDeviceQRCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateEndDeviceQRCodeRequest) ProtoMessage() {}

func (x *GenerateEndDeviceQRCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lorawan_stack_api_qrcodegenerator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateEndDeviceQRCodeRequest.ProtoReflect.Descriptor instead.
func (*GenerateEndDeviceQRCodeRequest) Descriptor() ([]byte, []int) {
	return file_lorawan_stack_api_qrcodegenerator_proto_rawDescGZIP(), []int{3}
}

func (x *GenerateEndDeviceQRCodeRequest) GetFormatId() string {
	if x != nil {
		return x.FormatId
	}
	return ""
}

func (x *GenerateEndDeviceQRCodeRequest) GetEndDevice() *EndDevice {
	if x != nil {
		return x.EndDevice
	}
	return nil
}

func (x *GenerateEndDeviceQRCodeRequest) GetImage() *GenerateEndDeviceQRCodeRequest_Image {
	if x != nil {
		return x.Image
	}
	return nil
}

type GenerateQRCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// QR code in PNG format, if requested.
	Image *Picture `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *GenerateQRCodeResponse) Reset() {
	*x = GenerateQRCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lorawan_stack_api_qrcodegenerator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateQRCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateQRCodeResponse) ProtoMessage() {}

func (x *GenerateQRCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lorawan_stack_api_qrcodegenerator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateQRCodeResponse.ProtoReflect.Descriptor instead.
func (*GenerateQRCodeResponse) Descriptor() ([]byte, []int) {
	return file_lorawan_stack_api_qrcodegenerator_proto_rawDescGZIP(), []int{4}
}

func (x *GenerateQRCodeResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *GenerateQRCodeResponse) GetImage() *Picture {
	if x != nil {
		return x.Image
	}
	return nil
}

type GenerateEndDeviceQRCodeRequest_Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageSize uint32 `protobuf:"varint,1,opt,name=image_size,json=imageSize,proto3" json:"image_size,omitempty"`
}

func (x *GenerateEndDeviceQRCodeRequest_Image) Reset() {
	*x = GenerateEndDeviceQRCodeRequest_Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lorawan_stack_api_qrcodegenerator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateEndDeviceQRCodeRequest_Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateEndDeviceQRCodeRequest_Image) ProtoMessage() {}

func (x *GenerateEndDeviceQRCodeRequest_Image) ProtoReflect() protoreflect.Message {
	mi := &file_lorawan_stack_api_qrcodegenerator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateEndDeviceQRCodeRequest_Image.ProtoReflect.Descriptor instead.
func (*GenerateEndDeviceQRCodeRequest_Image) Descriptor() ([]byte, []int) {
	return file_lorawan_stack_api_qrcodegenerator_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GenerateEndDeviceQRCodeRequest_Image) GetImageSize() uint32 {
	if x != nil {
		return x.ImageSize
	}
	return 0
}

var File_lorawan_stack_api_qrcodegenerator_proto protoreflect.FileDescriptor

var file_lorawan_stack_api_qrcodegenerator_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x74, 0x6e, 0x2e, 0x6c,
	0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61,
	0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x64, 0x5f,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6c, 0x6f,
	0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01,
	0x0a, 0x0c, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x1b,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x18, 0x64, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0xc8, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x09, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x3a, 0x04, 0xb8, 0xa0, 0x1f, 0x00, 0x22, 0xe3,
	0x01, 0x0a, 0x0d, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x73,
	0x12, 0x72, 0x0a, 0x07, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e,
	0x76, 0x33, 0x2e, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x73,
	0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x2c, 0xfa,
	0x42, 0x29, 0x9a, 0x01, 0x26, 0x22, 0x24, 0x72, 0x22, 0x18, 0x24, 0x32, 0x1e, 0x5e, 0x5b, 0x61,
	0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x28, 0x3f, 0x3a, 0x5b, 0x2d, 0x5d, 0x3f, 0x5b, 0x61, 0x2d,
	0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x7b, 0x32, 0x2c, 0x7d, 0x24, 0x52, 0x07, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x73, 0x1a, 0x58, 0x0a, 0x0c, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61,
	0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x04,
	0xb8, 0xa0, 0x1f, 0x00, 0x22, 0x70, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x51, 0x52, 0x43, 0x6f, 0x64,
	0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x50,
	0x0a, 0x09, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x33, 0xe2, 0xde, 0x1f, 0x08, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x49, 0x44, 0xfa,
	0x42, 0x24, 0x72, 0x22, 0x18, 0x24, 0x32, 0x1e, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39,
	0x5d, 0x28, 0x3f, 0x3a, 0x5b, 0x2d, 0x5d, 0x3f, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d,
	0x29, 0x7b, 0x32, 0x2c, 0x7d, 0x24, 0x52, 0x08, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x49, 0x64,
	0x3a, 0x04, 0xb8, 0xa0, 0x1f, 0x00, 0x22, 0xc0, 0x02, 0x0a, 0x1e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x51, 0x52, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x09, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xe2, 0xde,
	0x1f, 0x08, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x49, 0x44, 0xfa, 0x42, 0x24, 0x72, 0x22, 0x18,
	0x24, 0x32, 0x1e, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x28, 0x3f, 0x3a, 0x5b,
	0x2d, 0x5d, 0x3f, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x7b, 0x32, 0x2c, 0x7d,
	0x24, 0x52, 0x08, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x0a, 0x65,
	0x6e, 0x64, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33,
	0x2e, 0x45, 0x6e, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x0c, 0xc8, 0xde, 0x1f, 0x00,
	0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x34, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e,
	0x2e, 0x76, 0x33, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x1a,
	0x32, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0a, 0xfa, 0x42,
	0x07, 0x2a, 0x05, 0x18, 0xe8, 0x07, 0x28, 0x0a, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x3a, 0x04, 0xb8, 0xa0, 0x1f, 0x00, 0x22, 0x61, 0x0a, 0x16, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72,
	0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x3a, 0x04, 0xb8, 0xa0, 0x1f, 0x00, 0x32, 0x95, 0x03, 0x0a,
	0x18, 0x45, 0x6e, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x84, 0x01, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x26, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f,
	0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x52, 0x43, 0x6f,
	0x64, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33,
	0x2e, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x31, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x12, 0x29, 0x2f, 0x71, 0x72, 0x2d, 0x63, 0x6f, 0x64, 0x65, 0x73,
	0x2f, 0x65, 0x6e, 0x64, 0x2d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x73, 0x2f, 0x7b, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x7d,
	0x12, 0x6b, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x73, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f,
	0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x73, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d,
	0x2f, 0x71, 0x72, 0x2d, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2f, 0x65, 0x6e, 0x64, 0x2d, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x73, 0x12, 0x84, 0x01,
	0x0a, 0x08, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x2e, 0x74, 0x74, 0x6e,
	0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x51, 0x52, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x74, 0x74, 0x6e,
	0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x71, 0x72, 0x2d,
	0x63, 0x6f, 0x64, 0x65, 0x73, 0x2f, 0x65, 0x6e, 0x64, 0x2d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x3a, 0x01, 0x2a, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x6f, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6c, 0x6f, 0x72,
	0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lorawan_stack_api_qrcodegenerator_proto_rawDescOnce sync.Once
	file_lorawan_stack_api_qrcodegenerator_proto_rawDescData = file_lorawan_stack_api_qrcodegenerator_proto_rawDesc
)

func file_lorawan_stack_api_qrcodegenerator_proto_rawDescGZIP() []byte {
	file_lorawan_stack_api_qrcodegenerator_proto_rawDescOnce.Do(func() {
		file_lorawan_stack_api_qrcodegenerator_proto_rawDescData = protoimpl.X.CompressGZIP(file_lorawan_stack_api_qrcodegenerator_proto_rawDescData)
	})
	return file_lorawan_stack_api_qrcodegenerator_proto_rawDescData
}

var file_lorawan_stack_api_qrcodegenerator_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_lorawan_stack_api_qrcodegenerator_proto_goTypes = []interface{}{
	(*QRCodeFormat)(nil),                         // 0: ttn.lorawan.v3.QRCodeFormat
	(*QRCodeFormats)(nil),                        // 1: ttn.lorawan.v3.QRCodeFormats
	(*GetQRCodeFormatRequest)(nil),               // 2: ttn.lorawan.v3.GetQRCodeFormatRequest
	(*GenerateEndDeviceQRCodeRequest)(nil),       // 3: ttn.lorawan.v3.GenerateEndDeviceQRCodeRequest
	(*GenerateQRCodeResponse)(nil),               // 4: ttn.lorawan.v3.GenerateQRCodeResponse
	nil,                                          // 5: ttn.lorawan.v3.QRCodeFormats.FormatsEntry
	(*GenerateEndDeviceQRCodeRequest_Image)(nil), // 6: ttn.lorawan.v3.GenerateEndDeviceQRCodeRequest.Image
	(*fieldmaskpb.FieldMask)(nil),                // 7: google.protobuf.FieldMask
	(*EndDevice)(nil),                            // 8: ttn.lorawan.v3.EndDevice
	(*Picture)(nil),                              // 9: ttn.lorawan.v3.Picture
	(*emptypb.Empty)(nil),                        // 10: google.protobuf.Empty
}
var file_lorawan_stack_api_qrcodegenerator_proto_depIdxs = []int32{
	7,  // 0: ttn.lorawan.v3.QRCodeFormat.field_mask:type_name -> google.protobuf.FieldMask
	5,  // 1: ttn.lorawan.v3.QRCodeFormats.formats:type_name -> ttn.lorawan.v3.QRCodeFormats.FormatsEntry
	8,  // 2: ttn.lorawan.v3.GenerateEndDeviceQRCodeRequest.end_device:type_name -> ttn.lorawan.v3.EndDevice
	6,  // 3: ttn.lorawan.v3.GenerateEndDeviceQRCodeRequest.image:type_name -> ttn.lorawan.v3.GenerateEndDeviceQRCodeRequest.Image
	9,  // 4: ttn.lorawan.v3.GenerateQRCodeResponse.image:type_name -> ttn.lorawan.v3.Picture
	0,  // 5: ttn.lorawan.v3.QRCodeFormats.FormatsEntry.value:type_name -> ttn.lorawan.v3.QRCodeFormat
	2,  // 6: ttn.lorawan.v3.EndDeviceQRCodeGenerator.GetFormat:input_type -> ttn.lorawan.v3.GetQRCodeFormatRequest
	10, // 7: ttn.lorawan.v3.EndDeviceQRCodeGenerator.ListFormats:input_type -> google.protobuf.Empty
	3,  // 8: ttn.lorawan.v3.EndDeviceQRCodeGenerator.Generate:input_type -> ttn.lorawan.v3.GenerateEndDeviceQRCodeRequest
	0,  // 9: ttn.lorawan.v3.EndDeviceQRCodeGenerator.GetFormat:output_type -> ttn.lorawan.v3.QRCodeFormat
	1,  // 10: ttn.lorawan.v3.EndDeviceQRCodeGenerator.ListFormats:output_type -> ttn.lorawan.v3.QRCodeFormats
	4,  // 11: ttn.lorawan.v3.EndDeviceQRCodeGenerator.Generate:output_type -> ttn.lorawan.v3.GenerateQRCodeResponse
	9,  // [9:12] is the sub-list for method output_type
	6,  // [6:9] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_lorawan_stack_api_qrcodegenerator_proto_init() }
func file_lorawan_stack_api_qrcodegenerator_proto_init() {
	if File_lorawan_stack_api_qrcodegenerator_proto != nil {
		return
	}
	file_lorawan_stack_api_end_device_proto_init()
	file_lorawan_stack_api_picture_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_lorawan_stack_api_qrcodegenerator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QRCodeFormat); i {
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
		file_lorawan_stack_api_qrcodegenerator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QRCodeFormats); i {
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
		file_lorawan_stack_api_qrcodegenerator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQRCodeFormatRequest); i {
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
		file_lorawan_stack_api_qrcodegenerator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateEndDeviceQRCodeRequest); i {
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
		file_lorawan_stack_api_qrcodegenerator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateQRCodeResponse); i {
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
		file_lorawan_stack_api_qrcodegenerator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateEndDeviceQRCodeRequest_Image); i {
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
			RawDescriptor: file_lorawan_stack_api_qrcodegenerator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lorawan_stack_api_qrcodegenerator_proto_goTypes,
		DependencyIndexes: file_lorawan_stack_api_qrcodegenerator_proto_depIdxs,
		MessageInfos:      file_lorawan_stack_api_qrcodegenerator_proto_msgTypes,
	}.Build()
	File_lorawan_stack_api_qrcodegenerator_proto = out.File
	file_lorawan_stack_api_qrcodegenerator_proto_rawDesc = nil
	file_lorawan_stack_api_qrcodegenerator_proto_goTypes = nil
	file_lorawan_stack_api_qrcodegenerator_proto_depIdxs = nil
}
