// @generated by protoc-gen-es v0.4.0 with parameter "target=ts"
// @generated from file proto/uptime.proto (package uptime_dash.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message uptime_dash.v1.Target
 */
export class Target extends Message<Target> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: uptime_dash.v1.Target.Type type = 3;
   */
  type = Target_Type.UNSPECIFIED;

  /**
   * @generated from field: uint32 interval_seconds = 4;
   */
  intervalSeconds = 0;

  /**
   * @generated from field: uint32 timeout_seconds = 5;
   */
  timeoutSeconds = 0;

  /**
   * @generated from field: optional string url = 6;
   */
  url?: string;

  /**
   * @generated from field: optional string hostname = 7;
   */
  hostname?: string;

  /**
   * @generated from field: optional uint32 port = 8;
   */
  port?: number;

  constructor(data?: PartialMessage<Target>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "uptime_dash.v1.Target";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "type", kind: "enum", T: proto3.getEnumType(Target_Type) },
    { no: 4, name: "interval_seconds", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 5, name: "timeout_seconds", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 6, name: "url", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 7, name: "hostname", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 8, name: "port", kind: "scalar", T: 13 /* ScalarType.UINT32 */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Target {
    return new Target().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Target {
    return new Target().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Target {
    return new Target().fromJsonString(jsonString, options);
  }

  static equals(a: Target | PlainMessage<Target> | undefined, b: Target | PlainMessage<Target> | undefined): boolean {
    return proto3.util.equals(Target, a, b);
  }
}

/**
 * @generated from enum uptime_dash.v1.Target.Type
 */
export enum Target_Type {
  /**
   * @generated from enum value: TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: TYPE_HTTP = 1;
   */
  HTTP = 1,

  /**
   * @generated from enum value: TYPE_TCP = 2;
   */
  TCP = 2,

  /**
   * @generated from enum value: TYPE_PING = 3;
   */
  PING = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(Target_Type)
proto3.util.setEnumType(Target_Type, "uptime_dash.v1.Target.Type", [
  { no: 0, name: "TYPE_UNSPECIFIED" },
  { no: 1, name: "TYPE_HTTP" },
  { no: 2, name: "TYPE_TCP" },
  { no: 3, name: "TYPE_PING" },
]);

/**
 * @generated from message uptime_dash.v1.GetAllTargetsRequest
 */
export class GetAllTargetsRequest extends Message<GetAllTargetsRequest> {
  constructor(data?: PartialMessage<GetAllTargetsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "uptime_dash.v1.GetAllTargetsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetAllTargetsRequest {
    return new GetAllTargetsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetAllTargetsRequest {
    return new GetAllTargetsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetAllTargetsRequest {
    return new GetAllTargetsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetAllTargetsRequest | PlainMessage<GetAllTargetsRequest> | undefined, b: GetAllTargetsRequest | PlainMessage<GetAllTargetsRequest> | undefined): boolean {
    return proto3.util.equals(GetAllTargetsRequest, a, b);
  }
}

/**
 * @generated from message uptime_dash.v1.GetAllTargetsResponse
 */
export class GetAllTargetsResponse extends Message<GetAllTargetsResponse> {
  /**
   * @generated from field: repeated uptime_dash.v1.Target targets = 1;
   */
  targets: Target[] = [];

  constructor(data?: PartialMessage<GetAllTargetsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "uptime_dash.v1.GetAllTargetsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "targets", kind: "message", T: Target, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetAllTargetsResponse {
    return new GetAllTargetsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetAllTargetsResponse {
    return new GetAllTargetsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetAllTargetsResponse {
    return new GetAllTargetsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetAllTargetsResponse | PlainMessage<GetAllTargetsResponse> | undefined, b: GetAllTargetsResponse | PlainMessage<GetAllTargetsResponse> | undefined): boolean {
    return proto3.util.equals(GetAllTargetsResponse, a, b);
  }
}

/**
 * @generated from message uptime_dash.v1.GetTargetRequest
 */
export class GetTargetRequest extends Message<GetTargetRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<GetTargetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "uptime_dash.v1.GetTargetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTargetRequest {
    return new GetTargetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTargetRequest {
    return new GetTargetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTargetRequest {
    return new GetTargetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetTargetRequest | PlainMessage<GetTargetRequest> | undefined, b: GetTargetRequest | PlainMessage<GetTargetRequest> | undefined): boolean {
    return proto3.util.equals(GetTargetRequest, a, b);
  }
}

/**
 * @generated from message uptime_dash.v1.GetTargetResponse
 */
export class GetTargetResponse extends Message<GetTargetResponse> {
  /**
   * @generated from field: uptime_dash.v1.Target target = 1;
   */
  target?: Target;

  constructor(data?: PartialMessage<GetTargetResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "uptime_dash.v1.GetTargetResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "target", kind: "message", T: Target },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTargetResponse {
    return new GetTargetResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTargetResponse {
    return new GetTargetResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTargetResponse {
    return new GetTargetResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetTargetResponse | PlainMessage<GetTargetResponse> | undefined, b: GetTargetResponse | PlainMessage<GetTargetResponse> | undefined): boolean {
    return proto3.util.equals(GetTargetResponse, a, b);
  }
}

/**
 * @generated from message uptime_dash.v1.CreateTargetRequest
 */
export class CreateTargetRequest extends Message<CreateTargetRequest> {
  /**
   * @generated from field: uptime_dash.v1.Target target = 1;
   */
  target?: Target;

  constructor(data?: PartialMessage<CreateTargetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "uptime_dash.v1.CreateTargetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "target", kind: "message", T: Target },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateTargetRequest {
    return new CreateTargetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateTargetRequest {
    return new CreateTargetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateTargetRequest {
    return new CreateTargetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateTargetRequest | PlainMessage<CreateTargetRequest> | undefined, b: CreateTargetRequest | PlainMessage<CreateTargetRequest> | undefined): boolean {
    return proto3.util.equals(CreateTargetRequest, a, b);
  }
}

/**
 * @generated from message uptime_dash.v1.CreateTargetResponse
 */
export class CreateTargetResponse extends Message<CreateTargetResponse> {
  /**
   * @generated from field: uptime_dash.v1.Target target = 1;
   */
  target?: Target;

  constructor(data?: PartialMessage<CreateTargetResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "uptime_dash.v1.CreateTargetResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "target", kind: "message", T: Target },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateTargetResponse {
    return new CreateTargetResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateTargetResponse {
    return new CreateTargetResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateTargetResponse {
    return new CreateTargetResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateTargetResponse | PlainMessage<CreateTargetResponse> | undefined, b: CreateTargetResponse | PlainMessage<CreateTargetResponse> | undefined): boolean {
    return proto3.util.equals(CreateTargetResponse, a, b);
  }
}

/**
 * @generated from message uptime_dash.v1.UpdateTargetRequest
 */
export class UpdateTargetRequest extends Message<UpdateTargetRequest> {
  /**
   * @generated from field: uptime_dash.v1.Target target = 1;
   */
  target?: Target;

  constructor(data?: PartialMessage<UpdateTargetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "uptime_dash.v1.UpdateTargetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "target", kind: "message", T: Target },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateTargetRequest {
    return new UpdateTargetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateTargetRequest {
    return new UpdateTargetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateTargetRequest {
    return new UpdateTargetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateTargetRequest | PlainMessage<UpdateTargetRequest> | undefined, b: UpdateTargetRequest | PlainMessage<UpdateTargetRequest> | undefined): boolean {
    return proto3.util.equals(UpdateTargetRequest, a, b);
  }
}

/**
 * @generated from message uptime_dash.v1.UpdateTargetResponse
 */
export class UpdateTargetResponse extends Message<UpdateTargetResponse> {
  /**
   * @generated from field: uptime_dash.v1.Target target = 1;
   */
  target?: Target;

  constructor(data?: PartialMessage<UpdateTargetResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "uptime_dash.v1.UpdateTargetResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "target", kind: "message", T: Target },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateTargetResponse {
    return new UpdateTargetResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateTargetResponse {
    return new UpdateTargetResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateTargetResponse {
    return new UpdateTargetResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateTargetResponse | PlainMessage<UpdateTargetResponse> | undefined, b: UpdateTargetResponse | PlainMessage<UpdateTargetResponse> | undefined): boolean {
    return proto3.util.equals(UpdateTargetResponse, a, b);
  }
}

/**
 * @generated from message uptime_dash.v1.DeleteTargetRequest
 */
export class DeleteTargetRequest extends Message<DeleteTargetRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<DeleteTargetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "uptime_dash.v1.DeleteTargetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteTargetRequest {
    return new DeleteTargetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteTargetRequest {
    return new DeleteTargetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteTargetRequest {
    return new DeleteTargetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteTargetRequest | PlainMessage<DeleteTargetRequest> | undefined, b: DeleteTargetRequest | PlainMessage<DeleteTargetRequest> | undefined): boolean {
    return proto3.util.equals(DeleteTargetRequest, a, b);
  }
}

/**
 * @generated from message uptime_dash.v1.DeleteTargetResponse
 */
export class DeleteTargetResponse extends Message<DeleteTargetResponse> {
  constructor(data?: PartialMessage<DeleteTargetResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "uptime_dash.v1.DeleteTargetResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteTargetResponse {
    return new DeleteTargetResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteTargetResponse {
    return new DeleteTargetResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteTargetResponse {
    return new DeleteTargetResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteTargetResponse | PlainMessage<DeleteTargetResponse> | undefined, b: DeleteTargetResponse | PlainMessage<DeleteTargetResponse> | undefined): boolean {
    return proto3.util.equals(DeleteTargetResponse, a, b);
  }
}

