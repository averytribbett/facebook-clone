/*
A lot of the code in this file has to do with connecting frontend to backend via Auth0, which I was not able to get working.
I have the proxy-conf.json file connecting them at the moment
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

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
	println("get by post num: \n")
	test := feed.GetPostData(3)
	feed.DisplayPost(test)
	println("\n\n\n")

	println("get by post user: \n")
	var test2 [][]string
	test2 = feed.GetUserPosts(1)
	feed.DisplayPostArr(test2)
	print("\n\n\n")

	initialPostCount := 3

	println("\n\n\nStarting random sort: \n")
	var testarr [][]string
	var testdupe []int
	testarr, testdupe = feed.InitialFeedByRandom(initialPostCount)
	feed.DisplayPostArr(testarr)
	testarr = feed.FeedByRandom(testdupe)
	feed.DisplayPostArr(testarr)
	print("\n\n\n")

	println("\n\n\nStarting time sort: \n")

	feed.InitialFeedByTime(initialPostCount)
	feed.DisplayPostArr(testarr)
	testarr = feed.FeedByTime(initialPostCount)
	feed.DisplayPostArr(testarr)

	println("\n\n\nReplies to post #3 \n")

	testarr = feed.GetReplies(3)
	feed.DisplayPostArr(testarr)

	println("\n\n\nReactions to post #5 \n")

	testarr = feed.GetReactions(5)
	feed.DisplayPostArr(testarr)

	// setAuth0Variables()

	r := gin.Default()
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

	authorized.GET("/api/posts/initial/:numOfPosts", handlers.GetInitialFeedByTime)
	authorized.GET("/api/user/findUserByName/:fullName", handlers.FindUserByNameHandler)
	authorized.GET("/api/user/findUserByFirstAndLastName/:firstName/:lastName", handlers.FindUserByFullNameHandler)
	authorized.GET("/api/friends/findFriendList/:username", handlers.GetFriendsListHandler)
	authorized.PUT("/api/friends/addPendingFriendship/:requestor/:requestee", handlers.AddOneFriendHandler)
	authorized.GET("/api/friends/acceptFriendship/:originalRequestor/:acceptee", handlers.AcceptFriendshipHandler)
	authorized.DELETE("/api/friends/deleteFriendshipRequest/:originalRequestor/:deleter", handlers.DeleteFriendshipRequestHandler)
	authorized.DELETE("/api/friends/deleteFriendship/:friendToDelete/:deleter", handlers.DeleteFriendshipHandler)

	err := r.Run(":3000")
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
