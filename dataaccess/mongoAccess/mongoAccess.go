package mongoAccess

import "gopkg.in/mgo.v2"

var Session, nil = mgo.Dial("mongodb://localhost")

func GetDatabase(s *mgo.Session) *mgo.Database {
	return s.DB("atongraphql");
}
func GetCollection(s *mgo.Session, collectionName string) *mgo.Collection {
	return GetDatabase(s).C(collectionName);
}