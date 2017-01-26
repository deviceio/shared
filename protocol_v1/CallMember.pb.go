// Code generated by protoc-gen-go.
// source: CallMember.proto
// DO NOT EDIT!

package protocol_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CallMember struct {
	Reference string            `protobuf:"bytes,1,opt,name=Reference" json:"Reference,omitempty"`
	Name      string            `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Params    map[string][]byte `protobuf:"bytes,3,rep,name=Params" json:"Params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *CallMember) Reset()                    { *m = CallMember{} }
func (m *CallMember) String() string            { return proto.CompactTextString(m) }
func (*CallMember) ProtoMessage()               {}
func (*CallMember) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *CallMember) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *CallMember) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CallMember) GetParams() map[string][]byte {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterType((*CallMember)(nil), "protocol.v1.CallMember")
}

func init() { proto.RegisterFile("CallMember.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x70, 0x4e, 0xcc, 0xc9,
	0xf1, 0x4d, 0xcd, 0x4d, 0x4a, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x06, 0x53,
	0xc9, 0xf9, 0x39, 0x7a, 0x65, 0x86, 0x4a, 0xdb, 0x18, 0xb9, 0xb8, 0x10, 0x2a, 0x84, 0x64, 0xb8,
	0x38, 0x83, 0x52, 0xd3, 0x52, 0x8b, 0x52, 0xf3, 0x92, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0x10, 0x02, 0x42, 0x42, 0x5c, 0x2c, 0x7e, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x60, 0x09, 0x30,
	0x5b, 0xc8, 0x9a, 0x8b, 0x2d, 0x20, 0xb1, 0x28, 0x31, 0xb7, 0x58, 0x82, 0x59, 0x81, 0x59, 0x83,
	0xdb, 0x48, 0x59, 0x0f, 0xc9, 0x78, 0x3d, 0x24, 0xcb, 0x21, 0xaa, 0x5c, 0xf3, 0x4a, 0x8a, 0x2a,
	0x83, 0xa0, 0x5a, 0xa4, 0x2c, 0xb9, 0xb8, 0x91, 0x84, 0x85, 0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b,
	0xa1, 0xf6, 0x82, 0x98, 0x42, 0x22, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0x10, 0x2b, 0x79, 0x82,
	0x20, 0x1c, 0x2b, 0x26, 0x0b, 0xc6, 0x24, 0x36, 0xb0, 0x35, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xb3, 0x8f, 0xae, 0x24, 0xe0, 0x00, 0x00, 0x00,
}
