/* eslint-disable */
// @generated by protobuf-ts 2.4.0 with parameter generate_dependencies,ts_nocheck,eslint_disable,// @generated from protobuf file "er/schema/entity/v1/entity.proto" (package "er.schema.entity.v1", syntax proto3),// tslint:disable
// @ts-nocheck
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
import { DateTime } from "../../../../google/type/datetime";
import { Date } from "../../../../google/type/date";
import { LatLng } from "../../../../google/type/latlng";
/**
 * Address US PostalAddress
 *
 * @generated from protobuf message er.schema.entity.v1.Address
 */
export interface Address {
    /**
     * @generated from protobuf field: string suite = 1;
     */
    suite: string;
    /**
     * @generated from protobuf field: string street = 2;
     */
    street: string;
    /**
     * @generated from protobuf field: string city = 3;
     */
    city: string;
    /**
     * @generated from protobuf field: string state = 4;
     */
    state: string;
    /**
     * @generated from protobuf field: string code = 5;
     */
    code: string;
    /**
     * @generated from protobuf field: string country = 6;
     */
    country: string;
    /**
     * @generated from protobuf field: optional google.type.LatLng location = 7;
     */
    location?: LatLng;
}
// 
// message Location {
// double lat = 1 [(validate.rules).double = { gte: -90,  lte: 90 }];
// double lng = 2 [(validate.rules).double = { gte: -180, lte: 180 }];
// }

/**
 * @generated from protobuf message er.schema.entity.v1.Name
 */
export interface Name {
    /**
     * @generated from protobuf field: string first = 1;
     */
    first: string;
    /**
     * @generated from protobuf field: string last = 2;
     */
    last: string;
    /**
     * @generated from protobuf field: optional string middle = 3;
     */
    middle?: string;
    /**
     * @generated from protobuf field: optional string title = 4;
     */
    title?: string;
}
/**
 * @generated from protobuf message er.schema.entity.v1.Member
 */
export interface Member {
    /**
     * @generated from protobuf field: string external_id = 1;
     */
    externalId: string;
    /**
     * @generated from protobuf field: optional string national_id = 12;
     */
    nationalId?: string;
    /**
     * @generated from protobuf field: repeated er.schema.entity.v1.Name names = 2;
     */
    names: Name[];
    /**
     * @generated from protobuf field: google.type.Date dob = 3;
     */
    dob?: Date;
    /**
     * @generated from protobuf field: er.schema.entity.v1.Gender gender = 4;
     */
    gender: Gender;
    /**
     * @generated from protobuf field: map<string, string> emails = 5;
     */
    emails: {
        [key: string]: string;
    };
    /**
     * @generated from protobuf field: map<string, string> phones = 6;
     */
    phones: {
        [key: string]: string;
    };
    /**
     * map<string, google.type.PhoneNumber> phones = 6 ;
     *
     * @generated from protobuf field: repeated er.schema.entity.v1.Address addresses = 7;
     */
    addresses: Address[];
    /**
     *  repeated google.type.PostalAddress addresses = 7;
     *
     * @generated from protobuf field: optional google.type.DateTime created_on = 20;
     */
    createdOn?: DateTime;
}
/**
 * @generated from protobuf enum er.schema.entity.v1.Status
 */
export enum Status {
    /**
     * @generated from protobuf enum value: STATUS_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: STATUS_FULL_MATCH = 1;
     */
    FULL_MATCH = 1,
    /**
     * @generated from protobuf enum value: STATUS_PARTIAL_MATCH = 2;
     */
    PARTIAL_MATCH = 2,
    /**
     * @generated from protobuf enum value: STATUS_NO_HIT = 3;
     */
    NO_HIT = 3,
    /**
     * @generated from protobuf enum value: STATUS_NEW = 4;
     */
    NEW = 4
}
/**
 * @generated from protobuf enum er.schema.entity.v1.Gender
 */
export enum Gender {
    /**
     * @generated from protobuf enum value: GENDER_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: GENDER_MALE = 1;
     */
    MALE = 1,
    /**
     * @generated from protobuf enum value: GENDER_FEMALE = 2;
     */
    FEMALE = 2,
    /**
     * @generated from protobuf enum value: GENDER_OTHER = 3;
     */
    OTHER = 3
}
// @generated message type with reflection information, may provide speed optimized methods
class Address$Type extends MessageType<Address> {
    constructor() {
        super("er.schema.entity.v1.Address", [
            { no: 1, name: "suite", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "street", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "city", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "state", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 5, name: "code", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 6, name: "country", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 7, name: "location", kind: "message", T: () => LatLng }
        ]);
    }
    create(value?: PartialMessage<Address>): Address {
        const message = { suite: "", street: "", city: "", state: "", code: "", country: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Address>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Address): Address {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string suite */ 1:
                    message.suite = reader.string();
                    break;
                case /* string street */ 2:
                    message.street = reader.string();
                    break;
                case /* string city */ 3:
                    message.city = reader.string();
                    break;
                case /* string state */ 4:
                    message.state = reader.string();
                    break;
                case /* string code */ 5:
                    message.code = reader.string();
                    break;
                case /* string country */ 6:
                    message.country = reader.string();
                    break;
                case /* optional google.type.LatLng location */ 7:
                    message.location = LatLng.internalBinaryRead(reader, reader.uint32(), options, message.location);
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
    internalBinaryWrite(message: Address, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string suite = 1; */
        if (message.suite !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.suite);
        /* string street = 2; */
        if (message.street !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.street);
        /* string city = 3; */
        if (message.city !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.city);
        /* string state = 4; */
        if (message.state !== "")
            writer.tag(4, WireType.LengthDelimited).string(message.state);
        /* string code = 5; */
        if (message.code !== "")
            writer.tag(5, WireType.LengthDelimited).string(message.code);
        /* string country = 6; */
        if (message.country !== "")
            writer.tag(6, WireType.LengthDelimited).string(message.country);
        /* optional google.type.LatLng location = 7; */
        if (message.location)
            LatLng.internalBinaryWrite(message.location, writer.tag(7, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message er.schema.entity.v1.Address
 */
export const Address = new Address$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Name$Type extends MessageType<Name> {
    constructor() {
        super("er.schema.entity.v1.Name", [
            { no: 1, name: "first", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { maxBytes: "256", pattern: "^[^[0-9]A-Za-z]+( [^[0-9]A-Za-z]+)*$" } } } },
            { no: 2, name: "last", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "middle", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "title", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/, options: { "tagger.tags": "avro:\"title\" graphql:\"withNewTags,optional\"" } }
        ]);
    }
    create(value?: PartialMessage<Name>): Name {
        const message = { first: "", last: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Name>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Name): Name {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string first */ 1:
                    message.first = reader.string();
                    break;
                case /* string last */ 2:
                    message.last = reader.string();
                    break;
                case /* optional string middle */ 3:
                    message.middle = reader.string();
                    break;
                case /* optional string title */ 4:
                    message.title = reader.string();
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
    internalBinaryWrite(message: Name, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string first = 1; */
        if (message.first !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.first);
        /* string last = 2; */
        if (message.last !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.last);
        /* optional string middle = 3; */
        if (message.middle !== undefined)
            writer.tag(3, WireType.LengthDelimited).string(message.middle);
        /* optional string title = 4; */
        if (message.title !== undefined)
            writer.tag(4, WireType.LengthDelimited).string(message.title);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message er.schema.entity.v1.Name
 */
export const Name = new Name$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Member$Type extends MessageType<Member> {
    constructor() {
        super("er.schema.entity.v1.Member", [
            { no: 1, name: "external_id", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "validate.rules": { string: { uuid: true } } } },
            { no: 12, name: "national_id", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "names", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Name, options: { "validate.rules": { repeated: { minItems: "1" } } } },
            { no: 3, name: "dob", kind: "message", T: () => Date },
            { no: 4, name: "gender", kind: "enum", T: () => ["er.schema.entity.v1.Gender", Gender, "GENDER_"] },
            { no: 5, name: "emails", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "scalar", T: 9 /*ScalarType.STRING*/ }, options: { "validate.rules": { map: { maxPairs: "5", keys: {}, values: { string: { email: true } } } } } },
            { no: 6, name: "phones", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "scalar", T: 9 /*ScalarType.STRING*/ }, options: { "validate.rules": { map: { maxPairs: "5", keys: { string: { minLen: "5" } }, values: { string: { minLen: "5", maxLen: "15" } } } } } },
            { no: 7, name: "addresses", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Address },
            { no: 20, name: "created_on", kind: "message", T: () => DateTime }
        ]);
    }
    create(value?: PartialMessage<Member>): Member {
        const message = { externalId: "", names: [], gender: 0, emails: {}, phones: {}, addresses: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Member>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Member): Member {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string external_id */ 1:
                    message.externalId = reader.string();
                    break;
                case /* optional string national_id */ 12:
                    message.nationalId = reader.string();
                    break;
                case /* repeated er.schema.entity.v1.Name names */ 2:
                    message.names.push(Name.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* google.type.Date dob */ 3:
                    message.dob = Date.internalBinaryRead(reader, reader.uint32(), options, message.dob);
                    break;
                case /* er.schema.entity.v1.Gender gender */ 4:
                    message.gender = reader.int32();
                    break;
                case /* map<string, string> emails */ 5:
                    this.binaryReadMap5(message.emails, reader, options);
                    break;
                case /* map<string, string> phones */ 6:
                    this.binaryReadMap6(message.phones, reader, options);
                    break;
                case /* repeated er.schema.entity.v1.Address addresses */ 7:
                    message.addresses.push(Address.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* optional google.type.DateTime created_on */ 20:
                    message.createdOn = DateTime.internalBinaryRead(reader, reader.uint32(), options, message.createdOn);
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
    private binaryReadMap5(map: Member["emails"], reader: IBinaryReader, options: BinaryReadOptions): void {
        let len = reader.uint32(), end = reader.pos + len, key: keyof Member["emails"] | undefined, val: Member["emails"][any] | undefined;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case 1:
                    key = reader.string();
                    break;
                case 2:
                    val = reader.string();
                    break;
                default: throw new globalThis.Error("unknown map entry field for field er.schema.entity.v1.Member.emails");
            }
        }
        map[key ?? ""] = val ?? "";
    }
    private binaryReadMap6(map: Member["phones"], reader: IBinaryReader, options: BinaryReadOptions): void {
        let len = reader.uint32(), end = reader.pos + len, key: keyof Member["phones"] | undefined, val: Member["phones"][any] | undefined;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case 1:
                    key = reader.string();
                    break;
                case 2:
                    val = reader.string();
                    break;
                default: throw new globalThis.Error("unknown map entry field for field er.schema.entity.v1.Member.phones");
            }
        }
        map[key ?? ""] = val ?? "";
    }
    internalBinaryWrite(message: Member, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string external_id = 1; */
        if (message.externalId !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.externalId);
        /* optional string national_id = 12; */
        if (message.nationalId !== undefined)
            writer.tag(12, WireType.LengthDelimited).string(message.nationalId);
        /* repeated er.schema.entity.v1.Name names = 2; */
        for (let i = 0; i < message.names.length; i++)
            Name.internalBinaryWrite(message.names[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* google.type.Date dob = 3; */
        if (message.dob)
            Date.internalBinaryWrite(message.dob, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* er.schema.entity.v1.Gender gender = 4; */
        if (message.gender !== 0)
            writer.tag(4, WireType.Varint).int32(message.gender);
        /* map<string, string> emails = 5; */
        for (let k of Object.keys(message.emails))
            writer.tag(5, WireType.LengthDelimited).fork().tag(1, WireType.LengthDelimited).string(k).tag(2, WireType.LengthDelimited).string(message.emails[k]).join();
        /* map<string, string> phones = 6; */
        for (let k of Object.keys(message.phones))
            writer.tag(6, WireType.LengthDelimited).fork().tag(1, WireType.LengthDelimited).string(k).tag(2, WireType.LengthDelimited).string(message.phones[k]).join();
        /* repeated er.schema.entity.v1.Address addresses = 7; */
        for (let i = 0; i < message.addresses.length; i++)
            Address.internalBinaryWrite(message.addresses[i], writer.tag(7, WireType.LengthDelimited).fork(), options).join();
        /* optional google.type.DateTime created_on = 20; */
        if (message.createdOn)
            DateTime.internalBinaryWrite(message.createdOn, writer.tag(20, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message er.schema.entity.v1.Member
 */
export const Member = new Member$Type();
