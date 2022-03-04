package controllers

import (
	"github.com/lihemen/gin_api/models"
)

type Post = *models.Post
type User = models.User

// Create Post
func CreatePost(post *Post) (*Post, error) {
	return post, nil
}

// Get Post
func GetPost(post *Post) (*Post, error) {
	return post, nil
}

// Get all Posts 
func GetAllPosts(post *Post) ([]*Post, error) {
	return nil, nil
}


// Get all Posts for an Author
func GetAllPostsForAuthor(post *Post, author *User) ([]*Post, error) {
	return nil, nil
}


// Delete Post
func DeletePost(post *Post) error {
	return nil
}

// Update Post
func UpdatePost(post *Post) (*Post, error) {
	return post, nil
}

// Like Post
func LikePost(post *Post, user *User) ([]*Post, error) {
	return nil, nil
}