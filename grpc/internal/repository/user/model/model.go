package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID              int64
	Info            UserInfo
	Password        string
	PasswordConfirm string
	CreatedAt       time.Time
	UpdatedAt       sql.NullTime
}

type UserInfo struct {
	Name  string
	Email string
	Role  int32
}

type CreateUser struct {
	Info            UserInfo
	Password        string
	PasswordConfirm string
}

type GetUser struct {
	ID        int64
	Info      UserInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
