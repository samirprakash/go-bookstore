package db

import (
	"github.com/samirprakash/go-bookstore/oauth/domain/token"
	"github.com/samirprakash/go-bookstore/oauth/utils/errors"
)

type Repository interface {
	GetById(id string) (*token.Token, *errors.REST)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetById(id string) (*token.Token, *errors.REST) {
	return nil, nil
}
