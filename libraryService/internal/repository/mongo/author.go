package mongo

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"libraryService/internal/model/author"
)

type AuthorRepository struct {
	db *mongo.Collection
}

func NewAuthorRepository(db *mongo.Collection) AuthorRepository {
	return AuthorRepository{
		db: db,
	}
}

func (s *AuthorRepository) List(ctx context.Context) ([]author.Entity, error) {
	var authors []author.Entity
	cursor, err := s.db.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &authors); err != nil {
		return nil, err
	}
	return authors, nil
}

func (s *AuthorRepository) Add(ctx context.Context, req *author.Entity) (interface{}, error) {
	res, err := s.db.InsertOne(ctx, req)
	if err != nil {
		return nil, err
	}
	req.ObjectID = res.InsertedID.(primitive.ObjectID)
	return res.InsertedID, nil
}

func (s *AuthorRepository) Get(ctx context.Context, id string) (*author.Entity, error) {
	res := author.Entity{}
	err := s.db.FindOne(ctx, bson.D{{"_id", id}}).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *AuthorRepository) Update(ctx context.Context, req *author.Entity) error {
	options.Update().SetUpsert(true)
	b, _ := json.Marshal(&req)
	filter := bson.D{{"_id", req.ObjectID}}
	update := bson.D{{"$set", b}}
	_, err := s.db.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthorRepository) Delete(ctx context.Context, id string) error {
	_, err := s.db.DeleteOne(ctx, bson.D{{"_id", id}})
	if err != nil {
		return err
	}
	return nil
}
