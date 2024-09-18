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

	"fakebook.com/project/handlers"
	"github.com/auth0-community/go-auth0"
	"github.com/gin-gonic/gin"
	jose "gopkg.in/square/go-jose.v2"
)

var (
	audience string
	domain   string
)

func main() {
	setAuth0Variables()
	r := gin.Default()
	r.Use(CORSMiddleware())

	// This will ensure that the angular files are served correctly
	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		fmt.Println("request uri or whatever")
		fmt.Println(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("../Frontend/dist/frontend/index.html")
		} else {
			c.File("../Frontend/dist/frontend/" + path.Join(dir, file))
		}
	})

	authorized := r.Group("/")
	// authorized.Use(authRequired())
	authorized.GET("/api/users", handlers.GetUsersHandler)
	// authorized.POST("/todo", demoHandlers.AddTodoHandler)
	// authorized.DELETE("/todo/:id", demoHandlers.DeleteTodoHandler)
	// authorized.PUT("/todo", demoHandlers.CompleteTodoHandler)

	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
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
