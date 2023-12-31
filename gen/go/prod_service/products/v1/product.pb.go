// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: prod_service/products/v1/product.proto

package pb_product

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

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity           string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	PrintedWeight      string `protobuf:"bytes,2,opt,name=printed_weight,json=printedWeight,proto3" json:"printed_weight,omitempty"`
	GrossWeight        string `protobuf:"bytes,3,opt,name=gross_weight,json=grossWeight,proto3" json:"gross_weight,omitempty"`
	Batch              string `protobuf:"bytes,4,opt,name=batch,proto3" json:"batch,omitempty"`
	BoxNumber          string `protobuf:"bytes,5,opt,name=box_number,json=boxNumber,proto3" json:"box_number,omitempty"`
	PalletNumber       string `protobuf:"bytes,6,opt,name=pallet_number,json=palletNumber,proto3" json:"pallet_number,omitempty"`
	ContainerWeight    string `protobuf:"bytes,7,opt,name=container_weight,json=containerWeight,proto3" json:"container_weight,omitempty"`
	DeviceName         string `protobuf:"bytes,8,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	DeviceNumber       string `protobuf:"bytes,9,opt,name=device_number,json=deviceNumber,proto3" json:"device_number,omitempty"`
	Time               string `protobuf:"bytes,10,opt,name=time,proto3" json:"time,omitempty"`
	ArticleNumber      string `protobuf:"bytes,11,opt,name=article_number,json=articleNumber,proto3" json:"article_number,omitempty"`
	ArticleName        string `protobuf:"bytes,12,opt,name=article_name,json=articleName,proto3" json:"article_name,omitempty"`
	AmountUnitsProduct int64  `protobuf:"varint,13,opt,name=amount_units_product,json=amountUnitsProduct,proto3" json:"amount_units_product,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prod_service_products_v1_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_prod_service_products_v1_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_prod_service_products_v1_product_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *Product) GetPrintedWeight() string {
	if x != nil {
		return x.PrintedWeight
	}
	return ""
}

func (x *Product) GetGrossWeight() string {
	if x != nil {
		return x.GrossWeight
	}
	return ""
}

func (x *Product) GetBatch() string {
	if x != nil {
		return x.Batch
	}
	return ""
}

func (x *Product) GetBoxNumber() string {
	if x != nil {
		return x.BoxNumber
	}
	return ""
}

func (x *Product) GetPalletNumber() string {
	if x != nil {
		return x.PalletNumber
	}
	return ""
}

func (x *Product) GetContainerWeight() string {
	if x != nil {
		return x.ContainerWeight
	}
	return ""
}

func (x *Product) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *Product) GetDeviceNumber() string {
	if x != nil {
		return x.DeviceNumber
	}
	return ""
}

func (x *Product) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *Product) GetArticleNumber() string {
	if x != nil {
		return x.ArticleNumber
	}
	return ""
}

func (x *Product) GetArticleName() string {
	if x != nil {
		return x.ArticleName
	}
	return ""
}

func (x *Product) GetAmountUnitsProduct() int64 {
	if x != nil {
		return x.AmountUnitsProduct
	}
	return 0
}

var File_prod_service_products_v1_product_proto protoreflect.FileDescriptor

var file_prod_service_products_v1_product_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x22, 0xca, 0x03, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72,
	0x69, 0x6e, 0x74, 0x65, 0x64, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x64, 0x57, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x72, 0x6f, 0x73, 0x73, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x67, 0x72, 0x6f, 0x73, 0x73, 0x57, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f,
	0x78, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x62, 0x6f, 0x78, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x70, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x29,
	0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a,
	0x14, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x5f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42,
	0x23, 0x5a, 0x21, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x5f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prod_service_products_v1_product_proto_rawDescOnce sync.Once
	file_prod_service_products_v1_product_proto_rawDescData = file_prod_service_products_v1_product_proto_rawDesc
)

func file_prod_service_products_v1_product_proto_rawDescGZIP() []byte {
	file_prod_service_products_v1_product_proto_rawDescOnce.Do(func() {
		file_prod_service_products_v1_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_prod_service_products_v1_product_proto_rawDescData)
	})
	return file_prod_service_products_v1_product_proto_rawDescData
}

var file_prod_service_products_v1_product_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_prod_service_products_v1_product_proto_goTypes = []interface{}{
	(*Product)(nil), // 0: prod_service.products.v1.Product
}
var file_prod_service_products_v1_product_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_prod_service_products_v1_product_proto_init() }
func file_prod_service_products_v1_product_proto_init() {
	if File_prod_service_products_v1_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prod_service_products_v1_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
			RawDescriptor: file_prod_service_products_v1_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prod_service_products_v1_product_proto_goTypes,
		DependencyIndexes: file_prod_service_products_v1_product_proto_depIdxs,
		MessageInfos:      file_prod_service_products_v1_product_proto_msgTypes,
	}.Build()
	File_prod_service_products_v1_product_proto = out.File
	file_prod_service_products_v1_product_proto_rawDesc = nil
	file_prod_service_products_v1_product_proto_goTypes = nil
	file_prod_service_products_v1_product_proto_depIdxs = nil
}
