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
			avatar, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7
		)
	`

	FindUserById = `
		SELECT * FROM USERS
		WHERE id = $1
	`

	FindAllUser = `
		SELECT * FROM USERS
	`
)

type UserRepository interface {
	FetchUser(ctx context.Context, data model.User) error
	FindUserById(ctx context.Context, id int) (model.User, error)
	FindAllUser(ctx context.Context) ([]model.User, error)
}

type userRepo struct {
	db *database.DB
}

// FindAllUser implements UserRepository
func (u userRepo) FindAllUser(ctx context.Context) ([]model.User, error) {
	stmt, err := u.db.Postgres.Prepare(FindAllUser)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var users = []model.User{}
	for rows.Next() {
		var user = model.User{}
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Avatar,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil

}

// FindUserById implements UserRepository
func (u userRepo) FindUserById(ctx context.Context, id int) (model.User, error) {
	stmt, err := u.db.Postgres.Prepare(FindUserById)
	if err != nil {
		return model.User{}, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	user := model.User{}

	err = row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Avatar,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// FetchUser implements UserRepository
func (u userRepo) FetchUser(ctx context.Context, data model.User) error {
	stmt, err := u.db.Postgres.Prepare(FetchUser)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		data.ID, data.Email, data.FirstName, data.LastName, data.Avatar, data.CreatedAt, data.UpdatedAt,
	)

	return err
}

func NewUserRepo(db *database.DB) UserRepository {
	return userRepo{
		db: db,
	}
}
