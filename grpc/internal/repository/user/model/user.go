package model

import (
	"database/sql"
	"time"
)

type User struct {
	Id              int64        `db:"id"`
	Info            UserInfo     `db:""`
	Password        string       `db:"password"`
	PasswordConfirm string       `db:"password_confirm"`
	CreatedAt       time.Time    `db:"created_at"`
	UpdatedAt       sql.NullTime `db:"updated_at"`
}

type UserInfo struct {
	Name  string `db:"name"`
	Email string `db:"email"`
	Role  int32  `db:"role"`
}

type CreateUser struct {
	Info            UserInfo `db:""`
	Password        string   `db:"password"`
	PasswordConfirm string   `db:"password_confirm"`
}

type GetUser struct {
	Id        int64        `db:"id"`
	Info      UserInfo     `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
