package repository

import (
	"backup-tool/internal/model"

	"gorm.io/gorm"
)

// PathRepository 接口 定义了path数据库的相关操作
type PathRepository interface {
	GetAllDirName() ([]model.PathDTO, error)
	GetDirName(dirname string) (string, error)
	SavePath(path *model.Path) error
}

// pathRepositoryImpl 是 PathRepository 接口的实现
type pathRepositoryImpl struct {
	BaseRepository
}

// 在 model 包中添加一个新的 DTO 结构体
type PathDTO struct {
	DirName  string `json:"dirName"`
	FilePath string `json:"filePath"`
	BackPath string `json:"backPath"`
}

// 创建path数据库实例
func NewPathRepository(db *gorm.DB) PathRepository {
	return &pathRepositoryImpl{BaseRepository: NewBaseRepository(db)}
}

// GetAllDirName 获取所有目录路径
func (p *pathRepositoryImpl) GetAllDirName() ([]model.PathDTO, error) {
	var paths []model.PathDTO
	err := p.db.Model(&model.Path{}).Select("dir_name, file_path, back_path").Find(&paths).Error
	if err != nil {
		return nil, err
	}
	return paths, nil
}

// GetDirName 查询当前目录的名称
func (p *pathRepositoryImpl) GetDirName(dirname string) (string, error) {
	// 查询当前目录的名称
	var path model.Path
	err := p.db.Where("dirname = ?", dirname).First(&path).Error
	if err != nil {
		return "", err
	}
	return path.DirName, nil
}

// SavePath 保存路径配置
func (p *pathRepositoryImpl) SavePath(path *model.Path) error {
	return p.db.Create(path).Error
}
