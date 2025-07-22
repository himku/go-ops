package model

import (
	"database/sql"
	"errors"

	"go-ops/internal/pkg"

	"gorm.io/gorm"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// GetUserByUsername 查询用户
func GetUserByUsername(username string) (*User, error) {
	db, err := pkg.GetDB()
	if err != nil {
		return nil, err
	}
	var user User
	tx := db.Where("username = ?", username).First(&user)
	if tx.Error != nil {
		if errors.Is(tx.Error, sql.ErrNoRows) || errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tx.Error
	}
	return &user, nil
}
