// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var er_service_entity_v1_entity_service_pb = require('../../../../er/service/entity/v1/entity_service_pb.js');
var er_schema_entity_v1_entity_pb = require('../../../../er/schema/entity/v1/entity_pb.js');

function serialize_er_schema_entity_v1_Member(arg) {
  if (!(arg instanceof er_schema_entity_v1_entity_pb.Member)) {
    throw new Error('Expected argument of type er.schema.entity.v1.Member');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_er_schema_entity_v1_Member(buffer_arg) {
  return er_schema_entity_v1_entity_pb.Member.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_er_service_entity_v1_Response(arg) {
  if (!(arg instanceof er_service_entity_v1_entity_service_pb.Response)) {
    throw new Error('Expected argument of type er.service.entity.v1.Response');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_er_service_entity_v1_Response(buffer_arg) {
  return er_service_entity_v1_entity_service_pb.Response.deserializeBinary(new Uint8Array(buffer_arg));
}


// EntityService API
var EntityServiceService = exports.EntityServiceService = {
  ingest: {
    path: '/er.service.entity.v1.EntityService/Ingest',
    requestStream: false,
    responseStream: false,
    requestType: er_schema_entity_v1_entity_pb.Member,
    responseType: er_service_entity_v1_entity_service_pb.Response,
    requestSerialize: serialize_er_schema_entity_v1_Member,
    requestDeserialize: deserialize_er_schema_entity_v1_Member,
    responseSerialize: serialize_er_service_entity_v1_Response,
    responseDeserialize: deserialize_er_service_entity_v1_Response,
  },
  // rpc IngestStream(stream er.schema.entity.v1.Member) returns (stream Response) {}
inquiry: {
    path: '/er.service.entity.v1.EntityService/Inquiry',
    requestStream: false,
    responseStream: false,
    requestType: er_schema_entity_v1_entity_pb.Member,
    responseType: er_service_entity_v1_entity_service_pb.Response,
    requestSerialize: serialize_er_schema_entity_v1_Member,
    requestDeserialize: deserialize_er_schema_entity_v1_Member,
    responseSerialize: serialize_er_service_entity_v1_Response,
    responseDeserialize: deserialize_er_service_entity_v1_Response,
  },
  // rpc InquiryStream(stream er.schema.entity.v1.Member) returns (stream Response) {}
};

exports.EntityServiceClient = grpc.makeGenericClientConstructor(EntityServiceService);
