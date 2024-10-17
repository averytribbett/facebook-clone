package profile

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"sync"

	_ "github.com/go-sql-driver/mysql"

	"fakebook.com/project/models"
)

var (

	mtx      sync.RWMutex
	onceList sync.Once
	onceDB   sync.Once
	db       *sql.DB
)

func init() {
	onceDB.Do(initializeDB)
}

func initializeDB() {
	// Environmental variables set to connect to Cade's database
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/capstone", dbName, dbPass, dbHost)

	var err error
	// Open a connection to the database
	db, err = sql.Open("mysql", dsn)

	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(0)

	if err != nil {
		fmt.Print("are we in here?")
		panic(err)
	}

	// Ping the database to verify the connection is alive
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}


// Function returns a list of user models, with all the users from the database
func Get() []models.User {

	var list []models.User
	rows, err := db.Query("SELECT id, first_name, last_name, username, bio FROM users")

	// Printing error if the query does not run properly
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	//Looping through the result rows from the query, then appending each user ot the list	
	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Username, &user.Bio)
		if err != nil {
			panic(err)
		}

		log.Println(user)
		list = append(list, user)
	}
	return list
}

// Returns one user from the databas using the user id
func GetOneUser(userId int) models.User {

	var returnUser models.User

	err := db.QueryRow("SELECT id, first_Name, last_name, username, bio FROM users WHERE id = ?", userId).Scan(&returnUser.Id, &returnUser.FirstName, &returnUser.LastName, &returnUser.Username, &returnUser.Bio)

	if err != nil {
		log.Println(err)
	}

	return returnUser
}

// Returns one user from the databas using the username
func GetOneUserByUsername(username string) models.User {

	var returnUser models.User

	err := db.QueryRow("SELECT id, first_Name, last_name, username, bio FROM users WHERE username = ?", username).Scan(&returnUser.Id, &returnUser.FirstName, &returnUser.LastName, &returnUser.Username, &returnUser.Bio)

	if err != nil {
		log.Println(err)
	}

	return returnUser
}

// Adds one user to the database
func AddNewUser(newUser models.User) error {
	var oldUser = GetOneUserByUsername(newUser.Username)

	if oldUser.Id == 0 {
		query := "INSERT INTO users (first_name, last_name, bio, username) VALUES (?, ?, ?, ?)"

		_, err := db.Exec(query, newUser.FirstName, newUser.LastName, newUser.Bio, newUser.Username)

		fmt.Println(err)

		return err
	} else {
		// not sure if good but, for now it gets the function to run
		return nil
	}
}

func CreatePost(body string) error{

	query:= "INSERT INTO posts (user_id, post_text) values (?,?)"
	log.Println("THIS IS THE BODY: " + body)
	_, err := db.Exec(query,7,body)

	return err
}

/*
endpoint ideas for a user profile:


Requirements:
1. getFriendList
2. getUserPosts
3. probably something regarding logging out / logging in?

Optional:
1. getUserPhotos


*/
