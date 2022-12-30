package model

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	User    User   `json:"user"`
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
