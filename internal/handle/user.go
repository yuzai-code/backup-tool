// handler handler/user.go
package handler

import (
	"backup-tool/internal/model"
	"backup-tool/internal/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandle 处理与用户相关的请求
type UserHandler struct {
	userService user.UserService
}

// NewUserHandle 创建一个新的 UserHandle 实例
func NewUserHandler(userService user.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// Register 注册用户的处理函数
func (h *UserHandler) Register(ctx *gin.Context) {
	user := model.User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userService.Register(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}
