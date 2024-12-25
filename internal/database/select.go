package database

import (
	"Cognita_Go/internal/user"
	_ "github.com/lib/pq"
)

func SelectDataFromDB() []Post {
	db := ConnectToDB()
	defer func() {
		_ = db.Close()
	}()

	res, _ := db.Query("SELECT * FROM posts")
	posts := []Post{}

	for res.Next() {
		p := Post{}
		err := res.Scan(&p.Id, &p.Title, &p.Content, &p.Author)
		if err != nil {
			break
		}
		posts = append(posts, p)
	}

	return posts
}

func SelectUsersFromDB() []user.User {
	db := ConnectToDB()
	defer func() {
		_ = db.Close()
	}()

	res, _ := db.Query("SELECT * FROM users")
	users := []user.User{}

	for res.Next() {
		u := user.User{}
		err := res.Scan(&u.Id, &u.Email, &u.Username, &u.Password)
		if err != nil {
			break
		}
		users = append(users, u)
	}

	return users
}
