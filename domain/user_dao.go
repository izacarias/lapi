package domain

import (
	"context"

	"github.com/izacarias/lapi/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

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

func GetUsersByAccessPoint(accessPoint string) ([]User, error) {
	filter := bson.M{"access_point": accessPoint}
	var users []UserMongo
	cursor, err := userCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	userList := make([]User, len(users))
	for i, user := range users {
		userList[i] = User{
			Address:     user.Address,
			AccessPoint: user.AccessPoint,
			ZoneId:      "not defined",
		}
	}
	return userList, nil
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

func UpdateUser(user *User) error {
	filter := bson.M{"address": user.Address}
	update := bson.M{"$set": bson.M{"access_point": user.AccessPoint}}
	_, err := userCollection.UpdateOne(context.TODO(), filter, update)
	return err
}

func InsertUser(user *User) error {
	filter := bson.M{"address": user.Address}
	count, err := userCollection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return err
	}
	if count > 0 {
		return nil // User already exists
	}

	userMongo := UserMongo{
		Address:     user.Address,
		AccessPoint: user.AccessPoint,
	}
	_, err = userCollection.InsertOne(context.TODO(), userMongo)
	return err
}
