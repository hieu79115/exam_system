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

// Create godoc
// @Summary      Add question to existing exam
// @Description  Add a new question with selections to an existing examination
// @Tags         questions
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateQuestionRequest  true  "Question creation request"
// @Success      201      {object}  dto.QuestionRes
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /questions [post]
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

// GetList godoc
// @Summary      List questions by exam (Admin)
// @Description  Get list of questions for a specific examination (for admin question management)
// @Tags         questions
// @Accept       json
// @Produce      json
// @Param        examId  query     string  true  "Examination ID"
// @Success      200     {array}   dto.QuestionRes
// @Failure      400     {object}  map[string]string
// @Failure      500     {object}  map[string]string
// @Router       /questions [get]
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

// GetDetail godoc
// @Summary      Get question by ID
// @Description  Get detailed information about a specific question with its selections
// @Tags         questions
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Question ID"
// @Success      200  {object}  dto.QuestionRes
// @Failure      404  {object}  map[string]string
// @Router       /questions/{id} [get]
func (h *QuestionHandler) GetDetail(c *gin.Context) {
	id := c.Param("id")
	res, err := h.service.GetDetail(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// Update godoc
// @Summary      Update question
// @Description  Update an existing question. Old selections will be deleted and replaced with new ones.
// @Tags         questions
// @Accept       json
// @Produce      json
// @Param        id       path      string                      true  "Question ID"
// @Param        request  body      dto.UpdateQuestionRequest  true  "Question update request"
// @Success      200      {object}  dto.QuestionRes
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /questions/{id} [put]
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

// Delete godoc
// @Summary      Delete question
// @Description  Delete a question by ID. Selections will be cascade deleted.
// @Tags         questions
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Question ID"
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /questions/{id} [delete]
func (h *QuestionHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteQuestion(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
