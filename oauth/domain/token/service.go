package token

import (
	"strings"

	"github.com/samirprakash/go-bookstore/oauth/utils/errors"
)

type Repository interface {
	GetById(id string) (*Token, *errors.REST)
}

type Service interface {
	GetById(id string) (*Token, *errors.REST)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) GetById(id string) (*Token, *errors.REST) {
	t, err := s.repo.GetById(strings.TrimSpace(id))
	if err != nil {
		return nil, err
	}
	return t, nil
}
