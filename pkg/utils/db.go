package utils

import (
	"context"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetIndex(doc, field string, db *mongo.Database) error {

	coll := db.Collection(doc)
	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{field: 1}, Options: opt}
	if _, err := coll.Indexes().CreateOne(context.Background(), index); err != nil {
		return err
	}

	return nil
}

func InitCollection(m interface{}, db *mongo.Database) (*mongo.Collection, map[string]interface{}, error) {
	var name string

	t := reflect.TypeOf(m)
	if t.Kind() == reflect.Ptr {
		name = t.Elem().Name()
	} else {
		name = t.Name()
	}
	
	bm, err := bson.Marshal(m)
	if err != nil {
		return nil, nil, err
	}

	v := map[string]interface{}{}
	err = bson.Unmarshal(bm, &v)
	if err != nil {
		return nil, nil, err
	}

	colName := strings.ToLower(name + "s")
	return db.Collection(colName), v, nil
}
