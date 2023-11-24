// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: base/v1/errors.proto

package basev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorCode int32

const (
	ErrorCode_ERROR_CODE_UNSPECIFIED ErrorCode = 0
	// authn
	ErrorCode_ERROR_CODE_MISSING_BEARER_TOKEN ErrorCode = 1001
	ErrorCode_ERROR_CODE_UNAUTHENTICATED      ErrorCode = 1002
	ErrorCode_ERROR_CODE_MISSING_TENANT_ID    ErrorCode = 1003
	// validation
	ErrorCode_ERROR_CODE_VALIDATION                                        ErrorCode = 2000
	ErrorCode_ERROR_CODE_UNDEFINED_CHILD_TYPE                              ErrorCode = 2002
	ErrorCode_ERROR_CODE_UNDEFINED_CHILD_KIND                              ErrorCode = 2003
	ErrorCode_ERROR_CODE_UNDEFINED_RELATION_REFERENCE                      ErrorCode = 2006
	ErrorCode_ERROR_CODE_NOT_SUPPORTED_RELATION_WALK                       ErrorCode = 2007
	ErrorCode_ERROR_CODE_ENTITY_AND_SUBJECT_CANNOT_BE_EQUAL                ErrorCode = 2008
	ErrorCode_ERROR_CODE_DEPTH_NOT_ENOUGH                                  ErrorCode = 2009
	ErrorCode_ERROR_CODE_RELATION_REFERENCE_NOT_FOUND_IN_ENTITY_REFERENCES ErrorCode = 2010
	ErrorCode_ERROR_CODE_RELATION_REFERENCE_MUST_HAVE_ONE_ENTITY_REFERENCE ErrorCode = 2011
	ErrorCode_ERROR_CODE_DUPLICATED_ENTITY_REFERENCE                       ErrorCode = 2012
	ErrorCode_ERROR_CODE_DUPLICATED_RELATION_REFERENCE                     ErrorCode = 2013
	ErrorCode_ERROR_CODE_DUPLICATED_PERMISSION_REFERENCE                   ErrorCode = 2014
	ErrorCode_ERROR_CODE_SCHEMA_PARSE                                      ErrorCode = 2015
	ErrorCode_ERROR_CODE_SCHEMA_COMPILE                                    ErrorCode = 2016
	ErrorCode_ERROR_CODE_SUBJECT_RELATION_MUST_BE_EMPTY                    ErrorCode = 2017
	ErrorCode_ERROR_CODE_SUBJECT_RELATION_CANNOT_BE_EMPTY                  ErrorCode = 2018
	ErrorCode_ERROR_CODE_SCHEMA_MUST_HAVE_USER_ENTITY_DEFINITION           ErrorCode = 2019
	ErrorCode_ERROR_CODE_UNIQUE_CONSTRAINT                                 ErrorCode = 2020
	ErrorCode_ERROR_CODE_INVALID_CONTINUOUS_TOKEN                          ErrorCode = 2021
	ErrorCode_ERROR_CODE_INVALID_KEY                                       ErrorCode = 2022
	ErrorCode_ERROR_CODE_ENTITY_TYPE_REQUIRED                              ErrorCode = 2023
	ErrorCode_ERROR_CODE_NO_ENTITY_REFERENCES_FOUND_IN_SCHEMA              ErrorCode = 2024
	ErrorCode_ERROR_CODE_INVALID_ARGUMENT                                  ErrorCode = 2025
	ErrorCode_ERROR_CODE_INVALID_RULE_REFERENCE                            ErrorCode = 2026
	ErrorCode_ERROR_CODE_NOT_SUPPORTED_WALK                                ErrorCode = 2027
	ErrorCode_ERROR_CODE_MISSING_ARGUMENT                                  ErrorCode = 2028
	ErrorCode_ERROR_CODE_ALREADY_EXIST                                     ErrorCode = 2029
	// not found
	ErrorCode_ERROR_CODE_NOT_FOUND                       ErrorCode = 4000
	ErrorCode_ERROR_CODE_ENTITY_TYPE_NOT_FOUND           ErrorCode = 4001
	ErrorCode_ERROR_CODE_PERMISSION_NOT_FOUND            ErrorCode = 4002
	ErrorCode_ERROR_CODE_SCHEMA_NOT_FOUND                ErrorCode = 4003
	ErrorCode_ERROR_CODE_SUBJECT_TYPE_NOT_FOUND          ErrorCode = 4004
	ErrorCode_ERROR_CODE_ENTITY_DEFINITION_NOT_FOUND     ErrorCode = 4005
	ErrorCode_ERROR_CODE_PERMISSION_DEFINITION_NOT_FOUND ErrorCode = 4006
	ErrorCode_ERROR_CODE_RELATION_DEFINITION_NOT_FOUND   ErrorCode = 4007
	ErrorCode_ERROR_CODE_RECORD_NOT_FOUND                ErrorCode = 4008
	ErrorCode_ERROR_CODE_TENANT_NOT_FOUND                ErrorCode = 4009
	ErrorCode_ERROR_CODE_ATTRIBUTE_DEFINITION_NOT_FOUND  ErrorCode = 4010
	ErrorCode_ERROR_CODE_ATTRIBUTE_TYPE_MISMATCH         ErrorCode = 4011
	// internal
	ErrorCode_ERROR_CODE_INTERNAL                                  ErrorCode = 5000
	ErrorCode_ERROR_CODE_CANCELLED                                 ErrorCode = 5001
	ErrorCode_ERROR_CODE_SQL_BUILDER                               ErrorCode = 5002
	ErrorCode_ERROR_CODE_CIRCUIT_BREAKER                           ErrorCode = 5003
	ErrorCode_ERROR_CODE_EXECUTION                                 ErrorCode = 5005
	ErrorCode_ERROR_CODE_SCAN                                      ErrorCode = 5006
	ErrorCode_ERROR_CODE_MIGRATION                                 ErrorCode = 5007
	ErrorCode_ERROR_CODE_TYPE_CONVERSATION                         ErrorCode = 5008
	ErrorCode_ERROR_CODE_ERROR_MAX_RETRIES                         ErrorCode = 5009
	ErrorCode_ERROR_CODE_ROLLBACK                                  ErrorCode = 5010
	ErrorCode_ERROR_CODE_EXCLUSION_REQUIRES_MORE_THAN_ONE_FUNCTION ErrorCode = 5011
	ErrorCode_ERROR_CODE_NOT_IMPLEMENTED                           ErrorCode = 5012
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0:    "ERROR_CODE_UNSPECIFIED",
		1001: "ERROR_CODE_MISSING_BEARER_TOKEN",
		1002: "ERROR_CODE_UNAUTHENTICATED",
		1003: "ERROR_CODE_MISSING_TENANT_ID",
		2000: "ERROR_CODE_VALIDATION",
		2002: "ERROR_CODE_UNDEFINED_CHILD_TYPE",
		2003: "ERROR_CODE_UNDEFINED_CHILD_KIND",
		2006: "ERROR_CODE_UNDEFINED_RELATION_REFERENCE",
		2007: "ERROR_CODE_NOT_SUPPORTED_RELATION_WALK",
		2008: "ERROR_CODE_ENTITY_AND_SUBJECT_CANNOT_BE_EQUAL",
		2009: "ERROR_CODE_DEPTH_NOT_ENOUGH",
		2010: "ERROR_CODE_RELATION_REFERENCE_NOT_FOUND_IN_ENTITY_REFERENCES",
		2011: "ERROR_CODE_RELATION_REFERENCE_MUST_HAVE_ONE_ENTITY_REFERENCE",
		2012: "ERROR_CODE_DUPLICATED_ENTITY_REFERENCE",
		2013: "ERROR_CODE_DUPLICATED_RELATION_REFERENCE",
		2014: "ERROR_CODE_DUPLICATED_PERMISSION_REFERENCE",
		2015: "ERROR_CODE_SCHEMA_PARSE",
		2016: "ERROR_CODE_SCHEMA_COMPILE",
		2017: "ERROR_CODE_SUBJECT_RELATION_MUST_BE_EMPTY",
		2018: "ERROR_CODE_SUBJECT_RELATION_CANNOT_BE_EMPTY",
		2019: "ERROR_CODE_SCHEMA_MUST_HAVE_USER_ENTITY_DEFINITION",
		2020: "ERROR_CODE_UNIQUE_CONSTRAINT",
		2021: "ERROR_CODE_INVALID_CONTINUOUS_TOKEN",
		2022: "ERROR_CODE_INVALID_KEY",
		2023: "ERROR_CODE_ENTITY_TYPE_REQUIRED",
		2024: "ERROR_CODE_NO_ENTITY_REFERENCES_FOUND_IN_SCHEMA",
		2025: "ERROR_CODE_INVALID_ARGUMENT",
		2026: "ERROR_CODE_INVALID_RULE_REFERENCE",
		2027: "ERROR_CODE_NOT_SUPPORTED_WALK",
		2028: "ERROR_CODE_MISSING_ARGUMENT",
		2029: "ERROR_CODE_ALREADY_EXIST",
		4000: "ERROR_CODE_NOT_FOUND",
		4001: "ERROR_CODE_ENTITY_TYPE_NOT_FOUND",
		4002: "ERROR_CODE_PERMISSION_NOT_FOUND",
		4003: "ERROR_CODE_SCHEMA_NOT_FOUND",
		4004: "ERROR_CODE_SUBJECT_TYPE_NOT_FOUND",
		4005: "ERROR_CODE_ENTITY_DEFINITION_NOT_FOUND",
		4006: "ERROR_CODE_PERMISSION_DEFINITION_NOT_FOUND",
		4007: "ERROR_CODE_RELATION_DEFINITION_NOT_FOUND",
		4008: "ERROR_CODE_RECORD_NOT_FOUND",
		4009: "ERROR_CODE_TENANT_NOT_FOUND",
		4010: "ERROR_CODE_ATTRIBUTE_DEFINITION_NOT_FOUND",
		4011: "ERROR_CODE_ATTRIBUTE_TYPE_MISMATCH",
		5000: "ERROR_CODE_INTERNAL",
		5001: "ERROR_CODE_CANCELLED",
		5002: "ERROR_CODE_SQL_BUILDER",
		5003: "ERROR_CODE_CIRCUIT_BREAKER",
		5005: "ERROR_CODE_EXECUTION",
		5006: "ERROR_CODE_SCAN",
		5007: "ERROR_CODE_MIGRATION",
		5008: "ERROR_CODE_TYPE_CONVERSATION",
		5009: "ERROR_CODE_ERROR_MAX_RETRIES",
		5010: "ERROR_CODE_ROLLBACK",
		5011: "ERROR_CODE_EXCLUSION_REQUIRES_MORE_THAN_ONE_FUNCTION",
		5012: "ERROR_CODE_NOT_IMPLEMENTED",
	}
	ErrorCode_value = map[string]int32{
		"ERROR_CODE_UNSPECIFIED":                                       0,
		"ERROR_CODE_MISSING_BEARER_TOKEN":                              1001,
		"ERROR_CODE_UNAUTHENTICATED":                                   1002,
		"ERROR_CODE_MISSING_TENANT_ID":                                 1003,
		"ERROR_CODE_VALIDATION":                                        2000,
		"ERROR_CODE_UNDEFINED_CHILD_TYPE":                              2002,
		"ERROR_CODE_UNDEFINED_CHILD_KIND":                              2003,
		"ERROR_CODE_UNDEFINED_RELATION_REFERENCE":                      2006,
		"ERROR_CODE_NOT_SUPPORTED_RELATION_WALK":                       2007,
		"ERROR_CODE_ENTITY_AND_SUBJECT_CANNOT_BE_EQUAL":                2008,
		"ERROR_CODE_DEPTH_NOT_ENOUGH":                                  2009,
		"ERROR_CODE_RELATION_REFERENCE_NOT_FOUND_IN_ENTITY_REFERENCES": 2010,
		"ERROR_CODE_RELATION_REFERENCE_MUST_HAVE_ONE_ENTITY_REFERENCE": 2011,
		"ERROR_CODE_DUPLICATED_ENTITY_REFERENCE":                       2012,
		"ERROR_CODE_DUPLICATED_RELATION_REFERENCE":                     2013,
		"ERROR_CODE_DUPLICATED_PERMISSION_REFERENCE":                   2014,
		"ERROR_CODE_SCHEMA_PARSE":                                      2015,
		"ERROR_CODE_SCHEMA_COMPILE":                                    2016,
		"ERROR_CODE_SUBJECT_RELATION_MUST_BE_EMPTY":                    2017,
		"ERROR_CODE_SUBJECT_RELATION_CANNOT_BE_EMPTY":                  2018,
		"ERROR_CODE_SCHEMA_MUST_HAVE_USER_ENTITY_DEFINITION":           2019,
		"ERROR_CODE_UNIQUE_CONSTRAINT":                                 2020,
		"ERROR_CODE_INVALID_CONTINUOUS_TOKEN":                          2021,
		"ERROR_CODE_INVALID_KEY":                                       2022,
		"ERROR_CODE_ENTITY_TYPE_REQUIRED":                              2023,
		"ERROR_CODE_NO_ENTITY_REFERENCES_FOUND_IN_SCHEMA":              2024,
		"ERROR_CODE_INVALID_ARGUMENT":                                  2025,
		"ERROR_CODE_INVALID_RULE_REFERENCE":                            2026,
		"ERROR_CODE_NOT_SUPPORTED_WALK":                                2027,
		"ERROR_CODE_MISSING_ARGUMENT":                                  2028,
		"ERROR_CODE_ALREADY_EXIST":                                     2029,
		"ERROR_CODE_NOT_FOUND":                                         4000,
		"ERROR_CODE_ENTITY_TYPE_NOT_FOUND":                             4001,
		"ERROR_CODE_PERMISSION_NOT_FOUND":                              4002,
		"ERROR_CODE_SCHEMA_NOT_FOUND":                                  4003,
		"ERROR_CODE_SUBJECT_TYPE_NOT_FOUND":                            4004,
		"ERROR_CODE_ENTITY_DEFINITION_NOT_FOUND":                       4005,
		"ERROR_CODE_PERMISSION_DEFINITION_NOT_FOUND":                   4006,
		"ERROR_CODE_RELATION_DEFINITION_NOT_FOUND":                     4007,
		"ERROR_CODE_RECORD_NOT_FOUND":                                  4008,
		"ERROR_CODE_TENANT_NOT_FOUND":                                  4009,
		"ERROR_CODE_ATTRIBUTE_DEFINITION_NOT_FOUND":                    4010,
		"ERROR_CODE_ATTRIBUTE_TYPE_MISMATCH":                           4011,
		"ERROR_CODE_INTERNAL":                                          5000,
		"ERROR_CODE_CANCELLED":                                         5001,
		"ERROR_CODE_SQL_BUILDER":                                       5002,
		"ERROR_CODE_CIRCUIT_BREAKER":                                   5003,
		"ERROR_CODE_EXECUTION":                                         5005,
		"ERROR_CODE_SCAN":                                              5006,
		"ERROR_CODE_MIGRATION":                                         5007,
		"ERROR_CODE_TYPE_CONVERSATION":                                 5008,
		"ERROR_CODE_ERROR_MAX_RETRIES":                                 5009,
		"ERROR_CODE_ROLLBACK":                                          5010,
		"ERROR_CODE_EXCLUSION_REQUIRES_MORE_THAN_ONE_FUNCTION":         5011,
		"ERROR_CODE_NOT_IMPLEMENTED":                                   5012,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_base_v1_errors_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_base_v1_errors_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_base_v1_errors_proto_rawDescGZIP(), []int{0}
}

// ErrorResponse
type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    ErrorCode `protobuf:"varint,1,opt,name=code,proto3,enum=base.v1.ErrorCode" json:"code,omitempty"`
	Message string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_v1_errors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_base_v1_errors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_base_v1_errors_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorResponse) GetCode() ErrorCode {
	if x != nil {
		return x.Code
	}
	return ErrorCode_ERROR_CODE_UNSPECIFIED
}

func (x *ErrorResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_base_v1_errors_proto protoreflect.FileDescriptor

var file_base_v1_errors_proto_rawDesc = []byte{
	0x0a, 0x14, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x22,
	0x51, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2a, 0xe2, 0x10, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1a, 0x0a, 0x16, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x1f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49,
	0x4e, 0x47, 0x5f, 0x42, 0x45, 0x41, 0x52, 0x45, 0x52, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10,
	0xe9, 0x07, 0x12, 0x1f, 0x0a, 0x1a, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x45, 0x44,
	0x10, 0xea, 0x07, 0x12, 0x21, 0x0a, 0x1c, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54,
	0x5f, 0x49, 0x44, 0x10, 0xeb, 0x07, 0x12, 0x1a, 0x0a, 0x15, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0xd0, 0x0f, 0x12, 0x24, 0x0a, 0x1f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x5f, 0x43, 0x48, 0x49, 0x4c, 0x44,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0xd2, 0x0f, 0x12, 0x24, 0x0a, 0x1f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44,
	0x5f, 0x43, 0x48, 0x49, 0x4c, 0x44, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x10, 0xd3, 0x0f, 0x12, 0x2c,
	0x0a, 0x27, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x44,
	0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x10, 0xd6, 0x0f, 0x12, 0x2b, 0x0a, 0x26,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53,
	0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x57, 0x41, 0x4c, 0x4b, 0x10, 0xd7, 0x0f, 0x12, 0x32, 0x0a, 0x2d, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x41,
	0x4e, 0x44, 0x5f, 0x53, 0x55, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f,
	0x54, 0x5f, 0x42, 0x45, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0xd8, 0x0f, 0x12, 0x20, 0x0a,
	0x1b, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x44, 0x45, 0x50, 0x54,
	0x48, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x4f, 0x55, 0x47, 0x48, 0x10, 0xd9, 0x0f, 0x12,
	0x41, 0x0a, 0x3c, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x52, 0x45,
	0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x49, 0x4e, 0x5f, 0x45, 0x4e,
	0x54, 0x49, 0x54, 0x59, 0x5f, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x53, 0x10,
	0xda, 0x0f, 0x12, 0x41, 0x0a, 0x3c, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45,
	0x4e, 0x43, 0x45, 0x5f, 0x4d, 0x55, 0x53, 0x54, 0x5f, 0x48, 0x41, 0x56, 0x45, 0x5f, 0x4f, 0x4e,
	0x45, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e,
	0x43, 0x45, 0x10, 0xdb, 0x0f, 0x12, 0x2b, 0x0a, 0x26, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43,
	0x4f, 0x44, 0x45, 0x5f, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x45,
	0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x10,
	0xdc, 0x0f, 0x12, 0x2d, 0x0a, 0x28, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x4c, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x10, 0xdd,
	0x0f, 0x12, 0x2f, 0x0a, 0x2a, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49,
	0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x10,
	0xde, 0x0f, 0x12, 0x1c, 0x0a, 0x17, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x50, 0x41, 0x52, 0x53, 0x45, 0x10, 0xdf, 0x0f,
	0x12, 0x1e, 0x0a, 0x19, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53,
	0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x49, 0x4c, 0x45, 0x10, 0xe0, 0x0f,
	0x12, 0x2e, 0x0a, 0x29, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53,
	0x55, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x4d, 0x55, 0x53, 0x54, 0x5f, 0x42, 0x45, 0x5f, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x10, 0xe1, 0x0f,
	0x12, 0x30, 0x0a, 0x2b, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53,
	0x55, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x42, 0x45, 0x5f, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x10,
	0xe2, 0x0f, 0x12, 0x37, 0x0a, 0x32, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x4d, 0x55, 0x53, 0x54, 0x5f, 0x48, 0x41, 0x56,
	0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x44, 0x45,
	0x46, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0xe3, 0x0f, 0x12, 0x21, 0x0a, 0x1c, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x49, 0x51, 0x55, 0x45,
	0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x54, 0x10, 0xe4, 0x0f, 0x12, 0x28,
	0x0a, 0x23, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x49, 0x4e, 0x55, 0x4f, 0x55, 0x53, 0x5f,
	0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0xe5, 0x0f, 0x12, 0x1b, 0x0a, 0x16, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4b,
	0x45, 0x59, 0x10, 0xe6, 0x0f, 0x12, 0x24, 0x0a, 0x1f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43,
	0x4f, 0x44, 0x45, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0xe7, 0x0f, 0x12, 0x34, 0x0a, 0x2f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x5f, 0x45, 0x4e, 0x54,
	0x49, 0x54, 0x59, 0x5f, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x53, 0x5f, 0x46,
	0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x49, 0x4e, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x10, 0xe8,
	0x0f, 0x12, 0x20, 0x0a, 0x1b, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54,
	0x10, 0xe9, 0x0f, 0x12, 0x26, 0x0a, 0x21, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x52, 0x55, 0x4c, 0x45, 0x5f, 0x52,
	0x45, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x10, 0xea, 0x0f, 0x12, 0x22, 0x0a, 0x1d, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55,
	0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x57, 0x41, 0x4c, 0x4b, 0x10, 0xeb, 0x0f, 0x12,
	0x20, 0x0a, 0x1b, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4d, 0x49,
	0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0xec,
	0x0f, 0x12, 0x1d, 0x0a, 0x18, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0xed, 0x0f,
	0x12, 0x19, 0x0a, 0x14, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xa0, 0x1f, 0x12, 0x25, 0x0a, 0x20, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10,
	0xa1, 0x1f, 0x12, 0x24, 0x0a, 0x1f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xa2, 0x1f, 0x12, 0x20, 0x0a, 0x1b, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xa3, 0x1f, 0x12, 0x26, 0x0a, 0x21, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x55, 0x42, 0x4a, 0x45, 0x43, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10,
	0xa4, 0x1f, 0x12, 0x2b, 0x0a, 0x26, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x49, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xa5, 0x1f, 0x12,
	0x2f, 0x0a, 0x2a, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x50, 0x45,
	0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x49, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xa6, 0x1f,
	0x12, 0x2d, 0x0a, 0x28, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x52,
	0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x49, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xa7, 0x1f, 0x12,
	0x20, 0x0a, 0x1b, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x52, 0x45,
	0x43, 0x4f, 0x52, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xa8,
	0x1f, 0x12, 0x20, 0x0a, 0x1b, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44,
	0x10, 0xa9, 0x1f, 0x12, 0x2e, 0x0a, 0x29, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x5f, 0x44, 0x45, 0x46, 0x49,
	0x4e, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44,
	0x10, 0xaa, 0x1f, 0x12, 0x27, 0x0a, 0x22, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4d, 0x49, 0x53, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x10, 0xab, 0x1f, 0x12, 0x18, 0x0a, 0x13,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52,
	0x4e, 0x41, 0x4c, 0x10, 0x88, 0x27, 0x12, 0x19, 0x0a, 0x14, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x89,
	0x27, 0x12, 0x1b, 0x0a, 0x16, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x53, 0x51, 0x4c, 0x5f, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x45, 0x52, 0x10, 0x8a, 0x27, 0x12, 0x1f,
	0x0a, 0x1a, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x49, 0x52,
	0x43, 0x55, 0x49, 0x54, 0x5f, 0x42, 0x52, 0x45, 0x41, 0x4b, 0x45, 0x52, 0x10, 0x8b, 0x27, 0x12,
	0x19, 0x0a, 0x14, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x58,
	0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x8d, 0x27, 0x12, 0x14, 0x0a, 0x0f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x43, 0x41, 0x4e, 0x10, 0x8e, 0x27,
	0x12, 0x19, 0x0a, 0x14, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4d,
	0x49, 0x47, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x8f, 0x27, 0x12, 0x21, 0x0a, 0x1c, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43,
	0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x90, 0x27, 0x12, 0x21,
	0x0a, 0x1c, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x5f, 0x4d, 0x41, 0x58, 0x5f, 0x52, 0x45, 0x54, 0x52, 0x49, 0x45, 0x53, 0x10, 0x91,
	0x27, 0x12, 0x18, 0x0a, 0x13, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x52, 0x4f, 0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x92, 0x27, 0x12, 0x39, 0x0a, 0x34, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x58, 0x43, 0x4c, 0x55, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x53, 0x5f, 0x4d, 0x4f, 0x52,
	0x45, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x46, 0x55, 0x4e, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x93, 0x27, 0x12, 0x1f, 0x0a, 0x1a, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x4d, 0x45,
	0x4e, 0x54, 0x45, 0x44, 0x10, 0x94, 0x27, 0x42, 0x89, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x66, 0x79, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x66, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x62, 0x61, 0x73, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02,
	0x07, 0x42, 0x61, 0x73, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x42, 0x61, 0x73, 0x65, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x13, 0x42, 0x61, 0x73, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x42, 0x61, 0x73, 0x65, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_base_v1_errors_proto_rawDescOnce sync.Once
	file_base_v1_errors_proto_rawDescData = file_base_v1_errors_proto_rawDesc
)

func file_base_v1_errors_proto_rawDescGZIP() []byte {
	file_base_v1_errors_proto_rawDescOnce.Do(func() {
		file_base_v1_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_v1_errors_proto_rawDescData)
	})
	return file_base_v1_errors_proto_rawDescData
}

var file_base_v1_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_base_v1_errors_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_base_v1_errors_proto_goTypes = []interface{}{
	(ErrorCode)(0),        // 0: base.v1.ErrorCode
	(*ErrorResponse)(nil), // 1: base.v1.ErrorResponse
}
var file_base_v1_errors_proto_depIdxs = []int32{
	0, // 0: base.v1.ErrorResponse.code:type_name -> base.v1.ErrorCode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_base_v1_errors_proto_init() }
func file_base_v1_errors_proto_init() {
	if File_base_v1_errors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_v1_errors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_base_v1_errors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_v1_errors_proto_goTypes,
		DependencyIndexes: file_base_v1_errors_proto_depIdxs,
		EnumInfos:         file_base_v1_errors_proto_enumTypes,
		MessageInfos:      file_base_v1_errors_proto_msgTypes,
	}.Build()
	File_base_v1_errors_proto = out.File
	file_base_v1_errors_proto_rawDesc = nil
	file_base_v1_errors_proto_goTypes = nil
	file_base_v1_errors_proto_depIdxs = nil
}
