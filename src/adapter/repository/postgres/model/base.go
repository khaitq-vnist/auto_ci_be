package model

import "time"

type BaseModel struct {
	ID        int64     `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
