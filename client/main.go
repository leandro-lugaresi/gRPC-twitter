package main

import (
	"flag"
	"fmt"
	pb "github.com/leandro-lugaresi/gRPC-twitter/twitter"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

var (
	server = flag.String("server", "localhost:36060", "server address")
	mode   = flag.String("mode", "search", `one of "search" or "timeline"`)
	query  = flag.String("query", "test", "query string")
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
		timeline(client, *query)
	default:
		log.Fatalf("unknown mode: %q", *mode)
	}
}

// timeline runs a Timeline RPC and prints the result stream.
func timeline(client pb.TwitterClient, user string) {
	ctx, cancel := context.WithTimeout(context.Background(), 80*time.Millisecond)
	defer cancel()
	req := &pb.User{Uuser}
	res, err := client.GetTimeline(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

// search issues a search for query and prints the result.
func search(client pb.GoogleClient, query string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	req := &pb.Request{Query: query}
	stream, err := client.Watch(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("and now your watch is ended")
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)
	}
}
