// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_mocguestagent_certificate.proto

package security

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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
	return fileDescriptor_3663d8a2d284fe24, []int{0}
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

type CertificateResponse struct {
	Result               *wrappers.BoolValue `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CertificateResponse) Reset()         { *m = CertificateResponse{} }
func (m *CertificateResponse) String() string { return proto.CompactTextString(m) }
func (*CertificateResponse) ProtoMessage()    {}
func (*CertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3663d8a2d284fe24, []int{1}
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

type Certificate struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string         `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	NotBefore            int64          `protobuf:"varint,3,opt,name=notBefore,proto3" json:"notBefore,omitempty"`
	NotAfter             int64          `protobuf:"varint,4,opt,name=notAfter,proto3" json:"notAfter,omitempty"`
	Certificate          string         `protobuf:"bytes,5,opt,name=certificate,proto3" json:"certificate,omitempty"`
	Status               *common.Status `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Entity               *common.Entity `protobuf:"bytes,7,opt,name=entity,proto3" json:"entity,omitempty"`
	Tags                 *common.Tags   `protobuf:"bytes,8,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Certificate) Reset()         { *m = Certificate{} }
func (m *Certificate) String() string { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()    {}
func (*Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_3663d8a2d284fe24, []int{2}
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

func (m *Certificate) GetCertificate() string {
	if m != nil {
		return m.Certificate
	}
	return ""
}

func (m *Certificate) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Certificate) GetEntity() *common.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *Certificate) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
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
	return fileDescriptor_3663d8a2d284fe24, []int{3}
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

type CertificateSigningRequest struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Csr                  string         `protobuf:"bytes,2,opt,name=csr,proto3" json:"csr,omitempty"`
	OldCertificate       string         `protobuf:"bytes,3,opt,name=oldCertificate,proto3" json:"oldCertificate,omitempty"`
	Status               *common.Status `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	CaName               string         `protobuf:"bytes,5,opt,name=caName,proto3" json:"caName,omitempty"`
	Validity             int64          `protobuf:"varint,6,opt,name=validity,proto3" json:"validity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CertificateSigningRequest) Reset()         { *m = CertificateSigningRequest{} }
func (m *CertificateSigningRequest) String() string { return proto.CompactTextString(m) }
func (*CertificateSigningRequest) ProtoMessage()    {}
func (*CertificateSigningRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3663d8a2d284fe24, []int{4}
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

func (m *CertificateSigningRequest) GetValidity() int64 {
	if m != nil {
		return m.Validity
	}
	return 0
}

type RenewCSRResponse struct {
	Csr                  *CertificateSigningRequest `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
	Result               *wrappers.BoolValue        `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string                     `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *RenewCSRResponse) Reset()         { *m = RenewCSRResponse{} }
func (m *RenewCSRResponse) String() string { return proto.CompactTextString(m) }
func (*RenewCSRResponse) ProtoMessage()    {}
func (*RenewCSRResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3663d8a2d284fe24, []int{5}
}

func (m *RenewCSRResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenewCSRResponse.Unmarshal(m, b)
}
func (m *RenewCSRResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenewCSRResponse.Marshal(b, m, deterministic)
}
func (m *RenewCSRResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenewCSRResponse.Merge(m, src)
}
func (m *RenewCSRResponse) XXX_Size() int {
	return xxx_messageInfo_RenewCSRResponse.Size(m)
}
func (m *RenewCSRResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RenewCSRResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RenewCSRResponse proto.InternalMessageInfo

func (m *RenewCSRResponse) GetCsr() *CertificateSigningRequest {
	if m != nil {
		return m.Csr
	}
	return nil
}

func (m *RenewCSRResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *RenewCSRResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*CertificateRequest)(nil), "moc.mocguestagent.security.CertificateRequest")
	proto.RegisterType((*CertificateResponse)(nil), "moc.mocguestagent.security.CertificateResponse")
	proto.RegisterType((*Certificate)(nil), "moc.mocguestagent.security.Certificate")
	proto.RegisterType((*CSRRequest)(nil), "moc.mocguestagent.security.CSRRequest")
	proto.RegisterType((*CertificateSigningRequest)(nil), "moc.mocguestagent.security.CertificateSigningRequest")
	proto.RegisterType((*RenewCSRResponse)(nil), "moc.mocguestagent.security.RenewCSRResponse")
}

func init() {
	proto.RegisterFile("moc_mocguestagent_certificate.proto", fileDescriptor_3663d8a2d284fe24)
}

var fileDescriptor_3663d8a2d284fe24 = []byte{
	// 580 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xfd, 0x9c, 0xb8, 0xfe, 0x9a, 0x31, 0xaa, 0xca, 0x82, 0x82, 0x31, 0x3f, 0x8a, 0x5c, 0x09,
	0x72, 0x51, 0xd9, 0x52, 0xaa, 0x3e, 0x40, 0x13, 0x45, 0x15, 0x42, 0xe2, 0x62, 0xc3, 0x8f, 0xc4,
	0x4d, 0xb4, 0xd9, 0x6c, 0x8c, 0x25, 0xdb, 0x6b, 0x76, 0xd7, 0xad, 0xf2, 0x06, 0xbc, 0x08, 0x8f,
	0xc0, 0x6b, 0xf0, 0x0e, 0xbc, 0x07, 0x17, 0xc8, 0x6b, 0x9b, 0x6c, 0x02, 0x41, 0x81, 0xab, 0x64,
	0x66, 0xce, 0x9e, 0x99, 0x39, 0x67, 0x0c, 0x67, 0x19, 0xa7, 0xf3, 0x8c, 0xd3, 0xb8, 0x64, 0x52,
	0x91, 0x98, 0xe5, 0x6a, 0x4e, 0x99, 0x50, 0xc9, 0x2a, 0xa1, 0x44, 0xb1, 0xb0, 0x10, 0x5c, 0x71,
	0xe4, 0x67, 0x9c, 0x86, 0x5b, 0xa0, 0x50, 0x32, 0x5a, 0x8a, 0x44, 0xad, 0xfd, 0xa7, 0x31, 0xe7,
	0x71, 0xca, 0x22, 0x8d, 0x5c, 0x94, 0xab, 0xe8, 0x56, 0x90, 0xa2, 0x60, 0x42, 0xd6, 0x6f, 0xfd,
	0x47, 0xbb, 0x75, 0x96, 0x15, 0x6a, 0xdd, 0x14, 0x1f, 0x54, 0xdd, 0x29, 0xcf, 0x32, 0x9e, 0x37,
	0x3f, 0x75, 0x21, 0x20, 0x80, 0x26, 0x9b, 0x31, 0x30, 0xfb, 0x58, 0x75, 0x46, 0x2f, 0xe1, 0x8e,
	0x91, 0x95, 0x9e, 0x35, 0xe8, 0x0e, 0xdd, 0xd1, 0xf3, 0x70, 0xff, 0x78, 0xa1, 0xc9, 0xb2, 0xf5,
	0x38, 0x98, 0xc3, 0xbd, 0xad, 0x16, 0xb2, 0xe0, 0xb9, 0x64, 0x68, 0x04, 0x0e, 0x66, 0xb2, 0x4c,
	0x95, 0x67, 0x0d, 0xac, 0xa1, 0x3b, 0xf2, 0xc3, 0x7a, 0x81, 0xb0, 0x5d, 0x20, 0x1c, 0x73, 0x9e,
	0xbe, 0x25, 0x69, 0xc9, 0x70, 0x83, 0x44, 0xf7, 0xe1, 0x68, 0x2a, 0x04, 0x17, 0x5e, 0x67, 0x60,
	0x0d, 0x7b, 0xb8, 0x0e, 0x82, 0xef, 0x16, 0xb8, 0x46, 0x07, 0x84, 0xc0, 0xce, 0x49, 0xc6, 0x34,
	0x6f, 0x0f, 0xeb, 0xff, 0xe8, 0x04, 0x3a, 0xc9, 0xb2, 0x79, 0xd6, 0x49, 0x96, 0xe8, 0x31, 0xf4,
	0x72, 0xae, 0xc6, 0x6c, 0xc5, 0x05, 0xf3, 0xba, 0x03, 0x6b, 0xd8, 0xc5, 0x9b, 0x04, 0xf2, 0xe1,
	0x38, 0xe7, 0xea, 0x6a, 0xa5, 0x98, 0xf0, 0x6c, 0x5d, 0xfc, 0x19, 0xa3, 0x67, 0xe0, 0x1a, 0xc6,
	0x79, 0x47, 0x15, 0xe5, 0xd8, 0xfe, 0xf4, 0xc5, 0xb3, 0xb0, 0x59, 0x40, 0x67, 0xe0, 0x48, 0x45,
	0x54, 0x29, 0x3d, 0x47, 0xef, 0xe7, 0x6a, 0xf5, 0x66, 0x3a, 0x85, 0x9b, 0x52, 0x05, 0x62, 0xb9,
	0x4a, 0xd4, 0xda, 0xfb, 0xdf, 0x00, 0x4d, 0x75, 0x0a, 0x37, 0x25, 0xf4, 0x04, 0x6c, 0x45, 0x62,
	0xe9, 0x1d, 0x6b, 0x48, 0x4f, 0x43, 0x5e, 0x93, 0x58, 0x62, 0x9d, 0x0e, 0xde, 0x01, 0x4c, 0x66,
	0xb8, 0xb5, 0xee, 0x05, 0xd8, 0x93, 0x19, 0x6e, 0x2d, 0xbb, 0x3c, 0xd0, 0xb2, 0x59, 0x12, 0xe7,
	0x49, 0x1e, 0x37, 0x24, 0x58, 0x53, 0x04, 0x5f, 0x2d, 0x78, 0xb8, 0x17, 0xf3, 0x5b, 0x95, 0xfb,
	0xd0, 0xa5, 0xb2, 0x71, 0xa7, 0xd1, 0xa4, 0x4a, 0xa0, 0x73, 0x38, 0xe1, 0xe9, 0xd2, 0xe0, 0xd2,
	0x92, 0xb7, 0x90, 0x9d, 0x9a, 0xa1, 0x9c, 0xbd, 0x5f, 0xb9, 0x3e, 0x38, 0x94, 0xbc, 0xaa, 0x06,
	0xd0, 0x0e, 0xe0, 0x26, 0xaa, 0xac, 0xbb, 0x21, 0x69, 0xb2, 0xac, 0x34, 0x75, 0x6a, 0xeb, 0xda,
	0x38, 0xf8, 0x6c, 0xc1, 0x29, 0x66, 0x39, 0xbb, 0xd5, 0x7a, 0x35, 0x77, 0x78, 0x5d, 0xcf, 0x5c,
	0x1f, 0xe1, 0x3f, 0xea, 0xa5, 0x97, 0xdc, 0x1c, 0x74, 0xe7, 0xef, 0x0f, 0xba, 0x6b, 0x1c, 0xf4,
	0xe8, 0x9b, 0x05, 0xa7, 0x46, 0xb3, 0xab, 0x6a, 0x0c, 0xf4, 0x06, 0xdc, 0x6b, 0xa6, 0xda, 0xf1,
	0x51, 0xff, 0x17, 0xf6, 0x69, 0xf5, 0xbd, 0xfb, 0xe7, 0x7f, 0xda, 0x60, 0x77, 0xf9, 0xe0, 0x3f,
	0x74, 0x03, 0x77, 0x31, 0x57, 0x44, 0x31, 0xd3, 0x81, 0xf0, 0xd0, 0x2f, 0xbd, 0xde, 0xdf, 0x8f,
	0x0e, 0xc6, 0xb7, 0x7d, 0xc7, 0x97, 0xef, 0x2f, 0xe2, 0x44, 0x7d, 0x28, 0x17, 0x21, 0xe5, 0x59,
	0x94, 0x25, 0x54, 0x70, 0xc9, 0x57, 0x2a, 0xca, 0x38, 0x8d, 0x44, 0x41, 0xa3, 0x2d, 0xb2, 0xa8,
	0x25, 0x5b, 0x38, 0x7a, 0xdd, 0x8b, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x04, 0xe4, 0xb3,
	0x4f, 0x05, 0x00, 0x00,
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
	GetRenewCSR(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*RenewCSRResponse, error)
	RotateCertificate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
}

type certificateAgentClient struct {
	cc *grpc.ClientConn
}

func NewCertificateAgentClient(cc *grpc.ClientConn) CertificateAgentClient {
	return &certificateAgentClient{cc}
}

func (c *certificateAgentClient) GetRenewCSR(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*RenewCSRResponse, error) {
	out := new(RenewCSRResponse)
	err := c.cc.Invoke(ctx, "/moc.mocguestagent.security.CertificateAgent/GetRenewCSR", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAgentClient) RotateCertificate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.mocguestagent.security.CertificateAgent/RotateCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateAgentServer is the server API for CertificateAgent service.
type CertificateAgentServer interface {
	GetRenewCSR(context.Context, *empty.Empty) (*RenewCSRResponse, error)
	RotateCertificate(context.Context, *CertificateRequest) (*CertificateResponse, error)
}

// UnimplementedCertificateAgentServer can be embedded to have forward compatible implementations.
type UnimplementedCertificateAgentServer struct {
}

func (*UnimplementedCertificateAgentServer) GetRenewCSR(ctx context.Context, req *empty.Empty) (*RenewCSRResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRenewCSR not implemented")
}
func (*UnimplementedCertificateAgentServer) RotateCertificate(ctx context.Context, req *CertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RotateCertificate not implemented")
}

func RegisterCertificateAgentServer(s *grpc.Server, srv CertificateAgentServer) {
	s.RegisterService(&_CertificateAgent_serviceDesc, srv)
}

func _CertificateAgent_GetRenewCSR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).GetRenewCSR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.mocguestagent.security.CertificateAgent/GetRenewCSR",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).GetRenewCSR(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAgent_RotateCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).RotateCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.mocguestagent.security.CertificateAgent/RotateCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).RotateCertificate(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.mocguestagent.security.CertificateAgent",
	HandlerType: (*CertificateAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRenewCSR",
			Handler:    _CertificateAgent_GetRenewCSR_Handler,
		},
		{
			MethodName: "RotateCertificate",
			Handler:    _CertificateAgent_RotateCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_mocguestagent_certificate.proto",
}