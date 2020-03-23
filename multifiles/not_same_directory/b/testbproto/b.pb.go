// Code generated by protoc-gen-go. DO NOT EDIT.
// source: b.proto

package testbproto

import (
	aproto "./aproto"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Data struct {
	List                 []*aproto.Info `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_1848ba9bd6296c7e, []int{0}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetList() []*aproto.Info {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*Data)(nil), "testb.Data")
}

func init() { proto.RegisterFile("b.proto", fileDescriptor_1848ba9bd6296c7e) }

var fileDescriptor_1848ba9bd6296c7e = []byte{
	// 95 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x62, 0x4f, 0xd2, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0x49, 0x2d, 0x2e, 0x49, 0x92, 0x62, 0x4f, 0x84, 0xf0, 0x95,
	0xd4, 0xb9, 0x58, 0x5c, 0x12, 0x4b, 0x12, 0x85, 0xe4, 0xb9, 0x58, 0x72, 0x32, 0x8b, 0x4b, 0x24,
	0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0xb8, 0xf5, 0x40, 0xca, 0x12, 0xf5, 0x3c, 0xf3, 0xd2, 0xf2,
	0x83, 0xc0, 0x12, 0x4e, 0x7c, 0x51, 0x3c, 0x7a, 0xfa, 0x60, 0xcd, 0x60, 0x8d, 0x49, 0x6c, 0x60,
	0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x67, 0x64, 0xcb, 0x5a, 0x00, 0x00, 0x00,
}