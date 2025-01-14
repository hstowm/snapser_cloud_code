// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: storage_service.proto

package proto

import (
	context "context"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StorageServiceClient is the client API for StorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageServiceClient interface {
	// InsertBlob
	//
	// Inserts a blob. Will fail if the key already exists
	InsertBlob(ctx context.Context, in *InsertBlobRequest, opts ...grpc.CallOption) (*InsertBlobResponse, error)
	// ReplaceBlob
	//
	// ReplaceBlob inserts or updates a blob
	ReplaceBlob(ctx context.Context, in *ReplaceBlobRequest, opts ...grpc.CallOption) (*ReplaceBlobResponse, error)
	// GetBlob
	//
	// GetBlob returns the blob value for a key
	GetBlob(ctx context.Context, in *GetBlobRequest, opts ...grpc.CallOption) (*GetBlobResponse, error)
	// DeleteBlob
	//
	// DeleteBlob
	DeleteBlob(ctx context.Context, in *DeleteBlobRequest, opts ...grpc.CallOption) (*DeleteBlobResponse, error)
	// UpdateAppendBlob
	//
	// UpdateAppendBlob will append to the blob
	UpdateAppendBlob(ctx context.Context, in *UpdateAppendBlobRequest, opts ...grpc.CallOption) (*UpdateAppendBlobResponse, error)
	// GetAppendBlob
	//
	// GetAppendBlob returns the value of the append blob
	GetAppendBlob(ctx context.Context, in *GetAppendBlobRequest, opts ...grpc.CallOption) (*GetAppendBlobResponse, error)
	// DeleteAppendBlob
	//
	// DeleteAppendBlob deletes the append blob
	DeleteAppendBlob(ctx context.Context, in *DeleteAppendBlobRequest, opts ...grpc.CallOption) (*DeleteAppendBlobResponse, error)
	// DeleteUserData
	//
	// DeleteUserData deletes the blob owned by the user
	DeleteUserData(ctx context.Context, in *DeleteUserDataRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GetUserData
	//
	// GetUserData gets the blob owned by the user
	GetUserData(ctx context.Context, in *GetUserDataRequest, opts ...grpc.CallOption) (*GetUserDataResponse, error)
	// IncrementCounter
	//
	// IncrementCounter increments a counter value
	IncrementCounter(ctx context.Context, in *IncrementCounterRequest, opts ...grpc.CallOption) (*IncrementCounterResponse, error)
	// GetCounter
	//
	// GetCounter gets a counter value
	GetCounter(ctx context.Context, in *GetCounterRequest, opts ...grpc.CallOption) (*GetCounterResponse, error)
	// ResetCounter
	//
	// ResetCounter resets a counter
	ResetCounter(ctx context.Context, in *ResetCounterRequest, opts ...grpc.CallOption) (*ResetCounterResponse, error)
	AddKeys(ctx context.Context, in *AddKeysRequest, opts ...grpc.CallOption) (*AddKeysResponse, error)
	GetKeys(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetKeysResponse, error)
	DeleteKey(ctx context.Context, in *DeleteKeyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GetOpenApiSpec
	//
	// Returns the generated openapi swagger document
	GetOpenApiSpec(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	ExportSettings(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SettingsImportExportData, error)
	ImportSettings(ctx context.Context, in *SettingsImportExportData, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type storageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageServiceClient(cc grpc.ClientConnInterface) StorageServiceClient {
	return &storageServiceClient{cc}
}

func (c *storageServiceClient) InsertBlob(ctx context.Context, in *InsertBlobRequest, opts ...grpc.CallOption) (*InsertBlobResponse, error) {
	out := new(InsertBlobResponse)
	err := c.cc.Invoke(ctx, "/storage.StorageService/InsertBlob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) ReplaceBlob(ctx context.Context, in *ReplaceBlobRequest, opts ...grpc.CallOption) (*ReplaceBlobResponse, error) {
	out := new(ReplaceBlobResponse)
	err := c.cc.Invoke(ctx, "/storage.StorageService/ReplaceBlob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) GetBlob(ctx context.Context, in *GetBlobRequest, opts ...grpc.CallOption) (*GetBlobResponse, error) {
	out := new(GetBlobResponse)
	err := c.cc.Invoke(ctx, "/storage.StorageService/GetBlob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) DeleteBlob(ctx context.Context, in *DeleteBlobRequest, opts ...grpc.CallOption) (*DeleteBlobResponse, error) {
	out := new(DeleteBlobResponse)
	err := c.cc.Invoke(ctx, "/storage.StorageService/DeleteBlob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) UpdateAppendBlob(ctx context.Context, in *UpdateAppendBlobRequest, opts ...grpc.CallOption) (*UpdateAppendBlobResponse, error) {
	out := new(UpdateAppendBlobResponse)
	err := c.cc.Invoke(ctx, "/storage.StorageService/UpdateAppendBlob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) GetAppendBlob(ctx context.Context, in *GetAppendBlobRequest, opts ...grpc.CallOption) (*GetAppendBlobResponse, error) {
	out := new(GetAppendBlobResponse)
	err := c.cc.Invoke(ctx, "/storage.StorageService/GetAppendBlob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) DeleteAppendBlob(ctx context.Context, in *DeleteAppendBlobRequest, opts ...grpc.CallOption) (*DeleteAppendBlobResponse, error) {
	out := new(DeleteAppendBlobResponse)
	err := c.cc.Invoke(ctx, "/storage.StorageService/DeleteAppendBlob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) DeleteUserData(ctx context.Context, in *DeleteUserDataRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/storage.StorageService/DeleteUserData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) GetUserData(ctx context.Context, in *GetUserDataRequest, opts ...grpc.CallOption) (*GetUserDataResponse, error) {
	out := new(GetUserDataResponse)
	err := c.cc.Invoke(ctx, "/storage.StorageService/GetUserData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) IncrementCounter(ctx context.Context, in *IncrementCounterRequest, opts ...grpc.CallOption) (*IncrementCounterResponse, error) {
	out := new(IncrementCounterResponse)
	err := c.cc.Invoke(ctx, "/storage.StorageService/IncrementCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) GetCounter(ctx context.Context, in *GetCounterRequest, opts ...grpc.CallOption) (*GetCounterResponse, error) {
	out := new(GetCounterResponse)
	err := c.cc.Invoke(ctx, "/storage.StorageService/GetCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) ResetCounter(ctx context.Context, in *ResetCounterRequest, opts ...grpc.CallOption) (*ResetCounterResponse, error) {
	out := new(ResetCounterResponse)
	err := c.cc.Invoke(ctx, "/storage.StorageService/ResetCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) AddKeys(ctx context.Context, in *AddKeysRequest, opts ...grpc.CallOption) (*AddKeysResponse, error) {
	out := new(AddKeysResponse)
	err := c.cc.Invoke(ctx, "/storage.StorageService/AddKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) GetKeys(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetKeysResponse, error) {
	out := new(GetKeysResponse)
	err := c.cc.Invoke(ctx, "/storage.StorageService/GetKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) DeleteKey(ctx context.Context, in *DeleteKeyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/storage.StorageService/DeleteKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) GetOpenApiSpec(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/storage.StorageService/GetOpenApiSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) ExportSettings(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SettingsImportExportData, error) {
	out := new(SettingsImportExportData)
	err := c.cc.Invoke(ctx, "/storage.StorageService/ExportSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) ImportSettings(ctx context.Context, in *SettingsImportExportData, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/storage.StorageService/ImportSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServiceServer is the server API for StorageService service.
// All implementations must embed UnimplementedStorageServiceServer
// for forward compatibility
type StorageServiceServer interface {
	// InsertBlob
	//
	// Inserts a blob. Will fail if the key already exists
	InsertBlob(context.Context, *InsertBlobRequest) (*InsertBlobResponse, error)
	// ReplaceBlob
	//
	// ReplaceBlob inserts or updates a blob
	ReplaceBlob(context.Context, *ReplaceBlobRequest) (*ReplaceBlobResponse, error)
	// GetBlob
	//
	// GetBlob returns the blob value for a key
	GetBlob(context.Context, *GetBlobRequest) (*GetBlobResponse, error)
	// DeleteBlob
	//
	// DeleteBlob
	DeleteBlob(context.Context, *DeleteBlobRequest) (*DeleteBlobResponse, error)
	// UpdateAppendBlob
	//
	// UpdateAppendBlob will append to the blob
	UpdateAppendBlob(context.Context, *UpdateAppendBlobRequest) (*UpdateAppendBlobResponse, error)
	// GetAppendBlob
	//
	// GetAppendBlob returns the value of the append blob
	GetAppendBlob(context.Context, *GetAppendBlobRequest) (*GetAppendBlobResponse, error)
	// DeleteAppendBlob
	//
	// DeleteAppendBlob deletes the append blob
	DeleteAppendBlob(context.Context, *DeleteAppendBlobRequest) (*DeleteAppendBlobResponse, error)
	// DeleteUserData
	//
	// DeleteUserData deletes the blob owned by the user
	DeleteUserData(context.Context, *DeleteUserDataRequest) (*emptypb.Empty, error)
	// GetUserData
	//
	// GetUserData gets the blob owned by the user
	GetUserData(context.Context, *GetUserDataRequest) (*GetUserDataResponse, error)
	// IncrementCounter
	//
	// IncrementCounter increments a counter value
	IncrementCounter(context.Context, *IncrementCounterRequest) (*IncrementCounterResponse, error)
	// GetCounter
	//
	// GetCounter gets a counter value
	GetCounter(context.Context, *GetCounterRequest) (*GetCounterResponse, error)
	// ResetCounter
	//
	// ResetCounter resets a counter
	ResetCounter(context.Context, *ResetCounterRequest) (*ResetCounterResponse, error)
	AddKeys(context.Context, *AddKeysRequest) (*AddKeysResponse, error)
	GetKeys(context.Context, *emptypb.Empty) (*GetKeysResponse, error)
	DeleteKey(context.Context, *DeleteKeyRequest) (*emptypb.Empty, error)
	// GetOpenApiSpec
	//
	// Returns the generated openapi swagger document
	GetOpenApiSpec(context.Context, *emptypb.Empty) (*httpbody.HttpBody, error)
	ExportSettings(context.Context, *emptypb.Empty) (*SettingsImportExportData, error)
	ImportSettings(context.Context, *SettingsImportExportData) (*emptypb.Empty, error)
	mustEmbedUnimplementedStorageServiceServer()
}

// UnimplementedStorageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServiceServer struct {
}

func (UnimplementedStorageServiceServer) InsertBlob(context.Context, *InsertBlobRequest) (*InsertBlobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertBlob not implemented")
}
func (UnimplementedStorageServiceServer) ReplaceBlob(context.Context, *ReplaceBlobRequest) (*ReplaceBlobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceBlob not implemented")
}
func (UnimplementedStorageServiceServer) GetBlob(context.Context, *GetBlobRequest) (*GetBlobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlob not implemented")
}
func (UnimplementedStorageServiceServer) DeleteBlob(context.Context, *DeleteBlobRequest) (*DeleteBlobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlob not implemented")
}
func (UnimplementedStorageServiceServer) UpdateAppendBlob(context.Context, *UpdateAppendBlobRequest) (*UpdateAppendBlobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppendBlob not implemented")
}
func (UnimplementedStorageServiceServer) GetAppendBlob(context.Context, *GetAppendBlobRequest) (*GetAppendBlobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppendBlob not implemented")
}
func (UnimplementedStorageServiceServer) DeleteAppendBlob(context.Context, *DeleteAppendBlobRequest) (*DeleteAppendBlobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAppendBlob not implemented")
}
func (UnimplementedStorageServiceServer) DeleteUserData(context.Context, *DeleteUserDataRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserData not implemented")
}
func (UnimplementedStorageServiceServer) GetUserData(context.Context, *GetUserDataRequest) (*GetUserDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserData not implemented")
}
func (UnimplementedStorageServiceServer) IncrementCounter(context.Context, *IncrementCounterRequest) (*IncrementCounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrementCounter not implemented")
}
func (UnimplementedStorageServiceServer) GetCounter(context.Context, *GetCounterRequest) (*GetCounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCounter not implemented")
}
func (UnimplementedStorageServiceServer) ResetCounter(context.Context, *ResetCounterRequest) (*ResetCounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetCounter not implemented")
}
func (UnimplementedStorageServiceServer) AddKeys(context.Context, *AddKeysRequest) (*AddKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddKeys not implemented")
}
func (UnimplementedStorageServiceServer) GetKeys(context.Context, *emptypb.Empty) (*GetKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeys not implemented")
}
func (UnimplementedStorageServiceServer) DeleteKey(context.Context, *DeleteKeyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKey not implemented")
}
func (UnimplementedStorageServiceServer) GetOpenApiSpec(context.Context, *emptypb.Empty) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOpenApiSpec not implemented")
}
func (UnimplementedStorageServiceServer) ExportSettings(context.Context, *emptypb.Empty) (*SettingsImportExportData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportSettings not implemented")
}
func (UnimplementedStorageServiceServer) ImportSettings(context.Context, *SettingsImportExportData) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportSettings not implemented")
}
func (UnimplementedStorageServiceServer) mustEmbedUnimplementedStorageServiceServer() {}

// UnsafeStorageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServiceServer will
// result in compilation errors.
type UnsafeStorageServiceServer interface {
	mustEmbedUnimplementedStorageServiceServer()
}

func RegisterStorageServiceServer(s grpc.ServiceRegistrar, srv StorageServiceServer) {
	s.RegisterService(&StorageService_ServiceDesc, srv)
}

func _StorageService_InsertBlob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertBlobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).InsertBlob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/InsertBlob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).InsertBlob(ctx, req.(*InsertBlobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_ReplaceBlob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceBlobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).ReplaceBlob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/ReplaceBlob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).ReplaceBlob(ctx, req.(*ReplaceBlobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_GetBlob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetBlob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/GetBlob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetBlob(ctx, req.(*GetBlobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_DeleteBlob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBlobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).DeleteBlob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/DeleteBlob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).DeleteBlob(ctx, req.(*DeleteBlobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_UpdateAppendBlob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppendBlobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).UpdateAppendBlob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/UpdateAppendBlob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).UpdateAppendBlob(ctx, req.(*UpdateAppendBlobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_GetAppendBlob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppendBlobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetAppendBlob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/GetAppendBlob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetAppendBlob(ctx, req.(*GetAppendBlobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_DeleteAppendBlob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAppendBlobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).DeleteAppendBlob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/DeleteAppendBlob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).DeleteAppendBlob(ctx, req.(*DeleteAppendBlobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_DeleteUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).DeleteUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/DeleteUserData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).DeleteUserData(ctx, req.(*DeleteUserDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_GetUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/GetUserData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetUserData(ctx, req.(*GetUserDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_IncrementCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrementCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).IncrementCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/IncrementCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).IncrementCounter(ctx, req.(*IncrementCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_GetCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/GetCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetCounter(ctx, req.(*GetCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_ResetCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).ResetCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/ResetCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).ResetCounter(ctx, req.(*ResetCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_AddKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).AddKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/AddKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).AddKeys(ctx, req.(*AddKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_GetKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/GetKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetKeys(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_DeleteKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).DeleteKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/DeleteKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).DeleteKey(ctx, req.(*DeleteKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_GetOpenApiSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetOpenApiSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/GetOpenApiSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetOpenApiSpec(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_ExportSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).ExportSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/ExportSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).ExportSettings(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_ImportSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SettingsImportExportData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).ImportSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/ImportSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).ImportSettings(ctx, req.(*SettingsImportExportData))
	}
	return interceptor(ctx, in, info, handler)
}

// StorageService_ServiceDesc is the grpc.ServiceDesc for StorageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StorageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storage.StorageService",
	HandlerType: (*StorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertBlob",
			Handler:    _StorageService_InsertBlob_Handler,
		},
		{
			MethodName: "ReplaceBlob",
			Handler:    _StorageService_ReplaceBlob_Handler,
		},
		{
			MethodName: "GetBlob",
			Handler:    _StorageService_GetBlob_Handler,
		},
		{
			MethodName: "DeleteBlob",
			Handler:    _StorageService_DeleteBlob_Handler,
		},
		{
			MethodName: "UpdateAppendBlob",
			Handler:    _StorageService_UpdateAppendBlob_Handler,
		},
		{
			MethodName: "GetAppendBlob",
			Handler:    _StorageService_GetAppendBlob_Handler,
		},
		{
			MethodName: "DeleteAppendBlob",
			Handler:    _StorageService_DeleteAppendBlob_Handler,
		},
		{
			MethodName: "DeleteUserData",
			Handler:    _StorageService_DeleteUserData_Handler,
		},
		{
			MethodName: "GetUserData",
			Handler:    _StorageService_GetUserData_Handler,
		},
		{
			MethodName: "IncrementCounter",
			Handler:    _StorageService_IncrementCounter_Handler,
		},
		{
			MethodName: "GetCounter",
			Handler:    _StorageService_GetCounter_Handler,
		},
		{
			MethodName: "ResetCounter",
			Handler:    _StorageService_ResetCounter_Handler,
		},
		{
			MethodName: "AddKeys",
			Handler:    _StorageService_AddKeys_Handler,
		},
		{
			MethodName: "GetKeys",
			Handler:    _StorageService_GetKeys_Handler,
		},
		{
			MethodName: "DeleteKey",
			Handler:    _StorageService_DeleteKey_Handler,
		},
		{
			MethodName: "GetOpenApiSpec",
			Handler:    _StorageService_GetOpenApiSpec_Handler,
		},
		{
			MethodName: "ExportSettings",
			Handler:    _StorageService_ExportSettings_Handler,
		},
		{
			MethodName: "ImportSettings",
			Handler:    _StorageService_ImportSettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storage_service.proto",
}
