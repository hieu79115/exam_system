package repository

import (
	"exam-system/internal/models"

	"gorm.io/gorm"
)

type PassageRepository interface {
	Create(p *models.ReadingPassage) error
	FindAll() ([]models.ReadingPassage, error)
	FindByID(id string) (*models.ReadingPassage, error)
	Update(p *models.ReadingPassage) error
	Delete(id string) error
}

type passageRepo struct {
	db *gorm.DB
}

func NewPassageRepository(db *gorm.DB) PassageRepository {
	return &passageRepo{db: db}
}

func (r *passageRepo) Create(passage *models.ReadingPassage) error {
	return r.db.Create(passage).Error
}

func (r *passageRepo) FindAll() ([]models.ReadingPassage, error) {
	var passages []models.ReadingPassage
	err := r.db.Order("created_at desc").Find(&passages).Error
	return passages, err
}

func (r *passageRepo) FindByID(id string) (*models.ReadingPassage, error) {
	var passage models.ReadingPassage
	err := r.db.First(&passage, "id = ?", id).Error
	return &passage, err
}

func (r *passageRepo) Update(p *models.ReadingPassage) error {
	return r.db.Save(p).Error
}

func (r *passageRepo) Delete(id string) error {
	return r.db.Delete(&models.ReadingPassage{}, "id = ?", id).Error
}
