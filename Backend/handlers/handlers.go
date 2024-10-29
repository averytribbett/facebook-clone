package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"fakebook.com/project/feed"
	"fakebook.com/project/friends"
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

func GetUserPostsHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, feed.GetUserPosts(userID))
}

func GetInitialFeedByTimeHandler(c *gin.Context) {
	var numOfPosts, err = strconv.Atoi(c.Param("numOfPosts"))
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, feed.InitialFeedByTime(numOfPosts))
}

func GetFeedByTimeHandler(c *gin.Context) {
	var numOfPosts, err = strconv.Atoi(c.Param("numOfPosts"))
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, feed.FeedByTime(numOfPosts))
}

func FindUserByNameHandler(c *gin.Context) {
	var firstName string
	var lastName string

	fullName := c.Param("fullName")
	spaceIndex := strings.Index(fullName, " ")
	if fullName == "" {
		fmt.Println("not sure what to put here")
	} else if spaceIndex == -1 {
		firstName = fullName
		lastName = ""
	} else {
		firstName = fullName[0:spaceIndex]
		lastName = fullName[spaceIndex+1:]
	}

	c.JSON(http.StatusOK, profile.FindUserByName(firstName, lastName))
}

func FindUserByFullNameHandler(c *gin.Context) {
	firstName := c.Param("firstName")
	lastName := c.Param("lastName")
	c.JSON(http.StatusOK, profile.FindUserByFullName(firstName, lastName))
}

func GetFriendsListHandler(c *gin.Context) {
	username := c.Param("username")
	c.JSON(http.StatusOK, friends.GetFriendsList(username))
}

func GetFriendRequestListHandler(c *gin.Context) {
	username := c.Param("username")
	c.JSON(http.StatusOK, friends.GetFriendRequestList(username))
}

func AddOneFriendHandler(c *gin.Context) {
	requestor := c.Param("requestor")
	requestee := c.Param("requestee")
	c.JSON(http.StatusOK, friends.AddPendingFriend(requestor, requestee))
}

func AcceptFriendshipHandler(c *gin.Context) {
	originalRequestor := c.Param("originalRequestor")
	acceptee := c.Param("acceptee")
	c.JSON(http.StatusOK, friends.AcceptFriend(originalRequestor, acceptee))
}

func DeleteFriendshipRequestHandler(c *gin.Context) {
	originalRequestor := c.Param("originalRequestor")
	deleter := c.Param("deleter")
	c.JSON(http.StatusOK, friends.DeleteFriendRequest(originalRequestor, deleter))
}

func DeleteFriendshipHandler(c *gin.Context) {
	friendToDelete := c.Param("friendToDelete")
	deleter := c.Param("deleter")
	c.JSON(http.StatusOK, friends.DeleteFriend(friendToDelete, deleter))
}
