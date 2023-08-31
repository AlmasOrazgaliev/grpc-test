package book

import (
	desc "api-gateway/proto"
	"errors"
	"net/http"
)

type Request struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	ISBN  string `json:"isbn"`
}

func (s *Request) Bind(r *http.Request) error {
	if s.Name == "" {
		return errors.New("name: cannot be blank")
	}

	if s.Genre == "" {
		return errors.New("genre: cannot be blank")
	}

	if s.ISBN == "" {
		return errors.New("isbn: cannot be blank")
	}

	return nil
}

type Response struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	ISBN  string `json:"isbn"`
}

func ParseFromEntity(data *desc.BookData) (res Response) {
	res = Response{
		ID:    data.Id,
		Name:  data.Name,
		Genre: data.Genre,
		ISBN:  data.Isbn,
	}
	return
}

func ParseFromEntities(data *desc.ListBook) (res []Response) {
	res = make([]Response, 0)
	for _, object := range data.Data {
		res = append(res, ParseFromEntity(object))
	}
	return
}
