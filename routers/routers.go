package routers

import (
	"github.com/conglt-0917/gin-mysql/api"
	"github.com/gin-gonic/gin"
)

func SetUpRouters(router *gin.Engine) {
	router.POST("/login", api.Login)
	router.POST("/register", api.Register)

	router.GET("posts", api.GetAllPost)
	router.GET("/posts/{id}", api.GetPostByID)
	router.POST("/posts", api.CreatePost)
	router.PUT("/posts/{id}", api.EditPost)
	router.DELETE("/posts/{id}", api.DeletePost)

	router.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

}
