package handlers

import (
	"exam-system/internal/dto"
	"exam-system/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PassageHandler struct {
	service service.PassageService
}

func NewPassageHandler(s service.PassageService) *PassageHandler {
	return &PassageHandler{service: s}
}

// Create godoc
// @Summary      Create reading passage
// @Description  Create a new reading passage
// @Tags         passages
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreatePassageRequest  true  "Passage creation request"
// @Success      201      {object}  dto.PassageResponse
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /passages [post]
func (h *PassageHandler) Create(c *gin.Context) {
	var req dto.CreatePassageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// GetList godoc
// @Summary      List reading passages
// @Description  Get list of all reading passages
// @Tags         passages
// @Accept       json
// @Produce      json
// @Success      200  {array}   dto.PassageResponse
// @Failure      500  {object}  map[string]string
// @Router       /passages [get]
func (h *PassageHandler) GetList(c *gin.Context) {
	res, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetDetail godoc
// @Summary      Get passage by ID
// @Description  Get detailed information about a specific reading passage
// @Tags         passages
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Passage ID"
// @Success      200  {object}  dto.PassageResponse
// @Failure      404  {object}  map[string]string
// @Router       /passages/{id} [get]
func (h *PassageHandler) GetDetail(c *gin.Context) {
	id := c.Param("id")
	res, err := h.service.GetDetail(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Passage not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// Update godoc
// @Summary      Update passage
// @Description  Update an existing reading passage by ID
// @Tags         passages
// @Accept       json
// @Produce      json
// @Param        id       path      string                     true  "Passage ID"
// @Param        request  body      dto.UpdatePassageRequest  true  "Passage update request"
// @Success      200      {object}  dto.PassageResponse
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /passages/{id} [put]
func (h *PassageHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdatePassageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.Update(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// Delete godoc
// @Summary      Delete passage
// @Description  Delete a reading passage by ID
// @Tags         passages
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Passage ID"
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /passages/{id} [delete]
func (h *PassageHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
