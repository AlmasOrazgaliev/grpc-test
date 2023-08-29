package mongo

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	desc "libraryService/proto"
)

type BookRepository struct {
	db *mongo.Collection
}

func NewBookRepository(db *mongo.Collection) BookRepository {
	return BookRepository{
		db: db,
	}
}

func (s *BookRepository) List(ctx context.Context) (*desc.ListBook, error) {
	var books desc.ListBook
	cursor, err := s.db.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &books.Data); err != nil {
		return nil, err
	}
	return &books, nil
}

func (s *BookRepository) Add(ctx context.Context, req *desc.BookData) (*desc.BookData, error) {
	res, err := s.db.InsertOne(ctx, req)
	if err != nil {
		return nil, err
	}
	primitive.ObjectID{}.Hex()
	fmt.Println(res.InsertedID)
	return req, nil
}

func (s *BookRepository) Get(ctx context.Context, id string) (*desc.BookData, error) {
	res := desc.BookData{}
	err := s.db.FindOne(ctx, bson.D{{"_id", id}}).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *BookRepository) Update(ctx context.Context, req *desc.BookData) (*desc.BookData, error) {
	options.Update().SetUpsert(true)
	b, _ := json.Marshal(&req)
	filter := bson.D{{"_id", req.Id}}
	update := bson.D{{"$set", b}}
	_, err := s.db.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (s *BookRepository) Delete(ctx context.Context, req *desc.BookData) (*desc.BookData, error) {
	_, err := s.db.DeleteOne(ctx, bson.D{{"_id", req.Id}})
	if err != nil {
		return nil, err
	}
	return req, nil
}
