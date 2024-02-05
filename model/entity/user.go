package entity

import (
	"time"
)

type Users struct {
	ID        int64     `json:"id"`
	Phone     string    `json:"phone"`
	Password  string    `json:"-"`
	FullName  string    `json:"full_name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (u *Users) TableName() string {
	return "users"
}
