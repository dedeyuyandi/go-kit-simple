package user

import (
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/gofrs/uuid"
)

var RepoErr = errors.New("Unable to handle Repo Request")

type Repository interface {
	GetUser(id uuid.UUID) (*User, error)
	CreateUser(u CreateUserRequest) error
	DeleteUserByID(u uuid.UUID) error
}

type repo struct {
	db     *sql.DB
	logger log.Logger
}

var (
	createUser     = `INSERT INTO users (id, email, password) VALUES ($1, $2, $3)`
	getUserByID    = `SELECT * FROM users WHERE id=$1`
	deleteUserByID = `DELETE FROM users WHERE id=$1`
)

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreateUser(req CreateUserRequest) error {

	stmt, err := repo.db.Prepare(createUser)
	if err != nil {
		return errors.New("ErrStatement")
	}
	defer stmt.Close()
	uuid, _ := uuid.NewV4()
	_, err = stmt.Exec(uuid, req.Email, req.Password)
	if err != nil {
		return errors.New("ErrExecuteDatabaseStatement")
	}
	return nil
}

func (repo *repo) GetUser(id uuid.UUID) (*User, error) {
	// fmt.Println(getUserByID, "getUserByID")
	row := repo.db.QueryRow(getUserByID, id)
	var u User
	if err := row.Scan(&u.ID, &u.Email, &u.Password); err != nil {
		return nil, errors.New("ErrDataNotFound")
	}
	return &u, nil
}

func (repo *repo) DeleteUserByID(id uuid.UUID) error {
	resp, err := repo.db.Query(deleteUserByID, id)
	if err != nil {
		return err
	}
	defer resp.Close()
	return nil
}
