package author

import desc "libraryService/proto"

func ParseFromEntity(data *Entity) (res *desc.AuthorData) {
	res = &desc.AuthorData{
		Id:        data.ObjectID.Hex(),
		FullName:  *data.FullName,
		Pseudonym: *data.Pseudonym,
		Specialty: *data.Specialty,
	}
	return
}

func ParseFromEntities(data []Entity) (res []*desc.AuthorData) {
	res = make([]*desc.AuthorData, 0)
	for _, object := range data {
		res = append(res, ParseFromEntity(&object))
	}
	return
}
