package mongoAccess

import "gopkg.in/mgo.v2"

var session, nil = mgo.Dial("mongodb://localhost")

func GetSession() *mgo.Session {
	return session
}

func GetDatabase() *mgo.Database {
	return GetSession().DB("atongraphql");
}
func GetCollection(collectionName string) *mgo.Collection {
	return GetDatabase().C(collectionName);
}