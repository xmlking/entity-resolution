/* eslint-disable */
// @generated by protobuf-ts 2.4.0 with parameter generate_dependencies,ts_nocheck,eslint_disable,// @generated from protobuf file "google/type/latlng.proto" (package "google.type", syntax proto3),// tslint:disable
// @ts-nocheck
//
// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
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
/**
 * An object that represents a latitude/longitude pair. This is expressed as a
 * pair of doubles to represent degrees latitude and degrees longitude. Unless
 * specified otherwise, this must conform to the
 * <a href="http://www.unoosa.org/pdf/icg/2012/template/WGS_84.pdf">WGS84
 * standard</a>. Values must be within normalized ranges.
 *
 * @generated from protobuf message google.type.LatLng
 */
export interface LatLng {
    /**
     * The latitude in degrees. It must be in the range [-90.0, +90.0].
     *
     * @generated from protobuf field: double latitude = 1;
     */
    latitude: number;
    /**
     * The longitude in degrees. It must be in the range [-180.0, +180.0].
     *
     * @generated from protobuf field: double longitude = 2;
     */
    longitude: number;
}
// @generated message type with reflection information, may provide speed optimized methods
class LatLng$Type extends MessageType<LatLng> {
    constructor() {
        super("google.type.LatLng", [
            { no: 1, name: "latitude", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 2, name: "longitude", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ }
        ]);
    }
    create(value?: PartialMessage<LatLng>): LatLng {
        const message = { latitude: 0, longitude: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<LatLng>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: LatLng): LatLng {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* double latitude */ 1:
                    message.latitude = reader.double();
                    break;
                case /* double longitude */ 2:
                    message.longitude = reader.double();
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
    internalBinaryWrite(message: LatLng, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* double latitude = 1; */
        if (message.latitude !== 0)
            writer.tag(1, WireType.Bit64).double(message.latitude);
        /* double longitude = 2; */
        if (message.longitude !== 0)
            writer.tag(2, WireType.Bit64).double(message.longitude);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message google.type.LatLng
 */
export const LatLng = new LatLng$Type();
