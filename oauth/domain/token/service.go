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
	tid := strings.TrimSpace(id)
	if len(tid) == 0 {
		return nil, errors.NewBadRequestError("Invalid token id")
	}
	return s.repo.GetByID(tid)
}

func (s *service) Create(t Token) *errors.REST {
	if err := t.Validate(); err != nil {
		return err
	}
	return s.repo.Create(t)
}

func (s *service) UpdateExpiration(t Token) *errors.REST {
	if err := t.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateExpiration(t)
}
