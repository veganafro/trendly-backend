package main

import (
	"log"
	"time"
	"context"
	"google.golang.org/grpc"

	"github.com/veganafro/trendly-backend/nyt"
)

func main() {
	go nyt.Listen()
	opts := grpc.WithInsecure()
	conn, err := grpc.Dial("localhost:8080", opts)
	if err != nil {
		log.Fatal("main: Error dailing grpc server", err)
	}
	defer conn.Close()
	client := nyt.NewNytServiceClient(conn)
	request := nyt.NytRequest{ Token: "test", Period: 1 }
	context, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	mostShared, err := client.GetMostShared(
		context,
		&request,
	)
	if err != nil {
		log.Fatal("main: Error getting most shared", err)
	}
	log.Println("main: NytResponse is", mostShared)
}

