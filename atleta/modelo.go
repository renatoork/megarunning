package atleta

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Atleta struct {
	Id          bson.ObjectId `bson:"_id" json:"id"`
	Nome        string        `bson:"nome"`
	Email       string        `bson:"email"`
	Nascimento  time.Time     `bson:"nascimento"`
	Status      string        `bson:"status"`
	Membrodesde time.Time     `bson:"membrodesde"`
	Camiseta    string        `bson:"camiseta"`
}
