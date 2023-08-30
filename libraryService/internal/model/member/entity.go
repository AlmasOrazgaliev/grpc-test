package member

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"libraryService/internal/model/book"
)

type Entity struct {
	ObjectID primitive.ObjectID `bson:"_id" json:"_id"`
	FullName *string            `db:"full_name"`
	Books    []book.Entity      `db:"books"`
}
