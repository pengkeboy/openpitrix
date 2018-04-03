// Code generated by protoc-gen-go. DO NOT EDIT.
// source: drone.proto

/*
Package openpitrix_drone is a generated protocol buffer package.

It is generated from these files:
	drone.proto

It has these top-level messages:
	Info
	ConfdConfig
	ConfdBackendConfig
	StartConfdRequest
	ConfdStatus
	SubscribeCmdStatusRequest
	SubscribeCmdStatusResponse
	GetRegisterCmdStatusRequest
	GetRegisterCmdStatusResponse
	GetTemplateFilesRequest
	GetTemplateFilesResponse
	GetValuesRequest
	GetValuesResponse
	Empty
*/
package openpitrix_drone

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import google_protobuf1 "github.com/golang/protobuf/ptypes/wrappers"

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

type Info struct {
	DroneIp            string              `protobuf:"bytes,1,opt,name=drone_ip,json=droneIp" json:"drone_ip,omitempty"`
	ConfdConfig        *ConfdConfig        `protobuf:"bytes,2,opt,name=confd_config,json=confdConfig" json:"confd_config,omitempty"`
	ConfdBackendConfig *ConfdBackendConfig `protobuf:"bytes,3,opt,name=confd_backend_config,json=confdBackendConfig" json:"confd_backend_config,omitempty"`
}

func (m *Info) Reset()                    { *m = Info{} }
func (m *Info) String() string            { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()               {}
func (*Info) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Info) GetDroneIp() string {
	if m != nil {
		return m.DroneIp
	}
	return ""
}

func (m *Info) GetConfdConfig() *ConfdConfig {
	if m != nil {
		return m.ConfdConfig
	}
	return nil
}

func (m *Info) GetConfdBackendConfig() *ConfdBackendConfig {
	if m != nil {
		return m.ConfdBackendConfig
	}
	return nil
}

// See https://godoc.org/openpitrix.io/libconfd#Config
type ConfdConfig struct {
	ConfDir  *google_protobuf1.StringValue `protobuf:"bytes,1,opt,name=conf_dir,json=confDir" json:"conf_dir,omitempty"`
	Interval *google_protobuf1.Int32Value  `protobuf:"bytes,2,opt,name=interval" json:"interval,omitempty"`
	Noop     *google_protobuf1.BoolValue   `protobuf:"bytes,3,opt,name=noop" json:"noop,omitempty"`
	Prefix   *google_protobuf1.StringValue `protobuf:"bytes,4,opt,name=prefix" json:"prefix,omitempty"`
	SyncOnly *google_protobuf1.BoolValue   `protobuf:"bytes,5,opt,name=sync_only,json=syncOnly" json:"sync_only,omitempty"`
	LogLevel *google_protobuf1.StringValue `protobuf:"bytes,6,opt,name=log_level,json=logLevel" json:"log_level,omitempty"`
}

func (m *ConfdConfig) Reset()                    { *m = ConfdConfig{} }
func (m *ConfdConfig) String() string            { return proto.CompactTextString(m) }
func (*ConfdConfig) ProtoMessage()               {}
func (*ConfdConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ConfdConfig) GetConfDir() *google_protobuf1.StringValue {
	if m != nil {
		return m.ConfDir
	}
	return nil
}

func (m *ConfdConfig) GetInterval() *google_protobuf1.Int32Value {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *ConfdConfig) GetNoop() *google_protobuf1.BoolValue {
	if m != nil {
		return m.Noop
	}
	return nil
}

func (m *ConfdConfig) GetPrefix() *google_protobuf1.StringValue {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *ConfdConfig) GetSyncOnly() *google_protobuf1.BoolValue {
	if m != nil {
		return m.SyncOnly
	}
	return nil
}

func (m *ConfdConfig) GetLogLevel() *google_protobuf1.StringValue {
	if m != nil {
		return m.LogLevel
	}
	return nil
}

// See https://godoc.org/openpitrix.io/libconfd#BackendConfig
type ConfdBackendConfig struct {
	Type         *google_protobuf1.StringValue `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Host         []string                      `protobuf:"bytes,2,rep,name=host" json:"host,omitempty"`
	Username     *google_protobuf1.StringValue `protobuf:"bytes,3,opt,name=username" json:"username,omitempty"`
	Password     *google_protobuf1.StringValue `protobuf:"bytes,4,opt,name=password" json:"password,omitempty"`
	ClientCaKeys *google_protobuf1.StringValue `protobuf:"bytes,5,opt,name=client_ca_keys,json=clientCaKeys" json:"client_ca_keys,omitempty"`
	ClientCert   *google_protobuf1.StringValue `protobuf:"bytes,6,opt,name=client_cert,json=clientCert" json:"client_cert,omitempty"`
	ClientKey    *google_protobuf1.StringValue `protobuf:"bytes,7,opt,name=client_key,json=clientKey" json:"client_key,omitempty"`
}

func (m *ConfdBackendConfig) Reset()                    { *m = ConfdBackendConfig{} }
func (m *ConfdBackendConfig) String() string            { return proto.CompactTextString(m) }
func (*ConfdBackendConfig) ProtoMessage()               {}
func (*ConfdBackendConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ConfdBackendConfig) GetType() *google_protobuf1.StringValue {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *ConfdBackendConfig) GetHost() []string {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *ConfdBackendConfig) GetUsername() *google_protobuf1.StringValue {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *ConfdBackendConfig) GetPassword() *google_protobuf1.StringValue {
	if m != nil {
		return m.Password
	}
	return nil
}

func (m *ConfdBackendConfig) GetClientCaKeys() *google_protobuf1.StringValue {
	if m != nil {
		return m.ClientCaKeys
	}
	return nil
}

func (m *ConfdBackendConfig) GetClientCert() *google_protobuf1.StringValue {
	if m != nil {
		return m.ClientCert
	}
	return nil
}

func (m *ConfdBackendConfig) GetClientKey() *google_protobuf1.StringValue {
	if m != nil {
		return m.ClientKey
	}
	return nil
}

type StartConfdRequest struct {
	ConfdConfig        *ConfdConfig        `protobuf:"bytes,1,opt,name=confd_config,json=confdConfig" json:"confd_config,omitempty"`
	ConfdBackendConfig *ConfdBackendConfig `protobuf:"bytes,2,opt,name=confd_backend_config,json=confdBackendConfig" json:"confd_backend_config,omitempty"`
}

func (m *StartConfdRequest) Reset()                    { *m = StartConfdRequest{} }
func (m *StartConfdRequest) String() string            { return proto.CompactTextString(m) }
func (*StartConfdRequest) ProtoMessage()               {}
func (*StartConfdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StartConfdRequest) GetConfdConfig() *ConfdConfig {
	if m != nil {
		return m.ConfdConfig
	}
	return nil
}

func (m *StartConfdRequest) GetConfdBackendConfig() *ConfdBackendConfig {
	if m != nil {
		return m.ConfdBackendConfig
	}
	return nil
}

type ConfdStatus struct {
	ProcessId string                     `protobuf:"bytes,1,opt,name=process_id,json=processId" json:"process_id,omitempty"`
	Status    string                     `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	UpTime    *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=up_time,json=upTime" json:"up_time,omitempty"`
}

func (m *ConfdStatus) Reset()                    { *m = ConfdStatus{} }
func (m *ConfdStatus) String() string            { return proto.CompactTextString(m) }
func (*ConfdStatus) ProtoMessage()               {}
func (*ConfdStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ConfdStatus) GetProcessId() string {
	if m != nil {
		return m.ProcessId
	}
	return ""
}

func (m *ConfdStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ConfdStatus) GetUpTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpTime
	}
	return nil
}

type SubscribeCmdStatusRequest struct {
	SubtaskId string `protobuf:"bytes,1,opt,name=subtask_id,json=subtaskId" json:"subtask_id,omitempty"`
}

func (m *SubscribeCmdStatusRequest) Reset()                    { *m = SubscribeCmdStatusRequest{} }
func (m *SubscribeCmdStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*SubscribeCmdStatusRequest) ProtoMessage()               {}
func (*SubscribeCmdStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SubscribeCmdStatusRequest) GetSubtaskId() string {
	if m != nil {
		return m.SubtaskId
	}
	return ""
}

type SubscribeCmdStatusResponse struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *SubscribeCmdStatusResponse) Reset()                    { *m = SubscribeCmdStatusResponse{} }
func (m *SubscribeCmdStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*SubscribeCmdStatusResponse) ProtoMessage()               {}
func (*SubscribeCmdStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SubscribeCmdStatusResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type GetRegisterCmdStatusRequest struct {
	SubtaskId string `protobuf:"bytes,1,opt,name=subtask_id,json=subtaskId" json:"subtask_id,omitempty"`
}

func (m *GetRegisterCmdStatusRequest) Reset()                    { *m = GetRegisterCmdStatusRequest{} }
func (m *GetRegisterCmdStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRegisterCmdStatusRequest) ProtoMessage()               {}
func (*GetRegisterCmdStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetRegisterCmdStatusRequest) GetSubtaskId() string {
	if m != nil {
		return m.SubtaskId
	}
	return ""
}

type GetRegisterCmdStatusResponse struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *GetRegisterCmdStatusResponse) Reset()                    { *m = GetRegisterCmdStatusResponse{} }
func (m *GetRegisterCmdStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*GetRegisterCmdStatusResponse) ProtoMessage()               {}
func (*GetRegisterCmdStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetRegisterCmdStatusResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type GetTemplateFilesRequest struct {
	Regexp string `protobuf:"bytes,1,opt,name=regexp" json:"regexp,omitempty"`
}

func (m *GetTemplateFilesRequest) Reset()                    { *m = GetTemplateFilesRequest{} }
func (m *GetTemplateFilesRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTemplateFilesRequest) ProtoMessage()               {}
func (*GetTemplateFilesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetTemplateFilesRequest) GetRegexp() string {
	if m != nil {
		return m.Regexp
	}
	return ""
}

type GetTemplateFilesResponse struct {
	Files []string `protobuf:"bytes,1,rep,name=files" json:"files,omitempty"`
}

func (m *GetTemplateFilesResponse) Reset()                    { *m = GetTemplateFilesResponse{} }
func (m *GetTemplateFilesResponse) String() string            { return proto.CompactTextString(m) }
func (*GetTemplateFilesResponse) ProtoMessage()               {}
func (*GetTemplateFilesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *GetTemplateFilesResponse) GetFiles() []string {
	if m != nil {
		return m.Files
	}
	return nil
}

type GetValuesRequest struct {
	Keys []string `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
}

func (m *GetValuesRequest) Reset()                    { *m = GetValuesRequest{} }
func (m *GetValuesRequest) String() string            { return proto.CompactTextString(m) }
func (*GetValuesRequest) ProtoMessage()               {}
func (*GetValuesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *GetValuesRequest) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type GetValuesResponse struct {
	Values map[string]string `protobuf:"bytes,1,rep,name=values" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *GetValuesResponse) Reset()                    { *m = GetValuesResponse{} }
func (m *GetValuesResponse) String() string            { return proto.CompactTextString(m) }
func (*GetValuesResponse) ProtoMessage()               {}
func (*GetValuesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *GetValuesResponse) GetValues() map[string]string {
	if m != nil {
		return m.Values
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func init() {
	proto.RegisterType((*Info)(nil), "openpitrix.drone.Info")
	proto.RegisterType((*ConfdConfig)(nil), "openpitrix.drone.ConfdConfig")
	proto.RegisterType((*ConfdBackendConfig)(nil), "openpitrix.drone.ConfdBackendConfig")
	proto.RegisterType((*StartConfdRequest)(nil), "openpitrix.drone.StartConfdRequest")
	proto.RegisterType((*ConfdStatus)(nil), "openpitrix.drone.ConfdStatus")
	proto.RegisterType((*SubscribeCmdStatusRequest)(nil), "openpitrix.drone.SubscribeCmdStatusRequest")
	proto.RegisterType((*SubscribeCmdStatusResponse)(nil), "openpitrix.drone.SubscribeCmdStatusResponse")
	proto.RegisterType((*GetRegisterCmdStatusRequest)(nil), "openpitrix.drone.GetRegisterCmdStatusRequest")
	proto.RegisterType((*GetRegisterCmdStatusResponse)(nil), "openpitrix.drone.GetRegisterCmdStatusResponse")
	proto.RegisterType((*GetTemplateFilesRequest)(nil), "openpitrix.drone.GetTemplateFilesRequest")
	proto.RegisterType((*GetTemplateFilesResponse)(nil), "openpitrix.drone.GetTemplateFilesResponse")
	proto.RegisterType((*GetValuesRequest)(nil), "openpitrix.drone.GetValuesRequest")
	proto.RegisterType((*GetValuesResponse)(nil), "openpitrix.drone.GetValuesResponse")
	proto.RegisterType((*Empty)(nil), "openpitrix.drone.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DroneService service

type DroneServiceClient interface {
	GetInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Info, error)
	GetConfdConfig(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ConfdConfig, error)
	GetBackendConfig(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ConfdBackendConfig, error)
	StartConfd(ctx context.Context, in *StartConfdRequest, opts ...grpc.CallOption) (*Empty, error)
	StopConfd(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	GetConfdStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ConfdStatus, error)
	GetTemplateFiles(ctx context.Context, in *GetTemplateFilesRequest, opts ...grpc.CallOption) (*GetTemplateFilesResponse, error)
	GetValues(ctx context.Context, in *GetValuesRequest, opts ...grpc.CallOption) (*GetValuesResponse, error)
}

type droneServiceClient struct {
	cc *grpc.ClientConn
}

func NewDroneServiceClient(cc *grpc.ClientConn) DroneServiceClient {
	return &droneServiceClient{cc}
}

func (c *droneServiceClient) GetInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Info, error) {
	out := new(Info)
	err := grpc.Invoke(ctx, "/openpitrix.drone.DroneService/GetInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) GetConfdConfig(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ConfdConfig, error) {
	out := new(ConfdConfig)
	err := grpc.Invoke(ctx, "/openpitrix.drone.DroneService/GetConfdConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) GetBackendConfig(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ConfdBackendConfig, error) {
	out := new(ConfdBackendConfig)
	err := grpc.Invoke(ctx, "/openpitrix.drone.DroneService/GetBackendConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) StartConfd(ctx context.Context, in *StartConfdRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/openpitrix.drone.DroneService/StartConfd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) StopConfd(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/openpitrix.drone.DroneService/StopConfd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) GetConfdStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ConfdStatus, error) {
	out := new(ConfdStatus)
	err := grpc.Invoke(ctx, "/openpitrix.drone.DroneService/GetConfdStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) GetTemplateFiles(ctx context.Context, in *GetTemplateFilesRequest, opts ...grpc.CallOption) (*GetTemplateFilesResponse, error) {
	out := new(GetTemplateFilesResponse)
	err := grpc.Invoke(ctx, "/openpitrix.drone.DroneService/GetTemplateFiles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) GetValues(ctx context.Context, in *GetValuesRequest, opts ...grpc.CallOption) (*GetValuesResponse, error) {
	out := new(GetValuesResponse)
	err := grpc.Invoke(ctx, "/openpitrix.drone.DroneService/GetValues", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DroneService service

type DroneServiceServer interface {
	GetInfo(context.Context, *Empty) (*Info, error)
	GetConfdConfig(context.Context, *Empty) (*ConfdConfig, error)
	GetBackendConfig(context.Context, *Empty) (*ConfdBackendConfig, error)
	StartConfd(context.Context, *StartConfdRequest) (*Empty, error)
	StopConfd(context.Context, *Empty) (*Empty, error)
	GetConfdStatus(context.Context, *Empty) (*ConfdStatus, error)
	GetTemplateFiles(context.Context, *GetTemplateFilesRequest) (*GetTemplateFilesResponse, error)
	GetValues(context.Context, *GetValuesRequest) (*GetValuesResponse, error)
}

func RegisterDroneServiceServer(s *grpc.Server, srv DroneServiceServer) {
	s.RegisterService(&_DroneService_serviceDesc, srv)
}

func _DroneService_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.drone.DroneService/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).GetInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_GetConfdConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).GetConfdConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.drone.DroneService/GetConfdConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).GetConfdConfig(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_GetBackendConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).GetBackendConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.drone.DroneService/GetBackendConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).GetBackendConfig(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_StartConfd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartConfdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).StartConfd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.drone.DroneService/StartConfd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).StartConfd(ctx, req.(*StartConfdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_StopConfd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).StopConfd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.drone.DroneService/StopConfd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).StopConfd(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_GetConfdStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).GetConfdStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.drone.DroneService/GetConfdStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).GetConfdStatus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_GetTemplateFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTemplateFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).GetTemplateFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.drone.DroneService/GetTemplateFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).GetTemplateFiles(ctx, req.(*GetTemplateFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_GetValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).GetValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.drone.DroneService/GetValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).GetValues(ctx, req.(*GetValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DroneService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.drone.DroneService",
	HandlerType: (*DroneServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfo",
			Handler:    _DroneService_GetInfo_Handler,
		},
		{
			MethodName: "GetConfdConfig",
			Handler:    _DroneService_GetConfdConfig_Handler,
		},
		{
			MethodName: "GetBackendConfig",
			Handler:    _DroneService_GetBackendConfig_Handler,
		},
		{
			MethodName: "StartConfd",
			Handler:    _DroneService_StartConfd_Handler,
		},
		{
			MethodName: "StopConfd",
			Handler:    _DroneService_StopConfd_Handler,
		},
		{
			MethodName: "GetConfdStatus",
			Handler:    _DroneService_GetConfdStatus_Handler,
		},
		{
			MethodName: "GetTemplateFiles",
			Handler:    _DroneService_GetTemplateFiles_Handler,
		},
		{
			MethodName: "GetValues",
			Handler:    _DroneService_GetValues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "drone.proto",
}

// Client API for DroneServiceForFrontgate service

type DroneServiceForFrontgateClient interface {
	SubscribeCmdStatus(ctx context.Context, in *SubscribeCmdStatusRequest, opts ...grpc.CallOption) (DroneServiceForFrontgate_SubscribeCmdStatusClient, error)
	GetRegisterCmdStatus(ctx context.Context, in *GetRegisterCmdStatusRequest, opts ...grpc.CallOption) (*GetRegisterCmdStatusResponse, error)
}

type droneServiceForFrontgateClient struct {
	cc *grpc.ClientConn
}

func NewDroneServiceForFrontgateClient(cc *grpc.ClientConn) DroneServiceForFrontgateClient {
	return &droneServiceForFrontgateClient{cc}
}

func (c *droneServiceForFrontgateClient) SubscribeCmdStatus(ctx context.Context, in *SubscribeCmdStatusRequest, opts ...grpc.CallOption) (DroneServiceForFrontgate_SubscribeCmdStatusClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DroneServiceForFrontgate_serviceDesc.Streams[0], c.cc, "/openpitrix.drone.DroneServiceForFrontgate/SubscribeCmdStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &droneServiceForFrontgateSubscribeCmdStatusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DroneServiceForFrontgate_SubscribeCmdStatusClient interface {
	Recv() (*SubscribeCmdStatusResponse, error)
	grpc.ClientStream
}

type droneServiceForFrontgateSubscribeCmdStatusClient struct {
	grpc.ClientStream
}

func (x *droneServiceForFrontgateSubscribeCmdStatusClient) Recv() (*SubscribeCmdStatusResponse, error) {
	m := new(SubscribeCmdStatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *droneServiceForFrontgateClient) GetRegisterCmdStatus(ctx context.Context, in *GetRegisterCmdStatusRequest, opts ...grpc.CallOption) (*GetRegisterCmdStatusResponse, error) {
	out := new(GetRegisterCmdStatusResponse)
	err := grpc.Invoke(ctx, "/openpitrix.drone.DroneServiceForFrontgate/GetRegisterCmdStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DroneServiceForFrontgate service

type DroneServiceForFrontgateServer interface {
	SubscribeCmdStatus(*SubscribeCmdStatusRequest, DroneServiceForFrontgate_SubscribeCmdStatusServer) error
	GetRegisterCmdStatus(context.Context, *GetRegisterCmdStatusRequest) (*GetRegisterCmdStatusResponse, error)
}

func RegisterDroneServiceForFrontgateServer(s *grpc.Server, srv DroneServiceForFrontgateServer) {
	s.RegisterService(&_DroneServiceForFrontgate_serviceDesc, srv)
}

func _DroneServiceForFrontgate_SubscribeCmdStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeCmdStatusRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DroneServiceForFrontgateServer).SubscribeCmdStatus(m, &droneServiceForFrontgateSubscribeCmdStatusServer{stream})
}

type DroneServiceForFrontgate_SubscribeCmdStatusServer interface {
	Send(*SubscribeCmdStatusResponse) error
	grpc.ServerStream
}

type droneServiceForFrontgateSubscribeCmdStatusServer struct {
	grpc.ServerStream
}

func (x *droneServiceForFrontgateSubscribeCmdStatusServer) Send(m *SubscribeCmdStatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DroneServiceForFrontgate_GetRegisterCmdStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegisterCmdStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceForFrontgateServer).GetRegisterCmdStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.drone.DroneServiceForFrontgate/GetRegisterCmdStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceForFrontgateServer).GetRegisterCmdStatus(ctx, req.(*GetRegisterCmdStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DroneServiceForFrontgate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.drone.DroneServiceForFrontgate",
	HandlerType: (*DroneServiceForFrontgateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRegisterCmdStatus",
			Handler:    _DroneServiceForFrontgate_GetRegisterCmdStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeCmdStatus",
			Handler:       _DroneServiceForFrontgate_SubscribeCmdStatus_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "drone.proto",
}

func init() { proto.RegisterFile("drone.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 913 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdb, 0x6e, 0xdb, 0x46,
	0x10, 0x05, 0x65, 0x45, 0x97, 0x91, 0x11, 0x38, 0x0b, 0xc3, 0x61, 0x98, 0xb8, 0x35, 0x98, 0xa2,
	0x70, 0x6f, 0xb4, 0x2b, 0x07, 0x4d, 0xe2, 0x36, 0x40, 0x61, 0x27, 0x56, 0xd5, 0x14, 0x28, 0x2a,
	0x19, 0x79, 0x15, 0x28, 0x6a, 0xc4, 0x12, 0xa6, 0xb8, 0x9b, 0xdd, 0xa5, 0x63, 0xbe, 0xf7, 0xb9,
	0xaf, 0xfd, 0x8a, 0x7e, 0x41, 0xbf, 0xa4, 0x1f, 0x53, 0xa0, 0xe0, 0xee, 0x4a, 0x61, 0x42, 0xd1,
	0x66, 0xd0, 0xbe, 0x08, 0x9c, 0xe5, 0x39, 0x67, 0x66, 0x87, 0x33, 0x47, 0xd0, 0x9b, 0x71, 0x9a,
	0xa0, 0xc7, 0x38, 0x95, 0x94, 0x6c, 0x51, 0x86, 0x09, 0x8b, 0x24, 0x8f, 0xae, 0x3c, 0x75, 0xee,
	0x7c, 0x1c, 0x52, 0x1a, 0xc6, 0x78, 0xa0, 0xde, 0x4f, 0xd3, 0xf9, 0x81, 0x8c, 0x16, 0x28, 0xa4,
	0xbf, 0x60, 0x9a, 0xe2, 0x7c, 0xf4, 0x3e, 0xe0, 0x0d, 0xf7, 0x19, 0x43, 0x2e, 0xf4, 0x7b, 0xf7,
	0x2f, 0x0b, 0x9a, 0xc3, 0x64, 0x4e, 0xc9, 0x3d, 0xe8, 0x28, 0xc9, 0x49, 0xc4, 0x6c, 0x6b, 0xcf,
	0xda, 0xef, 0x8e, 0xda, 0x2a, 0x1e, 0x32, 0xf2, 0x3d, 0x6c, 0x06, 0x34, 0x99, 0xcf, 0x26, 0xf9,
	0x6f, 0x14, 0xda, 0x8d, 0x3d, 0x6b, 0xbf, 0xd7, 0xdf, 0xf5, 0xde, 0xaf, 0xc6, 0x3b, 0xcd, 0x51,
	0xa7, 0x0a, 0x34, 0xea, 0x05, 0x6f, 0x03, 0xf2, 0x0a, 0xb6, 0xb5, 0xc2, 0xd4, 0x0f, 0x2e, 0x30,
	0x59, 0x29, 0x6d, 0x28, 0xa5, 0x4f, 0x2a, 0x94, 0x4e, 0x34, 0xd8, 0x08, 0x92, 0xa0, 0x74, 0xe6,
	0xfe, 0xdd, 0x80, 0x5e, 0x21, 0x29, 0x79, 0x0c, 0x9d, 0x1c, 0x35, 0x99, 0x45, 0x5c, 0x5d, 0xa2,
	0xd7, 0x7f, 0xe0, 0xe9, 0x06, 0x78, 0xcb, 0x06, 0x78, 0x63, 0xc9, 0xa3, 0x24, 0x7c, 0xe5, 0xc7,
	0x29, 0x8e, 0xda, 0x39, 0xfa, 0x79, 0xc4, 0x73, 0x62, 0x94, 0x48, 0xe4, 0x97, 0x7e, 0x6c, 0xae,
	0x77, 0xbf, 0x44, 0x1c, 0x26, 0xf2, 0xa8, 0xaf, 0x79, 0x2b, 0x30, 0xf1, 0xa0, 0x99, 0x50, 0xca,
	0xcc, 0x4d, 0x9c, 0x12, 0xe9, 0x84, 0xd2, 0x58, 0x73, 0x14, 0x8e, 0x3c, 0x82, 0x16, 0xe3, 0x38,
	0x8f, 0xae, 0xec, 0x66, 0x8d, 0xfa, 0x0c, 0x96, 0x3c, 0x86, 0xae, 0xc8, 0x92, 0x60, 0x42, 0x93,
	0x38, 0xb3, 0x6f, 0xdd, 0x98, 0xaa, 0x93, 0x83, 0x7f, 0x4e, 0xe2, 0x8c, 0x3c, 0x85, 0x6e, 0x4c,
	0xc3, 0x49, 0x8c, 0x97, 0x18, 0xdb, 0xad, 0x1a, 0x19, 0x3b, 0x31, 0x0d, 0x7f, 0xca, 0xd1, 0xee,
	0xef, 0x1b, 0x40, 0xca, 0x9f, 0x81, 0x1c, 0x42, 0x53, 0x66, 0x0c, 0x6b, 0xb5, 0x57, 0x21, 0x09,
	0x81, 0xe6, 0xaf, 0x54, 0x48, 0xbb, 0xb1, 0xb7, 0xb1, 0xdf, 0x1d, 0xa9, 0x67, 0xf2, 0x04, 0x3a,
	0xa9, 0x40, 0x9e, 0xf8, 0x0b, 0x34, 0xad, 0xbb, 0xa1, 0xac, 0x25, 0x3a, 0x67, 0x32, 0x5f, 0x88,
	0x37, 0x94, 0xcf, 0x6a, 0xb5, 0x70, 0x85, 0x26, 0x27, 0x70, 0x3b, 0x88, 0x23, 0x4c, 0xe4, 0x24,
	0xf0, 0x27, 0x17, 0x98, 0x09, 0xd3, 0xc9, 0xeb, 0xf9, 0x9b, 0x9a, 0x73, 0xea, 0xbf, 0xc4, 0x4c,
	0x90, 0x67, 0xd0, 0x5b, 0x6a, 0x20, 0x97, 0xb5, 0x3a, 0x0a, 0x46, 0x00, 0xb9, 0x24, 0xdf, 0x82,
	0x89, 0xf2, 0xfc, 0x76, 0xbb, 0x06, 0xbb, 0xab, 0xf1, 0x2f, 0x31, 0x73, 0xff, 0xb4, 0xe0, 0xce,
	0x58, 0xfa, 0x5c, 0xaa, 0xaf, 0x32, 0xc2, 0xd7, 0x29, 0x0a, 0x59, 0x5a, 0x4e, 0xeb, 0x7f, 0x5b,
	0xce, 0xc6, 0x7f, 0x5c, 0xce, 0xcc, 0xec, 0xe6, 0x58, 0xfa, 0x32, 0x15, 0x64, 0x17, 0x80, 0x71,
	0x1a, 0xa0, 0x10, 0x93, 0x68, 0x66, 0x2c, 0xa6, 0x6b, 0x4e, 0x86, 0x33, 0xb2, 0x03, 0x2d, 0xa1,
	0x80, 0x2a, 0x6f, 0x77, 0x64, 0x22, 0x72, 0x04, 0xed, 0x94, 0x4d, 0x72, 0x5b, 0xab, 0xdc, 0xb1,
	0xf3, 0xa5, 0xe7, 0x8d, 0x5a, 0x29, 0xcb, 0x03, 0xf7, 0x18, 0xee, 0x8d, 0xd3, 0xa9, 0x08, 0x78,
	0x34, 0xc5, 0xd3, 0x85, 0xa9, 0x60, 0xd9, 0xb1, 0x5d, 0x00, 0x91, 0x4e, 0xa5, 0x2f, 0x2e, 0x0a,
	0x85, 0x98, 0x93, 0xe1, 0xcc, 0x7d, 0x04, 0xce, 0x3a, 0xae, 0x60, 0x34, 0x11, 0x58, 0x28, 0xd3,
	0x2a, 0x96, 0xe9, 0x7e, 0x07, 0xf7, 0x07, 0x28, 0x47, 0x18, 0x46, 0x42, 0x22, 0xff, 0xd0, 0x9c,
	0xdf, 0xc0, 0x83, 0xf5, 0xec, 0x1b, 0xb2, 0x7e, 0x0d, 0x77, 0x07, 0x28, 0xcf, 0x71, 0xc1, 0x62,
	0x5f, 0xe2, 0x59, 0x14, 0xe3, 0x2a, 0xe3, 0x0e, 0xb4, 0x38, 0x86, 0x78, 0xb5, 0x74, 0x73, 0x13,
	0xb9, 0x87, 0x60, 0x97, 0x29, 0x26, 0xcd, 0x36, 0xdc, 0x9a, 0xe7, 0x07, 0xb6, 0xa5, 0x56, 0x55,
	0x07, 0xee, 0xa7, 0xb0, 0x35, 0x40, 0xa9, 0xc6, 0x71, 0xa5, 0x4e, 0xa0, 0xa9, 0x36, 0x48, 0x03,
	0xd5, 0xb3, 0xfb, 0x87, 0x05, 0x77, 0x0a, 0x40, 0xa3, 0x39, 0x80, 0xd6, 0xa5, 0x3a, 0x51, 0xd8,
	0x5e, 0xff, 0xa0, 0x3c, 0x4f, 0x25, 0x92, 0xa7, 0xc3, 0x17, 0x89, 0xe4, 0xd9, 0xc8, 0xd0, 0x9d,
	0xa7, 0xd0, 0x2b, 0x1c, 0x93, 0x2d, 0xd8, 0xc8, 0x77, 0x48, 0x5f, 0x2e, 0x7f, 0xcc, 0xab, 0x57,
	0x50, 0x33, 0x40, 0x3a, 0x38, 0x6e, 0x3c, 0xb1, 0xdc, 0x36, 0xdc, 0x7a, 0xb1, 0x60, 0x32, 0xeb,
	0xff, 0xd3, 0x84, 0xcd, 0xe7, 0x79, 0xce, 0x31, 0xf2, 0xcb, 0x28, 0x40, 0x72, 0x0c, 0xed, 0x01,
	0x4a, 0xf5, 0x07, 0x78, 0xb7, 0x5c, 0x98, 0x22, 0x39, 0x3b, 0xe5, 0x17, 0x8a, 0xf0, 0x03, 0xdc,
	0x1e, 0xa0, 0x2c, 0xfe, 0xfd, 0x54, 0x4a, 0x5c, 0xbf, 0x8e, 0xe4, 0x17, 0xd5, 0xe1, 0x77, 0x7d,
	0xb6, 0x52, 0xab, 0xd6, 0x42, 0x92, 0x1f, 0x01, 0xde, 0x7a, 0x05, 0x79, 0x58, 0xe6, 0x94, 0x9c,
	0xc4, 0xa9, 0xca, 0x48, 0x9e, 0x41, 0x77, 0x2c, 0x29, 0xd3, 0x52, 0x95, 0x75, 0x55, 0xd2, 0x0b,
	0x7d, 0x32, 0x56, 0xf0, 0xc1, 0x7d, 0x32, 0xbc, 0x48, 0xf5, 0xe9, 0x9d, 0xd9, 0x25, 0x9f, 0xad,
	0x9d, 0xa7, 0x75, 0x2b, 0xe1, 0x7c, 0x5e, 0x07, 0x6a, 0xc6, 0xf6, 0x1c, 0xba, 0xab, 0xb1, 0x24,
	0xee, 0xb5, 0x33, 0xab, 0xc5, 0x1f, 0xd6, 0x98, 0xeb, 0xfe, 0x6f, 0x0d, 0xb0, 0x8b, 0xf3, 0x77,
	0x46, 0xf9, 0x19, 0xa7, 0x89, 0x0c, 0x7d, 0x89, 0xe4, 0x35, 0x90, 0xb2, 0xf1, 0x90, 0x2f, 0xd6,
	0x7c, 0xba, 0x2a, 0x6b, 0x73, 0xbe, 0xac, 0x07, 0xd6, 0xd5, 0x1c, 0x5a, 0x24, 0x85, 0xed, 0x75,
	0xbe, 0x43, 0xbe, 0x5a, 0x7b, 0x99, 0x2a, 0x77, 0x73, 0xbc, 0xba, 0x70, 0x9d, 0x78, 0xda, 0x52,
	0xd6, 0x7d, 0xf4, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x13, 0x2e, 0xc3, 0x03, 0xdd, 0x0a, 0x00,
	0x00,
}
