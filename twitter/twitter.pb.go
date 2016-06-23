// Code generated by protoc-gen-go.
// source: twitter/twitter.proto
// DO NOT EDIT!

/*
Package twitter is a generated protocol buffer package.

It is generated from these files:
	twitter/twitter.proto

It has these top-level messages:
	User
	Tweet
	Search
	Timeline
*/
package twitter

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

type User struct {
	ID uint64 `protobuf:"varint,1,opt,name=ID,json=iD" json:"ID,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Tweet struct {
	ID   uint64 `protobuf:"varint,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	User *User  `protobuf:"bytes,3,opt,name=user" json:"user,omitempty"`
}

func (m *Tweet) Reset()                    { *m = Tweet{} }
func (m *Tweet) String() string            { return proto.CompactTextString(m) }
func (*Tweet) ProtoMessage()               {}
func (*Tweet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Tweet) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type Search struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *Search) Reset()                    { *m = Search{} }
func (m *Search) String() string            { return proto.CompactTextString(m) }
func (*Search) ProtoMessage()               {}
func (*Search) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Timeline struct {
	Tweets []*Tweet `protobuf:"bytes,1,rep,name=tweets" json:"tweets,omitempty"`
}

func (m *Timeline) Reset()                    { *m = Timeline{} }
func (m *Timeline) String() string            { return proto.CompactTextString(m) }
func (*Timeline) ProtoMessage()               {}
func (*Timeline) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Timeline) GetTweets() []*Tweet {
	if m != nil {
		return m.Tweets
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "twitter.User")
	proto.RegisterType((*Tweet)(nil), "twitter.Tweet")
	proto.RegisterType((*Search)(nil), "twitter.Search")
	proto.RegisterType((*Timeline)(nil), "twitter.Timeline")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Twitter service

type TwitterClient interface {
	// We have a method called `GetTimeline` which takes
	// parameter called `User` and returns
	// the user `Timeline` (list of `Tweets`).
	GetTimeline(ctx context.Context, in *User, opts ...grpc.CallOption) (*Timeline, error)
	// We have a method called `UserStream` which takes
	// parameter called `Search` and returns
	// an stream of `Tweets`
	UserStream(ctx context.Context, in *Search, opts ...grpc.CallOption) (Twitter_UserStreamClient, error)
}

type twitterClient struct {
	cc *grpc.ClientConn
}

func NewTwitterClient(cc *grpc.ClientConn) TwitterClient {
	return &twitterClient{cc}
}

func (c *twitterClient) GetTimeline(ctx context.Context, in *User, opts ...grpc.CallOption) (*Timeline, error) {
	out := new(Timeline)
	err := grpc.Invoke(ctx, "/twitter.Twitter/GetTimeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitterClient) UserStream(ctx context.Context, in *Search, opts ...grpc.CallOption) (Twitter_UserStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Twitter_serviceDesc.Streams[0], c.cc, "/twitter.Twitter/UserStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &twitterUserStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Twitter_UserStreamClient interface {
	Recv() (*Tweet, error)
	grpc.ClientStream
}

type twitterUserStreamClient struct {
	grpc.ClientStream
}

func (x *twitterUserStreamClient) Recv() (*Tweet, error) {
	m := new(Tweet)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Twitter service

type TwitterServer interface {
	// We have a method called `GetTimeline` which takes
	// parameter called `User` and returns
	// the user `Timeline` (list of `Tweets`).
	GetTimeline(context.Context, *User) (*Timeline, error)
	// We have a method called `UserStream` which takes
	// parameter called `Search` and returns
	// an stream of `Tweets`
	UserStream(*Search, Twitter_UserStreamServer) error
}

func RegisterTwitterServer(s *grpc.Server, srv TwitterServer) {
	s.RegisterService(&_Twitter_serviceDesc, srv)
}

func _Twitter_GetTimeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitterServer).GetTimeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/twitter.Twitter/GetTimeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitterServer).GetTimeline(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Twitter_UserStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Search)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TwitterServer).UserStream(m, &twitterUserStreamServer{stream})
}

type Twitter_UserStreamServer interface {
	Send(*Tweet) error
	grpc.ServerStream
}

type twitterUserStreamServer struct {
	grpc.ServerStream
}

func (x *twitterUserStreamServer) Send(m *Tweet) error {
	return x.ServerStream.SendMsg(m)
}

var _Twitter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "twitter.Twitter",
	HandlerType: (*TwitterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTimeline",
			Handler:    _Twitter_GetTimeline_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UserStream",
			Handler:       _Twitter_UserStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("twitter/twitter.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x29, 0xcf, 0x2c,
	0x29, 0x49, 0x2d, 0xd2, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae,
	0x92, 0x18, 0x17, 0x4b, 0x68, 0x71, 0x6a, 0x91, 0x10, 0x1f, 0x17, 0x93, 0xa7, 0x8b, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x4b, 0x10, 0x53, 0xa6, 0x8b, 0x92, 0x1f, 0x17, 0x6b, 0x48, 0x79, 0x6a, 0x6a,
	0x09, 0xba, 0x84, 0x90, 0x10, 0x17, 0x4b, 0x49, 0x6a, 0x45, 0x89, 0x04, 0x13, 0x50, 0x84, 0x33,
	0x08, 0xcc, 0x16, 0x52, 0xe4, 0x62, 0x29, 0x05, 0x1a, 0x22, 0xc1, 0x0c, 0x14, 0xe3, 0x36, 0xe2,
	0xd5, 0x83, 0xd9, 0x05, 0x32, 0x39, 0x08, 0x2c, 0xa5, 0x24, 0xc3, 0xc5, 0x16, 0x9c, 0x9a, 0x58,
	0x94, 0x9c, 0x01, 0x37, 0x80, 0x11, 0x61, 0x80, 0x92, 0x11, 0x17, 0x47, 0x48, 0x66, 0x6e, 0x6a,
	0x4e, 0x66, 0x5e, 0xaa, 0x90, 0x1a, 0x17, 0x5b, 0x09, 0xc8, 0xe6, 0x62, 0xa0, 0x0a, 0x66, 0xa0,
	0x71, 0x7c, 0x70, 0xe3, 0xc0, 0x0e, 0x0a, 0x82, 0xca, 0x1a, 0xe5, 0x73, 0xb1, 0x87, 0x40, 0x24,
	0x84, 0x0c, 0xb9, 0xb8, 0xdd, 0x53, 0x4b, 0xe0, 0x26, 0xa0, 0x3a, 0x40, 0x4a, 0x10, 0x61, 0x00,
	0x54, 0x85, 0x12, 0x03, 0x50, 0x0b, 0x17, 0x48, 0x32, 0xb8, 0xa4, 0x28, 0x35, 0x31, 0x57, 0x88,
	0x1f, 0xae, 0x04, 0xe2, 0x48, 0x29, 0x34, 0x4b, 0x95, 0x18, 0x0c, 0x18, 0x93, 0xd8, 0xc0, 0x41,
	0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x0d, 0x23, 0xc2, 0x53, 0x01, 0x00, 0x00,
}
