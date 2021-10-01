package todo

import (
	"context"
	"io/ioutil"

	"github.com/elan8/propanedb-demo/pkg/pb"
	"github.com/elan8/propanedb-demo/pkg/util"
	"github.com/elan8/propanedb-go-driver/propane"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

type server struct {
	pb.UnimplementedTodoListServer
	config util.Config
	client *propane.Client
}

func Init() {
	log.Print("Start todo list service ")

	config, err := util.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config")
	}

	ctx := context.Background()
	serverAddress := "db:50051"
	databaseName := "todo"

	b, err := ioutil.ReadFile("/app/messages.bin")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	fds := &descriptorpb.FileDescriptorSet{}
	err = proto.Unmarshal(b, fds)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	util.WaitForGrpc()

	client, err := propane.Connect(ctx, serverAddress, databaseName, fds)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	d := server{config: config, client: client}

	opts := util.GetServerOptions(config)
	s := grpc.NewServer(opts...)

	pb.RegisterTodoListServer(s, &d)
	util.StartServer(s)
}
