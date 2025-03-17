package domain

import "time"

type Authentication struct {
	ID       int       `json:"id"`
	UserID   int       `json:"user_id"`
	Password string    `json:"password"`
	Token    string    `json:"token"`
	LoginAt  time.Time `json:"login_at"`
	LogoutAt time.Time `json:"logout_at"`
}
