// package repository repository/user.go
// 用户存储库包
package repository

import (
	"backup-tool/model"

	"gorm.io/gorm"
)

// 用户存储库接口，定义了用户相关的操作
type UserRepository interface {
	Create(user *model.User) error
	FindByUsername(username string) (*model.User, error)
}

// 用户存储库实现结构体
type userRepositoryImpl struct {
	BaseRepository
}

// 创建新的用户存储库实例
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{BaseRepository: NewBaseRepository(db)}
}

// 创建一个新用户
func (r *userRepositoryImpl) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// 根据用户名查找用户
func (r *userRepositoryImpl) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
