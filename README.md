# funcy-mock
[![Go Report Card](https://goreportcard.com/badge/github.com/Code-Hex/funcy-mock)](https://goreportcard.com/report/github.com/Code-Hex/funcy-mock)

funcy-mock generates mock file from interface go file

## Synopsis

You can try following:

- `git clone git@github.com:Code-Hex/funcy-mock.git && cd funcy-mock/tmp`
- `funcygen demo.go`

If you have a some question, You run `funcygen -h` or report to issue.

## Description

**THIS TOOL IS BETA QUALITY.**

`funcygen` just by specifying the Go file described the interface, we will generate a go file containing the mock code.

## Installation

    go get github.com/cmd/funcygen

## Generated file

from `demo.go` to `demo_mock_for_test.go`

```go
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
```

