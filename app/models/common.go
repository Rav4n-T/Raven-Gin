package models

import (
	"time"

	"gorm.io/gorm"
)

type ID struct {
	ID int `json:"id" gorm:"primarykey"`
}

type Timestamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SoftDelete struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
