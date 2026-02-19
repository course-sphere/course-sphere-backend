package domain

import "time"

type Role string

const (
	Student    Role = "student"
	Instructor Role = "instructor"
	Admin      Role = "admin"
)

type User struct {
	Email            string    `json:"email"`
	Name             string    `json:"name"`
	Username         string    `json:"username"`
	DisplayUsername  string    `json:"displayUsername"`
	EmailVerified    bool      `json:"emailVerified"`
	Image            string    `json:"image"`
	Role             Role      `json:"role"`
	TwoFactorEnabled bool      `json:"twoFactorEnabled"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
