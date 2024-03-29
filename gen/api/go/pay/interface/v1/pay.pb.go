// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        (unknown)
// source: pay/interface/v1/pay.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type GoPayReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel       string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	TransactionId string `protobuf:"bytes,2,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
}

func (x *GoPayReq) Reset() {
	*x = GoPayReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pay_interface_v1_pay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoPayReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoPayReq) ProtoMessage() {}

func (x *GoPayReq) ProtoReflect() protoreflect.Message {
	mi := &file_pay_interface_v1_pay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoPayReq.ProtoReflect.Descriptor instead.
func (*GoPayReq) Descriptor() ([]byte, []int) {
	return file_pay_interface_v1_pay_proto_rawDescGZIP(), []int{0}
}

func (x *GoPayReq) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *GoPayReq) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

type GoPayReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GoPayReply) Reset() {
	*x = GoPayReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pay_interface_v1_pay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoPayReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoPayReply) ProtoMessage() {}

func (x *GoPayReply) ProtoReflect() protoreflect.Message {
	mi := &file_pay_interface_v1_pay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoPayReply.ProtoReflect.Descriptor instead.
func (*GoPayReply) Descriptor() ([]byte, []int) {
	return file_pay_interface_v1_pay_proto_rawDescGZIP(), []int{1}
}

var File_pay_interface_v1_pay_proto protoreflect.FileDescriptor

var file_pay_interface_v1_pay_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x61, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a,
	0x0a, 0x08, 0x47, 0x6f, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x0c, 0x0a, 0x0a, 0x47, 0x6f,
	0x50, 0x61, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xde, 0x06, 0x0a, 0x03, 0x50, 0x61, 0x79,
	0x12, 0x5b, 0x0a, 0x05, 0x47, 0x6f, 0x50, 0x61, 0x79, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x61, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x6f, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61,
	0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x50,
	0x61, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a,
	0x01, 0x2a, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x70, 0x61, 0x79, 0x12, 0x5d, 0x0a,
	0x06, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61,
	0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x50,
	0x61, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x50, 0x61, 0x79,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a,
	0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x12, 0x7c, 0x0a, 0x0b,
	0x41, 0x73, 0x79, 0x6e, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x6f, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x61, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x6f, 0x50, 0x61, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x29, 0x3a, 0x01, 0x2a, 0x22, 0x24, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x2f, 0x7b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x7d, 0x2f, 0x7b, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x7d, 0x12, 0x7b, 0x0a, 0x0a, 0x53, 0x79,
	0x6e, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x61, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f,
	0x50, 0x61, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x50, 0x61,
	0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x01,
	0x2a, 0x22, 0x24, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x2f, 0x7b, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x7d, 0x2f, 0x7b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x7d, 0x12, 0x66, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x54, 0x72, 0x61, 0x64, 0x65, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x50, 0x61, 0x79,
	0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x50, 0x61, 0x79, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f,
	0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x12,
	0x68, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x12, 0x1c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x6f, 0x50, 0x61, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2f, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x12, 0x64, 0x0a, 0x09, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x50, 0x61,
	0x79, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x50, 0x61, 0x79, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22,
	0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x12,
	0x68, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x12, 0x1c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x6f, 0x50, 0x61, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x42, 0x4f, 0x0a, 0x14, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x61, 0x79, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x35, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c,
	0x61, 0x6c, 0x69, 0x66, 0x65, 0x69, 0x65, 0x72, 0x2f, 0x76, 0x76, 0x67, 0x6f, 0x2d, 0x6d, 0x61,
	0x6c, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pay_interface_v1_pay_proto_rawDescOnce sync.Once
	file_pay_interface_v1_pay_proto_rawDescData = file_pay_interface_v1_pay_proto_rawDesc
)

func file_pay_interface_v1_pay_proto_rawDescGZIP() []byte {
	file_pay_interface_v1_pay_proto_rawDescOnce.Do(func() {
		file_pay_interface_v1_pay_proto_rawDescData = protoimpl.X.CompressGZIP(file_pay_interface_v1_pay_proto_rawDescData)
	})
	return file_pay_interface_v1_pay_proto_rawDescData
}

var file_pay_interface_v1_pay_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pay_interface_v1_pay_proto_goTypes = []interface{}{
	(*GoPayReq)(nil),   // 0: api.pay.service.v1.GoPayReq
	(*GoPayReply)(nil), // 1: api.pay.service.v1.GoPayReply
}
var file_pay_interface_v1_pay_proto_depIdxs = []int32{
	0, // 0: api.pay.service.v1.Pay.GoPay:input_type -> api.pay.service.v1.GoPayReq
	0, // 1: api.pay.service.v1.Pay.Refund:input_type -> api.pay.service.v1.GoPayReq
	0, // 2: api.pay.service.v1.Pay.AsyncNotify:input_type -> api.pay.service.v1.GoPayReq
	0, // 3: api.pay.service.v1.Pay.SyncNotify:input_type -> api.pay.service.v1.GoPayReq
	0, // 4: api.pay.service.v1.Pay.QueryTrade:input_type -> api.pay.service.v1.GoPayReq
	0, // 5: api.pay.service.v1.Pay.QueryRefund:input_type -> api.pay.service.v1.GoPayReq
	0, // 6: api.pay.service.v1.Pay.QueryBill:input_type -> api.pay.service.v1.GoPayReq
	0, // 7: api.pay.service.v1.Pay.QuerySettle:input_type -> api.pay.service.v1.GoPayReq
	1, // 8: api.pay.service.v1.Pay.GoPay:output_type -> api.pay.service.v1.GoPayReply
	1, // 9: api.pay.service.v1.Pay.Refund:output_type -> api.pay.service.v1.GoPayReply
	1, // 10: api.pay.service.v1.Pay.AsyncNotify:output_type -> api.pay.service.v1.GoPayReply
	1, // 11: api.pay.service.v1.Pay.SyncNotify:output_type -> api.pay.service.v1.GoPayReply
	1, // 12: api.pay.service.v1.Pay.QueryTrade:output_type -> api.pay.service.v1.GoPayReply
	1, // 13: api.pay.service.v1.Pay.QueryRefund:output_type -> api.pay.service.v1.GoPayReply
	1, // 14: api.pay.service.v1.Pay.QueryBill:output_type -> api.pay.service.v1.GoPayReply
	1, // 15: api.pay.service.v1.Pay.QuerySettle:output_type -> api.pay.service.v1.GoPayReply
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pay_interface_v1_pay_proto_init() }
func file_pay_interface_v1_pay_proto_init() {
	if File_pay_interface_v1_pay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pay_interface_v1_pay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoPayReq); i {
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
		file_pay_interface_v1_pay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoPayReply); i {
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
			RawDescriptor: file_pay_interface_v1_pay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pay_interface_v1_pay_proto_goTypes,
		DependencyIndexes: file_pay_interface_v1_pay_proto_depIdxs,
		MessageInfos:      file_pay_interface_v1_pay_proto_msgTypes,
	}.Build()
	File_pay_interface_v1_pay_proto = out.File
	file_pay_interface_v1_pay_proto_rawDesc = nil
	file_pay_interface_v1_pay_proto_goTypes = nil
	file_pay_interface_v1_pay_proto_depIdxs = nil
}
