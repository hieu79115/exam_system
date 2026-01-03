package routes

import (
	"exam-system/internal/app"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *app.AppHandlers) {
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
	}

}
