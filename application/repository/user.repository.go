package repository

import (
	"context"
	"fd-test/application/database"
	"fd-test/application/model"
	"fmt"
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

	CreateUser = `
		INSERT INTO USERS (
			email, first_name, last_name,
			avatar, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
	`

	GetMaxId = `
		SELECT MAX(id) FROM users;
	`

	UpdateOneUser = `
		UPDATE USERS set email = $1, first_name = $2,
		last_name = $3, avatar = $4, updated_at = $5, deleted_at = $6
		WHERE id = $7
	`
)

type UserRepository interface {
	FetchUser(ctx context.Context, data model.User) error
	FindUserById(ctx context.Context, id int) (model.User, error)
	FindAllUser(ctx context.Context) ([]model.User, error)
	CreateUser(ctx context.Context, data model.User) error
	GetMaxUser(ctx context.Context) (int, error)
	UpdateSequenceId(ctx context.Context, newValue int) error
	UpdateUserById(ctx context.Context, data model.User) error
}

type userRepo struct {
	db *database.DB
}

// UpdateUserById implements UserRepository
func (u userRepo) UpdateUserById(ctx context.Context, data model.User) error {
	stmt, err := u.db.Postgres.Prepare(UpdateOneUser)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		data.Email,
		data.FirstName,
		data.LastName,
		data.Avatar,
		data.UpdatedAt,
		data.DeletedAt,
		data.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

// UpdateSequenceId implements UserRepository
func (u userRepo) UpdateSequenceId(ctx context.Context, newValue int) error {
	stmt, err := u.db.Postgres.Prepare("ALTER SEQUENCE users_id_seq RESTART WITH " + fmt.Sprintf("%d", newValue+1))
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}

// GetMaxUser implements UserRepository
func (u userRepo) GetMaxUser(ctx context.Context) (int, error) {
	stmt, err := u.db.Postgres.Prepare(GetMaxId)
	if err != nil {
		return 0, err
	}

	row := stmt.QueryRow()

	maxId := 0
	err = row.Scan(&maxId)
	if err != nil {
		return 0, err
	}

	return maxId, nil
}

// CreateUser implements UserRepository
func (u userRepo) CreateUser(ctx context.Context, data model.User) error {
	stmt, err := u.db.Postgres.Prepare(CreateUser)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		data.Email, data.FirstName, data.LastName, data.Avatar, data.CreatedAt, data.UpdatedAt,
	)

	return err
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
	if err != nil {
		return err
	}

	maxId, err := u.GetMaxUser(ctx)
	if err != nil {
		return err
	}

	err = u.UpdateSequenceId(ctx, maxId)
	if err != nil {
		return err
	}

	return nil
}

func NewUserRepo(db *database.DB) UserRepository {
	return userRepo{
		db: db,
	}
}
