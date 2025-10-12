package utils

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"

	"time"
)

type Token struct {
	UserID string
	Token string
	ExpiresAt time.Time
}

func generateToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func GenerateRefreshToken(userId int) (*Token) {
	refreshToken := generateToken()

	expiresAt := time.Now().Add(7 * 24 * time.Hour)

	data := &Token{
		UserID: strconv.Itoa(userId),
		Token: refreshToken,
		ExpiresAt: expiresAt,
	}

	return data
}

func GenerateVerificationCode() (*Token,error) {
	code := generateToken()

	data := &Token{
		Token: code,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	return data, nil
}