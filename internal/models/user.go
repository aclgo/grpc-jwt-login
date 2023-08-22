package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Tokens struct {
	Access  string
	Refresh string
}

type User struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Lastname  string    `json:"last_name" db:"last_name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

func (u *User) ClearPass() {
	u.Password = ""
}

func (u *User) HashPass() error {
	bc, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(bc)
	return nil
}

func (u *User) ComparePass(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}

	return nil
}
