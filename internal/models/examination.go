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

type Question struct {
	ID            string              `gorm:"primaryKey;type:varchar(36)"`
	ExaminationID string              `gorm:"not null;type:varchar(36)"`
	QuestionType  string              `gorm:"not null;type:enum('MULTIPLE_CHOICE', 'ESSAY', 'TRUE_FALSE', 'FILL_IN_BLANK')"`
	Description   string              `gorm:"not null"`
	Selections    []QuestionSelection `gorm:"foreignKey:QuestionID"`
}

type QuestionSelection struct {
	ID         string `gorm:"primaryKey;type:varchar(36)"`
	QuestionID string `gorm:"not null;type:varchar(36)"`
	Name       string `gorm:"not null"`
	Code       string `gorm:"not null"`
}
