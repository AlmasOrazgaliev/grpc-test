package author

import "go.mongodb.org/mongo-driver/bson/primitive"

type Entity struct {
	ObjectID  primitive.ObjectID `bson:"_id" json:"_id"`
	FullName  *string            `db:"full_name"`
	Pseudonym *string            `db:"pseudonym"`
	Specialty *string            `db:"specialty"`
}
