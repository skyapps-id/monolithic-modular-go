package model

import (
	"time"

	"github.com/skyapps-id/monolithic-modular-go/internal/modules/users/domain/user"
)

type CreateUserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r *CreateUserRequest) ToInput() *user.CreateUserInput {
	return &user.CreateUserInput{
		Name:  r.Name,
		Email: r.Email,
	}
}

func UserToResponse(u *user.User) *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func UsersToResponse(users []user.User) []UserResponse {
	res := make([]UserResponse, len(users))
	for i, u := range users {
		res[i] = *UserToResponse(&u)
	}
	return res
}
