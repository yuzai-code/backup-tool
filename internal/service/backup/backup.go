package backup

import (
	"archive/zip"
	"backup-tool/internal/repository"
	"backup-tool/utils"
	"fmt"
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
}

// NewPathRepository 创建一个新的 PathRepositoryImpl 实例
func NewPathRepository(pathRepo repository.PathRepository) *BackupServiceImpl {
	return &BackupServiceImpl{
		pathRepo: pathRepo,
	}
}

// BackupService 处理备份服务的逻辑
func (p *BackupServiceImpl) BackupService(id int) error {
	// 根据id获取需要备份的文件的信息
	path, err := p.pathRepo.GetPathByID(id)
	if err != nil {
		// 处理错误
		utils.Logger.Error("获取文件信息失败", zap.Error(err))
		return err
	}

	// 获取备份文件存储路径
	savePath := path.BackPath

	// 如果目录不存在就创建
	if err := os.MkdirAll(savePath, os.ModePerm); err != nil {
		utils.Logger.Error("创建目录失败", zap.Error(err))
		return err
	}

	// 备份文件到指定目录
	srcFilePath := path.FilePath
	dstFilePath := filepath.Join(savePath, filepath.Base(srcFilePath))

	// 检查文件是否已经是压缩包
	if isCompressedFile(srcFilePath) {
		utils.Logger.Info("文件路径", zap.String("srcFilePath", srcFilePath))
		// 直接复制文件
		if err := copyFile(srcFilePath, dstFilePath); err != nil {
			utils.Logger.Error("复制文件失败", zap.Error(err))
			return err
		}
	} else {
		// 压缩文件
		if err := compressFile(srcFilePath, dstFilePath); err != nil {
			utils.Logger.Error("压缩文件失败", zap.Error(err))
			return err
		}
	}

	utils.Logger.Info("文件备份成功", zap.String("srcFilePath", srcFilePath), zap.String("dstFilePath", dstFilePath))
	return nil
}

// isCompressedFile 检查文件是否是压缩文件
func isCompressedFile(filePath string) bool {
	// 转换为小写以进行不区分大小写的比较
	lowerPath := strings.ToLower(filePath)

	// 定义常见压缩文件的扩展名
	compressedExtensions := []string{
		".zip", // ZIP 压缩文件
		".rar", // RAR 压缩文件
		".7z",  // 7-Zip 压缩文件
		".gz",  // Gzip 压缩文件
		".bz2", // Bzip2 压缩文件
		".tar", // Tar 打包文件
		".tgz", // Tar Gzip 压缩文件
		".xz",  // XZ 压缩文件
		".iso", // ISO 镜像文件
		".dmg", // macOS 磁盘镜像文件
		".img", // 磁盘镜像文件
		".pkg", // 软件包文件
	}

	// 检查文件扩展名是否匹配任意一个压缩文件扩展名
	for _, ext := range compressedExtensions {
		if strings.HasSuffix(lowerPath, ext) {
			return true
		}
	}

	return false
}

// copyFile 复制文件
func copyFile(srcFilePath, dstFilePath string) error {
	// 首先检查源文件状态
	srcInfo, err := os.Stat(srcFilePath)
	if err != nil {
		utils.Logger.Error("获取源文件信息失败", zap.Error(err))
		return fmt.Errorf("获取源文件信息失败: %w", err)
	}

	// 检查是否为目录
	if srcInfo.IsDir() {
		errMsg := fmt.Sprintf("源路径 %s 是一个目录，不是文件", srcFilePath)
		utils.Logger.Error(errMsg)
		return fmt.Errorf(errMsg)
	}

	// 记录源文件信息
	utils.Logger.Info("开始复制文件",
		zap.String("srcFilePath", srcFilePath),
		zap.Int64("fileSize", srcInfo.Size()),
		zap.String("fileMode", srcInfo.Mode().String()),
	)

	// 打开源文件
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		utils.Logger.Error("打开源文件失败", zap.Error(err))
		return fmt.Errorf("打开源文件失败: %w", err)
	}
	defer srcFile.Close()

	// 创建目标文件
	utils.Logger.Info("准备创建目标文件", zap.String("dstFilePath", dstFilePath))
	dstFile, err := os.Create(dstFilePath)
	if err != nil {
		utils.Logger.Error("创建目标文件失败", zap.Error(err))
		return fmt.Errorf("创建目标文件失败: %w", err)
	}
	defer dstFile.Close()

	// 复制文件内容
	written, err := io.Copy(dstFile, srcFile)
	if err != nil {
		utils.Logger.Error("复制文件内容失败", zap.Error(err))
		return fmt.Errorf("复制文件内容失败: %w", err)
	}

	utils.Logger.Info("文件复制完成",
		zap.String("srcFilePath", srcFilePath),
		zap.String("dstFilePath", dstFilePath),
		zap.Int64("copiedBytes", written),
	)

	return nil
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
