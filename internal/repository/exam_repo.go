package repository

import (
	"exam-system/internal/models"

	"gorm.io/gorm"
)

type ExamRepository interface {
	Create(exam *models.Examination) error
	FindAll() ([]models.Examination, error)
	FindByID(id string) (*models.Examination, error)
	Update(exam *models.Examination) error
	Delete(id string) error
}

type examRepo struct {
	db *gorm.DB
}

func NewExamRepository(db *gorm.DB) ExamRepository {
	return &examRepo{db: db}
}

func (r *examRepo) FindByID(id string) (*models.Examination, error) {
	var exam models.Examination
	err := r.db.Preload("Questions.Selections").First(&exam, "id = ?", id).Error
	return &exam, err
}

func (r *examRepo) FindAll() ([]models.Examination, error) {
	var exams []models.Examination
	err := r.db.Find(&exams).Error
	return exams, err
}

func (r *examRepo) Create(exam *models.Examination) error {
	return r.db.Create(exam).Error
}

func (r *examRepo) Update(exam *models.Examination) error {
	return r.db.Save(exam).Error
}

func (r *examRepo) Delete(id string) error {
	return r.db.Delete(&models.Examination{}, "id = ?", id).Error
}
