package routes

import (
	"exam-system/internal/app"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *app.AppHandlers) {
	api := r.Group("/api/v1")

	exam := api.Group("/examinations")
	exam.GET("/:id", h.ExamHandler.GetDetail)
	exam.GET("", h.ExamHandler.GetList)
	exam.POST("", h.ExamHandler.Create)
	exam.PUT("/:id", h.ExamHandler.Update)
	exam.DELETE("/:id", h.ExamHandler.Delete)
}
