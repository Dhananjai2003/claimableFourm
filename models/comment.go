package models


type CreateCommentRequest struct {

	PostID int
	ParentID *int
	Content string
	IsAnonymous bool

}