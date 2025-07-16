package models

type CreatePostRequest type {

	Title string
	Content string
	UserID int
	IsAnonymous bool
	MediaURLs []string
	FileTypes []string

}