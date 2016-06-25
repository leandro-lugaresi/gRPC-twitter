// Code generated by protoc-gen-go.
// source: twitter/twitter.proto
// DO NOT EDIT!

/*
Package twitter is a generated protocol buffer package.

It is generated from these files:
	twitter/twitter.proto

It has these top-level messages:
	Token
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

type Token struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	SecretToken string `protobuf:"bytes,2,opt,name=secret_token,json=secretToken" json:"secret_token,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type User struct {
	Id         int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	ScreenName string `protobuf:"bytes,3,opt,name=screen_name,json=screenName" json:"screen_name,omitempty"`
	Url        string `protobuf:"bytes,4,opt,name=url" json:"url,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Tweet struct {
	Id            int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Text          string `protobuf:"bytes,2,opt,name=Text,json=text" json:"Text,omitempty"`
	User          *User  `protobuf:"bytes,3,opt,name=user" json:"user,omitempty"`
	Retweeted     bool   `protobuf:"varint,4,opt,name=retweeted" json:"retweeted,omitempty"`
	RetweetCount  int32  `protobuf:"varint,5,opt,name=retweet_count,json=retweetCount" json:"retweet_count,omitempty"`
	Favorited     bool   `protobuf:"varint,6,opt,name=favorited" json:"favorited,omitempty"`
	FavoriteCount int32  `protobuf:"varint,7,opt,name=favorite_count,json=favoriteCount" json:"favorite_count,omitempty"`
}

func (m *Tweet) Reset()                    { *m = Tweet{} }
func (m *Tweet) String() string            { return proto.CompactTextString(m) }
func (*Tweet) ProtoMessage()               {}
func (*Tweet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Tweet) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type Search struct {
	Text  string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	Token *Token `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *Search) Reset()                    { *m = Search{} }
func (m *Search) String() string            { return proto.CompactTextString(m) }
func (*Search) ProtoMessage()               {}
func (*Search) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Search) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type Timeline struct {
	Tweets []*Tweet `protobuf:"bytes,1,rep,name=tweets" json:"tweets,omitempty"`
}

func (m *Timeline) Reset()                    { *m = Timeline{} }
func (m *Timeline) String() string            { return proto.CompactTextString(m) }
func (*Timeline) ProtoMessage()               {}
func (*Timeline) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Timeline) GetTweets() []*Tweet {
	if m != nil {
		return m.Tweets
	}
	return nil
}

func init() {
	proto.RegisterType((*Token)(nil), "twitter.Token")
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
	// parameter called `Token` and returns
	// the user `Timeline` (list of `Tweets`).
	GetTimeline(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Timeline, error)
	// We have a method called `UserStream` which takes
	// parameter called `Token`and returns
	// an stream of `Tweets`
	UserStream(ctx context.Context, in *Token, opts ...grpc.CallOption) (Twitter_UserStreamClient, error)
	// We have a method called `Filter` which takes
	// parameter called `Search` and returns
	// an stream of `Tweets`
	Filter(ctx context.Context, in *Search, opts ...grpc.CallOption) (Twitter_FilterClient, error)
}

type twitterClient struct {
	cc *grpc.ClientConn
}

func NewTwitterClient(cc *grpc.ClientConn) TwitterClient {
	return &twitterClient{cc}
}

func (c *twitterClient) GetTimeline(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Timeline, error) {
	out := new(Timeline)
	err := grpc.Invoke(ctx, "/twitter.Twitter/GetTimeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitterClient) UserStream(ctx context.Context, in *Token, opts ...grpc.CallOption) (Twitter_UserStreamClient, error) {
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

func (c *twitterClient) Filter(ctx context.Context, in *Search, opts ...grpc.CallOption) (Twitter_FilterClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Twitter_serviceDesc.Streams[1], c.cc, "/twitter.Twitter/Filter", opts...)
	if err != nil {
		return nil, err
	}
	x := &twitterFilterClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Twitter_FilterClient interface {
	Recv() (*Tweet, error)
	grpc.ClientStream
}

type twitterFilterClient struct {
	grpc.ClientStream
}

func (x *twitterFilterClient) Recv() (*Tweet, error) {
	m := new(Tweet)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Twitter service

type TwitterServer interface {
	// We have a method called `GetTimeline` which takes
	// parameter called `Token` and returns
	// the user `Timeline` (list of `Tweets`).
	GetTimeline(context.Context, *Token) (*Timeline, error)
	// We have a method called `UserStream` which takes
	// parameter called `Token`and returns
	// an stream of `Tweets`
	UserStream(*Token, Twitter_UserStreamServer) error
	// We have a method called `Filter` which takes
	// parameter called `Search` and returns
	// an stream of `Tweets`
	Filter(*Search, Twitter_FilterServer) error
}

func RegisterTwitterServer(s *grpc.Server, srv TwitterServer) {
	s.RegisterService(&_Twitter_serviceDesc, srv)
}

func _Twitter_GetTimeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
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
		return srv.(TwitterServer).GetTimeline(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _Twitter_UserStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Token)
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

func _Twitter_Filter_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Search)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TwitterServer).Filter(m, &twitterFilterServer{stream})
}

type Twitter_FilterServer interface {
	Send(*Tweet) error
	grpc.ServerStream
}

type twitterFilterServer struct {
	grpc.ServerStream
}

func (x *twitterFilterServer) Send(m *Tweet) error {
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
		{
			StreamName:    "Filter",
			Handler:       _Twitter_Filter_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("twitter/twitter.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0x5d, 0x4f, 0xea, 0x40,
	0x10, 0xa5, 0xf4, 0x03, 0x98, 0x02, 0xf7, 0xde, 0x4d, 0x6e, 0x42, 0x6e, 0x6e, 0xa2, 0xd6, 0x8f,
	0xf8, 0x22, 0x92, 0xfa, 0x0f, 0x34, 0xd1, 0x27, 0x7d, 0x58, 0xea, 0xa3, 0x21, 0xb5, 0x8c, 0xb1,
	0x11, 0x5a, 0xb3, 0x5d, 0xc4, 0x5f, 0xe3, 0xcf, 0xf2, 0xf7, 0xb8, 0x3b, 0xbb, 0x2d, 0x11, 0x79,
	0x62, 0xe7, 0x9c, 0x33, 0x67, 0xce, 0x0c, 0x85, 0xbf, 0x72, 0x9d, 0x4b, 0x89, 0xe2, 0xdc, 0xfe,
	0x8e, 0x5f, 0x45, 0x29, 0x4b, 0xd6, 0xb1, 0x65, 0x74, 0x0b, 0x7e, 0x52, 0xbe, 0x60, 0xc1, 0x0e,
	0xa0, 0x9f, 0x66, 0x19, 0x56, 0xd5, 0x4c, 0xea, 0x7a, 0xe4, 0xec, 0x3b, 0xa7, 0x3d, 0x1e, 0x1a,
	0xac, 0x91, 0x54, 0x98, 0x09, 0x94, 0x56, 0xd2, 0x36, 0x12, 0x83, 0x91, 0x24, 0x7a, 0x00, 0xef,
	0xbe, 0x42, 0xc1, 0x86, 0xd0, 0xce, 0xe7, 0xe4, 0xe1, 0x72, 0xf5, 0x62, 0x0c, 0xbc, 0x22, 0x5d,
	0xa2, 0x6d, 0xa1, 0x37, 0xdb, 0x83, 0xb0, 0x52, 0x9d, 0x58, 0xcc, 0x88, 0x72, 0x89, 0x02, 0x03,
	0xdd, 0x69, 0xc1, 0x6f, 0x70, 0x57, 0x62, 0x31, 0xf2, 0x88, 0xd0, 0xcf, 0xe8, 0xd3, 0x51, 0x71,
	0xd7, 0x88, 0x72, 0xd7, 0x80, 0x04, 0xdf, 0x65, 0x3d, 0x40, 0xaa, 0xb7, 0xca, 0xeb, 0xad, 0x54,
	0x18, 0x72, 0x0e, 0xe3, 0xc1, 0xb8, 0x3e, 0x81, 0x4e, 0xc8, 0x89, 0x62, 0xff, 0xa1, 0xa7, 0xb2,
	0x6b, 0x47, 0x9c, 0xd3, 0xa0, 0x2e, 0xdf, 0x00, 0xec, 0x10, 0x06, 0xb6, 0x98, 0x65, 0xe5, 0xaa,
	0x90, 0x23, 0x5f, 0x29, 0x7c, 0xde, 0xb7, 0xe0, 0x95, 0xc6, 0xb4, 0xc5, 0x53, 0xfa, 0x56, 0x8a,
	0x5c, 0x5b, 0x04, 0xc6, 0xa2, 0x01, 0xd8, 0x31, 0x0c, 0xeb, 0xc2, 0x7a, 0x74, 0xc8, 0x63, 0x50,
	0xa3, 0x64, 0x12, 0x5d, 0x42, 0x30, 0xc5, 0x54, 0x64, 0xcf, 0x7a, 0x11, 0x1d, 0xde, 0xde, 0xdf,
	0x2c, 0x72, 0x04, 0xfe, 0xe6, 0xe2, 0x61, 0x3c, 0x6c, 0x36, 0xa1, 0xa3, 0x73, 0x43, 0x46, 0x31,
	0x74, 0x93, 0x7c, 0x89, 0x8b, 0xbc, 0x40, 0x76, 0x02, 0x01, 0x45, 0xac, 0x94, 0x8f, 0xfb, 0xbd,
	0x45, 0xc3, 0xdc, 0xb2, 0xf1, 0x87, 0x03, 0x9d, 0xc4, 0x30, 0x2c, 0x86, 0xf0, 0x46, 0xfd, 0x8f,
	0xb5, 0xc5, 0xd6, 0x94, 0x7f, 0x7f, 0x36, 0xb5, 0x95, 0x44, 0x2d, 0x36, 0x01, 0xd0, 0xd7, 0x9c,
	0x4a, 0x81, 0xe9, 0xf2, 0x47, 0xcb, 0xd6, 0xd4, 0xa8, 0x35, 0x71, 0xd8, 0x19, 0x04, 0xd7, 0xf9,
	0x42, 0xcf, 0xfb, 0xd5, 0xb0, 0x66, 0xf5, 0x5d, 0xf2, 0xc7, 0x80, 0xbe, 0xd7, 0x8b, 0xaf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x13, 0xdf, 0xa5, 0x27, 0xc8, 0x02, 0x00, 0x00,
}
