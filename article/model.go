package main

import "go.mongodb.org/mongo-driver/bson/primitive"

// Article is for article collection
type Article struct {
	ID          primitive.ObjectID `bson:"_id"`
	UserId      int32              `bson:"user_id"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
}
