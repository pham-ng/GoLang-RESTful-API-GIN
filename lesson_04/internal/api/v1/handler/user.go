package v1handler

import (
	"myapp/lesson_04/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

type GetUserByIDV1Param struct {
	ID int `uri:"id" binding:"gt=0"`
}

type GetUserByUUIDV1Param struct {
	UUID string `uri:"uuid" binding:"uuid"`
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetUserByUUIDV1(ctx *gin.Context) {
	var params GetUserByUUIDV1Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Get user by UUID",
		"user_uuid": params.UUID,
	})
}

func (u *UserHandler) GetUsersV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "List all users v1"})
}

func (u *UserHandler) GetUserByIDV1(ctx *gin.Context) {
	var params GetUserByIDV1Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by Id v1",
		"user_id": params.ID,
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
