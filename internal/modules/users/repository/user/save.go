package user

import (
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/users/domain/user"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/users/repository/model"
)

func (r *userRepository) Save(entity *user.User) error {
	row := model.UserFromEntity(entity)
	_, err := r.db.Exec(
		`INSERT INTO users (id, name, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`,
		row.ID, row.Name, row.Email, row.CreatedAt, row.UpdatedAt,
	)
	return err
}
