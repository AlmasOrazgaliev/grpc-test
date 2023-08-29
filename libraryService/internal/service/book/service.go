package book

import (
	"libraryService/internal/repository/mongo"
	desc "libraryService/proto"
)

type Configuration func(s *Service) error

type Service struct {
	desc.UnimplementedBookServer
	bookRepository mongo.BookRepository
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

func WithBookRepository(bookRepository mongo.BookRepository) Configuration {
	return func(s *Service) error {
		s.bookRepository = bookRepository
		return nil
	}
}
