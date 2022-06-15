package main

import (
	"context"
	"time"

	"github.com/trk54ylmz/blog-ms/proto/article"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ArticleService struct {
	Cancel context.CancelFunc

	client     *mongo.Client
	collection *mongo.Collection
}

// NewArticleService will create article service with active database connection
func NewArticleService() (*ArticleService, error) {
	a := new(ArticleService)

	opts := options.Client().ApplyURI("mongodb://localhost:27017")

	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	// Create database client
	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		defer cancel()

		return nil, err
	}

	// Select database and collection
	a.collection = client.Database("blog").Collection("article")

	a.client = client
	a.Cancel = cancel

	return a, nil
}

// Create will create new article
func (a *ArticleService) Create(ctx context.Context, c *article.Article) (*article.ArticleCreateResponse, error) {
	r := new(article.ArticleCreateResponse)

	record := bson.M{
		"id":          c.Id,
		"user_id":     c.UserId,
		"title":       c.Title,
		"description": c.Description,
	}

	// Create new article
	_, err := a.collection.InsertOne(ctx, record)

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
		var row bson.M

		// Decode result set to map
		if err := ac.Decode(&row); err != nil {
			return nil, err
		}

		a := new(article.Article)
		a.Id = row["id"].(string)
		a.Title = row["id"].(string)
		a.Description = row["id"].(string)
		a.UserId = row["user_id"].(int32)

		as = append(as, a)
	}

	r.Status = true
	r.Articles = as

	return r, nil
}