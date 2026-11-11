package v1handler

import (
	"myapp/lesson_04/utils"
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
}

// Danh sách category hợp lệ (dùng slug: chữ thường, không dấu cách, không ký tự đặc biệt).
// "C/C++" và "C#" không dùng được trên URL ("/" tách đường dẫn, "#" là fragment),
// nên đổi thành "cpp" và "csharp".
var validCategory = map[string]bool{
	"php":        true,
	"python":     true,
	"golang":     true,
	"java":       true,
	"cpp":        true,
	"csharp":     true,
	"javascript": true,
	"rust":       true,
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

// categoryList trả về danh sách category đã sắp xếp, dùng cho thông báo lỗi.
func categoryList() string {
	list := make([]string, 0, len(validCategory))
	for name := range validCategory {
		list = append(list, name)
	}
	sort.Strings(list)
	return strings.Join(list, ", ")
}

func (c *CategoryHandler) GetCategoriesV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":    "List all categories v1",
		"categories": categoryList(),
	})
}

func (c *CategoryHandler) GetCategoryByCategoryV1(ctx *gin.Context) {
	category := ctx.Param("category")
	if err := utils.ValidationInList("Category", category, validCategory); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Category found",
		"category": category,
	})
}

func (c *CategoryHandler) GetCategoryBySlugV1(ctx *gin.Context) {
	category := strings.ToLower(ctx.Param("slug"))
	if !validCategory[category] {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Category must be one of: " + categoryList(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Category found",
		"category": category,
	})
}

func (c *CategoryHandler) PostCategoryV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"message": "Create category v1"})
}

func (c *CategoryHandler) PutCategoryV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Update category v1"})
}

func (c *CategoryHandler) DeleteCategoryV1(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
