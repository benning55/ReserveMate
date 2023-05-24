package model

import "time"

type User struct {
	ID         uint       `gorm:"column:id;primary_key;" json:"id"`
	Name       string     `json:"name"`
	Email      string     `gorm:"index:idx_email,unique" json:"email"`
	LineUserID string     `json:"line_user_id"`
	IsAdmin    bool       `json:"is_admin"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

func (User) TableName() string { return "users" }
