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

// POST /passages
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

// GET /passages
func (h *PassageHandler) GetList(c *gin.Context) {
	res, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GET /passages/:id
func (h *PassageHandler) GetDetail(c *gin.Context) {
	id := c.Param("id")
	res, err := h.service.GetDetail(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Passage not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// PUT /passages/:id
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

// DELETE /passages/:id
func (h *PassageHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
