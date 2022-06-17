package main

import (
	"context"

	"github.com/trk54ylmz/blog-microservice-example/common"
	"github.com/trk54ylmz/blog-microservice-example/proto/article"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ArticleService is implementation of grpc service definition
type ArticleService struct {
	db         *common.Database
	collection *mongo.Collection
}

// NewArticleService will create article service with active database connection
func NewArticleService() (*ArticleService, error) {
	a := new(ArticleService)

	// Create new database connection
	db, err := common.NewDatabase()

	if err != nil {
		return nil, err
	}

	a.db = db

	// Select database and collection
	a.collection = db.Client().Database("blog").Collection("article")

	return a, nil
}

// Create will create new article
func (a *ArticleService) Create(ctx context.Context, c *article.Article) (*article.ArticleCreateResponse, error) {
	r := new(article.ArticleCreateResponse)

	// Generate new object id, if given id is invalid
	if len(c.Id) != 24 {
		c.Id = primitive.NewObjectID().Hex()
	}

	// Generate object id from raw value
	id, err := primitive.ObjectIDFromHex(c.Id)

	if err != nil {
		return nil, err
	}

	record := Article{
		ID:          id,
		UserID:      c.UserId,
		Title:       c.Title,
		Description: c.Description,
	}

	// Create new article
	_, err = a.collection.InsertOne(ctx, record)

	if err != nil {
		return nil, err
	}

	r.Status = true

	return r, nil
}

// List will return all articles
func (a *ArticleService) List(ctx context.Context, _ *emptypb.Empty) (*article.ArticleListResponse, error) {
	r := new(article.ArticleListResponse)

	// Find all articles
	ac, err := a.collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	// Close cursor after query
	defer ac.Close(ctx)

	var as []*article.Article

	for ac.Next(ctx) {
		var row Article

		// Decode result set to map
		if err := ac.Decode(&row); err != nil {
			return nil, err
		}

		a := new(article.Article)
		a.Id = row.ID.Hex()
		a.Title = row.Title
		a.Description = row.Description
		a.UserId = row.UserID

		as = append(as, a)
	}

	r.Status = true
	r.Articles = as

	return r, nil
}
