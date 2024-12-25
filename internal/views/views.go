package views

import (
	"Cognita_Go/internal/database"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	data := database.SelectDataFromDB()
	u := fmt.Sprintf("%s", username)
	if username == nil {
		u = "Noname"
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Posts":    data,
		"Username": u,
	})
}

func AddPost(c *gin.Context) {
	f := c.Request.FormValue
	_ = database.CreateEntries(f("Title"), f("Content"), 1)
	c.Redirect(302, "/")
}

func Create(c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", gin.H{})
}
