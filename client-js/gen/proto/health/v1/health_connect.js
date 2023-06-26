// @generated by protoc-gen-connect-es v0.10.1 with parameter "target=js"
// @generated from file proto/health/v1/health.proto (package proto.health.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CheckHealthRequest, CheckHealthResponse } from "./health_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service proto.health.v1.HealthService
 */
export const HealthService = {
  typeName: "proto.health.v1.HealthService",
  methods: {
    /**
     * @generated from rpc proto.health.v1.HealthService.CheckHealth
     */
    checkHealth: {
      name: "CheckHealth",
      I: CheckHealthRequest,
      O: CheckHealthResponse,
      kind: MethodKind.Unary,
    },
  }
};

