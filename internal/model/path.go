// Package model
package model

import "gorm.io/gorm"

type Path struct {
	gorm.Model
	DirName  string `gorm:"type:varchar(255);not null;unique" json:"dir_name"` // 文件名
	FilePath string `gorm:"type:varchar(255);not null" json:"file_path"`       // 文件路径
	BackPath string `gorm:"type:varchar(255);not null" json:"back_path"`       // 备份路径
}

// 用于返回需要的路径信息
type PathDTO struct {
	ID       uint   `json:"id"`
	DirName  string `json:"dir_name"`
	FilePath string `json:"file_path"`
	BackPath string `json:"back_path"`
}
