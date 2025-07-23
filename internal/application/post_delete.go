package application

import (
	"errors"

	"github.com/gihyeonsung/log/internal/domain"
)

type PostDelete struct {
	authnService   AuthnService
	postRepository domain.PostRepository
}

func NewPostDelete(authnService AuthnService, postRepository domain.PostRepository) *PostDelete {
	return &PostDelete{authnService: authnService, postRepository: postRepository}
}

func (c *PostDelete) Exec(key string, postID domain.PostID) error {
	ok, err := c.authnService.Login(key)
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("authn")
	}

	return c.postRepository.Delete(postID)
}
