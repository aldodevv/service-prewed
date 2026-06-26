package http

import (
	"net/http"
	"service-wedding/internal/domain"
	"service-wedding/internal/usecase"

	"github.com/gin-gonic/gin"
)

type PublicHandler struct {
	contextUsecase usecase.ContextUsecase
	guestUsecase   usecase.GuestUsecase
	themeUsecase   usecase.ThemeUsecase
	assetUsecase   usecase.AssetUsecase
}

func NewPublicHandler(
	cu usecase.ContextUsecase,
	gu usecase.GuestUsecase,
	tu usecase.ThemeUsecase,
	au usecase.AssetUsecase,
) *PublicHandler {
	return &PublicHandler{
		contextUsecase: cu,
		guestUsecase:   gu,
		themeUsecase:   tu,
		assetUsecase:   au,
	}
}

func (h *PublicHandler) GetPublicWedding(c *gin.Context) {
	contextSlug := c.Param("theme") // context slug is mapped to theme in path
	guestSlug := c.Param("guest")

	ctx := c.Request.Context()

	// 1. Get client context by slug
	clientCtx, err := h.contextUsecase.GetBySlug(ctx, contextSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if clientCtx == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wedding context not found"})
		return
	}

	// 2. Get guest by slug inside this context
	// Fetching all guests is fine since guest list per context is usually small, but searching by slug is better.
	// Let's check: can we use GetAllByContextID and find it? Yes, we can!
	guests, err := h.guestUsecase.GetAllByContextID(ctx, clientCtx.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var foundGuest domain.Guest
	found := false
	for _, g := range guests {
		if g.Slug == guestSlug {
			foundGuest = g
			found = true
			break
		}
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guest not found in this wedding invitation"})
		return
	}

	// 3. Load associated theme
	themeSlug := "default"
	htmlToRender := clientCtx.RenderHTML
	var themeData map[string]interface{}

	if clientCtx.ThemeID > 0 {
		theme, err := h.themeUsecase.GetByID(ctx, clientCtx.ThemeID)
		if err == nil && theme != nil {
			themeSlug = theme.Slug
			themeData = theme.ThemeData
			if htmlToRender == "" {
				htmlToRender = theme.RenderHTML
			}
		}
	}

	// 4. Fetch all assets to return
	assets, _ := h.assetUsecase.GetAll(ctx)
	if assets == nil {
		assets = []domain.Asset{}
	}

	// Respond with payload
	c.JSON(http.StatusOK, gin.H{
		"theme":       themeSlug,
		"guest":       foundGuest.Name,
		"context":     clientCtx.Name,
		"render_html": htmlToRender,
		"assets":      assets,
		"theme_data":  themeData,
	})
}
