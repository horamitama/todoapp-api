package entity

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title     string `json:"title"`
	Detail    string `json:"detail"`
	Status    string `json:"status"`
	UserRefer uint   `json:"user_refer"`
}
