package application

import (
	"errors"
	"time"

	"github.com/gihyeonsung/log/internal/domain"
)

type PostUpdate struct {
	authnService   AuthnService
	postRepository domain.PostRepository
}

func (c *PostUpdate) Exec(key string, postID domain.PostID, title string, content string) error {
	ok, err := c.authnService.Login(key)
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("authn")
	}

	post, err := c.postRepository.Get(postID)
	if err != nil {
		return err
	}

	if post == nil {
		return errors.New("nil")
	}

	post.Update(title, content, time.Now())
	return c.postRepository.Save(post)
}
