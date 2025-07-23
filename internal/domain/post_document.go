package domain

import (
	"time"

	"github.com/google/uuid"
)

type PostDocument struct {
	ID        PostDocumentID
	PostID    PostID
	Title     string
	Slug      PostSlug
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PostDocumentID uuid.UUID

func NewPostDocument(post *Post) (*PostDocument, error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	id := PostDocumentID(uuid)
	postID := post.ID
	title := post.Title
	slug := post.Slug
	content := post.Content
	createdAt := post.CreatedAt
	updatedAt := post.UpdatedAt

	return &PostDocument{
		ID:        id,
		PostID:    postID,
		Title:     title,
		Slug:      slug,
		Content:   content,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (p *PostDocument) Update(post *Post) {
	p.Title = post.Title
	p.Slug = post.Slug
	p.Content = post.Content
	p.UpdatedAt = post.UpdatedAt
}
