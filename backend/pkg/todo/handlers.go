package todo

import (
	"context"

	"github.com/elan8/propanedb-demo/pkg/pb"
	log "github.com/sirupsen/logrus"
)

func (s server) Get(ctx context.Context, in *pb.TodoId) (*pb.TodoItem, error) {
	entity, err := s.client.Get(ctx, in.Id)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	output := entity.(*pb.TodoItem)
	return output, nil
}

func (s server) GetAll(ctx context.Context, in *pb.Empty) (*pb.TodoItems, error) {
	output := &pb.TodoItems{}
	entities, err := s.client.Search(ctx, "todo.TodoItem", "*")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	for _, entity := range entities {
		item := entity.(*pb.TodoItem)
		output.Items = append(output.Items, item)
	}

	return output, nil
}

func (s server) Create(ctx context.Context, in *pb.TodoItem) (*pb.TodoItem, error) {
	id, err := s.client.Put(ctx, in)
	in.Id = id
	return in, err
}

func (s server) Update(ctx context.Context, in *pb.TodoItem) (*pb.TodoItem, error) {
	id, err := s.client.Put(ctx, in)
	in.Id = id
	return in, err
}

func (s server) Delete(ctx context.Context, in *pb.TodoId) (*pb.Result, error) {
	err := s.client.Delete(ctx, in.Id)
	output := &pb.Result{}
	return output, err
}
