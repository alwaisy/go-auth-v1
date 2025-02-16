package auth

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Index(user *User) error {
	return r.db.Create(user).Error
}
