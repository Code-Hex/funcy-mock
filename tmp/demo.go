package main

import (
	"context"

	"github.com/Code-Hex/funcy-mock/tmp/internal/auth"
	"github.com/Code-Hex/funcy-mock/tmp/internal/user"
)

type UserService interface {
	Echo() (*user.EchoResponse, error)
	Get() (*user.Response, error)
}

func main() {}

type AuthService interface {
	Auth(ctx context.Context) (*auth.Token, error)
}
