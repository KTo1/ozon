package model

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var ErrNotFound = errors.New("note not found")

type Note struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
