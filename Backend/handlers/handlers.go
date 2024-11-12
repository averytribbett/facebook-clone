package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

func EditNameHandler(c *gin.Context) {

	username := c.Param("username")
	newName := c.Param("newName")

	newNameSections := strings.Split(newName, " ")

	c.JSON(http.StatusOK, profile.EditName(username, newNameSections[0], newNameSections[1]))
}

func EditBioHandler(c *gin.Context) {

	username := c.Param("username")
	newBio := c.Param("newBio")

	c.JSON(http.StatusOK, profile.EditBio(username, newBio))
}

func EditUsernameHandler(c *gin.Context) {

	username := c.Param("username")
	newUsername := c.Param("newUsername")

	c.JSON(http.StatusOK, profile.EditUsername(username, newUsername))
}

func DeleteUserHandler(c *gin.Context) {

	username := c.Param("username")

	c.JSON(http.StatusOK, profile.DeleteUser(username))
}

func AddPostHandler(c *gin.Context) {
	var userId, err = strconv.Atoi(c.Param("userId"))
	if err != nil {
		fmt.Println(err)
	}
	postText := c.Param("postText")
	c.JSON(http.StatusOK, feed.AddPost(userId, postText))
}

func AddReplyHandler(c *gin.Context) {
	var reply models.Reply
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("i do not really care, I just want to graduate")
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	json.Unmarshal(body, &reply)

	c.JSON(http.StatusOK, feed.AddReply((reply)))
}

func GetAllRepliesHandler(c *gin.Context) {
	var postId int
	postId, err := strconv.Atoi(c.Param("postId"))
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, feed.GetReplies(postId))
}

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Limit the size of the upload (10MB in this case)
	err := r.ParseMultipartForm(10 << 20) // Max upload size: 10MB
	if err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	// Retrieve the file from the request form
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Ensure the "uploads" directory exists
	uploadDir := "uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		http.Error(w, "Unable to create upload directory", http.StatusInternalServerError)
		return
	}

	// Define the full path to save the uploaded file
	filePath := filepath.Join(uploadDir, handler.Filename)

	// Create the file on the server
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file data to the server's file
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	// Respond with success and the file path
	fmt.Fprintf(w, "File uploaded successfully: %s\n", filePath)
}
