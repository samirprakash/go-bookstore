package app

import (
	"github.com/gin-gonic/gin"
	"github.com/samirprakash/go-bookstore/oauth/domain/token"
	"github.com/samirprakash/go-bookstore/oauth/http"
	"github.com/samirprakash/go-bookstore/oauth/repository/db"
)

var router = gin.Default()

func Start() {
	s := token.NewService(db.NewRepository())
	h := http.NewTokenHandler(s)

	router.GET("/oauth/token/:id", h.GetByID)
	router.POST("/oauth/token", h.Create)

	router.Run(":8080")
}
