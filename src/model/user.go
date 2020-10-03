package model

import (
	"time"
)

// User model
type User struct {
	ID        uint      `json:"id"`
	Mail      string    `json:"mail"`
	Password  string    `json:"password"`
	Authkey   string    `json:"authkey"`
	Point     uint      `json:"point"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
