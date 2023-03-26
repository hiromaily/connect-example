// Copyright 2022 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// @generated by protoc-gen-es v1.2.0 with parameter "target=ts,import_extension=none"
// @generated from file eliza/v1/eliza.proto (package eliza.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * SayRequest describes the sentence said to the ELIZA program.
 *
 * @generated from message eliza.v1.SayRequest
 */
export class SayRequest extends Message<SayRequest> {
  /**
   * @generated from field: string sentence = 1;
   */
  sentence = "";

  constructor(data?: PartialMessage<SayRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "eliza.v1.SayRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sentence", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SayRequest {
    return new SayRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SayRequest {
    return new SayRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SayRequest {
    return new SayRequest().fromJsonString(jsonString, options);
  }

  static equals(a: SayRequest | PlainMessage<SayRequest> | undefined, b: SayRequest | PlainMessage<SayRequest> | undefined): boolean {
    return proto3.util.equals(SayRequest, a, b);
  }
}

/**
 * SayResponse describes the sentence responded by the ELIZA program.
 *
 * @generated from message eliza.v1.SayResponse
 */
export class SayResponse extends Message<SayResponse> {
  /**
   * @generated from field: string sentence = 1;
   */
  sentence = "";

  constructor(data?: PartialMessage<SayResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "eliza.v1.SayResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sentence", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SayResponse {
    return new SayResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SayResponse {
    return new SayResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SayResponse {
    return new SayResponse().fromJsonString(jsonString, options);
  }

  static equals(a: SayResponse | PlainMessage<SayResponse> | undefined, b: SayResponse | PlainMessage<SayResponse> | undefined): boolean {
    return proto3.util.equals(SayResponse, a, b);
  }
}

/**
 * ConverseRequest describes the sentence said to the ELIZA program.
 *
 * @generated from message eliza.v1.ConverseRequest
 */
export class ConverseRequest extends Message<ConverseRequest> {
  /**
   * @generated from field: string sentence = 1;
   */
  sentence = "";

  constructor(data?: PartialMessage<ConverseRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "eliza.v1.ConverseRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sentence", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ConverseRequest {
    return new ConverseRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ConverseRequest {
    return new ConverseRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ConverseRequest {
    return new ConverseRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ConverseRequest | PlainMessage<ConverseRequest> | undefined, b: ConverseRequest | PlainMessage<ConverseRequest> | undefined): boolean {
    return proto3.util.equals(ConverseRequest, a, b);
  }
}

/**
 * ConverseResponse describes the sentence responded by the ELIZA program.
 *
 * @generated from message eliza.v1.ConverseResponse
 */
export class ConverseResponse extends Message<ConverseResponse> {
  /**
   * @generated from field: string sentence = 1;
   */
  sentence = "";

  constructor(data?: PartialMessage<ConverseResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "eliza.v1.ConverseResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sentence", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ConverseResponse {
    return new ConverseResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ConverseResponse {
    return new ConverseResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ConverseResponse {
    return new ConverseResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ConverseResponse | PlainMessage<ConverseResponse> | undefined, b: ConverseResponse | PlainMessage<ConverseResponse> | undefined): boolean {
    return proto3.util.equals(ConverseResponse, a, b);
  }
}

/**
 * IntroduceRequest describes a request for details from the ELIZA program.
 *
 * @generated from message eliza.v1.IntroduceRequest
 */
export class IntroduceRequest extends Message<IntroduceRequest> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  constructor(data?: PartialMessage<IntroduceRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "eliza.v1.IntroduceRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): IntroduceRequest {
    return new IntroduceRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): IntroduceRequest {
    return new IntroduceRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): IntroduceRequest {
    return new IntroduceRequest().fromJsonString(jsonString, options);
  }

  static equals(a: IntroduceRequest | PlainMessage<IntroduceRequest> | undefined, b: IntroduceRequest | PlainMessage<IntroduceRequest> | undefined): boolean {
    return proto3.util.equals(IntroduceRequest, a, b);
  }
}

/**
 * IntroduceResponse describes the sentence responded by the ELIZA program.
 *
 * @generated from message eliza.v1.IntroduceResponse
 */
export class IntroduceResponse extends Message<IntroduceResponse> {
  /**
   * @generated from field: string sentence = 1;
   */
  sentence = "";

  constructor(data?: PartialMessage<IntroduceResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "eliza.v1.IntroduceResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sentence", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): IntroduceResponse {
    return new IntroduceResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): IntroduceResponse {
    return new IntroduceResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): IntroduceResponse {
    return new IntroduceResponse().fromJsonString(jsonString, options);
  }

  static equals(a: IntroduceResponse | PlainMessage<IntroduceResponse> | undefined, b: IntroduceResponse | PlainMessage<IntroduceResponse> | undefined): boolean {
    return proto3.util.equals(IntroduceResponse, a, b);
  }
}

