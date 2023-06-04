package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/kirakulakov/konvik-pb-server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = flag.String("addr", "localhost:50051", "grpc server address")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewKonvikClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ping, err := client.Ping(ctx, &pb.PingRequest{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(ping)

}
