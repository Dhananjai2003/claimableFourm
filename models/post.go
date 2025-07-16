package models

import (
	"time"
)

type Post struct {
	ID int
	Title string
	Content string
	UserID int
	IsAnonymous bool
	CreatedAt time.Time
	Media []Media
}