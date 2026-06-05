package user

import (
	"database/sql"

	"github.com/skyapps-id/monolithic-modular-go/internal/modules/users/domain/user"
)

var _ user.UserRepository = (*userRepository)(nil)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) user.UserRepository {
	return &userRepository{db: db}
}
