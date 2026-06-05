package user

import (
	"database/sql"

	"github.com/skyapps-id/monolithic-modular-go/internal/modules/users/domain/user"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/users/repository/model"
)

func (r *userRepository) FindByID(id string) (*user.User, error) {
	var row model.UserRow
	err := r.db.QueryRow(
		`SELECT id, name, email, created_at, updated_at FROM users WHERE id = ?`, id,
	).Scan(&row.ID, &row.Name, &row.Email, &row.CreatedAt, &row.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return row.ToEntity(), nil
}

var _ = (*sql.DB)(nil)
