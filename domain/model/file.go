package model

import "gopkg.in/mgo.v2/bson"

type File  struct {
	ID bson.ObjectId `bson:"_id" json:"_id"`
	Disabled bool `bson:"disabled" json:"disabled"`
	CreatedDate string `bson:"createdDate" json:"createdDate"`
	ModifiedDate string `bson:"modifiedDate" json:"modifiedDate"`
	ReferenceClass string `bson:"referenceClass" json:"referenceClass"`
	ReferenceId string `bson:"referenceId" json:"referenceId"`
	FilePath string `bson:"filePath" json:"filePath"`
	FileName string `bson:"fileName" json:"fileName"`
	FileContent string `bson:"fileContent" json:"fileContent"`
	Creator string `bson:"creator" json:"creator"`
}