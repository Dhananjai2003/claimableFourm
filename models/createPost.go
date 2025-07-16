package models

type CreatePostRequest struct {
	Title       string   
	Content     string   
	IsAnonymous bool     
	MediaURLs   []string 
	FileTypes   []string 
}