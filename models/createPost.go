package models

type CreatePostRequest struct {

	Title string
	Content string
	UserID int
	IsAnonymous bool
	MediaURLs []string
	FileTypes []string

}