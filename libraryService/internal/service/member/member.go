package member

import (
	"context"
	"go.uber.org/zap"
	"libraryService/pkg/log"
	desc "libraryService/proto"
)

func (s *Service) List(ctx context.Context, req *desc.MemberData) (res *desc.ListMember, err error) {
	logger := log.LoggerFromContext(ctx).Named("ListMembers")

	res, err = s.memberRepository.List(ctx)
	if err != nil {
		logger.Error("failed to select", zap.Error(err))
		return
	}
	return
}

func (s *Service) Add(ctx context.Context, req *desc.MemberData) (res *desc.MemberData, err error) {
	logger := log.LoggerFromContext(ctx).Named("AddMember")

	res, err = s.memberRepository.Add(ctx, req)
	if err != nil {
		logger.Error("failed to add", zap.Error(err))
		return
	}
	return
}

func (s *Service) Get(ctx context.Context, req *desc.MemberData) (res *desc.MemberData, err error) {
	logger := log.LoggerFromContext(ctx).Named("GetMember")

	res, err = s.memberRepository.Get(ctx, req.GetId())
	if err != nil {
		logger.Error("failed to get", zap.Error(err))
		return
	}
	return
}

func (s *Service) Update(ctx context.Context, req *desc.MemberData) (res *desc.MemberData, err error) {
	logger := log.LoggerFromContext(ctx).Named("UpdateMember")

	res, err = s.memberRepository.Update(ctx, req)
	if err != nil {
		logger.Error("failed to update", zap.Error(err))
		return
	}
	return
}

func (s *Service) Delete(ctx context.Context, req *desc.MemberData) (res *desc.MemberData, err error) {
	logger := log.LoggerFromContext(ctx).Named("DeleteMember")

	res, err = s.memberRepository.Update(ctx, req)
	if err != nil {
		logger.Error("failed to delete", zap.Error(err))
		return
	}
	return
}
