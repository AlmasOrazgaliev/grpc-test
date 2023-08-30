package mongo

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"libraryService/internal/model/book"
)

type BookRepository struct {
	db *mongo.Collection
}

func NewBookRepository(db *mongo.Collection) BookRepository {
	return BookRepository{
		db: db,
	}
}

func (s *BookRepository) List(ctx context.Context) (dest []book.Entity, err error) {
	cursor, err := s.db.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &dest); err != nil {
		return nil, err
	}
	return
}

func (s *BookRepository) Add(ctx context.Context, req *book.Entity) (interface{}, error) {
	res, err := s.db.InsertOne(ctx, req)
	if err != nil {
		return nil, err
	}
	req.ObjectID = res.InsertedID.(primitive.ObjectID)
	return res.InsertedID, nil
}

func (s *BookRepository) Get(ctx context.Context, id string) (*book.Entity, error) {
	res := book.Entity{}
	err := s.db.FindOne(ctx, bson.D{{"_id", id}}).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *BookRepository) Update(ctx context.Context, req *book.Entity) error {
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

func (s *BookRepository) Delete(ctx context.Context, id string) error {
	_, err := s.db.DeleteOne(ctx, bson.D{{"_id", id}})
	if err != nil {
		return err
	}
	return nil
}
