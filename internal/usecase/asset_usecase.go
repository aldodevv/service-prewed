package usecase

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"sort"
	"strings"
	"time"
	"service-wedding/internal/domain"
)

type AssetUsecase interface {
	GetAll(ctx context.Context) ([]domain.Asset, error)
	GetByID(ctx context.Context, id int64) (*domain.Asset, error)
	Create(ctx context.Context, asset *domain.Asset) error
	Update(ctx context.Context, id int64, asset *domain.Asset) error
	Delete(ctx context.Context, id int64) error
	Upload(ctx context.Context, fileBytes []byte, fileName string, fileType string) (*domain.Asset, error)
}

type assetUsecase struct {
	assetRepo   domain.AssetRepository
	cloudName   string
	apiKey      string
	apiSecret   string
}

func NewAssetUsecase(ar domain.AssetRepository, cloudName, apiKey, apiSecret string) AssetUsecase {
	return &assetUsecase{
		assetRepo:   ar,
		cloudName:   cloudName,
		apiKey:      apiKey,
		apiSecret:   apiSecret,
	}
}

func (u *assetUsecase) GetAll(ctx context.Context) ([]domain.Asset, error) {
	return u.assetRepo.GetAll(ctx)
}

func (u *assetUsecase) GetByID(ctx context.Context, id int64) (*domain.Asset, error) {
	return u.assetRepo.GetByID(ctx, id)
}

func (u *assetUsecase) Create(ctx context.Context, a *domain.Asset) error {
	if a.Name == "" || a.Type == "" || a.Url == "" {
		return errors.New("name, type, and url are required")
	}
	return u.assetRepo.Create(ctx, a)
}

func (u *assetUsecase) Update(ctx context.Context, id int64, a *domain.Asset) error {
	existing, err := u.assetRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("asset not found")
	}

	a.ID = id
	return u.assetRepo.Update(ctx, a)
}

func (u *assetUsecase) Delete(ctx context.Context, id int64) error {
	existing, err := u.assetRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("asset not found")
	}

	// 1. Delete from Cloudinary
	if existing.CloudinaryPublicID != "" {
		// Map type to Cloudinary resource type
		cldResourceType := "image"
		if existing.Type == "video" || existing.Type == "audio" {
			cldResourceType = "video"
		} else if existing.Type == "font" {
			cldResourceType = "raw"
		}

		err = u.destroyInCloudinary(ctx, existing.CloudinaryPublicID, cldResourceType)
		if err != nil {
			// Log error but proceed to delete DB record so users aren't stuck if Cloudinary file is already missing
			fmt.Printf("Warning: failed to delete from Cloudinary: %v\n", err)
		}
	}

	// 2. Delete from PostgreSQL
	return u.assetRepo.Delete(ctx, id)
}

func (u *assetUsecase) Upload(ctx context.Context, fileBytes []byte, fileName string, fileType string) (*domain.Asset, error) {
	if u.cloudName == "" || u.apiKey == "" || u.apiSecret == "" {
		return nil, errors.New("cloudinary credentials are not configured in usecase")
	}

	// 1. Map frontend types to Cloudinary resource types
	cldResourceType := "image"
	if fileType == "video" || fileType == "audio" {
		cldResourceType = "video"
	} else if fileType == "font" {
		cldResourceType = "raw"
	}

	// 2. Upload to Cloudinary
	secureURL, publicID, err := u.uploadToCloudinary(ctx, fileBytes, fileName, cldResourceType)
	if err != nil {
		return nil, fmt.Errorf("cloudinary upload error: %w", err)
	}

	// 3. Save to database
	asset := &domain.Asset{
		Name:               fileName,
		Type:               fileType,
		CloudinaryPublicID: publicID,
		Url:                secureURL,
	}

	err = u.assetRepo.Create(ctx, asset)
	if err != nil {
		return nil, fmt.Errorf("failed to save asset to database: %w", err)
	}

	return asset, nil
}

// Private helper to sign parameters for Cloudinary
func (u *assetUsecase) generateSignature(params map[string]string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var parts []string
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s=%s", k, params[k]))
	}
	sortedStr := strings.Join(parts, "&")

	signatureStr := sortedStr + u.apiSecret

	h := sha1.New()
	h.Write([]byte(signatureStr))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Private helper to perform the signed multipart upload
func (u *assetUsecase) uploadToCloudinary(ctx context.Context, fileBytes []byte, fileName string, resourceType string) (string, string, error) {
	cleanCloudName := strings.TrimPrefix(u.cloudName, "@")
	url := fmt.Sprintf("https://api.cloudinary.com/v1_1/%s/%s/upload", cleanCloudName, resourceType)

	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	folder := "wedding"

	params := map[string]string{
		"timestamp": timestamp,
		"folder":    folder,
	}

	signature := u.generateSignature(params)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return "", "", err
	}
	if _, err = part.Write(fileBytes); err != nil {
		return "", "", err
	}

	_ = writer.WriteField("api_key", u.apiKey)
	_ = writer.WriteField("timestamp", timestamp)
	_ = writer.WriteField("folder", folder)
	_ = writer.WriteField("signature", signature)

	err = writer.Close()
	if err != nil {
		return "", "", err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, body)
	if err != nil {
		return "", "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("cloudinary upload returned status %d: %s", resp.StatusCode, string(respBytes))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return "", "", err
	}

	secureURL, _ := result["secure_url"].(string)
	publicID, _ := result["public_id"].(string)

	if secureURL == "" || publicID == "" {
		return "", "", fmt.Errorf("cloudinary response missing secure_url or public_id: %s", string(respBytes))
	}

	return secureURL, publicID, nil
}

// Private helper to delete a file in Cloudinary
func (u *assetUsecase) destroyInCloudinary(ctx context.Context, publicID string, resourceType string) error {
	cleanCloudName := strings.TrimPrefix(u.cloudName, "@")
	url := fmt.Sprintf("https://api.cloudinary.com/v1_1/%s/%s/destroy", cleanCloudName, resourceType)

	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	params := map[string]string{
		"public_id": publicID,
		"timestamp": timestamp,
	}

	signature := u.generateSignature(params)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("public_id", publicID)
	_ = writer.WriteField("api_key", u.apiKey)
	_ = writer.WriteField("timestamp", timestamp)
	_ = writer.WriteField("signature", signature)

	_ = writer.Close()

	req, err := http.NewRequestWithContext(ctx, "POST", url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("cloudinary destroy returned status %d: %s", resp.StatusCode, string(respBytes))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return err
	}

	resultStr, _ := result["result"].(string)
	if resultStr != "ok" && resultStr != "not found" {
		return fmt.Errorf("cloudinary destroy failed: result is %s", resultStr)
	}

	return nil
}
