package todo

import (
	"context"

	"github.com/elan8/propanedb-demo/pkg/pb"
	"github.com/elan8/propanedb-go-driver/propane"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/anypb"
)

func (s server) Get(ctx context.Context, in *pb.TodoId) (*pb.TodoItem, error) {

	propaneId := &propane.PropaneId{
		Id:           in.Id,
		DatabaseName: databaseName,
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
	output := &pb.TodoItems{}

	searchQuery := propane.PropaneSearch{}
	searchQuery.DatabaseName = databaseName
	searchQuery.EntityType = "todo.TodoItem"
	searchQuery.Query = "*"
	entities, err := s.client.Search(ctx, &searchQuery)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	for _, entity := range entities.Entities {
		item := &pb.TodoItem{}
		if err := entity.Data.UnmarshalTo(item); err != nil {
			log.Fatalf("Error: %s", err)
		}
		output.Items = append(output.Items, item)
	}

	return output, nil
}

func (s server) Create(ctx context.Context, in *pb.TodoItem) (*pb.TodoItem, error) {
	propanePut := &propane.PropanePut{}
	propanePut.DatabaseName = databaseName

	propaneEntity := &propane.PropaneEntity{}

	any, err := anypb.New(in)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	propaneEntity.Data = any
	propanePut.Entity = propaneEntity

	id, err := s.client.Put(ctx, propanePut)
	in.Id = id.GetId()
	return in, err
}

func (s server) Delete(ctx context.Context, in *pb.TodoId) (*pb.Result, error) {
	propaneId := &propane.PropaneId{
		Id:           in.Id,
		DatabaseName: databaseName,
	}
	_, err := s.client.Delete(ctx, propaneId)
	output := &pb.Result{}
	return output, err
}

func (s server) Search(ctx context.Context, in *pb.TodoSearch) (*pb.TodoItems, error) {

	propaneSearch := &propane.PropaneSearch{}
	propaneSearch.DatabaseName = databaseName
	propaneSearch.EntityType = "todo.TodoItem"
	propaneSearch.Query = "*"

	entities, err := s.client.Search(ctx, propaneSearch)
	if err != nil {
		log.Fatalf("Error: %s", err)
		return nil, err
	}
	output := &pb.TodoItems{}
	items := output.Items

	for _, entity := range entities.Entities {
		todoItem := &pb.TodoItem{}

		any := entity.Data

		//m := new(propane.TodoItem)
		if err := any.UnmarshalTo(todoItem); err != nil {
			log.Fatalf("Error: %s", err)
			//t.Errorf("Cannot unmarshal to TodoItem")
		}

		items = append(items, todoItem)
	}

	return output, nil
}
