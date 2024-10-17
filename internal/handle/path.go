// package handler 定义了处理HTTP请求的处理器
package handler

import (
	"backup-tool/internal/service/path"
	"backup-tool/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PathHandler 结构体封装了路径相关的处理方法
type PathHandler struct {
	pathService path.PathService // pathService 用于处理路径相关的业务逻辑
}

// NewPathHandler 创建一个新的 PathHandler 实例
// 参数:
//   - pathService: 路径服务接口的实现
//
// 返回:
//   - *PathHandler: 新创建的 PathHandler 实例指针
func NewPathHandler(pathService path.PathService) *PathHandler {
	return &PathHandler{pathService: pathService}
}

// PathConfig 处理路径配置的HTTP请求
// 功能:
//   - 接收新的路径配置
//   - 检查路径名称是否重复
//   - 保存新的路径配置
//
// 参数:
//   - c *gin.Context: Gin的上下文对象，用于处理HTTP请求和响应
func (h *PathHandler) PathConfig(c *gin.Context) {
	var newPath model.Path
	// 绑定请求体到 newPath 结构体
	if err := c.ShouldBind(&newPath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 检查目录名称是否已存在
	existingPath, err := h.pathService.GetDirName(newPath.DirName)
	if err == nil && existingPath != "" {
		c.JSON(http.StatusConflict, gin.H{"error": "目录名称已存在"})
		return
	}

	// 保存新的路径配置
	err = h.pathService.SavePath(newPath.DirName, newPath.FilePath, newPath.BackPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存路径失败:" + err.Error()})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "路径配置成功", "path": newPath})
}

// GetAllPaths 处理获取所有路径的HTTP请求
// 功能:
//   - 获取所有已配置的路径
//   - 返回路径列表
//
// 参数:
//   - c *gin.Context: Gin的上下文对象，用于处理HTTP请求和响应
func (h *PathHandler) GetAllPaths(c *gin.Context) {
	// 获取所有路径
	paths, err := h.pathService.GetAllDirNames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取路径列表失败"})
		return
	}

	// 返回路径列表
	c.JSON(http.StatusOK, gin.H{"paths": paths})
}
