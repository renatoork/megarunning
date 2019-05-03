package prova

import (
	"gopkg.in/mgo.v2/bson"
	"projetos/megarunning/atleta"
	"time"
)

type Prova struct {
	Id            bson.ObjectId  `bson:"_id" json:"id"`
	Nome          string         `bson:"nome"`
	Url           string         `bson:"url"`
	Local         string         `bson:"local"`
	Data          time.Time      `bson:"data"`
	Valor         float32        `bson:"valor"`
	Participantes []Participacao `bson:"participantes"`
}

type Participacao struct {
	ValorPatrocinio float32       `bson:"valorpatrocinio"`
	Atleta          atleta.Atleta `bson:"atleta"`
}
