// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orc8r/protos/bootstrapper.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ChallengeKey_KeyType int32

const (
	ChallengeKey_ECHO                  ChallengeKey_KeyType = 0
	ChallengeKey_SOFTWARE_RSA_SHA256   ChallengeKey_KeyType = 1
	ChallengeKey_SOFTWARE_ECDSA_SHA256 ChallengeKey_KeyType = 2
)

var ChallengeKey_KeyType_name = map[int32]string{
	0: "ECHO",
	1: "SOFTWARE_RSA_SHA256",
	2: "SOFTWARE_ECDSA_SHA256",
}

var ChallengeKey_KeyType_value = map[string]int32{
	"ECHO":                  0,
	"SOFTWARE_RSA_SHA256":   1,
	"SOFTWARE_ECDSA_SHA256": 2,
}

func (x ChallengeKey_KeyType) String() string {
	return proto.EnumName(ChallengeKey_KeyType_name, int32(x))
}

func (ChallengeKey_KeyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b592b3c4e9ae6813, []int{1, 0}
}

type Challenge struct {
	KeyType              ChallengeKey_KeyType `protobuf:"varint,1,opt,name=key_type,json=keyType,proto3,enum=magma.orc8r.ChallengeKey_KeyType" json:"key_type,omitempty"`
	Challenge            []byte               `protobuf:"bytes,2,opt,name=challenge,proto3" json:"challenge,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Challenge) Reset()         { *m = Challenge{} }
func (m *Challenge) String() string { return proto.CompactTextString(m) }
func (*Challenge) ProtoMessage()    {}
func (*Challenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_b592b3c4e9ae6813, []int{0}
}

func (m *Challenge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Challenge.Unmarshal(m, b)
}
func (m *Challenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Challenge.Marshal(b, m, deterministic)
}
func (m *Challenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Challenge.Merge(m, src)
}
func (m *Challenge) XXX_Size() int {
	return xxx_messageInfo_Challenge.Size(m)
}
func (m *Challenge) XXX_DiscardUnknown() {
	xxx_messageInfo_Challenge.DiscardUnknown(m)
}

var xxx_messageInfo_Challenge proto.InternalMessageInfo

func (m *Challenge) GetKeyType() ChallengeKey_KeyType {
	if m != nil {
		return m.KeyType
	}
	return ChallengeKey_ECHO
}

func (m *Challenge) GetChallenge() []byte {
	if m != nil {
		return m.Challenge
	}
	return nil
}

// --------------------------------------------------------------------------
// Challenge key stores the key used for challenge-response during bootstrap.
// --------------------------------------------------------------------------
type ChallengeKey struct {
	KeyType ChallengeKey_KeyType `protobuf:"varint,1,opt,name=key_type,json=keyType,proto3,enum=magma.orc8r.ChallengeKey_KeyType" json:"key_type,omitempty"`
	// Public key encoded in DER format
	Key                  []byte   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChallengeKey) Reset()         { *m = ChallengeKey{} }
func (m *ChallengeKey) String() string { return proto.CompactTextString(m) }
func (*ChallengeKey) ProtoMessage()    {}
func (*ChallengeKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_b592b3c4e9ae6813, []int{1}
}

func (m *ChallengeKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChallengeKey.Unmarshal(m, b)
}
func (m *ChallengeKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChallengeKey.Marshal(b, m, deterministic)
}
func (m *ChallengeKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChallengeKey.Merge(m, src)
}
func (m *ChallengeKey) XXX_Size() int {
	return xxx_messageInfo_ChallengeKey.Size(m)
}
func (m *ChallengeKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ChallengeKey.DiscardUnknown(m)
}

var xxx_messageInfo_ChallengeKey proto.InternalMessageInfo

func (m *ChallengeKey) GetKeyType() ChallengeKey_KeyType {
	if m != nil {
		return m.KeyType
	}
	return ChallengeKey_ECHO
}

func (m *ChallengeKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type Response struct {
	HwId      *AccessGatewayID `protobuf:"bytes,1,opt,name=hw_id,json=hwId,proto3" json:"hw_id,omitempty"`
	Challenge []byte           `protobuf:"bytes,2,opt,name=challenge,proto3" json:"challenge,omitempty"`
	// Types that are valid to be assigned to Response:
	//	*Response_EchoResponse
	//	*Response_RsaResponse
	//	*Response_EcdsaResponse
	Response             isResponse_Response `protobuf_oneof:"response"`
	Csr                  *CSR                `protobuf:"bytes,6,opt,name=csr,proto3" json:"csr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_b592b3c4e9ae6813, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetHwId() *AccessGatewayID {
	if m != nil {
		return m.HwId
	}
	return nil
}

func (m *Response) GetChallenge() []byte {
	if m != nil {
		return m.Challenge
	}
	return nil
}

type isResponse_Response interface {
	isResponse_Response()
}

type Response_EchoResponse struct {
	EchoResponse *Response_Echo `protobuf:"bytes,3,opt,name=echo_response,json=echoResponse,proto3,oneof"`
}

type Response_RsaResponse struct {
	RsaResponse *Response_RSA `protobuf:"bytes,4,opt,name=rsa_response,json=rsaResponse,proto3,oneof"`
}

type Response_EcdsaResponse struct {
	EcdsaResponse *Response_ECDSA `protobuf:"bytes,5,opt,name=ecdsa_response,json=ecdsaResponse,proto3,oneof"`
}

func (*Response_EchoResponse) isResponse_Response() {}

func (*Response_RsaResponse) isResponse_Response() {}

func (*Response_EcdsaResponse) isResponse_Response() {}

func (m *Response) GetResponse() isResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *Response) GetEchoResponse() *Response_Echo {
	if x, ok := m.GetResponse().(*Response_EchoResponse); ok {
		return x.EchoResponse
	}
	return nil
}

func (m *Response) GetRsaResponse() *Response_RSA {
	if x, ok := m.GetResponse().(*Response_RsaResponse); ok {
		return x.RsaResponse
	}
	return nil
}

func (m *Response) GetEcdsaResponse() *Response_ECDSA {
	if x, ok := m.GetResponse().(*Response_EcdsaResponse); ok {
		return x.EcdsaResponse
	}
	return nil
}

func (m *Response) GetCsr() *CSR {
	if m != nil {
		return m.Csr
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Response) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Response_EchoResponse)(nil),
		(*Response_RsaResponse)(nil),
		(*Response_EcdsaResponse)(nil),
	}
}

type Response_Echo struct {
	Response             []byte   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response_Echo) Reset()         { *m = Response_Echo{} }
func (m *Response_Echo) String() string { return proto.CompactTextString(m) }
func (*Response_Echo) ProtoMessage()    {}
func (*Response_Echo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b592b3c4e9ae6813, []int{2, 0}
}

func (m *Response_Echo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_Echo.Unmarshal(m, b)
}
func (m *Response_Echo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_Echo.Marshal(b, m, deterministic)
}
func (m *Response_Echo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_Echo.Merge(m, src)
}
func (m *Response_Echo) XXX_Size() int {
	return xxx_messageInfo_Response_Echo.Size(m)
}
func (m *Response_Echo) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_Echo.DiscardUnknown(m)
}

var xxx_messageInfo_Response_Echo proto.InternalMessageInfo

func (m *Response_Echo) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

type Response_RSA struct {
	Signature            []byte   `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response_RSA) Reset()         { *m = Response_RSA{} }
func (m *Response_RSA) String() string { return proto.CompactTextString(m) }
func (*Response_RSA) ProtoMessage()    {}
func (*Response_RSA) Descriptor() ([]byte, []int) {
	return fileDescriptor_b592b3c4e9ae6813, []int{2, 1}
}

func (m *Response_RSA) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_RSA.Unmarshal(m, b)
}
func (m *Response_RSA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_RSA.Marshal(b, m, deterministic)
}
func (m *Response_RSA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_RSA.Merge(m, src)
}
func (m *Response_RSA) XXX_Size() int {
	return xxx_messageInfo_Response_RSA.Size(m)
}
func (m *Response_RSA) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_RSA.DiscardUnknown(m)
}

var xxx_messageInfo_Response_RSA proto.InternalMessageInfo

func (m *Response_RSA) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Response_ECDSA struct {
	R                    []byte   `protobuf:"bytes,1,opt,name=r,proto3" json:"r,omitempty"`
	S                    []byte   `protobuf:"bytes,2,opt,name=s,proto3" json:"s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response_ECDSA) Reset()         { *m = Response_ECDSA{} }
func (m *Response_ECDSA) String() string { return proto.CompactTextString(m) }
func (*Response_ECDSA) ProtoMessage()    {}
func (*Response_ECDSA) Descriptor() ([]byte, []int) {
	return fileDescriptor_b592b3c4e9ae6813, []int{2, 2}
}

func (m *Response_ECDSA) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_ECDSA.Unmarshal(m, b)
}
func (m *Response_ECDSA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_ECDSA.Marshal(b, m, deterministic)
}
func (m *Response_ECDSA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_ECDSA.Merge(m, src)
}
func (m *Response_ECDSA) XXX_Size() int {
	return xxx_messageInfo_Response_ECDSA.Size(m)
}
func (m *Response_ECDSA) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_ECDSA.DiscardUnknown(m)
}

var xxx_messageInfo_Response_ECDSA proto.InternalMessageInfo

func (m *Response_ECDSA) GetR() []byte {
	if m != nil {
		return m.R
	}
	return nil
}

func (m *Response_ECDSA) GetS() []byte {
	if m != nil {
		return m.S
	}
	return nil
}

func init() {
	proto.RegisterEnum("magma.orc8r.ChallengeKey_KeyType", ChallengeKey_KeyType_name, ChallengeKey_KeyType_value)
	proto.RegisterType((*Challenge)(nil), "magma.orc8r.Challenge")
	proto.RegisterType((*ChallengeKey)(nil), "magma.orc8r.ChallengeKey")
	proto.RegisterType((*Response)(nil), "magma.orc8r.Response")
	proto.RegisterType((*Response_Echo)(nil), "magma.orc8r.Response.Echo")
	proto.RegisterType((*Response_RSA)(nil), "magma.orc8r.Response.RSA")
	proto.RegisterType((*Response_ECDSA)(nil), "magma.orc8r.Response.ECDSA")
}

func init() { proto.RegisterFile("orc8r/protos/bootstrapper.proto", fileDescriptor_b592b3c4e9ae6813) }

var fileDescriptor_b592b3c4e9ae6813 = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x5d, 0x6f, 0xda, 0x30,
	0x14, 0x4d, 0x0a, 0xb4, 0xf4, 0x92, 0x56, 0xc8, 0x53, 0x37, 0x08, 0x48, 0xeb, 0xd2, 0x97, 0x3e,
	0x05, 0x8d, 0x69, 0xd3, 0x1e, 0xa6, 0x69, 0xe1, 0xa3, 0x50, 0xf5, 0xa1, 0x92, 0x53, 0x69, 0xd2,
	0x5e, 0x50, 0x30, 0x77, 0x21, 0x82, 0x26, 0x99, 0xed, 0x0a, 0xe5, 0x9f, 0xec, 0x1f, 0xec, 0x7f,
	0xec, 0x97, 0x4d, 0x31, 0x21, 0x21, 0x12, 0xeb, 0x4b, 0x9f, 0x62, 0xfb, 0x9c, 0x7b, 0xce, 0xcd,
	0xb1, 0x2f, 0xbc, 0x8d, 0x38, 0xfb, 0xcc, 0x7b, 0x31, 0x8f, 0x64, 0x24, 0x7a, 0xf3, 0x28, 0x92,
	0x42, 0x72, 0x2f, 0x8e, 0x91, 0xdb, 0xea, 0x8c, 0x34, 0x1e, 0x3d, 0xff, 0xd1, 0xb3, 0x15, 0xcd,
	0xec, 0x96, 0xd8, 0x0c, 0xb9, 0x0c, 0x7e, 0x06, 0x3b, 0xaa, 0xd9, 0x29, 0xa1, 0xc1, 0x02, 0x43,
	0x19, 0xc8, 0x64, 0x0b, 0x5a, 0x3e, 0x9c, 0x0e, 0x97, 0xde, 0x7a, 0x8d, 0xa1, 0x8f, 0xe4, 0x0b,
	0xd4, 0x57, 0x98, 0xcc, 0x64, 0x12, 0x63, 0x4b, 0xbf, 0xd4, 0xaf, 0xcf, 0xfb, 0xef, 0xec, 0x3d,
	0x1f, 0x3b, 0x67, 0xde, 0x61, 0x62, 0xdf, 0x61, 0xf2, 0x90, 0xc4, 0x48, 0x4f, 0x56, 0xdb, 0x05,
	0xe9, 0xc2, 0x29, 0xdb, 0x11, 0x5a, 0x47, 0x97, 0xfa, 0xb5, 0x41, 0x8b, 0x03, 0xeb, 0x8f, 0x0e,
	0xc6, 0x7e, 0xfd, 0x0b, 0xcd, 0x9a, 0x50, 0x59, 0x61, 0x92, 0xd9, 0xa4, 0x4b, 0x6b, 0x02, 0x27,
	0x19, 0x8b, 0xd4, 0xa1, 0x3a, 0x1e, 0x4e, 0xef, 0x9b, 0x1a, 0x79, 0x03, 0xaf, 0xdc, 0xfb, 0x9b,
	0x87, 0xef, 0x0e, 0x1d, 0xcf, 0xa8, 0xeb, 0xcc, 0xdc, 0xa9, 0xd3, 0xff, 0xf8, 0xa9, 0xa9, 0x93,
	0x36, 0x5c, 0xe4, 0xc0, 0x78, 0x38, 0x2a, 0xa0, 0x23, 0xeb, 0x6f, 0x05, 0xea, 0x14, 0x45, 0x1c,
	0x85, 0x02, 0xc9, 0x7b, 0xa8, 0x2d, 0x37, 0xb3, 0x60, 0xa1, 0x5a, 0x6c, 0xf4, 0xbb, 0xa5, 0x16,
	0x1d, 0xc6, 0x50, 0x88, 0x89, 0x27, 0x71, 0xe3, 0x25, 0xb7, 0x23, 0x5a, 0x5d, 0x6e, 0x6e, 0x17,
	0xcf, 0xe7, 0x40, 0x1c, 0x38, 0x43, 0xb6, 0x8c, 0x66, 0x3c, 0x73, 0x68, 0x55, 0x94, 0xb0, 0x59,
	0x12, 0xde, 0xd9, 0xdb, 0x63, 0xb6, 0x8c, 0xa6, 0x1a, 0x35, 0xd2, 0x92, 0xbc, 0xa7, 0xaf, 0x60,
	0x70, 0xe1, 0x15, 0x0a, 0x55, 0xa5, 0xd0, 0x3e, 0xac, 0x40, 0x5d, 0x67, 0xaa, 0xd1, 0x06, 0x17,
	0x5e, 0x5e, 0x3f, 0x82, 0x73, 0x64, 0x8b, 0x7d, 0x85, 0x9a, 0x52, 0xe8, 0xfc, 0xa7, 0x87, 0x34,
	0x9e, 0xa9, 0x46, 0xcf, 0x54, 0x51, 0xae, 0x62, 0x41, 0x85, 0x09, 0xde, 0x3a, 0x56, 0xa5, 0xcd,
	0xf2, 0xd5, 0xb9, 0x94, 0xa6, 0xa0, 0x69, 0x41, 0x35, 0xfd, 0x03, 0x62, 0x42, 0x3d, 0xf7, 0xd2,
	0x55, 0x22, 0xf9, 0xde, 0xbc, 0x82, 0x0a, 0x75, 0x9d, 0x34, 0x35, 0x11, 0xf8, 0xa1, 0x27, 0x9f,
	0xf8, 0x8e, 0x53, 0x1c, 0x98, 0x57, 0x50, 0x53, 0x6d, 0x10, 0x03, 0x74, 0x9e, 0xc1, 0x3a, 0x4f,
	0x77, 0x22, 0x8b, 0x58, 0x17, 0x03, 0x28, 0x5c, 0xfa, 0xbf, 0x75, 0x30, 0x06, 0x7b, 0x63, 0x43,
	0x6e, 0xc0, 0x98, 0xa0, 0x2c, 0xde, 0xfa, 0xb3, 0x37, 0x69, 0xbe, 0x3e, 0xfc, 0x14, 0x2d, 0x8d,
	0x7c, 0x83, 0x06, 0xc5, 0x5f, 0x4f, 0x28, 0xa4, 0x1b, 0xf8, 0x21, 0xb9, 0x38, 0x98, 0x99, 0xd9,
	0x2a, 0xd7, 0x6f, 0x27, 0x92, 0x79, 0x12, 0x2d, 0x6d, 0xd0, 0xf9, 0xd1, 0x56, 0x60, 0x6f, 0x3b,
	0x97, 0xeb, 0x60, 0xde, 0xf3, 0xa3, 0x6c, 0x3c, 0xe7, 0xc7, 0xea, 0xfb, 0xe1, 0x5f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x66, 0x62, 0xe2, 0x84, 0x01, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BootstrapperClient is the client API for Bootstrapper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BootstrapperClient interface {
	// get the challange for gateway specified in hw_id (AccessGatewayID)
	GetChallenge(ctx context.Context, in *AccessGatewayID, opts ...grpc.CallOption) (*Challenge, error)
	// send back response and csr for signing
	// Returns signed certificate.
	RequestSign(ctx context.Context, in *Response, opts ...grpc.CallOption) (*Certificate, error)
}

type bootstrapperClient struct {
	cc grpc.ClientConnInterface
}

func NewBootstrapperClient(cc grpc.ClientConnInterface) BootstrapperClient {
	return &bootstrapperClient{cc}
}

func (c *bootstrapperClient) GetChallenge(ctx context.Context, in *AccessGatewayID, opts ...grpc.CallOption) (*Challenge, error) {
	out := new(Challenge)
	err := c.cc.Invoke(ctx, "/magma.orc8r.Bootstrapper/GetChallenge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bootstrapperClient) RequestSign(ctx context.Context, in *Response, opts ...grpc.CallOption) (*Certificate, error) {
	out := new(Certificate)
	err := c.cc.Invoke(ctx, "/magma.orc8r.Bootstrapper/RequestSign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BootstrapperServer is the server API for Bootstrapper service.
type BootstrapperServer interface {
	// get the challange for gateway specified in hw_id (AccessGatewayID)
	GetChallenge(context.Context, *AccessGatewayID) (*Challenge, error)
	// send back response and csr for signing
	// Returns signed certificate.
	RequestSign(context.Context, *Response) (*Certificate, error)
}

// UnimplementedBootstrapperServer can be embedded to have forward compatible implementations.
type UnimplementedBootstrapperServer struct {
}

func (*UnimplementedBootstrapperServer) GetChallenge(ctx context.Context, req *AccessGatewayID) (*Challenge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChallenge not implemented")
}
func (*UnimplementedBootstrapperServer) RequestSign(ctx context.Context, req *Response) (*Certificate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestSign not implemented")
}

func RegisterBootstrapperServer(s *grpc.Server, srv BootstrapperServer) {
	s.RegisterService(&_Bootstrapper_serviceDesc, srv)
}

func _Bootstrapper_GetChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessGatewayID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BootstrapperServer).GetChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.Bootstrapper/GetChallenge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BootstrapperServer).GetChallenge(ctx, req.(*AccessGatewayID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bootstrapper_RequestSign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Response)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BootstrapperServer).RequestSign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.Bootstrapper/RequestSign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BootstrapperServer).RequestSign(ctx, req.(*Response))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bootstrapper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.Bootstrapper",
	HandlerType: (*BootstrapperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChallenge",
			Handler:    _Bootstrapper_GetChallenge_Handler,
		},
		{
			MethodName: "RequestSign",
			Handler:    _Bootstrapper_RequestSign_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orc8r/protos/bootstrapper.proto",
}
