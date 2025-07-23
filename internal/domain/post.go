package domain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        PostID
	Title     string
	Slug      PostSlug
	CreatedAt time.Time
	UpdatedAt time.Time
	Revision  int
	Content   string
}

func NewPost(now time.Time) (*Post, error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	id := PostID(uuid)
	title := ""
	slug := PostSlug("")
	createdAt := now
	updatedAt := now
	revision := 0
	content := ""

	return &Post{
		ID:        id,
		Title:     title,
		Slug:      slug,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Revision:  revision,
		Content:   content,
	}, nil
}

func (p *Post) Update(title string, content string, now time.Time) {
	slug := PostSlug(strings.ToLower(strings.ReplaceAll(title, " ", "-")))

	p.Title = title
	p.Slug = slug
	p.Content = content
	p.UpdatedAt = now
	p.Revision++
}

type PostID uuid.UUID

type PostSlug string
