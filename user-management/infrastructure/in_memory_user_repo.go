package infrastructure

import (
	"errors"
	"sync"
	"time"

	"user-management/internal/domain"
)

type InMemoryUserRepo struct {
	users           map[int]domain.User
	authentications map[int]domain.Authentication
	mu              sync.Mutex
	nextID          int
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users:           make(map[int]domain.User),
		authentications: make(map[int]domain.Authentication),
		nextID:          1,
	}
}

func (repo *InMemoryUserRepo) SaveUser(user domain.User) domain.User {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	user.ID = repo.nextID
	user.CreatedAt = time.Now()
	repo.users[user.ID] = user
	repo.nextID++

	return user
}

func (repo *InMemoryUserRepo) FindUserByEmail(email string) (*domain.User, bool) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	for _, user := range repo.users {
		if user.Email == email {
			return &user, true
		}
	}
	return nil, false
}

func (repo *InMemoryUserRepo) SaveAuthentication(auth domain.Authentication) domain.Authentication {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	auth.ID = len(repo.authentications) + 1
	auth.LoginAt = time.Now()
	repo.authentications[auth.UserID] = auth

	return auth
}

func (repo *InMemoryUserRepo) FindAuthenticationByUserID(userID int) (*domain.Authentication, bool) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	auth, exists := repo.authentications[userID]
	if exists {
		return &auth, true
	}
	return nil, false
}

func (repo *InMemoryUserRepo) Logout(userID int) bool {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	auth, exists := repo.authentications[userID]
	if exists {
		auth.LogoutAt = time.Now()
		repo.authentications[userID] = auth
		return true
	}
	return false
}

func (r *InMemoryUserRepo) GetUserByID(userID int) (*domain.User, error) {
	for _, user := range r.users {
		if user.ID == userID {
			return &user, nil
		}
	}
	return nil, errors.New("User not found")
}
