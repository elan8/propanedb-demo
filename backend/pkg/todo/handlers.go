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

	// searchQuery := propane.PropaneSearch{}
	// searchQuery.DatabaseName = databaseName
	// searchQuery.EntityType = "todo.TodoItem"
	// searchQuery.Query = "*"
	entities, err := s.client.Search(ctx, "todo.TodoItem", "*")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	for _, entity := range entities {
		item := entity.(*pb.TodoItem)
		// if err := entity.Data.UnmarshalTo(item); err != nil {
		// 	log.Fatalf("Error: %s", err)
		// }
		output.Items = append(output.Items, item)
	}

	return output, nil
}

func (s server) Create(ctx context.Context, in *pb.TodoItem) (*pb.TodoItem, error) {
	// propanePut := &propane.PropanePut{}
	// propanePut.DatabaseName = databaseName

	// propaneEntity := &propane.PropaneEntity{}

	// any, err := anypb.New(in)
	// if err != nil {
	// 	log.Fatalf("Error: %s", err)
	// }
	// propaneEntity.Data = any
	// propanePut.Entity = propaneEntity

	id, err := s.client.Put(ctx, in)
	in.Id = id
	return in, err
}

func (s server) Delete(ctx context.Context, in *pb.TodoId) (*pb.Result, error) {
	// propaneId := &propane.PropaneId{
	// 	Id:           in.Id,
	// 	DatabaseName: databaseName,
	// }
	err := s.client.Delete(ctx, in.Id)
	output := &pb.Result{}
	return output, err
}
