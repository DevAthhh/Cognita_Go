package handlers

import (
	"Cognita_Go/internal/views"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var (
	store = cookie.NewStore([]byte("tEy_7kylWJsDPlMfT5G69f0LFFQeUMtavA07DlwbrFY"))
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(sessions.Sessions("mysession", store))

	router.LoadHTMLGlob("../internal/web/templates/*")
	router.Static("/css", "../internal/web/static/css/")

	router.GET("/", views.Index)
	router.POST("/", views.AddPost)
	router.GET("/create", views.Create)

	auth := router.Group("/u")
	{
		auth.GET("/login", views.Login)
		auth.GET("/reg", views.Register)
		auth.POST("/login", views.LogPost)
		auth.POST("/reg", views.RegPost)
	}

	router.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.Redirect(302, "/")
	})

	return router
}
