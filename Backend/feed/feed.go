package feed

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"fakebook.com/project/models"
	_ "github.com/go-sql-driver/mysql"
)

// func to get individual post
func GetPostData(post_id int) []string {
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
	query := "SELECT posts.post_id, posts.post_text, users.username AS post_author, users.first_name AS post_author_first_name, users.last_name AS post_author_last_name FROM posts JOIN users ON posts.user_id = users.id LEFT JOIN reactions ON posts.post_id = reactions.post_id WHERE posts.post_id=\"" + strconv.Itoa(post_id) + "\"ORDER BY posts.post_id;"

	// x rows of sql result
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// format each row of the result
	var data []string
	for rows.Next() {
		var postID string
		var postText string
		var postAuthor string
		var postAuthorFirstName string
		var postAuthorLastName string

		// scan result and set the values to each variable
		err = rows.Scan(&postID, &postText, &postAuthor, &postAuthorFirstName, &postAuthorLastName)
		if err != nil {
			panic(err)
		}

		data = append(data, postID)
		data = append(data, postText)
		data = append(data, postAuthor)
		data = append(data, postAuthorFirstName)
		data = append(data, postAuthorLastName)
	}
	db.Close()
	return data
}

// func to get all the posts from a user
func GetUserPosts(user_id int) [][]string {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/capstone", dbName, dbPass, dbHost)
	// arr of posts
	var data [][]string

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
	query := "SELECT posts.post_id, posts.post_text, users.username AS post_author, users.first_name AS post_author_first_name, users.last_name AS post_author_last_name FROM posts JOIN users ON posts.user_id = users.id LEFT JOIN reactions ON posts.post_id = reactions.post_id WHERE posts.user_id=\"" + strconv.Itoa(user_id) + "\"ORDER BY posts.post_id;"

	// x rows of sql result
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// format each row of the result
	for rows.Next() {
		var postID string
		var postText string
		var postAuthor string
		var postAuthorFirstName string
		var postAuthorLastName string

		// scan result and set the values to each variable
		err = rows.Scan(&postID, &postText, &postAuthor, &postAuthorFirstName, &postAuthorLastName)
		if err != nil {
			panic(err)
		}

		var post []string
		post = append(post, postID)
		post = append(post, postText)
		post = append(post, postAuthor)
		post = append(post, postAuthorFirstName)
		post = append(post, postAuthorLastName)
		data = append(data, post)
	}
	db.Close()
	return data
}

// func to initialize feed by time sort
func InitialFeedByTime(numOfPosts int) []models.Post {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/capstone", dbName, dbPass, dbHost)
	var data []models.Post

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
	query := "SELECT posts.post_id, posts.post_text, users.id, users.first_name, users.last_name FROM posts JOIN users ON posts.user_id = users.id LEFT JOIN reactions ON posts.post_id = reactions.post_id ORDER BY posts.post_id DESC;"

	// x rows of sql result
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// format each row of the result
	for rows.Next() {

		var post models.Post
		// scan result and set the values to each variable
		err = rows.Scan(&post.Id, &post.Text, &post.AuthorId, &post.AuthorFirstName, &post.AuthorLastName)
		if err != nil {
			panic(err)
		}

		data = append(data, post)
	}
	db.Close()
	return data
}

// func to sort feed by post time
func FeedByTime(numOfPosts int) [][]string {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/capstone", dbName, dbPass, dbHost)
	var data [][]string

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
	query := "SELECT posts.post_id, posts.post_text, users.username AS post_author, users.first_name AS post_author_first_name, users.last_name AS post_author_last_name FROM posts JOIN users ON posts.user_id = users.id LEFT JOIN reactions ON posts.post_id = reactions.post_id ORDER BY posts.post_id DESC LIMIT " + strconv.Itoa(numOfPosts) + ";"

	// x rows of sql result
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// skip first x posts (+1 because post starts from 1)
	for i := 1; i < numOfPosts+1; i++ {
		rows.Next()
	}

	// format each row of the result
	for rows.Next() {
		var postID string
		var postText string
		var postAuthor string
		var postAuthorFirstName string
		var postAuthorLastName string
		println(postID, postText, postAuthor, postAuthorFirstName, postAuthorLastName, "qwdqwd")

		var post []string
		post = append(post, postID)
		post = append(post, postText)
		post = append(post, postAuthor)
		post = append(post, postAuthorFirstName)
		post = append(post, postAuthorLastName)
		data = append(data, post)
		print(data)
	}
	db.Close()
	return data
}

// func to initialize feed for random sorting
func InitialFeedByRandom(numOfPosts int) ([][]string, []int) {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/capstone", dbName, dbPass, dbHost)
	var data [][]string
	var usedPosts []int

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
	query := "SELECT posts.post_id, posts.post_text, users.username AS post_author, users.first_name AS post_author_first_name, users.last_name AS post_author_last_name FROM posts JOIN users ON posts.user_id = users.id LEFT JOIN reactions ON posts.post_id = reactions.post_id ORDER BY RAND() DESC LIMIT " + strconv.Itoa(numOfPosts) + ";"

	// x rows of sql result
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// format each row of the result
	for rows.Next() {
		var postID string
		var postText string
		var postAuthor string
		var postAuthorFirstName string
		var postAuthorLastName string

		// scan result and set the values to each variable
		err = rows.Scan(&postID, &postText, &postAuthor, &postAuthorFirstName, &postAuthorLastName)
		if err != nil {
			panic(err)
		}

		// add post_id to used posts
		var add int
		add, err := strconv.Atoi(postID)
		if err != nil {
			panic(err)
		}
		usedPosts = append(usedPosts, add)

		var post []string
		post = append(post, postID)
		post = append(post, postText)
		post = append(post, postAuthor)
		post = append(post, postAuthorFirstName)
		post = append(post, postAuthorLastName)
		data = append(data, post)
	}
	db.Close()
	// return the x posts used and also the array of used posts so they arent printed again
	return data, usedPosts
}

// func to sort feed by random
func FeedByRandom(exclude []int) [][]string {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/capstone", dbName, dbPass, dbHost)
	var data [][]string

	// map of used posts (used to skip posts that have already been printed)
	used := make(map[int]int)
	for i := 0; i < len(exclude); i++ {
		used[exclude[i]] = 0
	}

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
	query := "SELECT posts.post_id, posts.post_text, users.username AS post_author, users.first_name AS post_author_first_name, users.last_name AS post_author_last_name FROM posts JOIN users ON posts.user_id = users.id LEFT JOIN reactions ON posts.post_id = reactions.post_id ORDER BY RAND() DESC;"

	// x rows of sql result
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// format each row of the result
	for rows.Next() {
		var postID string
		var postText string
		var postAuthor string
		var postAuthorFirstName string
		var postAuthorLastName string

		// scan result and set the values to each variable
		err = rows.Scan(&postID, &postText, &postAuthor, &postAuthorFirstName, &postAuthorLastName)
		if err != nil {
			panic(err)
		}

		// check if post is already used
		currID, err := strconv.Atoi(postID)
		if err != nil {
			panic(err)
		}
		// get to the existence check
		_, exists := used[currID]

		// if key exists in map, skip
		if exists {
			continue
		}

		var post []string
		post = append(post, postID)
		post = append(post, postText)
		post = append(post, postAuthor)
		post = append(post, postAuthorFirstName)
		post = append(post, postAuthorLastName)
		data = append(data, post)
	}
	db.Close()
	return data
}

// func to display posts
func DisplayPost(posts []string) {
	for i := 0; i < len(posts); i++ {
		println(posts[i] + "\n")
	}
}

// func to display posts
func DisplayPostArr(posts [][]string) {
	for i := 0; i < len(posts); i++ {
		for j := 0; j < len(posts[i]); j++ {
			println(posts[i][j])
		}
	}
}

// func to get replies on a post
func GetReplies(post_id int) [][]string {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/capstone", dbName, dbPass, dbHost)
	var data [][]string

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

	// sql query to grab reply info
	query := "SELECT replier.username AS replier_username, replier.first_name AS replier_first_name, replier.last_name AS replier_last_name, replies.reply_text AS reply_text FROM posts JOIN users ON posts.user_id = users.id LEFT JOIN replies ON posts.post_id = replies.post_id LEFT JOIN users AS replier ON replies.user_id = replier.id WHERE posts.post_id = " + strconv.Itoa(post_id) + ";"
	// x rows of sql result
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// format each row of the result
	for rows.Next() {
		var replierUsername string
		var replierFirstName string
		var replierLastName string
		var replyText string

		// scan result and set the values to each variable
		err = rows.Scan(&replierUsername, &replierFirstName, &replierLastName, &replyText)
		if err != nil {
			panic(err)
		}

		var reply []string
		reply = append(reply, replierUsername)
		reply = append(reply, replierFirstName)
		reply = append(reply, replierLastName)
		reply = append(reply, replyText)
		data = append(data, reply)
	}
	db.Close()
	// return the reply data
	return data
}

// func to get reactions to each post
func GetReactions(post_id int) [][]string {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/capstone", dbName, dbPass, dbHost)
	var data [][]string

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
	query := "SELECT reacter.username AS reacter_username, reacter.first_name AS reacter_first_name, reacter.last_name AS reacter_last_name, reactions.reaction AS reaction FROM posts JOIN users ON posts.user_id = users.id LEFT JOIN reactions ON posts.post_id = reactions.post_id LEFT JOIN users AS reacter ON reactions.user_id = reacter.id WHERE posts.post_id = " + strconv.Itoa(post_id) + ";"
	// x rows of sql result
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// format each row of the result
	for rows.Next() {
		var reacterUsername string
		var reacterFirstName string
		var reacterLastName string
		var reaction string

		// scan result and set the values to each variable
		err = rows.Scan(&reacterUsername, &reacterFirstName, &reacterLastName, &reaction)
		if err != nil {
			panic(err)
		}

		var reply []string
		reply = append(reply, reacterUsername)
		reply = append(reply, reacterFirstName)
		reply = append(reply, reacterLastName)
		reply = append(reply, reaction)
		data = append(data, reply)
	}
	db.Close()
	// return the reply data
	return data
}
