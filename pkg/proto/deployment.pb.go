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
	ReleaseName      string       `protobuf:"bytes,1,opt,name=release_name,json=releaseName" json:"release_name,omitempty"`
	RawConfig        string       `protobuf:"bytes,2,opt,name=raw_config,json=rawConfig" json:"raw_config,omitempty"`
	Chart            *Chart       `protobuf:"bytes,3,opt,name=chart" json:"chart,omitempty"`
	OrganizationUuid *UUID        `protobuf:"bytes,6,opt,name=organization_uuid,json=organizationUuid" json:"organization_uuid,omitempty"`
	Secrets          []*Secret    `protobuf:"bytes,10,rep,name=secrets" json:"secrets,omitempty"`
	EnvSecret        []*EnvSecret `protobuf:"bytes,11,rep,name=env_secret,json=envSecret" json:"env_secret,omitempty"`
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

func (m *CreateDeploymentRequest) GetOrganizationUuid() *UUID {
	if m != nil {
		return m.OrganizationUuid
	}
	return nil
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
	ReleaseName string `protobuf:"bytes,1,opt,name=release_name,json=releaseName" json:"release_name,omitempty"`
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
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x94, 0xcf, 0x8b, 0xda, 0x40,
	0x14, 0xc7, 0x51, 0xab, 0x92, 0x17, 0xa1, 0x3a, 0x28, 0x4e, 0x0b, 0x05, 0x9b, 0x43, 0xb1, 0x2d,
	0x78, 0x50, 0x4a, 0x2f, 0x9e, 0xaa, 0xa5, 0xf4, 0xd2, 0x43, 0x4a, 0x2e, 0xbd, 0x84, 0x69, 0xf2,
	0x6a, 0x03, 0xc9, 0x4c, 0x3a, 0x33, 0x51, 0xda, 0xee, 0xde, 0xf6, 0x3f, 0xd9, 0x7f, 0x74, 0xc9,
	0xf8, 0x6b, 0xd8, 0xec, 0xb2, 0xec, 0x1e, 0xdc, 0xcb, 0x9e, 0x7c, 0xbe, 0xf7, 0x79, 0x99, 0xef,
	0x7c, 0xbf, 0x21, 0xd0, 0x8d, 0x31, 0x4f, 0xc5, 0xdf, 0x0c, 0xb9, 0x9e, 0xe4, 0x52, 0x68, 0x41,
	0x9c, 0x48, 0x64, 0x19, 0xe3, 0x31, 0xca, 0x97, 0x9d, 0xb2, 0x14, 0x7c, 0x3b, 0xf0, 0xde, 0x41,
	0xff, 0x0b, 0xea, 0xe5, 0x81, 0xf7, 0xf1, 0x4f, 0x81, 0x4a, 0x13, 0x02, 0xcf, 0x38, 0xcb, 0x90,
	0xd6, 0x46, 0xb5, 0xb1, 0xe3, 0x9b, 0xda, 0xfb, 0x0f, 0x83, 0x6b, 0xac, 0xca, 0x05, 0x57, 0x48,
	0xde, 0x42, 0x4b, 0xa2, 0x2a, 0x52, 0x6d, 0x70, 0x77, 0xda, 0x9b, 0x1c, 0x8e, 0x9b, 0xf8, 0x66,
	0xe0, 0xef, 0x00, 0xf2, 0x11, 0xdc, 0xa3, 0x38, 0x45, 0xeb, 0xa3, 0xc6, 0xd8, 0x9d, 0x0e, 0x2c,
	0xde, 0x7a, 0xbc, 0x4d, 0x7a, 0x97, 0x75, 0x18, 0x2e, 0x24, 0x32, 0x8d, 0x55, 0xb1, 0xaf, 0xa1,
	0x23, 0x31, 0x45, 0xa6, 0x30, 0xb4, 0x44, 0xbb, 0xbb, 0xde, 0x37, 0x96, 0x21, 0x79, 0x05, 0x20,
	0xd9, 0x26, 0x8c, 0x04, 0xff, 0x95, 0xac, 0x68, 0xdd, 0x00, 0x8e, 0x64, 0x9b, 0x85, 0x69, 0x90,
	0x37, 0xd0, 0x8c, 0x7e, 0x33, 0xa9, 0x69, 0xc3, 0x5c, 0xa0, 0x6b, 0x09, 0x5a, 0x94, 0x7d, 0x7f,
	0x3b, 0x26, 0x73, 0xe8, 0x09, 0xb9, 0x62, 0x3c, 0xf9, 0xc7, 0x74, 0x22, 0x78, 0x58, 0x14, 0x49,
	0x4c, 0x5b, 0x66, 0xe7, 0xb9, 0xb5, 0x13, 0x04, 0x5f, 0x97, 0x7e, 0xd7, 0x26, 0x83, 0x22, 0x89,
	0xc9, 0x7b, 0x68, 0x2b, 0x8c, 0x24, 0x6a, 0x45, 0xc1, 0x5c, 0xdc, 0x36, 0xea, 0xbb, 0x99, 0xf8,
	0x7b, 0x82, 0xcc, 0x00, 0x90, 0xaf, 0xc3, 0xed, 0x5f, 0xea, 0x1a, 0xbe, 0x6f, 0xf1, 0x9f, 0xf9,
	0x7a, 0xb7, 0xe2, 0xe0, 0xbe, 0xf4, 0xce, 0x80, 0x56, 0x4d, 0xba, 0x7f, 0x4a, 0x1f, 0x00, 0x8e,
	0xde, 0x1b, 0xb7, 0x6e, 0x0d, 0xc9, 0x02, 0x4d, 0x46, 0x41, 0x1e, 0x3f, 0x65, 0x74, 0x57, 0x46,
	0x55, 0x93, 0x4e, 0x96, 0xd1, 0x45, 0xad, 0x3c, 0x7e, 0x25, 0x59, 0xfc, 0x98, 0x21, 0x79, 0xe7,
	0xf0, 0xe2, 0x06, 0x15, 0x27, 0x73, 0x61, 0x0e, 0xc3, 0x25, 0xa6, 0xf8, 0xb0, 0x17, 0xb5, 0x4c,
	0xb0, 0xba, 0x7d, 0x2a, 0xed, 0x9f, 0xda, 0x3f, 0x9a, 0xe6, 0xdb, 0xfd, 0xb3, 0x65, 0x7e, 0x66,
	0x57, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x14, 0x7e, 0x8b, 0xef, 0x05, 0x00, 0x00,
}
