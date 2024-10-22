package path

import (
	"backup-tool/internal/model"
	"backup-tool/internal/repository"
)

// PathService 定义了路径服务的接口
type PathService interface {
	GetAllDirNames() ([]model.PathDTO, error)
	GetDirName(dirname string) (model.Path, error)
	SavePath(dirName, filePath, backPath string) error
	DeletePath(id int) error
	GetPathByID(id int) (model.PathDTO, error)
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

// SavePath 保存路径配置
func (s *pathServiceImpl) SavePath(dirName, filePath, backPath string) error {
	// 创建一个新的 model.Path 实例
	pathModel := &model.Path{
		DirName:  dirName,
		BackPath: backPath,
	}
	return s.pathRepo.SavePath(pathModel)
}
