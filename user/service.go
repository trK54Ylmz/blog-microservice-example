package main

import (
	"context"
	"errors"

	"github.com/trk54ylmz/blog-microservice-example/common"
	"github.com/trk54ylmz/blog-microservice-example/proto/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	// ErrUserExist is result of user already exists on the collection
	ErrUserExist = errors.New("the user exists")

	// ErrUserNoExist is result of a user check on the collection, but does not exist
	ErrUserNoExist = errors.New("the user does not exist")
)

// UserService is implementation of grpc service definition
type UserService struct {
	db         *common.Database
	collection *mongo.Collection
}

// NewUserService will create user service with active database connection
func NewUserService() (*UserService, error) {
	u := new(UserService)

	// Create new database connection
	db, err := common.NewDatabase()

	if err != nil {
		return nil, err
	}

	u.db = db

	// Select database and collection
	u.collection = db.Client().Database("blog").Collection("user")

	return u, nil
}

// SignUp will create new user account, if username does not exist
func (u *UserService) SignUp(ctx context.Context, c *user.SignUpRequest) (*user.SignUpResponse, error) {
	r := new(user.SignUpResponse)

	q := bson.M{
		"username": c.Username,
	}

	// Get user by given user name
	rs := u.collection.FindOne(ctx, q)

	if rs.Err() != mongo.ErrNoDocuments {
		return nil, ErrUserExist
	}

	// Generate object id
	id := primitive.NewObjectID()

	record := User{
		ID:       id,
		Username: c.Username,
		Password: c.Password,
	}

	// Create new article
	_, err := u.collection.InsertOne(ctx, record)

	if err != nil {
		return nil, err
	}

	r.Status = true
	r.Id = id.Hex()

	return r, nil
}

// SignIn will returns user ID, if user exists or not
func (u *UserService) SignIn(ctx context.Context, c *user.SignInRequest) (*user.SignInResponse, error) {
	r := new(user.SignInResponse)

	q := bson.M{
		"username": c.Username,
	}

	// Get user by given user name
	rs := u.collection.FindOne(ctx, q)

	if rs.Err() == mongo.ErrNoDocuments {
		return nil, ErrUserNoExist
	}

	us := new(User)

	// Convert raw bson value to user instance
	if err := rs.Decode(us); err != nil {
		return nil, err
	}

	r.Status = true
	r.Id = us.ID.Hex()

	return r, nil
}
