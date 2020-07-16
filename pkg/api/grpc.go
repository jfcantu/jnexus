package api

import (
	"net"

	"github.com/jfcantu/jnexus/pb"
	"google.golang.org/grpc"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func (s *Server) serveGRPC() {
	logrus.Infof("Starting GRPC listener on: %v", viper.GetString("grpc-server-bind-addr"))

	l, err := net.Listen("tcp", viper.GetString("grpc-server-bind-addr"))
	if err != nil {
		logrus.Fatal(err.Error())
	}

	s.GRPCServer = grpc.NewServer()
	pb.RegisterAPIServer(s.GRPCServer, s)

	s.WaitGroup.Add(1)

	// Blocking call, so start a goroutine that decrements the waitgroup when it finishes
	go func() {
		defer s.WaitGroup.Done()
		if err := s.GRPCServer.Serve(l); err != nil {
			logrus.Fatal(err.Error())
		}
	}()
}
