package member

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"libraryService/internal/model/member"
	"libraryService/pkg/log"
	desc "libraryService/proto"
)

func (s *Service) List(ctx context.Context, req *desc.MemberData) (*desc.ListMember, error) {
	logger := log.LoggerFromContext(ctx).Named("ListMembers")
	res := &desc.ListMember{}
	data, err := s.memberRepository.List(ctx)
	if err != nil {
		logger.Error("failed to select", zap.Error(err))
		return nil, err
	}
	res.Data = member.ParseFromEntities(data)
	return res, nil
}

func (s *Service) Add(ctx context.Context, req *desc.MemberData) (res *desc.MemberData, err error) {
	logger := log.LoggerFromContext(ctx).Named("AddMember")
	data := &member.Entity{
		ObjectID: primitive.NewObjectID(),
		FullName: &req.FullName,
		Books:    req.Books,
	}
	err = s.memberRepository.Add(ctx, data)
	if err != nil {
		logger.Error("failed to add", zap.Error(err))
		return
	}
	res = member.ParseFromEntity(data)
	return
}

func (s *Service) Get(ctx context.Context, req *desc.MemberData) (res *desc.MemberData, err error) {
	logger := log.LoggerFromContext(ctx).Named("GetMember")
	id, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		logger.Error("failed to convert string to ObjectID", zap.Error(err))
		return
	}
	data, err := s.memberRepository.Get(ctx, id)
	if err != nil {
		logger.Error("failed to get", zap.Error(err))
		return
	}
	res = member.ParseFromEntity(data)
	return
}

func (s *Service) Update(ctx context.Context, req *desc.MemberData) (res *desc.MemberData, err error) {
	logger := log.LoggerFromContext(ctx).Named("UpdateMember")
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		logger.Error("failed to convert string to ObjectID", zap.Error(err))
		return
	}
	data := &member.Entity{
		ObjectID: id,
		FullName: &req.FullName,
		Books:    req.Books,
	}
	err = s.memberRepository.Update(ctx, data)
	if err != nil {
		logger.Error("failed to update", zap.Error(err))
		return
	}
	res = req
	return
}

func (s *Service) Delete(ctx context.Context, req *desc.MemberData) (res *desc.MemberData, err error) {
	logger := log.LoggerFromContext(ctx).Named("DeleteMember")
	id, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		logger.Error("failed to convert string to ObjectID", zap.Error(err))
		return
	}
	err = s.memberRepository.Delete(ctx, id)
	if err != nil {
		logger.Error("failed to delete", zap.Error(err))
		return
	}
	res = req
	return
}
