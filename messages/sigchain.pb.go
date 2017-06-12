// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sigchain.proto

/*
Package messages is a generated protocol buffer package.

It is generated from these files:
	sigchain.proto

It has these top-level messages:
	ServerInfo
	Signed
	NewDevice
	Link
*/
package messages

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ServerInfo struct {
	Version  string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	SenderId string `protobuf:"bytes,2,opt,name=sender_id,json=senderId" json:"sender_id,omitempty"`
}

func (m *ServerInfo) Reset()                    { *m = ServerInfo{} }
func (m *ServerInfo) String() string            { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()               {}
func (*ServerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServerInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServerInfo) GetSenderId() string {
	if m != nil {
		return m.SenderId
	}
	return ""
}

type Signed struct {
	Body      []byte `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	PublicKey []byte `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (m *Signed) Reset()                    { *m = Signed{} }
func (m *Signed) String() string            { return proto.CompactTextString(m) }
func (*Signed) ProtoMessage()               {}
func (*Signed) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Signed) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Signed) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Signed) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

type NewDevice struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	PublicKey []byte `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	FCMToken  string `protobuf:"bytes,3,opt,name=FCM_token,json=FCMToken" json:"FCM_token,omitempty"`
}

func (m *NewDevice) Reset()                    { *m = NewDevice{} }
func (m *NewDevice) String() string            { return proto.CompactTextString(m) }
func (*NewDevice) ProtoMessage()               {}
func (*NewDevice) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NewDevice) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NewDevice) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *NewDevice) GetFCMToken() string {
	if m != nil {
		return m.FCMToken
	}
	return ""
}

type Link struct {
	Prev           []byte `protobuf:"bytes,1,opt,name=prev,proto3" json:"prev,omitempty"`
	SequenceNumber uint32 `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"`
	// Types that are valid to be assigned to Body:
	//	*Link_NewDevice
	Body isLink_Body `protobuf_oneof:"body"`
}

func (m *Link) Reset()                    { *m = Link{} }
func (m *Link) String() string            { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()               {}
func (*Link) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isLink_Body interface {
	isLink_Body()
}

type Link_NewDevice struct {
	NewDevice *NewDevice `protobuf:"bytes,3,opt,name=new_device,json=newDevice,oneof"`
}

func (*Link_NewDevice) isLink_Body() {}

func (m *Link) GetBody() isLink_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Link) GetPrev() []byte {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *Link) GetSequenceNumber() uint32 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *Link) GetNewDevice() *NewDevice {
	if x, ok := m.GetBody().(*Link_NewDevice); ok {
		return x.NewDevice
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Link) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Link_OneofMarshaler, _Link_OneofUnmarshaler, _Link_OneofSizer, []interface{}{
		(*Link_NewDevice)(nil),
	}
}

func _Link_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Link)
	// body
	switch x := m.Body.(type) {
	case *Link_NewDevice:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NewDevice); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Link.Body has unexpected type %T", x)
	}
	return nil
}

func _Link_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Link)
	switch tag {
	case 3: // body.new_device
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NewDevice)
		err := b.DecodeMessage(msg)
		m.Body = &Link_NewDevice{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Link_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Link)
	// body
	switch x := m.Body.(type) {
	case *Link_NewDevice:
		s := proto.Size(x.NewDevice)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ServerInfo)(nil), "ServerInfo")
	proto.RegisterType((*Signed)(nil), "Signed")
	proto.RegisterType((*NewDevice)(nil), "NewDevice")
	proto.RegisterType((*Link)(nil), "Link")
}

func init() { proto.RegisterFile("sigchain.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x14, 0x84, 0x6d, 0x2d, 0xb5, 0xfb, 0xac, 0x15, 0xf6, 0x14, 0x50, 0x41, 0x72, 0x51, 0x10, 0x7a,
	0xd0, 0x7f, 0xd0, 0x48, 0xb1, 0x68, 0x7b, 0xd8, 0x7a, 0x51, 0x0f, 0x21, 0xc9, 0x3e, 0xe3, 0x12,
	0xf3, 0x36, 0xee, 0x26, 0xa9, 0xf9, 0xf7, 0x92, 0x8d, 0x51, 0xe8, 0x6d, 0x66, 0x1e, 0x33, 0xf9,
	0xc8, 0xc2, 0xcc, 0xaa, 0x34, 0xf9, 0x88, 0x14, 0xcd, 0x0b, 0xa3, 0x4b, 0xed, 0x07, 0x00, 0x5b,
	0x34, 0x35, 0x9a, 0x15, 0xbd, 0x6b, 0xee, 0xc1, 0x51, 0x8d, 0xc6, 0x2a, 0x4d, 0xde, 0xe0, 0x72,
	0x70, 0xcd, 0x44, 0x6f, 0xf9, 0x19, 0x30, 0x8b, 0x24, 0xd1, 0x84, 0x4a, 0x7a, 0x43, 0x77, 0x9b,
	0x74, 0xc1, 0x4a, 0xfa, 0x2f, 0x30, 0xde, 0xaa, 0x94, 0x50, 0x72, 0x0e, 0xa3, 0x58, 0xcb, 0xc6,
	0xb5, 0xa7, 0xc2, 0x69, 0x7e, 0x0e, 0xcc, 0xaa, 0x94, 0xa2, 0xb2, 0x32, 0xe8, 0xaa, 0x53, 0xf1,
	0x1f, 0xf0, 0x0b, 0x80, 0xa2, 0x8a, 0x3f, 0x55, 0x12, 0x66, 0xd8, 0x78, 0x87, 0xdd, 0xb9, 0x4b,
	0x1e, 0xb1, 0xf1, 0xdf, 0x80, 0x6d, 0x70, 0x77, 0x8f, 0xb5, 0x4a, 0xb0, 0x5d, 0xa7, 0x28, 0xc7,
	0x5f, 0x36, 0xa7, 0xf7, 0xfa, 0xc3, 0xbd, 0x7e, 0xcb, 0xbd, 0x0c, 0xd6, 0x61, 0xa9, 0x33, 0x24,
	0xb7, 0xce, 0xc4, 0x64, 0x19, 0xac, 0x9f, 0x5b, 0xef, 0x7f, 0xc3, 0xe8, 0x49, 0x51, 0xd6, 0xee,
	0x16, 0x06, 0xeb, 0x9e, 0xba, 0xd5, 0xfc, 0x0a, 0x4e, 0x2d, 0x7e, 0x55, 0x48, 0x09, 0x86, 0x54,
	0xe5, 0x31, 0x1a, 0x37, 0x7e, 0x22, 0x66, 0x7d, 0xbc, 0x71, 0x29, 0xbf, 0x01, 0x20, 0xdc, 0x85,
	0xd2, 0x21, 0xba, 0x4f, 0x1c, 0xdf, 0xc2, 0xfc, 0x0f, 0xfa, 0xe1, 0x40, 0x30, 0xea, 0xcd, 0x62,
	0xdc, 0xfd, 0x9f, 0x05, 0xbc, 0x4e, 0x72, 0xb4, 0x36, 0x4a, 0xd1, 0xc6, 0x63, 0xf7, 0x12, 0x77,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x09, 0xbd, 0xf6, 0xd6, 0x9b, 0x01, 0x00, 0x00,
}
