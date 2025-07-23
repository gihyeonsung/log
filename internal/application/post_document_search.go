package application

import "github.com/gihyeonsung/log/internal/domain"

type PostDocumentSearch struct {
	postDocumentRepository domain.PostDocumentRepository
}

func (c *PostDocumentSearch) Exec(query string) ([]*domain.PostDocument, error) {
	ps, err := c.postDocumentRepository.Search(query)
	if err != nil {
		return nil, err
	}
	return ps, nil
}
