// @generated by protoc-gen-connect-web v0.3.3 with parameter "target=js+dts"
// @generated from file proto/uptime.proto (package uptime_dash.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateTargetRequest, CreateTargetResponse, DeleteTargetRequest, DeleteTargetResponse, GetAllTargetsRequest, GetAllTargetsResponse, GetTargetRequest, GetTargetResponse, UpdateTargetRequest, UpdateTargetResponse } from "./uptime_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service uptime_dash.v1.UptimeService
 */
export declare const UptimeService: {
  readonly typeName: "uptime_dash.v1.UptimeService",
  readonly methods: {
    /**
     * @generated from rpc uptime_dash.v1.UptimeService.GetAllTargets
     */
    readonly getAllTargets: {
      readonly name: "GetAllTargets",
      readonly I: typeof GetAllTargetsRequest,
      readonly O: typeof GetAllTargetsResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc uptime_dash.v1.UptimeService.GetTarget
     */
    readonly getTarget: {
      readonly name: "GetTarget",
      readonly I: typeof GetTargetRequest,
      readonly O: typeof GetTargetResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc uptime_dash.v1.UptimeService.CreateTarget
     */
    readonly createTarget: {
      readonly name: "CreateTarget",
      readonly I: typeof CreateTargetRequest,
      readonly O: typeof CreateTargetResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc uptime_dash.v1.UptimeService.UpdateTarget
     */
    readonly updateTarget: {
      readonly name: "UpdateTarget",
      readonly I: typeof UpdateTargetRequest,
      readonly O: typeof UpdateTargetResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc uptime_dash.v1.UptimeService.DeleteTarget
     */
    readonly deleteTarget: {
      readonly name: "DeleteTarget",
      readonly I: typeof DeleteTargetRequest,
      readonly O: typeof DeleteTargetResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

