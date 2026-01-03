package models

import (
	"time"
)

type Examination struct {
	ID            string `gorm:"primaryKey;type:varchar(36)"`
	Title         string `gorm:"not null"`
	Description   string
	Duration      int `gorm:"not null"`
	QuestionCount int `gorm:"default:0"`

	Questions []Question `gorm:"foreignKey:ExaminationID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
