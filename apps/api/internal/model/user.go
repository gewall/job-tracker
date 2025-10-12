package models

import (
	"time"

	"gorm.io/gorm"
)

// ======================
// Model Database
// ======================

type User struct {
	ID            uint               `gorm:"primaryKey" json:"id"`
	Name          string             `gorm:"size:100;not null" json:"name"`
	Email         string             `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Password      string             `gorm:"size:255;not null" json:"-"`
	IsVerified    bool               `gorm:"default:false" json:"is_verified"`
	RefreshTokens []RefreshToken     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
	Verifications []EmailVerification `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
	DeletedAt     gorm.DeletedAt     `gorm:"index" json:"-"`
}

type RefreshToken struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	Token     string    `gorm:"size:500;not null" json:"token"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

type EmailVerification struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	Code      string    `gorm:"size:500;not null" json:"code"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	IsUsed    bool      `gorm:"default:false" json:"is_used"`
	CreatedAt time.Time `json:"created_at"`
}
// ======================
// DTO (Data Transfer Object) dengan validator
// ======================

// Register DTO
type UserRegisterDTO struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// Signin DTO
type UserSigninDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Email verification DTO
type VerifyEmailDTO struct {
	ExpiresAt time.Time `json:"expires_at" validate:"required"`
	IsUsed bool `json:"is_used" validate:"required"`
	Code  string `json:"code" validate:"required,len=6"`
}

// Response DTO user
type UserResponseDTO struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	IsVerified bool   `json:"is_verified"`
	Verifications []VerifyEmailDTO `json:"verifications"`
}

// Response DTO token
type TokenResponseDTO struct {
	AccessToken  string    `json:"access_token"`  // JWT
	RefreshToken string    `json:"refresh_token"` // disimpan di DB
	ExpiresAt    time.Time `json:"expires_at"`
}


