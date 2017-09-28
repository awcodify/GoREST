package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/whatdacode/GoREST/config"
	"github.com/whatdacode/GoREST/controllers"
	"github.com/whatdacode/GoREST/database"
	"github.com/whatdacode/GoREST/models"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	database.Migrations()

	router := gin.Default()

	router.Use(RequestLogger())

	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "OurRealm",
		Key:        []byte("OurSecretKey"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(email string, password string, c *gin.Context) (string, bool) {
			var user models.User
			db := config.Connect()
			db.First(&user, "email = ?", email)

			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
				return email, false
			}

			return email, true
		},
		// Authorizator: func(userId string, c *gin.Context) bool {
		// 	if userId == "admin" {
		// 		return true
		// 	}

		// 	return false
		// },
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

// RequestLogger is used for logging each http request in our API.
// thanks to https://stackoverflow.com/users/3011570/emb
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because rdr1 will be read.

		fmt.Println(readBody(rdr1)) // Print request body

		c.Request.Body = rdr2
		c.Next()
	}
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}
