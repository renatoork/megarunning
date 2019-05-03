package conexao

import (
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func GetDb() *mgo.Database {
	if session == nil {
		var err error
		session, err = mgo.Dial("mongodb://mega:megamega@kahana.mongohq.com:10094/MegaRunning")
		if err != nil {
			panic(err) // no, not really
		}
	}
	return session.Clone().DB("MegaRunning")
}

func CloseSession() {
	session.Close()
}
