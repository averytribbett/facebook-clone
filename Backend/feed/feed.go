package main

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	println("get by post num: \n")
	test := getPostData(1)
	displayPost(test)
	println("\n\n\n")

	println("get by post user: \n")
	var test2 [][]string
	test2 = getUserPosts(1)
	displayPostArr(test2)
	print("\n\n\n")

	initialPostCount := 3

	println("\n\n\nStarting random sort: \n")
	var testarr [][]string
	var testdupe []int
	testarr, testdupe = initialFeedByRandom(initialPostCount)
	displayPostArr(testarr)
	testarr = feedByRandom(testdupe)
	displayPostArr(testarr)
	print("\n\n\n")

	println("\n\n\nStarting time sort: \n")

	arrarr := initialFeedByTime(initialPostCount)
	displayPostArr(arrarr)
	testarr = feedByTime(initialPostCount)
	displayPostArr(testarr)
}

// func to get individual post
func getPostData(post_id int) []string {
	dsn := "capstone:csc450@tcp(71.89.73.28:3306)/capstone"

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
func getUserPosts(user_id int) [][]string {
	dsn := "capstone:csc450@tcp(71.89.73.28:3306)/capstone"
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

		var tempArr []string
		tempArr = append(tempArr, postID)
		tempArr = append(tempArr, postText)
		tempArr = append(tempArr, postAuthor)
		tempArr = append(tempArr, postAuthorFirstName)
		tempArr = append(tempArr, postAuthorLastName)
		data = append(data, tempArr)
	}
	db.Close()
	return data
}

// func to initialize feed by time sort
func initialFeedByTime(numOfPosts int) [][]string {
	dsn := "capstone:csc450@tcp(71.89.73.28:3306)/capstone"
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
	query := "SELECT posts.post_id, posts.post_text, users.username AS post_author, users.first_name AS post_author_first_name, users.last_name AS post_author_last_name FROM posts JOIN users ON posts.user_id = users.id LEFT JOIN reactions ON posts.post_id = reactions.post_id ORDER BY posts.post_id DESC;"

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
		var tempArr []string
		tempArr = append(tempArr, postID)
		tempArr = append(tempArr, postText)
		tempArr = append(tempArr, postAuthor)
		tempArr = append(tempArr, postAuthorFirstName)
		tempArr = append(tempArr, postAuthorLastName)
		data = append(data, tempArr)
	}
	db.Close()
	return data
}

// func to sort feed by post time
func feedByTime(numOfPosts int) [][]string {
	dsn := "capstone:csc450@tcp(71.89.73.28:3306)/capstone"
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

		var tempArr []string
		tempArr = append(tempArr, postID)
		tempArr = append(tempArr, postText)
		tempArr = append(tempArr, postAuthor)
		tempArr = append(tempArr, postAuthorFirstName)
		tempArr = append(tempArr, postAuthorLastName)
		data = append(data, tempArr)
		print(data)
	}
	db.Close()
	return data
}

// func to initialize feed for random sorting
func initialFeedByRandom(numOfPosts int) ([][]string, []int) {
	dsn := "capstone:csc450@tcp(71.89.73.28:3306)/capstone"
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

		var tempArr []string
		tempArr = append(tempArr, postID)
		tempArr = append(tempArr, postText)
		tempArr = append(tempArr, postAuthor)
		tempArr = append(tempArr, postAuthorFirstName)
		tempArr = append(tempArr, postAuthorLastName)
		data = append(data, tempArr)
	}
	db.Close()
	// return the x posts used and also the array of used posts so they arent printed again
	return data, usedPosts
}

// func to sort feed by random
func feedByRandom(exclude []int) [][]string {
	dsn := "capstone:csc450@tcp(71.89.73.28:3306)/capstone"
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

		var tempArr []string
		tempArr = append(tempArr, postID)
		tempArr = append(tempArr, postText)
		tempArr = append(tempArr, postAuthor)
		tempArr = append(tempArr, postAuthorFirstName)
		tempArr = append(tempArr, postAuthorLastName)
		data = append(data, tempArr)
	}
	db.Close()
	return data
}

// func to display posts
func displayPost(posts []string) {
	for i := 0; i < len(posts); i++ {
		println(posts[i] + "\n")
	}
}

// func to display posts
func displayPostArr(posts [][]string) {
	for i := 0; i < len(posts); i++ {
		for j := 0; j < len(posts[i]); j++ {
			println(posts[i][j])
		}
	}
}
