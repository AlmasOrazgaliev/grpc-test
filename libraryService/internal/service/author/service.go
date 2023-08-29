package author

import (
	"libraryService/internal/repository/mongo"
	desc "libraryService/proto"
)

type Configuration func(s *Service) error

type Service struct {
	desc.UnimplementedAuthorServer
	authorRepository mongo.AuthorRepository
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

func WithAuthorRepository(authorRepository mongo.AuthorRepository) Configuration {
	return func(s *Service) error {
		s.authorRepository = authorRepository
		return nil
	}
}
