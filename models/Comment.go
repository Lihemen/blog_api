package models

import (
	"time"
)

type Comment struct {
	Id      int64 `json:"id"`
	Author *User 	`json:"author"`
	Body []byte 	`json:"body"`
	Upvotes []*User `json:"upvotes"`
	Downvotes []*User `json:"downvotes"`
	Created_At *time.Time  `json:"created_at"`	
}