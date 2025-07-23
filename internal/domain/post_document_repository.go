package domain

type PostDocumentRepository interface {
	GetByPostID(postID PostID) (*PostDocument, error)
	Search(query string) ([]*PostDocument, error)
	Save(postDocument *PostDocument) error
}
