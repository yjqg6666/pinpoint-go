// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: Service.proto

package grpc

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_Service_proto protoreflect.FileDescriptor

var file_Service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0a, 0x53, 0x70, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x53, 0x74,
	0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x44, 0x75, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x43, 0x6d, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x40, 0x0a, 0x04, 0x53, 0x70, 0x61, 0x6e, 0x12, 0x38, 0x0a,
	0x08, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x70, 0x61, 0x6e, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x53, 0x70, 0x61, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x32, 0x65, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x12, 0x31, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x09, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x50, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x32, 0xb5,
	0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x12, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x71, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x53, 0x71, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x44,
	0x61, 0x74, 0x61, 0x1a, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x00, 0x12, 0x35, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x70, 0x69,
	0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x41,
	0x70, 0x69, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0b, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x15, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d,
	0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x32, 0x45, 0x0a, 0x04, 0x53, 0x74, 0x61, 0x74, 0x12, 0x3d,
	0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x12,
	0x10, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x53, 0x74, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x32, 0xd8, 0x03,
	0x0a, 0x16, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x43, 0x6d, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0f, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x43, 0x6d, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x03, 0x88, 0x02, 0x01,
	0x28, 0x01, 0x30, 0x01, 0x12, 0x39, 0x0a, 0x0f, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x56, 0x32, 0x12, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x43, 0x6d,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x43,
	0x6d, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12,
	0x3d, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x14,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x43, 0x6d, 0x64, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x5a,
	0x0a, 0x1e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x43, 0x6d, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12, 0x50, 0x0a, 0x17, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x43, 0x6d, 0x64, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x44, 0x75, 0x6d, 0x70, 0x52,
	0x65, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x1c,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x68, 0x72,
	0x65, 0x61, 0x64, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x20, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x43, 0x6d, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x44, 0x75, 0x6d, 0x70, 0x52, 0x65, 0x73, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x58, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e,
	0x6e, 0x61, 0x76, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x70, 0x69, 0x6e, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x42, 0x0c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x23, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x6e, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x67, 0x6f, 0x70, 0x69, 0x6e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_Service_proto_goTypes = []interface{}{
	(*PSpanMessage)(nil),                 // 0: v1.PSpanMessage
	(*PAgentInfo)(nil),                   // 1: v1.PAgentInfo
	(*PPing)(nil),                        // 2: v1.PPing
	(*PSqlMetaData)(nil),                 // 3: v1.PSqlMetaData
	(*PApiMetaData)(nil),                 // 4: v1.PApiMetaData
	(*PStringMetaData)(nil),              // 5: v1.PStringMetaData
	(*PStatMessage)(nil),                 // 6: v1.PStatMessage
	(*PCmdMessage)(nil),                  // 7: v1.PCmdMessage
	(*PCmdEchoResponse)(nil),             // 8: v1.PCmdEchoResponse
	(*PCmdActiveThreadCountRes)(nil),     // 9: v1.PCmdActiveThreadCountRes
	(*PCmdActiveThreadDumpRes)(nil),      // 10: v1.PCmdActiveThreadDumpRes
	(*PCmdActiveThreadLightDumpRes)(nil), // 11: v1.PCmdActiveThreadLightDumpRes
	(*empty.Empty)(nil),                  // 12: google.protobuf.Empty
	(*PResult)(nil),                      // 13: v1.PResult
	(*PCmdRequest)(nil),                  // 14: v1.PCmdRequest
}
var file_Service_proto_depIdxs = []int32{
	0,  // 0: v1.Span.SendSpan:input_type -> v1.PSpanMessage
	1,  // 1: v1.Agent.RequestAgentInfo:input_type -> v1.PAgentInfo
	2,  // 2: v1.Agent.PingSession:input_type -> v1.PPing
	3,  // 3: v1.Metadata.RequestSqlMetaData:input_type -> v1.PSqlMetaData
	4,  // 4: v1.Metadata.RequestApiMetaData:input_type -> v1.PApiMetaData
	5,  // 5: v1.Metadata.RequestStringMetaData:input_type -> v1.PStringMetaData
	6,  // 6: v1.Stat.SendAgentStat:input_type -> v1.PStatMessage
	7,  // 7: v1.ProfilerCommandService.HandleCommand:input_type -> v1.PCmdMessage
	7,  // 8: v1.ProfilerCommandService.HandleCommandV2:input_type -> v1.PCmdMessage
	8,  // 9: v1.ProfilerCommandService.CommandEcho:input_type -> v1.PCmdEchoResponse
	9,  // 10: v1.ProfilerCommandService.CommandStreamActiveThreadCount:input_type -> v1.PCmdActiveThreadCountRes
	10, // 11: v1.ProfilerCommandService.CommandActiveThreadDump:input_type -> v1.PCmdActiveThreadDumpRes
	11, // 12: v1.ProfilerCommandService.CommandActiveThreadLightDump:input_type -> v1.PCmdActiveThreadLightDumpRes
	12, // 13: v1.Span.SendSpan:output_type -> google.protobuf.Empty
	13, // 14: v1.Agent.RequestAgentInfo:output_type -> v1.PResult
	2,  // 15: v1.Agent.PingSession:output_type -> v1.PPing
	13, // 16: v1.Metadata.RequestSqlMetaData:output_type -> v1.PResult
	13, // 17: v1.Metadata.RequestApiMetaData:output_type -> v1.PResult
	13, // 18: v1.Metadata.RequestStringMetaData:output_type -> v1.PResult
	12, // 19: v1.Stat.SendAgentStat:output_type -> google.protobuf.Empty
	14, // 20: v1.ProfilerCommandService.HandleCommand:output_type -> v1.PCmdRequest
	14, // 21: v1.ProfilerCommandService.HandleCommandV2:output_type -> v1.PCmdRequest
	12, // 22: v1.ProfilerCommandService.CommandEcho:output_type -> google.protobuf.Empty
	12, // 23: v1.ProfilerCommandService.CommandStreamActiveThreadCount:output_type -> google.protobuf.Empty
	12, // 24: v1.ProfilerCommandService.CommandActiveThreadDump:output_type -> google.protobuf.Empty
	12, // 25: v1.ProfilerCommandService.CommandActiveThreadLightDump:output_type -> google.protobuf.Empty
	13, // [13:26] is the sub-list for method output_type
	0,  // [0:13] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_Service_proto_init() }
func file_Service_proto_init() {
	if File_Service_proto != nil {
		return
	}
	file_Span_proto_init()
	file_Stat_proto_init()
	file_ThreadDump_proto_init()
	file_Cmd_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   5,
		},
		GoTypes:           file_Service_proto_goTypes,
		DependencyIndexes: file_Service_proto_depIdxs,
	}.Build()
	File_Service_proto = out.File
	file_Service_proto_rawDesc = nil
	file_Service_proto_goTypes = nil
	file_Service_proto_depIdxs = nil
}
