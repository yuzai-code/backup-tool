package repository

import (
	"backup-tool/internal/model"

	"gorm.io/gorm"
)

// PathRepository 接口 定义了path数据库的相关操作
type PathRepository interface {
	GetAllDirName() ([]model.PathDTO, error)
	GetDirName(dirname string) (model.Path, error)
	SavePath(path *model.Path) error
	DeletePath(id int) error
	GetPathByID(id int) (model.PathDTO, error)
	UpdatePath(id int, path *model.PathDTO) error
}

// pathRepositoryImpl 是 PathRepository 接口的实现
type pathRepositoryImpl struct {
	db *gorm.DB
}

// NewPathRepository 创建path数据库实例
func NewPathRepository(db *gorm.DB) PathRepository {
	return &pathRepositoryImpl{db: db}
}

func (p *pathRepositoryImpl) UpdatePath(id int, path *model.PathDTO) error {
	var oldPath model.PathDTO
	// 更新数据
	err := p.db.Model(&model.Path{}).Where("id = ?", id).Updates(&path).Error
	if err != nil {
		return err
	}
	// 如果文件路径有变化，就更新
	if oldPath.FilePath != path.FilePath {
		// 更新文件路径
		err = p.db.Model(&model.Path{}).Where("id = ?", id).Update("file_path", path.FilePath).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// GetPathByID 获取备份文件配置的详情
func (p *pathRepositoryImpl) GetPathByID(id int) (model.PathDTO, error) {
	var path model.PathDTO
	err := p.db.Model(&model.Path{}).Where("id = ?", id).First(&path).Error
	if err != nil {
		return model.PathDTO{}, err
	}
	return path, nil
}

// 删除路径
func (p *pathRepositoryImpl) DeletePath(id int) error {
	return p.db.Where("id = ?", id).Delete(&model.Path{}).Error
}

// GetAllDirName 获取所有目录路径
func (p *pathRepositoryImpl) GetAllDirName() ([]model.PathDTO, error) {
	var paths []model.PathDTO
	err := p.db.Model(&model.Path{}).Select("id, dir_name, file_path, back_path").Find(&paths).Error
	if err != nil {
		return nil, err
	}
	return paths, nil
}

// GetDirName 查询配置备份文件的名称
func (p *pathRepositoryImpl) GetDirName(dirname string) (model.Path, error) {
	// 查询当前配置备份文件的名称
	var path model.Path
	err := p.db.Where("dirname = ?", dirname).First(&path).Error
	if err != nil {
		return model.Path{}, err
	}
	return path, nil
}

// SavePath 保存路径配置
func (p *pathRepositoryImpl) SavePath(path *model.Path) error {
	return p.db.Create(path).Error
}
