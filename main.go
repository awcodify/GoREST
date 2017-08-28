package main

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/whatdacode/sporta/config"
	"github.com/whatdacode/sporta/controllers"
	"github.com/whatdacode/sporta/models"
	"time"
)

func main() {
	db := config.Connect()
	db.AutoMigrate(&models.User{})

	router := gin.Default()

	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "OurRealm",
		Key:        []byte("OurSecretKey"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			if (userId == "admin" && password == "admin") || (userId == "test" && password == "test") {
				return userId, true
			}

			return userId, false
		},
		Authorizator: func(userId string, c *gin.Context) bool {
			if userId == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header:Authorization",
		TimeFunc:    time.Now,
	}

	api := router.Group("/api/v1/")
	{
		api.POST("/login", authMiddleware.LoginHandler)
		usersWithAuth := api.Group("/users")
		usersWithAuth.Use(authMiddleware.MiddlewareFunc())
		{
			usersWithAuth.GET("/", controllers.GetUsers)
			usersWithAuth.GET("/:id", controllers.GetUserDetail)
			usersWithAuth.PATCH("/:id", controllers.UpdateUserDetail)
			usersWithAuth.DELETE("/:id", controllers.DeleteUser)
		}

		users := api.Group("/users")
		{
			users.POST("/", controllers.CreateUser)
		}

	}
	router.Run()

}
