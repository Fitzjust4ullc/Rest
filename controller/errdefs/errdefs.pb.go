// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: errdefs.proto

package errdefs

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/moby/buildkit/solver/pb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Build struct {
	Ref                  string   `protobuf:"bytes,1,opt,name=Ref,proto3" json:"Ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Build) Reset()         { *m = Build{} }
func (m *Build) String() string { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()    {}
func (*Build) Descriptor() ([]byte, []int) {
	return fileDescriptor_689dc58a5060aff5, []int{0}
}
func (m *Build) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Build.Unmarshal(m, b)
}
func (m *Build) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Build.Marshal(b, m, deterministic)
}
func (m *Build) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Build.Merge(m, src)
}
func (m *Build) XXX_Size() int {
	return xxx_messageInfo_Build.Size(m)
}
func (m *Build) XXX_DiscardUnknown() {
	xxx_messageInfo_Build.DiscardUnknown(m)
}

var xxx_messageInfo_Build proto.InternalMessageInfo

func (m *Build) GetRef() string {
	if m != nil {
		return m.Ref
	}
	return ""
}

func init() {
	proto.RegisterType((*Build)(nil), "errdefs.Build")
}

func init() { proto.RegisterFile("errdefs.proto", fileDescriptor_689dc58a5060aff5) }

var fileDescriptor_689dc58a5060aff5 = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2d, 0x2a, 0x4a,
	0x49, 0x4d, 0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0xa5, 0x74, 0xd2,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x73, 0xf3, 0x93, 0x2a, 0xf5, 0x93,
	0x4a, 0x33, 0x73, 0x52, 0xb2, 0x33, 0x4b, 0xf4, 0x8b, 0xf3, 0x73, 0xca, 0x52, 0x8b, 0xf4, 0x0b,
	0x92, 0xf4, 0xf3, 0x0b, 0xa0, 0xda, 0x94, 0x24, 0xb9, 0x58, 0x9d, 0x40, 0xf2, 0x42, 0x02, 0x5c,
	0xcc, 0x41, 0xa9, 0x69, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0x66, 0x12, 0x1b, 0x58,
	0x85, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x56, 0x52, 0x41, 0x91, 0x69, 0x00, 0x00, 0x00,
}
