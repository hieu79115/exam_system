package repository

import (
	"exam-system/internal/models"

	"gorm.io/gorm"
)

type QuestionRepository interface {
	Create(q *models.Question) error
	FindAll(examID string) ([]models.Question, error)
	FindByID(id string) (*models.Question, error)
	Update(q *models.Question) error
	Delete(id string) error
}

type questionRepo struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) QuestionRepository {
	return &questionRepo{db: db}
}

func (r *questionRepo) Create(q *models.Question) error {
	return r.db.Create(q).Error
}

func (r *questionRepo) FindAll(examID string) ([]models.Question, error) {
	var questions []models.Question
	err := r.db.Preload("Selections").Where("examination_id = ?", examID).Find(&questions).Error
	return questions, err
}

func (r *questionRepo) FindByID(id string) (*models.Question, error) {
	var question models.Question
	err := r.db.Preload("Selections").First(&question, "id = ?", id).Error
	return &question, err
}

func (r *questionRepo) Update(q *models.Question) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&models.QuestionSelection{}, "question_id = ?", q.ID).Error; err != nil {
			return err
		}

		if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(q).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *questionRepo) Delete(id string) error {
	return r.db.Delete(&models.Question{}, "id = ?", id).Error
}
