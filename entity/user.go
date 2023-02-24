package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserID int64
type Role string

type User struct {
	ID       UserID    `json:"id"  db:"id"`
	Name     string    `json:"name" db:"name"`
	Password string    `json:"password" db:"password"`
	Role     Role      `json:"role" db:"role"`
	Created  time.Time `josn:"created" db:"created"`
	Modified time.Time `josn:"modified" db:"modified"`
}

func (u *User) ComparePassword(pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pw))
}
