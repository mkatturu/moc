// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_common_security.proto

package common

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

type Algorithm int32

const (
	Algorithm_A_UNKNOWN  Algorithm = 0
	Algorithm_RSA15      Algorithm = 1
	Algorithm_RSAOAEP    Algorithm = 2
	Algorithm_RSAOAEP256 Algorithm = 3
	Algorithm_A256KW     Algorithm = 4
)

var Algorithm_name = map[int32]string{
	0: "A_UNKNOWN",
	1: "RSA15",
	2: "RSAOAEP",
	3: "RSAOAEP256",
	4: "A256KW",
}

var Algorithm_value = map[string]int32{
	"A_UNKNOWN":  0,
	"RSA15":      1,
	"RSAOAEP":    2,
	"RSAOAEP256": 3,
	"A256KW":     4,
}

func (x Algorithm) String() string {
	return proto.EnumName(Algorithm_name, int32(x))
}

func (Algorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d3874efde778ac1, []int{0}
}

type KeyOperation int32

const (
	KeyOperation_ENCRYPT   KeyOperation = 0
	KeyOperation_DECRYPT   KeyOperation = 1
	KeyOperation_WRAPKEY   KeyOperation = 2
	KeyOperation_UNWRAPKEY KeyOperation = 3
)

var KeyOperation_name = map[int32]string{
	0: "ENCRYPT",
	1: "DECRYPT",
	2: "WRAPKEY",
	3: "UNWRAPKEY",
}

var KeyOperation_value = map[string]int32{
	"ENCRYPT":   0,
	"DECRYPT":   1,
	"WRAPKEY":   2,
	"UNWRAPKEY": 3,
}

func (x KeyOperation) String() string {
	return proto.EnumName(KeyOperation_name, int32(x))
}

func (KeyOperation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d3874efde778ac1, []int{1}
}

// https://docs.microsoft.com/en-us/rest/api/keyvault/createkey/createkey#jsonwebkeytype
type JsonWebKeyType int32

const (
	JsonWebKeyType_EC      JsonWebKeyType = 0
	JsonWebKeyType_EC_HSM  JsonWebKeyType = 1
	JsonWebKeyType_RSA     JsonWebKeyType = 2
	JsonWebKeyType_RSA_HSM JsonWebKeyType = 3
	JsonWebKeyType_OCT     JsonWebKeyType = 4
	JsonWebKeyType_AES     JsonWebKeyType = 5
)

var JsonWebKeyType_name = map[int32]string{
	0: "EC",
	1: "EC_HSM",
	2: "RSA",
	3: "RSA_HSM",
	4: "OCT",
	5: "AES",
}

var JsonWebKeyType_value = map[string]int32{
	"EC":      0,
	"EC_HSM":  1,
	"RSA":     2,
	"RSA_HSM": 3,
	"OCT":     4,
	"AES":     5,
}

func (x JsonWebKeyType) String() string {
	return proto.EnumName(JsonWebKeyType_name, int32(x))
}

func (JsonWebKeyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d3874efde778ac1, []int{2}
}

type JsonWebKeyCurveName int32

const (
	JsonWebKeyCurveName_P_256  JsonWebKeyCurveName = 0
	JsonWebKeyCurveName_P_256K JsonWebKeyCurveName = 1
	JsonWebKeyCurveName_P_384  JsonWebKeyCurveName = 2
	JsonWebKeyCurveName_P_521  JsonWebKeyCurveName = 3
)

var JsonWebKeyCurveName_name = map[int32]string{
	0: "P_256",
	1: "P_256K",
	2: "P_384",
	3: "P_521",
}

var JsonWebKeyCurveName_value = map[string]int32{
	"P_256":  0,
	"P_256K": 1,
	"P_384":  2,
	"P_521":  3,
}

func (x JsonWebKeyCurveName) String() string {
	return proto.EnumName(JsonWebKeyCurveName_name, int32(x))
}

func (JsonWebKeyCurveName) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d3874efde778ac1, []int{3}
}

type KeySize int32

const (
	KeySize_K_UNKNOWN KeySize = 0
	KeySize__256      KeySize = 1
	KeySize__2048     KeySize = 2
	KeySize__3072     KeySize = 3
	KeySize__4096     KeySize = 4
)

var KeySize_name = map[int32]string{
	0: "K_UNKNOWN",
	1: "_256",
	2: "_2048",
	3: "_3072",
	4: "_4096",
}

var KeySize_value = map[string]int32{
	"K_UNKNOWN": 0,
	"_256":      1,
	"_2048":     2,
	"_3072":     3,
	"_4096":     4,
}

func (x KeySize) String() string {
	return proto.EnumName(KeySize_name, int32(x))
}

func (KeySize) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d3874efde778ac1, []int{4}
}

type IdentityOperation int32

const (
	IdentityOperation_REVOKE IdentityOperation = 0
	IdentityOperation_ROTATE IdentityOperation = 1
)

var IdentityOperation_name = map[int32]string{
	0: "REVOKE",
	1: "ROTATE",
}

var IdentityOperation_value = map[string]int32{
	"REVOKE": 0,
	"ROTATE": 1,
}

func (x IdentityOperation) String() string {
	return proto.EnumName(IdentityOperation_name, int32(x))
}

func (IdentityOperation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d3874efde778ac1, []int{5}
}

type IdentityCertificateOperation int32

const (
	IdentityCertificateOperation_CREATE_CERTIFICATE IdentityCertificateOperation = 0
	IdentityCertificateOperation_RENEW_CERTIFICATE  IdentityCertificateOperation = 1
)

var IdentityCertificateOperation_name = map[int32]string{
	0: "CREATE_CERTIFICATE",
	1: "RENEW_CERTIFICATE",
}

var IdentityCertificateOperation_value = map[string]int32{
	"CREATE_CERTIFICATE": 0,
	"RENEW_CERTIFICATE":  1,
}

func (x IdentityCertificateOperation) String() string {
	return proto.EnumName(IdentityCertificateOperation_name, int32(x))
}

func (IdentityCertificateOperation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d3874efde778ac1, []int{6}
}

type KeyWrappingAlgorithm int32

const (
	KeyWrappingAlgorithm_CKM_RSA_AES_KEY_WRAP KeyWrappingAlgorithm = 0
)

var KeyWrappingAlgorithm_name = map[int32]string{
	0: "CKM_RSA_AES_KEY_WRAP",
}

var KeyWrappingAlgorithm_value = map[string]int32{
	"CKM_RSA_AES_KEY_WRAP": 0,
}

func (x KeyWrappingAlgorithm) String() string {
	return proto.EnumName(KeyWrappingAlgorithm_name, int32(x))
}

func (KeyWrappingAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d3874efde778ac1, []int{7}
}

type Scope struct {
	Location             string       `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	ResourceGroup        string       `protobuf:"bytes,2,opt,name=resourceGroup,proto3" json:"resourceGroup,omitempty"`
	ProviderType         ProviderType `protobuf:"varint,3,opt,name=providerType,proto3,enum=moc.ProviderType" json:"providerType,omitempty"`
	Resource             string       `protobuf:"bytes,4,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Scope) Reset()         { *m = Scope{} }
func (m *Scope) String() string { return proto.CompactTextString(m) }
func (*Scope) ProtoMessage()    {}
func (*Scope) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d3874efde778ac1, []int{0}
}

func (m *Scope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Scope.Unmarshal(m, b)
}
func (m *Scope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Scope.Marshal(b, m, deterministic)
}
func (m *Scope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scope.Merge(m, src)
}
func (m *Scope) XXX_Size() int {
	return xxx_messageInfo_Scope.Size(m)
}
func (m *Scope) XXX_DiscardUnknown() {
	xxx_messageInfo_Scope.DiscardUnknown(m)
}

var xxx_messageInfo_Scope proto.InternalMessageInfo

func (m *Scope) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Scope) GetResourceGroup() string {
	if m != nil {
		return m.ResourceGroup
	}
	return ""
}

func (m *Scope) GetProviderType() ProviderType {
	if m != nil {
		return m.ProviderType
	}
	return ProviderType_AnyProvider
}

func (m *Scope) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func init() {
	proto.RegisterEnum("moc.Algorithm", Algorithm_name, Algorithm_value)
	proto.RegisterEnum("moc.KeyOperation", KeyOperation_name, KeyOperation_value)
	proto.RegisterEnum("moc.JsonWebKeyType", JsonWebKeyType_name, JsonWebKeyType_value)
	proto.RegisterEnum("moc.JsonWebKeyCurveName", JsonWebKeyCurveName_name, JsonWebKeyCurveName_value)
	proto.RegisterEnum("moc.KeySize", KeySize_name, KeySize_value)
	proto.RegisterEnum("moc.IdentityOperation", IdentityOperation_name, IdentityOperation_value)
	proto.RegisterEnum("moc.IdentityCertificateOperation", IdentityCertificateOperation_name, IdentityCertificateOperation_value)
	proto.RegisterEnum("moc.KeyWrappingAlgorithm", KeyWrappingAlgorithm_name, KeyWrappingAlgorithm_value)
	proto.RegisterType((*Scope)(nil), "moc.Scope")
}

func init() { proto.RegisterFile("moc_common_security.proto", fileDescriptor_0d3874efde778ac1) }

var fileDescriptor_0d3874efde778ac1 = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x52, 0x6d, 0x4f, 0xdb, 0x30,
	0x10, 0x6e, 0x9a, 0xd2, 0xd2, 0xe3, 0x45, 0xc6, 0x63, 0x5b, 0x87, 0xf6, 0x01, 0xed, 0x45, 0x42,
	0x99, 0xd4, 0x96, 0x40, 0x19, 0xfb, 0x68, 0x82, 0xb7, 0x31, 0x8f, 0xa4, 0x72, 0xc2, 0x22, 0xf6,
	0x25, 0x2a, 0xc6, 0x40, 0x24, 0x52, 0x47, 0x6e, 0x8a, 0x94, 0xfd, 0x96, 0xfd, 0xd8, 0xc9, 0x09,
	0x1d, 0xe5, 0x93, 0xef, 0x79, 0xee, 0xee, 0xb9, 0x17, 0x1f, 0xbc, 0xc9, 0x94, 0x48, 0x84, 0xca,
	0x32, 0x35, 0x4d, 0x66, 0x52, 0xcc, 0x75, 0x5a, 0x94, 0xfd, 0x5c, 0xab, 0x42, 0x61, 0x3b, 0x53,
	0x62, 0xe7, 0xf5, 0x92, 0xbf, 0x7e, 0x6a, 0xef, 0xbb, 0xbf, 0x16, 0xac, 0x84, 0x42, 0xe5, 0x12,
	0xef, 0xc0, 0xea, 0xbd, 0x12, 0x93, 0x22, 0x55, 0xd3, 0x9e, 0xb5, 0x6b, 0xed, 0x75, 0xf9, 0x7f,
	0x8c, 0x3f, 0xc0, 0x86, 0x96, 0x33, 0x35, 0xd7, 0x42, 0x7e, 0xd3, 0x6a, 0x9e, 0xf7, 0x9a, 0x55,
	0xc0, 0x73, 0x12, 0x8f, 0x60, 0x3d, 0xd7, 0xea, 0x21, 0xbd, 0x96, 0x3a, 0x2a, 0x73, 0xd9, 0xb3,
	0x77, 0xad, 0xbd, 0x4d, 0x77, 0xab, 0x9f, 0x29, 0xd1, 0x1f, 0x2f, 0x39, 0xf8, 0xb3, 0x30, 0x53,
	0x78, 0xa1, 0xd3, 0x6b, 0xd5, 0x85, 0x17, 0xd8, 0xf1, 0xa1, 0x4b, 0xee, 0x6f, 0x95, 0x4e, 0x8b,
	0xbb, 0x0c, 0x6f, 0x40, 0x97, 0x24, 0x17, 0x3e, 0xf3, 0x83, 0xd8, 0x47, 0x0d, 0xdc, 0x85, 0x15,
	0x1e, 0x92, 0xfd, 0x11, 0xb2, 0xf0, 0x1a, 0x74, 0x78, 0x48, 0x02, 0x42, 0xc7, 0xa8, 0x89, 0x37,
	0x01, 0x1e, 0x81, 0x3b, 0x3a, 0x42, 0x36, 0x06, 0x68, 0x13, 0x77, 0x74, 0xc4, 0x62, 0xd4, 0x72,
	0x4e, 0x61, 0x9d, 0xc9, 0x32, 0xc8, 0xa5, 0xae, 0x07, 0x5b, 0x83, 0x0e, 0xf5, 0x3d, 0x7e, 0x39,
	0x8e, 0x50, 0xc3, 0x80, 0x53, 0x5a, 0x83, 0x4a, 0x32, 0xe6, 0x64, 0xcc, 0xe8, 0x25, 0x6a, 0x9a,
	0xca, 0x17, 0xfe, 0x02, 0xda, 0xce, 0x4f, 0xd8, 0xfc, 0x31, 0x53, 0xd3, 0x58, 0x5e, 0x31, 0x59,
	0x56, 0x33, 0xb4, 0xa1, 0x49, 0x3d, 0xd4, 0x30, 0xb5, 0xa8, 0x97, 0x7c, 0x0f, 0xcf, 0x91, 0x85,
	0x3b, 0x60, 0xf3, 0x90, 0xa0, 0xe6, 0x63, 0x77, 0x15, 0x6b, 0x1b, 0x36, 0xf0, 0x22, 0xd4, 0x32,
	0x06, 0xa1, 0x21, 0x5a, 0x71, 0x4e, 0xe0, 0xc5, 0x93, 0x9a, 0x37, 0xd7, 0x0f, 0xd2, 0x9f, 0x64,
	0xd2, 0x8c, 0x37, 0x4e, 0xcc, 0x04, 0x95, 0x6a, 0x65, 0x32, 0x64, 0xd5, 0xf4, 0xc1, 0xf1, 0x21,
	0x6a, 0xd6, 0xe6, 0xc8, 0xdd, 0x47, 0xb6, 0xe3, 0x41, 0x87, 0xc9, 0x32, 0x4c, 0xff, 0x48, 0xd3,
	0x2b, 0x5b, 0xda, 0xd2, 0x2a, 0xb4, 0x2a, 0x95, 0x2a, 0x33, 0x71, 0x87, 0x87, 0xc7, 0x75, 0x66,
	0x72, 0x30, 0xfc, 0xec, 0x22, 0xbb, 0x32, 0x0f, 0x87, 0x5f, 0x8e, 0x50, 0xcb, 0xf9, 0x04, 0x5b,
	0x67, 0xd7, 0x72, 0x5a, 0xa4, 0xc5, 0xd2, 0x86, 0x00, 0xda, 0x9c, 0xfe, 0x0a, 0x18, 0xad, 0xfb,
	0xe0, 0x41, 0x44, 0x22, 0x8a, 0x2c, 0xe7, 0x1c, 0xde, 0x2e, 0x82, 0x3d, 0xa9, 0x8b, 0xf4, 0x26,
	0x15, 0x93, 0x42, 0x3e, 0xe5, 0xbd, 0x02, 0xec, 0x71, 0x4a, 0x22, 0x9a, 0x78, 0x94, 0x47, 0x67,
	0x5f, 0xcf, 0x3c, 0x93, 0xd7, 0xc0, 0x2f, 0x61, 0x8b, 0x53, 0x9f, 0xc6, 0xcf, 0x68, 0xcb, 0x19,
	0xc2, 0x36, 0x93, 0x65, 0xac, 0x27, 0x79, 0x9e, 0x4e, 0x6f, 0x9f, 0xfe, 0xbc, 0x07, 0xdb, 0x1e,
	0x3b, 0x4f, 0xcc, 0xfe, 0x08, 0x0d, 0x13, 0x46, 0x2f, 0x13, 0xf3, 0x0f, 0xa8, 0x71, 0xf2, 0xf1,
	0xf7, 0xfb, 0xdb, 0xb4, 0xb8, 0x9b, 0x5f, 0xf5, 0x85, 0xca, 0x06, 0x59, 0x2a, 0xb4, 0x9a, 0xa9,
	0x9b, 0x62, 0x90, 0x29, 0x31, 0xd0, 0xb9, 0x18, 0xd4, 0x67, 0x7e, 0xd5, 0xae, 0xee, 0xfc, 0xe0,
	0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x37, 0xd7, 0x59, 0xa6, 0x22, 0x03, 0x00, 0x00,
}
