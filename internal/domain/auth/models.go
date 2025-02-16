package auth

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Username     string    `gorm:"type:varchar(255);not null;unique" json:"username"`
	Email        string    `gorm:"type:varchar(255);not null;unique" json:"email"`
	FirstName    string    `gorm:"type:varchar(255)" json:"first_name,omitempty"`
	LastName     string    `gorm:"type:varchar(255)" json:"last_name,omitempty"`
	Role         string    `gorm:"type:varchar(255);not null;default:'user'" json:"role"`
	PasswordHash string    `gorm:"type:varchar(255);not null" json:"-"`
	IsVerified   bool      `gorm:"default:false" json:"is_verified"`
	CreatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp" json:"updated_at"`
}

// BeforeCreate hook to set UUID and timestamps
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate hook to update UpdatedAt timestamp
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return nil
}
