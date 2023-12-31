// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: library.proto

package library

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LibraryClient is the client API for Library service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LibraryClient interface {
	// CUD == C-create, U-update, D-delete DataBase sql functionality
	HandleAuthorCUD(ctx context.Context, in *CUDAuthorRequest, opts ...grpc.CallOption) (*DefaultResponse, error)
	HandleBookCUD(ctx context.Context, in *CUDBookRequest, opts ...grpc.CallOption) (*DefaultResponse, error)
	// Get handlers
	GetAllAuthors(ctx context.Context, in *DefaultRequest, opts ...grpc.CallOption) (*Authors, error)
	GetAllBooks(ctx context.Context, in *DefaultRequest, opts ...grpc.CallOption) (*Books, error)
	GetAuthorById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Author, error)
	GetBookById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Book, error)
	GetAuthorsByBookName(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Authors, error)
	GetBooksByAuthorId(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Books, error)
}

type libraryClient struct {
	cc grpc.ClientConnInterface
}

func NewLibraryClient(cc grpc.ClientConnInterface) LibraryClient {
	return &libraryClient{cc}
}

func (c *libraryClient) HandleAuthorCUD(ctx context.Context, in *CUDAuthorRequest, opts ...grpc.CallOption) (*DefaultResponse, error) {
	out := new(DefaultResponse)
	err := c.cc.Invoke(ctx, "/Library/HandleAuthorCUD", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryClient) HandleBookCUD(ctx context.Context, in *CUDBookRequest, opts ...grpc.CallOption) (*DefaultResponse, error) {
	out := new(DefaultResponse)
	err := c.cc.Invoke(ctx, "/Library/HandleBookCUD", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryClient) GetAllAuthors(ctx context.Context, in *DefaultRequest, opts ...grpc.CallOption) (*Authors, error) {
	out := new(Authors)
	err := c.cc.Invoke(ctx, "/Library/GetAllAuthors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryClient) GetAllBooks(ctx context.Context, in *DefaultRequest, opts ...grpc.CallOption) (*Books, error) {
	out := new(Books)
	err := c.cc.Invoke(ctx, "/Library/GetAllBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryClient) GetAuthorById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, "/Library/GetAuthorById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryClient) GetBookById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/Library/GetBookById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryClient) GetAuthorsByBookName(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Authors, error) {
	out := new(Authors)
	err := c.cc.Invoke(ctx, "/Library/GetAuthorsByBookName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryClient) GetBooksByAuthorId(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Books, error) {
	out := new(Books)
	err := c.cc.Invoke(ctx, "/Library/GetBooksByAuthorId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LibraryServer is the server API for Library service.
// All implementations must embed UnimplementedLibraryServer
// for forward compatibility
type LibraryServer interface {
	// CUD == C-create, U-update, D-delete DataBase sql functionality
	HandleAuthorCUD(context.Context, *CUDAuthorRequest) (*DefaultResponse, error)
	HandleBookCUD(context.Context, *CUDBookRequest) (*DefaultResponse, error)
	// Get handlers
	GetAllAuthors(context.Context, *DefaultRequest) (*Authors, error)
	GetAllBooks(context.Context, *DefaultRequest) (*Books, error)
	GetAuthorById(context.Context, *IdRequest) (*Author, error)
	GetBookById(context.Context, *IdRequest) (*Book, error)
	GetAuthorsByBookName(context.Context, *Book) (*Authors, error)
	GetBooksByAuthorId(context.Context, *IdRequest) (*Books, error)
	mustEmbedUnimplementedLibraryServer()
}

// UnimplementedLibraryServer must be embedded to have forward compatible implementations.
type UnimplementedLibraryServer struct {
}

func (UnimplementedLibraryServer) HandleAuthorCUD(context.Context, *CUDAuthorRequest) (*DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleAuthorCUD not implemented")
}
func (UnimplementedLibraryServer) HandleBookCUD(context.Context, *CUDBookRequest) (*DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleBookCUD not implemented")
}
func (UnimplementedLibraryServer) GetAllAuthors(context.Context, *DefaultRequest) (*Authors, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAuthors not implemented")
}
func (UnimplementedLibraryServer) GetAllBooks(context.Context, *DefaultRequest) (*Books, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllBooks not implemented")
}
func (UnimplementedLibraryServer) GetAuthorById(context.Context, *IdRequest) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthorById not implemented")
}
func (UnimplementedLibraryServer) GetBookById(context.Context, *IdRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookById not implemented")
}
func (UnimplementedLibraryServer) GetAuthorsByBookName(context.Context, *Book) (*Authors, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthorsByBookName not implemented")
}
func (UnimplementedLibraryServer) GetBooksByAuthorId(context.Context, *IdRequest) (*Books, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBooksByAuthorId not implemented")
}
func (UnimplementedLibraryServer) mustEmbedUnimplementedLibraryServer() {}

// UnsafeLibraryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LibraryServer will
// result in compilation errors.
type UnsafeLibraryServer interface {
	mustEmbedUnimplementedLibraryServer()
}

func RegisterLibraryServer(s grpc.ServiceRegistrar, srv LibraryServer) {
	s.RegisterService(&Library_ServiceDesc, srv)
}

func _Library_HandleAuthorCUD_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CUDAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServer).HandleAuthorCUD(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Library/HandleAuthorCUD",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServer).HandleAuthorCUD(ctx, req.(*CUDAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Library_HandleBookCUD_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CUDBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServer).HandleBookCUD(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Library/HandleBookCUD",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServer).HandleBookCUD(ctx, req.(*CUDBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Library_GetAllAuthors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DefaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServer).GetAllAuthors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Library/GetAllAuthors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServer).GetAllAuthors(ctx, req.(*DefaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Library_GetAllBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DefaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServer).GetAllBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Library/GetAllBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServer).GetAllBooks(ctx, req.(*DefaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Library_GetAuthorById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServer).GetAuthorById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Library/GetAuthorById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServer).GetAuthorById(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Library_GetBookById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServer).GetBookById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Library/GetBookById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServer).GetBookById(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Library_GetAuthorsByBookName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServer).GetAuthorsByBookName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Library/GetAuthorsByBookName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServer).GetAuthorsByBookName(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _Library_GetBooksByAuthorId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServer).GetBooksByAuthorId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Library/GetBooksByAuthorId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServer).GetBooksByAuthorId(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Library_ServiceDesc is the grpc.ServiceDesc for Library service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Library_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Library",
	HandlerType: (*LibraryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleAuthorCUD",
			Handler:    _Library_HandleAuthorCUD_Handler,
		},
		{
			MethodName: "HandleBookCUD",
			Handler:    _Library_HandleBookCUD_Handler,
		},
		{
			MethodName: "GetAllAuthors",
			Handler:    _Library_GetAllAuthors_Handler,
		},
		{
			MethodName: "GetAllBooks",
			Handler:    _Library_GetAllBooks_Handler,
		},
		{
			MethodName: "GetAuthorById",
			Handler:    _Library_GetAuthorById_Handler,
		},
		{
			MethodName: "GetBookById",
			Handler:    _Library_GetBookById_Handler,
		},
		{
			MethodName: "GetAuthorsByBookName",
			Handler:    _Library_GetAuthorsByBookName_Handler,
		},
		{
			MethodName: "GetBooksByAuthorId",
			Handler:    _Library_GetBooksByAuthorId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "library.proto",
}
