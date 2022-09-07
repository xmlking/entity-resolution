// @generated by protoc-gen-es v0.1.1 with parameter "target=ts"
// @generated from file er/schema/entity/v1/entity.proto (package er.schema.entity.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import type {BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage} from "@bufbuild/protobuf";
import {Message, proto3} from "@bufbuild/protobuf";
import {LatLng} from "../../../../google/type/latlng_pb.js";
import {Date} from "../../../../google/type/date_pb.js";
import {DateTime} from "../../../../google/type/datetime_pb.js";

/**
 * @generated from enum er.schema.entity.v1.Status
 */
export enum Status {
  /**
   * @generated from enum value: STATUS_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: STATUS_FULL_MATCH = 1;
   */
  FULL_MATCH = 1,

  /**
   * @generated from enum value: STATUS_PARTIAL_MATCH = 2;
   */
  PARTIAL_MATCH = 2,

  /**
   * @generated from enum value: STATUS_NO_HIT = 3;
   */
  NO_HIT = 3,

  /**
   * @generated from enum value: STATUS_NEW = 4;
   */
  NEW = 4,
}
// Retrieve enum metadata with: proto3.getEnumType(Status)
proto3.util.setEnumType(Status, "er.schema.entity.v1.Status", [
  { no: 0, name: "STATUS_UNSPECIFIED" },
  { no: 1, name: "STATUS_FULL_MATCH" },
  { no: 2, name: "STATUS_PARTIAL_MATCH" },
  { no: 3, name: "STATUS_NO_HIT" },
  { no: 4, name: "STATUS_NEW" },
]);

/**
 * @generated from enum er.schema.entity.v1.Gender
 */
export enum Gender {
  /**
   * @generated from enum value: GENDER_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: GENDER_MALE = 1;
   */
  MALE = 1,

  /**
   * @generated from enum value: GENDER_FEMALE = 2;
   */
  FEMALE = 2,

  /**
   * @generated from enum value: GENDER_OTHER = 3;
   */
  OTHER = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(Gender)
proto3.util.setEnumType(Gender, "er.schema.entity.v1.Gender", [
  { no: 0, name: "GENDER_UNSPECIFIED" },
  { no: 1, name: "GENDER_MALE" },
  { no: 2, name: "GENDER_FEMALE" },
  { no: 3, name: "GENDER_OTHER" },
]);

/**
 * Address US PostalAddress
 *
 * @generated from message er.schema.entity.v1.Address
 */
export class Address extends Message<Address> {
  /**
   * @generated from field: string suite = 1;
   */
  suite = "";

  /**
   * @generated from field: string street = 2;
   */
  street = "";

  /**
   * @generated from field: string city = 3;
   */
  city = "";

  /**
   * @generated from field: string state = 4;
   */
  state = "";

  /**
   * @generated from field: string code = 5;
   */
  code = "";

  /**
   * @generated from field: string country = 6;
   */
  country = "";

  /**
   * @generated from field: optional google.type.LatLng location = 7;
   */
  location?: LatLng;

  constructor(data?: PartialMessage<Address>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "er.schema.entity.v1.Address";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "suite", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "street", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "city", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "state", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "code", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "country", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "location", kind: "message", T: LatLng, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Address {
    return new Address().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Address {
    return new Address().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Address {
    return new Address().fromJsonString(jsonString, options);
  }

  static equals(a: Address | PlainMessage<Address> | undefined, b: Address | PlainMessage<Address> | undefined): boolean {
    return proto3.util.equals(Address, a, b);
  }
}

/**
 * @generated from message er.schema.entity.v1.Name
 */
export class Name extends Message<Name> {
  /**
   * @generated from field: string first = 1;
   */
  first = "";

  /**
   * @generated from field: string last = 2;
   */
  last = "";

  /**
   * @generated from field: optional string middle = 3;
   */
  middle?: string;

  /**
   * @generated from field: optional string title = 4;
   */
  title?: string;

  constructor(data?: PartialMessage<Name>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "er.schema.entity.v1.Name";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "first", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "last", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "middle", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 4, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Name {
    return new Name().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Name {
    return new Name().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Name {
    return new Name().fromJsonString(jsonString, options);
  }

  static equals(a: Name | PlainMessage<Name> | undefined, b: Name | PlainMessage<Name> | undefined): boolean {
    return proto3.util.equals(Name, a, b);
  }
}

/**
 * @generated from message er.schema.entity.v1.Member
 */
export class Member extends Message<Member> {
  /**
   * @generated from field: string external_id = 1;
   */
  externalId = "";

  /**
   * @generated from field: optional string national_id = 12;
   */
  nationalId?: string;

  /**
   * @generated from field: repeated er.schema.entity.v1.Name names = 2;
   */
  names: Name[] = [];

  /**
   * @generated from field: google.type.Date dob = 3;
   */
  dob?: Date;

  /**
   * @generated from field: er.schema.entity.v1.Gender gender = 4;
   */
  gender = Gender.UNSPECIFIED;

  /**
   * @generated from field: map<string, string> emails = 5;
   */
  emails: { [key: string]: string } = {};

  /**
   * @generated from field: map<string, string> phones = 6;
   */
  phones: { [key: string]: string } = {};

  /**
   * map<string, google.type.PhoneNumber> phones = 6 ;
   *
   * @generated from field: repeated er.schema.entity.v1.Address addresses = 7;
   */
  addresses: Address[] = [];

  /**
   *  repeated google.type.PostalAddress addresses = 7;
   *
   * @generated from field: optional google.type.DateTime created_on = 20;
   */
  createdOn?: DateTime;

  constructor(data?: PartialMessage<Member>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "er.schema.entity.v1.Member";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "external_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 12, name: "national_id", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 2, name: "names", kind: "message", T: Name, repeated: true },
    { no: 3, name: "dob", kind: "message", T: Date },
    { no: 4, name: "gender", kind: "enum", T: proto3.getEnumType(Gender) },
    { no: 5, name: "emails", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 6, name: "phones", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 7, name: "addresses", kind: "message", T: Address, repeated: true },
    { no: 20, name: "created_on", kind: "message", T: DateTime, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Member {
    return new Member().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Member {
    return new Member().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Member {
    return new Member().fromJsonString(jsonString, options);
  }

  static equals(a: Member | PlainMessage<Member> | undefined, b: Member | PlainMessage<Member> | undefined): boolean {
    return proto3.util.equals(Member, a, b);
  }
}

