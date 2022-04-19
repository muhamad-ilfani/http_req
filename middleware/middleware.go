package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	USERNAME = "batman"
	PASSWORD = "secret"
)

func Auth(c *gin.Context) bool {
	username, password, ok := c.Request.BasicAuth()
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "unaouthorized",
		})
		return false
	}
	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "unaouthorized",
		})
		return false
	}
	return true
}
