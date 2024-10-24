package backup

import (
	"archive/zip"
	"backup-tool/internal/repository"
	"io"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
)

type BackupService interface {
	BackupService(id int) error
}

// BackupServiceImpl
type BackupServiceImpl struct {
	pathRepo repository.PathRepository
	logger   *zap.Logger
}

// NewPathRepository 创建一个新的 PathRepositoryImpl 实例
func NewPathRepository(pathRepo repository.PathRepository) *BackupServiceImpl {
	return &BackupServiceImpl{pathRepo: pathRepo}
}

// BackupService 处理备份服务的逻辑
func (p *BackupServiceImpl) BackupService(id int) error {
	// 根据id获取需要备份的文件的信息
	path, err := p.pathRepo.GetPathByID(id)
	if err != nil {
		// 处理错误
		p.logger.Error("获取文件信息失败", zap.Error(err))
		return err
	}

	// 获取备份文件存储路径
	savePath := path.BackPath

	// 如果目录不存在就创建
	if err := os.MkdirAll(savePath, os.ModePerm); err != nil {
		p.logger.Error("创建目录失败", zap.Error(err))
		return err
	}

	// 备份文件到指定目录
	srcFilePath := path.FilePath
	dstFilePath := filepath.Join(savePath, filepath.Base(srcFilePath))

	// 检查文件是否已经是压缩包
	if isZipFile(srcFilePath) {
		// 直接复制文件
		if err := copyFile(srcFilePath, dstFilePath); err != nil {
			p.logger.Error("复制文件失败", zap.Error(err))
			return err
		}
	} else {
		// 压缩文件
		if err := compressFile(srcFilePath, dstFilePath); err != nil {
			p.logger.Error("压缩文件失败", zap.Error(err))
			return err
		}
	}

	p.logger.Info("文件备份成功", zap.String("srcFilePath", srcFilePath), zap.String("dstFilePath", dstFilePath))
	return nil
}

// isZipFile 检查文件是否是压缩包
func isZipFile(filePath string) bool {
	return strings.HasSuffix(filePath, ".zip")
}

// copyFile 复制文件
func copyFile(srcFilePath, dstFilePath string) error {
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dstFilePath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

// compressFile 压缩文件
func compressFile(srcFilePath, dstFilePath string) error {
	zipFile, err := os.Create(dstFilePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	info, err := srcFile.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = filepath.Base(srcFilePath)
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, srcFile)
	return err
}
