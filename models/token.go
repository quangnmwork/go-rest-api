package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	ID        uuid.UUID `json:"id"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}
