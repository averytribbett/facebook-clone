package profile

import (
	"database/sql"
	"fmt"
	// "log"
	"sync"

	_ "github.com/go-sql-driver/mysql"

	"fakebook.com/project/models"
)

var (
	list []models.User
	mtx  sync.RWMutex
	onceList sync.Once
	onceDB sync.Once
	db *sql.DB
)

func init() {
	onceDB.Do(initializeDB)
	onceList.Do(initializeList)
}

func initializeDB(){

	dsn := "root:mysql@tcp(127.0.0.1:3306)/capstone"

	var err error
	// Open a connection to the database
	db, err = sql.Open("mysql", dsn)

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

func initializeList() {
	// this is simply a placeholder, and to show functionality of requests
	// in the future, we probably would not even need this
	list = []models.User{
		{
			Id:        0,
			FirstName: "Melissa",
			LastName:  "Brown",
			Username:  "melissa.cat.brown02@gmail.com",
		},
		{
			Id:        1,
			FirstName: "Avery",
			LastName:  "Tribbett",
			Username:  "averytribbett",
		},
		{
			Id:        2,
			FirstName: "Cade",
			LastName:  "Becker",
			Username:  "cadegithub",
		},
		{
			Id:        3,
			FirstName: "Youssef",
			LastName:  "Ibrahim",
			Username:  "youssefgithub",

		},
	}
}

func Get() []models.User {
	return list
}

func GetOneUser(userId int) models.User { // will return one user
	var returnUser models.User
	// replace lines 58-62 with call to database, this function might be obsolete eventually
	for _, user := range list {
		if user.Id == userId {
			returnUser = user
		}
	}
	return returnUser
}

func GetOneUserByUsername(username string) models.User {
	var returnUser models.User
	// replace lines 69-73 with calls to database
	for _, user := range list {
		if user.Username == username {
			returnUser = user
		}
	}
	return returnUser
}

func AddNewUser(newUser models.User) error {

	//todo disallow email duplicates
	fmt.Println(newUser)

	query :="INSERT INTO users (first_name, last_name, bio, username) VALUES (?, ?, ?, ?)"

	_, err := db.Exec(query, newUser.FirstName, newUser.LastName, newUser.Bio, newUser.Username)

	db.Close()

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
