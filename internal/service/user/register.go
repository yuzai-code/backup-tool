// package user service/user/register.go
package user

import (
	"backup-tool/internal/repository"
	"backup-tool/model"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user *model.User) error
}

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{userRepo: userRepo}
}

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
