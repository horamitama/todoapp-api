package entity

type Task struct {
	Base
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Status string `json:"status"`
	User   User   `json:"user" gorm:"embedded"`
}
