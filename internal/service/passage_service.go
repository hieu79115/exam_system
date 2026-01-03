package service

import (
	"exam-system/internal/dto"
	"exam-system/internal/models"
	"exam-system/internal/repository"

	"github.com/google/uuid"
)

type PassageService interface {
	Create(req dto.CreatePassageRequest) (*dto.PassageResponse, error)
	GetAll() ([]dto.PassageResponse, error)
	GetDetail(id string) (*dto.PassageResponse, error)
	Update(id string, req dto.UpdatePassageRequest) (*dto.PassageResponse, error)
	Delete(id string) error
}

type passageService struct {
	repo repository.PassageRepository
}

func NewPassageService(repo repository.PassageRepository) PassageService {
	return &passageService{repo: repo}
}

func (s *passageService) Create(req dto.CreatePassageRequest) (*dto.PassageResponse, error) {
	finalID := req.ID
	if finalID == "" {
		finalID = uuid.New().String()
	}

	model := &models.ReadingPassage{
		ID:          finalID,
		ContentText: req.Text,
		Description: req.Description,
	}

	if err := s.repo.Create(model); err != nil {
		return nil, err
	}

	return &dto.PassageResponse{
		ID:          model.ID,
		Text:        model.ContentText,
		Description: model.Description,
	}, nil
}

func (s *passageService) GetAll() ([]dto.PassageResponse, error) {
	modelsList, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var result []dto.PassageResponse
	for _, m := range modelsList {
		result = append(result, dto.PassageResponse{
			ID:          m.ID,
			Text:        m.ContentText,
			Description: m.Description,
		})
	}
	if result == nil {
		result = []dto.PassageResponse{}
	}
	return result, nil
}

func (s *passageService) GetDetail(id string) (*dto.PassageResponse, error) {
	model, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.PassageResponse{
		ID:          model.ID,
		Text:        model.ContentText,
		Description: model.Description,
	}, nil
}

func (s *passageService) Update(id string, req dto.UpdatePassageRequest) (*dto.PassageResponse, error) {

	model, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	model.ContentText = req.Text
	model.Description = req.Description

	if err := s.repo.Update(model); err != nil {
		return nil, err
	}

	return &dto.PassageResponse{
		ID:          model.ID,
		Text:        model.ContentText,
		Description: model.Description,
	}, nil
}

func (s *passageService) Delete(id string) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
