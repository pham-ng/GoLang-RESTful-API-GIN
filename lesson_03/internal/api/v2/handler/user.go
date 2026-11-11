package v2handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetUsersV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "List all users v2"})
}

func (u *UserHandler) GetUserByIDV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Get user by Id v2"})
}

func (u *UserHandler) CreateUserV2(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"message": "Create users v2"})
}

func (u *UserHandler) UpdateUserV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Update users v2"})
}

func (u *UserHandler) DeleteUserV2(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
