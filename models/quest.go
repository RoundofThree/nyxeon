package models

import (
	"context"
	"errors"
	"time"

	"github.com/RoundofThree/nyxeon/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Quest struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Categories []string           `bson:"categories,omitempty" json:"categories"`
	Content    string             `bson:"content,omitempty" json:"content"`
	Date       time.Time          `bson:"date,omitempty" json:"date"`
}

type QuestManager struct{}

// fetch all quests from user defined by the session token
func (q *QuestManager) GetByUser(u *User) ([]Quest, error) {
	// ids := u.Quests
	// bulk operation
	var questCollection = db.GetDB().Collection("quests")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ids := u.Quests
	query := bson.M{"_id": bson.M{"$in": ids}}
	cursor, err := questCollection.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	quests := make([]Quest, len(ids))
	if err := cursor.All(ctx, &quests); err != nil {
		return nil, err
	}
	return quests, nil
}

// TODO: Get by category
func (q *QuestManager) GetByCategory(u *User, category string) ([]Quest, error) {
	return nil, nil
}

// Create quest
func (q *QuestManager) Create(u *User, content string, categories []string) error {
	if len(content) <= 0 || len(categories) <= 0 {
		return errors.New("Empty quest not allowed")
	}
	var questCollection = db.GetDB().Collection("quests")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	quest := Quest{
		ID:         primitive.NewObjectID(),
		Date:       time.Now(),
		Content:    content,
		Categories: categories,
	}
	result, err := questCollection.InsertOne(ctx, quest)
	if err != nil {
		return err
	}
	// insert the quest to user collection
	var userCollection = db.GetDB().Collection("users")
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"_id": u.Email}
	update := bson.M{"$push": bson.M{"quests": result.InsertedID.(primitive.ObjectID)}}
	_, err = userCollection.UpdateOne(ctx, filter, update)
	return err
}

// No update operation

// delete quest, not for now
func (q *QuestManager) Delete(u *User, questID string) error {
	return nil
}
