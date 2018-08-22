// Code generated by protoc-gen-go. DO NOT EDIT.
// source: drivers.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	drivers.proto

It has these top-level messages:
	Empty
	DriverFlags
	Flag
	Default
	DriverOptions
	StringSlice
	ClusterInfo
	KubernetesVersion
	NodeCount
	Capabilities
	CreateRequest
	UpdateRequest
	SetVersionRequest
	SetNodeCountRequest
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DriverFlags struct {
	Options map[string]*Flag `protobuf:"bytes,1,rep,name=options" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *DriverFlags) Reset()                    { *m = DriverFlags{} }
func (m *DriverFlags) String() string            { return proto.CompactTextString(m) }
func (*DriverFlags) ProtoMessage()               {}
func (*DriverFlags) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DriverFlags) GetOptions() map[string]*Flag {
	if m != nil {
		return m.Options
	}
	return nil
}

type Flag struct {
	Type    string   `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Usage   string   `protobuf:"bytes,2,opt,name=usage" json:"usage,omitempty"`
	Value   string   `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	Default *Default `protobuf:"bytes,4,opt,name=default" json:"default,omitempty"`
}

func (m *Flag) Reset()                    { *m = Flag{} }
func (m *Flag) String() string            { return proto.CompactTextString(m) }
func (*Flag) ProtoMessage()               {}
func (*Flag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Flag) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Flag) GetUsage() string {
	if m != nil {
		return m.Usage
	}
	return ""
}

func (m *Flag) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Flag) GetDefault() *Default {
	if m != nil {
		return m.Default
	}
	return nil
}

type Default struct {
	DefaultBool        bool         `protobuf:"varint,1,opt,name=defaultBool" json:"defaultBool,omitempty"`
	DefaultString      string       `protobuf:"bytes,2,opt,name=defaultString" json:"defaultString,omitempty"`
	DefaultStringSlice *StringSlice `protobuf:"bytes,3,opt,name=defaultStringSlice" json:"defaultStringSlice,omitempty"`
	DefaultInt         int64        `protobuf:"varint,4,opt,name=defaultInt" json:"defaultInt,omitempty"`
}

func (m *Default) Reset()                    { *m = Default{} }
func (m *Default) String() string            { return proto.CompactTextString(m) }
func (*Default) ProtoMessage()               {}
func (*Default) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Default) GetDefaultBool() bool {
	if m != nil {
		return m.DefaultBool
	}
	return false
}

func (m *Default) GetDefaultString() string {
	if m != nil {
		return m.DefaultString
	}
	return ""
}

func (m *Default) GetDefaultStringSlice() *StringSlice {
	if m != nil {
		return m.DefaultStringSlice
	}
	return nil
}

func (m *Default) GetDefaultInt() int64 {
	if m != nil {
		return m.DefaultInt
	}
	return 0
}

type DriverOptions struct {
	BoolOptions        map[string]bool         `protobuf:"bytes,1,rep,name=bool_options,json=boolOptions" json:"bool_options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	StringOptions      map[string]string       `protobuf:"bytes,2,rep,name=string_options,json=stringOptions" json:"string_options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	IntOptions         map[string]int64        `protobuf:"bytes,3,rep,name=int_options,json=intOptions" json:"int_options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	StringSliceOptions map[string]*StringSlice `protobuf:"bytes,4,rep,name=string_slice_options,json=stringSliceOptions" json:"string_slice_options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *DriverOptions) Reset()                    { *m = DriverOptions{} }
func (m *DriverOptions) String() string            { return proto.CompactTextString(m) }
func (*DriverOptions) ProtoMessage()               {}
func (*DriverOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DriverOptions) GetBoolOptions() map[string]bool {
	if m != nil {
		return m.BoolOptions
	}
	return nil
}

func (m *DriverOptions) GetStringOptions() map[string]string {
	if m != nil {
		return m.StringOptions
	}
	return nil
}

func (m *DriverOptions) GetIntOptions() map[string]int64 {
	if m != nil {
		return m.IntOptions
	}
	return nil
}

func (m *DriverOptions) GetStringSliceOptions() map[string]*StringSlice {
	if m != nil {
		return m.StringSliceOptions
	}
	return nil
}

type StringSlice struct {
	Value []string `protobuf:"bytes,1,rep,name=value" json:"value,omitempty"`
}

func (m *StringSlice) Reset()                    { *m = StringSlice{} }
func (m *StringSlice) String() string            { return proto.CompactTextString(m) }
func (*StringSlice) ProtoMessage()               {}
func (*StringSlice) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *StringSlice) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type ClusterInfo struct {
	Version             string            `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	ServiceAccountToken string            `protobuf:"bytes,2,opt,name=service_account_token,json=serviceAccountToken" json:"service_account_token,omitempty"`
	Endpoint            string            `protobuf:"bytes,3,opt,name=endpoint" json:"endpoint,omitempty"`
	Username            string            `protobuf:"bytes,4,opt,name=username" json:"username,omitempty"`
	Password            string            `protobuf:"bytes,5,opt,name=password" json:"password,omitempty"`
	RootCaCertificate   string            `protobuf:"bytes,6,opt,name=root_ca_certificate,json=rootCaCertificate" json:"root_ca_certificate,omitempty"`
	ClientCertificate   string            `protobuf:"bytes,7,opt,name=client_certificate,json=clientCertificate" json:"client_certificate,omitempty"`
	ClientKey           string            `protobuf:"bytes,8,opt,name=client_key,json=clientKey" json:"client_key,omitempty"`
	NodeCount           int64             `protobuf:"varint,9,opt,name=node_count,json=nodeCount" json:"node_count,omitempty"`
	Metadata            map[string]string `protobuf:"bytes,10,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Status              string            `protobuf:"bytes,11,opt,name=status" json:"status,omitempty"`
	CreateError         string            `protobuf:"bytes,12,opt,name=create_error,json=createError" json:"create_error,omitempty"`
}

func (m *ClusterInfo) Reset()                    { *m = ClusterInfo{} }
func (m *ClusterInfo) String() string            { return proto.CompactTextString(m) }
func (*ClusterInfo) ProtoMessage()               {}
func (*ClusterInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ClusterInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ClusterInfo) GetServiceAccountToken() string {
	if m != nil {
		return m.ServiceAccountToken
	}
	return ""
}

func (m *ClusterInfo) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *ClusterInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ClusterInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ClusterInfo) GetRootCaCertificate() string {
	if m != nil {
		return m.RootCaCertificate
	}
	return ""
}

func (m *ClusterInfo) GetClientCertificate() string {
	if m != nil {
		return m.ClientCertificate
	}
	return ""
}

func (m *ClusterInfo) GetClientKey() string {
	if m != nil {
		return m.ClientKey
	}
	return ""
}

func (m *ClusterInfo) GetNodeCount() int64 {
	if m != nil {
		return m.NodeCount
	}
	return 0
}

func (m *ClusterInfo) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ClusterInfo) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ClusterInfo) GetCreateError() string {
	if m != nil {
		return m.CreateError
	}
	return ""
}

type KubernetesVersion struct {
	Version string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
}

func (m *KubernetesVersion) Reset()                    { *m = KubernetesVersion{} }
func (m *KubernetesVersion) String() string            { return proto.CompactTextString(m) }
func (*KubernetesVersion) ProtoMessage()               {}
func (*KubernetesVersion) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *KubernetesVersion) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type NodeCount struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *NodeCount) Reset()                    { *m = NodeCount{} }
func (m *NodeCount) String() string            { return proto.CompactTextString(m) }
func (*NodeCount) ProtoMessage()               {}
func (*NodeCount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *NodeCount) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Capabilities struct {
	Capabilities map[int64]bool `protobuf:"bytes,1,rep,name=capabilities" json:"capabilities,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *Capabilities) Reset()                    { *m = Capabilities{} }
func (m *Capabilities) String() string            { return proto.CompactTextString(m) }
func (*Capabilities) ProtoMessage()               {}
func (*Capabilities) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Capabilities) GetCapabilities() map[int64]bool {
	if m != nil {
		return m.Capabilities
	}
	return nil
}

type CreateRequest struct {
	DriverOptions *DriverOptions `protobuf:"bytes,1,opt,name=driver_options,json=driverOptions" json:"driver_options,omitempty"`
	ClusterInfo   *ClusterInfo   `protobuf:"bytes,2,opt,name=cluster_info,json=clusterInfo" json:"cluster_info,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *CreateRequest) GetDriverOptions() *DriverOptions {
	if m != nil {
		return m.DriverOptions
	}
	return nil
}

func (m *CreateRequest) GetClusterInfo() *ClusterInfo {
	if m != nil {
		return m.ClusterInfo
	}
	return nil
}

type UpdateRequest struct {
	ClusterInfo   *ClusterInfo   `protobuf:"bytes,1,opt,name=cluster_info,json=clusterInfo" json:"cluster_info,omitempty"`
	DriverOptions *DriverOptions `protobuf:"bytes,2,opt,name=driver_options,json=driverOptions" json:"driver_options,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *UpdateRequest) GetClusterInfo() *ClusterInfo {
	if m != nil {
		return m.ClusterInfo
	}
	return nil
}

func (m *UpdateRequest) GetDriverOptions() *DriverOptions {
	if m != nil {
		return m.DriverOptions
	}
	return nil
}

type SetVersionRequest struct {
	Info    *ClusterInfo       `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
	Version *KubernetesVersion `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *SetVersionRequest) Reset()                    { *m = SetVersionRequest{} }
func (m *SetVersionRequest) String() string            { return proto.CompactTextString(m) }
func (*SetVersionRequest) ProtoMessage()               {}
func (*SetVersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *SetVersionRequest) GetInfo() *ClusterInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *SetVersionRequest) GetVersion() *KubernetesVersion {
	if m != nil {
		return m.Version
	}
	return nil
}

type SetNodeCountRequest struct {
	Info  *ClusterInfo `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
	Count *NodeCount   `protobuf:"bytes,2,opt,name=count" json:"count,omitempty"`
}

func (m *SetNodeCountRequest) Reset()                    { *m = SetNodeCountRequest{} }
func (m *SetNodeCountRequest) String() string            { return proto.CompactTextString(m) }
func (*SetNodeCountRequest) ProtoMessage()               {}
func (*SetNodeCountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *SetNodeCountRequest) GetInfo() *ClusterInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *SetNodeCountRequest) GetCount() *NodeCount {
	if m != nil {
		return m.Count
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "types.Empty")
	proto.RegisterType((*DriverFlags)(nil), "types.DriverFlags")
	proto.RegisterType((*Flag)(nil), "types.Flag")
	proto.RegisterType((*Default)(nil), "types.Default")
	proto.RegisterType((*DriverOptions)(nil), "types.DriverOptions")
	proto.RegisterType((*StringSlice)(nil), "types.StringSlice")
	proto.RegisterType((*ClusterInfo)(nil), "types.ClusterInfo")
	proto.RegisterType((*KubernetesVersion)(nil), "types.KubernetesVersion")
	proto.RegisterType((*NodeCount)(nil), "types.NodeCount")
	proto.RegisterType((*Capabilities)(nil), "types.Capabilities")
	proto.RegisterType((*CreateRequest)(nil), "types.CreateRequest")
	proto.RegisterType((*UpdateRequest)(nil), "types.UpdateRequest")
	proto.RegisterType((*SetVersionRequest)(nil), "types.SetVersionRequest")
	proto.RegisterType((*SetNodeCountRequest)(nil), "types.SetNodeCountRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Driver service

type DriverClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*ClusterInfo, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*ClusterInfo, error)
	PostCheck(ctx context.Context, in *ClusterInfo, opts ...grpc.CallOption) (*ClusterInfo, error)
	Remove(ctx context.Context, in *ClusterInfo, opts ...grpc.CallOption) (*Empty, error)
	GetDriverCreateOptions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DriverFlags, error)
	GetDriverUpdateOptions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DriverFlags, error)
	GetVersion(ctx context.Context, in *ClusterInfo, opts ...grpc.CallOption) (*KubernetesVersion, error)
	SetVersion(ctx context.Context, in *SetVersionRequest, opts ...grpc.CallOption) (*Empty, error)
	GetNodeCount(ctx context.Context, in *ClusterInfo, opts ...grpc.CallOption) (*NodeCount, error)
	SetNodeCount(ctx context.Context, in *SetNodeCountRequest, opts ...grpc.CallOption) (*Empty, error)
	GetCapabilities(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Capabilities, error)
}

type driverClient struct {
	cc *grpc.ClientConn
}

func NewDriverClient(cc *grpc.ClientConn) DriverClient {
	return &driverClient{cc}
}

func (c *driverClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*ClusterInfo, error) {
	out := new(ClusterInfo)
	err := grpc.Invoke(ctx, "/types.Driver/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*ClusterInfo, error) {
	out := new(ClusterInfo)
	err := grpc.Invoke(ctx, "/types.Driver/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) PostCheck(ctx context.Context, in *ClusterInfo, opts ...grpc.CallOption) (*ClusterInfo, error) {
	out := new(ClusterInfo)
	err := grpc.Invoke(ctx, "/types.Driver/PostCheck", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) Remove(ctx context.Context, in *ClusterInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/types.Driver/Remove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) GetDriverCreateOptions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DriverFlags, error) {
	out := new(DriverFlags)
	err := grpc.Invoke(ctx, "/types.Driver/GetDriverCreateOptions", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) GetDriverUpdateOptions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DriverFlags, error) {
	out := new(DriverFlags)
	err := grpc.Invoke(ctx, "/types.Driver/GetDriverUpdateOptions", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) GetVersion(ctx context.Context, in *ClusterInfo, opts ...grpc.CallOption) (*KubernetesVersion, error) {
	out := new(KubernetesVersion)
	err := grpc.Invoke(ctx, "/types.Driver/GetVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) SetVersion(ctx context.Context, in *SetVersionRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/types.Driver/SetVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) GetNodeCount(ctx context.Context, in *ClusterInfo, opts ...grpc.CallOption) (*NodeCount, error) {
	out := new(NodeCount)
	err := grpc.Invoke(ctx, "/types.Driver/GetNodeCount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) SetNodeCount(ctx context.Context, in *SetNodeCountRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/types.Driver/SetNodeCount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) GetCapabilities(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Capabilities, error) {
	out := new(Capabilities)
	err := grpc.Invoke(ctx, "/types.Driver/GetCapabilities", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Driver service

type DriverServer interface {
	Create(context.Context, *CreateRequest) (*ClusterInfo, error)
	Update(context.Context, *UpdateRequest) (*ClusterInfo, error)
	PostCheck(context.Context, *ClusterInfo) (*ClusterInfo, error)
	Remove(context.Context, *ClusterInfo) (*Empty, error)
	GetDriverCreateOptions(context.Context, *Empty) (*DriverFlags, error)
	GetDriverUpdateOptions(context.Context, *Empty) (*DriverFlags, error)
	GetVersion(context.Context, *ClusterInfo) (*KubernetesVersion, error)
	SetVersion(context.Context, *SetVersionRequest) (*Empty, error)
	GetNodeCount(context.Context, *ClusterInfo) (*NodeCount, error)
	SetNodeCount(context.Context, *SetNodeCountRequest) (*Empty, error)
	GetCapabilities(context.Context, *Empty) (*Capabilities, error)
}

func RegisterDriverServer(s *grpc.Server, srv DriverServer) {
	s.RegisterService(&_Driver_serviceDesc, srv)
}

func _Driver_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Driver/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Driver/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_PostCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).PostCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Driver/PostCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).PostCheck(ctx, req.(*ClusterInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Driver/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Remove(ctx, req.(*ClusterInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_GetDriverCreateOptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).GetDriverCreateOptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Driver/GetDriverCreateOptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).GetDriverCreateOptions(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_GetDriverUpdateOptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).GetDriverUpdateOptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Driver/GetDriverUpdateOptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).GetDriverUpdateOptions(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Driver/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).GetVersion(ctx, req.(*ClusterInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_SetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).SetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Driver/SetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).SetVersion(ctx, req.(*SetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_GetNodeCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).GetNodeCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Driver/GetNodeCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).GetNodeCount(ctx, req.(*ClusterInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_SetNodeCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetNodeCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).SetNodeCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Driver/SetNodeCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).SetNodeCount(ctx, req.(*SetNodeCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_GetCapabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).GetCapabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Driver/GetCapabilities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).GetCapabilities(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Driver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.Driver",
	HandlerType: (*DriverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Driver_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Driver_Update_Handler,
		},
		{
			MethodName: "PostCheck",
			Handler:    _Driver_PostCheck_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Driver_Remove_Handler,
		},
		{
			MethodName: "GetDriverCreateOptions",
			Handler:    _Driver_GetDriverCreateOptions_Handler,
		},
		{
			MethodName: "GetDriverUpdateOptions",
			Handler:    _Driver_GetDriverUpdateOptions_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _Driver_GetVersion_Handler,
		},
		{
			MethodName: "SetVersion",
			Handler:    _Driver_SetVersion_Handler,
		},
		{
			MethodName: "GetNodeCount",
			Handler:    _Driver_GetNodeCount_Handler,
		},
		{
			MethodName: "SetNodeCount",
			Handler:    _Driver_SetNodeCount_Handler,
		},
		{
			MethodName: "GetCapabilities",
			Handler:    _Driver_GetCapabilities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "drivers.proto",
}

func init() { proto.RegisterFile("drivers.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1025 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0x5b, 0x73, 0xdb, 0x44,
	0x14, 0xb6, 0xe2, 0xfb, 0x91, 0x9c, 0x26, 0x27, 0xa6, 0x68, 0x3c, 0x03, 0x38, 0xa2, 0x14, 0x3f,
	0xb4, 0x7e, 0x30, 0xb4, 0x03, 0x6d, 0xb9, 0xd5, 0x04, 0x93, 0xe9, 0x50, 0x18, 0x05, 0x98, 0xe1,
	0x05, 0x8f, 0x2c, 0x6f, 0x82, 0x26, 0x8e, 0xd6, 0x68, 0x57, 0x61, 0xfc, 0x0a, 0xbf, 0x80, 0x57,
	0xfe, 0x07, 0xff, 0x8a, 0x77, 0x5e, 0x99, 0xbd, 0x48, 0x5e, 0xd9, 0x32, 0x49, 0xde, 0xbc, 0xe7,
	0x7c, 0xe7, 0x3b, 0xb7, 0x6f, 0xd7, 0x82, 0xce, 0x3c, 0x89, 0xae, 0x49, 0xc2, 0x86, 0xcb, 0x84,
	0x72, 0x8a, 0x75, 0xbe, 0x5a, 0x12, 0xe6, 0x35, 0xa1, 0x7e, 0x72, 0xb5, 0xe4, 0x2b, 0xef, 0x4f,
	0x0b, 0xec, 0x2f, 0x25, 0xe2, 0xab, 0x45, 0x70, 0xc1, 0xf0, 0x63, 0x68, 0xd2, 0x25, 0x8f, 0x68,
	0xcc, 0x5c, 0xab, 0x5f, 0x1d, 0xd8, 0xa3, 0x77, 0x86, 0x32, 0x62, 0x68, 0x80, 0x86, 0xdf, 0x2a,
	0xc4, 0x49, 0xcc, 0x93, 0x95, 0x9f, 0xe1, 0x7b, 0x13, 0x70, 0x4c, 0x07, 0x1e, 0x40, 0xf5, 0x92,
	0xac, 0x5c, 0xab, 0x6f, 0x0d, 0xda, 0xbe, 0xf8, 0x89, 0xc7, 0x50, 0xbf, 0x0e, 0x16, 0x29, 0x71,
	0xf7, 0xfa, 0xd6, 0xc0, 0x1e, 0xd9, 0x9a, 0x5a, 0x90, 0xfa, 0xca, 0xf3, 0x6c, 0xef, 0x23, 0xcb,
	0x5b, 0x42, 0x4d, 0x98, 0x10, 0xa1, 0x26, 0x00, 0x9a, 0x41, 0xfe, 0xc6, 0x2e, 0xd4, 0x53, 0x16,
	0x5c, 0x28, 0x8a, 0xb6, 0xaf, 0x0e, 0xc2, 0xaa, 0x88, 0xab, 0xca, 0x2a, 0x0f, 0x38, 0x80, 0xe6,
	0x9c, 0x9c, 0x07, 0xe9, 0x82, 0xbb, 0x35, 0x99, 0x70, 0x3f, 0xeb, 0x45, 0x59, 0xfd, 0xcc, 0xed,
	0xfd, 0x6d, 0x41, 0x53, 0x1b, 0xb1, 0x0f, 0xb6, 0x36, 0xbf, 0xa4, 0x74, 0x21, 0x93, 0xb7, 0x7c,
	0xd3, 0x84, 0x0f, 0xa0, 0xa3, 0x8f, 0x67, 0x3c, 0x89, 0xe2, 0x0b, 0x5d, 0x4b, 0xd1, 0x88, 0x2f,
	0x01, 0x0b, 0x86, 0xb3, 0x45, 0x14, 0xaa, 0x02, 0xed, 0x11, 0xea, 0x42, 0x0c, 0x8f, 0x5f, 0x82,
	0xc6, 0xb7, 0x01, 0xb4, 0xf5, 0x34, 0x56, 0x4d, 0x54, 0x7d, 0xc3, 0xe2, 0xfd, 0x53, 0x83, 0x8e,
	0x5a, 0x8c, 0x9e, 0x3c, 0x7e, 0x0d, 0xce, 0x8c, 0xd2, 0xc5, 0xb4, 0xb8, 0xc4, 0xf7, 0x0a, 0x4b,
	0xd4, 0xd8, 0xa1, 0x68, 0xa6, 0xb0, 0x4a, 0x7b, 0xb6, 0xb6, 0xe0, 0x6b, 0xd8, 0x67, 0xb2, 0x94,
	0x9c, 0x6b, 0x4f, 0x72, 0xbd, 0x5f, 0xca, 0xa5, 0xaa, 0x2e, 0xb0, 0x75, 0x98, 0x69, 0xc3, 0x13,
	0xb0, 0xa3, 0x98, 0xe7, 0x64, 0x55, 0x49, 0xf6, 0xa0, 0x94, 0xec, 0x34, 0xe6, 0x05, 0x26, 0x88,
	0x72, 0x03, 0xfe, 0x0c, 0x5d, 0x5d, 0x16, 0x13, 0x23, 0xca, 0xf9, 0x6a, 0x92, 0xef, 0xd1, 0xff,
	0x14, 0x27, 0x47, 0x5a, 0xe0, 0x45, 0xb6, 0xe5, 0xe8, 0x7d, 0x0a, 0x07, 0x9b, 0x73, 0x29, 0x51,
	0x72, 0xd7, 0x54, 0x72, 0xcb, 0x10, 0x6f, 0xef, 0x73, 0xc0, 0xed, 0x59, 0xdc, 0xc4, 0xd0, 0x36,
	0x19, 0x3e, 0x81, 0x7b, 0x1b, 0x03, 0xb8, 0x29, 0xbc, 0x6a, 0x86, 0xff, 0x04, 0x6f, 0xee, 0xe8,
	0xb7, 0x84, 0x66, 0x50, 0xbc, 0x91, 0x65, 0xba, 0x34, 0x2e, 0xe6, 0xbb, 0x60, 0x9b, 0xea, 0xcc,
	0x6b, 0x10, 0x22, 0xcb, 0x5a, 0xf0, 0x7e, 0xaf, 0x81, 0x3d, 0x5e, 0xa4, 0x8c, 0x93, 0xe4, 0x34,
	0x3e, 0xa7, 0xe8, 0x42, 0x53, 0xbc, 0x3f, 0x11, 0x8d, 0x75, 0xe2, 0xec, 0x88, 0x23, 0x78, 0x83,
	0x91, 0xe4, 0x5a, 0x6c, 0x31, 0x08, 0x43, 0x9a, 0xc6, 0x7c, 0xca, 0xe9, 0x25, 0x89, 0xf5, 0x48,
	0x8e, 0xb4, 0xf3, 0x0b, 0xe5, 0xfb, 0x5e, 0xb8, 0xb0, 0x07, 0x2d, 0x12, 0xcf, 0x97, 0x34, 0x8a,
	0xb9, 0xbe, 0xec, 0xf9, 0x59, 0xf8, 0x52, 0x46, 0x92, 0x38, 0xb8, 0x22, 0xf2, 0xae, 0xb4, 0xfd,
	0xfc, 0x2c, 0x7c, 0xcb, 0x80, 0xb1, 0xdf, 0x68, 0x32, 0x77, 0xeb, 0xca, 0x97, 0x9d, 0x71, 0x08,
	0x47, 0x09, 0xa5, 0x7c, 0x1a, 0x06, 0xd3, 0x90, 0x24, 0x3c, 0x3a, 0x8f, 0xc2, 0x80, 0x13, 0xb7,
	0x21, 0x61, 0x87, 0xc2, 0x35, 0x0e, 0xc6, 0x6b, 0x07, 0x3e, 0x06, 0x0c, 0x17, 0x11, 0x89, 0x79,
	0x01, 0xde, 0x54, 0x70, 0xe5, 0x31, 0xe1, 0x6f, 0x01, 0x68, 0xb8, 0x18, 0x7e, 0x4b, 0xc2, 0xda,
	0xca, 0xf2, 0x8a, 0xac, 0x84, 0x3b, 0xa6, 0x73, 0x32, 0x95, 0x4d, 0xba, 0x6d, 0xb9, 0xce, 0xb6,
	0xb0, 0x8c, 0x85, 0x01, 0x5f, 0x40, 0xeb, 0x8a, 0xf0, 0x60, 0x1e, 0xf0, 0xc0, 0x05, 0xa9, 0xf1,
	0xbe, 0x5e, 0x92, 0x31, 0xe4, 0xe1, 0x37, 0x1a, 0xa2, 0x74, 0x9d, 0x47, 0xe0, 0x7d, 0x68, 0x30,
	0x1e, 0xf0, 0x94, 0xb9, 0xb6, 0xcc, 0xab, 0x4f, 0x78, 0x0c, 0x4e, 0x98, 0x90, 0x80, 0x93, 0x29,
	0x49, 0x12, 0x9a, 0xb8, 0x8e, 0xf4, 0xda, 0xca, 0x76, 0x22, 0x4c, 0xbd, 0xe7, 0xd0, 0x29, 0xb0,
	0xde, 0x45, 0xc3, 0xde, 0x63, 0x38, 0x7c, 0x95, 0xce, 0x48, 0x12, 0x13, 0x4e, 0xd8, 0x8f, 0x7a,
	0xdf, 0x3b, 0x95, 0xe0, 0x1d, 0x43, 0xfb, 0x75, 0xde, 0x71, 0x17, 0xea, 0x6a, 0x16, 0x96, 0x92,
	0xb6, 0x3c, 0x78, 0x7f, 0x59, 0xe0, 0x8c, 0x83, 0x65, 0x30, 0x8b, 0x16, 0x11, 0x8f, 0x08, 0xc3,
	0x53, 0x70, 0x42, 0xe3, 0xbc, 0xf1, 0xd2, 0x99, 0xd0, 0xc2, 0x41, 0x4d, 0xa8, 0x10, 0xda, 0xfb,
	0x0c, 0x0e, 0xb7, 0x20, 0x66, 0xbb, 0xd5, 0x1b, 0x2e, 0xbd, 0xf7, 0x87, 0x05, 0x9d, 0xb1, 0x9c,
	0x9d, 0x4f, 0x7e, 0x4d, 0x09, 0xe3, 0xf8, 0x1c, 0xf6, 0xd5, 0x1f, 0xaf, 0xf1, 0x12, 0x8b, 0x1b,
	0xd6, 0x2d, 0x7b, 0xa0, 0x7c, 0xfd, 0x27, 0x9d, 0xbd, 0x71, 0x4f, 0xc0, 0x09, 0xd5, 0x72, 0xa7,
	0x51, 0x7c, 0x4e, 0x37, 0x2e, 0xa7, 0xb1, 0x77, 0xdf, 0x0e, 0xd7, 0x07, 0x59, 0xc5, 0x0f, 0xcb,
	0xb9, 0x51, 0xc5, 0x26, 0x91, 0x75, 0x2b, 0xa2, 0x92, 0xe2, 0xf7, 0x6e, 0x5d, 0xbc, 0x47, 0xe1,
	0xf0, 0x8c, 0x70, 0xbd, 0xf3, 0xac, 0x90, 0x87, 0x50, 0xbb, 0xa1, 0x00, 0xe9, 0xc7, 0xd1, 0x5a,
	0x22, 0x2a, 0xa5, 0xab, 0xa1, 0x5b, 0x6a, 0x5a, 0x8b, 0x87, 0xc0, 0xd1, 0x19, 0xe1, 0xb9, 0x7e,
	0xee, 0x9a, 0xf2, 0x61, 0x26, 0x37, 0x95, 0xf0, 0x40, 0x03, 0xd7, 0x7c, 0xca, 0x3d, 0xfa, 0xb7,
	0x06, 0x0d, 0xd5, 0x38, 0x7e, 0x08, 0x0d, 0xb5, 0x6d, 0xcc, 0x26, 0x52, 0x58, 0x7e, 0xaf, 0x24,
	0x99, 0x57, 0x11, 0x51, 0x6a, 0x3b, 0x79, 0x54, 0x61, 0x59, 0x3b, 0xa2, 0x9e, 0x40, 0xfb, 0x3b,
	0xca, 0xf8, 0xf8, 0x17, 0x12, 0x5e, 0x62, 0x09, 0x64, 0x47, 0xd8, 0x23, 0x68, 0xf8, 0xe4, 0x8a,
	0x5e, 0x93, 0xd2, 0x18, 0x47, 0xdb, 0xd4, 0x37, 0x60, 0x05, 0x5f, 0xc0, 0xfd, 0x09, 0xe1, 0xaa,
	0x3b, 0xd5, 0x4a, 0x26, 0xc5, 0x02, 0x32, 0xcf, 0x65, 0x7c, 0x0c, 0x6e, 0x44, 0xab, 0x96, 0xee,
	0x16, 0x0d, 0x93, 0x5c, 0x2f, 0xa5, 0xd5, 0xee, 0xd4, 0x80, 0x57, 0xc1, 0xa7, 0x00, 0x6b, 0xb5,
	0x61, 0x86, 0xdc, 0x12, 0xe0, 0x56, 0xc7, 0x4f, 0xc1, 0x99, 0x18, 0xa2, 0x29, 0xcd, 0xbb, 0x25,
	0x05, 0xaf, 0x82, 0xcf, 0xc0, 0x31, 0xc5, 0x86, 0xbd, 0x75, 0xc6, 0x4d, 0x05, 0x96, 0xe4, 0xbc,
	0x37, 0x21, 0xbc, 0xf0, 0x88, 0x15, 0x07, 0x74, 0x54, 0xf2, 0x78, 0x79, 0x95, 0x59, 0x43, 0x7e,
	0xba, 0x7f, 0xf0, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0xf1, 0xdc, 0xb9, 0xcb, 0x0b, 0x00,
	0x00,
}