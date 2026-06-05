package user

import "context"

type CreateUserInput struct {
	Name  string
	Email string
}

type UserUsecase interface {
	Create(ctx context.Context, input *CreateUserInput) (*User, error)
	FindByID(ctx context.Context, id string) (*User, error)
	FindAll(ctx context.Context) ([]User, error)
}
