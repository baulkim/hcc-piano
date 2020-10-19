// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rpcharp

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	rpcmsgType "hcc/piano/action/grpc/pb/rpcmsgType"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HarpClient is the client API for Harp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HarpClient interface {
	// Subnet
	CreateSubnet(ctx context.Context, in *ReqCreateSubnet, opts ...grpc.CallOption) (*ResCreateSubnet, error)
	GetSubnet(ctx context.Context, in *ReqGetSubnet, opts ...grpc.CallOption) (*ResGetSubnet, error)
	GetSubnetByServer(ctx context.Context, in *ReqGetSubnetByServer, opts ...grpc.CallOption) (*ResGetSubnetByServer, error)
	GetSubnetList(ctx context.Context, in *ReqGetSubnetList, opts ...grpc.CallOption) (*ResGetSubnetList, error)
	GetSubnetNum(ctx context.Context, in *rpcmsgType.Empty, opts ...grpc.CallOption) (*ResGetSubnetNum, error)
	UpdateSubnet(ctx context.Context, in *ReqUpdateSubnet, opts ...grpc.CallOption) (*ResUpdateSubnet, error)
	DeleteSubnet(ctx context.Context, in *ReqDeleteSubnet, opts ...grpc.CallOption) (*ResDeleteSubnet, error)
	// AdaptiveIP
	CreateAdaptiveIPSetting(ctx context.Context, in *ReqCreateAdaptiveIPSetting, opts ...grpc.CallOption) (*ResCreateAdaptiveIPSetting, error)
	GetAdaptiveIPSetting(ctx context.Context, in *rpcmsgType.Empty, opts ...grpc.CallOption) (*ResGetAdaptiveIPSetting, error)
	GetAdaptiveIPAvailableIPList(ctx context.Context, in *rpcmsgType.Empty, opts ...grpc.CallOption) (*ResGetAdaptiveIPAvailableIPList, error)
	CreateAdaptiveIPServer(ctx context.Context, in *ReqCreateAdaptiveIPServer, opts ...grpc.CallOption) (*ResCreateAdaptiveIPServer, error)
	GetAdaptiveIPServer(ctx context.Context, in *ReqGetAdaptiveIPServer, opts ...grpc.CallOption) (*ResGetAdaptiveIPServer, error)
	GetAdaptiveIPServerList(ctx context.Context, in *ReqGetAdaptiveIPServerList, opts ...grpc.CallOption) (*ResGetAdaptiveIPServerList, error)
	GetAdaptiveIPServerNum(ctx context.Context, in *rpcmsgType.Empty, opts ...grpc.CallOption) (*ResGetAdaptiveIPServerNum, error)
	DeleteAdaptiveIPServer(ctx context.Context, in *ReqDeleteAdaptiveIPServer, opts ...grpc.CallOption) (*ResDeleteAdaptiveIPServer, error)
	// DHCPD
	CreateDHCPDConf(ctx context.Context, in *ReqCreateDHCPDConf, opts ...grpc.CallOption) (*ResCreateDHCPDConf, error)
	DeleteDHCPDConf(ctx context.Context, in *ReqDeleteDHCPDConf, opts ...grpc.CallOption) (*ResDeleteDHCPDConf, error)
}

type harpClient struct {
	cc grpc.ClientConnInterface
}

func NewHarpClient(cc grpc.ClientConnInterface) HarpClient {
	return &harpClient{cc}
}

func (c *harpClient) CreateSubnet(ctx context.Context, in *ReqCreateSubnet, opts ...grpc.CallOption) (*ResCreateSubnet, error) {
	out := new(ResCreateSubnet)
	err := c.cc.Invoke(ctx, "/RpcHarp.Harp/CreateSubnet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harpClient) GetSubnet(ctx context.Context, in *ReqGetSubnet, opts ...grpc.CallOption) (*ResGetSubnet, error) {
	out := new(ResGetSubnet)
	err := c.cc.Invoke(ctx, "/RpcHarp.Harp/GetSubnet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harpClient) GetSubnetByServer(ctx context.Context, in *ReqGetSubnetByServer, opts ...grpc.CallOption) (*ResGetSubnetByServer, error) {
	out := new(ResGetSubnetByServer)
	err := c.cc.Invoke(ctx, "/RpcHarp.Harp/GetSubnetByServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harpClient) GetSubnetList(ctx context.Context, in *ReqGetSubnetList, opts ...grpc.CallOption) (*ResGetSubnetList, error) {
	out := new(ResGetSubnetList)
	err := c.cc.Invoke(ctx, "/RpcHarp.Harp/GetSubnetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harpClient) GetSubnetNum(ctx context.Context, in *rpcmsgType.Empty, opts ...grpc.CallOption) (*ResGetSubnetNum, error) {
	out := new(ResGetSubnetNum)
	err := c.cc.Invoke(ctx, "/RpcHarp.Harp/GetSubnetNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harpClient) UpdateSubnet(ctx context.Context, in *ReqUpdateSubnet, opts ...grpc.CallOption) (*ResUpdateSubnet, error) {
	out := new(ResUpdateSubnet)
	err := c.cc.Invoke(ctx, "/RpcHarp.Harp/UpdateSubnet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harpClient) DeleteSubnet(ctx context.Context, in *ReqDeleteSubnet, opts ...grpc.CallOption) (*ResDeleteSubnet, error) {
	out := new(ResDeleteSubnet)
	err := c.cc.Invoke(ctx, "/RpcHarp.Harp/DeleteSubnet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harpClient) CreateAdaptiveIPSetting(ctx context.Context, in *ReqCreateAdaptiveIPSetting, opts ...grpc.CallOption) (*ResCreateAdaptiveIPSetting, error) {
	out := new(ResCreateAdaptiveIPSetting)
	err := c.cc.Invoke(ctx, "/RpcHarp.Harp/CreateAdaptiveIPSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harpClient) GetAdaptiveIPSetting(ctx context.Context, in *rpcmsgType.Empty, opts ...grpc.CallOption) (*ResGetAdaptiveIPSetting, error) {
	out := new(ResGetAdaptiveIPSetting)
	err := c.cc.Invoke(ctx, "/RpcHarp.Harp/GetAdaptiveIPSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harpClient) GetAdaptiveIPAvailableIPList(ctx context.Context, in *rpcmsgType.Empty, opts ...grpc.CallOption) (*ResGetAdaptiveIPAvailableIPList, error) {
	out := new(ResGetAdaptiveIPAvailableIPList)
	err := c.cc.Invoke(ctx, "/RpcHarp.Harp/GetAdaptiveIPAvailableIPList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harpClient) CreateAdaptiveIPServer(ctx context.Context, in *ReqCreateAdaptiveIPServer, opts ...grpc.CallOption) (*ResCreateAdaptiveIPServer, error) {
	out := new(ResCreateAdaptiveIPServer)
	err := c.cc.Invoke(ctx, "/RpcHarp.Harp/CreateAdaptiveIPServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harpClient) GetAdaptiveIPServer(ctx context.Context, in *ReqGetAdaptiveIPServer, opts ...grpc.CallOption) (*ResGetAdaptiveIPServer, error) {
	out := new(ResGetAdaptiveIPServer)
	err := c.cc.Invoke(ctx, "/RpcHarp.Harp/GetAdaptiveIPServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harpClient) GetAdaptiveIPServerList(ctx context.Context, in *ReqGetAdaptiveIPServerList, opts ...grpc.CallOption) (*ResGetAdaptiveIPServerList, error) {
	out := new(ResGetAdaptiveIPServerList)
	err := c.cc.Invoke(ctx, "/RpcHarp.Harp/GetAdaptiveIPServerList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harpClient) GetAdaptiveIPServerNum(ctx context.Context, in *rpcmsgType.Empty, opts ...grpc.CallOption) (*ResGetAdaptiveIPServerNum, error) {
	out := new(ResGetAdaptiveIPServerNum)
	err := c.cc.Invoke(ctx, "/RpcHarp.Harp/GetAdaptiveIPServerNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harpClient) DeleteAdaptiveIPServer(ctx context.Context, in *ReqDeleteAdaptiveIPServer, opts ...grpc.CallOption) (*ResDeleteAdaptiveIPServer, error) {
	out := new(ResDeleteAdaptiveIPServer)
	err := c.cc.Invoke(ctx, "/RpcHarp.Harp/DeleteAdaptiveIPServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harpClient) CreateDHCPDConf(ctx context.Context, in *ReqCreateDHCPDConf, opts ...grpc.CallOption) (*ResCreateDHCPDConf, error) {
	out := new(ResCreateDHCPDConf)
	err := c.cc.Invoke(ctx, "/RpcHarp.Harp/CreateDHCPDConf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *harpClient) DeleteDHCPDConf(ctx context.Context, in *ReqDeleteDHCPDConf, opts ...grpc.CallOption) (*ResDeleteDHCPDConf, error) {
	out := new(ResDeleteDHCPDConf)
	err := c.cc.Invoke(ctx, "/RpcHarp.Harp/DeleteDHCPDConf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HarpServer is the server API for Harp service.
// All implementations must embed UnimplementedHarpServer
// for forward compatibility
type HarpServer interface {
	// Subnet
	CreateSubnet(context.Context, *ReqCreateSubnet) (*ResCreateSubnet, error)
	GetSubnet(context.Context, *ReqGetSubnet) (*ResGetSubnet, error)
	GetSubnetByServer(context.Context, *ReqGetSubnetByServer) (*ResGetSubnetByServer, error)
	GetSubnetList(context.Context, *ReqGetSubnetList) (*ResGetSubnetList, error)
	GetSubnetNum(context.Context, *rpcmsgType.Empty) (*ResGetSubnetNum, error)
	UpdateSubnet(context.Context, *ReqUpdateSubnet) (*ResUpdateSubnet, error)
	DeleteSubnet(context.Context, *ReqDeleteSubnet) (*ResDeleteSubnet, error)
	// AdaptiveIP
	CreateAdaptiveIPSetting(context.Context, *ReqCreateAdaptiveIPSetting) (*ResCreateAdaptiveIPSetting, error)
	GetAdaptiveIPSetting(context.Context, *rpcmsgType.Empty) (*ResGetAdaptiveIPSetting, error)
	GetAdaptiveIPAvailableIPList(context.Context, *rpcmsgType.Empty) (*ResGetAdaptiveIPAvailableIPList, error)
	CreateAdaptiveIPServer(context.Context, *ReqCreateAdaptiveIPServer) (*ResCreateAdaptiveIPServer, error)
	GetAdaptiveIPServer(context.Context, *ReqGetAdaptiveIPServer) (*ResGetAdaptiveIPServer, error)
	GetAdaptiveIPServerList(context.Context, *ReqGetAdaptiveIPServerList) (*ResGetAdaptiveIPServerList, error)
	GetAdaptiveIPServerNum(context.Context, *rpcmsgType.Empty) (*ResGetAdaptiveIPServerNum, error)
	DeleteAdaptiveIPServer(context.Context, *ReqDeleteAdaptiveIPServer) (*ResDeleteAdaptiveIPServer, error)
	// DHCPD
	CreateDHCPDConf(context.Context, *ReqCreateDHCPDConf) (*ResCreateDHCPDConf, error)
	DeleteDHCPDConf(context.Context, *ReqDeleteDHCPDConf) (*ResDeleteDHCPDConf, error)
	mustEmbedUnimplementedHarpServer()
}

// UnimplementedHarpServer must be embedded to have forward compatible implementations.
type UnimplementedHarpServer struct {
}

func (*UnimplementedHarpServer) CreateSubnet(context.Context, *ReqCreateSubnet) (*ResCreateSubnet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubnet not implemented")
}
func (*UnimplementedHarpServer) GetSubnet(context.Context, *ReqGetSubnet) (*ResGetSubnet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubnet not implemented")
}
func (*UnimplementedHarpServer) GetSubnetByServer(context.Context, *ReqGetSubnetByServer) (*ResGetSubnetByServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubnetByServer not implemented")
}
func (*UnimplementedHarpServer) GetSubnetList(context.Context, *ReqGetSubnetList) (*ResGetSubnetList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubnetList not implemented")
}
func (*UnimplementedHarpServer) GetSubnetNum(context.Context, *rpcmsgType.Empty) (*ResGetSubnetNum, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubnetNum not implemented")
}
func (*UnimplementedHarpServer) UpdateSubnet(context.Context, *ReqUpdateSubnet) (*ResUpdateSubnet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSubnet not implemented")
}
func (*UnimplementedHarpServer) DeleteSubnet(context.Context, *ReqDeleteSubnet) (*ResDeleteSubnet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubnet not implemented")
}
func (*UnimplementedHarpServer) CreateAdaptiveIPSetting(context.Context, *ReqCreateAdaptiveIPSetting) (*ResCreateAdaptiveIPSetting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAdaptiveIPSetting not implemented")
}
func (*UnimplementedHarpServer) GetAdaptiveIPSetting(context.Context, *rpcmsgType.Empty) (*ResGetAdaptiveIPSetting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdaptiveIPSetting not implemented")
}
func (*UnimplementedHarpServer) GetAdaptiveIPAvailableIPList(context.Context, *rpcmsgType.Empty) (*ResGetAdaptiveIPAvailableIPList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdaptiveIPAvailableIPList not implemented")
}
func (*UnimplementedHarpServer) CreateAdaptiveIPServer(context.Context, *ReqCreateAdaptiveIPServer) (*ResCreateAdaptiveIPServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAdaptiveIPServer not implemented")
}
func (*UnimplementedHarpServer) GetAdaptiveIPServer(context.Context, *ReqGetAdaptiveIPServer) (*ResGetAdaptiveIPServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdaptiveIPServer not implemented")
}
func (*UnimplementedHarpServer) GetAdaptiveIPServerList(context.Context, *ReqGetAdaptiveIPServerList) (*ResGetAdaptiveIPServerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdaptiveIPServerList not implemented")
}
func (*UnimplementedHarpServer) GetAdaptiveIPServerNum(context.Context, *rpcmsgType.Empty) (*ResGetAdaptiveIPServerNum, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdaptiveIPServerNum not implemented")
}
func (*UnimplementedHarpServer) DeleteAdaptiveIPServer(context.Context, *ReqDeleteAdaptiveIPServer) (*ResDeleteAdaptiveIPServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAdaptiveIPServer not implemented")
}
func (*UnimplementedHarpServer) CreateDHCPDConf(context.Context, *ReqCreateDHCPDConf) (*ResCreateDHCPDConf, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDHCPDConf not implemented")
}
func (*UnimplementedHarpServer) DeleteDHCPDConf(context.Context, *ReqDeleteDHCPDConf) (*ResDeleteDHCPDConf, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDHCPDConf not implemented")
}
func (*UnimplementedHarpServer) mustEmbedUnimplementedHarpServer() {}

func RegisterHarpServer(s *grpc.Server, srv HarpServer) {
	s.RegisterService(&_Harp_serviceDesc, srv)
}

func _Harp_CreateSubnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqCreateSubnet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarpServer).CreateSubnet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcHarp.Harp/CreateSubnet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarpServer).CreateSubnet(ctx, req.(*ReqCreateSubnet))
	}
	return interceptor(ctx, in, info, handler)
}

func _Harp_GetSubnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetSubnet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarpServer).GetSubnet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcHarp.Harp/GetSubnet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarpServer).GetSubnet(ctx, req.(*ReqGetSubnet))
	}
	return interceptor(ctx, in, info, handler)
}

func _Harp_GetSubnetByServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetSubnetByServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarpServer).GetSubnetByServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcHarp.Harp/GetSubnetByServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarpServer).GetSubnetByServer(ctx, req.(*ReqGetSubnetByServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Harp_GetSubnetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetSubnetList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarpServer).GetSubnetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcHarp.Harp/GetSubnetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarpServer).GetSubnetList(ctx, req.(*ReqGetSubnetList))
	}
	return interceptor(ctx, in, info, handler)
}

func _Harp_GetSubnetNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcmsgType.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarpServer).GetSubnetNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcHarp.Harp/GetSubnetNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarpServer).GetSubnetNum(ctx, req.(*rpcmsgType.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Harp_UpdateSubnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqUpdateSubnet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarpServer).UpdateSubnet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcHarp.Harp/UpdateSubnet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarpServer).UpdateSubnet(ctx, req.(*ReqUpdateSubnet))
	}
	return interceptor(ctx, in, info, handler)
}

func _Harp_DeleteSubnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqDeleteSubnet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarpServer).DeleteSubnet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcHarp.Harp/DeleteSubnet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarpServer).DeleteSubnet(ctx, req.(*ReqDeleteSubnet))
	}
	return interceptor(ctx, in, info, handler)
}

func _Harp_CreateAdaptiveIPSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqCreateAdaptiveIPSetting)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarpServer).CreateAdaptiveIPSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcHarp.Harp/CreateAdaptiveIPSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarpServer).CreateAdaptiveIPSetting(ctx, req.(*ReqCreateAdaptiveIPSetting))
	}
	return interceptor(ctx, in, info, handler)
}

func _Harp_GetAdaptiveIPSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcmsgType.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarpServer).GetAdaptiveIPSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcHarp.Harp/GetAdaptiveIPSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarpServer).GetAdaptiveIPSetting(ctx, req.(*rpcmsgType.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Harp_GetAdaptiveIPAvailableIPList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcmsgType.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarpServer).GetAdaptiveIPAvailableIPList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcHarp.Harp/GetAdaptiveIPAvailableIPList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarpServer).GetAdaptiveIPAvailableIPList(ctx, req.(*rpcmsgType.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Harp_CreateAdaptiveIPServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqCreateAdaptiveIPServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarpServer).CreateAdaptiveIPServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcHarp.Harp/CreateAdaptiveIPServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarpServer).CreateAdaptiveIPServer(ctx, req.(*ReqCreateAdaptiveIPServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Harp_GetAdaptiveIPServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetAdaptiveIPServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarpServer).GetAdaptiveIPServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcHarp.Harp/GetAdaptiveIPServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarpServer).GetAdaptiveIPServer(ctx, req.(*ReqGetAdaptiveIPServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Harp_GetAdaptiveIPServerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetAdaptiveIPServerList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarpServer).GetAdaptiveIPServerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcHarp.Harp/GetAdaptiveIPServerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarpServer).GetAdaptiveIPServerList(ctx, req.(*ReqGetAdaptiveIPServerList))
	}
	return interceptor(ctx, in, info, handler)
}

func _Harp_GetAdaptiveIPServerNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcmsgType.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarpServer).GetAdaptiveIPServerNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcHarp.Harp/GetAdaptiveIPServerNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarpServer).GetAdaptiveIPServerNum(ctx, req.(*rpcmsgType.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Harp_DeleteAdaptiveIPServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqDeleteAdaptiveIPServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarpServer).DeleteAdaptiveIPServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcHarp.Harp/DeleteAdaptiveIPServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarpServer).DeleteAdaptiveIPServer(ctx, req.(*ReqDeleteAdaptiveIPServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Harp_CreateDHCPDConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqCreateDHCPDConf)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarpServer).CreateDHCPDConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcHarp.Harp/CreateDHCPDConf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarpServer).CreateDHCPDConf(ctx, req.(*ReqCreateDHCPDConf))
	}
	return interceptor(ctx, in, info, handler)
}

func _Harp_DeleteDHCPDConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqDeleteDHCPDConf)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HarpServer).DeleteDHCPDConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcHarp.Harp/DeleteDHCPDConf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HarpServer).DeleteDHCPDConf(ctx, req.(*ReqDeleteDHCPDConf))
	}
	return interceptor(ctx, in, info, handler)
}

var _Harp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RpcHarp.Harp",
	HandlerType: (*HarpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSubnet",
			Handler:    _Harp_CreateSubnet_Handler,
		},
		{
			MethodName: "GetSubnet",
			Handler:    _Harp_GetSubnet_Handler,
		},
		{
			MethodName: "GetSubnetByServer",
			Handler:    _Harp_GetSubnetByServer_Handler,
		},
		{
			MethodName: "GetSubnetList",
			Handler:    _Harp_GetSubnetList_Handler,
		},
		{
			MethodName: "GetSubnetNum",
			Handler:    _Harp_GetSubnetNum_Handler,
		},
		{
			MethodName: "UpdateSubnet",
			Handler:    _Harp_UpdateSubnet_Handler,
		},
		{
			MethodName: "DeleteSubnet",
			Handler:    _Harp_DeleteSubnet_Handler,
		},
		{
			MethodName: "CreateAdaptiveIPSetting",
			Handler:    _Harp_CreateAdaptiveIPSetting_Handler,
		},
		{
			MethodName: "GetAdaptiveIPSetting",
			Handler:    _Harp_GetAdaptiveIPSetting_Handler,
		},
		{
			MethodName: "GetAdaptiveIPAvailableIPList",
			Handler:    _Harp_GetAdaptiveIPAvailableIPList_Handler,
		},
		{
			MethodName: "CreateAdaptiveIPServer",
			Handler:    _Harp_CreateAdaptiveIPServer_Handler,
		},
		{
			MethodName: "GetAdaptiveIPServer",
			Handler:    _Harp_GetAdaptiveIPServer_Handler,
		},
		{
			MethodName: "GetAdaptiveIPServerList",
			Handler:    _Harp_GetAdaptiveIPServerList_Handler,
		},
		{
			MethodName: "GetAdaptiveIPServerNum",
			Handler:    _Harp_GetAdaptiveIPServerNum_Handler,
		},
		{
			MethodName: "DeleteAdaptiveIPServer",
			Handler:    _Harp_DeleteAdaptiveIPServer_Handler,
		},
		{
			MethodName: "CreateDHCPDConf",
			Handler:    _Harp_CreateDHCPDConf_Handler,
		},
		{
			MethodName: "DeleteDHCPDConf",
			Handler:    _Harp_DeleteDHCPDConf_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "harp.proto",
}
