package path

import (
	"backup-tool/internal/model"
	"backup-tool/internal/repository"
	"fmt"
	"strings"
)

// PathService 定义了路径服务的接口
type PathService interface {
	GetAllDirNames() ([]model.PathDTO, error)
	GetDirName(dirname string) (model.Path, error)
	SavePath(dirName, filePath, backPath string) error
	DeletePath(id int) error
	GetPathByID(id int) (model.PathDTO, error)
	UpdatePath(id int, path model.PathDTO) error
}

// pathServiceImpl 是 PathService 接口的实现
type pathServiceImpl struct {
	pathRepo repository.PathRepository
}

// NewPathService 创建一个新的 PathService 实例
func NewPathService(pathRepo repository.PathRepository) PathService {
	return &pathServiceImpl{
		pathRepo: pathRepo,
	}
}

func (s *pathServiceImpl) UpdatePath(id int, path model.PathDTO) error {
	// 判断文件是否存在
	_, err := s.pathRepo.GetPathByID(id)
	if err != nil {
		return err
	}
	// 更新数据库
	err = s.pathRepo.UpdatePath(id, &path) // 使用赋值运算符来重新赋值 err
	if err != nil {
		return err
	}
	return nil
}

func (s *pathServiceImpl) GetPathByID(id int) (model.PathDTO, error) {
	PathDTO, err := s.pathRepo.GetPathByID(id)
	return PathDTO, err
}

// DeletePath 删除路径
func (s *pathServiceImpl) DeletePath(id int) error {
	return s.pathRepo.DeletePath(id)
}

// GetAllDirNames 获取所有目录名称
func (s *pathServiceImpl) GetAllDirNames() ([]model.PathDTO, error) {
	return s.pathRepo.GetAllDirName()
}

// GetDirName 获取指定目录名称
func (s *pathServiceImpl) GetDirName(dirname string) (model.Path, error) {
	path, err := s.pathRepo.GetDirName(dirname)
	return path, err
}

func (s *pathServiceImpl) SavePath(dirName, filePath, backPath string) error {
	// 判断文件是否存在
	_, err := s.pathRepo.GetDirName(dirName)
	if err == nil {
		return fmt.Errorf("目录名已存在: %s", dirName)
	}

	// 创建一个新的 model.Path 实例
	pathModel := &model.Path{
		DirName:  dirName,
		FilePath: filePath,
		BackPath: backPath,
	}

	// 保存到数据库
	err = s.pathRepo.SavePath(pathModel)
	if err != nil {
		// 检查是否是唯一性约束失败
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fmt.Errorf("目录名已存在: %s", dirName)
		}
		return err
	}

	return nil
}
