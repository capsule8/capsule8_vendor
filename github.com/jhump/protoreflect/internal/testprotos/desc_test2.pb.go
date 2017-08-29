// Code generated by protoc-gen-go. DO NOT EDIT.
// source: desc_test2.proto

package testprotos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import jhump_protoreflect_desc "github.com/jhump/protoreflect/internal/testprotos/pkg"
import desc_test_nopkg "github.com/jhump/protoreflect/internal/testprotos/nopkg"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Frobnitz struct {
	A *TestMessage        `protobuf:"bytes,1,opt,name=a" json:"a,omitempty"`
	B *AnotherTestMessage `protobuf:"bytes,2,opt,name=b" json:"b,omitempty"`
	// Types that are valid to be assigned to Abc:
	//	*Frobnitz_C1
	//	*Frobnitz_C2
	Abc isFrobnitz_Abc             `protobuf_oneof:"abc"`
	D   *TestMessage_NestedMessage `protobuf:"bytes,5,opt,name=d" json:"d,omitempty"`
	E   *TestMessage_NestedEnum    `protobuf:"varint,6,opt,name=e,enum=testprotos.TestMessage_NestedEnum,def=2" json:"e,omitempty"`
	F   []string                   `protobuf:"bytes,7,rep,name=f" json:"f,omitempty"`
	// Types that are valid to be assigned to Def:
	//	*Frobnitz_G1
	//	*Frobnitz_G2
	//	*Frobnitz_G3
	Def              isFrobnitz_Def `protobuf_oneof:"def"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Frobnitz) Reset()                    { *m = Frobnitz{} }
func (m *Frobnitz) String() string            { return proto.CompactTextString(m) }
func (*Frobnitz) ProtoMessage()               {}
func (*Frobnitz) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

const Default_Frobnitz_E TestMessage_NestedEnum = TestMessage_VALUE2

type isFrobnitz_Abc interface {
	isFrobnitz_Abc()
}
type isFrobnitz_Def interface {
	isFrobnitz_Def()
}

type Frobnitz_C1 struct {
	C1 *TestMessage_NestedMessage `protobuf:"bytes,3,opt,name=c1,oneof"`
}
type Frobnitz_C2 struct {
	C2 TestMessage_NestedEnum `protobuf:"varint,4,opt,name=c2,enum=testprotos.TestMessage_NestedEnum,oneof"`
}
type Frobnitz_G1 struct {
	G1 int32 `protobuf:"varint,8,opt,name=g1,oneof"`
}
type Frobnitz_G2 struct {
	G2 int32 `protobuf:"zigzag32,9,opt,name=g2,oneof"`
}
type Frobnitz_G3 struct {
	G3 uint32 `protobuf:"varint,10,opt,name=g3,oneof"`
}

func (*Frobnitz_C1) isFrobnitz_Abc() {}
func (*Frobnitz_C2) isFrobnitz_Abc() {}
func (*Frobnitz_G1) isFrobnitz_Def() {}
func (*Frobnitz_G2) isFrobnitz_Def() {}
func (*Frobnitz_G3) isFrobnitz_Def() {}

func (m *Frobnitz) GetAbc() isFrobnitz_Abc {
	if m != nil {
		return m.Abc
	}
	return nil
}
func (m *Frobnitz) GetDef() isFrobnitz_Def {
	if m != nil {
		return m.Def
	}
	return nil
}

func (m *Frobnitz) GetA() *TestMessage {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *Frobnitz) GetB() *AnotherTestMessage {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *Frobnitz) GetC1() *TestMessage_NestedMessage {
	if x, ok := m.GetAbc().(*Frobnitz_C1); ok {
		return x.C1
	}
	return nil
}

func (m *Frobnitz) GetC2() TestMessage_NestedEnum {
	if x, ok := m.GetAbc().(*Frobnitz_C2); ok {
		return x.C2
	}
	return TestMessage_VALUE1
}

func (m *Frobnitz) GetD() *TestMessage_NestedMessage {
	if m != nil {
		return m.D
	}
	return nil
}

func (m *Frobnitz) GetE() TestMessage_NestedEnum {
	if m != nil && m.E != nil {
		return *m.E
	}
	return Default_Frobnitz_E
}

func (m *Frobnitz) GetF() []string {
	if m != nil {
		return m.F
	}
	return nil
}

func (m *Frobnitz) GetG1() int32 {
	if x, ok := m.GetDef().(*Frobnitz_G1); ok {
		return x.G1
	}
	return 0
}

func (m *Frobnitz) GetG2() int32 {
	if x, ok := m.GetDef().(*Frobnitz_G2); ok {
		return x.G2
	}
	return 0
}

func (m *Frobnitz) GetG3() uint32 {
	if x, ok := m.GetDef().(*Frobnitz_G3); ok {
		return x.G3
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Frobnitz) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Frobnitz_OneofMarshaler, _Frobnitz_OneofUnmarshaler, _Frobnitz_OneofSizer, []interface{}{
		(*Frobnitz_C1)(nil),
		(*Frobnitz_C2)(nil),
		(*Frobnitz_G1)(nil),
		(*Frobnitz_G2)(nil),
		(*Frobnitz_G3)(nil),
	}
}

func _Frobnitz_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Frobnitz)
	// abc
	switch x := m.Abc.(type) {
	case *Frobnitz_C1:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.C1); err != nil {
			return err
		}
	case *Frobnitz_C2:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.C2))
	case nil:
	default:
		return fmt.Errorf("Frobnitz.Abc has unexpected type %T", x)
	}
	// def
	switch x := m.Def.(type) {
	case *Frobnitz_G1:
		b.EncodeVarint(8<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.G1))
	case *Frobnitz_G2:
		b.EncodeVarint(9<<3 | proto.WireVarint)
		b.EncodeZigzag32(uint64(x.G2))
	case *Frobnitz_G3:
		b.EncodeVarint(10<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.G3))
	case nil:
	default:
		return fmt.Errorf("Frobnitz.Def has unexpected type %T", x)
	}
	return nil
}

func _Frobnitz_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Frobnitz)
	switch tag {
	case 3: // abc.c1
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TestMessage_NestedMessage)
		err := b.DecodeMessage(msg)
		m.Abc = &Frobnitz_C1{msg}
		return true, err
	case 4: // abc.c2
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Abc = &Frobnitz_C2{TestMessage_NestedEnum(x)}
		return true, err
	case 8: // def.g1
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Def = &Frobnitz_G1{int32(x)}
		return true, err
	case 9: // def.g2
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeZigzag32()
		m.Def = &Frobnitz_G2{int32(x)}
		return true, err
	case 10: // def.g3
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Def = &Frobnitz_G3{uint32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Frobnitz_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Frobnitz)
	// abc
	switch x := m.Abc.(type) {
	case *Frobnitz_C1:
		s := proto.Size(x.C1)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Frobnitz_C2:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.C2))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// def
	switch x := m.Def.(type) {
	case *Frobnitz_G1:
		n += proto.SizeVarint(8<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.G1))
	case *Frobnitz_G2:
		n += proto.SizeVarint(9<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64((uint32(x.G2) << 1) ^ uint32((int32(x.G2) >> 31))))
	case *Frobnitz_G3:
		n += proto.SizeVarint(10<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.G3))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Whatchamacallit struct {
	Foos             *jhump_protoreflect_desc.Foo `protobuf:"varint,1,req,name=foos,enum=jhump.protoreflect.desc.Foo" json:"foos,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *Whatchamacallit) Reset()                    { *m = Whatchamacallit{} }
func (m *Whatchamacallit) String() string            { return proto.CompactTextString(m) }
func (*Whatchamacallit) ProtoMessage()               {}
func (*Whatchamacallit) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Whatchamacallit) GetFoos() jhump_protoreflect_desc.Foo {
	if m != nil && m.Foos != nil {
		return *m.Foos
	}
	return jhump_protoreflect_desc.Foo_ABC
}

type Whatzit struct {
	Gyzmeau          []*jhump_protoreflect_desc.Bar `protobuf:"bytes,1,rep,name=gyzmeau" json:"gyzmeau,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *Whatzit) Reset()                    { *m = Whatzit{} }
func (m *Whatzit) String() string            { return proto.CompactTextString(m) }
func (*Whatzit) ProtoMessage()               {}
func (*Whatzit) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Whatzit) GetGyzmeau() []*jhump_protoreflect_desc.Bar {
	if m != nil {
		return m.Gyzmeau
	}
	return nil
}

type GroupX struct {
	Groupxi          *int64  `protobuf:"varint,1041,opt,name=groupxi" json:"groupxi,omitempty"`
	Groupxs          *string `protobuf:"bytes,1042,opt,name=groupxs" json:"groupxs,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GroupX) Reset()                    { *m = GroupX{} }
func (m *GroupX) String() string            { return proto.CompactTextString(m) }
func (*GroupX) ProtoMessage()               {}
func (*GroupX) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *GroupX) GetGroupxi() int64 {
	if m != nil && m.Groupxi != nil {
		return *m.Groupxi
	}
	return 0
}

func (m *GroupX) GetGroupxs() string {
	if m != nil && m.Groupxs != nil {
		return *m.Groupxs
	}
	return ""
}

var E_Otl = &proto.ExtensionDesc{
	ExtendedType:  (*desc_test_nopkg.TopLevel)(nil),
	ExtensionType: (*desc_test_nopkg.TopLevel)(nil),
	Field:         100,
	Name:          "testprotos.otl",
	Tag:           "bytes,100,opt,name=otl",
	Filename:      "desc_test2.proto",
}

var E_Groupx = &proto.ExtensionDesc{
	ExtendedType:  (*desc_test_nopkg.TopLevel)(nil),
	ExtensionType: (*GroupX)(nil),
	Field:         104,
	Name:          "testprotos.groupx",
	Tag:           "group,104,opt,name=GroupX,json=groupx",
	Filename:      "desc_test2.proto",
}

func init() {
	proto.RegisterType((*Frobnitz)(nil), "testprotos.Frobnitz")
	proto.RegisterType((*Whatchamacallit)(nil), "testprotos.Whatchamacallit")
	proto.RegisterType((*Whatzit)(nil), "testprotos.Whatzit")
	proto.RegisterType((*GroupX)(nil), "testprotos.GroupX")
	proto.RegisterExtension(E_Otl)
	proto.RegisterExtension(E_Groupx)
}

func init() { proto.RegisterFile("desc_test2.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x19, 0xbb, 0x4d, 0x9a, 0xa9, 0x28, 0x61, 0x2f, 0x5d, 0x0a, 0x42, 0x56, 0xa4, 0x22,
	0x1f, 0x90, 0x83, 0x37, 0xfc, 0x53, 0x0e, 0x48, 0x09, 0x6a, 0xe9, 0xa1, 0x70, 0x58, 0x95, 0x3f,
	0xe2, 0x52, 0x6d, 0xec, 0x8d, 0x6d, 0xea, 0x78, 0x2d, 0xef, 0x06, 0x41, 0xbe, 0x05, 0xf0, 0x85,
	0x91, 0xd7, 0x09, 0x4e, 0x0f, 0x54, 0x70, 0xf2, 0xce, 0x7b, 0xbf, 0x99, 0xf1, 0xe8, 0x61, 0x3f,
	0x96, 0x3a, 0xba, 0x34, 0x52, 0x1b, 0x16, 0x94, 0x95, 0x32, 0x8a, 0x60, 0x5d, 0xd8, 0xa7, 0x3e,
	0x6a, 0xdd, 0xb0, 0x71, 0x8f, 0x0e, 0xcb, 0xab, 0x64, 0xf8, 0x47, 0xbd, 0x2c, 0xaf, 0x92, 0xb5,
	0x71, 0xbf, 0x50, 0xd7, 0x2d, 0x5b, 0x37, 0xe6, 0xe0, 0x97, 0x8b, 0x7b, 0xa7, 0x95, 0x9a, 0x15,
	0x99, 0x59, 0x91, 0x63, 0x04, 0x41, 0xc1, 0x03, 0x7f, 0x9f, 0x1d, 0x06, 0xed, 0xb2, 0xe0, 0x42,
	0x6a, 0xf3, 0x56, 0x6a, 0x2d, 0x12, 0xc9, 0x41, 0x90, 0xc7, 0x08, 0x33, 0xea, 0x58, 0xec, 0xe1,
	0x36, 0x36, 0x29, 0x94, 0x49, 0x65, 0x75, 0x8d, 0x9e, 0x91, 0x17, 0xe8, 0x44, 0x21, 0x75, 0x2d,
	0x7e, 0xfc, 0x97, 0xa9, 0xc1, 0x3b, 0xa9, 0x8d, 0x8c, 0xd7, 0xd5, 0xd9, 0x2d, 0xee, 0x44, 0x21,
	0x79, 0x8a, 0x4e, 0xc4, 0xe8, 0x8e, 0x07, 0xfe, 0x01, 0x1b, 0xdc, 0xdc, 0x78, 0x52, 0x2c, 0x17,
	0xb6, 0x8b, 0x91, 0x11, 0x42, 0x4c, 0x77, 0xff, 0x63, 0x1b, 0x87, 0x98, 0xbc, 0x44, 0x90, 0xb4,
	0xf3, 0xaf, 0x9b, 0xc6, 0x9d, 0x0f, 0x93, 0xf3, 0xf7, 0x27, 0x8c, 0x83, 0x24, 0x7d, 0x84, 0x39,
	0xed, 0x7a, 0xae, 0xdf, 0x9b, 0x3a, 0x14, 0x38, 0xcc, 0x49, 0x1f, 0x9d, 0x24, 0xa4, 0x7b, 0x1e,
	0xf8, 0xbb, 0x67, 0xc0, 0x9d, 0x24, 0xb4, 0x0a, 0xa3, 0x3d, 0x0f, 0xfc, 0xbb, 0x56, 0x61, 0x56,
	0x19, 0x51, 0xf4, 0xc0, 0xbf, 0x6d, 0x95, 0xd1, 0x74, 0x17, 0x5d, 0x31, 0x8b, 0xea, 0x4f, 0x2c,
	0xe7, 0x83, 0xd7, 0x78, 0xe7, 0x63, 0x2a, 0x4c, 0x94, 0x8a, 0x85, 0x88, 0x44, 0x9e, 0x67, 0x86,
	0x3c, 0xc1, 0x9d, 0xb9, 0x52, 0x9a, 0x82, 0xe7, 0xf8, 0x07, 0xec, 0x41, 0xf0, 0x25, 0x5d, 0x2e,
	0xca, 0x26, 0xc4, 0x4a, 0xce, 0x73, 0x19, 0x99, 0xa0, 0x4e, 0x38, 0x38, 0x55, 0x8a, 0x5b, 0x72,
	0x30, 0xc1, 0x6e, 0x3d, 0x64, 0x95, 0x19, 0xf2, 0x1c, 0xbb, 0xc9, 0xf7, 0xd5, 0x42, 0x8a, 0x25,
	0x05, 0xcf, 0xf5, 0xf7, 0x6f, 0xe8, 0x9f, 0x8a, 0x8a, 0x6f, 0xe0, 0xc1, 0x2b, 0xec, 0xbc, 0xa9,
	0xd4, 0xb2, 0xfc, 0x44, 0xee, 0x61, 0x37, 0xa9, 0x5f, 0xdf, 0x32, 0xfa, 0xa3, 0xbe, 0xcd, 0xe5,
	0x9b, 0xba, 0xb5, 0x34, 0xfd, 0x59, 0x5b, 0xbd, 0x8d, 0xa5, 0xc7, 0x8f, 0xd0, 0x55, 0x26, 0x27,
	0xbd, 0xe0, 0x42, 0x95, 0xe7, 0xf2, 0xab, 0xcc, 0x69, 0x6c, 0x93, 0x69, 0x05, 0x5e, 0x03, 0xe3,
	0x67, 0xd8, 0x69, 0x5a, 0xb6, 0xd1, 0xd4, 0x03, 0x1f, 0x19, 0xd9, 0xce, 0xa3, 0xf9, 0x19, 0xbe,
	0x86, 0xa7, 0xa3, 0xcf, 0x61, 0x92, 0x99, 0x74, 0x39, 0x0b, 0x22, 0xb5, 0x18, 0xda, 0x8b, 0x86,
	0xdb, 0x17, 0x0d, 0xb3, 0xc2, 0xc8, 0xaa, 0x10, 0xf9, 0xb0, 0x9d, 0xf1, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0x8e, 0xcf, 0x50, 0x97, 0x58, 0x03, 0x00, 0x00,
}
