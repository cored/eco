package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "github.com/cored/eco/protos"
	"google.golang.org/grpc"
)

func main() {
	server := flag.String("b", "localhost:8080", "echo server address")
	flag.Parse()

	grpc.WithInsecure()
	conn, err := grpc.Dial(*server, grpc.WithInsecure())
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	client := pb.NewEcoClient(conn)

	text := &pb.Text{Text: "Here we go"}

	for {
		res, err := client.Echo(context.Background(), text)
		if err != nil {
			log.Panicln(err)
		}
		fmt.Printf("%v: ", res)
	}
}
