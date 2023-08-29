package repository

import (
	"context"
	"libraryService/internal/repository/mongo"
	"libraryService/pkg/store"
)

type Configuration func(r *Repository) error

// Repository is an implementation of the Repository
type Repository struct {
	mongo  *store.Database
	Book   mongo.BookRepository
	Author mongo.AuthorRepository
	Member mongo.MemberRepository
}

// New takes a variable amount of Configuration functions and returns a new Repository
// Each Configuration will be called in the order they are passed in
func New(configs ...Configuration) (s *Repository, err error) {
	// Create the repository
	s = &Repository{}

	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the repository into the configuration function
		if err = cfg(s); err != nil {
			return
		}
	}

	return
}

// Close closes the repository and prevents new queries from starting.
// Close then waits for all queries that have started processing on the server to finish.
func (r *Repository) Close() {
	if r.mongo != nil {
		err := r.mongo.Client.Disconnect(context.TODO())
		if err != nil {
			return
		}
	}
}

func WithMongoStore(db string) Configuration {
	return func(s *Repository) (err error) {
		// Create the postgres store, if we needed parameters, such as connection strings they could be inputted here
		s.mongo, err = store.NewDatabase()
		if err != nil {
			return
		}

		s.Book = mongo.NewBookRepository(s.mongo.Client.Database(db).Collection("books"))
		s.Author = mongo.NewAuthorRepository(s.mongo.Client.Database(db).Collection("authors"))
		s.Member = mongo.NewMemberRepository(s.mongo.Client.Database(db).Collection("members"))

		return
	}
}
