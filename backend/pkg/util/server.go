package util

import (
	"net"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func GetServerOptions(config Config) []grpc.ServerOption {
	opts := []grpc.ServerOption{}

	return opts
}

func WaitForGrpc() {
	time.Sleep(5 * time.Second)
}

func StartServer(s *grpc.Server) {
	lis, err := net.Listen("tcp", ":"+viper.GetString("port"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
