// package user service/user/register.go
// 用户服务包，包含用户注册功能
package user

import (
	"backup-tool/internal/model"
	"backup-tool/internal/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// 用户服务接口，定义了用户相关操作
type UserService interface {
	Register(user *model.User) error
}

// 用户服务实现结构体，包含用户仓库
type UserServiceImpl struct {
	userRepo repository.UserRepository
}

// 创建新的用户服务
func NewUserService(userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{userRepo: userRepo}
}

// 注册新用户，进行验证和密码哈希处理
func (s *UserServiceImpl) Register(user *model.User) error {
	// 检查用户是否存在
	existingUser, _ := s.userRepo.FindByUsername(user.Username)
	if existingUser != nil {
		return errors.New("用户名已经存在")
	}

	// 哈希密码
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashPassword)
	// 创建用户
	return s.userRepo.Create(user)
}
