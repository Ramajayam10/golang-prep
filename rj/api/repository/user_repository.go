package repository

import (
	"database/sql"
	"api/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository {
		db: db,
	}
}

func (u *UserRepository) CreateUser(user *model.User) error {
	query := `
		INSERT INTO users(name, email, bio) VALUES($1, $2, $3)
		RETURNING id, created_at, updated_at, is_active
	`
	return u.db.QueryRow(query, user.Name, user.Email, user.Bio).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.IsActive,
	)
}

func (u *UserRepository) GetAllUsers() ([]model.User, error) {
	query := `
		SELECT id, name, email, bio, created_at, updated_at, is_active from users
		WHERE is_active = true
		ORDER BY created_at DESC
	`
	users := []model.User{}
	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user model.User
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Bio,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.IsActive,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

    if err := rows.Err(); err != nil {
        return nil, err
    }

	return users, nil
}

func (u *UserRepository) GetUserById(id string) (model.User, error) {
	query := `
		SELECT id, name, email, bio, created_at, updated_at, is_active from users
		where id = $1
	`
	var user model.User
	row := u.db.QueryRow(query, id)
	if err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Bio,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.IsActive,
	); err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserRepository) UpdateUser(id string, user *model.User) error {
	query := `
	UPDATE users SET name=$1, email=$2, bio=$3, updated_at=now()
	WHERE id = $4
	`
	result, err := u.db.Exec(query, user.Name, user.Email, user.Bio, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return sql.ErrNoRows
	}
	return nil
}

func (u *UserRepository) HardDeleteUser(id string) error {
	query := `
	DELETE FROM users WHERE id = $1
	`
	result, err := u.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return sql.ErrNoRows
	}
	return nil
}

func (u *UserRepository) SoftDeleteUser(id string) error {
	query := `
	UPDATE users SET is_active=false, updated_at=now()
	WHERE id = $1
	`
	result, err := u.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return sql.ErrNoRows
	}
	return nil
}