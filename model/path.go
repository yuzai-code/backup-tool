// Package model
package model

import "gorm.io/gorm"

type Path struct {
	gorm.Model
	DirName  string `gorm:"type:varchar(255);not null" json:"dirName"`  // 文件名
	FilePath string `gorm:"type:varchar(255);not null" json:"filePath"` // 文件路径
	BackPath string `gorm:"type:varchar(255);not null" json:"backPath"` // 备份路径
}
