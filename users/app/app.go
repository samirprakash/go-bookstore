package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func Start() {
	mapRoutes()
	router.Run(":8080")
}
