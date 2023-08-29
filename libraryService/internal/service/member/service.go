package member

import (
	"libraryService/internal/repository/mongo"
	desc "libraryService/proto"
)

type Configuration func(s *Service) error

type Service struct {
	desc.UnimplementedMemberServer
	memberRepository mongo.MemberRepository
}

func New(configs ...Configuration) (s *Service, err error) {
	s = &Service{}

	for _, cfg := range configs {
		if err = cfg(s); err != nil {
			return
		}
	}
	return
}

func WithMemberRepository(memberRepository mongo.MemberRepository) Configuration {
	return func(s *Service) error {
		s.memberRepository = memberRepository
		return nil
	}
}
