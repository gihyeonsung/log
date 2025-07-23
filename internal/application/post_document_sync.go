package application

import (
	"github.com/gihyeonsung/log/internal/domain"
)

type PostDocumentSync struct {
	postDocumentRepository domain.PostDocumentRepository
	postRepository         domain.PostRepository
}

func NewPostDocumentSync(postDocumentRepository domain.PostDocumentRepository, postRepository domain.PostRepository) *PostDocumentSync {
	return &PostDocumentSync{postDocumentRepository: postDocumentRepository, postRepository: postRepository}
}

func (c *PostDocumentSync) Exec(postID domain.PostID) error {
	var postDocument *domain.PostDocument
	var err error
	postDocument, err = c.postDocumentRepository.GetByPostID(postID)
	if err != nil {
		return err
	}

	post, err := c.postRepository.Get(postID)
	if err != nil || post == nil {
		return err
	}

	if postDocument == nil {
		postDocument, err = domain.NewPostDocument(post)
	}
	if err != nil {
		return err
	}

	postDocument.Update(post)

	err = c.postDocumentRepository.Save(postDocument)
	if err != nil {
		return err
	}

	return nil
}
