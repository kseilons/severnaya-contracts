// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: prod_service/products/v1/service.proto

package pb_product

import (
	v1 "gen/go/common/filter/v1"
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

type AllProductsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination    *v1.Pagination        `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Batch         *v1.StringFieldFilter `protobuf:"bytes,2,opt,name=batch,proto3" json:"batch,omitempty"`
	PalletNumber  *v1.StringFieldFilter `protobuf:"bytes,3,opt,name=pallet_number,json=palletNumber,proto3" json:"pallet_number,omitempty"`
	DeviceName    *v1.StringFieldFilter `protobuf:"bytes,4,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	DeviceNumber  *v1.IntFieldFilter    `protobuf:"bytes,5,opt,name=device_number,json=deviceNumber,proto3" json:"device_number,omitempty"`
	Time          *v1.IntFieldFilter    `protobuf:"bytes,6,opt,name=time,proto3" json:"time,omitempty"`
	ArticleNumber *v1.IntFieldFilter    `protobuf:"bytes,7,opt,name=article_number,json=articleNumber,proto3" json:"article_number,omitempty"`
	ArticleName   *v1.IntFieldFilter    `protobuf:"bytes,8,opt,name=article_name,json=articleName,proto3" json:"article_name,omitempty"`
	Sort          *v1.Sort              `protobuf:"bytes,9,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *AllProductsRequest) Reset() {
	*x = AllProductsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prod_service_products_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllProductsRequest) ProtoMessage() {}

func (x *AllProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prod_service_products_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllProductsRequest.ProtoReflect.Descriptor instead.
func (*AllProductsRequest) Descriptor() ([]byte, []int) {
	return file_prod_service_products_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *AllProductsRequest) GetPagination() *v1.Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *AllProductsRequest) GetBatch() *v1.StringFieldFilter {
	if x != nil {
		return x.Batch
	}
	return nil
}

func (x *AllProductsRequest) GetPalletNumber() *v1.StringFieldFilter {
	if x != nil {
		return x.PalletNumber
	}
	return nil
}

func (x *AllProductsRequest) GetDeviceName() *v1.StringFieldFilter {
	if x != nil {
		return x.DeviceName
	}
	return nil
}

func (x *AllProductsRequest) GetDeviceNumber() *v1.IntFieldFilter {
	if x != nil {
		return x.DeviceNumber
	}
	return nil
}

func (x *AllProductsRequest) GetTime() *v1.IntFieldFilter {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *AllProductsRequest) GetArticleNumber() *v1.IntFieldFilter {
	if x != nil {
		return x.ArticleNumber
	}
	return nil
}

func (x *AllProductsRequest) GetArticleName() *v1.IntFieldFilter {
	if x != nil {
		return x.ArticleName
	}
	return nil
}

func (x *AllProductsRequest) GetSort() *v1.Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

type AllProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *AllProductsResponse) Reset() {
	*x = AllProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prod_service_products_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllProductsResponse) ProtoMessage() {}

func (x *AllProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_prod_service_products_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllProductsResponse.ProtoReflect.Descriptor instead.
func (*AllProductsResponse) Descriptor() ([]byte, []int) {
	return file_prod_service_products_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *AllProductsResponse) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type ProductByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProductByIDRequest) Reset() {
	*x = ProductByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prod_service_products_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductByIDRequest) ProtoMessage() {}

func (x *ProductByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prod_service_products_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductByIDRequest.ProtoReflect.Descriptor instead.
func (*ProductByIDRequest) Descriptor() ([]byte, []int) {
	return file_prod_service_products_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *ProductByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ProductByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *ProductByIDResponse) Reset() {
	*x = ProductByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prod_service_products_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductByIDResponse) ProtoMessage() {}

func (x *ProductByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_prod_service_products_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductByIDResponse.ProtoReflect.Descriptor instead.
func (*ProductByIDResponse) Descriptor() ([]byte, []int) {
	return file_prod_service_products_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *ProductByIDResponse) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

var File_prod_service_products_v1_service_proto protoreflect.FileDescriptor

var file_prod_service_products_v1_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x1d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26,
	0x70, 0x72, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x04, 0x0a, 0x12, 0x41, 0x6c, 0x6c, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x05, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x12, 0x48, 0x0a, 0x0d, 0x70, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x0c, 0x70, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x44, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x0d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x34, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0d, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x0c,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x2a, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x54, 0x0a,
	0x13, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x52, 0x0a, 0x13, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3b, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x32, 0xec, 0x01,
	0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x6c, 0x0a, 0x0b, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12,
	0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6c,
	0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x2c, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prod_service_products_v1_service_proto_rawDescOnce sync.Once
	file_prod_service_products_v1_service_proto_rawDescData = file_prod_service_products_v1_service_proto_rawDesc
)

func file_prod_service_products_v1_service_proto_rawDescGZIP() []byte {
	file_prod_service_products_v1_service_proto_rawDescOnce.Do(func() {
		file_prod_service_products_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_prod_service_products_v1_service_proto_rawDescData)
	})
	return file_prod_service_products_v1_service_proto_rawDescData
}

var file_prod_service_products_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_prod_service_products_v1_service_proto_goTypes = []interface{}{
	(*AllProductsRequest)(nil),   // 0: prod_service.products.v1.AllProductsRequest
	(*AllProductsResponse)(nil),  // 1: prod_service.products.v1.AllProductsResponse
	(*ProductByIDRequest)(nil),   // 2: prod_service.products.v1.ProductByIDRequest
	(*ProductByIDResponse)(nil),  // 3: prod_service.products.v1.ProductByIDResponse
	(*v1.Pagination)(nil),        // 4: common.filter.v1.Pagination
	(*v1.StringFieldFilter)(nil), // 5: common.filter.v1.StringFieldFilter
	(*v1.IntFieldFilter)(nil),    // 6: common.filter.v1.IntFieldFilter
	(*v1.Sort)(nil),              // 7: common.filter.v1.Sort
	(*Product)(nil),              // 8: prod_service.products.v1.Product
}
var file_prod_service_products_v1_service_proto_depIdxs = []int32{
	4,  // 0: prod_service.products.v1.AllProductsRequest.pagination:type_name -> common.filter.v1.Pagination
	5,  // 1: prod_service.products.v1.AllProductsRequest.batch:type_name -> common.filter.v1.StringFieldFilter
	5,  // 2: prod_service.products.v1.AllProductsRequest.pallet_number:type_name -> common.filter.v1.StringFieldFilter
	5,  // 3: prod_service.products.v1.AllProductsRequest.device_name:type_name -> common.filter.v1.StringFieldFilter
	6,  // 4: prod_service.products.v1.AllProductsRequest.device_number:type_name -> common.filter.v1.IntFieldFilter
	6,  // 5: prod_service.products.v1.AllProductsRequest.time:type_name -> common.filter.v1.IntFieldFilter
	6,  // 6: prod_service.products.v1.AllProductsRequest.article_number:type_name -> common.filter.v1.IntFieldFilter
	6,  // 7: prod_service.products.v1.AllProductsRequest.article_name:type_name -> common.filter.v1.IntFieldFilter
	7,  // 8: prod_service.products.v1.AllProductsRequest.sort:type_name -> common.filter.v1.Sort
	8,  // 9: prod_service.products.v1.AllProductsResponse.products:type_name -> prod_service.products.v1.Product
	8,  // 10: prod_service.products.v1.ProductByIDResponse.product:type_name -> prod_service.products.v1.Product
	0,  // 11: prod_service.products.v1.ProductService.AllProducts:input_type -> prod_service.products.v1.AllProductsRequest
	2,  // 12: prod_service.products.v1.ProductService.ProductByID:input_type -> prod_service.products.v1.ProductByIDRequest
	1,  // 13: prod_service.products.v1.ProductService.AllProducts:output_type -> prod_service.products.v1.AllProductsResponse
	3,  // 14: prod_service.products.v1.ProductService.ProductByID:output_type -> prod_service.products.v1.ProductByIDResponse
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_prod_service_products_v1_service_proto_init() }
func file_prod_service_products_v1_service_proto_init() {
	if File_prod_service_products_v1_service_proto != nil {
		return
	}
	file_prod_service_products_v1_product_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_prod_service_products_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllProductsRequest); i {
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
		file_prod_service_products_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllProductsResponse); i {
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
		file_prod_service_products_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductByIDRequest); i {
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
		file_prod_service_products_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductByIDResponse); i {
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
			RawDescriptor: file_prod_service_products_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_prod_service_products_v1_service_proto_goTypes,
		DependencyIndexes: file_prod_service_products_v1_service_proto_depIdxs,
		MessageInfos:      file_prod_service_products_v1_service_proto_msgTypes,
	}.Build()
	File_prod_service_products_v1_service_proto = out.File
	file_prod_service_products_v1_service_proto_rawDesc = nil
	file_prod_service_products_v1_service_proto_goTypes = nil
	file_prod_service_products_v1_service_proto_depIdxs = nil
}
