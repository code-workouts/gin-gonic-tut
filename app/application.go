package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapurls()
	router.Run("0.0.0.0:8080")
	// router.RunTLS("0.0.0.0:8080", "/path/to/certfile", "/path/to/keyfile")
}
