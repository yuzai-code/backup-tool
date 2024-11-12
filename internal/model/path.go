// Package model
package model

import "gorm.io/gorm"

type Path struct {
	gorm.Model
	DirName  string `gorm:"type:varchar(255);not null;unique" json:"dir_name" bind:"required" form:"dir_name"` // 文件名
	FilePath string `gorm:"type:varchar(255);not null" json:"file_path" bind:"required" form:"file_path"`      // 文件路径
	BackPath string `gorm:"type:varchar(255);not null" json:"back_path" bind:"required" form:"back_path"`      // 备份路径
}

// 用于返回需要的路径信息
type PathDTO struct {
	DirName  string `json:"dir_name"`  // 16 bytes
	FilePath string `json:"file_path"` // 16 bytes
	BackPath string `json:"back_path"` // 16 bytes
	ID       uint   `json:"id"`        // 8 bytes
}
