package server

import (
	"context"
	"fmt"
	"net"

	pb "github.com/cored/eco/protos"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Start the server
func Start(address string) {
	lis, err := net.Listen("tcp", address)
	logrus.Infof("eco is serving on %s", fmt.Sprintf("localhost%s", address))
	if err != nil {
		logrus.Fatalf("error connecting to %s - %v", address, err)
	}

	s := grpc.NewServer()
	pb.RegisterEcoServer(s, Server{})
	err = s.Serve(lis)
	if err != nil {
		logrus.Fatalf("error registering server - %v", err)
	}

}

// Server eco service implementation
type Server struct{}

// Echo endpoint implementation
func (s Server) Echo(ctx context.Context, text *pb.Text) (*pb.Text, error) {
	logrus.Info("Received data: %v", text)
	return text, nil
}
