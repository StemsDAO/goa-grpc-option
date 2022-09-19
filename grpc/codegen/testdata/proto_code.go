package testdata

const UnaryRPCsProtoCode = `
syntax = "proto3";

package service_unary_rp_cs;

option go_package = "/service_unary_rp_cspb";

// Service is the ServiceUnaryRPCs service interface.
service ServiceUnaryRPCs {
	// MethodUnaryRPCA implements MethodUnaryRPCA.
	rpc MethodUnaryRPCA (MethodUnaryRPCARequest) returns (MethodUnaryRPCAResponse);
	// MethodUnaryRPCB implements MethodUnaryRPCB.
	rpc MethodUnaryRPCB (MethodUnaryRPCBRequest) returns (MethodUnaryRPCBResponse);
}

message MethodUnaryRPCARequest {
	sint32 int = 1;
	string string_ = 2;
}

message MethodUnaryRPCAResponse {
	repeated bool array_field = 1;
	map<string, double> map_field = 2;
}

message MethodUnaryRPCBRequest {
	uint32 u_int = 1;
	float float32 = 2;
}

message MethodUnaryRPCBResponse {
	repeated bool array_field = 1;
	map<string, double> map_field = 2;
}
`

const UnaryRPCNoPayloadProtoCode = `
syntax = "proto3";

package service_unary_rpc_no_payload;

option go_package = "/service_unary_rpc_no_payloadpb";

// Service is the ServiceUnaryRPCNoPayload service interface.
service ServiceUnaryRPCNoPayload {
	// MethodUnaryRPCNoPayload implements MethodUnaryRPCNoPayload.
	rpc MethodUnaryRPCNoPayload (MethodUnaryRPCNoPayloadRequest) returns (MethodUnaryRPCNoPayloadResponse);
}

message MethodUnaryRPCNoPayloadRequest {
}

message MethodUnaryRPCNoPayloadResponse {
	string field = 1;
}
`

const UnaryRPCNoResultProtoCode = `
syntax = "proto3";

package service_unary_rpc_no_result;

option go_package = "/service_unary_rpc_no_resultpb";

// Service is the ServiceUnaryRPCNoResult service interface.
service ServiceUnaryRPCNoResult {
	// MethodUnaryRPCNoResult implements MethodUnaryRPCNoResult.
	rpc MethodUnaryRPCNoResult (MethodUnaryRPCNoResultRequest) returns (MethodUnaryRPCNoResultResponse);
}

message MethodUnaryRPCNoResultRequest {
	repeated string field = 1;
}

message MethodUnaryRPCNoResultResponse {
}
`

const ServerStreamingRPCProtoCode = `
syntax = "proto3";

package service_server_streaming_rpc;

option go_package = "/service_server_streaming_rpcpb";

// Service is the ServiceServerStreamingRPC service interface.
service ServiceServerStreamingRPC {
	// MethodServerStreamingRPC implements MethodServerStreamingRPC.
	rpc MethodServerStreamingRPC (MethodServerStreamingRPCRequest) returns (stream MethodServerStreamingRPCResponse);
}

message MethodServerStreamingRPCRequest {
	sint32 field = 1;
}

message MethodServerStreamingRPCResponse {
	string field = 1;
}
`

const ClientStreamingRPCProtoCode = `
syntax = "proto3";

package service_client_streaming_rpc;

option go_package = "/service_client_streaming_rpcpb";

// Service is the ServiceClientStreamingRPC service interface.
service ServiceClientStreamingRPC {
	// MethodClientStreamingRPC implements MethodClientStreamingRPC.
	rpc MethodClientStreamingRPC (stream MethodClientStreamingRPCStreamingRequest) returns (MethodClientStreamingRPCResponse);
}

message MethodClientStreamingRPCStreamingRequest {
	sint32 field = 1;
}

message MethodClientStreamingRPCResponse {
	string field = 1;
}
`

const BidirectionalStreamingRPCProtoCode = `
syntax = "proto3";

package service_bidirectional_streaming_rpc;

option go_package = "/service_bidirectional_streaming_rpcpb";

// Service is the ServiceBidirectionalStreamingRPC service interface.
service ServiceBidirectionalStreamingRPC {
	// MethodBidirectionalStreamingRPC implements MethodBidirectionalStreamingRPC.
	rpc MethodBidirectionalStreamingRPC (stream MethodBidirectionalStreamingRPCStreamingRequest) returns (stream MethodBidirectionalStreamingRPCResponse);
}

message MethodBidirectionalStreamingRPCStreamingRequest {
	sint32 field = 1;
}

message MethodBidirectionalStreamingRPCResponse {
	sint32 a = 1;
	string b = 2;
}
`

const MessageWithServiceNameProtoCode = `
syntax = "proto3";

package my_name_conflicts;

option go_package = "/my_name_conflictspb";

// Service is the MyNameConflicts service interface.
service MyNameConflicts {
	// MyNameConflictsMethod implements MyNameConflictsMethod.
	rpc MyNameConflictsMethod (MyNameConflictsMethodRequest) returns (MyNameConflictsMethodResponse);
}

message MyNameConflictsMethodRequest {
	MyNameConflicts2 conflict = 1;
}

message MyNameConflicts2 {
	bool boolean_field = 1;
}

message MyNameConflictsMethodResponse {
}
`

const MessageUserTypeWithPrimitivesMessageCode = `
message MethodMessageUserTypeWithPrimitivesRequest {
	bool boolean_field = 1;
	sint32 int_field = 2;
	sint32 int32_field = 3;
	sint64 int64_field = 4;
	uint32 u_int_field = 5;
	uint32 u_int32_field = 6;
	uint64 u_int64_field = 7;
}

message MethodMessageUserTypeWithPrimitivesResponse {
	float float32_field = 1;
	double float64_field = 2;
	string string_field = 3;
	bytes bytes_field = 4;
}
`

const MessageUserTypeWithAliasMessageCode = `
message MethodMessageUserTypeWithAliasRequest {
	sint32 int_alias_field = 1;
	sint32 optional_int_alias_field = 2;
}

message MethodMessageUserTypeWithAliasResponse {
	sint32 int_alias_field = 1;
	sint32 optional_int_alias_field = 2;
}
`

const MessageUserTypeWithNestedUserTypesCode = `
message MethodMessageUserTypeWithNestedUserTypesRequest {
	bool boolean_field = 1;
	sint32 int_field = 2;
	UTLevel1 ut_level1 = 3;
}

message UTLevel1 {
	sint32 int32_field = 1;
	sint64 int64_field = 2;
	UTLevel2 ut_level2 = 3;
}

message UTLevel2 {
	sint64 int64_field = 2;
}

message MethodMessageUserTypeWithNestedUserTypesResponse {
	RecursiveT recursive = 1;
}

message RecursiveT {
	RecursiveT recursive = 1;
}
`

const MessageResultTypeListCode = `
message MethodMessageUserTypeWithNestedUserTypesRequest {
}

message RTList {
	repeated RT field = 1;
}

message RT {
	sint32 int_field = 1;
	string string_field = 2;
}
`

const MessageUserTypeWithListCode = `
message MethodMessageUserTypeWithPrimitivesRequest {
}

message MethodMessageUserTypeWithPrimitivesResponse {
	RTList list_field = 1;
}

message RTList {
	repeated RT field = 1;
}

message RT {
	sint32 int_field = 1;
}
`

const MessageArrayCode = `
message MethodMessageArrayRequest {
	repeated uint32 array_of_primitives = 1;
	repeated ArrayOfBytes two_d_array = 2;
	repeated ArrayOfArrayOfBytes three_d_array = 3;
	repeated MapOfStringDouble array_of_maps = 4;
}

message ArrayOfBytes {
	repeated bytes field = 1;
}

message ArrayOfArrayOfBytes {
	repeated ArrayOfBytes field = 1;
}

message MapOfStringDouble {
	map<string, double> field = 1;
}

message MethodMessageArrayResponse {
	repeated UT field = 1;
}

message UT {
	repeated uint32 array_of_primitives = 1;
	repeated ArrayOfBytes two_d_array = 2;
	repeated ArrayOfArrayOfBytes three_d_array = 3;
	repeated MapOfStringDouble array_of_maps = 4;
}
`

const MessageMapCode = `
message MethodMessageMapRequest {
	map<sint32, UT> field = 1;
}

message UT {
	map<uint32, bool> map_of_primitives = 1;
	map<sint32, ArrayOfUTLevel1> map_of_primitive_ut_array = 2;
}

message ArrayOfUTLevel1 {
	repeated UTLevel1 field = 1;
}

message UTLevel1 {
	map<string, MapOfSint32Uint32> map_of_map_of_primitives = 1;
}

message MapOfSint32Uint32 {
	map<sint32, uint32> field = 1;
}

message MethodMessageMapResponse {
	map<uint32, bool> map_of_primitives = 1;
	map<sint32, ArrayOfUTLevel1> map_of_primitive_ut_array = 2;
}
`

const MessagePrimitiveCode = `
message MethodMessagePrimitiveRequest {
	uint32 field = 1;
}

message MethodMessagePrimitiveResponse {
	sint32 field = 1;
}
`

const MessageWithMetadataCode = `
message MethodMessageWithMetadataRequest {
	bool boolean_field = 1;
	UTLevel1 ut_level1 = 3;
}

message UTLevel1 {
	sint32 int32_field = 1;
	sint64 int64_field = 2;
}

message MethodMessageWithMetadataResponse {
	UTLevel1 ut_level1 = 3;
}
`

const MessageWithSecurityAttrsCode = `
message MethodMessageWithSecurityRequest {
	string oauth_token = 3;
	bool boolean_field = 1;
}

message MethodMessageWithSecurityResponse {
}
`

const MethodWithReservedNameProtoCode = `
syntax = "proto3";

package method_with_reserved_name;

option go_package = "/method_with_reserved_namepb";

// Service is the MethodWithReservedName service interface.
service MethodWithReservedName {
	// String implements string.
	rpc String (StringRequest) returns (StringResponse);
}

message StringRequest {
}

message StringResponse {
}
`

const MultipleMethodsSameResultListProtoCode = `
syntax = "proto3";

package multiple_methods_same_result_list;

option go_package = "/multiple_methods_same_result_listpb";

// Service is the MultipleMethodsSameResultList service interface.
service MultipleMethodsSameResultList {
	// MethodA implements method_a.
	rpc MethodA (MethodARequest) returns (ResultTList);
	// MethodB implements method_b.
	rpc MethodB (MethodBRequest) returns (ResultTList);
}

message MethodARequest {
}

message ResultTList {
	repeated ResultT field = 1;
}

message ResultT {
	bool boolean_field = 1;
}

message MethodBRequest {
}
`

const MethodWithAcronymProtoCode = `
syntax = "proto3";

package method_with_acronym;

option go_package = "/method_with_acronympb";

// Service is the MethodWithAcronym service interface.
service MethodWithAcronym {
	// MethodJWT implements method_jwt.
	rpc MethodJWT (MethodJWTRequest) returns (MethodJWTResponse);
}

message MethodJWTRequest {
}

message MethodJWTResponse {
}
`

const ServiceWithPackageCode = `
syntax = "proto3";

package custom;

option go_package = "/custompb";

// Service is the ServiceWithPackageName service interface.
service ServiceWithPackageName {
	// Method implements method.
	rpc Method (MethodRequest) returns (MethodResponse);
}

message MethodRequest {
}

message MethodResponse {
}
`

const StructMetaTypePackageCode = `
syntax = "proto3";

package using_meta_types;

option go_package = "/using_meta_typespb";

// Service is the UsingMetaTypes service interface.
service UsingMetaTypes {
	// Method implements Method.
	rpc Method (MethodRequest) returns (MethodResponse);
}

message MethodRequest {
	sint64 a = 1;
	sint64 b = 2;
	repeated sint64 c = 3;
	sint64 d = 4;
}

message MethodResponse {
	sint64 a = 1;
	sint64 b = 2;
	repeated sint64 c = 3;
	sint64 d = 4;
}
`

const DefaultFieldsPackageCode = `
syntax = "proto3";

package default_fields;

option go_package = "/default_fieldspb";

// Service is the DefaultFields service interface.
service DefaultFields {
	// Method implements Method.
	rpc Method (MethodRequest) returns (MethodResponse);
}

message MethodRequest {
	sint64 req = 1;
	sint64 opt = 2;
	sint64 def0 = 3;
	sint64 def1 = 4;
	sint64 def2 = 5;
	string reqs = 6;
	string opts = 7;
	string defs = 8;
	string defe = 9;
	double rat = 10;
	double flt = 11;
	double flt0 = 12;
	double flt1 = 13;
}

message MethodResponse {
}
`
