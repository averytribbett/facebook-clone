/*
A lot of the code in this file has to do with connecting frontend to backend via Auth0, which I was not able to get working.
I have the proxy-conf.json file connecting them at the moment
*/

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"fakebook.com/project/feed"
	"fakebook.com/project/handlers"
	"fakebook.com/project/models"
	"github.com/auth0-community/go-auth0"
	"github.com/gin-gonic/gin"
	jose "gopkg.in/square/go-jose.v2"
)

var (
	audience string
	domain   string
)

type User models.User

func main() {
	// initialPostCount := 3

	// println("\n\n\nStarting time sort: \n")

	// testarr := feed.InitialFeedByTime(initialPostCount)
	// feed.DisplayModel(testarr)
	// testarr = feed.FeedByTime(initialPostCount)
	// feed.DisplayModel(testarr)

	// println("\n\n\nStarting random sort: \n")

	// var used []int
	// testarr, used = feed.InitialFeedByRandom(3)
	// feed.DisplayModel(testarr)
	// testarr = feed.FeedByRandom(used)
	// feed.DisplayModel(testarr)

	// AddReaction("thumbs_up", 3, 3)

	// setAuth0Variables()

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

	r := gin.Default()
	r.Use(CORSMiddleware())
	// r.Use(CORSMiddleware())
	// This will ensure that the angular files are served correctly
	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("../Frontend/dist/frontend/index.html")
		} else if ext == ".woff2?dd67030699838ea613ee6dbda90effa6" {
			c.File("../Frontend/dist/frontend/bootstrap-icons.bfa90bda92a84a6a.woff2") // not sure why this one is so weird, will troubleshoot eventually
		} else if ext == ".jpg" {
			c.File("../Frontend/src/app/" + path.Join(dir, file))
		} else {
			c.File("../Frontend/dist/frontend/" + path.Join(dir, file))
		}
	})

	authorized := r.Group("/")
	// authorized.Use(authRequired())
	authorized.GET("/api/users", handlers.GetUsersHandler)
	authorized.GET("/api/user/:id", handlers.GetOneUserHandler)
	authorized.GET("/api/username/:username", handlers.GetOneUserbyUsernameHandler)
	authorized.PUT("/api/user/addNewUser", handlers.AddNewUserHandler)
	authorized.PATCH("/api/user/editFullName/:newName/:username", handlers.EditNameHandler)
	authorized.PATCH("/api/user/editFirstName/:newFirstName/:username", handlers.EditFirstNameHandler)
	authorized.PATCH("/api/user/editLastName/:newLastName/:username", handlers.EditLastNameHandler)
	authorized.PATCH("/api/user/editBio/:newBio/:username", handlers.EditBioHandler)
	authorized.PATCH("/api/user/editUsername/:newUsername/:username", handlers.EditUsernameHandler)
	authorized.DELETE("/api/user/deleteUser/:username", handlers.DeleteUserHandler)

	authorized.GET("/api/checkAdmin/:adminId", handlers.CheckAdminHandler)
	authorized.PUT("/api/makeAdmin/:userId/:adminId", handlers.MakeUserAdminHandler)
	authorized.DELETE("/api/unmakeAdmin/:userId/:adminId", handlers.UnmakeUserAdminHandler)
	authorized.DELETE("/api/deletePostAdmin/:postId/:adminId", handlers.DeletePostAdminHandler)
	authorized.DELETE("/api/deleteUserAdmin/:username/:adminId", handlers.DeleteUserProfileAdminHandler)

	authorized.GET("/api/posts/user/:userID/:loggedInUserId", handlers.GetUserPostsHandler)
	authorized.GET("/api/posts/initial/:numOfPosts/:loggedInUserId", handlers.GetInitialFeedByTimeHandler)
	authorized.GET("/api/posts/:numOfPosts/:loggedInUserId", handlers.GetFeedByTimeHandler)
	authorized.POST("/api/posts/:userId/:postText", handlers.AddPostHandler)
	authorized.POST("/api/posts/reply", handlers.AddReplyHandler)
	authorized.GET("/api/posts/getAllReplies/:postId", handlers.GetAllRepliesHandler)
	authorized.GET("/api/user/findUserByName/:fullName", handlers.FindUserByNameHandler)
	authorized.GET("/api/user/findUserByFirstAndLastName/:firstName/:lastName", handlers.FindUserByFullNameHandler)
	authorized.GET("/api/friends/findFriendList/:username", handlers.GetFriendsListHandler)
	authorized.GET("/api/friends/findFriendRequestList/:username", handlers.GetFriendRequestListHandler)
	authorized.PUT("/api/friends/addPendingFriendship/:requestor/:requestee", handlers.AddOneFriendHandler)
	authorized.GET("/api/friends/acceptFriendship/:originalRequestor/:acceptee", handlers.AcceptFriendshipHandler)
	authorized.DELETE("/api/friends/deleteFriendshipRequest/:originalRequestor/:deleter", handlers.DeleteFriendshipRequestHandler)
	authorized.DELETE("/api/friends/deleteFriendship/:friendToDelete/:deleter", handlers.DeleteFriendshipHandler)

	authorized.POST("/api/reactions/addReaction/:emoji/:post_id/:user_id", handlers.AddReactionHandler)
	authorized.PUT("/api/reactions/updateReaction/:emoji/:post_id/:user_id", handlers.UpdateReactionHandler)
	authorized.DELETE("/api/reactions/deleteReaction/:post_id/:user_id", handlers.DeleteReactionHandler)

	// authorized.GET("/uploads/*filepath", handlers.FileServerHandler)
	authorized.POST("/upload", handlers.UploadImageHandler)
	authorized.GET("/getProfilePicture", handlers.GetProfilePictureHandler(db))

	err = r.Run(":3000")
	if err != nil {
		panic(err)
	}

	feed.GetPostData(2)

}

func setAuth0Variables() {
	audience = os.Getenv("AUTH0_API_IDENTIFIER")
	domain = os.Getenv("AUTH0_DOMAIN")
	fmt.Println(audience)
	fmt.Println(domain)
}

// ValidateRequest will verify that a token received from an http request
// is valid and signyed by Auth0
func authRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		var auth0Domain = "https://" + domain + "/"
		client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: auth0Domain + ".well-known/jwks.json"}, nil)
		configuration := auth0.NewConfiguration(client, []string{audience}, auth0Domain, jose.RS256)
		validator := auth0.NewValidator(configuration, nil)

		_, err := validator.ValidateRequest(c.Request)

		if err != nil {
			log.Println(err)
			terminateWithError(http.StatusUnauthorized, "token is not valid", c)
			return
		}
		c.Next()
	}
}

func terminateWithError(statusCode int, message string, c *gin.Context) {
	c.JSON(statusCode, gin.H{"error": message})
	c.Abort()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func AddReply(text string, post_id int, username string) {
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

	// sql query
	query := fmt.Sprintf("INSERT INTO replies VALUES (%s, %s, %s);", strconv.Itoa(post_id), username, text)

	// execute sql
	_, err = db.Query(query)
	if err != nil {
		panic(err)
	}
}

// func to get the friends list for a user
func GetFriendsList(username string) []models.Friendlist {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/capstone", dbName, dbPass, dbHost)
	var data []models.Friendlist

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

	// sql query to get reaction info
	query := "SELECT first_name, last_name, username FROM users JOIN friends ON users.username = friends.friend_id WHERE friends.user_id = '" + username + "' AND friends.friend_status = 'friends';"
	// x rows of sql result
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// format each row of the result
	for rows.Next() {
		var friend models.Friendlist
		// scan result and set the values to each variable
		err = rows.Scan(&friend.FirstName, &friend.LastName, &friend.Username)
		if err != nil {
			panic(err)
		}
		data = append(data, friend)
	}
	db.Close()
	return data
}

// checks the friend status from a user to another user, returns a bool of the status
// example status = "friends" would return true if two users are friends
// statuses can be friends or pending
func StatusCheck(username string, friendUsername string, status string) bool {
	dsn := "capstone:csc450@tcp(71.89.73.28:3306)/capstone"

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

	// sql query
	query := fmt.Sprintf("SELECT * FROM friends where user_id = '%s' and friend_id = '%s' and friend_status = '%s';", username, friendUsername, status)

	// x rows of sql result
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		count++
	}

	if count < 1 {
		return false
	}
	return true
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from http://localhost:4200
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
