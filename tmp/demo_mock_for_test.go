package main

import (
	"context"
	"github.com/Code-Hex/funcy-mock/tmp/internal/auth"
	"github.com/Code-Hex/funcy-mock/tmp/internal/user"
)

type UserServiceMock struct {
	auth func(context.Context) (*auth.Token, error)
	get  func() (*user.Response, error)
}

func NewUserServiceMock() *UserServiceMock {
	return &UserServiceMock{
		auth: func(context.Context) (*auth.Token, error) { return nil, nil },
		get:  func() (*user.Response, error) { return nil, nil },
	}
}

func (u *UserServiceMock) Auth(ctx context.Context) (*auth.Token, error) {
	return u.auth()
}

func (u *UserServiceMock) SetAuth(f func(context.Context) (*auth.Token, error)) {
	if f == nil {
		panic("You should specify the mock function")
	}
	u.auth = f
}

func (u *UserServiceMock) Get() (*user.Response, error) {
	return u.get()
}

func (u *UserServiceMock) SetGet(f func() (*user.Response, error)) {
	if f == nil {
		panic("You should specify the mock function")
	}
	u.get = f
}

type AuthServiceMock struct {
	auth func(context.Context) (*auth.Token, error)
}

func NewAuthServiceMock() *AuthServiceMock {
	return &AuthServiceMock{
		auth: func(context.Context) (*auth.Token, error) { return nil, nil },
	}
}

func (a *AuthServiceMock) Auth(ctx context.Context) (*auth.Token, error) {
	return a.auth()
}

func (a *AuthServiceMock) SetAuth(f func(context.Context) (*auth.Token, error)) {
	if f == nil {
		panic("You should specify the mock function")
	}
	a.auth = f
}
