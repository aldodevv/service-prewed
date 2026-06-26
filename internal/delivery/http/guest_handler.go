package http

import (
	"net/http"
	"service-wedding/internal/domain"
	"service-wedding/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GuestHandler struct {
	guestUsecase usecase.GuestUsecase
}

func NewGuestHandler(gu usecase.GuestUsecase) *GuestHandler {
	return &GuestHandler{guestUsecase: gu}
}

func (h *GuestHandler) GetAll(c *gin.Context) {
	contextIDStr := c.Param("id")
	contextID, err := strconv.ParseInt(contextIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid context ID"})
		return
	}

	guests, err := h.guestUsecase.GetAllByContextID(c.Request.Context(), contextID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, guests)
}

func (h *GuestHandler) Create(c *gin.Context) {
	contextIDStr := c.Param("id")
	contextID, err := strconv.ParseInt(contextIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid context ID"})
		return
	}

	var guest domain.Guest
	if err := c.ShouldBindJSON(&guest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	guest.ContextID = contextID
	err = h.guestUsecase.Create(c.Request.Context(), &guest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, guest)
}

func (h *GuestHandler) Update(c *gin.Context) {
	contextIDStr := c.Param("id")
	contextID, err := strconv.ParseInt(contextIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid context ID"})
		return
	}

	guestIDStr := c.Param("guestId")
	guestID, err := strconv.ParseInt(guestIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid guest ID"})
		return
	}

	var req struct {
		Name string `json:"name" binding:"required"`
		Slug string `json:"slug" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.guestUsecase.Update(c.Request.Context(), contextID, guestID, req.Name, req.Slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Guest updated successfully"})
}

func (h *GuestHandler) Delete(c *gin.Context) {
	contextIDStr := c.Param("id")
	contextID, err := strconv.ParseInt(contextIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid context ID"})
		return
	}

	guestIDStr := c.Param("guestId")
	guestID, err := strconv.ParseInt(guestIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid guest ID"})
		return
	}

	err = h.guestUsecase.Delete(c.Request.Context(), contextID, guestID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Guest deleted successfully"})
}
