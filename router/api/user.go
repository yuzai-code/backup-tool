// api router/api/user.go
package api

import (
	handler "backup-tool/internal/handle"
	repository "backup-tool/internal/repository"
	service "backup-tool/internal/service/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserRoutes(r *gin.RouterGroup, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/register", userHandler.Register)
		// 添加其他用户相关路由...
	}
}
