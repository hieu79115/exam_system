package handlers

import (
	"exam-system/internal/dto"
	"exam-system/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExamHandler struct {
	service service.ExamService
}

func NewExamHandler(s service.ExamService) *ExamHandler {
	return &ExamHandler{service: s}
}

// GET /examinations/:id
func (h *ExamHandler) GetDetail(c *gin.Context) {
	id := c.Param("id")

	examModel, err := h.service.GetExamDetail(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exam not found"})
		return
	}

	respone := dto.ExamResponse{
		ID:            examModel.ID,
		Title:         examModel.Title,
		Description:   examModel.Description,
		Duration:      examModel.Duration,
		QuestionCount: examModel.QuestionCount,
	}

	c.JSON(http.StatusOK, respone)
}

// GET /examinations
func (h *ExamHandler) GetList(c *gin.Context) {
	exams, err := h.service.GetExamList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, exams)
}

// POST /examinations
func (h *ExamHandler) Create(c *gin.Context) {
	var req dto.CreateExamRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createExam, err := h.service.CreateExam(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createExam)
}

// PUT /examinations/:id
func (h *ExamHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateExamRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedExam, err := h.service.UpdateExam(id, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedExam)
}

// DELETE /examinations/:id
func (h *ExamHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteExam(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete examination successfully"})
}
