package v1handler

import (
	"myapp/lesson_04/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetUserByUUIDV1(ctx *gin.Context) {
	uuidStr := ctx.Param("uuid")
	uid, err := utils.ValidationUuid("UUID", uuidStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Get user by UUID",
		"user_uuid": uid,
	})
}

func (u *UserHandler) GetUsersV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "List all users v1"})
}

func (u *UserHandler) GetUserByIDV1(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := utils.ValidationPositiveInt("ID", idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by Id v1",
		"user_id": id,
	})
}

func (u *UserHandler) CreateUserV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"message": "Create users v1"})
}

func (u *UserHandler) UpdateUserV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Update users v1"})
}

func (u *UserHandler) DeleteUserV1(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
