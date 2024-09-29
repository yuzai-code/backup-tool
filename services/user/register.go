package user

import "github.com/gin-gonic/gin"

// UserRegisterService 管理注册用户服务
type UserRegisterService struct {
	UserName string `form:"username" json:"username" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required，min=4,max=16"`
}

// 注册用户
func (service *UserRegisterService) Register(c *gin.Context) serializer.Response {
}
