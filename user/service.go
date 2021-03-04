package user

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

type service struct {
	repostory Repository
	logger    log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repostory: rep,
		logger:    logger,
	}
}

type Service interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (string, error)
	GetUser(ctx context.Context, id uuid.UUID) (*User, error)
	DeleteUserByID(ctx context.Context, id uuid.UUID) (string, error)
}

func (s service) CreateUser(ctx context.Context, req CreateUserRequest) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")
	if err := s.repostory.CreateUser(req); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	return "Success created", nil
}

func (s service) GetUser(ctx context.Context, id uuid.UUID) (*User, error) {
	logger := log.With(s.logger, "method", "GetUser")
	fmt.Println(id, "service -> 3")
	resp, err := s.repostory.GetUser(id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}
	fmt.Println(resp, "response from repo -> 5")
	return &User{
		ID:       resp.ID,
		Email:    resp.Email,
		Password: resp.Password,
	}, nil
}

func (s service) DeleteUserByID(ctx context.Context, id uuid.UUID) (string, error) {
	logger := log.With(s.logger, "method", "DeleteUserByID")
	err := s.repostory.DeleteUserByID(id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	return "Success Deleted", nil
}
