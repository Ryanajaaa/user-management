package usecase

import (
	"errors"
	"time"
	"user-management/infrastructure"
	"user-management/internal/domain"

	"github.com/golang-jwt/jwt/v5"
)

type AuthUseCase struct {
	repo *infrastructure.InMemoryUserRepo
}

// Secret key untuk JWT
var jwtSecret = []byte("supersecretkey")

// Constructor untuk Authentication Use Case
func NewAuthUseCase(repo *infrastructure.InMemoryUserRepo) *AuthUseCase {
	return &AuthUseCase{repo: repo}
}

// Login user dan buat token JWT
func (uc *AuthUseCase) Login(email, password string) (string, error) {
	user, exists := uc.repo.FindUserByEmail(email)
	if !exists {
		return "", errors.New("user tidak ditemukan")
	}

	// (Untuk tugas ini, kita abaikan verifikasi password)
	// Buat token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 2).Unix(), // Expired dalam 2 jam
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	// Simpan autentikasi
	auth := domain.Authentication{
		UserID:  user.ID,
		Token:   tokenString,
		LoginAt: time.Now(),
	}
	uc.repo.SaveAuthentication(auth)

	return tokenString, nil
}

// Verifikasi token JWT
func (uc *AuthUseCase) VerifyToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}
	return nil, errors.New("token tidak valid")
}

// Logout user
func (uc *AuthUseCase) Logout(userID int) error {
	if !uc.repo.Logout(userID) {
		return errors.New("user tidak ditemukan atau sudah logout")
	}
	return nil
}
