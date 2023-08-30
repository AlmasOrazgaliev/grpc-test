package member

import (
	"libraryService/internal/model/book"
	desc "libraryService/proto"
)

func ParseFromEntity(data Entity) (res desc.MemberData) {
	res = desc.MemberData{
		Id:       data.ObjectID.Hex(),
		FullName: *data.FullName,
		Books:    book.ParseFromEntities(data.Books),
	}
	return
}

func ParseFromEntities(data []Entity) (res []desc.MemberData) {
	res = make([]desc.MemberData, 0)
	for _, object := range data {
		res = append(res, ParseFromEntity(object))
	}
	return
}
