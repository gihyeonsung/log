package domain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        PostID
	Title     string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Revision  int
	Content   string
}

type PostID uuid.UUID

type PostSlug string

func NewPostSlug(title string) (PostSlug, error) {
	slug := strings.ToLower(strings.ReplaceAll(title, " ", "-"))
	return PostSlug(slug), nil
}

func NewPostID() (PostID, error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return PostID{}, err
	}
	return PostID(uuid), nil
}
