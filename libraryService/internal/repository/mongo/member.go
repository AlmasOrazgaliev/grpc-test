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

type MemberRepository struct {
	db *mongo.Collection
}

func NewMemberRepository(db *mongo.Collection) MemberRepository {
	return MemberRepository{
		db: db,
	}
}

func (s *MemberRepository) List(ctx context.Context) (*desc.ListMember, error) {
	var members desc.ListMember
	cursor, err := s.db.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &members.Data); err != nil {
		return nil, err
	}
	return &members, nil
}

func (s *MemberRepository) ListBooks(ctx context.Context, id string) (*desc.ListBook, error) {
	var books desc.ListBook
	cursor, err := s.db.Find(ctx, bson.D{{"_id", id}})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &books.Data); err != nil {
		return nil, err
	}
	return &books, nil
}

func (s *MemberRepository) Add(ctx context.Context, req *desc.MemberData) (*desc.MemberData, error) {
	res, err := s.db.InsertOne(ctx, req)
	if err != nil {
		return nil, err
	}
	fmt.Println(res.InsertedID.(string))
	return req, nil
}

func (s *MemberRepository) Get(ctx context.Context, id string) (*desc.MemberData, error) {
	res := desc.MemberData{}
	err := s.db.FindOne(ctx, bson.D{{"_id", id}}).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *MemberRepository) Update(ctx context.Context, req *desc.MemberData) (*desc.MemberData, error) {
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

func (s *MemberRepository) Delete(ctx context.Context, req *desc.MemberData) (*desc.MemberData, error) {
	_, err := s.db.DeleteOne(ctx, bson.D{{"_id", req.Id}})
	if err != nil {
		return nil, err
	}
	return req, nil
}
