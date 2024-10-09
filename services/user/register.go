// package user service/user/register.go
package user

import (
	"backup-tool/model"
	"backup-tool/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// UserService
type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Register(user *model.User) error {
	// 检查用户是否存在
	existingUser, _ := s.userRepo.FindByUsername(user.Username)
	if existingUser != nil {
		return errors.New("username already exists")
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
