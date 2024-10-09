// package repository repository/user_repository.go
package repository

import (
	"backup-tool/model"

	"gorm.io/gorm"
)

// userRepositoryImpl 实现UserRepository接口的具体结构体
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository 创建UserRepository实例
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create 创建用户
func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// FindByUsername 根据用户名查找用户
func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
