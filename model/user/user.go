// package user /model/user
package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string `gorm:"column:name" json:"name"`
	Email      string `gorm:"column:email;uniqueIndex;not null" json:"email"`
	Password   string `gorm:"columnn:password;not null" json:"password"`
	BackupPath string `gorm:"column:backup_path;default:null" json:"backup_path"`
}
