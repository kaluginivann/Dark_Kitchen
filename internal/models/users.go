package models

import (
	"context"
	"time"
)

type User struct {
	Id        int       `db:"id"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UsersRepository interface {
	Create(ctx context.Context, user *User) (*int, error)
	GetByID(ctx context.Context, id int) (*User, error)
	//GetByUsername(ctx context.Context, username string) (*User, error) // TODO: You can do it easily
	//Update(ctx context.Context, user *User) error	// TODO: You can do it easily
	//Delete(ctx context.Context, id int) error	// TODO: You can do it easily
}
