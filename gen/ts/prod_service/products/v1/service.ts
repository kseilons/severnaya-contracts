// @generated by protobuf-ts 2.2.2 with parameter long_type_string
// @generated from protobuf file "prod_service/products/v1/service.proto" (package "prod_service.products.v1", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { Product } from "./product";
import { Sort } from "../../../common/filter/v1/sort";
import { IntFieldFilter } from "../../../common/filter/v1/filter";
import { StringFieldFilter } from "../../../common/filter/v1/filter";
import { Pagination } from "../../../common/filter/v1/filter";
/**
 * @generated from protobuf message prod_service.products.v1.AllProductsRequest
 */
export interface AllProductsRequest {
    /**
     * @generated from protobuf field: common.filter.v1.Pagination pagination = 1;
     */
    pagination?: Pagination;
    /**
     * @generated from protobuf field: common.filter.v1.StringFieldFilter Batch = 2 [json_name = "Batch"];
     */
    batch?: StringFieldFilter;
    /**
     * @generated from protobuf field: common.filter.v1.StringFieldFilter PalletNumber = 3 [json_name = "PalletNumber"];
     */
    palletNumber?: StringFieldFilter;
    /**
     * @generated from protobuf field: common.filter.v1.StringFieldFilter DeviceName = 4 [json_name = "DeviceName"];
     */
    deviceName?: StringFieldFilter;
    /**
     * @generated from protobuf field: common.filter.v1.IntFieldFilter DeviceNumber = 5 [json_name = "DeviceNumber"];
     */
    deviceNumber?: IntFieldFilter;
    /**
     * @generated from protobuf field: common.filter.v1.IntFieldFilter Time = 6 [json_name = "Time"];
     */
    time?: IntFieldFilter;
    /**
     * @generated from protobuf field: common.filter.v1.IntFieldFilter ArticleNumber = 7 [json_name = "ArticleNumber"];
     */
    articleNumber?: IntFieldFilter;
    /**
     * @generated from protobuf field: common.filter.v1.IntFieldFilter ArticleName = 8 [json_name = "ArticleName"];
     */
    articleName?: IntFieldFilter;
    /**
     * @generated from protobuf field: common.filter.v1.Sort sort = 9;
     */
    sort?: Sort;
}
/**
 * @generated from protobuf message prod_service.products.v1.AllProductsResponse
 */
export interface AllProductsResponse {
    /**
     * @generated from protobuf field: repeated prod_service.products.v1.Product products = 1;
     */
    products: Product[];
}
/**
 * @generated from protobuf message prod_service.products.v1.ProductByIDRequest
 */
export interface ProductByIDRequest {
    /**
     * ID
     *
     * @generated from protobuf field: string id = 1;
     */
    id: string;
}
/**
 * @generated from protobuf message prod_service.products.v1.ProductByIDResponse
 */
export interface ProductByIDResponse {
    /**
     * @generated from protobuf field: prod_service.products.v1.Product product = 1;
     */
    product?: Product;
}
// @generated message type with reflection information, may provide speed optimized methods
class AllProductsRequest$Type extends MessageType<AllProductsRequest> {
    constructor() {
        super("prod_service.products.v1.AllProductsRequest", [
            { no: 1, name: "pagination", kind: "message", T: () => Pagination },
            { no: 2, name: "Batch", kind: "message", jsonName: "Batch", T: () => StringFieldFilter },
            { no: 3, name: "PalletNumber", kind: "message", jsonName: "PalletNumber", T: () => StringFieldFilter },
            { no: 4, name: "DeviceName", kind: "message", jsonName: "DeviceName", T: () => StringFieldFilter },
            { no: 5, name: "DeviceNumber", kind: "message", jsonName: "DeviceNumber", T: () => IntFieldFilter },
            { no: 6, name: "Time", kind: "message", jsonName: "Time", T: () => IntFieldFilter },
            { no: 7, name: "ArticleNumber", kind: "message", jsonName: "ArticleNumber", T: () => IntFieldFilter },
            { no: 8, name: "ArticleName", kind: "message", jsonName: "ArticleName", T: () => IntFieldFilter },
            { no: 9, name: "sort", kind: "message", T: () => Sort }
        ]);
    }
    create(value?: PartialMessage<AllProductsRequest>): AllProductsRequest {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<AllProductsRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: AllProductsRequest): AllProductsRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* common.filter.v1.Pagination pagination */ 1:
                    message.pagination = Pagination.internalBinaryRead(reader, reader.uint32(), options, message.pagination);
                    break;
                case /* common.filter.v1.StringFieldFilter Batch = 2 [json_name = "Batch"];*/ 2:
                    message.batch = StringFieldFilter.internalBinaryRead(reader, reader.uint32(), options, message.batch);
                    break;
                case /* common.filter.v1.StringFieldFilter PalletNumber = 3 [json_name = "PalletNumber"];*/ 3:
                    message.palletNumber = StringFieldFilter.internalBinaryRead(reader, reader.uint32(), options, message.palletNumber);
                    break;
                case /* common.filter.v1.StringFieldFilter DeviceName = 4 [json_name = "DeviceName"];*/ 4:
                    message.deviceName = StringFieldFilter.internalBinaryRead(reader, reader.uint32(), options, message.deviceName);
                    break;
                case /* common.filter.v1.IntFieldFilter DeviceNumber = 5 [json_name = "DeviceNumber"];*/ 5:
                    message.deviceNumber = IntFieldFilter.internalBinaryRead(reader, reader.uint32(), options, message.deviceNumber);
                    break;
                case /* common.filter.v1.IntFieldFilter Time = 6 [json_name = "Time"];*/ 6:
                    message.time = IntFieldFilter.internalBinaryRead(reader, reader.uint32(), options, message.time);
                    break;
                case /* common.filter.v1.IntFieldFilter ArticleNumber = 7 [json_name = "ArticleNumber"];*/ 7:
                    message.articleNumber = IntFieldFilter.internalBinaryRead(reader, reader.uint32(), options, message.articleNumber);
                    break;
                case /* common.filter.v1.IntFieldFilter ArticleName = 8 [json_name = "ArticleName"];*/ 8:
                    message.articleName = IntFieldFilter.internalBinaryRead(reader, reader.uint32(), options, message.articleName);
                    break;
                case /* common.filter.v1.Sort sort */ 9:
                    message.sort = Sort.internalBinaryRead(reader, reader.uint32(), options, message.sort);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: AllProductsRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* common.filter.v1.Pagination pagination = 1; */
        if (message.pagination)
            Pagination.internalBinaryWrite(message.pagination, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* common.filter.v1.StringFieldFilter Batch = 2 [json_name = "Batch"]; */
        if (message.batch)
            StringFieldFilter.internalBinaryWrite(message.batch, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* common.filter.v1.StringFieldFilter PalletNumber = 3 [json_name = "PalletNumber"]; */
        if (message.palletNumber)
            StringFieldFilter.internalBinaryWrite(message.palletNumber, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* common.filter.v1.StringFieldFilter DeviceName = 4 [json_name = "DeviceName"]; */
        if (message.deviceName)
            StringFieldFilter.internalBinaryWrite(message.deviceName, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* common.filter.v1.IntFieldFilter DeviceNumber = 5 [json_name = "DeviceNumber"]; */
        if (message.deviceNumber)
            IntFieldFilter.internalBinaryWrite(message.deviceNumber, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        /* common.filter.v1.IntFieldFilter Time = 6 [json_name = "Time"]; */
        if (message.time)
            IntFieldFilter.internalBinaryWrite(message.time, writer.tag(6, WireType.LengthDelimited).fork(), options).join();
        /* common.filter.v1.IntFieldFilter ArticleNumber = 7 [json_name = "ArticleNumber"]; */
        if (message.articleNumber)
            IntFieldFilter.internalBinaryWrite(message.articleNumber, writer.tag(7, WireType.LengthDelimited).fork(), options).join();
        /* common.filter.v1.IntFieldFilter ArticleName = 8 [json_name = "ArticleName"]; */
        if (message.articleName)
            IntFieldFilter.internalBinaryWrite(message.articleName, writer.tag(8, WireType.LengthDelimited).fork(), options).join();
        /* common.filter.v1.Sort sort = 9; */
        if (message.sort)
            Sort.internalBinaryWrite(message.sort, writer.tag(9, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message prod_service.products.v1.AllProductsRequest
 */
export const AllProductsRequest = new AllProductsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class AllProductsResponse$Type extends MessageType<AllProductsResponse> {
    constructor() {
        super("prod_service.products.v1.AllProductsResponse", [
            { no: 1, name: "products", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Product }
        ]);
    }
    create(value?: PartialMessage<AllProductsResponse>): AllProductsResponse {
        const message = { products: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<AllProductsResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: AllProductsResponse): AllProductsResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated prod_service.products.v1.Product products */ 1:
                    message.products.push(Product.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: AllProductsResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated prod_service.products.v1.Product products = 1; */
        for (let i = 0; i < message.products.length; i++)
            Product.internalBinaryWrite(message.products[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message prod_service.products.v1.AllProductsResponse
 */
export const AllProductsResponse = new AllProductsResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ProductByIDRequest$Type extends MessageType<ProductByIDRequest> {
    constructor() {
        super("prod_service.products.v1.ProductByIDRequest", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<ProductByIDRequest>): ProductByIDRequest {
        const message = { id: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<ProductByIDRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ProductByIDRequest): ProductByIDRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string id */ 1:
                    message.id = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: ProductByIDRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string id = 1; */
        if (message.id !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.id);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message prod_service.products.v1.ProductByIDRequest
 */
export const ProductByIDRequest = new ProductByIDRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ProductByIDResponse$Type extends MessageType<ProductByIDResponse> {
    constructor() {
        super("prod_service.products.v1.ProductByIDResponse", [
            { no: 1, name: "product", kind: "message", T: () => Product }
        ]);
    }
    create(value?: PartialMessage<ProductByIDResponse>): ProductByIDResponse {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<ProductByIDResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ProductByIDResponse): ProductByIDResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* prod_service.products.v1.Product product */ 1:
                    message.product = Product.internalBinaryRead(reader, reader.uint32(), options, message.product);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: ProductByIDResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* prod_service.products.v1.Product product = 1; */
        if (message.product)
            Product.internalBinaryWrite(message.product, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message prod_service.products.v1.ProductByIDResponse
 */
export const ProductByIDResponse = new ProductByIDResponse$Type();
/**
 * @generated ServiceType for protobuf service prod_service.products.v1.ProductService
 */
export const ProductService = new ServiceType("prod_service.products.v1.ProductService", [
    { name: "AllProducts", options: {}, I: AllProductsRequest, O: AllProductsResponse },
    { name: "ProductByID", options: {}, I: ProductByIDRequest, O: ProductByIDResponse }
]);
