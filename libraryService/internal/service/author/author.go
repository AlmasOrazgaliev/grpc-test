package author

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"libraryService/internal/model/author"
	"libraryService/pkg/log"
	desc "libraryService/proto"
)

func (s *Service) List(ctx context.Context, req *desc.AuthorData) (res *desc.ListAuthor, err error) {
	logger := log.LoggerFromContext(ctx).Named("ListAuthors")

	data, err := s.authorRepository.List(ctx)
	if err != nil {
		logger.Error("failed to select", zap.Error(err))
		return
	}
	res.Data = author.ParseFromEntities(data)
	return
}

func (s *Service) Add(ctx context.Context, req *desc.AuthorData) (res *desc.AuthorData, err error) {
	logger := log.LoggerFromContext(ctx).Named("AddAuthor")
	data := &author.Entity{
		ObjectID:  primitive.NewObjectID(),
		FullName:  &req.FullName,
		Pseudonym: &req.Pseudonym,
		Specialty: &req.Specialty,
	}
	err = s.authorRepository.Add(ctx, data)
	if err != nil {
		logger.Error("failed to add", zap.Error(err))
		return
	}
	res = author.ParseFromEntity(data)
	return
}

func (s *Service) Get(ctx context.Context, req *desc.AuthorData) (res *desc.AuthorData, err error) {
	logger := log.LoggerFromContext(ctx).Named("GetAuthor")
	id, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		logger.Error("failed to convert string to ObjectID", zap.Error(err))
		return
	}
	data, err := s.authorRepository.Get(ctx, id)
	if err != nil {
		logger.Error("failed to get", zap.Error(err))
		return
	}
	res = author.ParseFromEntity(data)
	return
}

func (s *Service) Update(ctx context.Context, req *desc.AuthorData) (res *desc.AuthorData, err error) {
	logger := log.LoggerFromContext(ctx).Named("UpdateAuthor")
	id, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		logger.Error("failed to convert string to ObjectID", zap.Error(err))
		return
	}
	data := &author.Entity{
		ObjectID:  id,
		FullName:  &req.FullName,
		Pseudonym: &req.Pseudonym,
		Specialty: &req.Specialty,
	}
	err = s.authorRepository.Update(ctx, data)
	if err != nil {
		logger.Error("failed to update", zap.Error(err))
		return
	}
	res = author.ParseFromEntity(data)
	return
}

func (s *Service) Delete(ctx context.Context, req *desc.AuthorData) (res *desc.AuthorData, err error) {
	logger := log.LoggerFromContext(ctx).Named("DeleteAuthor")
	id, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		logger.Error("failed to convert string to ObjectID", zap.Error(err))
		return
	}
	err = s.authorRepository.Delete(ctx, id)
	if err != nil {
		logger.Error("failed to delete", zap.Error(err))
		return
	}
	res = req
	return
}
