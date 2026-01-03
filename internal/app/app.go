package app

import (
	"exam-system/internal/handlers"
	"exam-system/internal/repository"
	"exam-system/internal/service"

	"gorm.io/gorm"
)

type AppHandlers struct {
	ExamHandler     *handlers.ExamHandler
	PassageHandler  *handlers.PassageHandler
	QuestionHandler *handlers.QuestionHandler
}

func InitApp(db *gorm.DB) *AppHandlers {
	// 1. Init Repos
	examRepo := repository.NewExamRepository(db)
	passageRepo := repository.NewPassageRepository(db)
	questionRepo := repository.NewQuestionRepository(db)

	// 2. Init Services
	examService := service.NewExamService(examRepo)
	passageService := service.NewPassageService(passageRepo)
	questionService := service.NewQuestionService(questionRepo, examRepo)

	// 3. Init Handlers
	return &AppHandlers{
		ExamHandler:     handlers.NewExamHandler(examService),
		PassageHandler:  handlers.NewPassageHandler(passageService),
		QuestionHandler: handlers.NewQuestionHandler(questionService),
	}
}
