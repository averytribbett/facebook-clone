package friends

import (
	"database/sql"
	"fmt"
	"os"

	"fakebook.com/project/models"
	_ "github.com/go-sql-driver/mysql"
)

const (
	FRIENDS = "friends"
	PENDING = "pending"
	BLOCKED = "blocked"
)

func GetFriendsList(username string) []models.Friend {
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

	var friendList []models.Friend
	// sql query
	rows, err2 := db.Query("SELECT user_id, friend_id, friend_status FROM friends WHERE user_id = ? OR friend_id = ?", username, username)

	if err2 != nil {
		panic(err2)
	}

	for rows.Next() {
		var friend models.Friend

		err = rows.Scan(&friend.User_id, &friend.Friend_id, &friend.Friend_status)
		if err != nil {
			panic(err)
		}

		friendList = append(friendList, friend)
	}

	return friendList
}

func AddPendingFriend(friendRequestor string, friendRequestee string) error {
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

	query := "INSERT INTO friends (user_id, friend_id, friend_status) VALUES (?, ?, ?)"

	_, err3 := db.Exec(query, friendRequestor, friendRequestee, PENDING)

	return err3
}

func AcceptFriend(originalRequestor string, acceptee string) error {
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

	query1 := "INSERT INTO friends (user_id, friend_id, friend_status) VALUES (?, ?, ?)"
	_, err2 := db.Exec(query1, acceptee, originalRequestor, FRIENDS)
	if err2 != nil {
		panic(err2)
	}

	query2 := "UPDATE friends SET friend_status = ? WHERE user_id = ? AND friend_id = ?"
	_, err3 := db.Exec(query2, FRIENDS, originalRequestor, acceptee)
	if err3 != nil {
		panic(err3)
	}
	return err3
}

func DeleteFriendRequest(originalRequestor string, deleter string) error {
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

	query := "DELETE FROM friends WHERE user_id = ? AND friend_id = ? AND friend_status = ?"
	_, err2 := db.Exec(query, originalRequestor, deleter, PENDING)
	if err2 != nil {
		panic(err2)
	}

	return err2
}

func DeleteFriend(friendToDelete string, deleter string) error {
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

	query := "DELETE FROM friends WHERE user_id = ? AND friend_id = ? AND friend_status = ?"
	_, err2 := db.Exec(query, friendToDelete, deleter, FRIENDS)
	if err2 != nil {
		panic(err2)
	}

	query2 := "DELETE FROM friends WHERE user_id = ? AND friend_id = ? AND friend_status = ?"
	_, err3 := db.Exec(query2, deleter, friendToDelete, FRIENDS)
	if err3 != nil {
		panic(err3)
	}

	return err3
}
