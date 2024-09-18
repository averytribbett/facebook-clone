package handlers

import (
	"net/http"

	"fakebook.com/project/profile"
	"github.com/gin-gonic/gin"
)

// main calls handlers, handlers calls... the other things?
// GetUsersHandler returns all users
func GetUsersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, profile.Get())
}
