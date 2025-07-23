package application

import "github.com/gihyeonsung/log/internal/domain"

type PostFind struct {
	postRepository domain.PostRepository
}

func NewPostFind(postRepository domain.PostRepository) *PostFind {
	return &PostFind{postRepository: postRepository}
}

func (c *PostFind) Exec() ([]*domain.Post, error) {
	return c.postRepository.Find()
}
