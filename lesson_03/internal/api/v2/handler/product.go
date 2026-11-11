package v2handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (p *ProductHandler) GetProductsV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List all Products v2"})
}

func (p *ProductHandler) GetProductByIdV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get product by ID v2"})
}

func (p *ProductHandler) PostProductV2(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Create product v2"})
}

func (p *ProductHandler) PutProductV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update product v2"})
}

func (p *ProductHandler) DeleteProductV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete product v2"})
}
