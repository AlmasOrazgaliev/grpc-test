package book

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"libraryService/internal/model/book"
	"libraryService/pkg/log"
	desc "libraryService/proto"
)

func (s *Service) List(ctx context.Context, req *desc.BookData) (res *desc.ListBook, err error) {
	logger := log.LoggerFromContext(ctx).Named("ListBooks")

	data, err := s.bookRepository.List(ctx)
	if err != nil {
		logger.Error("failed to select", zap.Error(err))
		return
	}
	res.Data = book.ParseFromEntities(data)
	return
}

func (s *Service) Add(ctx context.Context, req *desc.BookData) (res *desc.BookData, err error) {
	logger := log.LoggerFromContext(ctx).Named("AddBook")
	data := &book.Entity{
		ObjectID: primitive.NewObjectID(),
		Name:     &req.Name,
		Genre:    &req.Genre,
		ISBN:     &req.Isbn,
	}
	err = s.bookRepository.Add(ctx, data)
	if err != nil {
		logger.Error("failed to add", zap.Error(err))
		return
	}
	res = book.ParseFromEntity(data)
	return
}

func (s *Service) Get(ctx context.Context, req *desc.BookData) (res *desc.BookData, err error) {
	logger := log.LoggerFromContext(ctx).Named("GetBook")
	id, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		logger.Error("failed to convert string to ObjectID", zap.Error(err))
		return
	}
	data, err := s.bookRepository.Get(ctx, id)
	if err != nil {
		logger.Error("failed to get", zap.Error(err))
		return
	}
	res = book.ParseFromEntity(data)
	return
}

func (s *Service) Update(ctx context.Context, req *desc.BookData) (res *desc.BookData, err error) {
	logger := log.LoggerFromContext(ctx).Named("UpdateBook")
	id, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		logger.Error("failed to convert string to ObjectID", zap.Error(err))
		return
	}
	data := &book.Entity{
		ObjectID: id,
		Name:     &req.Name,
		Genre:    &req.Genre,
		ISBN:     &req.Isbn,
	}
	err = s.bookRepository.Update(ctx, data)
	if err != nil {
		logger.Error("failed to update", zap.Error(err))
		return
	}
	res = book.ParseFromEntity(data)
	return
}

func (s *Service) Delete(ctx context.Context, req *desc.BookData) (res *desc.BookData, err error) {
	logger := log.LoggerFromContext(ctx).Named("DeleteBook")
	id, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		logger.Error("failed to convert string to ObjectID", zap.Error(err))
		return
	}
	err = s.bookRepository.Delete(ctx, id)
	if err != nil {
		logger.Error("failed to delete", zap.Error(err))
		return
	}
	res = req
	return
}
