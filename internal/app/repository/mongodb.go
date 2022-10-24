package repository

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/cavelms/config"
	"github.com/cavelms/internal/model"
	"github.com/cavelms/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func newMongoDBRepository() MongoDB {
	config.LoadConfig()
	dbi := config.NewDBConnection()
	db := &mongodb{dbi}

	// Set email as unique
	err := utils.SetIndex("users", "email", dbi.MongoDB)
	if err != nil {
		log.Panic(err.Error())
	}

	// Setup Admin User
	password := os.Getenv("ADMIN_PASS")
	hashedPass, _ := utils.EncryptPassword(password)
	user := &model.User{
		FullName:     "Bezaleel Onojeta",
		Email:        os.Getenv("ADMIN_EMAIL"),
		PasswordHash: hashedPass,
		Role:         "Administrator",
		IsVerified:   true,
	}
	//Save User To DB
	if err := db.Create(user); err != nil {
		e := err.(mongo.WriteException)
		if c := e.WriteErrors[0].Code; c != 11000 {
			log.Panic(err.Error())
		}
	}

	return db
}

/**
CRUD functions
*/

// Create creates a new user record
func (db *mongodb) Create(m interface{}) error {
	col, v, err := utils.InitCollection(m, db.MongoDB)
	if err != nil {
		return err
	}

	now := time.Now()
	v["createdAt"] = now
	v["updatedAt"] = now
	v["_id"] = primitive.NewObjectID().Hex()

	re, err := col.InsertOne(context.TODO(), v)
	if err != nil {
		return err
	}

	v["id"] = re.InsertedID
	bm, err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bm, &m)
	if err != nil {
		return err
	}

	return nil
}

// FetchByID fetches User by id
func (db *mongodb) FetchByID(m interface{}) error {
	col, v, err := utils.InitCollection(m, db.MongoDB)
	if err != nil {
		return err
	}

	err = col.FindOne(context.TODO(), bson.M{"_id": v["_id"]}).Decode(m)
	if err != nil {
		return err
	}

	return nil
}

// FetchByEmail fetches User by email
func (db *mongodb) FetchByEmail(m interface{}) error {
	col, v, err := utils.InitCollection(m, db.MongoDB)
	if err != nil {
		return err
	}

	err = col.FindOne(context.TODO(), bson.M{"email": v["email"]}).Decode(m)
	if err != nil {
		return err
	}
	return nil
}

// FetchAll fetchs all User
func (db *mongodb) FetchAll(ml interface{}, m interface{}) error {
	col, _, err := utils.InitCollection(m, db.MongoDB)
	if err != nil {
		return err
	}
	cursor, err := col.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}

	if err = cursor.All(context.TODO(), ml); err != nil {
		return err
	}

	return nil
}

func (db *mongodb) FetchByUserID(ml interface{}, m interface{}) error {

	col, v, err := utils.InitCollection(m, db.MongoDB)
	if err != nil {
		return err
	}
	cursor, err := col.Find(context.TODO(), bson.M{"userId": v["userId"].(string)})
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), ml); err != nil {
		return err
	}
	return nil
}

// UpdateOne updates a given user
func (db *mongodb) UpdateOne(m interface{}) error {
	col, v, err := utils.InitCollection(m, db.MongoDB)
	if err != nil {
		return err
	}

	v["updatedAt"] = time.Now()
	bm, err := bson.Marshal(v)
	if err != nil {
		return err
	}

	var val bson.M
	err = bson.Unmarshal(bm, &val)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": v["_id"]}
	update := bson.D{{Key: "$set", Value: val}}
	_, err = col.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

// UpdateManyWhere: updates many logBooks where field and value matches
func (db *mongodb) UpdateManyWhere(m interface{}, field, value string) error {
	col, v, err := utils.InitCollection(m, db.MongoDB)
	if err != nil {
		return err
	}

	t := time.Now()
	v["UpdatedAt"] = &t

	bm, err := bson.Marshal(v)
	if err != nil {
		return err
	}

	var val bson.M
	err = bson.Unmarshal(bm, &val)
	if err != nil {
		return err
	}

	filter := bson.M{field: value}
	update := bson.D{{Key: "$set", Value: val}}
	_, err = col.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes user by id
func (db *mongodb) Delete(m interface{}) error {
	col, v, err := utils.InitCollection(m, db.MongoDB)
	if err != nil {
		return err
	}

	v["deletedAt"] = time.Now()

	_, err = col.DeleteOne(context.TODO(), bson.M{"_id": v["_id"]})
	if err != nil {
		return err
	}

	return nil
}

func (db *mongodb) DeleteMany(m interface{}) error {
	col, _, err := utils.InitCollection(m, db.MongoDB)
	if err != nil {
		return err
	}

	_, err = col.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
