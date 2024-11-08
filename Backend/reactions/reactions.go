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

	// swap emoji name with the html code for the emoji
	code := ""
	switch emoji {
	case "thumbs_up":
		code = "&#128077;"
	case "thumbs_down":
		code = "&#128078;"
	case "heart":
		code = "&#129505;"
	default:
		return false
	}

	// SQL query
	query := "INSERT INTO reactions (post_id, user_id, reaction) VALUES (?, ?, ?);"
	_, err = db.Exec(query, post_id, user_id, code)
	if err != nil {
		panic(err)
	}
	return true
}

// @TODO update reaction

// @TODO delete reaction
