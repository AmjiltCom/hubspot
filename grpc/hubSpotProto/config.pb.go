// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: config.proto

package hubspot_integration

import (
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

type HubspotIntegration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CompanyID    string `protobuf:"bytes,2,opt,name=CompanyID,proto3" json:"CompanyID,omitempty"`
	AccessToken  string `protobuf:"bytes,3,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	SyncPartner  bool   `protobuf:"varint,4,opt,name=SyncPartner,proto3" json:"SyncPartner,omitempty"`
	SyncProduct  bool   `protobuf:"varint,5,opt,name=SyncProduct,proto3" json:"SyncProduct,omitempty"`
	SyncPosOrder bool   `protobuf:"varint,6,opt,name=SyncPosOrder,proto3" json:"SyncPosOrder,omitempty"`
	SyncSales    bool   `protobuf:"varint,7,opt,name=SyncSales,proto3" json:"SyncSales,omitempty"`
	CreatedAt    string `protobuf:"bytes,8,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt    string `protobuf:"bytes,9,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *HubspotIntegration) Reset() {
	*x = HubspotIntegration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HubspotIntegration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HubspotIntegration) ProtoMessage() {}

func (x *HubspotIntegration) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HubspotIntegration.ProtoReflect.Descriptor instead.
func (*HubspotIntegration) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

func (x *HubspotIntegration) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *HubspotIntegration) GetCompanyID() string {
	if x != nil {
		return x.CompanyID
	}
	return ""
}

func (x *HubspotIntegration) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *HubspotIntegration) GetSyncPartner() bool {
	if x != nil {
		return x.SyncPartner
	}
	return false
}

func (x *HubspotIntegration) GetSyncProduct() bool {
	if x != nil {
		return x.SyncProduct
	}
	return false
}

func (x *HubspotIntegration) GetSyncPosOrder() bool {
	if x != nil {
		return x.SyncPosOrder
	}
	return false
}

func (x *HubspotIntegration) GetSyncSales() bool {
	if x != nil {
		return x.SyncSales
	}
	return false
}

func (x *HubspotIntegration) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *HubspotIntegration) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type ConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyID string `protobuf:"bytes,1,opt,name=CompanyID,proto3" json:"CompanyID,omitempty"`
}

func (x *ConfigRequest) Reset() {
	*x = ConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest) ProtoMessage() {}

func (x *ConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest.ProtoReflect.Descriptor instead.
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{1}
}

func (x *ConfigRequest) GetCompanyID() string {
	if x != nil {
		return x.CompanyID
	}
	return ""
}

type ConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  bool                `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Message string              `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Config  *HubspotIntegration `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *ConfigResponse) Reset() {
	*x = ConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigResponse) ProtoMessage() {}

func (x *ConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigResponse.ProtoReflect.Descriptor instead.
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *ConfigResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ConfigResponse) GetConfig() *HubspotIntegration {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x68, 0x75, 0x62, 0x73, 0x70, 0x6f, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xa6, 0x02, 0x0a, 0x12, 0x48, 0x75, 0x62, 0x73, 0x70, 0x6f, 0x74, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x79,
	0x6e, 0x63, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b,
	0x53, 0x79, 0x6e, 0x63, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6f, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6f, 0x73, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x61, 0x6c, 0x65, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2d, 0x0a, 0x0d,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x22, 0x83, 0x01, 0x0a, 0x0e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x3f, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x68, 0x75, 0x62, 0x73, 0x70, 0x6f, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x48, 0x75, 0x62, 0x73, 0x70, 0x6f, 0x74, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x32, 0x61, 0x0a, 0x07, 0x48, 0x75, 0x62, 0x53, 0x70, 0x6f, 0x74, 0x12, 0x56, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x2e, 0x68, 0x75, 0x62, 0x73,
	0x70, 0x6f, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x68, 0x75, 0x62, 0x73, 0x70, 0x6f, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x2e, 0x2f, 0x3b, 0x68, 0x75, 0x62, 0x73, 0x70,
	0x6f, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_config_proto_goTypes = []interface{}{
	(*HubspotIntegration)(nil), // 0: hubspot_integration.HubspotIntegration
	(*ConfigRequest)(nil),      // 1: hubspot_integration.ConfigRequest
	(*ConfigResponse)(nil),     // 2: hubspot_integration.ConfigResponse
}
var file_config_proto_depIdxs = []int32{
	0, // 0: hubspot_integration.ConfigResponse.config:type_name -> hubspot_integration.HubspotIntegration
	1, // 1: hubspot_integration.HubSpot.GetConfig:input_type -> hubspot_integration.ConfigRequest
	2, // 2: hubspot_integration.HubSpot.GetConfig:output_type -> hubspot_integration.ConfigResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HubspotIntegration); i {
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
		file_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest); i {
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
		file_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigResponse); i {
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
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}
