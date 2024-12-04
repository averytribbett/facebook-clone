package profile

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"fakebook.com/project/models"
)



func CheckAdmin(adminId int) bool{

	var admin bool
	var adminUsername string

	err:= db.QueryRow("SELECT username FROM admins WHERE user_id = ?",adminId).Scan(adminUsername)

	if err != nil{

		log.Println(err)
		admin = false
	}
	if len(adminUsername)>0 {
		admin = true
	}


	return admin
}

func MakeUserAdmin(userId int, adminId int){

	var username string

	if CheckAdmin(adminId){

		err1 := db.QueryRow("SELECT username, bio FROM users WHERE id = ?", userId).Scan(username)

		if err1 != nil {
			log.Println(err1)
		}

		query :=  ("INSERT INTO admins (username, user_id) VALUES (?, ?)")

		_, err2 := db.Exec(query,username,userId)

		if err2 != nil {
			log.Println(err2)
		}

	}

}

func DeleteUserAdmin(username string, adminId int) error {

	if CheckAdmin(adminId){
		//Beginning transactions for the database, so they can be rolled back if an error occurs midway.
		txn, err := db.Begin()

		if err != nil {
			return err
		}

		// deferring the function to either commit the transactions, or roll them back depending on if an error is thrown.
		defer func() {
			if err != nil {
				txn.Rollback()
			} else {
				err = txn.Commit()
			}
		}()

		var userID int
		// getting username to delete from Friends table
		err = txn.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&userID)

		// removing from reactions table
		_, err = txn.Exec("DELETE FROM reactions WHERE user_id = ?", userID)
		if err != nil {
			return err
		}

		// removing from replies table
		_, err = txn.Exec("DELETE FROM replies WHERE username = ?", username)

		if err != nil {
			return err
		}

		// removing from posts table
		_, err = txn.Exec("DELETE FROM posts WHERE user_id = ?", userID)

		if err != nil {
			return err
		}

		// removing from friends table
		_, err = txn.Exec("DELETE FROM friends WHERE user_id = ? or friend_id =?", username)

		if err != nil {
			return err
		}

		// Deleting the user
		_, err = txn.Exec("DELETE FROM users WHERE username = ?", username)

		if err != nil {
			return err
		}

		return nil
	}

	return nil
}



func AddNewUserAdmin(newUser models.User, adminId int) error {

	if CheckAdmin(adminId){

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
	}else{
		return nil
	}
}



