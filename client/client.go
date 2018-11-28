package main

import (
	"context"
	"flag"
	"fmt"

	pb "github.com/cored/eco/protos"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	server := flag.String("s", "localhost:8080", "echo server address")
	argText := flag.String("t", "hello to eco", "text to send to the server")
	flag.Parse()

	conn, err := grpc.Dial(*server, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("could not connect to the server, make sure is running - %v", err)
	}
	defer conn.Close()

	client := pb.NewEcoClient(conn)

	text := &pb.Text{Text: *argText}

	res, err := client.Echo(context.Background(), text)
	if err != nil {
		logrus.Fatalf("echo endpoint returned the following error - %v", err)
	}
	fmt.Printf("%v: ", res)
}
