package mongoAccess

import "gopkg.in/mgo.v2"

var (
	Session *mgo.Session
)

func GetDatabase(s *mgo.Session) *mgo.Database {
	return s.DB("atongraphql");
}
func GetCollection(s *mgo.Session, collectionName string) *mgo.Collection {
	return GetDatabase(s).C(collectionName);
}

func init() {
	Session, _ = mgo.Dial("mongodb://localhost")
}