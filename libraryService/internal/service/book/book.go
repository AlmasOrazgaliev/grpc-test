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
		Name:  &req.Name,
		Genre: &req.Genre,
		ISBN:  &req.Isbn,
	}
	id, err := s.bookRepository.Add(ctx, data)
	data.ObjectID = id.(primitive.ObjectID)
	if err != nil {
		logger.Error("failed to add", zap.Error(err))
		return
	}
	res = book.ParseFromEntity(data)
	return
}

func (s *Service) Get(ctx context.Context, req *desc.BookData) (res *desc.BookData, err error) {
	logger := log.LoggerFromContext(ctx).Named("GetBook")

	data, err := s.bookRepository.Get(ctx, req.GetId())
	if err != nil {
		logger.Error("failed to get", zap.Error(err))
		return
	}
	res = book.ParseFromEntity(data)
	return
}

func (s *Service) Update(ctx context.Context, req *desc.BookData) (res *desc.BookData, err error) {
	logger := log.LoggerFromContext(ctx).Named("UpdateBook")
	data := &book.Entity{
		Name:  &req.Name,
		Genre: &req.Genre,
		ISBN:  &req.Isbn,
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

	err = s.bookRepository.Delete(ctx, req.GetId())
	if err != nil {
		logger.Error("failed to delete", zap.Error(err))
		return
	}
	res = req
	return
}
