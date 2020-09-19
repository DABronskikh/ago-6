// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: fine.proto

package v1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AutopaysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *AutopaysRequest) Reset() {
	*x = AutopaysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutopaysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutopaysRequest) ProtoMessage() {}

func (x *AutopaysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutopaysRequest.ProtoReflect.Descriptor instead.
func (*AutopaysRequest) Descriptor() ([]byte, []int) {
	return file_fine_proto_rawDescGZIP(), []int{0}
}

func (x *AutopaysRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type AutopayRequestById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AutopayId int64 `protobuf:"varint,1,opt,name=autopayId,proto3" json:"autopayId,omitempty"`
}

func (x *AutopayRequestById) Reset() {
	*x = AutopayRequestById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutopayRequestById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutopayRequestById) ProtoMessage() {}

func (x *AutopayRequestById) ProtoReflect() protoreflect.Message {
	mi := &file_fine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutopayRequestById.ProtoReflect.Descriptor instead.
func (*AutopayRequestById) Descriptor() ([]byte, []int) {
	return file_fine_proto_rawDescGZIP(), []int{1}
}

func (x *AutopayRequestById) GetAutopayId() int64 {
	if x != nil {
		return x.AutopayId
	}
	return 0
}

type AutopayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Autopay `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *AutopayResponse) Reset() {
	*x = AutopayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutopayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutopayResponse) ProtoMessage() {}

func (x *AutopayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutopayResponse.ProtoReflect.Descriptor instead.
func (*AutopayResponse) Descriptor() ([]byte, []int) {
	return file_fine_proto_rawDescGZIP(), []int{2}
}

func (x *AutopayResponse) GetItems() []*Autopay {
	if x != nil {
		return x.Items
	}
	return nil
}

type Autopay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`          //id
	Name    string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`       //название
	Phone   string               `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`     //номер телефона
	Created *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"` //дата создания
	Updated *timestamp.Timestamp `protobuf:"bytes,5,opt,name=updated,proto3" json:"updated,omitempty"` //дата обновления
}

func (x *Autopay) Reset() {
	*x = Autopay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Autopay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Autopay) ProtoMessage() {}

func (x *Autopay) ProtoReflect() protoreflect.Message {
	mi := &file_fine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Autopay.ProtoReflect.Descriptor instead.
func (*Autopay) Descriptor() ([]byte, []int) {
	return file_fine_proto_rawDescGZIP(), []int{3}
}

func (x *Autopay) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Autopay) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Autopay) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Autopay) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Autopay) GetUpdated() *timestamp.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

var File_fine_proto protoreflect.FileDescriptor

var file_fine_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x70, 0x6b,
	0x67, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x29, 0x0a, 0x0f, 0x41, 0x75, 0x74, 0x6f, 0x70, 0x61, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x12, 0x41, 0x75,
	0x74, 0x6f, 0x70, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x6f, 0x70, 0x61, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x75, 0x74, 0x6f, 0x70, 0x61, 0x79, 0x49, 0x64, 0x22, 0x42,
	0x0a, 0x0f, 0x41, 0x75, 0x74, 0x6f, 0x70, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x69, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x70, 0x61, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0xaf, 0x01, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x6f, 0x70, 0x61, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34,
	0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x32, 0xf0, 0x02, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x6f, 0x70, 0x61, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x19, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x70, 0x61, 0x79, 0x1a, 0x19, 0x2e, 0x70,
	0x6b, 0x67, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x75, 0x74, 0x6f, 0x70, 0x61, 0x79, 0x12, 0x51, 0x0a, 0x09, 0x46, 0x69, 0x6e, 0x64, 0x42,
	0x79, 0x41, 0x6c, 0x6c, 0x12, 0x21, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x66, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x70, 0x61, 0x79, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x70,
	0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x08, 0x46, 0x69,
	0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x24, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x70, 0x61,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x19, 0x2e, 0x70,
	0x6b, 0x67, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x75, 0x74, 0x6f, 0x70, 0x61, 0x79, 0x12, 0x3e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x19, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x70, 0x61, 0x79, 0x1a, 0x19, 0x2e, 0x70,
	0x6b, 0x67, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x75, 0x74, 0x6f, 0x70, 0x61, 0x79, 0x12, 0x3e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x19, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x70, 0x61, 0x79, 0x1a, 0x19, 0x2e, 0x70,
	0x6b, 0x67, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x75, 0x74, 0x6f, 0x70, 0x61, 0x79, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x41, 0x42, 0x72, 0x6f, 0x6e, 0x73, 0x6b, 0x69, 0x6b,
	0x68, 0x2f, 0x61, 0x67, 0x6f, 0x2d, 0x36, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x69, 0x6e, 0x65,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fine_proto_rawDescOnce sync.Once
	file_fine_proto_rawDescData = file_fine_proto_rawDesc
)

func file_fine_proto_rawDescGZIP() []byte {
	file_fine_proto_rawDescOnce.Do(func() {
		file_fine_proto_rawDescData = protoimpl.X.CompressGZIP(file_fine_proto_rawDescData)
	})
	return file_fine_proto_rawDescData
}

var file_fine_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_fine_proto_goTypes = []interface{}{
	(*AutopaysRequest)(nil),     // 0: pkg.grpc.fine.v1.AutopaysRequest
	(*AutopayRequestById)(nil),  // 1: pkg.grpc.fine.v1.AutopayRequestById
	(*AutopayResponse)(nil),     // 2: pkg.grpc.fine.v1.AutopayResponse
	(*Autopay)(nil),             // 3: pkg.grpc.fine.v1.Autopay
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_fine_proto_depIdxs = []int32{
	3, // 0: pkg.grpc.fine.v1.AutopayResponse.items:type_name -> pkg.grpc.fine.v1.Autopay
	4, // 1: pkg.grpc.fine.v1.Autopay.created:type_name -> google.protobuf.Timestamp
	4, // 2: pkg.grpc.fine.v1.Autopay.updated:type_name -> google.protobuf.Timestamp
	3, // 3: pkg.grpc.fine.v1.AutopayService.Create:input_type -> pkg.grpc.fine.v1.Autopay
	0, // 4: pkg.grpc.fine.v1.AutopayService.FindByAll:input_type -> pkg.grpc.fine.v1.AutopaysRequest
	1, // 5: pkg.grpc.fine.v1.AutopayService.FindById:input_type -> pkg.grpc.fine.v1.AutopayRequestById
	3, // 6: pkg.grpc.fine.v1.AutopayService.Update:input_type -> pkg.grpc.fine.v1.Autopay
	3, // 7: pkg.grpc.fine.v1.AutopayService.Delete:input_type -> pkg.grpc.fine.v1.Autopay
	3, // 8: pkg.grpc.fine.v1.AutopayService.Create:output_type -> pkg.grpc.fine.v1.Autopay
	2, // 9: pkg.grpc.fine.v1.AutopayService.FindByAll:output_type -> pkg.grpc.fine.v1.AutopayResponse
	3, // 10: pkg.grpc.fine.v1.AutopayService.FindById:output_type -> pkg.grpc.fine.v1.Autopay
	3, // 11: pkg.grpc.fine.v1.AutopayService.Update:output_type -> pkg.grpc.fine.v1.Autopay
	3, // 12: pkg.grpc.fine.v1.AutopayService.Delete:output_type -> pkg.grpc.fine.v1.Autopay
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_fine_proto_init() }
func file_fine_proto_init() {
	if File_fine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutopaysRequest); i {
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
		file_fine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutopayRequestById); i {
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
		file_fine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutopayResponse); i {
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
		file_fine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Autopay); i {
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
			RawDescriptor: file_fine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fine_proto_goTypes,
		DependencyIndexes: file_fine_proto_depIdxs,
		MessageInfos:      file_fine_proto_msgTypes,
	}.Build()
	File_fine_proto = out.File
	file_fine_proto_rawDesc = nil
	file_fine_proto_goTypes = nil
	file_fine_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AutopayServiceClient is the client API for AutopayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AutopayServiceClient interface {
	Create(ctx context.Context, in *Autopay, opts ...grpc.CallOption) (*Autopay, error)
	FindByAll(ctx context.Context, in *AutopaysRequest, opts ...grpc.CallOption) (*AutopayResponse, error)
	FindById(ctx context.Context, in *AutopayRequestById, opts ...grpc.CallOption) (*Autopay, error)
	Update(ctx context.Context, in *Autopay, opts ...grpc.CallOption) (*Autopay, error)
	Delete(ctx context.Context, in *Autopay, opts ...grpc.CallOption) (*Autopay, error)
}

type autopayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAutopayServiceClient(cc grpc.ClientConnInterface) AutopayServiceClient {
	return &autopayServiceClient{cc}
}

func (c *autopayServiceClient) Create(ctx context.Context, in *Autopay, opts ...grpc.CallOption) (*Autopay, error) {
	out := new(Autopay)
	err := c.cc.Invoke(ctx, "/pkg.grpc.fine.v1.AutopayService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autopayServiceClient) FindByAll(ctx context.Context, in *AutopaysRequest, opts ...grpc.CallOption) (*AutopayResponse, error) {
	out := new(AutopayResponse)
	err := c.cc.Invoke(ctx, "/pkg.grpc.fine.v1.AutopayService/FindByAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autopayServiceClient) FindById(ctx context.Context, in *AutopayRequestById, opts ...grpc.CallOption) (*Autopay, error) {
	out := new(Autopay)
	err := c.cc.Invoke(ctx, "/pkg.grpc.fine.v1.AutopayService/FindById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autopayServiceClient) Update(ctx context.Context, in *Autopay, opts ...grpc.CallOption) (*Autopay, error) {
	out := new(Autopay)
	err := c.cc.Invoke(ctx, "/pkg.grpc.fine.v1.AutopayService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autopayServiceClient) Delete(ctx context.Context, in *Autopay, opts ...grpc.CallOption) (*Autopay, error) {
	out := new(Autopay)
	err := c.cc.Invoke(ctx, "/pkg.grpc.fine.v1.AutopayService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AutopayServiceServer is the server API for AutopayService service.
type AutopayServiceServer interface {
	Create(context.Context, *Autopay) (*Autopay, error)
	FindByAll(context.Context, *AutopaysRequest) (*AutopayResponse, error)
	FindById(context.Context, *AutopayRequestById) (*Autopay, error)
	Update(context.Context, *Autopay) (*Autopay, error)
	Delete(context.Context, *Autopay) (*Autopay, error)
}

// UnimplementedAutopayServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAutopayServiceServer struct {
}

func (*UnimplementedAutopayServiceServer) Create(context.Context, *Autopay) (*Autopay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedAutopayServiceServer) FindByAll(context.Context, *AutopaysRequest) (*AutopayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByAll not implemented")
}
func (*UnimplementedAutopayServiceServer) FindById(context.Context, *AutopayRequestById) (*Autopay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (*UnimplementedAutopayServiceServer) Update(context.Context, *Autopay) (*Autopay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedAutopayServiceServer) Delete(context.Context, *Autopay) (*Autopay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterAutopayServiceServer(s *grpc.Server, srv AutopayServiceServer) {
	s.RegisterService(&_AutopayService_serviceDesc, srv)
}

func _AutopayService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Autopay)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutopayServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.grpc.fine.v1.AutopayService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutopayServiceServer).Create(ctx, req.(*Autopay))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutopayService_FindByAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AutopaysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutopayServiceServer).FindByAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.grpc.fine.v1.AutopayService/FindByAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutopayServiceServer).FindByAll(ctx, req.(*AutopaysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutopayService_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AutopayRequestById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutopayServiceServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.grpc.fine.v1.AutopayService/FindById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutopayServiceServer).FindById(ctx, req.(*AutopayRequestById))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutopayService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Autopay)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutopayServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.grpc.fine.v1.AutopayService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutopayServiceServer).Update(ctx, req.(*Autopay))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutopayService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Autopay)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutopayServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.grpc.fine.v1.AutopayService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutopayServiceServer).Delete(ctx, req.(*Autopay))
	}
	return interceptor(ctx, in, info, handler)
}

var _AutopayService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pkg.grpc.fine.v1.AutopayService",
	HandlerType: (*AutopayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AutopayService_Create_Handler,
		},
		{
			MethodName: "FindByAll",
			Handler:    _AutopayService_FindByAll_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _AutopayService_FindById_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AutopayService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AutopayService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fine.proto",
}