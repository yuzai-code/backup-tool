// Package model package model /model/user
package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string `gorm:"column:Username" json:"username"`
	Email      string `gorm:"column:email;uniqueIndex;not null" json:"email"`
	Password   string `gorm:"columnn:password;not null" json:"password"`
	BackupPath string `gorm:"column:backup_path;default:null" json:"backup_path"`
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
