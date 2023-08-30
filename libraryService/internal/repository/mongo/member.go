package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"libraryService/internal/model/member"
)

type MemberRepository struct {
	db *mongo.Collection
}

func NewMemberRepository(db *mongo.Collection) MemberRepository {
	return MemberRepository{
		db: db,
	}
}

func (s *MemberRepository) List(ctx context.Context) ([]member.Entity, error) {
	var members []member.Entity
	cursor, err := s.db.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &members); err != nil {
		return nil, err
	}
	return members, nil
}

//func (s *MemberRepository) ListBooks(ctx context.Context, id string) (*member.Entity, error) {
//	var books member.Entity
//	cursor, err := s.db.Find(ctx, bson.D{{"_id", id}})
//	if err != nil {
//		return nil, err
//	}
//	if err = cursor.All(context.TODO(), &books.Data); err != nil {
//		return nil, err
//	}
//	return &books, nil
//}

func (s *MemberRepository) Add(ctx context.Context, req *member.Entity) error {
	_, err := s.db.InsertOne(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (s *MemberRepository) Get(ctx context.Context, id primitive.ObjectID) (*member.Entity, error) {
	res := member.Entity{}
	err := s.db.FindOne(ctx, bson.D{{"_id", id}}).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *MemberRepository) Update(ctx context.Context, req *member.Entity) error {
	options.Update().SetUpsert(true)
	pByte, err := bson.Marshal(req)
	if err != nil {
		return err
	}
	var update bson.M
	err = bson.Unmarshal(pByte, &update)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", req.ObjectID}}
	_, err = s.db.UpdateOne(ctx, filter, bson.D{{"$set", update}})
	if err != nil {
		return err
	}
	return nil
}

func (s *MemberRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := s.db.DeleteOne(ctx, bson.D{{"_id", id}})
	if err != nil {
		return err
	}
	return nil
}
