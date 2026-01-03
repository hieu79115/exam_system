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

// GetDetail godoc
// @Summary      Get examination by ID
// @Description  Get FULL examination details including questions and selections for taking the exam
// @Tags         examinations
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Examination ID"
// @Success      200  {object}  dto.ExamDetailResponse
// @Failure      404  {object}  map[string]string
// @Router       /examinations/{id} [get]
func (h *ExamHandler) GetDetail(c *gin.Context) {
	id := c.Param("id")

	examModel, err := h.service.GetExamDetail(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exam not found"})
		return
	}

	// Build exam info
	examResponse := dto.ExamResponse{
		ID:            examModel.ID,
		Title:         examModel.Title,
		Description:   examModel.Description,
		Duration:      examModel.Duration,
		QuestionCount: examModel.QuestionCount,
	}

	// Build questions with selections
	var questions []dto.QuestionResponse
	for _, q := range examModel.Questions {
		var selections []dto.SelectionResponse
		for _, sel := range q.Selections {
			selections = append(selections, dto.SelectionResponse{
				ID:   sel.ID,
				Name: sel.Name,
				Code: sel.Code,
			})
		}

		questions = append(questions, dto.QuestionResponse{
			ID:               q.ID,
			Type:             q.QuestionType,
			Description:      q.Description,
			MaxText:          q.MaxTextLength,
			ReadingPassageID: q.ReadingPassageID,
			Selections:       selections,
		})
	}

	response := dto.ExamDetailResponse{
		Exam:      examResponse,
		Questions: questions,
		Passages:  []dto.PassageResponse{}, // Add passage logic if needed
	}

	c.JSON(http.StatusOK, response)
}

// GetList godoc
// @Summary      List examinations
// @Description  Get list of all examinations with basic info only (ID, Title, Duration)
// @Tags         examinations
// @Accept       json
// @Produce      json
// @Success      200  {array}   dto.ExamListItemResponse
// @Failure      500  {object}  map[string]string
// @Router       /examinations [get]
func (h *ExamHandler) GetList(c *gin.Context) {
	exams, err := h.service.GetExamList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, exams)
}

// Create godoc
// @Summary      Create examination
// @Description  Create a new examination with questions and selections inline
// @Tags         examinations
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateExamRequest  true  "Exam creation request with questions"
// @Success      201      {object}  dto.ExamResponse
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /examinations [post]
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

// Update godoc
// @Summary      Update examination basic info
// @Description  Update examination basic info only (Title, Description, Duration). Does not modify questions.
// @Tags         examinations
// @Accept       json
// @Produce      json
// @Param        id       path      string                  true  "Examination ID"
// @Param        request  body      dto.UpdateExamRequest  true  "Exam update request"
// @Success      200      {object}  dto.ExamResponse
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /examinations/{id} [put]
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

// Delete godoc
// @Summary      Delete examination
// @Description  Delete an examination by ID. Questions belonging to this exam will be cascade deleted.
// @Tags         examinations
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Examination ID"
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /examinations/{id} [delete]
func (h *ExamHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteExam(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete examination successfully"})
}
