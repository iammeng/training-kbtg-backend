package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	Email         string         `gorm:"uniqueIndex;not null" json:"email"`
	Password      string         `gorm:"not null" json:"-"`
	FirstName     string         `json:"first_name"`
	LastName      string         `json:"last_name"`
	Phone         string         `json:"phone"`
	MembershipID  string         `gorm:"uniqueIndex" json:"membership_id"`
	MemberLevel   string         `gorm:"default:Gold" json:"member_level"`
	Points        int            `gorm:"default:0" json:"points"`
}

type RegisterRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Phone     string `json:"phone"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateProfileRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type ProfileResponse struct {
	User User `json:"user"`
}
