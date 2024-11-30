package main

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"fakebook.com/project/feed"
	"fakebook.com/project/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Database configuration
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

	// Create a Gin router
	r := gin.Default()
	r.Use(CORSMiddleware())

	// Serve frontend files for unmatched routes
	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("../Frontend/dist/frontend/index.html")
		} else if ext == ".woff2?dd67030699838ea613ee6dbda90effa6" {
			c.File("../Frontend/dist/frontend/bootstrap-icons.bfa90bda92a84a6a.woff2")
		} else if ext == ".jpg" {
			c.File("../Frontend/src/app/" + path.Join(dir, file))
		} else {
			c.File("../Frontend/dist/frontend/" + path.Join(dir, file))
		}
	})

	// Add authorized routes
	authorized := r.Group("/")
	// authorized.Use(authRequired())
	authorized.GET("/api/users", handlers.GetUsersHandler)
	authorized.GET("/api/user/:id", handlers.GetOneUserHandler)
	authorized.GET("/api/username/:username", handlers.GetOneUserbyUsernameHandler)
	authorized.POST("/api/user/addNewUser", handlers.AddNewUserHandler)
	authorized.PATCH("/api/user/editFullName/:newName/:username", handlers.EditNameHandler)
	authorized.PATCH("/api/user/editFirstName/:newFirstName/:username", handlers.EditFirstNameHandler)
	authorized.PATCH("/api/user/editLastName/:newLastName/:username", handlers.EditLastNameHandler)
	authorized.PATCH("/api/user/editBio/:newBio/:username", handlers.EditBioHandler)
	authorized.PATCH("/api/user/editUsername/:newUsername/:username", handlers.EditUsernameHandler)
	authorized.DELETE("/api/user/deleteUser/:username", handlers.DeleteUserHandler)

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

	authorized.GET("/uploads/*filepath", handlers.FileServerHandler)
	authorized.POST("/upload", handlers.UploadImageHandler)
	authorized.GET("/getProfilePicture", handlers.GetProfilePictureHandler(db))

	err = r.Run(":3000")
	if err != nil {
		panic(err)
	}

	feed.GetPostData(2)
}

func addRoutes(group *gin.RouterGroup, db *sql.DB) {
	group.GET("/api/users", handlers.GetUsersHandler)
	group.GET("/api/user/:id", handlers.GetOneUserHandler)
	group.GET("/api/username/:username", handlers.GetOneUserbyUsernameHandler)
	group.POST("/api/user/addNewUser", handlers.AddNewUserHandler)
	group.PATCH("/api/user/editFullName/:newName/:username", handlers.EditNameHandler)
	group.PATCH("/api/user/editFirstName/:newFirstName/:username", handlers.EditFirstNameHandler)
	group.PATCH("/api/user/editLastName/:newLastName/:username", handlers.EditLastNameHandler)
	group.PATCH("/api/user/editBio/:newBio/:username", handlers.EditBioHandler)
	group.PATCH("/api/user/editUsername/:newUsername/:username", handlers.EditUsernameHandler)
	group.DELETE("/api/user/deleteUser/:username", handlers.DeleteUserHandler)

	group.GET("/api/posts/user/:userID/:loggedInUserId", handlers.GetUserPostsHandler)
	group.GET("/api/posts/initial/:numOfPosts/:loggedInUserId", handlers.GetInitialFeedByTimeHandler)
	group.GET("/api/posts/:numOfPosts/:loggedInUserId", handlers.GetFeedByTimeHandler)
	group.POST("/api/posts/:userId/:postText", handlers.AddPostHandler)
	group.POST("/api/posts/reply", handlers.AddReplyHandler)
	group.GET("/api/posts/getAllReplies/:postId", handlers.GetAllRepliesHandler)
	group.GET("/api/user/findUserByName/:fullName", handlers.FindUserByNameHandler)
	group.GET("/api/user/findUserByFirstAndLastName/:firstName/:lastName", handlers.FindUserByFullNameHandler)
	group.GET("/api/friends/findFriendList/:username", handlers.GetFriendsListHandler)
	group.GET("/api/friends/findFriendRequestList/:username", handlers.GetFriendRequestListHandler)
	group.PUT("/api/friends/addPendingFriendship/:requestor/:requestee", handlers.AddOneFriendHandler)
	group.GET("/api/friends/acceptFriendship/:originalRequestor/:acceptee", handlers.AcceptFriendshipHandler)
	group.DELETE("/api/friends/deleteFriendshipRequest/:originalRequestor/:deleter", handlers.DeleteFriendshipRequestHandler)
	group.DELETE("/api/friends/deleteFriendship/:friendToDelete/:deleter", handlers.DeleteFriendshipHandler)

	group.POST("/api/reactions/addReaction/:emoji/:post_id/:user_id", handlers.AddReactionHandler)
	group.DELETE("/api/reactions/deleteReaction/:post_id/:user_id", handlers.DeleteReactionHandler)

	group.GET("/uploads/*filepath", handlers.FileServerHandler)
	group.POST("/upload", handlers.UploadImageHandler)
	group.GET("/getProfilePicture", handlers.GetProfilePictureHandler(db))
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
