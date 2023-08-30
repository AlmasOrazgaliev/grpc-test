package book

import "go.mongodb.org/mongo-driver/bson/primitive"

type Entity struct {
	ObjectID primitive.ObjectID `bson:"_id" json:"_id"`
	Name     *string            `db:"name"`
	Genre    *string            `db:"genre"`
	ISBN     *string            `db:"isbn"`
}
