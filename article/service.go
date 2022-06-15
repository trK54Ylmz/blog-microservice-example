package main

import (
	"context"

	"github.com/trk54ylmz/blog-ms/proto/article"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ArticleService struct {
}

// Create will create new article
func (a *ArticleService) Create(ctx context.Context, c *article.Article) (*article.ArticleCreateResponse, error) {
	r := new(article.ArticleCreateResponse)

	return r, nil
}

// List will return all articles
func (a *ArticleService) List(ctx context.Context, _ *emptypb.Empty) (*article.ArticleListResponse, error) {
	r := new(article.ArticleListResponse)

	return r, nil
}
