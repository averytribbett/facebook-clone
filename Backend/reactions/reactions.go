package reactions

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func AddReaction(emoji string, post_id int, user_id int) bool {
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

	// SQL query
	query := "INSERT INTO reactions (post_id, user_id, reaction) VALUES (?, ?, ?);"
	_, err = db.Exec(query, post_id, user_id, emoji)
	if err != nil {
		panic(err)
	}
	return true
}

// @TODO update reaction
func UpdateReaction(post_id int, user_id int, emoji string) bool {
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

	query := "UPDATE reactions SET reaction = ? WHERE post_id = ? AND user_id = ?;"
	_, err = db.Exec(query, emoji, post_id, user_id)
	if err != nil {
		panic(err)
	}
	return true
}

func DeleteReaction(post_id int, user_id int) bool {
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

	query := "DELETE FROM reactions WHERE post_id = ? AND user_id = ?;"
	_, err = db.Exec(query, post_id, user_id)
	if err != nil {
		panic(err)
	}
	return true
}
