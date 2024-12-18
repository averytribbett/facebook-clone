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
func GetUserPosts(user_id int, loggedInUserId int) []models.Post {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/capstone", dbName, dbPass, dbHost)
	// arr of posts
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
	query := "SELECT posts.post_id, posts.post_text, users.id AS user_id, users.first_name, users.last_name, users.username AS author_username, (SELECT COUNT(*) FROM replies WHERE replies.post_id = posts.post_id) AS reply_count, (SELECT COUNT(*) FROM reactions WHERE reactions.post_id = posts.post_id) AS reaction_count, (SELECT reaction FROM reactions WHERE reactions.post_id = posts.post_id AND reactions.user_id = ?) AS reaction_by_user FROM posts JOIN users ON posts.user_id = users.id LEFT JOIN reactions ON posts.post_id = reactions.post_id WHERE posts.user_id=\"" + strconv.Itoa(user_id) + "\"ORDER BY posts.post_id DESC;"

	// x rows of sql result
	rows, err := db.Query(query, loggedInUserId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// format each row of the result
	for rows.Next() {
		var post models.Post
		var reactionByUser sql.NullString

		// scan result and set the values to each variable
		err = rows.Scan(&post.Id, &post.Text, &post.AuthorId, &post.AuthorFirstName, &post.AuthorLastName, &post.AuthorUsername, &post.ReplyCount, &post.ReactionCount, &reactionByUser)
		if err != nil {
			panic(err)
		}

		// Handle NULL for `reaction_by_user`
		post.ReactionByUser = ""
		if reactionByUser.Valid {
			post.ReactionByUser = reactionByUser.String
		}

		data = append(data, post)
	}
	db.Close()
	return data
}

// func to initialize feed by time sort
func InitialFeedByTime(numOfPosts int, loggedInUserId int) []models.Post {
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
	query := `WITH friend_usernames
			  AS (SELECT CASE WHEN user_id = (SELECT username FROM capstone.users WHERE id = ?)
			  	  THEN friend_id WHEN friend_id = (SELECT username FROM capstone.users WHERE id = ?)
				  THEN user_id END AS friend_username FROM capstone.friends
				  WHERE (user_id = (SELECT username FROM capstone.users WHERE id = ?)
				  	OR friend_id = (SELECT username FROM capstone.users WHERE id = ?))
				  AND friend_status = 'friends')
			  SELECT posts.post_id, posts.post_text, users.id
			  AS user_id, users.first_name, users.last_name, users.username,
			  (SELECT COUNT(*) FROM replies
			  	WHERE replies.post_id = posts.post_id)
			  AS reply_count,
			  (SELECT COUNT(*) FROM reactions WHERE reactions.post_id = posts.post_id)
			  AS reaction_count,
			  (SELECT reaction FROM reactions WHERE reactions.post_id = posts.post_id AND reactions.user_id = ?)
			  AS reaction_by_user
			  FROM posts
			  JOIN users
			  ON posts.user_id = users.id
			  WHERE users.username
			  IN (SELECT friend_username FROM friend_usernames)
			  ORDER BY posts.post_id DESC LIMIT ?;`

	// x rows of sql result
	rows, err := db.Query(query, loggedInUserId, loggedInUserId, loggedInUserId, loggedInUserId, loggedInUserId, numOfPosts)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// format each row of the result
	for rows.Next() {
		var post models.Post
		var reactionByUser sql.NullString

		// scan result and set the values to each variable
		err = rows.Scan(&post.Id, &post.Text, &post.AuthorId, &post.AuthorFirstName, &post.AuthorLastName, &post.AuthorUsername, &post.ReplyCount, &post.ReactionCount, &reactionByUser)
		if err != nil {
			panic(err)
		}

		// Handle NULL for `reaction_by_user`
		post.ReactionByUser = ""
		if reactionByUser.Valid {
			post.ReactionByUser = reactionByUser.String
		}

		data = append(data, post)
	}
	db.Close()
	return data
}

// func to sort feed by post time
func FeedByTime(numOfPosts int, loggedInUserId int) []models.Post {
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
	query := `WITH friend_usernames
			  AS (SELECT CASE WHEN user_id = (SELECT username FROM capstone.users WHERE id = ?)
			  	  THEN friend_id WHEN friend_id = (SELECT username FROM capstone.users WHERE id = ?)
				  THEN user_id END AS friend_username FROM capstone.friends
				  WHERE (user_id = (SELECT username FROM capstone.users WHERE id = ?)
				  	OR friend_id = (SELECT username FROM capstone.users WHERE id = ?))
				  AND friend_status = 'friends')
			  SELECT posts.post_id, posts.post_text, users.id
			  AS user_id, users.first_name, users.last_name, users.username,
			  (SELECT COUNT(*) FROM replies
			  	WHERE replies.post_id = posts.post_id)
			  AS reply_count,
			  (SELECT COUNT(*) FROM reactions WHERE reactions.post_id = posts.post_id)
			  AS reaction_count,
			  (SELECT reaction FROM reactions WHERE reactions.post_id = posts.post_id AND reactions.user_id = ?)
			  AS reaction_by_user
			  FROM posts
			  JOIN users
			  ON posts.user_id = users.id
			  WHERE users.username
			  IN (SELECT friend_username FROM friend_usernames)
			  ORDER BY posts.post_id DESC;`

	// x rows of sql result
	rows, err := db.Query(query, loggedInUserId, loggedInUserId, loggedInUserId, loggedInUserId, loggedInUserId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// skip first x posts
	for i := 0; i < numOfPosts; i++ {
		// End of feed
		if !rows.Next() {
			return data
		}
	}

	// format each row of the result
	for j := 0; j < 20; j++ {
		// End of feed not a multiple of 20
		if !rows.Next() {
			break
		}
		var post models.Post
		var reactionByUser sql.NullString

		// scan result and set the values to each variable
		err = rows.Scan(&post.Id, &post.Text, &post.AuthorId, &post.AuthorFirstName, &post.AuthorLastName, &post.AuthorUsername, &post.ReplyCount, &post.ReactionCount, &reactionByUser)
		if err != nil {
			panic(err)
		}

		// Handle NULL for `reaction_by_user`
		post.ReactionByUser = ""
		if reactionByUser.Valid {
			post.ReactionByUser = reactionByUser.String
		}

		data = append(data, post)
	}
	db.Close()
	return data
}

// func to initialize feed for random sorting
func InitialFeedByRandom(numOfPosts int) ([]models.Post, []int) {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/capstone", dbName, dbPass, dbHost)
	var data []models.Post
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
	query := "SELECT posts.post_id, posts.post_text, users.id AS user_id, users.first_name, users.last_name, (SELECT COUNT(*) FROM replies WHERE replies.post_id = posts.post_id) AS reply_count, (SELECT COUNT(*) FROM reactions WHERE reactions.post_id = posts.post_id) AS reaction_count FROM posts JOIN users ON posts.user_id = users.id LEFT JOIN reactions ON posts.post_id = reactions.post_id ORDER BY RAND() DESC LIMIT " + strconv.Itoa(numOfPosts) + ";"

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
		err = rows.Scan(&post.Id, &post.Text, &post.AuthorId, &post.AuthorFirstName, &post.AuthorLastName, &post.ReactionCount, &post.ReplyCount)
		if err != nil {
			panic(err)
		}

		data = append(data, post)

		add := post.Id
		usedPosts = append(usedPosts, add)
	}
	db.Close()
	// return the x posts used and also the array of used posts so they arent printed again
	return data, usedPosts
}

// func to sort feed by random
func FeedByRandom(exclude []int) []models.Post {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/capstone", dbName, dbPass, dbHost)
	var data []models.Post

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
	query := "SELECT posts.post_id, posts.post_text, users.id AS user_id, users.first_name, users.last_name, (SELECT COUNT(*) FROM replies WHERE replies.post_id = posts.post_id) AS reply_count, (SELECT COUNT(*) FROM reactions WHERE reactions.post_id = posts.post_id) AS reaction_count FROM posts JOIN users ON posts.user_id = users.id LEFT JOIN reactions ON posts.post_id = reactions.post_id ORDER BY posts.post_id DESC;"

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
		err = rows.Scan(&post.Id, &post.Text, &post.AuthorId, &post.AuthorFirstName, &post.AuthorLastName, &post.ReactionCount, &post.ReplyCount)
		if err != nil {
			panic(err)
		}

		// check if post is already used
		currID := post.Id
		if err != nil {
			panic(err)
		}
		// get to the existence check
		_, exists := used[currID]

		// if key exists in map, skip
		if exists {
			continue
		}

		data = append(data, post)
	}
	db.Close()
	return data
}

// func to add post
func AddPost(user_id int, post_text string) bool {
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
	query := fmt.Sprintf("insert into posts (user_id, post_text) values (%d, '%s');", user_id, post_text)
	println(query)

	// execute query
	_, err = db.Query(query)

	return !(err != nil)
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

func AddReply(reply models.Reply) error {
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
	query := "INSERT INTO replies (post_id, username, reply_text) VALUES (?, ?, ?)"

	// execute sql
	_, err2 := db.Exec(query, reply.PostId, reply.UserId, reply.ReplyText)
	if err2 != nil {
		panic(err2)
	}
	return err
}

// func to get replies on a post
func GetReplies(post_id int) []models.Reply {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/capstone", dbName, dbPass, dbHost)
	var data []models.Reply

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
	query := "SELECT replier.username AS replier_username, replier.first_name AS replier_first_name, replier.last_name AS replier_last_name, replies.reply_text AS reply_text FROM posts JOIN users ON posts.user_id = users.id LEFT JOIN replies ON posts.post_id = replies.post_id LEFT JOIN users AS replier ON replies.username = replier.username WHERE posts.post_id = " + strconv.Itoa(post_id) + ";"
	// x rows of sql result
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// format each row of the result
	// if rows.NextResultSet() {
	// 	fmt.Println("we have a row")
	// } else {
	// 	fmt.Println("we don't have shit")
	// }
	for rows.Next() {
		fmt.Println("we got into rows.Next")
		var replierUsername sql.NullString
		var replierFirstName sql.NullString
		var replierLastName sql.NullString
		var replyText sql.NullString

		// scan result and set the values to each variable
		err = rows.Scan(&replierUsername, &replierFirstName, &replierLastName, &replyText)
		if err != nil {
			panic(err)
		}

		if replierUsername.Valid {
			reply := models.Reply{
				ReplierUsername:  replierUsername.String,
				ReplierFirstName: replierFirstName.String,
				ReplierLastName:  replierLastName.String,
				ReplyText:        replyText.String,
				PostId:           post_id,
			}
			data = append(data, reply)
		}
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

// func to display user info
func DisplayModel(posts []models.Post) {
	for i := 0; i < len(posts); i++ {
		println("post_id: ", posts[i].Id)
		println("post_text: ", posts[i].Text)
		println("author_id: ", posts[i].AuthorId)
		println("first_name: ", posts[i].AuthorFirstName)
		println("last_name: ", posts[i].AuthorLastName)
		println("reply_count: ", posts[i].ReplyCount)
		println("reaction_count: ", posts[i].ReactionCount, "\n")
	}
}
