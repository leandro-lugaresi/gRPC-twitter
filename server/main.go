package main

import (
	"flag"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	pb "github.com/leandro-lugaresi/gRPC-twitter/twitter"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"net/url"
	"os"
)

var (
	index          = flag.Int("index", 0, "RPC port is 6666+index; debug port is 9999+index")
	consumerKey    = flag.String("consumerKey", os.Getenv("TWITTER_CONSUMER_KEY"), "consumer key for access Twitter apis")
	consumerSecret = flag.String("consumerSecret", os.Getenv("TWITTER_CONSUMER_SECRET"), "consumer secret for access Twitter apis")
)

type server struct{}

// We have a method called `GetTimeline` which takes
// parameter called `User` and returns
// the user `Timeline` (list of `Tweets`).
func (s *server) GetTimeline(cx context.Context, token *pb.Token) (*pb.Timeline, error) {
	api := anaconda.NewTwitterApi(token.AccessToken, token.SecretToken)
	v := url.Values{}
	timeline, err := api.GetHomeTimeline(v)
	if err != nil {
		return nil, err
	}
	t := make([]pb.Tweet, len(timeline))
	for i, tweet := range timeline {
		t[i] = pb.Tweet{
			Id:            tweet.Id,
			Text:          tweet.Text,
			User:          pb.User{Id: tweet.User.Id, Name: tweet.User.Name, ScreenName: tweet.User.ScreenName},
			Retweeted:     tweet.Retweeted,
			RetweetCount:  tweet.RetweetCount,
			Favorited:     tweet.Favorited,
			FavoriteCount: tweet.FavoriteCount,
		}
	}
	return &pb.Timeline{t}, nil
}

// We have a method called `UserStream` which takes
// parameter called `Search` and returns
// an stream of `Tweets`.
func (s *server) UserStream(search *Search, str pb.Twitter_UserStreamServer) error {

}

func main() {
	flag.Parse()

	anaconda.SetConsumerKey(*consumerKey)
	anaconda.SetConsumerSecret(*consumerSecret)

	go http.ListenAndServe(fmt.Sprintf(":%d", 9999+*index), nil)   //HTTP debugging
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 6666+*index)) //RPC port
	if err != nil {
		log.Fatalf("filed to listen: %v", err)
	}
	g := grpc.NewServer()

	pb.RegisterTwitterServer(g, new(server))
	g.Serve(lis)
}
