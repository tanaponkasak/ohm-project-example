package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tanaponkasak/ohm-project-example/controller"
	"github.com/tanaponkasak/ohm-project-example/entity"
)

func main() {
	// เชื่อมต่อกับฐานข้อมูล
	entity.SetupDatabase()

	// สร้าง Gin engine
	r := gin.Default()

	// เปิดใช้งาน CORS middleware
	r.Use(CORSMiddleware())

	// Routes: Users
	r.GET("/users", controller.ListUsers)
	r.GET("/user/:id", controller.GetUser)
	r.POST("/users", controller.CreateUser)
	r.PATCH("/users", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)

	// เริ่มรัน server ที่ port 8080
	r.Run(":8080")
}

// CORS Middleware รองรับ frontend เช่น React (port อื่น)
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
