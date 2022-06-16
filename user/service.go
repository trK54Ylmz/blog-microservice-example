package main

import (
	"context"

	"github.com/trk54ylmz/blog-microservice-example/proto/user"
)

// UserService is implementation of grpc service definition
type UserService struct {
}

// SignIn will returns if user exists or not
func (u *UserService) SignIn(ctx context.Context, r *user.SignInRequest) (*user.SignInResponse, error) {
	rs := new(user.SignInResponse)

	if r.Username == "test" && r.Password == "test" {
		rs.Status = true

		rs.Id = 1
	} else {
		rs.Status = false
	}

	return rs, nil
}
