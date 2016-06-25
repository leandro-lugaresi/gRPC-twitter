package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "github.com/leandro-lugaresi/gRPC-twitter/twitter"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	server      = flag.String("server", "localhost:6666", "server address")
	mode        = flag.String("mode", "search", `one of "search" or "timeline"`)
	query       = flag.String("query", "test", "query string")
	accessToken = flag.String("accessToken", os.Getenv("TWITTER_ACCESS_TOKEN"), "access token")
	secretToken = flag.String("secretToken", os.Getenv("TWITTER_SECRET_TOKEN"), "secret token")
)

func main() {
	flag.Parse()

	// Connect to the server.
	conn, err := grpc.Dial(*server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewTwitterClient(conn)

	// Run the RPC.
	switch *mode {
	case "search":
		search(client, *query)
	case "timeline":
		timeline(client)
	default:
		log.Fatalf("unknown mode: %q", *mode)
	}
}

// timeline runs a Timeline RPC and prints the result stream.
func timeline(client pb.TwitterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	res, err := client.GetTimeline(ctx, getToken())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Last Tweets:")
	for _, t := range res.GetTweets() {
		fmt.Printf("%v : %s \n", t.User.ScreenName, t.Text)
	}
}

// search issues a search for query and prints the result.
func search(client pb.TwitterClient, query string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	req := &pb.Search{Token: getToken(), Text: query}
	stream, err := client.Filter(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Search ended :O")
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v : %s \n", res.User.ScreenName, res.Text)
	}
}

func getToken() *pb.Token {
	return &pb.Token{
		AccessToken: *accessToken,
		SecretToken: *secretToken,
	}
}
