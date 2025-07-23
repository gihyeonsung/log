package domain

type PostUpdated struct {
	PostID PostID
}

func NewPostUpdated(postID PostID) *PostUpdated {
	return &PostUpdated{PostID: postID}
}
