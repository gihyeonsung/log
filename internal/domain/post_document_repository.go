package domain

type PostDocumentRepository interface {
	Get(id PostDocumentID) (*PostDocument, error)
	Search(query string) ([]*PostDocument, error)
	Save(postDocument *PostDocument) error
}
