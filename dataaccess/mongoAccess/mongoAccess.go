package mongoAccess

import "gopkg.in/mgo.v2"

var session, nil = mgo.Dial("mongodb://localhost")

func GetSession() *mgo.Session {
	return session
}