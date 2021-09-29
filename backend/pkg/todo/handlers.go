package todo

import (
	"context"

	"github.com/elan8/propanedb-demo/pkg/pb"
	"github.com/elan8/propanedb-go-driver/propane"
	log "github.com/sirupsen/logrus"
)

func (s server) Get(ctx context.Context, in *pb.TodoId) (*pb.TodoItem, error) {

	propaneId := &propane.PropaneId{
		Id: in.Id,
	}
	entity, err := s.client.Get(ctx, propaneId)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	output := &pb.TodoItem{}

	if err := entity.Data.UnmarshalTo(output); err != nil {
		log.Fatalf("Error: %s", err)

	}

	return output, nil
}

func (s server) GetAll(ctx context.Context, in *pb.Empty) (*pb.TodoItems, error) {
	//TBD
	//entity := s.client.Search()

	output := &pb.TodoItems{}
	return output, nil
}

func (s server) Create(ctx context.Context, in *pb.TodoItem) (*pb.TodoItem, error) {
	propaneEntity := &propane.PropaneEntity{}
	id, err := s.client.Put(ctx, propaneEntity)
	in.Id = id.GetId()
	return in, err
}

func (s server) Delete(ctx context.Context, in *pb.TodoId) (*pb.Result, error) {
	propaneId := &propane.PropaneId{
		Id: in.Id,
	}
	_, err := s.client.Delete(ctx, propaneId)
	output := &pb.Result{}
	return output, err
}

func (s server) Search(ctx context.Context, in *pb.TodoSearch) (*pb.TodoItems, error) {
	//TBD
	output := &pb.TodoItems{}
	return output, nil
}
