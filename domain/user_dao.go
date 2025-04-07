package domain

import (
	"context"

	"github.com/izacarias/lapi/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "locations")

type UserMongo struct {
	Address     string `bson:"address"`
	AccessPoint string `bson:"access_point"`
}

func GetUserByAddress(address string) (*User, error) {
	filter := bson.M{"address": address}
	var user UserMongo
	err := userCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &User{
		Address:     user.Address,
		AccessPoint: user.AccessPoint,
		ZoneId:      "not defined",
	}, nil
}

func GetAllUsers() ([]User, error) {
	cursor, err := userCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	var users []UserMongo
	if err := cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	usersList := make([]User, len(users))
	for i, user := range users {
		usersList[i] = User{
			Address:     user.Address,
			AccessPoint: user.AccessPoint,
			ZoneId:      "not defined",
		}
	}
	return usersList, nil
}
