package models

import "./basemodel.go"

type PostType struct {
	BaseModel
	ParentID int64
	Title string
	Summary string
	Status int
}
