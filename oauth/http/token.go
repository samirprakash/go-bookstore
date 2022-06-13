package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samirprakash/go-bookstore/oauth/domain/token"
)

type TokenHandler interface {
	GetById(*gin.Context)
}

type tokenHandler struct {
	service token.Service
}

func NewTokenHandler(s token.Service) TokenHandler {
	return &tokenHandler{
		service: s,
	}
}

func (h *tokenHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	t, err := h.service.GetById(id)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, t)
}
