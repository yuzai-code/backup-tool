// service/backup.go
package services

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// 处理备份请求
func HandleBackup(c *gin.Context) {
	// 获取需要备份文件的路径参数
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "无法解析参数"})
		return
	}

	files, ok := form.File["files"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "请选择文件或文件夹"})
	}

	// 文件备份的路径
	filePath := "D:\\Project\\recode\\go\\backup-tool\\backupdir"
	// 如果文件夹不存在，则创建
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "无法创建备份目录"})
		return
	}

	// 备份的文件
	for _, file := range files {
		dst := filepath.Join(filePath, file.Filename)
		// 保存文件
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "备份失败", "error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "备份成功"})

}
