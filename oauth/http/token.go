package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samirprakash/go-bookstore/oauth/domain/token"
	"github.com/samirprakash/go-bookstore/oauth/utils/errors"
)

type TokenHandler interface {
	GetByID(*gin.Context)
	Create(*gin.Context)
}

type tokenHandler struct {
	service token.Service
}

func NewTokenHandler(s token.Service) TokenHandler {
	return &tokenHandler{
		service: s,
	}
}

func (h *tokenHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	t, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, t)
}

func (h *tokenHandler) Create(c *gin.Context) {
	var t token.Token
	if err := c.ShouldBindJSON(&t); err != nil {
		rErr := errors.NewBadRequestError("invalid json body")
		c.JSON(rErr.Status, rErr)
		return
	}

	if err := h.service.Create(t); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, t)
}

func (h *tokenHandler) UpdateExpiration(c *gin.Context) {
	var t token.Token
	if err := c.ShouldBindJSON(&t); err != nil {
		rErr := errors.NewBadRequestError("invalid json body")
		c.JSON(rErr.Status, rErr)
		return
	}

	if err := h.service.UpdateExpiration(t); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, t)
}
