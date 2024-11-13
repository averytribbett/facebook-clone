package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

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

	username := r.FormValue("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// creates a filename based on the username and when it was uploaded
	extension := filepath.Ext(handler.Filename)
	newFileName := fmt.Sprintf("%s_%d%s", username, time.Now().UnixNano(), extension)

	uploadDir := "uploads"
	os.MkdirAll(uploadDir, os.ModePerm)
	filePath := filepath.Join(uploadDir, newFileName)

	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/capstone", dbName, dbPass, dbHost)

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

	stmt, err := db.Prepare("delete from capstone.pfp where username = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(username)
	if err != nil {
		panic(err)
	}

	stmt, err = db.Prepare("INSERT INTO pfp (username, image_name) VALUES (?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, newFileName)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "File uploaded successfully for user %s: %s\n", username, newFileName)
}

func GetProfilePictureHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the username from the query parameter
		username := r.URL.Query().Get("username")
		if username == "" {
			http.Error(w, "Username is required", http.StatusBadRequest)
			return
		}

		var imageName string
		query := "SELECT image_name FROM pfp WHERE username = ?"
		err := db.QueryRow(query, username).Scan(&imageName)
		if err == sql.ErrNoRows {
			http.Error(w, "No profile picture found for this user", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			fmt.Println("Error retrieving profile picture:", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		jsonResponse := map[string]string{"imageName": imageName}
		json.NewEncoder(w).Encode(jsonResponse)
	}
}
