package handler

import (
	"backup-tool/model"
	"github.com/gin-gonic/gin"
)

func PathConfig(c *gin.Context) {
	var newpath model.Path
	if err := c.ShouldBind(&newpath); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	// 判断名字是否重复
}
