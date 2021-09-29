package main

import (
	"net/http"

	"github.com/elan8/propanedb-demo/pkg/pb"
	"github.com/elan8/propanedb-demo/pkg/util"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	cors "github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

func run() error {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	config, err := util.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config")
	}
	opts, err := util.GetDialOptions(config)
	viper.SetDefault("todolistEndpoint", "todo-service:50000")

	err = viper.BindEnv("todolistEndpoint", "TODOLISTHOST")
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	todolistEndpoint := viper.GetString("todolistEndpoint")

	err = pb.RegisterTodoListHandlerFromEndpoint(ctx, mux, todolistEndpoint, opts)
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowCredentials: true,
		Debug:            false,
	})

	return http.ListenAndServe(":8080", c.Handler(mux))
}

func main() {
	log.Println("Start GRPC gateway")
	if err := run(); err != nil {
		log.Println(err)
	}
}
