package http

import (
	"io"
	"net/http"
	"path/filepath"
	"service-wedding/internal/domain"
	"service-wedding/internal/usecase"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type AssetHandler struct {
	assetUsecase usecase.AssetUsecase
}

func NewAssetHandler(au usecase.AssetUsecase) *AssetHandler {
	return &AssetHandler{assetUsecase: au}
}

func (h *AssetHandler) GetAll(c *gin.Context) {
	assets, err := h.assetUsecase.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, assets)
}

func (h *AssetHandler) Create(c *gin.Context) {
	var asset domain.Asset
	if err := c.ShouldBindJSON(&asset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.assetUsecase.Create(c.Request.Context(), &asset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, asset)
}

func (h *AssetHandler) Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	name := c.PostForm("name")
	if name == "" {
		name = header.Filename
	}

	fileType := c.PostForm("type")
	if fileType == "" {
		mimeType := header.Header.Get("Content-Type")
		if strings.HasPrefix(mimeType, "image/") {
			fileType = "image"
		} else if strings.HasPrefix(mimeType, "audio/") {
			fileType = "audio"
		} else if strings.HasPrefix(mimeType, "video/") {
			fileType = "video"
		} else {
			ext := strings.ToLower(filepath.Ext(header.Filename))
			if ext == ".ttf" || ext == ".otf" || ext == ".woff" || ext == ".woff2" {
				fileType = "font"
			} else if ext == ".mp3" || ext == ".wav" || ext == ".m4a" || ext == ".ogg" {
				fileType = "audio"
			} else if ext == ".mp4" || ext == ".mov" || ext == ".webm" {
				fileType = "video"
			} else {
				fileType = "image"
			}
		}
	}

	asset, err := h.assetUsecase.Upload(c.Request.Context(), fileBytes, name, fileType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, asset)
}

func (h *AssetHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid asset ID"})
		return
	}

	var asset domain.Asset
	if err := c.ShouldBindJSON(&asset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.assetUsecase.Update(c.Request.Context(), id, &asset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, asset)
}

func (h *AssetHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid asset ID"})
		return
	}

	err = h.assetUsecase.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Asset deleted successfully"})
}
