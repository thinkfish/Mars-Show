package models

import "./basemodel.go"

type Posts struct {
	BaseModel
	Title string
	Summary string
	Content string
	LikeCount int
	UnlikeCount int
	CommentCount int
	Status int
}
