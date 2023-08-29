package book

import (
	"context"
	"go.uber.org/zap"
	"libraryService/pkg/log"
	desc "libraryService/proto"
)

func (s *Service) List(ctx context.Context, req *desc.BookData) (res *desc.ListBook, err error) {
	logger := log.LoggerFromContext(ctx).Named("ListBooks")

	res, err = s.bookRepository.List(ctx)
	if err != nil {
		logger.Error("failed to select", zap.Error(err))
		return
	}
	return
}

func (s *Service) Add(ctx context.Context, req *desc.BookData) (res *desc.BookData, err error) {
	logger := log.LoggerFromContext(ctx).Named("AddBook")

	res, err = s.bookRepository.Add(ctx, req)
	if err != nil {
		logger.Error("failed to add", zap.Error(err))
		return
	}
	return
}

func (s *Service) Get(ctx context.Context, req *desc.BookData) (res *desc.BookData, err error) {
	logger := log.LoggerFromContext(ctx).Named("GetBook")

	res, err = s.bookRepository.Get(ctx, req.GetId())
	if err != nil {
		logger.Error("failed to get", zap.Error(err))
		return
	}
	return
}

func (s *Service) Update(ctx context.Context, req *desc.BookData) (res *desc.BookData, err error) {
	logger := log.LoggerFromContext(ctx).Named("UpdateBook")

	res, err = s.bookRepository.Update(ctx, req)
	if err != nil {
		logger.Error("failed to update", zap.Error(err))
		return
	}
	return
}

func (s *Service) Delete(ctx context.Context, req *desc.BookData) (res *desc.BookData, err error) {
	logger := log.LoggerFromContext(ctx).Named("DeleteBook")

	res, err = s.bookRepository.Delete(ctx, req)
	if err != nil {
		logger.Error("failed to delete", zap.Error(err))
		return
	}
	return
}
