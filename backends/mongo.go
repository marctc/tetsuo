package backends

import (
	"time"

	mgo "gopkg.in/mgo.v2"
)

type MongoBackend struct {
	Collection *mgo.Collection
}

type MongoPayload struct {
	Metric    string
	Name      string
	TimeStamp time.Time
	Value     interface{}
}

func NewMongoBackend(dbName, collectionName string) (mb *MongoBackend) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	c := session.DB(dbName).C(collectionName)
	return &MongoBackend{Collection: c}
}

func (mb *MongoBackend) Report(metric, name string, value interface{}) (err error) {
	item := MongoPayload{
		Metric:    metric,
		Name:      name,
		TimeStamp: time.Now(),
		Value:     value,
	}
	err = mb.Collection.Insert(&item)
	return err
}
