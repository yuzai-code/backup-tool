// api router/api/user.go
package api

import (
	"backup-tool/handler"
	"backup-tool/repository"
	"backup-tool/services/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserRoutes(r *gin.RouterGroup, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/register", userHandler.Register)
		// 添加其他用户相关路由...
	}
}