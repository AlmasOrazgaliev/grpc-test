package mongo

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	desc "libraryService/proto"
)

type AuthorRepository struct {
	db *mongo.Collection
}

func NewAuthorRepository(db *mongo.Collection) AuthorRepository {
	return AuthorRepository{
		db: db,
	}
}

func (s *AuthorRepository) List(ctx context.Context) (*desc.ListAuthor, error) {
	var books desc.ListAuthor
	cursor, err := s.db.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &books.Data); err != nil {
		return nil, err
	}
	return &books, nil
}

func (s *AuthorRepository) Add(ctx context.Context, req *desc.AuthorData) (*desc.AuthorData, error) {
	res, err := s.db.InsertOne(ctx, req)
	if err != nil {
		return nil, err
	}
	fmt.Println(res.InsertedID.(string))
	return req, nil
}

func (s *AuthorRepository) Get(ctx context.Context, id string) (*desc.AuthorData, error) {
	res := desc.AuthorData{}
	err := s.db.FindOne(ctx, bson.D{{"_id", id}}).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *AuthorRepository) Update(ctx context.Context, req *desc.AuthorData) (*desc.AuthorData, error) {
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

func (s *AuthorRepository) Delete(ctx context.Context, req *desc.AuthorData) (*desc.AuthorData, error) {
	_, err := s.db.DeleteOne(ctx, bson.D{{"_id", req.Id}})
	if err != nil {
		return nil, err
	}
	return req, nil
}
