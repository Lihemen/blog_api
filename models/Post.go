package models

import (
	"time"
)

type Post struct {
	Id      int64  `json:"id"`
	Author  *User  `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Comments []*Comment `json:"comments"`
	Likes []*User `json:"likes"`
	Published bool `json:"published"`
	Created_At *time.Time  `json:"created_at"`
}