package accountMgr

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/google/uuid"
)

type service struct {
	logger log.Logger
}

type Service interface {
	CreateUser(ctx context.Context, email, password string) (*User, error)
	GetUser(ctx context.Context, id string) (*User, error)
}

func NewService(logger log.Logger) Service {
	return &service{logger:logger}
}

func (s *service) CreateUser(ctx context.Context, email, password string) (*User, error) {
	return &User{
		ID:       uuid.New().String(),
		Email:    email,
		Password: password,
	}, nil
}

func (s *service) GetUser(ctx context.Context, id string) (*User, error) {
	return &User{}, nil
}
