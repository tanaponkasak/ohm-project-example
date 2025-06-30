package main

import (
	"github.com/tanaponkasak/ohm-project-example/controller"
	"github.com/tanaponkasak/ohm-project-example/entity"
	"github.com/tanaponkasak/ohm-project-example/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group(""){
		protected := api.Use(middlewares.Authorizes())
		{

			// User Routes
			protected.GET("/users", controller.ListUsers)
			protected.GET("/user/:id", controller.GetUser)
			protected.POST("/users", controller.CreateUser)
			protected.PATCH("/users", controller.UpdateUser)
			protected.DELETE("/users/:id", controller.DeleteUser)
		}
		
		r.POST("/users", controller.CreateUser)

		// Run server
		r.Run()
	}
}
