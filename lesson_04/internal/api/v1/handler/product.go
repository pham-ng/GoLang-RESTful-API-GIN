package v1handler

import (
	"net/http"
	"time"

	"myapp/lesson_04/utils"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

type GetProductBySlugV1Param struct {
	Slug string `uri:"slug" binding:"slug,min=5,max=100"`
}

type GetProductsV1Param struct {
	Search string `form:"search" binding:"required,min=3,max=50,search"`
	Limit  int    `form:"limit" binding:"omitempty,gte=1,lte=100"`
	Email  string `form:"email" binding:"omitempty,email"`
	Date   string `form:"date" binding:"omitempty,datetime=2-1-2006"`
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (p *ProductHandler) GetProductsV1(ctx *gin.Context) {
	// search là bắt buộc: phải có, dài 3-50 ký tự, chỉ gồm chữ, số và dấu cách.
	var params GetProductsV1Param
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}
	if params.Limit == 0 {
		params.Limit = 1
	}
	if params.Email == "" {
		params.Email = "No email"
	}
	if params.Date == "" {
		params.Date = time.Now().Format("2-1-2006")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List all products(v1)",
		"search":  params.Search,
		"limit":   params.Limit,
		"email":   params.Email,
		"date":    params.Date,
	})
}

func (p *ProductHandler) GetProductBySlugV1(ctx *gin.Context) {
	var params GetProductBySlugV1Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get product by slug v1",
		"slug":    params.Slug,
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
