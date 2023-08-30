package member

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	desc "libraryService/proto"
)

type Entity struct {
	ObjectID primitive.ObjectID `bson:"_id" json:"_id"`
	FullName *string            `db:"full_name"`
	Books    []*desc.BookData   `db:"books"`
}
