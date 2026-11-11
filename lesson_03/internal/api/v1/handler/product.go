package v1handler

import (
	"net/http"
	"regexp"
	"strings"

	"myapp/lesson_03/utils"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

var (
	slugRegex   = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
	searchRegex = regexp.MustCompile(`^[a-zA-Z0-9 ]+$`)
)

const maxLimit = 100

func (p *ProductHandler) GetProductsV1(ctx *gin.Context) {
	// search là bắt buộc: phải có, dài 3-50 ký tự, chỉ gồm chữ, số và dấu cách.
	search := strings.TrimSpace(ctx.Query("search"))
	if err := utils.ValidationRequired("Search", search); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := utils.ValidationStringLength("Search", search, 3, 50); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := utils.ValidationRegex("Search", search, searchRegex, "must contain only letters, numbers and spaces"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	limit, err := utils.ValidationPositiveInt("Limit", ctx.DefaultQuery("limit", "10"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if limit > maxLimit {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Limit must not exceed 100"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List all Products v1",
		"search":  search,
		"limit":   limit,
	})
}

func (p *ProductHandler) GetProductBySlugV1(ctx *gin.Context) {
	slug := ctx.Param("slug")

	if err := utils.ValidationRegex("Slug", slug, slugRegex, "must contain only lowercase letters, numbers, hyphens and dots"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get product by slug v1",
		"slug":    slug,
	})
}

func (p *ProductHandler) PostProductV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Create product v1"})
}

func (p *ProductHandler) PutProductV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update product v1"})
}

func (p *ProductHandler) DeleteProductV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete product v1"})
}
