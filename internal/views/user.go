package views

import (
	"Cognita_Go/internal/views/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func LogPost(c *gin.Context) {
	resp := utils.SignIn(c)

	if resp == "" {
		c.Redirect(302, "/")
		return
	}
	c.HTML(http.StatusOK, "login.html", gin.H{
		"error": resp,
	})
}

func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

func RegPost(c *gin.Context) {
	utils.SignUp(c)

	c.Redirect(302, "/")
}
