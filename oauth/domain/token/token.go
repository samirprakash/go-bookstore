package token

import "time"

const expirationTime = time.Hour * 24

type Token struct {
	ID        string `json:"id"`
	UserID    int64  `json:"user_id"`
	ClientID  int64  `json:"client_id"`
	ExpiresAt int64  `json:"expires_at"`
}

func NewToken() *Token {
	return &Token{
		ExpiresAt: time.Now().UTC().Add(expirationTime).Unix(),
	}
}

func (t Token) isExpired() bool {
	return time.Now().UTC().Unix() > t.ExpiresAt
}
