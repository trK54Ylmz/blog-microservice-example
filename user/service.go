package main

import (
	"context"

	"github.com/trk54ylmz/blog-ms/proto/user"
)

type UserService struct {
}

// SignIn will returns if user exists or not
func (u *UserService) SignIn(ctx context.Context, r *user.SignInRequest) (*user.SignInResponse, error) {
	rs := new(user.SignInResponse)

	if r.Username == "test" && r.Password == "test" {
		rs.Status = true
	} else {
		rs.Status = false
	}

	return rs, nil
}
