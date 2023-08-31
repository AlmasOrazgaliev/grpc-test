package member

import (
	desc "api-gateway/proto"
	"errors"
	"net/http"
)

type Request struct {
	ID       string `json:"id"`
	FullName string `json:"fullName"`
}

func (s *Request) Bind(r *http.Request) error {
	if s.FullName == "" {
		return errors.New("fullName: cannot be blank")
	}

	return nil
}

type Response struct {
	ID       string `json:"id"`
	FullName string `json:"fullName"`
}

func ParseFromEntity(data *desc.MemberData) (res Response) {
	res = Response{
		ID:       data.Id,
		FullName: data.FullName,
	}
	return
}

func ParseFromEntities(data *desc.ListMember) (res []Response) {
	res = make([]Response, 0)
	for _, object := range data.Data {
		res = append(res, ParseFromEntity(object))
	}
	return
}
