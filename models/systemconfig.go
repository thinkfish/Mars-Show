package models

import "./basemodel.go"

type SystemConfig struct {
	BaseModel
	ConfigKey string
	ConfigValue string
	Status int
}