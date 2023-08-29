package author

import (
	"context"
	"go.uber.org/zap"
	"libraryService/pkg/log"
	desc "libraryService/proto"
)

func (s *Service) List(ctx context.Context, req *desc.AuthorData) (res *desc.ListAuthor, err error) {
	logger := log.LoggerFromContext(ctx).Named("ListAuthors")

	res, err = s.authorRepository.List(ctx)
	if err != nil {
		logger.Error("failed to select", zap.Error(err))
		return
	}
	return
}

func (s *Service) Add(ctx context.Context, req *desc.AuthorData) (res *desc.AuthorData, err error) {
	logger := log.LoggerFromContext(ctx).Named("AddAuthor")

	res, err = s.authorRepository.Add(ctx, req)
	if err != nil {
		logger.Error("failed to add", zap.Error(err))
		return
	}
	return
}

func (s *Service) Get(ctx context.Context, req *desc.AuthorData) (res *desc.AuthorData, err error) {
	logger := log.LoggerFromContext(ctx).Named("GetAuthor")

	res, err = s.authorRepository.Get(ctx, req.GetId())
	if err != nil {
		logger.Error("failed to get", zap.Error(err))
		return
	}
	return
}

func (s *Service) Update(ctx context.Context, req *desc.AuthorData) (res *desc.AuthorData, err error) {
	logger := log.LoggerFromContext(ctx).Named("UpdateAuthor")

	res, err = s.authorRepository.Update(ctx, req)
	if err != nil {
		logger.Error("failed to update", zap.Error(err))
		return
	}
	return
}

func (s *Service) Delete(ctx context.Context, req *desc.AuthorData) (res *desc.AuthorData, err error) {
	logger := log.LoggerFromContext(ctx).Named("DeleteAuthor")

	res, err = s.authorRepository.Delete(ctx, req)
	if err != nil {
		logger.Error("failed to delete", zap.Error(err))
		return
	}
	return
}
