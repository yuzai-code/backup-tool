package middleware

import "gorm.io/gorm"

type DB struct {
	db *gorm.DB
}
