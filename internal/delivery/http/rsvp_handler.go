package http

import (
	"net/http"
	"service-wedding/internal/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RSVPHandler struct {
	rsvpUsecase domain.RSVPUsecase
}

func NewRSVPHandler(ru domain.RSVPUsecase) *RSVPHandler {
	return &RSVPHandler{rsvpUsecase: ru}
}

func (h *RSVPHandler) SubmitRSVP(c *gin.Context) {
	var input struct {
		GuestID    int64  `json:"guest_id"`
		Attendance string `json:"attendance"`
		GuestCount int    `json:"guest_count"`
		Message    string `json:"message"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body payload"})
		return
	}

	rsvp := domain.RSVP{
		GuestID:    input.GuestID,
		Attendance: input.Attendance,
		GuestCount: input.GuestCount,
		Message:    input.Message,
	}

	err := h.rsvpUsecase.SubmitRSVP(c.Request.Context(), &rsvp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "RSVP submitted successfully!",
		"rsvp":    rsvp,
	})
}

func (h *RSVPHandler) GetAllByContextID(c *gin.Context) {
	idStr := c.Param("id")
	contextID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid context ID"})
		return
	}

	rsvps, err := h.rsvpUsecase.GetAllByContextID(c.Request.Context(), contextID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rsvps)
}
