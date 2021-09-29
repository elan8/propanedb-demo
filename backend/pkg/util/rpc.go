package util

import (
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func GetDialOptions(config Config) ([]grpc.DialOption, error) {

	opts := []grpc.DialOption{grpc.WithInsecure()}

	return opts, nil

}

func Dial(address string, config Config) (*grpc.ClientConn, error) {

	var err error
	var conn *grpc.ClientConn

	conn, err = grpc.Dial(address,
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor()),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor()),
	)
	if err != nil {
		log.Printf("Error: address: %s", address)
		return nil, err
	}

	return conn, nil
}
