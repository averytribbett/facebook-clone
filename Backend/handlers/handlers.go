package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"fakebook.com/project/feed"
	"fakebook.com/project/friends"
	"fakebook.com/project/models"
	"fakebook.com/project/profile"

	"fakebook.com/project/reactions"
	"github.com/gin-gonic/gin"
)

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
	var loggedInUserId, err2 = strconv.Atoi(c.Param("loggedInUserId"))
	if err2 != nil {
		fmt.Println(err2)
	}
	c.JSON(http.StatusOK, feed.GetUserPosts(userID, loggedInUserId))
}

func GetInitialFeedByTimeHandler(c *gin.Context) {
	var numOfPosts, err = strconv.Atoi(c.Param("numOfPosts"))
	if err != nil {
		fmt.Println(err)
	}
	var loggedInUserId, err2 = strconv.Atoi(c.Param("loggedInUserId"))
	if err2 != nil {
		fmt.Println(err2)
	}
	c.JSON(http.StatusOK, feed.InitialFeedByTime(numOfPosts, loggedInUserId))
}

func GetFeedByTimeHandler(c *gin.Context) {
	var numOfPosts, err = strconv.Atoi(c.Param("numOfPosts"))
	if err != nil {
		fmt.Println(err)
	}
	var loggedInUserId, err2 = strconv.Atoi(c.Param("loggedInUserId"))
	if err2 != nil {
		fmt.Println(err2)
	}
	c.JSON(http.StatusOK, feed.FeedByTime(numOfPosts, loggedInUserId))
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

func EditFirstNameHandler(c *gin.Context) {
	username := c.Param("username")
	newFirstName := c.Param("newFirstName")

	c.JSON(http.StatusOK, profile.EditFirstName(username, newFirstName))
}

func EditLastNameHandler(c *gin.Context) {
	username := c.Param("username")
	newLastName := c.Param("newLastName")

	c.JSON(http.StatusOK, profile.EditLastName(username, newLastName))
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

func CheckAdminHandler(c *gin.Context) {

	adminId, err := strconv.Atoi(c.Param("adminId"))
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, profile.CheckAdmin(adminId))
}
func MakeUserAdminHandler(c *gin.Context) {

	adminId, err := strconv.Atoi(c.Param("adminId"))
	userId, err := strconv.Atoi(c.Param("userId"))

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, profile.MakeUserAdmin(userId, adminId))
}
func UnmakeUserAdminHandler(c *gin.Context) {

	adminId, err := strconv.Atoi(c.Param("adminId"))
	userId, err := strconv.Atoi(c.Param("userId"))

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, profile.UnmakeUserAdmin(userId, adminId))
}
func DeletePostAdminHandler(c *gin.Context) {

	postId, err := strconv.Atoi(c.Param("postId"))
	adminId, err := strconv.Atoi(c.Param("adminId"))

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, profile.DeletePostAdmin(postId, adminId))
}
func DeleteUserProfileAdminHandler(c *gin.Context) {

	adminId, err := strconv.Atoi(c.Param("adminId"))
	username := c.Param("username")

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, profile.DeleteUserProfileAdmin(username, adminId))
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

func AddReactionHandler(c *gin.Context) {
	emoji := c.Param("emoji")
	post_id, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		fmt.Println(err)
	}
	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, reactions.AddReaction(emoji, post_id, user_id))
}

func UpdateReactionHandler(c *gin.Context) {
	emoji := c.Param("emoji")
	post_id, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		fmt.Println(err)
	}
	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, reactions.UpdateReaction(post_id, user_id, emoji))
}

func DeleteReactionHandler(c *gin.Context) {
	post_id, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		fmt.Println(err)
	}
	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, reactions.DeleteReaction(post_id, user_id))
}

func UploadImageHandler(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20) // Limit file size to 10 MB
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse form"})
		return
	}

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to retrieve file"})
		return
	}
	defer file.Close()

	username := c.PostForm("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	timestamp := time.Now().Format("20060102150405") // Format: YYYYMMDDHHMMSS
	fileExtension := ""
	if extIndex := strings.LastIndex(fileHeader.Filename, "."); extIndex != -1 {
		fileExtension = fileHeader.Filename[extIndex:]
	}
	fileName := fmt.Sprintf("%s_%s%s", username, timestamp, fileExtension)
	savePath := "uploads/" + fileName

	out, err := os.Create(savePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
		return
	}

	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/capstone", dbName, dbPass, dbHost)

	// Open a connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Ping the database to verify the connection is alive
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM pfp WHERE username = ?)", username).Scan(&exists)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query error"})
		return
	}

	if !exists {
		_, err = db.Exec("INSERT INTO pfp (username, image_name) VALUES (?, ?)", username, "none")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert new user"})
			return
		}
	}

	_, err = db.Exec("UPDATE pfp SET image_name = ? WHERE username = ?", fileName, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update image name"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "File uploaded successfully",
		"file_name": fileName,
		"file_path": savePath,
	})
}

func GetProfilePictureHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username") // Get username from query parameters

		if username == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
			return
		}

		//fmt.Println("Querying for username:", username)

		var imageName string
		err := db.QueryRow("SELECT image_name FROM pfp WHERE username = ?", username).Scan(&imageName)

		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			} else {
				fmt.Println("Database query error:", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			}
			return
		}

		//fmt.Println(imageName)

		c.JSON(http.StatusOK, gin.H{
			"imageName": imageName,
		})
	}
}
