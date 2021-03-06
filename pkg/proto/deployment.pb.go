// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deployment.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GetDeploymentRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetDeploymentRequest) Reset()                    { *m = GetDeploymentRequest{} }
func (m *GetDeploymentRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetDeploymentRequest) ProtoMessage()               {}
func (*GetDeploymentRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *GetDeploymentRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetDeploymentResponse struct {
	Result      *Result       `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	Deployments []*Deployment `protobuf:"bytes,2,rep,name=deployments" json:"deployments,omitempty"`
}

func (m *GetDeploymentResponse) Reset()                    { *m = GetDeploymentResponse{} }
func (m *GetDeploymentResponse) String() string            { return proto1.CompactTextString(m) }
func (*GetDeploymentResponse) ProtoMessage()               {}
func (*GetDeploymentResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *GetDeploymentResponse) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GetDeploymentResponse) GetDeployments() []*Deployment {
	if m != nil {
		return m.Deployments
	}
	return nil
}

type CreateDeploymentRequest struct {
	ReleaseName string       `protobuf:"bytes,1,opt,name=release_name,json=releaseName" json:"release_name,omitempty"`
	RawConfig   string       `protobuf:"bytes,2,opt,name=raw_config,json=rawConfig" json:"raw_config,omitempty"`
	Chart       *Chart       `protobuf:"bytes,3,opt,name=chart" json:"chart,omitempty"`
	Namespace   string       `protobuf:"bytes,4,opt,name=namespace" json:"namespace,omitempty"`
	Secrets     []*Secret    `protobuf:"bytes,10,rep,name=secrets" json:"secrets,omitempty"`
	EnvSecret   []*EnvSecret `protobuf:"bytes,11,rep,name=env_secret,json=envSecret" json:"env_secret,omitempty"`
}

func (m *CreateDeploymentRequest) Reset()                    { *m = CreateDeploymentRequest{} }
func (m *CreateDeploymentRequest) String() string            { return proto1.CompactTextString(m) }
func (*CreateDeploymentRequest) ProtoMessage()               {}
func (*CreateDeploymentRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *CreateDeploymentRequest) GetReleaseName() string {
	if m != nil {
		return m.ReleaseName
	}
	return ""
}

func (m *CreateDeploymentRequest) GetRawConfig() string {
	if m != nil {
		return m.RawConfig
	}
	return ""
}

func (m *CreateDeploymentRequest) GetChart() *Chart {
	if m != nil {
		return m.Chart
	}
	return nil
}

func (m *CreateDeploymentRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CreateDeploymentRequest) GetSecrets() []*Secret {
	if m != nil {
		return m.Secrets
	}
	return nil
}

func (m *CreateDeploymentRequest) GetEnvSecret() []*EnvSecret {
	if m != nil {
		return m.EnvSecret
	}
	return nil
}

type CreateDeploymentResponse struct {
	Result     *Result     `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	Deployment *Deployment `protobuf:"bytes,2,opt,name=deployment" json:"deployment,omitempty"`
}

func (m *CreateDeploymentResponse) Reset()                    { *m = CreateDeploymentResponse{} }
func (m *CreateDeploymentResponse) String() string            { return proto1.CompactTextString(m) }
func (*CreateDeploymentResponse) ProtoMessage()               {}
func (*CreateDeploymentResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *CreateDeploymentResponse) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *CreateDeploymentResponse) GetDeployment() *Deployment {
	if m != nil {
		return m.Deployment
	}
	return nil
}

type UpdateDeploymentRequest struct {
	ReleaseName      string       `protobuf:"bytes,1,opt,name=release_name,json=releaseName" json:"release_name,omitempty"`
	RawConfig        string       `protobuf:"bytes,2,opt,name=raw_config,json=rawConfig" json:"raw_config,omitempty"`
	Chart            *Chart       `protobuf:"bytes,3,opt,name=chart" json:"chart,omitempty"`
	OrganizationUuid *UUID        `protobuf:"bytes,6,opt,name=organization_uuid,json=organizationUuid" json:"organization_uuid,omitempty"`
	Secrets          []*Secret    `protobuf:"bytes,10,rep,name=secrets" json:"secrets,omitempty"`
	EnvSecret        []*EnvSecret `protobuf:"bytes,11,rep,name=env_secret,json=envSecret" json:"env_secret,omitempty"`
}

func (m *UpdateDeploymentRequest) Reset()                    { *m = UpdateDeploymentRequest{} }
func (m *UpdateDeploymentRequest) String() string            { return proto1.CompactTextString(m) }
func (*UpdateDeploymentRequest) ProtoMessage()               {}
func (*UpdateDeploymentRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *UpdateDeploymentRequest) GetReleaseName() string {
	if m != nil {
		return m.ReleaseName
	}
	return ""
}

func (m *UpdateDeploymentRequest) GetRawConfig() string {
	if m != nil {
		return m.RawConfig
	}
	return ""
}

func (m *UpdateDeploymentRequest) GetChart() *Chart {
	if m != nil {
		return m.Chart
	}
	return nil
}

func (m *UpdateDeploymentRequest) GetOrganizationUuid() *UUID {
	if m != nil {
		return m.OrganizationUuid
	}
	return nil
}

func (m *UpdateDeploymentRequest) GetSecrets() []*Secret {
	if m != nil {
		return m.Secrets
	}
	return nil
}

func (m *UpdateDeploymentRequest) GetEnvSecret() []*EnvSecret {
	if m != nil {
		return m.EnvSecret
	}
	return nil
}

type UpdateDeploymentResponse struct {
	Result     *Result     `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	Deployment *Deployment `protobuf:"bytes,2,opt,name=deployment" json:"deployment,omitempty"`
}

func (m *UpdateDeploymentResponse) Reset()                    { *m = UpdateDeploymentResponse{} }
func (m *UpdateDeploymentResponse) String() string            { return proto1.CompactTextString(m) }
func (*UpdateDeploymentResponse) ProtoMessage()               {}
func (*UpdateDeploymentResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *UpdateDeploymentResponse) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *UpdateDeploymentResponse) GetDeployment() *Deployment {
	if m != nil {
		return m.Deployment
	}
	return nil
}

type UpgradeDeploymentRequest struct {
	ReleaseName string `protobuf:"bytes,1,opt,name=release_name,json=releaseName" json:"release_name,omitempty"`
	RawConfig   string `protobuf:"bytes,2,opt,name=raw_config,json=rawConfig" json:"raw_config,omitempty"`
	Chart       *Chart `protobuf:"bytes,3,opt,name=chart" json:"chart,omitempty"`
}

func (m *UpgradeDeploymentRequest) Reset()                    { *m = UpgradeDeploymentRequest{} }
func (m *UpgradeDeploymentRequest) String() string            { return proto1.CompactTextString(m) }
func (*UpgradeDeploymentRequest) ProtoMessage()               {}
func (*UpgradeDeploymentRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *UpgradeDeploymentRequest) GetReleaseName() string {
	if m != nil {
		return m.ReleaseName
	}
	return ""
}

func (m *UpgradeDeploymentRequest) GetRawConfig() string {
	if m != nil {
		return m.RawConfig
	}
	return ""
}

func (m *UpgradeDeploymentRequest) GetChart() *Chart {
	if m != nil {
		return m.Chart
	}
	return nil
}

type UpgradeDeploymentResponse struct {
	Result     *Result     `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	Deployment *Deployment `protobuf:"bytes,2,opt,name=deployment" json:"deployment,omitempty"`
}

func (m *UpgradeDeploymentResponse) Reset()                    { *m = UpgradeDeploymentResponse{} }
func (m *UpgradeDeploymentResponse) String() string            { return proto1.CompactTextString(m) }
func (*UpgradeDeploymentResponse) ProtoMessage()               {}
func (*UpgradeDeploymentResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *UpgradeDeploymentResponse) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *UpgradeDeploymentResponse) GetDeployment() *Deployment {
	if m != nil {
		return m.Deployment
	}
	return nil
}

type DeleteDeploymentRequest struct {
	ReleaseName     string `protobuf:"bytes,1,opt,name=release_name,json=releaseName" json:"release_name,omitempty"`
	Namespace       string `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	DeleteNamespace bool   `protobuf:"varint,5,opt,name=delete_namespace,json=deleteNamespace" json:"delete_namespace,omitempty"`
}

func (m *DeleteDeploymentRequest) Reset()                    { *m = DeleteDeploymentRequest{} }
func (m *DeleteDeploymentRequest) String() string            { return proto1.CompactTextString(m) }
func (*DeleteDeploymentRequest) ProtoMessage()               {}
func (*DeleteDeploymentRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *DeleteDeploymentRequest) GetReleaseName() string {
	if m != nil {
		return m.ReleaseName
	}
	return ""
}

func (m *DeleteDeploymentRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DeleteDeploymentRequest) GetDeleteNamespace() bool {
	if m != nil {
		return m.DeleteNamespace
	}
	return false
}

type DeleteDeploymentResponse struct {
	Result     *Result     `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	Deployment *Deployment `protobuf:"bytes,2,opt,name=deployment" json:"deployment,omitempty"`
}

func (m *DeleteDeploymentResponse) Reset()                    { *m = DeleteDeploymentResponse{} }
func (m *DeleteDeploymentResponse) String() string            { return proto1.CompactTextString(m) }
func (*DeleteDeploymentResponse) ProtoMessage()               {}
func (*DeleteDeploymentResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *DeleteDeploymentResponse) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *DeleteDeploymentResponse) GetDeployment() *Deployment {
	if m != nil {
		return m.Deployment
	}
	return nil
}

func init() {
	proto1.RegisterType((*GetDeploymentRequest)(nil), "commander.GetDeploymentRequest")
	proto1.RegisterType((*GetDeploymentResponse)(nil), "commander.GetDeploymentResponse")
	proto1.RegisterType((*CreateDeploymentRequest)(nil), "commander.CreateDeploymentRequest")
	proto1.RegisterType((*CreateDeploymentResponse)(nil), "commander.CreateDeploymentResponse")
	proto1.RegisterType((*UpdateDeploymentRequest)(nil), "commander.UpdateDeploymentRequest")
	proto1.RegisterType((*UpdateDeploymentResponse)(nil), "commander.UpdateDeploymentResponse")
	proto1.RegisterType((*UpgradeDeploymentRequest)(nil), "commander.UpgradeDeploymentRequest")
	proto1.RegisterType((*UpgradeDeploymentResponse)(nil), "commander.UpgradeDeploymentResponse")
	proto1.RegisterType((*DeleteDeploymentRequest)(nil), "commander.DeleteDeploymentRequest")
	proto1.RegisterType((*DeleteDeploymentResponse)(nil), "commander.DeleteDeploymentResponse")
}

func init() { proto1.RegisterFile("deployment.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x86, 0x95, 0xb4, 0xbb, 0x25, 0x93, 0x4a, 0xdd, 0x5a, 0xad, 0xd6, 0x20, 0x90, 0x96, 0x1c,
	0xd0, 0x16, 0xa4, 0x3d, 0xb4, 0x42, 0x5c, 0x38, 0xb1, 0x8b, 0x10, 0x97, 0x1e, 0x8c, 0x72, 0xe1,
	0x12, 0x99, 0x64, 0x58, 0x22, 0x25, 0x76, 0xb0, 0x9d, 0x56, 0x7c, 0x9d, 0x80, 0x5f, 0xc2, 0x1f,
	0x45, 0x71, 0xf6, 0xc3, 0x34, 0x70, 0xa0, 0x87, 0xec, 0x29, 0xce, 0xcc, 0x63, 0xcf, 0xf8, 0x7d,
	0x6d, 0xc3, 0x28, 0xc3, 0xaa, 0x90, 0x9f, 0x4a, 0x14, 0x66, 0x56, 0x29, 0x69, 0x24, 0x09, 0x52,
	0x59, 0x96, 0x5c, 0x64, 0xa8, 0xee, 0x1d, 0x36, 0x43, 0x29, 0xda, 0x44, 0xf4, 0x18, 0x4e, 0x5e,
	0xa1, 0x59, 0x6c, 0x78, 0x86, 0x1f, 0x6b, 0xd4, 0x86, 0x10, 0xd8, 0x17, 0xbc, 0x44, 0xea, 0x4d,
	0xbc, 0x69, 0xc0, 0xec, 0x38, 0xfa, 0x02, 0xa7, 0x37, 0x58, 0x5d, 0x49, 0xa1, 0x91, 0x9c, 0xc1,
	0x50, 0xa1, 0xae, 0x0b, 0x63, 0xf1, 0xf0, 0xfc, 0x78, 0xb6, 0x29, 0x37, 0x63, 0x36, 0xc1, 0x56,
	0x00, 0x79, 0x06, 0xe1, 0xb6, 0x39, 0x4d, 0xfd, 0xc9, 0xde, 0x34, 0x3c, 0x3f, 0x75, 0x78, 0x67,
	0x79, 0x97, 0x8c, 0xbe, 0xfb, 0x30, 0x9e, 0x2b, 0xe4, 0x06, 0xbb, 0xcd, 0x3e, 0x84, 0x43, 0x85,
	0x05, 0x72, 0x8d, 0x89, 0xd3, 0x74, 0xb8, 0x8a, 0x5d, 0xf2, 0x12, 0xc9, 0x03, 0x00, 0xc5, 0xaf,
	0x93, 0x54, 0x8a, 0xf7, 0xf9, 0x92, 0xfa, 0x16, 0x08, 0x14, 0xbf, 0x9e, 0xdb, 0x00, 0x79, 0x04,
	0x83, 0xf4, 0x03, 0x57, 0x86, 0xee, 0xd9, 0x0d, 0x8c, 0x9c, 0x86, 0xe6, 0x4d, 0x9c, 0xb5, 0x69,
	0x72, 0x1f, 0x82, 0xa6, 0x82, 0xae, 0x78, 0x8a, 0x74, 0xbf, 0x5d, 0x65, 0x13, 0x20, 0x4f, 0xe0,
	0x40, 0x63, 0xaa, 0xd0, 0x68, 0x0a, 0x76, 0x63, 0xae, 0x10, 0x6f, 0x6c, 0x86, 0xad, 0x09, 0x72,
	0x01, 0x80, 0xe2, 0x2a, 0x69, 0x7f, 0x69, 0x68, 0xf9, 0x13, 0x87, 0x7f, 0x29, 0xae, 0x56, 0x53,
	0x02, 0x5c, 0x0f, 0xa3, 0xaf, 0x40, 0xbb, 0x22, 0xfc, 0xbf, 0x0b, 0x4f, 0x01, 0xb6, 0xda, 0x5a,
	0x35, 0xfe, 0x69, 0x82, 0x03, 0x46, 0xbf, 0x7c, 0x18, 0xc7, 0x55, 0xb6, 0x63, 0x0f, 0x9e, 0xc3,
	0xb1, 0x54, 0x4b, 0x2e, 0xf2, 0xcf, 0xdc, 0xe4, 0x52, 0x24, 0x75, 0x9d, 0x67, 0x74, 0x68, 0xe7,
	0x1c, 0x39, 0x73, 0xe2, 0xf8, 0xf5, 0x82, 0x8d, 0x5c, 0x32, 0xae, 0xf3, 0xac, 0x1f, 0x8f, 0xba,
	0x22, 0xf5, 0xe6, 0xd1, 0x0f, 0xaf, 0x29, 0xbf, 0x54, 0x3c, 0xdb, 0xa5, 0x49, 0xd1, 0x37, 0xb8,
	0xfb, 0x97, 0x2e, 0x7a, 0x53, 0xe1, 0xa7, 0x07, 0xe3, 0x05, 0x16, 0x78, 0xcb, 0x93, 0xfa, 0xc7,
	0x35, 0xf7, 0x6f, 0x5e, 0xf3, 0xb3, 0xe6, 0x81, 0x6d, 0xd6, 0x4e, 0xb6, 0xd0, 0x60, 0xe2, 0x4d,
	0xef, 0xb0, 0xa3, 0x36, 0x7e, 0xb9, 0x0e, 0x37, 0x67, 0xa1, 0xdb, 0x46, 0x5f, 0x2a, 0xbc, 0x38,
	0x78, 0x3b, 0xb0, 0xaf, 0xfc, 0xbb, 0xa1, 0xfd, 0x5c, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x3a,
	0x8c, 0x47, 0xa4, 0x19, 0x06, 0x00, 0x00,
}
