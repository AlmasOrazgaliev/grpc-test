package service

import (
	"libraryService/internal/service/author"
	"libraryService/internal/service/book"
	"libraryService/internal/service/member"
)

type Configuration func(s *Service) error

type Service struct {
	BookService   book.Service
	AuthorService author.Service
	MemberService member.Service
}

func NewService(bookService book.Service, authorService author.Service, memberService member.Service) *Service {
	return &Service{
		BookService:   bookService,
		AuthorService: authorService,
		MemberService: memberService,
	}
}
