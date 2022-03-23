package redirect

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Redirect(c *gin.Context) {
	c.Redirect(http.StatusFound, "https://google.com")
}