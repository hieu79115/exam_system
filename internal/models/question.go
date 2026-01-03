package models

type Question struct {
	ID               string  `gorm:"primaryKey;type:varchar(36)"`
	ExaminationID    string  `gorm:"not null;type:varchar(36)"`
	ReadingPassageID *string `gorm:"type:varchar(36)"`

	QuestionType  string `gorm:"not null"`
	Description   string `gorm:"not null;type:text"`
	MaxTextLength int

	Selections []QuestionSelection `gorm:"foreignKey:QuestionID"`
}

type QuestionSelection struct {
	ID         string `gorm:"primaryKey;type:varchar(36)"`
	QuestionID string `gorm:"not null;type:varchar(36)"`
	Name       string `gorm:"not null"`
	Code       string `gorm:"not null"`
}
