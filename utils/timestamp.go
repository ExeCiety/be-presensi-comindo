package utils

import (
	"time"

	"gorm.io/gorm"
)

type Timestamp struct {
	CreatedAt time.Time       `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
