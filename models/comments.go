package models

import "./basemodel.go"

type Comments struct {
	BaseModel
	Content string
	PostID int64
	ParentID int64
	PostTypeID int
	Status int
	LikeCount int
	UnlikeCount int
}
