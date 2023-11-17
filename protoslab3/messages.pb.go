// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protoslab3/messages.proto

package messages

import (
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

type Comando struct {
	Accion               string   `protobuf:"bytes,1,opt,name=accion,proto3" json:"accion,omitempty"`
	Sector               string   `protobuf:"bytes,2,opt,name=sector,proto3" json:"sector,omitempty"`
	Base                 string   `protobuf:"bytes,3,opt,name=base,proto3" json:"base,omitempty"`
	Valor                string   `protobuf:"bytes,4,opt,name=valor,proto3" json:"valor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Comando) Reset()         { *m = Comando{} }
func (m *Comando) String() string { return proto.CompactTextString(m) }
func (*Comando) ProtoMessage()    {}
func (*Comando) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2e02856b5ed9737, []int{0}
}

func (m *Comando) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Comando.Unmarshal(m, b)
}
func (m *Comando) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Comando.Marshal(b, m, deterministic)
}
func (m *Comando) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comando.Merge(m, src)
}
func (m *Comando) XXX_Size() int {
	return xxx_messageInfo_Comando.Size(m)
}
func (m *Comando) XXX_DiscardUnknown() {
	xxx_messageInfo_Comando.DiscardUnknown(m)
}

var xxx_messageInfo_Comando proto.InternalMessageInfo

func (m *Comando) GetAccion() string {
	if m != nil {
		return m.Accion
	}
	return ""
}

func (m *Comando) GetSector() string {
	if m != nil {
		return m.Sector
	}
	return ""
}

func (m *Comando) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

func (m *Comando) GetValor() string {
	if m != nil {
		return m.Valor
	}
	return ""
}

type Direccion struct {
	Direccion            string   `protobuf:"bytes,1,opt,name=direccion,proto3" json:"direccion,omitempty"`
	Fulcrum              string   `protobuf:"bytes,2,opt,name=fulcrum,proto3" json:"fulcrum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Direccion) Reset()         { *m = Direccion{} }
func (m *Direccion) String() string { return proto.CompactTextString(m) }
func (*Direccion) ProtoMessage()    {}
func (*Direccion) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2e02856b5ed9737, []int{1}
}

func (m *Direccion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Direccion.Unmarshal(m, b)
}
func (m *Direccion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Direccion.Marshal(b, m, deterministic)
}
func (m *Direccion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Direccion.Merge(m, src)
}
func (m *Direccion) XXX_Size() int {
	return xxx_messageInfo_Direccion.Size(m)
}
func (m *Direccion) XXX_DiscardUnknown() {
	xxx_messageInfo_Direccion.DiscardUnknown(m)
}

var xxx_messageInfo_Direccion proto.InternalMessageInfo

func (m *Direccion) GetDireccion() string {
	if m != nil {
		return m.Direccion
	}
	return ""
}

func (m *Direccion) GetFulcrum() string {
	if m != nil {
		return m.Fulcrum
	}
	return ""
}

type Vector struct {
	X                    string   `protobuf:"bytes,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    string   `protobuf:"bytes,2,opt,name=y,proto3" json:"y,omitempty"`
	Z                    string   `protobuf:"bytes,3,opt,name=z,proto3" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vector) Reset()         { *m = Vector{} }
func (m *Vector) String() string { return proto.CompactTextString(m) }
func (*Vector) ProtoMessage()    {}
func (*Vector) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2e02856b5ed9737, []int{2}
}

func (m *Vector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vector.Unmarshal(m, b)
}
func (m *Vector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vector.Marshal(b, m, deterministic)
}
func (m *Vector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vector.Merge(m, src)
}
func (m *Vector) XXX_Size() int {
	return xxx_messageInfo_Vector.Size(m)
}
func (m *Vector) XXX_DiscardUnknown() {
	xxx_messageInfo_Vector.DiscardUnknown(m)
}

var xxx_messageInfo_Vector proto.InternalMessageInfo

func (m *Vector) GetX() string {
	if m != nil {
		return m.X
	}
	return ""
}

func (m *Vector) GetY() string {
	if m != nil {
		return m.Y
	}
	return ""
}

func (m *Vector) GetZ() string {
	if m != nil {
		return m.Z
	}
	return ""
}

type Sector struct {
	Sector               string   `protobuf:"bytes,1,opt,name=sector,proto3" json:"sector,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sector) Reset()         { *m = Sector{} }
func (m *Sector) String() string { return proto.CompactTextString(m) }
func (*Sector) ProtoMessage()    {}
func (*Sector) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2e02856b5ed9737, []int{3}
}

func (m *Sector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sector.Unmarshal(m, b)
}
func (m *Sector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sector.Marshal(b, m, deterministic)
}
func (m *Sector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sector.Merge(m, src)
}
func (m *Sector) XXX_Size() int {
	return xxx_messageInfo_Sector.Size(m)
}
func (m *Sector) XXX_DiscardUnknown() {
	xxx_messageInfo_Sector.DiscardUnknown(m)
}

var xxx_messageInfo_Sector proto.InternalMessageInfo

func (m *Sector) GetSector() string {
	if m != nil {
		return m.Sector
	}
	return ""
}

type Base struct {
	Base                 string   `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Base) Reset()         { *m = Base{} }
func (m *Base) String() string { return proto.CompactTextString(m) }
func (*Base) ProtoMessage()    {}
func (*Base) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2e02856b5ed9737, []int{4}
}

func (m *Base) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Base.Unmarshal(m, b)
}
func (m *Base) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Base.Marshal(b, m, deterministic)
}
func (m *Base) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Base.Merge(m, src)
}
func (m *Base) XXX_Size() int {
	return xxx_messageInfo_Base.Size(m)
}
func (m *Base) XXX_DiscardUnknown() {
	xxx_messageInfo_Base.DiscardUnknown(m)
}

var xxx_messageInfo_Base proto.InternalMessageInfo

func (m *Base) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

type Confirmar struct {
	Flag                 string   `protobuf:"bytes,1,opt,name=flag,proto3" json:"flag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Confirmar) Reset()         { *m = Confirmar{} }
func (m *Confirmar) String() string { return proto.CompactTextString(m) }
func (*Confirmar) ProtoMessage()    {}
func (*Confirmar) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2e02856b5ed9737, []int{5}
}

func (m *Confirmar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Confirmar.Unmarshal(m, b)
}
func (m *Confirmar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Confirmar.Marshal(b, m, deterministic)
}
func (m *Confirmar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Confirmar.Merge(m, src)
}
func (m *Confirmar) XXX_Size() int {
	return xxx_messageInfo_Confirmar.Size(m)
}
func (m *Confirmar) XXX_DiscardUnknown() {
	xxx_messageInfo_Confirmar.DiscardUnknown(m)
}

var xxx_messageInfo_Confirmar proto.InternalMessageInfo

func (m *Confirmar) GetFlag() string {
	if m != nil {
		return m.Flag
	}
	return ""
}

type Soldados struct {
	Sector               string   `protobuf:"bytes,1,opt,name=sector,proto3" json:"sector,omitempty"`
	Base                 string   `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Soldados) Reset()         { *m = Soldados{} }
func (m *Soldados) String() string { return proto.CompactTextString(m) }
func (*Soldados) ProtoMessage()    {}
func (*Soldados) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2e02856b5ed9737, []int{6}
}

func (m *Soldados) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Soldados.Unmarshal(m, b)
}
func (m *Soldados) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Soldados.Marshal(b, m, deterministic)
}
func (m *Soldados) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Soldados.Merge(m, src)
}
func (m *Soldados) XXX_Size() int {
	return xxx_messageInfo_Soldados.Size(m)
}
func (m *Soldados) XXX_DiscardUnknown() {
	xxx_messageInfo_Soldados.DiscardUnknown(m)
}

var xxx_messageInfo_Soldados proto.InternalMessageInfo

func (m *Soldados) GetSector() string {
	if m != nil {
		return m.Sector
	}
	return ""
}

func (m *Soldados) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

type Numero struct {
	Numero               string   `protobuf:"bytes,1,opt,name=numero,proto3" json:"numero,omitempty"`
	X                    string   `protobuf:"bytes,2,opt,name=x,proto3" json:"x,omitempty"`
	Y                    string   `protobuf:"bytes,3,opt,name=y,proto3" json:"y,omitempty"`
	Z                    string   `protobuf:"bytes,4,opt,name=z,proto3" json:"z,omitempty"`
	Fulcrum              string   `protobuf:"bytes,5,opt,name=fulcrum,proto3" json:"fulcrum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Numero) Reset()         { *m = Numero{} }
func (m *Numero) String() string { return proto.CompactTextString(m) }
func (*Numero) ProtoMessage()    {}
func (*Numero) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2e02856b5ed9737, []int{7}
}

func (m *Numero) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Numero.Unmarshal(m, b)
}
func (m *Numero) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Numero.Marshal(b, m, deterministic)
}
func (m *Numero) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Numero.Merge(m, src)
}
func (m *Numero) XXX_Size() int {
	return xxx_messageInfo_Numero.Size(m)
}
func (m *Numero) XXX_DiscardUnknown() {
	xxx_messageInfo_Numero.DiscardUnknown(m)
}

var xxx_messageInfo_Numero proto.InternalMessageInfo

func (m *Numero) GetNumero() string {
	if m != nil {
		return m.Numero
	}
	return ""
}

func (m *Numero) GetX() string {
	if m != nil {
		return m.X
	}
	return ""
}

func (m *Numero) GetY() string {
	if m != nil {
		return m.Y
	}
	return ""
}

func (m *Numero) GetZ() string {
	if m != nil {
		return m.Z
	}
	return ""
}

func (m *Numero) GetFulcrum() string {
	if m != nil {
		return m.Fulcrum
	}
	return ""
}

type Cambio struct {
	Nombre               string   `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Tiempo1              string   `protobuf:"bytes,2,opt,name=tiempo1,proto3" json:"tiempo1,omitempty"`
	Valor                string   `protobuf:"bytes,3,opt,name=valor,proto3" json:"valor,omitempty"`
	Tiempo2              string   `protobuf:"bytes,4,opt,name=tiempo2,proto3" json:"tiempo2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cambio) Reset()         { *m = Cambio{} }
func (m *Cambio) String() string { return proto.CompactTextString(m) }
func (*Cambio) ProtoMessage()    {}
func (*Cambio) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2e02856b5ed9737, []int{8}
}

func (m *Cambio) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cambio.Unmarshal(m, b)
}
func (m *Cambio) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cambio.Marshal(b, m, deterministic)
}
func (m *Cambio) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cambio.Merge(m, src)
}
func (m *Cambio) XXX_Size() int {
	return xxx_messageInfo_Cambio.Size(m)
}
func (m *Cambio) XXX_DiscardUnknown() {
	xxx_messageInfo_Cambio.DiscardUnknown(m)
}

var xxx_messageInfo_Cambio proto.InternalMessageInfo

func (m *Cambio) GetNombre() string {
	if m != nil {
		return m.Nombre
	}
	return ""
}

func (m *Cambio) GetTiempo1() string {
	if m != nil {
		return m.Tiempo1
	}
	return ""
}

func (m *Cambio) GetValor() string {
	if m != nil {
		return m.Valor
	}
	return ""
}

func (m *Cambio) GetTiempo2() string {
	if m != nil {
		return m.Tiempo2
	}
	return ""
}

type SectorVector struct {
	Sector               string   `protobuf:"bytes,1,opt,name=sector,proto3" json:"sector,omitempty"`
	X                    string   `protobuf:"bytes,2,opt,name=x,proto3" json:"x,omitempty"`
	Y                    string   `protobuf:"bytes,3,opt,name=y,proto3" json:"y,omitempty"`
	Z                    string   `protobuf:"bytes,4,opt,name=z,proto3" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SectorVector) Reset()         { *m = SectorVector{} }
func (m *SectorVector) String() string { return proto.CompactTextString(m) }
func (*SectorVector) ProtoMessage()    {}
func (*SectorVector) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2e02856b5ed9737, []int{9}
}

func (m *SectorVector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SectorVector.Unmarshal(m, b)
}
func (m *SectorVector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SectorVector.Marshal(b, m, deterministic)
}
func (m *SectorVector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SectorVector.Merge(m, src)
}
func (m *SectorVector) XXX_Size() int {
	return xxx_messageInfo_SectorVector.Size(m)
}
func (m *SectorVector) XXX_DiscardUnknown() {
	xxx_messageInfo_SectorVector.DiscardUnknown(m)
}

var xxx_messageInfo_SectorVector proto.InternalMessageInfo

func (m *SectorVector) GetSector() string {
	if m != nil {
		return m.Sector
	}
	return ""
}

func (m *SectorVector) GetX() string {
	if m != nil {
		return m.X
	}
	return ""
}

func (m *SectorVector) GetY() string {
	if m != nil {
		return m.Y
	}
	return ""
}

func (m *SectorVector) GetZ() string {
	if m != nil {
		return m.Z
	}
	return ""
}

func init() {
	proto.RegisterType((*Comando)(nil), "Comando")
	proto.RegisterType((*Direccion)(nil), "Direccion")
	proto.RegisterType((*Vector)(nil), "Vector")
	proto.RegisterType((*Sector)(nil), "Sector")
	proto.RegisterType((*Base)(nil), "Base")
	proto.RegisterType((*Confirmar)(nil), "Confirmar")
	proto.RegisterType((*Soldados)(nil), "Soldados")
	proto.RegisterType((*Numero)(nil), "Numero")
	proto.RegisterType((*Cambio)(nil), "Cambio")
	proto.RegisterType((*SectorVector)(nil), "SectorVector")
}

func init() { proto.RegisterFile("protoslab3/messages.proto", fileDescriptor_d2e02856b5ed9737) }

var fileDescriptor_d2e02856b5ed9737 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x6d, 0x6b, 0xd4, 0x40,
	0x10, 0xbe, 0xbd, 0x97, 0xa4, 0x99, 0x9e, 0x2d, 0xae, 0x22, 0x31, 0x28, 0xd6, 0x2d, 0xa8, 0x20,
	0xae, 0x78, 0x05, 0x7f, 0xc0, 0x9d, 0x28, 0x7e, 0x50, 0x8a, 0x07, 0xfd, 0xe0, 0x17, 0xd9, 0x24,
	0x7b, 0x65, 0x35, 0xc9, 0x96, 0xdd, 0x5c, 0x69, 0xfb, 0x03, 0xfc, 0x21, 0xfe, 0x52, 0xc9, 0xbe,
	0xe4, 0x45, 0xa8, 0x5c, 0x3f, 0x65, 0x9e, 0x61, 0x66, 0x9e, 0x79, 0x9e, 0x59, 0x02, 0x8f, 0x2f,
	0x94, 0xac, 0xa5, 0x2e, 0x58, 0x7a, 0xf2, 0xb6, 0xe4, 0x5a, 0xb3, 0x73, 0xae, 0xa9, 0xc9, 0x91,
	0x0c, 0xc2, 0x95, 0x2c, 0x59, 0x95, 0x4b, 0xfc, 0x08, 0x02, 0x96, 0x65, 0x42, 0x56, 0x31, 0x3a,
	0x42, 0xaf, 0xa2, 0x6f, 0x0e, 0x35, 0x79, 0xcd, 0xb3, 0x5a, 0xaa, 0x78, 0x6c, 0xf3, 0x16, 0x61,
	0x0c, 0xd3, 0x94, 0x69, 0x1e, 0x4f, 0x4c, 0xd6, 0xc4, 0xf8, 0x21, 0xcc, 0x2e, 0x59, 0x21, 0x55,
	0x3c, 0x35, 0x49, 0x0b, 0xc8, 0x0a, 0xa2, 0x0f, 0x42, 0x71, 0x3b, 0xee, 0x09, 0x44, 0xb9, 0x07,
	0x8e, 0xa9, 0x4b, 0xe0, 0x18, 0xc2, 0xcd, 0xb6, 0xc8, 0xd4, 0xb6, 0x74, 0x6c, 0x1e, 0x92, 0x05,
	0x04, 0x67, 0x96, 0x78, 0x0e, 0xe8, 0xca, 0x75, 0xa2, 0xab, 0x06, 0x5d, 0xbb, 0x5a, 0x74, 0xdd,
	0xa0, 0x1b, 0xb7, 0x11, 0xba, 0x21, 0x47, 0x10, 0xac, 0x6d, 0x4f, 0x27, 0x02, 0xf5, 0x45, 0x90,
	0x04, 0xa6, 0xcb, 0x66, 0x71, 0x2f, 0x06, 0x75, 0x62, 0xc8, 0x33, 0x88, 0x56, 0xb2, 0xda, 0x08,
	0x55, 0x32, 0xa3, 0x76, 0x53, 0xb0, 0x73, 0x5f, 0xd0, 0xc4, 0xe4, 0x3d, 0xec, 0xad, 0x65, 0x91,
	0xb3, 0x5c, 0xea, 0xdb, 0x08, 0xda, 0xc1, 0xe3, 0xde, 0xe0, 0x1c, 0x82, 0xaf, 0xdb, 0x92, 0x2b,
	0xe3, 0x79, 0x65, 0x22, 0xdf, 0x65, 0x91, 0x95, 0x38, 0x1e, 0x48, 0x9c, 0x0c, 0x24, 0x4e, 0x9d,
	0xc4, 0xbe, 0x61, 0xb3, 0xa1, 0x61, 0x3f, 0x21, 0x58, 0xb1, 0x32, 0x15, 0x96, 0x45, 0x96, 0xa9,
	0xe2, 0x2d, 0x8b, 0x41, 0x4d, 0x6f, 0x2d, 0x78, 0x79, 0x21, 0xdf, 0x79, 0xb3, 0x1d, 0xec, 0xee,
	0x38, 0xe9, 0xdd, 0xb1, 0xab, 0x5f, 0x38, 0x7e, 0x0f, 0xc9, 0x29, 0xcc, 0xad, 0xd1, 0x67, 0xff,
	0xb5, 0x7b, 0x77, 0x5d, 0x8b, 0xdf, 0x08, 0x0e, 0x96, 0x4a, 0xfe, 0xe2, 0xea, 0xc7, 0x9a, 0xab,
	0x4b, 0x91, 0x71, 0xfc, 0x02, 0xe6, 0x5f, 0x58, 0x95, 0x33, 0xe5, 0x64, 0xed, 0x51, 0xf7, 0x74,
	0x13, 0xa0, 0xed, 0xfb, 0x22, 0x23, 0x7c, 0x0c, 0xfb, 0x9f, 0x78, 0xdd, 0x5e, 0x26, 0xa2, 0x3e,
	0x4c, 0x42, 0x6a, 0x7d, 0x27, 0x23, 0xfc, 0x12, 0x0e, 0x3e, 0x57, 0x99, 0xac, 0xb4, 0xd0, 0x35,
	0xaf, 0x32, 0xc1, 0x70, 0x48, 0xad, 0x84, 0x04, 0x68, 0x7b, 0x76, 0x32, 0x5a, 0xfc, 0x19, 0xc3,
	0xe1, 0x47, 0x6b, 0x69, 0xbb, 0xc9, 0xf1, 0xad, 0x9b, 0x84, 0xd4, 0x3a, 0x70, 0x07, 0x06, 0xfc,
	0x1c, 0xf6, 0x4f, 0x79, 0x2e, 0xbc, 0x77, 0x6d, 0x55, 0x6f, 0xd6, 0x6b, 0x4f, 0xe8, 0x6a, 0xee,
	0xd1, 0xbe, 0xdd, 0xff, 0xcc, 0x7b, 0x0a, 0x91, 0x99, 0x67, 0x1e, 0xf6, 0x8c, 0x36, 0x9f, 0x24,
	0xa4, 0x76, 0x55, 0x32, 0xc2, 0x6f, 0xe0, 0x41, 0x7f, 0x79, 0xa7, 0x6d, 0xe0, 0x66, 0x7f, 0xda,
	0x2e, 0x6e, 0x2e, 0xef, 0x7f, 0x3f, 0xb4, 0xff, 0x98, 0xf6, 0xff, 0x92, 0x06, 0x26, 0x71, 0xf2,
	0x37, 0x00, 0x00, 0xff, 0xff, 0xea, 0xf8, 0xd5, 0x73, 0x7d, 0x04, 0x00, 0x00,
}
