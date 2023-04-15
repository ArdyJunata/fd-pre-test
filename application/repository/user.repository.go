package repository

import (
	"context"
	"fd-test/application/database"
	"fd-test/application/model"
)

var (
	FetchUser = `
		INSERT INTO USERS (
			id, email, first_name, last_name,
			created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
	`

	FindUserById = `
		SELECT * FROM USERS
		WHERE id = $1
	`
)

type UserRepository interface {
	FetchUser(ctx context.Context, data model.User) error
}

type userRepo struct {
	db *database.DB
}

// FetchUser implements UserRepository
func (u userRepo) FetchUser(ctx context.Context, data model.User) error {
	stmt, err := u.db.Postgres.Prepare(FetchUser)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		data.ID, data.Email, data.FirstName, data.LastName, data.CreatedAt, data.UpdatedAt,
	)

	return err
}

func NewUserRepo(db *database.DB) UserRepository {
	return userRepo{
		db: db,
	}
}
