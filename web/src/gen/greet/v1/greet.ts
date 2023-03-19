/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "greet.v1";

export interface GreetRequest {
  name: string;
}

export interface GreetResponse {
  greeting: string;
}

function createBaseGreetRequest(): GreetRequest {
  return { name: "" };
}

export const GreetRequest = {
  encode(message: GreetRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GreetRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGreetRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GreetRequest {
    return { name: isSet(object.name) ? String(object.name) : "" };
  },

  toJSON(message: GreetRequest): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  create<I extends Exact<DeepPartial<GreetRequest>, I>>(base?: I): GreetRequest {
    return GreetRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GreetRequest>, I>>(object: I): GreetRequest {
    const message = createBaseGreetRequest();
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseGreetResponse(): GreetResponse {
  return { greeting: "" };
}

export const GreetResponse = {
  encode(message: GreetResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.greeting !== "") {
      writer.uint32(10).string(message.greeting);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GreetResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGreetResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.greeting = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GreetResponse {
    return { greeting: isSet(object.greeting) ? String(object.greeting) : "" };
  },

  toJSON(message: GreetResponse): unknown {
    const obj: any = {};
    message.greeting !== undefined && (obj.greeting = message.greeting);
    return obj;
  },

  create<I extends Exact<DeepPartial<GreetResponse>, I>>(base?: I): GreetResponse {
    return GreetResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GreetResponse>, I>>(object: I): GreetResponse {
    const message = createBaseGreetResponse();
    message.greeting = object.greeting ?? "";
    return message;
  },
};

export interface GreetService {
  Greet(request: GreetRequest): Promise<GreetResponse>;
}

export class GreetServiceClientImpl implements GreetService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "greet.v1.GreetService";
    this.rpc = rpc;
    this.Greet = this.Greet.bind(this);
  }
  Greet(request: GreetRequest): Promise<GreetResponse> {
    const data = GreetRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Greet", data);
    return promise.then((data) => GreetResponse.decode(_m0.Reader.create(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
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
