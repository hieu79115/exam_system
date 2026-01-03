package app

import (
	"exam-system/internal/handlers"
	"exam-system/internal/repository"
	"exam-system/internal/service"

	"gorm.io/gorm"
)

type AppHandlers struct {
	ExamHandler *handlers.ExamHandler
}

func InitApp(db *gorm.DB) *AppHandlers {
	// 1. Init Repos
	examRepo := repository.NewExamRepository(db)

	// 2. Init Services
	examService := service.NewExamService(examRepo)

	// 3. Init Handlers và trả về struct
	return &AppHandlers{
		ExamHandler: handlers.NewExamHandler(examService),
	}
}
