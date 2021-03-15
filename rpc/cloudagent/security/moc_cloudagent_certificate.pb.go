// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_cloudagent_certificate.proto

package security

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CertificateType int32

const (
	CertificateType_Client CertificateType = 0
	CertificateType_Server CertificateType = 1
)

var CertificateType_name = map[int32]string{
	0: "Client",
	1: "Server",
}

var CertificateType_value = map[string]int32{
	"Client": 0,
	"Server": 1,
}

func (x CertificateType) String() string {
	return proto.EnumName(CertificateType_name, int32(x))
}

func (CertificateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_373194e28b9c267a, []int{0}
}

type CertificateRequest struct {
	Certificates         []*Certificate `protobuf:"bytes,1,rep,name=Certificates,proto3" json:"Certificates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CertificateRequest) Reset()         { *m = CertificateRequest{} }
func (m *CertificateRequest) String() string { return proto.CompactTextString(m) }
func (*CertificateRequest) ProtoMessage()    {}
func (*CertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_373194e28b9c267a, []int{0}
}

func (m *CertificateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateRequest.Unmarshal(m, b)
}
func (m *CertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateRequest.Marshal(b, m, deterministic)
}
func (m *CertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateRequest.Merge(m, src)
}
func (m *CertificateRequest) XXX_Size() int {
	return xxx_messageInfo_CertificateRequest.Size(m)
}
func (m *CertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateRequest proto.InternalMessageInfo

func (m *CertificateRequest) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

type CSRRequest struct {
	CSRs                 []*CertificateSigningRequest `protobuf:"bytes,1,rep,name=CSRs,proto3" json:"CSRs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *CSRRequest) Reset()         { *m = CSRRequest{} }
func (m *CSRRequest) String() string { return proto.CompactTextString(m) }
func (*CSRRequest) ProtoMessage()    {}
func (*CSRRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_373194e28b9c267a, []int{1}
}

func (m *CSRRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CSRRequest.Unmarshal(m, b)
}
func (m *CSRRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CSRRequest.Marshal(b, m, deterministic)
}
func (m *CSRRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CSRRequest.Merge(m, src)
}
func (m *CSRRequest) XXX_Size() int {
	return xxx_messageInfo_CSRRequest.Size(m)
}
func (m *CSRRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CSRRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CSRRequest proto.InternalMessageInfo

func (m *CSRRequest) GetCSRs() []*CertificateSigningRequest {
	if m != nil {
		return m.CSRs
	}
	return nil
}

type CertificateResponse struct {
	Certificates         []*Certificate      `protobuf:"bytes,1,rep,name=Certificates,proto3" json:"Certificates,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CertificateResponse) Reset()         { *m = CertificateResponse{} }
func (m *CertificateResponse) String() string { return proto.CompactTextString(m) }
func (*CertificateResponse) ProtoMessage()    {}
func (*CertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_373194e28b9c267a, []int{2}
}

func (m *CertificateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateResponse.Unmarshal(m, b)
}
func (m *CertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateResponse.Marshal(b, m, deterministic)
}
func (m *CertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateResponse.Merge(m, src)
}
func (m *CertificateResponse) XXX_Size() int {
	return xxx_messageInfo_CertificateResponse.Size(m)
}
func (m *CertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateResponse proto.InternalMessageInfo

func (m *CertificateResponse) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

func (m *CertificateResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *CertificateResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type CertificateSigningRequest struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Csr                  string         `protobuf:"bytes,2,opt,name=csr,proto3" json:"csr,omitempty"`
	OldCertificate       string         `protobuf:"bytes,3,opt,name=oldCertificate,proto3" json:"oldCertificate,omitempty"`
	Status               *common.Status `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	CaName               string         `protobuf:"bytes,5,opt,name=caName,proto3" json:"caName,omitempty"`
	GroupName            string         `protobuf:"bytes,6,opt,name=groupName,proto3" json:"groupName,omitempty"`
	VaultName            string         `protobuf:"bytes,7,opt,name=vaultName,proto3" json:"vaultName,omitempty"`
	LocationName         string         `protobuf:"bytes,8,opt,name=locationName,proto3" json:"locationName,omitempty"`
	Identity             string         `protobuf:"bytes,9,opt,name=identity,proto3" json:"identity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CertificateSigningRequest) Reset()         { *m = CertificateSigningRequest{} }
func (m *CertificateSigningRequest) String() string { return proto.CompactTextString(m) }
func (*CertificateSigningRequest) ProtoMessage()    {}
func (*CertificateSigningRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_373194e28b9c267a, []int{3}
}

func (m *CertificateSigningRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateSigningRequest.Unmarshal(m, b)
}
func (m *CertificateSigningRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateSigningRequest.Marshal(b, m, deterministic)
}
func (m *CertificateSigningRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateSigningRequest.Merge(m, src)
}
func (m *CertificateSigningRequest) XXX_Size() int {
	return xxx_messageInfo_CertificateSigningRequest.Size(m)
}
func (m *CertificateSigningRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateSigningRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateSigningRequest proto.InternalMessageInfo

func (m *CertificateSigningRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CertificateSigningRequest) GetCsr() string {
	if m != nil {
		return m.Csr
	}
	return ""
}

func (m *CertificateSigningRequest) GetOldCertificate() string {
	if m != nil {
		return m.OldCertificate
	}
	return ""
}

func (m *CertificateSigningRequest) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CertificateSigningRequest) GetCaName() string {
	if m != nil {
		return m.CaName
	}
	return ""
}

func (m *CertificateSigningRequest) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *CertificateSigningRequest) GetVaultName() string {
	if m != nil {
		return m.VaultName
	}
	return ""
}

func (m *CertificateSigningRequest) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *CertificateSigningRequest) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

type Certificate struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string          `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	NotBefore            int64           `protobuf:"varint,3,opt,name=notBefore,proto3" json:"notBefore,omitempty"`
	NotAfter             int64           `protobuf:"varint,4,opt,name=notAfter,proto3" json:"notAfter,omitempty"`
	Certificate          []byte          `protobuf:"bytes,5,opt,name=certificate,proto3" json:"certificate,omitempty"` // Deprecated: Do not use.
	Status               *common.Status  `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Type                 CertificateType `protobuf:"varint,7,opt,name=type,proto3,enum=moc.cloudagent.security.CertificateType" json:"type,omitempty"`
	GroupName            string          `protobuf:"bytes,8,opt,name=groupName,proto3" json:"groupName,omitempty"`
	VaultName            string          `protobuf:"bytes,9,opt,name=vaultName,proto3" json:"vaultName,omitempty"`
	LocationName         string          `protobuf:"bytes,10,opt,name=locationName,proto3" json:"locationName,omitempty"`
	Tags                 *common.Tags    `protobuf:"bytes,11,opt,name=tags,proto3" json:"tags,omitempty"`
	NewCertificate       string          `protobuf:"bytes,12,opt,name=newCertificate,proto3" json:"newCertificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Certificate) Reset()         { *m = Certificate{} }
func (m *Certificate) String() string { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()    {}
func (*Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_373194e28b9c267a, []int{4}
}

func (m *Certificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificate.Unmarshal(m, b)
}
func (m *Certificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificate.Marshal(b, m, deterministic)
}
func (m *Certificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificate.Merge(m, src)
}
func (m *Certificate) XXX_Size() int {
	return xxx_messageInfo_Certificate.Size(m)
}
func (m *Certificate) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificate.DiscardUnknown(m)
}

var xxx_messageInfo_Certificate proto.InternalMessageInfo

func (m *Certificate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Certificate) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Certificate) GetNotBefore() int64 {
	if m != nil {
		return m.NotBefore
	}
	return 0
}

func (m *Certificate) GetNotAfter() int64 {
	if m != nil {
		return m.NotAfter
	}
	return 0
}

// Deprecated: Do not use.
func (m *Certificate) GetCertificate() []byte {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func (m *Certificate) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Certificate) GetType() CertificateType {
	if m != nil {
		return m.Type
	}
	return CertificateType_Client
}

func (m *Certificate) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *Certificate) GetVaultName() string {
	if m != nil {
		return m.VaultName
	}
	return ""
}

func (m *Certificate) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *Certificate) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Certificate) GetNewCertificate() string {
	if m != nil {
		return m.NewCertificate
	}
	return ""
}

func init() {
	proto.RegisterEnum("moc.cloudagent.security.CertificateType", CertificateType_name, CertificateType_value)
	proto.RegisterType((*CertificateRequest)(nil), "moc.cloudagent.security.CertificateRequest")
	proto.RegisterType((*CSRRequest)(nil), "moc.cloudagent.security.CSRRequest")
	proto.RegisterType((*CertificateResponse)(nil), "moc.cloudagent.security.CertificateResponse")
	proto.RegisterType((*CertificateSigningRequest)(nil), "moc.cloudagent.security.CertificateSigningRequest")
	proto.RegisterType((*Certificate)(nil), "moc.cloudagent.security.Certificate")
}

func init() { proto.RegisterFile("moc_cloudagent_certificate.proto", fileDescriptor_373194e28b9c267a) }

var fileDescriptor_373194e28b9c267a = []byte{
	// 641 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x6d, 0x62, 0xd7, 0xbf, 0x66, 0x12, 0xe5, 0x17, 0x2d, 0x88, 0x9a, 0x08, 0x50, 0xe4, 0x56,
	0x28, 0xfc, 0x91, 0x8d, 0xcc, 0x95, 0x4b, 0x13, 0xfe, 0x9d, 0x40, 0xda, 0x14, 0x0e, 0x80, 0xa8,
	0x1c, 0x67, 0x62, 0x2c, 0xd9, 0xbb, 0x66, 0x77, 0xdd, 0x2a, 0x9f, 0x80, 0x0f, 0xc1, 0x8d, 0x1b,
	0xdf, 0x12, 0x79, 0xed, 0x24, 0x4e, 0xd5, 0x40, 0x90, 0xe8, 0xc9, 0xde, 0xf7, 0x66, 0xde, 0xcc,
	0x3c, 0xed, 0x0e, 0x0c, 0x52, 0x1e, 0x9e, 0x85, 0x09, 0xcf, 0x67, 0x41, 0x84, 0x4c, 0x9d, 0x85,
	0x28, 0x54, 0x3c, 0x8f, 0xc3, 0x40, 0xa1, 0x9b, 0x09, 0xae, 0x38, 0x39, 0x4c, 0x79, 0xe8, 0xae,
	0x23, 0x5c, 0x89, 0x61, 0x2e, 0x62, 0xb5, 0xe8, 0xdf, 0x8b, 0x38, 0x8f, 0x12, 0xf4, 0x74, 0xd8,
	0x34, 0x9f, 0x7b, 0x17, 0x22, 0xc8, 0x32, 0x14, 0xb2, 0x4c, 0xec, 0x1f, 0x6a, 0x69, 0x9e, 0xa6,
	0x9c, 0x55, 0x9f, 0x92, 0x70, 0x3e, 0x03, 0x19, 0xaf, 0xcb, 0x50, 0xfc, 0x9a, 0xa3, 0x54, 0xe4,
	0x35, 0x74, 0x6a, 0xa8, 0xb4, 0x1b, 0x03, 0x63, 0xd8, 0xf6, 0x8f, 0xdd, 0x2d, 0xe5, 0xdd, 0xba,
	0xc4, 0x46, 0xa6, 0x73, 0x0a, 0x30, 0x9e, 0xd0, 0xa5, 0xee, 0x4b, 0x30, 0xc7, 0x13, 0xba, 0xd4,
	0xf3, 0x77, 0xd1, 0x9b, 0xc4, 0x11, 0x8b, 0x59, 0x54, 0x29, 0x50, 0x9d, 0xef, 0xfc, 0x6c, 0xc0,
	0x8d, 0x8d, 0xb6, 0x65, 0xc6, 0x99, 0xc4, 0x7f, 0xd7, 0x37, 0xf1, 0xc1, 0xa2, 0x28, 0xf3, 0x44,
	0xd9, 0xcd, 0x41, 0x63, 0xd8, 0xf6, 0xfb, 0x6e, 0xe9, 0xb0, 0xbb, 0x74, 0xd8, 0x1d, 0x71, 0x9e,
	0xbc, 0x0f, 0x92, 0x1c, 0x69, 0x15, 0x49, 0x6e, 0xc2, 0xfe, 0x0b, 0x21, 0xb8, 0xb0, 0x8d, 0x41,
	0x63, 0xd8, 0xa2, 0xe5, 0xc1, 0xf9, 0xd1, 0x84, 0xdb, 0x5b, 0xe7, 0x21, 0x04, 0x4c, 0x16, 0xa4,
	0x68, 0x37, 0x74, 0x8a, 0xfe, 0x27, 0x3d, 0x30, 0x42, 0x29, 0x74, 0xe1, 0x16, 0x2d, 0x7e, 0xc9,
	0x7d, 0xe8, 0xf2, 0x64, 0x56, 0x53, 0xa9, 0x4a, 0x5c, 0x42, 0xc9, 0x11, 0x58, 0x52, 0x05, 0x2a,
	0x97, 0xb6, 0xa9, 0xbb, 0x6e, 0xeb, 0xc9, 0x27, 0x1a, 0xa2, 0x15, 0x45, 0x6e, 0x81, 0x15, 0x06,
	0x6f, 0x8a, 0xa2, 0xfb, 0x5a, 0xa4, 0x3a, 0x91, 0x3b, 0xd0, 0x8a, 0x04, 0xcf, 0x33, 0x4d, 0x59,
	0x9a, 0x5a, 0x03, 0x05, 0x7b, 0x1e, 0xe4, 0x89, 0xd2, 0xec, 0x7f, 0x25, 0xbb, 0x02, 0x88, 0x03,
	0x9d, 0x84, 0x87, 0x81, 0x8a, 0x39, 0xd3, 0x01, 0x07, 0x3a, 0x60, 0x03, 0x23, 0x7d, 0x38, 0x88,
	0x67, 0xc8, 0x54, 0xac, 0x16, 0x76, 0x4b, 0xf3, 0xab, 0xb3, 0xf3, 0xdd, 0x80, 0x76, 0x7d, 0x90,
	0xab, 0x6c, 0xe9, 0x42, 0x33, 0x9e, 0x55, 0xae, 0x34, 0xe3, 0x59, 0xd1, 0x11, 0xe3, 0x6a, 0x84,
	0x73, 0x2e, 0x4a, 0x3f, 0x0c, 0xba, 0x06, 0x8a, 0x6a, 0x8c, 0xab, 0x93, 0xb9, 0x42, 0xa1, 0xcd,
	0x30, 0xe8, 0xea, 0x4c, 0x8e, 0xa1, 0x5d, 0x7b, 0x5b, 0xda, 0x86, 0xce, 0xa8, 0x69, 0x37, 0x68,
	0x1d, 0xae, 0x99, 0x69, 0x6d, 0x37, 0xf3, 0x19, 0x98, 0x6a, 0x91, 0x95, 0x8e, 0x74, 0xfd, 0xe1,
	0x2e, 0x37, 0xed, 0x74, 0x91, 0x21, 0xd5, 0x59, 0x9b, 0x96, 0x1f, 0xfc, 0xd6, 0xf2, 0xd6, 0x9f,
	0x2c, 0x87, 0x2b, 0x2c, 0xbf, 0x0b, 0xa6, 0x0a, 0x22, 0x69, 0xb7, 0xf5, 0x00, 0x2d, 0xdd, 0xdd,
	0x69, 0x10, 0x49, 0xaa, 0xe1, 0xe2, 0x5a, 0x31, 0xbc, 0xa8, 0x5f, 0xab, 0x4e, 0x79, 0xad, 0x36,
	0xd1, 0x87, 0x0f, 0xe0, 0xff, 0x4b, 0xfd, 0x13, 0x00, 0x6b, 0x9c, 0xc4, 0xc8, 0x54, 0x6f, 0xaf,
	0xf8, 0x9f, 0xa0, 0x38, 0x47, 0xd1, 0x6b, 0xf8, 0xdf, 0x4c, 0xe8, 0xd5, 0x62, 0x4f, 0x0a, 0x17,
	0x48, 0x0a, 0xdd, 0xb1, 0xc0, 0x40, 0xe1, 0x5b, 0xf1, 0x2e, 0x9b, 0x15, 0xde, 0x3e, 0xda, 0xe9,
	0x49, 0x96, 0x6f, 0xa4, 0xff, 0x78, 0xb7, 0xe0, 0x72, 0x07, 0x38, 0x7b, 0x64, 0x0a, 0xc6, 0x2b,
	0x54, 0xd7, 0x5b, 0x03, 0xc1, 0x7a, 0x8e, 0x09, 0x5e, 0xf7, 0x28, 0x1f, 0xc1, 0x2c, 0x16, 0x06,
	0x39, 0xda, 0x9e, 0xb7, 0xda, 0xae, 0x7f, 0x2d, 0xfe, 0x09, 0xf6, 0x29, 0x32, 0xbc, 0xb8, 0x16,
	0xf5, 0x91, 0xff, 0xe1, 0x49, 0x14, 0xab, 0x2f, 0xf9, 0xd4, 0x0d, 0x79, 0xea, 0xa5, 0x71, 0x28,
	0xb8, 0xe4, 0x73, 0xe5, 0xa5, 0x3c, 0xf4, 0x44, 0x16, 0x7a, 0x6b, 0x25, 0x6f, 0xa9, 0x34, 0xb5,
	0xf4, 0x76, 0x7d, 0xfa, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x15, 0xd9, 0x7c, 0x67, 0x0a, 0x07, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CertificateAgentClient is the client API for CertificateAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CertificateAgentClient interface {
	CreateOrUpdate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
	Get(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
	Delete(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
	Sign(ctx context.Context, in *CSRRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
	Renew(ctx context.Context, in *CSRRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
}

type certificateAgentClient struct {
	cc *grpc.ClientConn
}

func NewCertificateAgentClient(cc *grpc.ClientConn) CertificateAgentClient {
	return &certificateAgentClient{cc}
}

func (c *certificateAgentClient) CreateOrUpdate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.security.CertificateAgent/CreateOrUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAgentClient) Get(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.security.CertificateAgent/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAgentClient) Delete(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.security.CertificateAgent/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAgentClient) Sign(ctx context.Context, in *CSRRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.security.CertificateAgent/Sign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAgentClient) Renew(ctx context.Context, in *CSRRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.security.CertificateAgent/Renew", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateAgentServer is the server API for CertificateAgent service.
type CertificateAgentServer interface {
	CreateOrUpdate(context.Context, *CertificateRequest) (*CertificateResponse, error)
	Get(context.Context, *CertificateRequest) (*CertificateResponse, error)
	Delete(context.Context, *CertificateRequest) (*CertificateResponse, error)
	Sign(context.Context, *CSRRequest) (*CertificateResponse, error)
	Renew(context.Context, *CSRRequest) (*CertificateResponse, error)
}

// UnimplementedCertificateAgentServer can be embedded to have forward compatible implementations.
type UnimplementedCertificateAgentServer struct {
}

func (*UnimplementedCertificateAgentServer) CreateOrUpdate(ctx context.Context, req *CertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdate not implemented")
}
func (*UnimplementedCertificateAgentServer) Get(ctx context.Context, req *CertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCertificateAgentServer) Delete(ctx context.Context, req *CertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedCertificateAgentServer) Sign(ctx context.Context, req *CSRRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}
func (*UnimplementedCertificateAgentServer) Renew(ctx context.Context, req *CSRRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Renew not implemented")
}

func RegisterCertificateAgentServer(s *grpc.Server, srv CertificateAgentServer) {
	s.RegisterService(&_CertificateAgent_serviceDesc, srv)
}

func _CertificateAgent_CreateOrUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).CreateOrUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.security.CertificateAgent/CreateOrUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).CreateOrUpdate(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAgent_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.security.CertificateAgent/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).Get(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAgent_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.security.CertificateAgent/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).Delete(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAgent_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.security.CertificateAgent/Sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).Sign(ctx, req.(*CSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAgent_Renew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).Renew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.security.CertificateAgent/Renew",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).Renew(ctx, req.(*CSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.security.CertificateAgent",
	HandlerType: (*CertificateAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrUpdate",
			Handler:    _CertificateAgent_CreateOrUpdate_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CertificateAgent_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CertificateAgent_Delete_Handler,
		},
		{
			MethodName: "Sign",
			Handler:    _CertificateAgent_Sign_Handler,
		},
		{
			MethodName: "Renew",
			Handler:    _CertificateAgent_Renew_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_cloudagent_certificate.proto",
}