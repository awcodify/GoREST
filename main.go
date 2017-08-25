package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/whatdacode/sporta/config"
	"github.com/whatdacode/sporta/controllers"
	"github.com/whatdacode/sporta/models"
)

func main() {

	//Migrate the schema
	db := config.Connect()
	db.AutoMigrate(&models.User{})

	router := gin.Default()

	api := router.Group("/api/v1/")
	{
		users := api.Group("/users")
		{
			users.GET("/", controllers.GetUsers)
			users.POST("/", controllers.CreateUser)
			users.GET("/:id", controllers.GetUserDetail)
			users.PATCH("/:id", controllers.UpdateUserDetail)
			users.DELETE("/:id", controllers.DeleteUser)
		}

	}
	router.Run()

}
