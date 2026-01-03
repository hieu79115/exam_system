package service

import (
	"exam-system/internal/dto"
	"exam-system/internal/models"
	"exam-system/internal/repository"

	"github.com/google/uuid"
)

type QuestionService interface {
	CreateQuestion(req dto.CreateQuestionRequest) (*dto.QuestionRes, error)
	GetListByExam(examID string) ([]dto.QuestionRes, error)
	GetDetail(id string) (*dto.QuestionRes, error)
	UpdateQuestion(id string, req dto.UpdateQuestionRequest) (*dto.QuestionRes, error)
	DeleteQuestion(id string) error
}

type questionService struct {
	repo     repository.QuestionRepository
	examRepo repository.ExamRepository
}

func NewQuestionService(repo repository.QuestionRepository, examRepo repository.ExamRepository) QuestionService {
	return &questionService{repo: repo, examRepo: examRepo}
}

func (s *questionService) CreateQuestion(req dto.CreateQuestionRequest) (*dto.QuestionRes, error) {
	quesID := uuid.New().String()

	var selectionsModel []models.QuestionSelection
	var selectionsRes []dto.SelectionRes

	for _, sel := range req.Selections {
		selID := uuid.New().String()

		selectionsModel = append(selectionsModel, models.QuestionSelection{
			ID:         selID,
			QuestionID: quesID,
			Name:       sel.Name,
			Code:       sel.Code,
		})

		selectionsRes = append(selectionsRes, dto.SelectionRes{
			ID:   selID,
			Name: sel.Name,
			Code: sel.Code,
		})
	}

	questionModel := &models.Question{
		ID:               quesID,
		ExaminationID:    req.ExaminationID,
		ReadingPassageID: req.ReadingPassageID,
		QuestionType:     req.Type,
		Description:      req.Description,
		MaxTextLength:    req.MaxText,
		Selections:       selectionsModel,
	}

	if err := s.repo.Create(questionModel); err != nil {
		return nil, err
	}

	// Update question count
	exam, err := s.examRepo.FindByID(req.ExaminationID)
	if err == nil {
		exam.QuestionCount++
		s.examRepo.Update(exam)
	}

	return &dto.QuestionRes{
		ID:               questionModel.ID,
		ExaminationID:    questionModel.ExaminationID,
		ReadingPassageID: questionModel.ReadingPassageID,
		Type:             questionModel.QuestionType,
		Description:      questionModel.Description,
		Selections:       selectionsRes,
	}, nil
}

func (s *questionService) GetListByExam(examID string) ([]dto.QuestionRes, error) {
	questions, err := s.repo.FindAll(examID)
	if err != nil {
		return nil, err
	}

	var res []dto.QuestionRes
	for _, q := range questions {
		var selections []dto.SelectionRes
		for _, sel := range q.Selections {
			selections = append(selections, dto.SelectionRes{
				ID:   sel.ID,
				Name: sel.Name,
				Code: sel.Code,
			})
		}

		res = append(res, dto.QuestionRes{
			ID:               q.ID,
			ExaminationID:    q.ExaminationID,
			ReadingPassageID: q.ReadingPassageID,
			Type:             q.QuestionType,
			Description:      q.Description,
			Selections:       selections,
		})
	}
	if res == nil {
		res = []dto.QuestionRes{}
	}
	return res, nil
}

func (s *questionService) GetDetail(id string) (*dto.QuestionRes, error) {
	q, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	var selRes []dto.SelectionRes
	for _, sel := range q.Selections {
		selRes = append(selRes, dto.SelectionRes{
			ID:   sel.ID,
			Name: sel.Name,
			Code: sel.Code,
		})
	}

	return &dto.QuestionRes{
		ID:               q.ID,
		ExaminationID:    q.ExaminationID,
		ReadingPassageID: q.ReadingPassageID,
		Type:             q.QuestionType,
		Description:      q.Description,
		Selections:       selRes,
	}, nil
}

func (s *questionService) UpdateQuestion(id string, req dto.UpdateQuestionRequest) (*dto.QuestionRes, error) {
	if _, err := s.repo.FindByID(id); err != nil {
		return nil, err
	}

	var selectionsModel []models.QuestionSelection
	var selRes []dto.SelectionRes

	for _, sel := range req.Selections {
		newSelID := uuid.New().String()
		selectionsModel = append(selectionsModel, models.QuestionSelection{
			ID:         newSelID,
			QuestionID: id,
			Name:       sel.Name,
			Code:       sel.Code,
		})
		selRes = append(selRes, dto.SelectionRes{
			ID: newSelID, Name: sel.Name, Code: sel.Code,
		})
	}

	model := &models.Question{
		ID:               id,
		ExaminationID:    req.ExaminationID,
		ReadingPassageID: req.ReadingPassageID,
		QuestionType:     req.Type,
		Description:      req.Description,
		MaxTextLength:    req.MaxText,
		Selections:       selectionsModel,
	}

	if err := s.repo.Update(model); err != nil {
		return nil, err
	}

	return &dto.QuestionRes{
		ID:               model.ID,
		ExaminationID:    model.ExaminationID,
		ReadingPassageID: model.ReadingPassageID,
		Type:             model.QuestionType,
		Description:      model.Description,
		Selections:       selRes,
	}, nil
}

func (s *questionService) DeleteQuestion(id string) error {
	// Get question to find exam ID before deleting
	q, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if err := s.repo.Delete(id); err != nil {
		return err
	}

	// Update question count
	exam, err := s.examRepo.FindByID(q.ExaminationID)
	if err == nil && exam.QuestionCount > 0 {
		exam.QuestionCount--
		s.examRepo.Update(exam)
	}

	return nil
}
