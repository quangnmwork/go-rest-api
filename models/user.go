package models

import (
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

type User struct {
	gorm.Model
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name" validate:"required,min=3,max=32"`
	LastName  string    `json:"last_name" validate:"required,min=3,max=32"`
	Username  string    `json:"username" validate:"required,min=3,max=32"`
	Password  string    `json:"password" validate:"required,min=8"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Tokens    []Token   `json:"tokens"`
}
