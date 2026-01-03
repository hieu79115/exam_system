package service

import (
	"exam-system/internal/dto"
	"exam-system/internal/models"
	"exam-system/internal/repository"
)

type ExamService interface {
	GetExamDetail(id string) (*models.Examination, error)
	GetExamList() ([]dto.ExamListItemResponse, error)
	CreateExam(req dto.CreateExamRequest) (*models.Examination, error)
	UpdateExam(id string, req dto.UpdateExamRequest) (*models.Examination, error)
	DeleteExam(id string) error
}

type examService struct {
	repo repository.ExamRepository
}

func NewExamService(repo repository.ExamRepository) ExamService {
	return &examService{repo: repo}
}

func (s *examService) GetExamDetail(id string) (*models.Examination, error) {
	return s.repo.FindByID(id)
}

func (s *examService) GetExamList() ([]dto.ExamListItemResponse, error) {
	examsDB, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var response []dto.ExamListItemResponse
	for _, e := range examsDB {
		response = append(response, dto.ExamListItemResponse{
			ID:            e.ID,
			Title:         e.Title,
			Duration:      e.Duration,
			QuestionCount: e.QuestionCount,
		})
	}

	if response == nil {
		return []dto.ExamListItemResponse{}, nil
	}

	return response, nil
}

func (s *examService) CreateExam(req dto.CreateExamRequest) (*models.Examination, error) {
	newExam := &models.Examination{
		ID:            req.ID,
		Title:         req.Title,
		Description:   req.Description,
		Duration:      req.Duration,
		QuestionCount: 0,
	}

	if err := s.repo.Create(newExam); err != nil {
		return nil, err
	}

	return newExam, nil
}

func (s *examService) UpdateExam(id string, req dto.UpdateExamRequest) (*models.Examination, error) {

	exam, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	exam.Title = req.Title
	exam.Description = req.Description
	exam.Duration = req.Duration

	if err := s.repo.Update(exam); err != nil {
		return nil, err
	}

	return exam, nil
}

func (s *examService) DeleteExam(id string) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
