package user

import (
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/users/domain/user"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/users/repository/model"
)

func (r *userRepository) FindAll() ([]user.User, error) {
	rows, err := r.db.Query(`SELECT id, name, email, created_at, updated_at FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var row model.UserRow
		if err := rows.Scan(&row.ID, &row.Name, &row.Email, &row.CreatedAt, &row.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, *row.ToEntity())
	}
	return users, nil
}
