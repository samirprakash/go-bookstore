package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExpirationTimeContstant(t *testing.T) {
	assert.Equal(t, time.Hour*24, expirationTime, "expiration time should be 24 hours")
}

func TestNewToken(t *testing.T) {
	token := New()
	assert.False(t, token.isExpired(), "token should not be expired")
	assert.NotEmpty(t, token.ID, "token should not be empty")
	assert.NotEmpty(t, token.UserID, "token should have a user id")
	assert.NotEmpty(t, token.ClientID, "token should have a client id")
	assert.NotEmpty(t, token.ExpiresAt, "token should have an expiration time")
}

func TestIsExpired(t *testing.T) {
	token := New()
	assert.False(t, token.isExpired(), "token should not be expired")
	token.ExpiresAt = time.Now().UTC().Add(-time.Hour).Unix()
	assert.True(t, token.isExpired(), "token should be expired")
}
