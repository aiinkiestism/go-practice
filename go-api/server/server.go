package server

import (
	"go/src/go-api/controllers"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	u := r.Group("/users")
	{
		ctrl := controllers.UserControllers{}
		u.GET("", ctrl.Index)
		u.POST("", ctrl.Create)
		u.GET("/:id", ctrl.Show)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	p := r.Group("/posts")
	{
		ctrl := controllers.PostControllers{}
		p.GET("", ctrl.Index)
		p.POST("", ctrl.Create)
		p.GET("/:id", ctrl.Show)
		p.PUT("/:id", ctrl.Update)
		p.DELETE("/:id", ctrl.Delete)
	}

	return r
}
