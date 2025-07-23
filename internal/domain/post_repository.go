package domain

type PostRepository interface {
	Get(id PostID) (*Post, error)
	Find() ([]*Post, error)
	Save(post *Post) error
	Delete(id PostID) error
}
