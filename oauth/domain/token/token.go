package token

import (
	"time"

	"github.com/samirprakash/go-bookstore/oauth/utils/errors"
)

const expirationTime = time.Hour * 24

type Token struct {
	ID        string `json:"id"`
	UserID    int64  `json:"user_id"`
	ClientID  int64  `json:"client_id"`
	ExpiresAt int64  `json:"expires_at"`
}

func New() *Token {
	return &Token{
		ExpiresAt: time.Now().UTC().Add(expirationTime).Unix(),
	}
}

func (t *Token) isExpired() bool {
	return time.Now().UTC().Unix() > t.ExpiresAt
}

func (t *Token) Validate() *errors.REST {
	if t.ID == "" {
		return errors.NewBadRequestError("Invalid token id")
	}
	if t.UserID == 0 {
		return errors.NewBadRequestError("Invalid user id")
	}
	if t.ClientID == 0 {
		return errors.NewBadRequestError("Invalid client id")
	}
	if t.ExpiresAt == 0 {
		return errors.NewBadRequestError("Invalid expiration time")
	}
	return nil
}
