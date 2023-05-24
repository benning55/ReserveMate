package model

import "time"

type Merchant struct {
	ID          uint       `gorm:"column:id;primary_key;" json:"id"`
	Name        string     `json:"name"`
	Address     string     `json:"address"`
	PhoneNumber string     `json:"phone_number"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func (Merchant) TableName() string { return "merchants" }
