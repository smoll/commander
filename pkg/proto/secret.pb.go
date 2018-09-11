// Code generated by protoc-gen-go. DO NOT EDIT.
// source: secret.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	secret.proto
	common.proto
	deployment.proto
	commander.proto

It has these top-level messages:
	GetSecretRequest
	GetSecretResponse
	PingRequest
	PingResponse
	UUID
	Env
	EnvSecret
	Secret
	Result
	Chart
	Deployment
	GetDeploymentRequest
	GetDeploymentResponse
	CreateDeploymentRequest
	CreateDeploymentResponse
	UpdateDeploymentRequest
	UpdateDeploymentResponse
	UpgradeDeploymentRequest
	UpgradeDeploymentResponse
	DeleteDeploymentRequest
	DeleteDeploymentResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type GetSecretRequest struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *GetSecretRequest) Reset()                    { *m = GetSecretRequest{} }
func (m *GetSecretRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetSecretRequest) ProtoMessage()               {}
func (*GetSecretRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetSecretRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetSecretRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetSecretResponse struct {
	Result *Result `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	Secret *Secret `protobuf:"bytes,2,opt,name=secret" json:"secret,omitempty"`
}

func (m *GetSecretResponse) Reset()                    { *m = GetSecretResponse{} }
func (m *GetSecretResponse) String() string            { return proto1.CompactTextString(m) }
func (*GetSecretResponse) ProtoMessage()               {}
func (*GetSecretResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetSecretResponse) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GetSecretResponse) GetSecret() *Secret {
	if m != nil {
		return m.Secret
	}
	return nil
}

func init() {
	proto1.RegisterType((*GetSecretRequest)(nil), "commander.GetSecretRequest")
	proto1.RegisterType((*GetSecretResponse)(nil), "commander.GetSecretResponse")
}

func init() { proto1.RegisterFile("secret.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x4d, 0x2e,
	0x4a, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0xce, 0xcf, 0xcd, 0x4d, 0xcc,
	0x4b, 0x49, 0x2d, 0x92, 0xe2, 0x01, 0x31, 0xf3, 0xf3, 0x20, 0x12, 0x4a, 0x2e, 0x5c, 0x02, 0xee,
	0xa9, 0x25, 0xc1, 0x60, 0xb5, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x32, 0x5c, 0x9c,
	0x79, 0x89, 0xb9, 0xa9, 0xc5, 0x05, 0x89, 0xc9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x08, 0x01, 0x21, 0x21, 0x2e, 0x16, 0x10, 0x47, 0x82, 0x09, 0x2c, 0x01, 0x66, 0x2b, 0x65, 0x72,
	0x09, 0x22, 0x99, 0x52, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0xa4, 0xc9, 0xc5, 0x56, 0x94, 0x5a,
	0x5c, 0x9a, 0x53, 0x02, 0x36, 0x83, 0xdb, 0x48, 0x50, 0x0f, 0xee, 0x08, 0xbd, 0x20, 0xb0, 0x44,
	0x10, 0x54, 0x01, 0x48, 0x29, 0xc4, 0xb9, 0x60, 0x53, 0x51, 0x95, 0x42, 0x4d, 0x85, 0x2a, 0x70,
	0x62, 0x8f, 0x62, 0x05, 0xbb, 0x3c, 0x89, 0x0d, 0x4c, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x1f, 0x17, 0x76, 0x35, 0xe9, 0x00, 0x00, 0x00,
}
