package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"fakebook.com/project/feed"
	"fakebook.com/project/models"
	"fakebook.com/project/profile"
	"github.com/gin-gonic/gin"
)

// main calls handlers, handlers calls... the other things?
// GetUsersHandler returns all users
func GetUsersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, profile.Get())
}

func GetOneUserHandler(c *gin.Context) {
	var1, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, http.StatusOK) // idk how to format this rn
	}
	c.JSON(http.StatusOK, profile.GetOneUser(var1))
}

func GetOneUserbyUsernameHandler(c *gin.Context) {
	var username = c.Param("username")
	c.JSON(http.StatusOK, profile.GetOneUserByUsername(username))
}

func AddNewUserHandler(c *gin.Context) {
	var newUser models.User
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		fmt.Println("we will deal with you later")
	}
	c.JSON(http.StatusOK, profile.AddNewUser(newUser))
}

func GetInitialFeedByTime(c *gin.Context) {
	var numOfPosts, err = strconv.Atoi(c.Param("numOfPosts"))
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, feed.InitialFeedByTime(numOfPosts))
}
