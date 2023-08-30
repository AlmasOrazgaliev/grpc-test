package book

import desc "libraryService/proto"

func ParseFromEntity(data Entity) (res *desc.BookData) {
	res = &desc.BookData{
		Id:    data.ObjectID.Hex(),
		Name:  *data.Name,
		Genre: *data.Genre,
		Isbn:  *data.ISBN,
	}
	return
}

func ParseFromEntities(data []Entity) []*desc.BookData {
	res := make([]*desc.BookData, 0)
	for _, object := range data {
		res = append(res, ParseFromEntity(object))
	}
	return res
}
