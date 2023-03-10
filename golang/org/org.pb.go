// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: org.proto

package org

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

type OrgPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *OrgPayload) Reset() {
	*x = OrgPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_org_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgPayload) ProtoMessage() {}

func (x *OrgPayload) ProtoReflect() protoreflect.Message {
	mi := &file_org_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgPayload.ProtoReflect.Descriptor instead.
func (*OrgPayload) Descriptor() ([]byte, []int) {
	return file_org_proto_rawDescGZIP(), []int{0}
}

func (x *OrgPayload) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrgPayload) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrgPayload) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type OrgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OrgResponse) Reset() {
	*x = OrgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_org_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgResponse) ProtoMessage() {}

func (x *OrgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_org_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgResponse.ProtoReflect.Descriptor instead.
func (*OrgResponse) Descriptor() ([]byte, []int) {
	return file_org_proto_rawDescGZIP(), []int{1}
}

func (x *OrgResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_org_proto protoreflect.FileDescriptor

var file_org_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x0a, 0x4f,
	0x72, 0x67, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x1d, 0x0a, 0x0b, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xa1,
	0x01, 0x0a, 0x03, 0x4f, 0x72, 0x67, 0x12, 0x26, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x67, 0x12, 0x0b, 0x2e, 0x4f, 0x72, 0x67, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x1a, 0x0c, 0x2e, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x12, 0x0b, 0x2e, 0x4f, 0x72, 0x67, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x0b, 0x2e, 0x4f, 0x72, 0x67, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x26, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x12,
	0x0b, 0x2e, 0x4f, 0x72, 0x67, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x0c, 0x2e, 0x4f,
	0x72, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x09, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x12, 0x0b, 0x2e, 0x4f, 0x72, 0x67, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x0c, 0x2e, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x76, 0x69, 0x6a, 0x65, 0x79, 0x61, 0x73, 0x68, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6f, 0x72, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_org_proto_rawDescOnce sync.Once
	file_org_proto_rawDescData = file_org_proto_rawDesc
)

func file_org_proto_rawDescGZIP() []byte {
	file_org_proto_rawDescOnce.Do(func() {
		file_org_proto_rawDescData = protoimpl.X.CompressGZIP(file_org_proto_rawDescData)
	})
	return file_org_proto_rawDescData
}

var file_org_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_org_proto_goTypes = []interface{}{
	(*OrgPayload)(nil),  // 0: OrgPayload
	(*OrgResponse)(nil), // 1: OrgResponse
}
var file_org_proto_depIdxs = []int32{
	0, // 0: Org.CreateOrg:input_type -> OrgPayload
	0, // 1: Org.GetOrg:input_type -> OrgPayload
	0, // 2: Org.UpdateOrg:input_type -> OrgPayload
	0, // 3: Org.DeleteOrg:input_type -> OrgPayload
	1, // 4: Org.CreateOrg:output_type -> OrgResponse
	0, // 5: Org.GetOrg:output_type -> OrgPayload
	1, // 6: Org.UpdateOrg:output_type -> OrgResponse
	1, // 7: Org.DeleteOrg:output_type -> OrgResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_org_proto_init() }
func file_org_proto_init() {
	if File_org_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_org_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgPayload); i {
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
		file_org_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgResponse); i {
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
			RawDescriptor: file_org_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_org_proto_goTypes,
		DependencyIndexes: file_org_proto_depIdxs,
		MessageInfos:      file_org_proto_msgTypes,
	}.Build()
	File_org_proto = out.File
	file_org_proto_rawDesc = nil
	file_org_proto_goTypes = nil
	file_org_proto_depIdxs = nil
}
