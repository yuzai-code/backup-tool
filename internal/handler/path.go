// package handler 定义了处理HTTP请求的处理器
package handler

import (
	"backup-tool/internal/model"
	"backup-tool/internal/service/path"
	"backup-tool/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

// UpdatePath 处理更新路径的HTTP请求
func (h *PathHandler) UpdatePath(c *gin.Context) {
	var path model.PathDTO
	pathID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的路径id"})
		return
	}
	// 绑定请求体到path结构体
	if err := c.ShouldBindJSON(&path); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 更新路径信息
	err = h.pathService.UpdatePath(pathID, path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "路径更新成功"})
}

// GetPathByID 处理根据路径ID获取路径信息的HTTP请求
func (h *PathHandler) GetPathByID(c *gin.Context) {
	var path model.PathDTO
	id := c.Param("id")
	pathID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的路径ID"})
		return
	}
	// 调用路径服务根据id获取路径信息
	path, err = h.pathService.GetPathByID(pathID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "路径不存在"})
		return
	}
	// 返回路径信息
	c.JSON(http.StatusOK, path)
}

// 删除路径
func (h *PathHandler) DeletePath(c *gin.Context) {
	// 获取要删除的路径ID
	id := c.Param("id")
	// 将id转换为整数
	pathID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的路径ID"})
		return
	}
	// 调用路径服务的删除路径方法
	err = h.pathService.DeletePath(pathID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "路径删除成功"})
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

	// 打印接收到的表单数据，用于调试
	utils.Logger.Debug("收到的表单数据",
		zap.String("dir_name", c.PostForm("dir_name")),
		zap.String("file_path", c.PostForm("file_path")),
		zap.String("back_path", c.PostForm("back_path")))

	// 使用 ShouldBind 来绑定表单数据
	if err := c.ShouldBind(&newPath); err != nil {
		utils.Logger.Error("无效的请求数据",
			zap.Error(err),
			zap.String("dir_name", c.PostForm("dir_name")))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的请求数据",
		})
		return
	}

	// 验证必要字段
	if newPath.DirName == "" {
		utils.Logger.Error("目录名称不能为空")
		c.JSON(http.StatusBadRequest, gin.H{"error": "目录名称不能为空"})
		return
	}

	// 保存新的路径配置
	err := h.pathService.SavePath(newPath.DirName, newPath.FilePath, newPath.BackPath)
	if err != nil {
		utils.Logger.Error("保存路径失败",
			zap.Error(err),
			zap.String("dir_name", newPath.DirName))
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("保存路径失败: %v", err)})
		return
	}

	utils.Logger.Info("路径配置成功", zap.String("dir_name", newPath.DirName))
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
	c.JSON(http.StatusOK, paths)
}
