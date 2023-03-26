/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Observable } from "rxjs";
import { map } from "rxjs/operators";

export const protobufPackage = "eliza.v1";

/** SayRequest describes the sentence said to the ELIZA program. */
export interface SayRequest {
  sentence: string;
}

/** SayResponse describes the sentence responded by the ELIZA program. */
export interface SayResponse {
  sentence: string;
}

/** ConverseRequest describes the sentence said to the ELIZA program. */
export interface ConverseRequest {
  sentence: string;
}

/** ConverseResponse describes the sentence responded by the ELIZA program. */
export interface ConverseResponse {
  sentence: string;
}

/** IntroduceRequest describes a request for details from the ELIZA program. */
export interface IntroduceRequest {
  name: string;
}

/** IntroduceResponse describes the sentence responded by the ELIZA program. */
export interface IntroduceResponse {
  sentence: string;
}

function createBaseSayRequest(): SayRequest {
  return { sentence: "" };
}

export const SayRequest = {
  encode(message: SayRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sentence !== "") {
      writer.uint32(10).string(message.sentence);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SayRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSayRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.sentence = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SayRequest {
    return { sentence: isSet(object.sentence) ? String(object.sentence) : "" };
  },

  toJSON(message: SayRequest): unknown {
    const obj: any = {};
    message.sentence !== undefined && (obj.sentence = message.sentence);
    return obj;
  },

  create<I extends Exact<DeepPartial<SayRequest>, I>>(base?: I): SayRequest {
    return SayRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SayRequest>, I>>(object: I): SayRequest {
    const message = createBaseSayRequest();
    message.sentence = object.sentence ?? "";
    return message;
  },
};

function createBaseSayResponse(): SayResponse {
  return { sentence: "" };
}

export const SayResponse = {
  encode(message: SayResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sentence !== "") {
      writer.uint32(10).string(message.sentence);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SayResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSayResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.sentence = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SayResponse {
    return { sentence: isSet(object.sentence) ? String(object.sentence) : "" };
  },

  toJSON(message: SayResponse): unknown {
    const obj: any = {};
    message.sentence !== undefined && (obj.sentence = message.sentence);
    return obj;
  },

  create<I extends Exact<DeepPartial<SayResponse>, I>>(base?: I): SayResponse {
    return SayResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SayResponse>, I>>(object: I): SayResponse {
    const message = createBaseSayResponse();
    message.sentence = object.sentence ?? "";
    return message;
  },
};

function createBaseConverseRequest(): ConverseRequest {
  return { sentence: "" };
}

export const ConverseRequest = {
  encode(message: ConverseRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sentence !== "") {
      writer.uint32(10).string(message.sentence);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ConverseRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseConverseRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.sentence = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ConverseRequest {
    return { sentence: isSet(object.sentence) ? String(object.sentence) : "" };
  },

  toJSON(message: ConverseRequest): unknown {
    const obj: any = {};
    message.sentence !== undefined && (obj.sentence = message.sentence);
    return obj;
  },

  create<I extends Exact<DeepPartial<ConverseRequest>, I>>(base?: I): ConverseRequest {
    return ConverseRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ConverseRequest>, I>>(object: I): ConverseRequest {
    const message = createBaseConverseRequest();
    message.sentence = object.sentence ?? "";
    return message;
  },
};

function createBaseConverseResponse(): ConverseResponse {
  return { sentence: "" };
}

export const ConverseResponse = {
  encode(message: ConverseResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sentence !== "") {
      writer.uint32(10).string(message.sentence);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ConverseResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseConverseResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.sentence = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ConverseResponse {
    return { sentence: isSet(object.sentence) ? String(object.sentence) : "" };
  },

  toJSON(message: ConverseResponse): unknown {
    const obj: any = {};
    message.sentence !== undefined && (obj.sentence = message.sentence);
    return obj;
  },

  create<I extends Exact<DeepPartial<ConverseResponse>, I>>(base?: I): ConverseResponse {
    return ConverseResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ConverseResponse>, I>>(object: I): ConverseResponse {
    const message = createBaseConverseResponse();
    message.sentence = object.sentence ?? "";
    return message;
  },
};

function createBaseIntroduceRequest(): IntroduceRequest {
  return { name: "" };
}

export const IntroduceRequest = {
  encode(message: IntroduceRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IntroduceRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIntroduceRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.name = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): IntroduceRequest {
    return { name: isSet(object.name) ? String(object.name) : "" };
  },

  toJSON(message: IntroduceRequest): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  create<I extends Exact<DeepPartial<IntroduceRequest>, I>>(base?: I): IntroduceRequest {
    return IntroduceRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<IntroduceRequest>, I>>(object: I): IntroduceRequest {
    const message = createBaseIntroduceRequest();
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseIntroduceResponse(): IntroduceResponse {
  return { sentence: "" };
}

export const IntroduceResponse = {
  encode(message: IntroduceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sentence !== "") {
      writer.uint32(10).string(message.sentence);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IntroduceResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIntroduceResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.sentence = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): IntroduceResponse {
    return { sentence: isSet(object.sentence) ? String(object.sentence) : "" };
  },

  toJSON(message: IntroduceResponse): unknown {
    const obj: any = {};
    message.sentence !== undefined && (obj.sentence = message.sentence);
    return obj;
  },

  create<I extends Exact<DeepPartial<IntroduceResponse>, I>>(base?: I): IntroduceResponse {
    return IntroduceResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<IntroduceResponse>, I>>(object: I): IntroduceResponse {
    const message = createBaseIntroduceResponse();
    message.sentence = object.sentence ?? "";
    return message;
  },
};

/**
 * ElizaService provides a way to talk to the ELIZA, which is a port of
 * the DOCTOR script for Joseph Weizenbaum's original ELIZA program.
 * Created in the mid-1960s at the MIT Artificial Intelligence Laboratory,
 * ELIZA demonstrates the superficiality of human-computer communication.
 * DOCTOR simulates a psychotherapist, and is commonly found as an Easter
 * egg in emacs distributions.
 */
export interface ElizaService {
  /**
   * Say is a unary request demo. This method should allow for a one sentence
   * response given a one sentence request.
   */
  Say(request: SayRequest): Promise<SayResponse>;
  /**
   * Converse is a bi-directional streaming request demo. This method should allow for
   * many requests and many responses.
   */
  Converse(request: Observable<ConverseRequest>): Observable<ConverseResponse>;
  /**
   * Introduce is a server-streaming request demo.  This method allows for a single request that will return a series
   * of responses
   */
  Introduce(request: IntroduceRequest): Observable<IntroduceResponse>;
}

export class ElizaServiceClientImpl implements ElizaService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "eliza.v1.ElizaService";
    this.rpc = rpc;
    this.Say = this.Say.bind(this);
    this.Converse = this.Converse.bind(this);
    this.Introduce = this.Introduce.bind(this);
  }
  Say(request: SayRequest): Promise<SayResponse> {
    const data = SayRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Say", data);
    return promise.then((data) => SayResponse.decode(_m0.Reader.create(data)));
  }

  Converse(request: Observable<ConverseRequest>): Observable<ConverseResponse> {
    const data = request.pipe(map((request) => ConverseRequest.encode(request).finish()));
    const result = this.rpc.bidirectionalStreamingRequest(this.service, "Converse", data);
    return result.pipe(map((data) => ConverseResponse.decode(_m0.Reader.create(data))));
  }

  Introduce(request: IntroduceRequest): Observable<IntroduceResponse> {
    const data = IntroduceRequest.encode(request).finish();
    const result = this.rpc.serverStreamingRequest(this.service, "Introduce", data);
    return result.pipe(map((data) => IntroduceResponse.decode(_m0.Reader.create(data))));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
  clientStreamingRequest(service: string, method: string, data: Observable<Uint8Array>): Promise<Uint8Array>;
  serverStreamingRequest(service: string, method: string, data: Uint8Array): Observable<Uint8Array>;
  bidirectionalStreamingRequest(service: string, method: string, data: Observable<Uint8Array>): Observable<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
