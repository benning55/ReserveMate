package model

import "time"

type User struct {
	ID         uint       `gorm:"column:id;primary_key;" json:"id"`
	Name       string     `json:"name"`
	Age        string     `json:"age"`
	Email      string     `json:"email"`
	LineUserID string     `json:"line_user_id"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

func (User) TableName() string { return "users" }
