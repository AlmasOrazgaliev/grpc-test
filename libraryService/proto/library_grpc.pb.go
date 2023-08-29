// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.1
// source: proto/library.proto

package __

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

// BookClient is the client API for Book service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookClient interface {
	List(ctx context.Context, in *BookData, opts ...grpc.CallOption) (*ListBook, error)
	Add(ctx context.Context, in *BookData, opts ...grpc.CallOption) (*BookData, error)
	Get(ctx context.Context, in *BookData, opts ...grpc.CallOption) (*BookData, error)
	Update(ctx context.Context, in *BookData, opts ...grpc.CallOption) (*BookData, error)
	Delete(ctx context.Context, in *BookData, opts ...grpc.CallOption) (*BookData, error)
}

type bookClient struct {
	cc grpc.ClientConnInterface
}

func NewBookClient(cc grpc.ClientConnInterface) BookClient {
	return &bookClient{cc}
}

func (c *bookClient) List(ctx context.Context, in *BookData, opts ...grpc.CallOption) (*ListBook, error) {
	out := new(ListBook)
	err := c.cc.Invoke(ctx, "/bookService.Book/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookClient) Add(ctx context.Context, in *BookData, opts ...grpc.CallOption) (*BookData, error) {
	out := new(BookData)
	err := c.cc.Invoke(ctx, "/bookService.Book/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookClient) Get(ctx context.Context, in *BookData, opts ...grpc.CallOption) (*BookData, error) {
	out := new(BookData)
	err := c.cc.Invoke(ctx, "/bookService.Book/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookClient) Update(ctx context.Context, in *BookData, opts ...grpc.CallOption) (*BookData, error) {
	out := new(BookData)
	err := c.cc.Invoke(ctx, "/bookService.Book/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookClient) Delete(ctx context.Context, in *BookData, opts ...grpc.CallOption) (*BookData, error) {
	out := new(BookData)
	err := c.cc.Invoke(ctx, "/bookService.Book/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServer is the server API for Book service.
// All implementations must embed UnimplementedBookServer
// for forward compatibility
type BookServer interface {
	List(context.Context, *BookData) (*ListBook, error)
	Add(context.Context, *BookData) (*BookData, error)
	Get(context.Context, *BookData) (*BookData, error)
	Update(context.Context, *BookData) (*BookData, error)
	Delete(context.Context, *BookData) (*BookData, error)
	mustEmbedUnimplementedBookServer()
}

// UnimplementedBookServer must be embedded to have forward compatible implementations.
type UnimplementedBookServer struct {
}

func (UnimplementedBookServer) List(context.Context, *BookData) (*ListBook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedBookServer) Add(context.Context, *BookData) (*BookData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedBookServer) Get(context.Context, *BookData) (*BookData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBookServer) Update(context.Context, *BookData) (*BookData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBookServer) Delete(context.Context, *BookData) (*BookData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBookServer) mustEmbedUnimplementedBookServer() {}

// UnsafeBookServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServer will
// result in compilation errors.
type UnsafeBookServer interface {
	mustEmbedUnimplementedBookServer()
}

func RegisterBookServer(s grpc.ServiceRegistrar, srv BookServer) {
	s.RegisterService(&Book_ServiceDesc, srv)
}

func _Book_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookService.Book/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).List(ctx, req.(*BookData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Book_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookService.Book/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).Add(ctx, req.(*BookData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Book_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookService.Book/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).Get(ctx, req.(*BookData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Book_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookService.Book/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).Update(ctx, req.(*BookData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Book_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookService.Book/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).Delete(ctx, req.(*BookData))
	}
	return interceptor(ctx, in, info, handler)
}

// Book_ServiceDesc is the grpc.ServiceDesc for Book service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Book_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bookService.Book",
	HandlerType: (*BookServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Book_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Book_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Book_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Book_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Book_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/library.proto",
}

// AuthorClient is the client API for Author service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorClient interface {
	List(ctx context.Context, in *AuthorData, opts ...grpc.CallOption) (*ListAuthor, error)
	Add(ctx context.Context, in *AuthorData, opts ...grpc.CallOption) (*AuthorData, error)
	Get(ctx context.Context, in *AuthorData, opts ...grpc.CallOption) (*AuthorData, error)
	Update(ctx context.Context, in *AuthorData, opts ...grpc.CallOption) (*AuthorData, error)
	Delete(ctx context.Context, in *AuthorData, opts ...grpc.CallOption) (*AuthorData, error)
}

type authorClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorClient(cc grpc.ClientConnInterface) AuthorClient {
	return &authorClient{cc}
}

func (c *authorClient) List(ctx context.Context, in *AuthorData, opts ...grpc.CallOption) (*ListAuthor, error) {
	out := new(ListAuthor)
	err := c.cc.Invoke(ctx, "/bookService.Author/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorClient) Add(ctx context.Context, in *AuthorData, opts ...grpc.CallOption) (*AuthorData, error) {
	out := new(AuthorData)
	err := c.cc.Invoke(ctx, "/bookService.Author/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorClient) Get(ctx context.Context, in *AuthorData, opts ...grpc.CallOption) (*AuthorData, error) {
	out := new(AuthorData)
	err := c.cc.Invoke(ctx, "/bookService.Author/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorClient) Update(ctx context.Context, in *AuthorData, opts ...grpc.CallOption) (*AuthorData, error) {
	out := new(AuthorData)
	err := c.cc.Invoke(ctx, "/bookService.Author/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorClient) Delete(ctx context.Context, in *AuthorData, opts ...grpc.CallOption) (*AuthorData, error) {
	out := new(AuthorData)
	err := c.cc.Invoke(ctx, "/bookService.Author/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorServer is the server API for Author service.
// All implementations must embed UnimplementedAuthorServer
// for forward compatibility
type AuthorServer interface {
	List(context.Context, *AuthorData) (*ListAuthor, error)
	Add(context.Context, *AuthorData) (*AuthorData, error)
	Get(context.Context, *AuthorData) (*AuthorData, error)
	Update(context.Context, *AuthorData) (*AuthorData, error)
	Delete(context.Context, *AuthorData) (*AuthorData, error)
	mustEmbedUnimplementedAuthorServer()
}

// UnimplementedAuthorServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorServer struct {
}

func (UnimplementedAuthorServer) List(context.Context, *AuthorData) (*ListAuthor, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAuthorServer) Add(context.Context, *AuthorData) (*AuthorData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedAuthorServer) Get(context.Context, *AuthorData) (*AuthorData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAuthorServer) Update(context.Context, *AuthorData) (*AuthorData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAuthorServer) Delete(context.Context, *AuthorData) (*AuthorData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAuthorServer) mustEmbedUnimplementedAuthorServer() {}

// UnsafeAuthorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorServer will
// result in compilation errors.
type UnsafeAuthorServer interface {
	mustEmbedUnimplementedAuthorServer()
}

func RegisterAuthorServer(s grpc.ServiceRegistrar, srv AuthorServer) {
	s.RegisterService(&Author_ServiceDesc, srv)
}

func _Author_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookService.Author/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServer).List(ctx, req.(*AuthorData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Author_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookService.Author/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServer).Add(ctx, req.(*AuthorData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Author_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookService.Author/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServer).Get(ctx, req.(*AuthorData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Author_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookService.Author/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServer).Update(ctx, req.(*AuthorData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Author_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookService.Author/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServer).Delete(ctx, req.(*AuthorData))
	}
	return interceptor(ctx, in, info, handler)
}

// Author_ServiceDesc is the grpc.ServiceDesc for Author service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Author_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bookService.Author",
	HandlerType: (*AuthorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Author_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Author_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Author_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Author_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Author_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/library.proto",
}

// MemberClient is the client API for Member service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MemberClient interface {
	List(ctx context.Context, in *MemberData, opts ...grpc.CallOption) (*ListMember, error)
	Add(ctx context.Context, in *MemberData, opts ...grpc.CallOption) (*MemberData, error)
	Get(ctx context.Context, in *MemberData, opts ...grpc.CallOption) (*MemberData, error)
	Update(ctx context.Context, in *MemberData, opts ...grpc.CallOption) (*MemberData, error)
	Delete(ctx context.Context, in *MemberData, opts ...grpc.CallOption) (*MemberData, error)
}

type memberClient struct {
	cc grpc.ClientConnInterface
}

func NewMemberClient(cc grpc.ClientConnInterface) MemberClient {
	return &memberClient{cc}
}

func (c *memberClient) List(ctx context.Context, in *MemberData, opts ...grpc.CallOption) (*ListMember, error) {
	out := new(ListMember)
	err := c.cc.Invoke(ctx, "/bookService.Member/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberClient) Add(ctx context.Context, in *MemberData, opts ...grpc.CallOption) (*MemberData, error) {
	out := new(MemberData)
	err := c.cc.Invoke(ctx, "/bookService.Member/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberClient) Get(ctx context.Context, in *MemberData, opts ...grpc.CallOption) (*MemberData, error) {
	out := new(MemberData)
	err := c.cc.Invoke(ctx, "/bookService.Member/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberClient) Update(ctx context.Context, in *MemberData, opts ...grpc.CallOption) (*MemberData, error) {
	out := new(MemberData)
	err := c.cc.Invoke(ctx, "/bookService.Member/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberClient) Delete(ctx context.Context, in *MemberData, opts ...grpc.CallOption) (*MemberData, error) {
	out := new(MemberData)
	err := c.cc.Invoke(ctx, "/bookService.Member/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemberServer is the server API for Member service.
// All implementations must embed UnimplementedMemberServer
// for forward compatibility
type MemberServer interface {
	List(context.Context, *MemberData) (*ListMember, error)
	Add(context.Context, *MemberData) (*MemberData, error)
	Get(context.Context, *MemberData) (*MemberData, error)
	Update(context.Context, *MemberData) (*MemberData, error)
	Delete(context.Context, *MemberData) (*MemberData, error)
	mustEmbedUnimplementedMemberServer()
}

// UnimplementedMemberServer must be embedded to have forward compatible implementations.
type UnimplementedMemberServer struct {
}

func (UnimplementedMemberServer) List(context.Context, *MemberData) (*ListMember, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedMemberServer) Add(context.Context, *MemberData) (*MemberData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedMemberServer) Get(context.Context, *MemberData) (*MemberData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedMemberServer) Update(context.Context, *MemberData) (*MemberData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMemberServer) Delete(context.Context, *MemberData) (*MemberData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedMemberServer) mustEmbedUnimplementedMemberServer() {}

// UnsafeMemberServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MemberServer will
// result in compilation errors.
type UnsafeMemberServer interface {
	mustEmbedUnimplementedMemberServer()
}

func RegisterMemberServer(s grpc.ServiceRegistrar, srv MemberServer) {
	s.RegisterService(&Member_ServiceDesc, srv)
}

func _Member_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookService.Member/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).List(ctx, req.(*MemberData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Member_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookService.Member/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).Add(ctx, req.(*MemberData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Member_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookService.Member/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).Get(ctx, req.(*MemberData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Member_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookService.Member/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).Update(ctx, req.(*MemberData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Member_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookService.Member/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).Delete(ctx, req.(*MemberData))
	}
	return interceptor(ctx, in, info, handler)
}

// Member_ServiceDesc is the grpc.ServiceDesc for Member service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Member_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bookService.Member",
	HandlerType: (*MemberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Member_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Member_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Member_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Member_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Member_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/library.proto",
}
