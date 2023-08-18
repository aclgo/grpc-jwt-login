package models

import "time"

type User struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Lastname  string    `json:"last_name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TokenJwtUserPayload struct {
	UserID   string        `json:"user_id"`
	ExpireAt time.Duration `json:"expire_at"`
	Role     string        `json:"role"`
}