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

func Get() []models.User {
	var list []models.User
	rows, err := db.Query("SELECT id, first_name, last_name, username, bio FROM users")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

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

func GetOneUser(userId int) models.User {

	var returnUser models.User

	err := db.QueryRow("SELECT id, first_Name, last_name, username, bio FROM users WHERE id = ?", userId).Scan(&returnUser.Id, &returnUser.FirstName, &returnUser.LastName, &returnUser.Username, &returnUser.Bio)

	if err != nil {
		log.Println(err)
	}

	return returnUser
}

func GetOneUserByUsername(username string) models.User {

	var returnUser models.User

	err := db.QueryRow("SELECT id, first_Name, last_name, username, bio FROM users WHERE username = ?", username).Scan(&returnUser.Id, &returnUser.FirstName, &returnUser.LastName, &returnUser.Username, &returnUser.Bio)

	if err != nil {
		log.Println(err)
	}

	return returnUser
}

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

func FindUserByName(firstName string, lastName string) []models.User {
	var rows *sql.Rows
	var returnList []models.User
	var err error
	if len(lastName) > 0 {
		rows, err = db.Query("SELECT id, first_name, last_name, username, bio FROM users WHERE first_name LIKE '%" + firstName + "%' AND last_name LIKE '%" + lastName + "%'")
	} else {
		rows, err = db.Query("SELECT id, first_name, last_name, username, bio FROM users WHERE first_name LIKE '%" + firstName + "%'")
	}

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Username, &user.Bio)
		if err != nil {
			panic(err)
		}

		log.Println(user)
		returnList = append(returnList, user)
	}
	return returnList
}

/*
endpoint ideas for a user profile:

Tdo:
1. emails/usernames need to be uinque
2. change to prepared statements
3. do rows need to be closed?
4. does the DB need to be closed
5. more descriptive errors

Requirements:
1. getFriendList
2. getUserPosts
3. probably something regarding logging out / logging in?

Optional:
1. getUserPhotos


*/
