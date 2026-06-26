package http

import (
	"net/http"
	"service-wedding/internal/domain"
	"service-wedding/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContextHandler struct {
	contextUsecase usecase.ContextUsecase
}

func NewContextHandler(cu usecase.ContextUsecase) *ContextHandler {
	return &ContextHandler{contextUsecase: cu}
}

func (h *ContextHandler) GetAll(c *gin.Context) {
	contexts, err := h.contextUsecase.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contexts)
}

func (h *ContextHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid context ID"})
		return
	}

	ctxVal, err := h.contextUsecase.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if ctxVal == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Context not found"})
		return
	}

	c.JSON(http.StatusOK, ctxVal)
}

func (h *ContextHandler) Create(c *gin.Context) {
	var ctxVal domain.Context
	if err := c.ShouldBindJSON(&ctxVal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.contextUsecase.Create(c.Request.Context(), &ctxVal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, ctxVal)
}

func (h *ContextHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid context ID"})
		return
	}

	var ctxVal domain.Context
	if err := c.ShouldBindJSON(&ctxVal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.contextUsecase.Update(c.Request.Context(), id, &ctxVal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ctxVal)
}

func (h *ContextHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid context ID"})
		return
	}

	err = h.contextUsecase.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Context deleted successfully"})
}
