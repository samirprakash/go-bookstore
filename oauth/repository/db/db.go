package db

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/samirprakash/go-bookstore/oauth/clients/cassandra"
	"github.com/samirprakash/go-bookstore/oauth/domain/token"
	"github.com/samirprakash/go-bookstore/oauth/utils/errors"
)

type Repository interface {
	GetByID(string) (*token.Token, *errors.REST)
	Create(token.Token) *errors.REST
	UpdateExpiration(token.Token) *errors.REST
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Create(t token.Token) *errors.REST {
	queryCreate := "INSERT INTO oauth.tokens (id, user_id, client_id, expires_at) VALUES (?, ?, ?, ?);"
	if err := cassandra.GetSession().Query(queryCreate, t.ID, t.UserID, t.ClientID, t.ExpiresAt).WithContext(context.Background()).Exec(); err != nil {
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (r *repository) GetByID(id string) (*token.Token, *errors.REST) {
	token := token.Token{}
	getTokenByIDQuery := "SELECT id, user_id, client_id, expires_at FROM oauth.tokens WHERE id = ?;"
	if err := cassandra.GetSession().Query(getTokenByIDQuery, id).WithContext(context.Background()).Scan(&token.ID, &token.UserID, &token.ClientID, &token.ExpiresAt); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("token not found in database")
		}
		return nil, errors.NewInternalServerError("database error")
	}

	return &token, nil
}

func (r *repository) UpdateExpiration(t token.Token) *errors.REST {
	queryUpdate := "UPDATE oauth.tokens SET expires_at = ? WHERE id = ?;"
	if err := cassandra.GetSession().Query(queryUpdate, t.ExpiresAt, t.ID).WithContext(context.Background()).Exec(); err != nil {
		return errors.NewInternalServerError("database error")
	}
	return nil
}
