package models

import (
	"context"
	"time"

	"github.com/RoundofThree/nyxeon/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Email  string               `bson:"_id,omitempty"`
	Quests []primitive.ObjectID `bson:"quests,omitempty"`
}

type UserManager struct{}

// Get the user given the userID.
func (u *UserManager) GetByUserID(userID string) (*User, error) {
	var userCollection = db.GetDB().Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result := userCollection.FindOne(ctx, bson.M{"_id": userID})
	if result.Err() != nil {
		return nil, result.Err()
	}
	var user User
	if err := result.Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

// Create a user given the userID.
func (u *UserManager) CreateUser(userID string) (*User, error) {
	var userCollection = db.GetDB().Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user := User{
		Email:  userID,
		Quests: []primitive.ObjectID{},
	}
	_, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
