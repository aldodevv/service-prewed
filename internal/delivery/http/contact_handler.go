package http

import (
	"net/http"
	"service-wedding/internal/domain"
	"service-wedding/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	contactUsecase usecase.ContactUsecase
}

func NewContactHandler(cu usecase.ContactUsecase) *ContactHandler {
	return &ContactHandler{contactUsecase: cu}
}

func (h *ContactHandler) Create(c *gin.Context) {
	var msg domain.ContactMessage
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	if err := h.contactUsecase.Create(ctx, &msg); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, msg)
}

func (h *ContactHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	messages, err := h.contactUsecase.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if messages == nil {
		messages = []domain.ContactMessage{}
	}

	c.JSON(http.StatusOK, messages)
}
