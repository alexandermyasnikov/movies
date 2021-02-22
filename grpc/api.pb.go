// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: grpc/api.proto

package moviesgrpc

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_grpc_api_proto_rawDescGZIP(), []int{0}
}

type MovieKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MovieKey) Reset() {
	*x = MovieKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieKey) ProtoMessage() {}

func (x *MovieKey) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieKey.ProtoReflect.Descriptor instead.
func (*MovieKey) Descriptor() ([]byte, []int) {
	return file_grpc_api_proto_rawDescGZIP(), []int{1}
}

func (x *MovieKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Json string `protobuf:"bytes,2,opt,name=json,proto3" json:"json,omitempty"`
}

func (x *Movie) Reset() {
	*x = Movie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_grpc_api_proto_rawDescGZIP(), []int{2}
}

func (x *Movie) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Movie) GetJson() string {
	if x != nil {
		return x.Json
	}
	return ""
}

type Stats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MoviesCount int32 `protobuf:"varint,1,opt,name=moviesCount,proto3" json:"moviesCount,omitempty"`
}

func (x *Stats) Reset() {
	*x = Stats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stats) ProtoMessage() {}

func (x *Stats) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stats.ProtoReflect.Descriptor instead.
func (*Stats) Descriptor() ([]byte, []int) {
	return file_grpc_api_proto_rawDescGZIP(), []int{3}
}

func (x *Stats) GetMoviesCount() int32 {
	if x != nil {
		return x.MoviesCount
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOk  bool   `protobuf:"varint,1,opt,name=isOk,proto3" json:"isOk,omitempty"`
	Stats *Stats `protobuf:"bytes,2,opt,name=stats,proto3" json:"stats,omitempty"`
	Movie *Movie `protobuf:"bytes,3,opt,name=movie,proto3" json:"movie,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_grpc_api_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetIsOk() bool {
	if x != nil {
		return x.IsOk
	}
	return false
}

func (x *Response) GetStats() *Stats {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *Response) GetMovie() *Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

var File_grpc_api_proto protoreflect.FileDescriptor

var file_grpc_api_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x67, 0x72, 0x70, 0x63, 0x22, 0x06, 0x0a, 0x04,
	0x56, 0x6f, 0x69, 0x64, 0x22, 0x1a, 0x0a, 0x08, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4b, 0x65, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x2b, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6a, 0x73, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x29, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x70, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6b, 0x12, 0x27, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x27, 0x0a, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x52, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x32, 0xe0, 0x01, 0x0a, 0x0d, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x73, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x08,
	0x67, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x10, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x73, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x11, 0x2e, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x73, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x33, 0x0a,
	0x08, 0x67, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x14, 0x2e, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x73, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4b, 0x65, 0x79, 0x1a,
	0x11, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x12, 0x35, 0x0a, 0x0e, 0x67, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x61,
	0x6e, 0x64, 0x6f, 0x6d, 0x12, 0x10, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x11, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x32, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x11, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x73, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x1a, 0x10, 0x2e, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x73, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_api_proto_rawDescOnce sync.Once
	file_grpc_api_proto_rawDescData = file_grpc_api_proto_rawDesc
)

func file_grpc_api_proto_rawDescGZIP() []byte {
	file_grpc_api_proto_rawDescOnce.Do(func() {
		file_grpc_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_api_proto_rawDescData)
	})
	return file_grpc_api_proto_rawDescData
}

var file_grpc_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_grpc_api_proto_goTypes = []interface{}{
	(*Void)(nil),     // 0: moviesgrpc.Void
	(*MovieKey)(nil), // 1: moviesgrpc.MovieKey
	(*Movie)(nil),    // 2: moviesgrpc.Movie
	(*Stats)(nil),    // 3: moviesgrpc.Stats
	(*Response)(nil), // 4: moviesgrpc.Response
}
var file_grpc_api_proto_depIdxs = []int32{
	3, // 0: moviesgrpc.Response.stats:type_name -> moviesgrpc.Stats
	2, // 1: moviesgrpc.Response.movie:type_name -> moviesgrpc.Movie
	0, // 2: moviesgrpc.MoviesStorage.getStats:input_type -> moviesgrpc.Void
	1, // 3: moviesgrpc.MoviesStorage.getMovie:input_type -> moviesgrpc.MovieKey
	0, // 4: moviesgrpc.MoviesStorage.getMovieRandom:input_type -> moviesgrpc.Void
	2, // 5: moviesgrpc.MoviesStorage.updateMovie:input_type -> moviesgrpc.Movie
	3, // 6: moviesgrpc.MoviesStorage.getStats:output_type -> moviesgrpc.Stats
	2, // 7: moviesgrpc.MoviesStorage.getMovie:output_type -> moviesgrpc.Movie
	2, // 8: moviesgrpc.MoviesStorage.getMovieRandom:output_type -> moviesgrpc.Movie
	0, // 9: moviesgrpc.MoviesStorage.updateMovie:output_type -> moviesgrpc.Void
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_grpc_api_proto_init() }
func file_grpc_api_proto_init() {
	if File_grpc_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieKey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Movie); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stats); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_api_proto_goTypes,
		DependencyIndexes: file_grpc_api_proto_depIdxs,
		MessageInfos:      file_grpc_api_proto_msgTypes,
	}.Build()
	File_grpc_api_proto = out.File
	file_grpc_api_proto_rawDesc = nil
	file_grpc_api_proto_goTypes = nil
	file_grpc_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MoviesStorageClient is the client API for MoviesStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MoviesStorageClient interface {
	GetStats(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Stats, error)
	GetMovie(ctx context.Context, in *MovieKey, opts ...grpc.CallOption) (*Movie, error)
	GetMovieRandom(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Movie, error)
	UpdateMovie(ctx context.Context, in *Movie, opts ...grpc.CallOption) (*Void, error)
}

type moviesStorageClient struct {
	cc grpc.ClientConnInterface
}

func NewMoviesStorageClient(cc grpc.ClientConnInterface) MoviesStorageClient {
	return &moviesStorageClient{cc}
}

func (c *moviesStorageClient) GetStats(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Stats, error) {
	out := new(Stats)
	err := c.cc.Invoke(ctx, "/moviesgrpc.MoviesStorage/getStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesStorageClient) GetMovie(ctx context.Context, in *MovieKey, opts ...grpc.CallOption) (*Movie, error) {
	out := new(Movie)
	err := c.cc.Invoke(ctx, "/moviesgrpc.MoviesStorage/getMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesStorageClient) GetMovieRandom(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Movie, error) {
	out := new(Movie)
	err := c.cc.Invoke(ctx, "/moviesgrpc.MoviesStorage/getMovieRandom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesStorageClient) UpdateMovie(ctx context.Context, in *Movie, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/moviesgrpc.MoviesStorage/updateMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MoviesStorageServer is the server API for MoviesStorage service.
type MoviesStorageServer interface {
	GetStats(context.Context, *Void) (*Stats, error)
	GetMovie(context.Context, *MovieKey) (*Movie, error)
	GetMovieRandom(context.Context, *Void) (*Movie, error)
	UpdateMovie(context.Context, *Movie) (*Void, error)
}

// UnimplementedMoviesStorageServer can be embedded to have forward compatible implementations.
type UnimplementedMoviesStorageServer struct {
}

func (*UnimplementedMoviesStorageServer) GetStats(context.Context, *Void) (*Stats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}
func (*UnimplementedMoviesStorageServer) GetMovie(context.Context, *MovieKey) (*Movie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovie not implemented")
}
func (*UnimplementedMoviesStorageServer) GetMovieRandom(context.Context, *Void) (*Movie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovieRandom not implemented")
}
func (*UnimplementedMoviesStorageServer) UpdateMovie(context.Context, *Movie) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMovie not implemented")
}

func RegisterMoviesStorageServer(s *grpc.Server, srv MoviesStorageServer) {
	s.RegisterService(&_MoviesStorage_serviceDesc, srv)
}

func _MoviesStorage_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesStorageServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moviesgrpc.MoviesStorage/GetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesStorageServer).GetStats(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoviesStorage_GetMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesStorageServer).GetMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moviesgrpc.MoviesStorage/GetMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesStorageServer).GetMovie(ctx, req.(*MovieKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoviesStorage_GetMovieRandom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesStorageServer).GetMovieRandom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moviesgrpc.MoviesStorage/GetMovieRandom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesStorageServer).GetMovieRandom(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoviesStorage_UpdateMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Movie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesStorageServer).UpdateMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moviesgrpc.MoviesStorage/UpdateMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesStorageServer).UpdateMovie(ctx, req.(*Movie))
	}
	return interceptor(ctx, in, info, handler)
}

var _MoviesStorage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moviesgrpc.MoviesStorage",
	HandlerType: (*MoviesStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getStats",
			Handler:    _MoviesStorage_GetStats_Handler,
		},
		{
			MethodName: "getMovie",
			Handler:    _MoviesStorage_GetMovie_Handler,
		},
		{
			MethodName: "getMovieRandom",
			Handler:    _MoviesStorage_GetMovieRandom_Handler,
		},
		{
			MethodName: "updateMovie",
			Handler:    _MoviesStorage_UpdateMovie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/api.proto",
}
