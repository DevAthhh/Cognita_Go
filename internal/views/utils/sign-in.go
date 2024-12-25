package utils

import (
	"Cognita_Go/internal/database"
	"Cognita_Go/internal/user"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) string {
	f := c.Request.FormValue
	u := user.User{Username: f("username"), Password: ToHash(f("password"))}
	us := database.SelectUsersFromDB()
	for _, v := range us {
		if v.Username == u.Username && v.Password == u.Password {
			session := sessions.Default(c)
			session.Set("username", u.Username)
			session.Save()
			return ""
		}
	}
	return "Пользователь не найден"
}

func SignUp(c *gin.Context) {
	f := c.Request.FormValue
	u := user.User{
		Username: f("name"),
		Email:    f("email"),
		Password: ToHash(f("password")),
	}
	if err := database.CreateUser(u); err != nil {
		fmt.Println(err)
		return
	}

	session := sessions.Default(c)
	session.Set("username", u.Username)
	session.Save()
}
