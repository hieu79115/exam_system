package models

import "time"

type ReadingPassage struct {
	ID          string `gorm:"primaryKey;type:varchar(36)"`
	Description string `gorm:"type:text"`
	ContentText string `gorm:"column:content_text;type:mediumtext;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
