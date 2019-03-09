package models

import "time"


type BaseModel struct {
	ID int64
	CreatedDate time.Time
	UpdatedDate time.Time
}