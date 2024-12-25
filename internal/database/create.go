package database

import (
	"Cognita_Go/internal/user"
	"fmt"
)

func CreateEntries(title, content string, author int) error {
	db := ConnectToDB()
	defer func() {
		_ = db.Close()
	}()

	_, err := db.Exec(fmt.Sprintf("INSERT INTO posts (Title, Content, Author) VALUES ('%s', '%s',%d)", title, content, author))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func CreateUser(u user.User) error {
	db := ConnectToDB()
	defer func() {
		_ = db.Close()
	}()
	_, err := db.Exec(fmt.Sprintf("INSERT INTO users (Email, Name, Password) VALUES ('%s', '%s', '%s')", u.Email, u.Username, u.Password))
	if err != nil {
		return err
	}
	return nil
}
