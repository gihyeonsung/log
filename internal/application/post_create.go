package application

import (
	"errors"
	"time"

	"github.com/gihyeonsung/log/internal/domain"
)

type PostCreate struct {
	authnService   AuthnService
	postRepository domain.PostRepository
}

func NewPostCreate(authnService AuthnService, postRepository domain.PostRepository) *PostCreate {
	return &PostCreate{authnService: authnService, postRepository: postRepository}
}

func (c *PostCreate) Exec(key string) error {
	ok, err := c.authnService.Login(key)
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("authn")
	}

	post, err := domain.NewPost(time.Now())
	if err != nil {
		return err
	}

	return c.postRepository.Save(post)
}
