package main

import "go.mongodb.org/mongo-driver/bson/primitive"

// User is for user collection
type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}
