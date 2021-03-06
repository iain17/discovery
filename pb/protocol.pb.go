// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/protocol.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/protocol.proto

It has these top-level messages:
	Hearbeat
	DPeerInfo
	Transfer
	DPeer
	Peers
	Shutdown
	Message
*/
package pb

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

type Hearbeat struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Hearbeat) Reset()                    { *m = Hearbeat{} }
func (m *Hearbeat) String() string            { return proto.CompactTextString(m) }
func (*Hearbeat) ProtoMessage()               {}
func (*Hearbeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Hearbeat) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DPeerInfo struct {
	Network string            `protobuf:"bytes,1,opt,name=network" json:"network,omitempty"`
	Id      string            `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Info    map[string]string `protobuf:"bytes,3,rep,name=info" json:"info,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *DPeerInfo) Reset()                    { *m = DPeerInfo{} }
func (m *DPeerInfo) String() string            { return proto.CompactTextString(m) }
func (*DPeerInfo) ProtoMessage()               {}
func (*DPeerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DPeerInfo) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *DPeerInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DPeerInfo) GetInfo() map[string]string {
	if m != nil {
		return m.Info
	}
	return nil
}

type Transfer struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Transfer) Reset()                    { *m = Transfer{} }
func (m *Transfer) String() string            { return proto.CompactTextString(m) }
func (*Transfer) ProtoMessage()               {}
func (*Transfer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Transfer) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type DPeer struct {
	Ip   string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
}

func (m *DPeer) Reset()                    { *m = DPeer{} }
func (m *DPeer) String() string            { return proto.CompactTextString(m) }
func (*DPeer) ProtoMessage()               {}
func (*DPeer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DPeer) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *DPeer) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type Peers struct {
	Peers []*DPeer `protobuf:"bytes,1,rep,name=peers" json:"peers,omitempty"`
}

func (m *Peers) Reset()                    { *m = Peers{} }
func (m *Peers) String() string            { return proto.CompactTextString(m) }
func (*Peers) ProtoMessage()               {}
func (*Peers) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Peers) GetPeers() []*DPeer {
	if m != nil {
		return m.Peers
	}
	return nil
}

type Shutdown struct {
}

func (m *Shutdown) Reset()                    { *m = Shutdown{} }
func (m *Shutdown) String() string            { return proto.CompactTextString(m) }
func (*Shutdown) ProtoMessage()               {}
func (*Shutdown) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Message struct {
	Version int64 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	// Types that are valid to be assigned to Msg:
	//	*Message_Heartbeat
	//	*Message_PeerInfo
	//	*Message_Transfer
	//	*Message_Peers
	//	*Message_Shutdown
	Msg isMessage_Msg `protobuf_oneof:"msg"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type isMessage_Msg interface {
	isMessage_Msg()
}

type Message_Heartbeat struct {
	Heartbeat *Hearbeat `protobuf:"bytes,2,opt,name=heartbeat,oneof"`
}
type Message_PeerInfo struct {
	PeerInfo *DPeerInfo `protobuf:"bytes,3,opt,name=peerInfo,oneof"`
}
type Message_Transfer struct {
	Transfer *Transfer `protobuf:"bytes,4,opt,name=transfer,oneof"`
}
type Message_Peers struct {
	Peers *Peers `protobuf:"bytes,5,opt,name=peers,oneof"`
}
type Message_Shutdown struct {
	Shutdown *Shutdown `protobuf:"bytes,6,opt,name=shutdown,oneof"`
}

func (*Message_Heartbeat) isMessage_Msg() {}
func (*Message_PeerInfo) isMessage_Msg()  {}
func (*Message_Transfer) isMessage_Msg()  {}
func (*Message_Peers) isMessage_Msg()     {}
func (*Message_Shutdown) isMessage_Msg()  {}

func (m *Message) GetMsg() isMessage_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *Message) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Message) GetHeartbeat() *Hearbeat {
	if x, ok := m.GetMsg().(*Message_Heartbeat); ok {
		return x.Heartbeat
	}
	return nil
}

func (m *Message) GetPeerInfo() *DPeerInfo {
	if x, ok := m.GetMsg().(*Message_PeerInfo); ok {
		return x.PeerInfo
	}
	return nil
}

func (m *Message) GetTransfer() *Transfer {
	if x, ok := m.GetMsg().(*Message_Transfer); ok {
		return x.Transfer
	}
	return nil
}

func (m *Message) GetPeers() *Peers {
	if x, ok := m.GetMsg().(*Message_Peers); ok {
		return x.Peers
	}
	return nil
}

func (m *Message) GetShutdown() *Shutdown {
	if x, ok := m.GetMsg().(*Message_Shutdown); ok {
		return x.Shutdown
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Message) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Message_OneofMarshaler, _Message_OneofUnmarshaler, _Message_OneofSizer, []interface{}{
		(*Message_Heartbeat)(nil),
		(*Message_PeerInfo)(nil),
		(*Message_Transfer)(nil),
		(*Message_Peers)(nil),
		(*Message_Shutdown)(nil),
	}
}

func _Message_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Message)
	// msg
	switch x := m.Msg.(type) {
	case *Message_Heartbeat:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Heartbeat); err != nil {
			return err
		}
	case *Message_PeerInfo:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PeerInfo); err != nil {
			return err
		}
	case *Message_Transfer:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Transfer); err != nil {
			return err
		}
	case *Message_Peers:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Peers); err != nil {
			return err
		}
	case *Message_Shutdown:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Shutdown); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Message.Msg has unexpected type %T", x)
	}
	return nil
}

func _Message_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Message)
	switch tag {
	case 2: // msg.heartbeat
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Hearbeat)
		err := b.DecodeMessage(msg)
		m.Msg = &Message_Heartbeat{msg}
		return true, err
	case 3: // msg.peerInfo
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DPeerInfo)
		err := b.DecodeMessage(msg)
		m.Msg = &Message_PeerInfo{msg}
		return true, err
	case 4: // msg.transfer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Transfer)
		err := b.DecodeMessage(msg)
		m.Msg = &Message_Transfer{msg}
		return true, err
	case 5: // msg.peers
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Peers)
		err := b.DecodeMessage(msg)
		m.Msg = &Message_Peers{msg}
		return true, err
	case 6: // msg.shutdown
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Shutdown)
		err := b.DecodeMessage(msg)
		m.Msg = &Message_Shutdown{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Message_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Message)
	// msg
	switch x := m.Msg.(type) {
	case *Message_Heartbeat:
		s := proto.Size(x.Heartbeat)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_PeerInfo:
		s := proto.Size(x.PeerInfo)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Transfer:
		s := proto.Size(x.Transfer)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Peers:
		s := proto.Size(x.Peers)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Shutdown:
		s := proto.Size(x.Shutdown)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Hearbeat)(nil), "pb.Hearbeat")
	proto.RegisterType((*DPeerInfo)(nil), "pb.DPeerInfo")
	proto.RegisterType((*Transfer)(nil), "pb.Transfer")
	proto.RegisterType((*DPeer)(nil), "pb.DPeer")
	proto.RegisterType((*Peers)(nil), "pb.Peers")
	proto.RegisterType((*Shutdown)(nil), "pb.Shutdown")
	proto.RegisterType((*Message)(nil), "pb.Message")
}

func init() { proto.RegisterFile("pb/protocol.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x1b, 0xa7, 0x69, 0x93, 0x69, 0x41, 0x60, 0x21, 0x61, 0x71, 0x80, 0x12, 0x71, 0xa8,
	0x28, 0x0a, 0x52, 0x39, 0x80, 0x38, 0x56, 0x20, 0x85, 0x03, 0x12, 0x32, 0xfb, 0x02, 0xce, 0xc6,
	0x6d, 0xa3, 0xb6, 0xb1, 0xe5, 0xb8, 0xad, 0x7a, 0xdb, 0xf7, 0xd8, 0x97, 0x5d, 0x79, 0x6a, 0xa7,
	0xda, 0x4b, 0x34, 0x63, 0xff, 0xf9, 0x66, 0xfe, 0x3f, 0x81, 0xd7, 0xba, 0xfa, 0xaa, 0x8d, 0xb2,
	0xea, 0x5e, 0xed, 0x0b, 0x2c, 0x28, 0xd1, 0x55, 0xfe, 0x09, 0xd2, 0x52, 0x0a, 0x53, 0x49, 0x61,
	0x29, 0x83, 0xf1, 0x41, 0x76, 0x9d, 0xd8, 0x48, 0x16, 0xcd, 0xa2, 0x79, 0xc6, 0x43, 0x9b, 0x3f,
	0x46, 0x90, 0xfd, 0xfa, 0x27, 0xa5, 0xf9, 0xd3, 0xae, 0x95, 0xd3, 0xb5, 0xd2, 0x9e, 0x95, 0xd9,
	0x05, 0x9d, 0x6f, 0xe9, 0x4b, 0x20, 0x4d, 0xcd, 0x08, 0x1e, 0x92, 0xa6, 0xa6, 0x0b, 0x18, 0x36,
	0xed, 0x5a, 0xb1, 0x78, 0x16, 0xcf, 0x27, 0xcb, 0xb7, 0x85, 0xae, 0x8a, 0x1e, 0x53, 0xb8, 0xc7,
	0xef, 0xd6, 0x9a, 0x0b, 0x47, 0xd1, 0xbb, 0xef, 0x90, 0xf5, 0x47, 0xf4, 0x15, 0xc4, 0x3b, 0x79,
	0xf1, 0x7c, 0x57, 0xd2, 0x37, 0x90, 0x9c, 0xc4, 0xfe, 0x28, 0x3d, 0xfe, 0xda, 0xfc, 0x24, 0x3f,
	0xa2, 0xfc, 0x3d, 0xa4, 0x77, 0x46, 0xb4, 0xdd, 0x5a, 0x1a, 0x4a, 0x61, 0x58, 0x0b, 0x2b, 0xf0,
	0xc5, 0x29, 0xc7, 0x3a, 0x5f, 0x40, 0x82, 0x53, 0x71, 0x3d, 0xed, 0x99, 0xa4, 0xd1, 0x4e, 0xac,
	0x95, 0xb1, 0x48, 0x4c, 0x38, 0xd6, 0xf9, 0x1c, 0x12, 0xa7, 0xed, 0xe8, 0x07, 0x48, 0xb4, 0x2b,
	0x58, 0x84, 0xcb, 0x67, 0xfd, 0xf2, 0xfc, 0x7a, 0x9e, 0x03, 0xa4, 0xff, 0xb7, 0x47, 0x5b, 0xab,
	0x73, 0x9b, 0x3f, 0x10, 0x18, 0xff, 0xbd, 0x86, 0xe5, 0xe2, 0x39, 0x49, 0xd3, 0x35, 0xaa, 0xc5,
	0x51, 0x31, 0x0f, 0x2d, 0xfd, 0x02, 0xd9, 0x56, 0x0a, 0x63, 0x5d, 0xda, 0x38, 0x74, 0xb2, 0x9c,
	0x3a, 0x6c, 0xf8, 0x02, 0xe5, 0x80, 0xdf, 0x04, 0x74, 0x01, 0xa9, 0xf6, 0x59, 0xb1, 0x18, 0xc5,
	0x2f, 0x9e, 0x05, 0x58, 0x0e, 0x78, 0x2f, 0xa0, 0x9f, 0x21, 0xb5, 0x3e, 0x03, 0x36, 0xbc, 0x91,
	0x43, 0x2e, 0x4e, 0x1b, 0xee, 0xe9, 0xc7, 0xe0, 0x2c, 0x41, 0x21, 0x3a, 0x43, 0xcf, 0xe5, 0xc0,
	0x7b, 0x73, 0xb8, 0xce, 0x7b, 0x63, 0xa3, 0x1b, 0x2e, 0xf8, 0x75, 0xb8, 0x70, 0xbf, 0x4a, 0x20,
	0x3e, 0x74, 0x9b, 0x15, 0x29, 0xe3, 0x6a, 0x84, 0x3f, 0xd6, 0xb7, 0xa7, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xf8, 0x9c, 0xe4, 0x2b, 0x6d, 0x02, 0x00, 0x00,
}
