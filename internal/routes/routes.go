package routes

import (
	"exam-system/internal/app"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "exam-system/docs" // Import generated docs
)

func SetupRoutes(r *gin.Engine, h *app.AppHandlers) {
	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	{
		// exam Routes
		exam := api.Group("/examinations")

		exam.GET("/:id", h.ExamHandler.GetDetail)
		exam.GET("", h.ExamHandler.GetList)
		exam.POST("", h.ExamHandler.Create)
		exam.PUT("/:id", h.ExamHandler.Update)
		exam.DELETE("/:id", h.ExamHandler.Delete)

		// Passage Routes
		passages := api.Group("/passages")

		passages.POST("", h.PassageHandler.Create)
		passages.GET("", h.PassageHandler.GetList)
		passages.GET("/:id", h.PassageHandler.GetDetail)
		passages.PUT("/:id", h.PassageHandler.Update)
		passages.DELETE("/:id", h.PassageHandler.Delete)

		// Question Routes
		questions := api.Group("/questions")
		questions.POST("", h.QuestionHandler.Create)
		questions.GET("", h.QuestionHandler.GetList)
		questions.GET("/:id", h.QuestionHandler.GetDetail)
		questions.PUT("/:id", h.QuestionHandler.Update)
		questions.DELETE("/:id", h.QuestionHandler.Delete)
	}

}
