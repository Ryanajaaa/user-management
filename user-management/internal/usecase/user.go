package usecase

import (
	"errors"
	"user-management/infrastructure"
	"user-management/internal/domain"
)

type UserUseCase struct {
	repo *infrastructure.InMemoryUserRepo
}

// Constructor untuk Use Case
func NewUserUseCase(repo *infrastructure.InMemoryUserRepo) *UserUseCase {
	return &UserUseCase{repo: repo}
}

// Registrasi user baru
func (uc *UserUseCase) RegisterUser(name, email string) (domain.User, error) {
	// Cek apakah email sudah terdaftar
	if _, exists := uc.repo.FindUserByEmail(email); exists {
		return domain.User{}, errors.New("email sudah digunakan")
	}

	// Simpan user baru
	user := domain.User{
		Name:  name,
		Email: email,
	}
	return uc.repo.SaveUser(user), nil
}

// Ambil data user berdasarkan email
func (uc *UserUseCase) GetUserByID(userID int) (*domain.User, error) {
	return uc.repo.GetUserByID(userID)
}
