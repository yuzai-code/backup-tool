package repository

import (
	"backup-tool/model"
	"gorm.io/gorm"
)

type PathRepository interface {
	GetDirName()
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func (p *userRepositoryImpl) GetDirName(dirname string) (string, error) {
	// 查询当前目录的名称
	var path model.Path
	err := p.db.Where("dirname = ?", dirname).First(&path).Error
	if err != nil {
		return "", err
	}
}
