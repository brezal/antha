// Code generated by protoc-gen-go.
// source: github.com/antha-lang/antha/api/v1/blob.proto
// DO NOT EDIT!

package org_antha_lang_antha_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Unstructured data
type Blob struct {
	// A descriptive name.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Types that are valid to be assigned to From:
	//	*Blob_Bytes
	//	*Blob_HostFile
	From isBlob_From `protobuf_oneof:"from"`
}

func (m *Blob) Reset()                    { *m = Blob{} }
func (m *Blob) String() string            { return proto.CompactTextString(m) }
func (*Blob) ProtoMessage()               {}
func (*Blob) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

type isBlob_From interface {
	isBlob_From()
}

type Blob_Bytes struct {
	Bytes *FromBytes `protobuf:"bytes,2,opt,name=bytes,oneof"`
}
type Blob_HostFile struct {
	HostFile *FromHostFile `protobuf:"bytes,3,opt,name=host_file,json=hostFile,oneof"`
}

func (*Blob_Bytes) isBlob_From()    {}
func (*Blob_HostFile) isBlob_From() {}

func (m *Blob) GetFrom() isBlob_From {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Blob) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Blob) GetBytes() *FromBytes {
	if x, ok := m.GetFrom().(*Blob_Bytes); ok {
		return x.Bytes
	}
	return nil
}

func (m *Blob) GetHostFile() *FromHostFile {
	if x, ok := m.GetFrom().(*Blob_HostFile); ok {
		return x.HostFile
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Blob) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Blob_OneofMarshaler, _Blob_OneofUnmarshaler, _Blob_OneofSizer, []interface{}{
		(*Blob_Bytes)(nil),
		(*Blob_HostFile)(nil),
	}
}

func _Blob_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Blob)
	// from
	switch x := m.From.(type) {
	case *Blob_Bytes:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Bytes); err != nil {
			return err
		}
	case *Blob_HostFile:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HostFile); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Blob.From has unexpected type %T", x)
	}
	return nil
}

func _Blob_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Blob)
	switch tag {
	case 2: // from.bytes
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FromBytes)
		err := b.DecodeMessage(msg)
		m.From = &Blob_Bytes{msg}
		return true, err
	case 3: // from.host_file
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FromHostFile)
		err := b.DecodeMessage(msg)
		m.From = &Blob_HostFile{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Blob_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Blob)
	// from
	switch x := m.From.(type) {
	case *Blob_Bytes:
		s := proto.Size(x.Bytes)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Blob_HostFile:
		s := proto.Size(x.HostFile)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type FromBytes struct {
	Bytes []byte `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (m *FromBytes) Reset()                    { *m = FromBytes{} }
func (m *FromBytes) String() string            { return proto.CompactTextString(m) }
func (*FromBytes) ProtoMessage()               {}
func (*FromBytes) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *FromBytes) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

type FromHostFile struct {
	Filename string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
}

func (m *FromHostFile) Reset()                    { *m = FromHostFile{} }
func (m *FromHostFile) String() string            { return proto.CompactTextString(m) }
func (*FromHostFile) ProtoMessage()               {}
func (*FromHostFile) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *FromHostFile) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func init() {
	proto.RegisterType((*Blob)(nil), "org.antha_lang.antha.v1.Blob")
	proto.RegisterType((*FromBytes)(nil), "org.antha_lang.antha.v1.FromBytes")
	proto.RegisterType((*FromHostFile)(nil), "org.antha_lang.antha.v1.FromHostFile")
}

func init() { proto.RegisterFile("github.com/antha-lang/antha/api/v1/blob.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcc, 0x2b, 0xc9, 0x48, 0xd4, 0xcd, 0x49, 0xcc,
	0x4b, 0x87, 0x30, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0xcb, 0x0c, 0xf5, 0x93, 0x72, 0xf2, 0x93, 0xf4,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xc4, 0xf3, 0x8b, 0xd2, 0xf5, 0xc0, 0x92, 0xf1, 0x20, 0x75,
	0x10, 0xa6, 0x5e, 0x99, 0xa1, 0xd2, 0x12, 0x46, 0x2e, 0x16, 0xa7, 0x9c, 0xfc, 0x24, 0x21, 0x21,
	0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0xc8,
	0x8a, 0x8b, 0x35, 0xa9, 0xb2, 0x24, 0xb5, 0x58, 0x82, 0x49, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x49,
	0x0f, 0x87, 0x29, 0x7a, 0x6e, 0x45, 0xf9, 0xb9, 0x4e, 0x20, 0x95, 0x1e, 0x0c, 0x41, 0x10, 0x2d,
	0x42, 0x2e, 0x5c, 0x9c, 0x19, 0xf9, 0xc5, 0x25, 0xf1, 0x69, 0x99, 0x39, 0xa9, 0x12, 0xcc, 0x60,
	0xfd, 0xaa, 0x78, 0xf5, 0x7b, 0xe4, 0x17, 0x97, 0xb8, 0x65, 0xe6, 0xa4, 0x7a, 0x30, 0x04, 0x71,
	0x64, 0x40, 0xd9, 0x4e, 0x6c, 0x5c, 0x2c, 0x69, 0x45, 0xf9, 0xb9, 0x4a, 0x8a, 0x5c, 0x9c, 0x70,
	0x3b, 0x84, 0x44, 0x60, 0xce, 0x02, 0xb9, 0x95, 0x07, 0x6a, 0xa1, 0x92, 0x16, 0x17, 0x0f, 0xb2,
	0x31, 0x42, 0x52, 0x5c, 0x1c, 0x20, 0xbb, 0x91, 0x3c, 0x05, 0xe7, 0x27, 0xb1, 0x81, 0x43, 0xc5,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x77, 0x6e, 0x5f, 0x33, 0x46, 0x01, 0x00, 0x00,
}
