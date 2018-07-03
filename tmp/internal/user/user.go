package user

import "time"

type EchoResponse struct {
	code    int
	message string
}

type User struct {
	Name      string
	Age       uint
	CreatedAt time.Time
}

type Response struct {
	User    *User
	code    int
	message string
}
