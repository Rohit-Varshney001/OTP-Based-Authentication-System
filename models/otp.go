package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OTP struct {
	MobileNumber string    `bson:"mobileNumber"`
	OTP          string    `bson:"otp"`
	ExpiresAt    time.Time `bson:"expiresAt"`
}

func CreateOTP(collection *mongo.Collection, otp OTP) error {
	_, err := collection.InsertOne(context.TODO(), otp)
	return err
}

func ValidateOTP(collection *mongo.Collection, mobile, otp string) (bool, error) {
	var result OTP
	err := collection.FindOne(context.TODO(), bson.M{
		"mobileNumber": mobile,
		"otp":          otp,
		"expiresAt":    bson.M{"$gte": time.Now()},
	}).Decode(&result)
	return err == nil, err
}
