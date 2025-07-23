package infrastructure

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/gihyeonsung/log/internal/domain"
)

type EsPostDocumentRepository struct {
	es *elasticsearch.Client
}

var _ domain.PostDocumentRepository = (*EsPostDocumentRepository)(nil)

func NewEsPostDocumentRepository(es *elasticsearch.Client) *EsPostDocumentRepository {
	return &EsPostDocumentRepository{es: es}
}

const esPostDocumentIndex = "post_documents"

func (e *EsPostDocumentRepository) GetByPostID(postID domain.PostID) (*domain.PostDocument, error) {
	res, err := e.es.Get(esPostDocumentIndex, postID.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.IsError() {
		if res.StatusCode == 404 {
			return nil, nil
		}
		return nil, fmt.Errorf("elasticsearch get error: %s", res.String())
	}
	var doc struct {
		Source domain.PostDocument `json:"_source"`
	}
	if err := json.NewDecoder(res.Body).Decode(&doc); err != nil {
		return nil, err
	}
	return &doc.Source, nil
}

func (e *EsPostDocumentRepository) Save(postDocument *domain.PostDocument) error {
	body, err := json.Marshal(postDocument)
	if err != nil {
		return err
	}
	res, err := e.es.Index(
		esPostDocumentIndex,
		bytes.NewReader(body),
		e.es.Index.WithDocumentID(postDocument.ID.String()),
		// upsert
		e.es.Index.WithRefresh("true"),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("elasticsearch index error: %s", res.String())
	}
	return nil
}

func (e *EsPostDocumentRepository) Search(query string) ([]*domain.PostDocument, error) {
	q := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  query,
				"fields": []string{"title^2", "content", "slug"},
			},
		},
	}
	body, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}
	res, err := e.es.Search(
		e.es.Search.WithContext(context.Background()),
		e.es.Search.WithIndex(esPostDocumentIndex),
		e.es.Search.WithBody(bytes.NewReader(body)),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.IsError() {
		return nil, fmt.Errorf("elasticsearch search error: %s", res.String())
	}
	var resp struct {
		Hits struct {
			Hits []struct {
				Source domain.PostDocument `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, err
	}
	var docs []*domain.PostDocument
	for _, hit := range resp.Hits.Hits {
		d := hit.Source
		docs = append(docs, &d)
	}
	return docs, nil
}
