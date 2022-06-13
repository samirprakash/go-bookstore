package token

import (
	"strings"

	"github.com/samirprakash/go-bookstore/oauth/utils/errors"
)

type Repository interface {
	GetByID(string) (*Token, *errors.REST)
	Create(Token) *errors.REST
	UpdateExpiration(Token) *errors.REST
}

type Service interface {
	GetByID(string) (*Token, *errors.REST)
	Create(Token) *errors.REST
	UpdateExpiration(Token) *errors.REST
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) GetByID(id string) (*Token, *errors.REST) {
	t, err := s.repo.GetByID(strings.TrimSpace(id))
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (s *service) Create(t Token) *errors.REST {
	err := s.repo.Create(t)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateExpiration(t Token) *errors.REST {
	err := s.repo.UpdateExpiration(t)
	if err != nil {
		return err
	}
	return nil
}
