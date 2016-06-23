package main

import (
	"flag"
	"fmt"
	pb "github.com/leandro-lugaresi/gRPC-twitter/twitter"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
)

var (
	index          = flag.Int("index", 0, "RPC port is 6666+index; debug port is 9999+index")
	consumerKey    = flag.Int("consumerKey", os.Getenv("TWITTER_CONSUMER_KEY"), "consumer key for access Twitter apis")
	consumerSecret = flag.Int("consumerSecret", os.Getenv("TWITTER_CONSUMER_SECRET"), "consumer secret for access Twitter apis")
)

type server struct{}

// We have a method called `GetTimeline` which takes
// parameter called `User` and returns
// the user `Timeline` (list of `Tweets`).
func (s *server) GetTimeline(cx context.Context, user *pb.User) (*pb.Timeline, error) {

}

// We have a method called `UserStream` which takes
// parameter called `Search` and returns
// an stream of `Tweets`.
func (s *server) UserStream(search *Search, str pb.Twitter_UserStreamServer) error {

}

func main() {
	flag.Parse()
	go http.ListenAndServe(fmt.Sprintf(":%d", 9999+*index), nil)   //HTTP debugging
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 6666+*index)) //RPC port
	if err != nil {
		log.Fatalf("filed to listen: %v", err)
	}
	g := grpc.NewServer()
	pb.RegisterTwitterServer(g, new(server))
	g.Serve(lis)
}
