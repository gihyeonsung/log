package application

import (
	"errors"

	"github.com/gihyeonsung/log/internal/domain"
)

type PostDocumentSync struct {
	postDocumentRepository domain.PostDocumentRepository
	postRepository         domain.PostRepository
}

func (c *PostDocumentSync) Exec(postDocumentID domain.PostDocumentID) error {
	postDocument, err := c.postDocumentRepository.Get(postDocumentID)
	if err != nil {
		return err
	}

	post, err := c.postRepository.Get(postDocument.PostID)
	if err != nil {
		return err
	}

	if post == nil {
		return errors.New("post nil")
	}

	if postDocument == nil {
		postDocument, err = domain.NewPostDocument(post)
		if err != nil {
			return err
		}
	} else {
		postDocument.Update(post)
	}

	err = c.postDocumentRepository.Save(postDocument)
	if err != nil {
		return err
	}

	return nil
}
