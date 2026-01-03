package handlers

import (
	"exam-system/internal/dto"
	"exam-system/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QuestionHandler struct {
	service service.QuestionService
}

func NewQuestionHandler(s service.QuestionService) *QuestionHandler {
	return &QuestionHandler{service: s}
}

// POST /questions
func (h *QuestionHandler) Create(c *gin.Context) {
	var req dto.CreateQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.CreateQuestion(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// GET /questions?examId=...
func (h *QuestionHandler) GetList(c *gin.Context) {
	examID := c.Query("examId")
	if examID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "examId is required"})
		return
	}

	res, err := h.service.GetListByExam(examID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GET /questions/:id
func (h *QuestionHandler) GetDetail(c *gin.Context) {
	id := c.Param("id")
	res, err := h.service.GetDetail(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// PUT /questions/:id
func (h *QuestionHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.UpdateQuestion(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DELETE /questions/:id
func (h *QuestionHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteQuestion(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
