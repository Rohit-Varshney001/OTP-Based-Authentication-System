package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	MobileNumber      string    `bson:"mobileNumber"`
	Name              string    `bson:"name"`
	DeviceFingerprint string    `bson:"deviceFingerprint"`
	CreatedAt         time.Time `bson:"createdAt"`
}

func CreateUser(collection *mongo.Collection, user User) error {
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}

func FindUserByMobile(collection *mongo.Collection, mobile string) (User, error) {
	var user User
	err := collection.FindOne(context.TODO(), bson.M{"mobileNumber": mobile}).Decode(&user)
	return user, err
}
