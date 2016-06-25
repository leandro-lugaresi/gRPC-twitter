package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	pb "github.com/leandro-lugaresi/gRPC-twitter/twitter"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
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
	t := make([]*pb.Tweet, len(timeline))
	for i, tweet := range timeline {
		newTweet, err := convertTweet(tweet)
		if err != nil {
			return nil, err
		}
		t[i] = newTweet
	}
	return &pb.Timeline{Tweets: t}, nil
}

// We have a method called `UserStream` which takes
// parameter called `Search` and returns
// an stream of `Tweets`.
func (s *server) UserStream(token *pb.Token, str pb.Twitter_UserStreamServer) error {
	api := anaconda.NewTwitterApi(token.AccessToken, token.SecretToken)
	uStr := api.UserStream(nil)
	ctx := str.Context()
	defer uStr.Stop()
	for {
		select {
		case t := <-uStr.C:
			tweet, err := convertTweet(t)
			if err != nil {
				return err
			}
			err = str.Send(tweet)
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (s *server) Filter(search *pb.Search, str pb.Twitter_FilterServer) error {
	api := anaconda.NewTwitterApi(search.Token.AccessToken, search.Token.SecretToken)
	v := url.Values{}
	v.Set("track", search.Text)
	pStr := api.PublicStreamFilter(v)
	ctx := str.Context()
	defer pStr.Stop()
	for {
		select {
		case t := <-pStr.C:
			tweet, err := convertTweet(t)
			if err != nil {
				return err
			}
			err = str.Send(tweet)
			if err != nil {
				return err
			}

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func convertTweet(tweet interface{}) (*pb.Tweet, error) {
	switch v := tweet.(type) {
	case anaconda.Tweet:
		return &pb.Tweet{
			Id:            v.Id,
			Text:          v.Text,
			User:          &pb.User{Id: v.User.Id, Name: v.User.Name, ScreenName: v.User.ScreenName},
			Retweeted:     v.Retweeted,
			RetweetCount:  int32(v.RetweetCount),
			Favorited:     v.Favorited,
			FavoriteCount: int32(v.FavoriteCount),
		}, nil
	case anaconda.EventTweet:
		t := v.TargetObject
		return &pb.Tweet{
			Id:            t.Id,
			Text:          fmt.Sprintf("%S by %s: %s", v.Event.Event, v.Source.ScreenName, t.Text),
			User:          &pb.User{Id: t.User.Id, Name: t.User.Name, ScreenName: t.User.ScreenName},
			Retweeted:     t.Retweeted,
			RetweetCount:  int32(t.RetweetCount),
			Favorited:     t.Favorited,
			FavoriteCount: int32(t.FavoriteCount),
		}, nil
	}

	return nil, errors.New("Tweet type unknown")
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
