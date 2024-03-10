// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: shop/v1/order.proto

package shopv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version         int32                  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Id              string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	LastUpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=last_updated_at,json=lastUpdatedAt,proto3" json:"last_updated_at,omitempty"`
	DeliveredAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=delivered_at,json=deliveredAt,proto3" json:"delivered_at,omitempty"`
	CompletedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=completed_at,json=completedAt,proto3" json:"completed_at,omitempty"`
	Customer        *Customer              `protobuf:"bytes,7,opt,name=customer,proto3" json:"customer,omitempty"`
	OrderValue      int32                  `protobuf:"varint,8,opt,name=order_value,json=orderValue,proto3" json:"order_value,omitempty"`
	LineItems       []*Order_LineItem      `protobuf:"bytes,9,rep,name=line_items,json=lineItems,proto3" json:"line_items,omitempty"`
	Payment         *Order_Payment         `protobuf:"bytes,10,opt,name=payment,proto3" json:"payment,omitempty"`
	DeliveryAddress *Address               `protobuf:"bytes,11,opt,name=delivery_address,json=deliveryAddress,proto3" json:"delivery_address,omitempty"`
	Revision        int32                  `protobuf:"varint,12,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_v1_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_shop_v1_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_shop_v1_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Order) GetLastUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdatedAt
	}
	return nil
}

func (x *Order) GetDeliveredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeliveredAt
	}
	return nil
}

func (x *Order) GetCompletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CompletedAt
	}
	return nil
}

func (x *Order) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

func (x *Order) GetOrderValue() int32 {
	if x != nil {
		return x.OrderValue
	}
	return 0
}

func (x *Order) GetLineItems() []*Order_LineItem {
	if x != nil {
		return x.LineItems
	}
	return nil
}

func (x *Order) GetPayment() *Order_Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

func (x *Order) GetDeliveryAddress() *Address {
	if x != nil {
		return x.DeliveryAddress
	}
	return nil
}

func (x *Order) GetRevision() int32 {
	if x != nil {
		return x.Revision
	}
	return 0
}

type Order_LineItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArticleId    string `protobuf:"bytes,1,opt,name=article_id,json=articleId,proto3" json:"article_id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Quantity     int32  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	QuantityUnit string `protobuf:"bytes,4,opt,name=quantity_unit,json=quantityUnit,proto3" json:"quantity_unit,omitempty"`
	UnitPrice    int32  `protobuf:"varint,5,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"`
	TotalPrice   int32  `protobuf:"varint,6,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
}

func (x *Order_LineItem) Reset() {
	*x = Order_LineItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_v1_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order_LineItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order_LineItem) ProtoMessage() {}

func (x *Order_LineItem) ProtoReflect() protoreflect.Message {
	mi := &file_shop_v1_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order_LineItem.ProtoReflect.Descriptor instead.
func (*Order_LineItem) Descriptor() ([]byte, []int) {
	return file_shop_v1_order_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Order_LineItem) GetArticleId() string {
	if x != nil {
		return x.ArticleId
	}
	return ""
}

func (x *Order_LineItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Order_LineItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Order_LineItem) GetQuantityUnit() string {
	if x != nil {
		return x.QuantityUnit
	}
	return ""
}

func (x *Order_LineItem) GetUnitPrice() int32 {
	if x != nil {
		return x.UnitPrice
	}
	return 0
}

func (x *Order_LineItem) GetTotalPrice() int32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

type Order_Payment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentId string `protobuf:"bytes,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	Method    string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *Order_Payment) Reset() {
	*x = Order_Payment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_v1_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order_Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order_Payment) ProtoMessage() {}

func (x *Order_Payment) ProtoReflect() protoreflect.Message {
	mi := &file_shop_v1_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order_Payment.ProtoReflect.Descriptor instead.
func (*Order_Payment) Descriptor() ([]byte, []int) {
	return file_shop_v1_order_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Order_Payment) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

func (x *Order_Payment) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

var File_shop_v1_order_proto protoreflect.FileDescriptor

var file_shop_v1_order_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4,
	0x06, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x42, 0x0a,
	0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0b, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x2d, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1f,
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x36, 0x0a, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x6c, 0x69,
	0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x30, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x10, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x1a, 0xbe, 0x01, 0x0a, 0x08, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23,
	0x0a, 0x0d, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x55,
	0x6e, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x1a, 0x40, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x90, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x68,
	0x6f, 0x70, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x68, 0x75, 0x74, 0x2f, 0x6f, 0x77, 0x6c, 0x2d, 0x73, 0x68, 0x6f,
	0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x73,
	0x68, 0x6f, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x68, 0x6f, 0x70, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x53, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x53, 0x68, 0x6f, 0x70, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07,
	0x53, 0x68, 0x6f, 0x70, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x53, 0x68, 0x6f, 0x70, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08,
	0x53, 0x68, 0x6f, 0x70, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shop_v1_order_proto_rawDescOnce sync.Once
	file_shop_v1_order_proto_rawDescData = file_shop_v1_order_proto_rawDesc
)

func file_shop_v1_order_proto_rawDescGZIP() []byte {
	file_shop_v1_order_proto_rawDescOnce.Do(func() {
		file_shop_v1_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_shop_v1_order_proto_rawDescData)
	})
	return file_shop_v1_order_proto_rawDescData
}

var file_shop_v1_order_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_shop_v1_order_proto_goTypes = []interface{}{
	(*Order)(nil),                 // 0: shop.v1.Order
	(*Order_LineItem)(nil),        // 1: shop.v1.Order.LineItem
	(*Order_Payment)(nil),         // 2: shop.v1.Order.Payment
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*Customer)(nil),              // 4: shop.v1.Customer
	(*Address)(nil),               // 5: shop.v1.Address
}
var file_shop_v1_order_proto_depIdxs = []int32{
	3, // 0: shop.v1.Order.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: shop.v1.Order.last_updated_at:type_name -> google.protobuf.Timestamp
	3, // 2: shop.v1.Order.delivered_at:type_name -> google.protobuf.Timestamp
	3, // 3: shop.v1.Order.completed_at:type_name -> google.protobuf.Timestamp
	4, // 4: shop.v1.Order.customer:type_name -> shop.v1.Customer
	1, // 5: shop.v1.Order.line_items:type_name -> shop.v1.Order.LineItem
	2, // 6: shop.v1.Order.payment:type_name -> shop.v1.Order.Payment
	5, // 7: shop.v1.Order.delivery_address:type_name -> shop.v1.Address
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_shop_v1_order_proto_init() }
func file_shop_v1_order_proto_init() {
	if File_shop_v1_order_proto != nil {
		return
	}
	file_shop_v1_address_proto_init()
	file_shop_v1_customer_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_shop_v1_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_shop_v1_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order_LineItem); i {
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
		file_shop_v1_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order_Payment); i {
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
			RawDescriptor: file_shop_v1_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shop_v1_order_proto_goTypes,
		DependencyIndexes: file_shop_v1_order_proto_depIdxs,
		MessageInfos:      file_shop_v1_order_proto_msgTypes,
	}.Build()
	File_shop_v1_order_proto = out.File
	file_shop_v1_order_proto_rawDesc = nil
	file_shop_v1_order_proto_goTypes = nil
	file_shop_v1_order_proto_depIdxs = nil
}
