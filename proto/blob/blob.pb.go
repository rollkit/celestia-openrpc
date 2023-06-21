// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: blob/blob.proto

package blob

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

// Proof is a collection of nmt.Proofs that verifies the inclusion of the data.
type Proof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start uint32   `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End   uint32   `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	Nodes [][]byte `protobuf:"bytes,3,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *Proof) Reset() {
	*x = Proof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blob_blob_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proof) ProtoMessage() {}

func (x *Proof) ProtoReflect() protoreflect.Message {
	mi := &file_blob_blob_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proof.ProtoReflect.Descriptor instead.
func (*Proof) Descriptor() ([]byte, []int) {
	return file_blob_blob_proto_rawDescGZIP(), []int{0}
}

func (x *Proof) GetStart() uint32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Proof) GetEnd() uint32 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *Proof) GetNodes() [][]byte {
	if x != nil {
		return x.Nodes
	}
	return nil
}

// Blob (named after binary large object) is a chunk of data submitted by a user
// to be published to the Celestia blockchain. The data of a Blob is published
// to a namespace and is encoded into shares based on the format specified by
// share_version.
type Blob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace        []byte `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Data             []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	ShareVersion     uint32 `protobuf:"varint,3,opt,name=share_version,json=shareVersion,proto3" json:"share_version,omitempty"`
	NamespaceVersion uint32 `protobuf:"varint,4,opt,name=namespace_version,json=namespaceVersion,proto3" json:"namespace_version,omitempty"`
	Commitment       []byte `protobuf:"bytes,5,opt,name=commitment,proto3" json:"commitment,omitempty"`
}

func (x *Blob) Reset() {
	*x = Blob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blob_blob_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blob) ProtoMessage() {}

func (x *Blob) ProtoReflect() protoreflect.Message {
	mi := &file_blob_blob_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blob.ProtoReflect.Descriptor instead.
func (*Blob) Descriptor() ([]byte, []int) {
	return file_blob_blob_proto_rawDescGZIP(), []int{1}
}

func (x *Blob) GetNamespace() []byte {
	if x != nil {
		return x.Namespace
	}
	return nil
}

func (x *Blob) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Blob) GetShareVersion() uint32 {
	if x != nil {
		return x.ShareVersion
	}
	return 0
}

func (x *Blob) GetNamespaceVersion() uint32 {
	if x != nil {
		return x.NamespaceVersion
	}
	return 0
}

func (x *Blob) GetCommitment() []byte {
	if x != nil {
		return x.Commitment
	}
	return nil
}

var File_blob_blob_proto protoreflect.FileDescriptor

var file_blob_blob_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x62, 0x6c, 0x6f, 0x62, 0x22, 0x45, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x6f, 0x66,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0xaa,
	0x01, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2b,
	0x0a, 0x11, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x30, 0x5a, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x6c, 0x6c, 0x6b, 0x69,
	0x74, 0x2f, 0x63, 0x65, 0x6c, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blob_blob_proto_rawDescOnce sync.Once
	file_blob_blob_proto_rawDescData = file_blob_blob_proto_rawDesc
)

func file_blob_blob_proto_rawDescGZIP() []byte {
	file_blob_blob_proto_rawDescOnce.Do(func() {
		file_blob_blob_proto_rawDescData = protoimpl.X.CompressGZIP(file_blob_blob_proto_rawDescData)
	})
	return file_blob_blob_proto_rawDescData
}

var file_blob_blob_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_blob_blob_proto_goTypes = []interface{}{
	(*Proof)(nil), // 0: blob.Proof
	(*Blob)(nil),  // 1: blob.Blob
}
var file_blob_blob_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_blob_blob_proto_init() }
func file_blob_blob_proto_init() {
	if File_blob_blob_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blob_blob_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proof); i {
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
		file_blob_blob_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blob); i {
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
			RawDescriptor: file_blob_blob_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blob_blob_proto_goTypes,
		DependencyIndexes: file_blob_blob_proto_depIdxs,
		MessageInfos:      file_blob_blob_proto_msgTypes,
	}.Build()
	File_blob_blob_proto = out.File
	file_blob_blob_proto_rawDesc = nil
	file_blob_blob_proto_goTypes = nil
	file_blob_blob_proto_depIdxs = nil
}