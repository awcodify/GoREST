package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/whatdacode/GoREST/config"
	"github.com/whatdacode/GoREST/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
)

func CreateUser(c *gin.Context) {
	hash, err := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	user := models.User{Username: c.PostForm("username"), Password: string(hash), Email: c.PostForm("email")}

	if errs := validator.Validate(user); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": errs})
		return
	}

	db := config.Connect()
	db.Save(&user)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User successfully created!", "resourceId": user.ID})
}

func GetUsers(c *gin.Context) {
	var users []models.User
	var _users []models.TransformedUser

	db := config.Connect()
	db.Find(&users)

	for _, item := range users {
		_users = append(_users, models.TransformedUser{ID: item.ID, Username: item.Username, Email: item.Email})
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _users, "count": len(_users)})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	userId := c.Param("id")
	db := config.Connect()
	db.First(&user, userId)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User not found!"})
		return
	}

	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User successfully removed!"})
}

func GetUserDetail(c *gin.Context) {
	var user models.User
	userId := c.Param("id")

	db := config.Connect()
	db.First(&user, userId)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User not found!"})
		return
	}

	db.Find(&user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}

func UpdateUserDetail(c *gin.Context) {
	var user models.User
	userId := c.Param("id")
	db := config.Connect()
	db.First(&user, userId)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User not found!"})
		return
	}

	db.Model(&user).Update("username", c.PostForm("username"))
	db.Model(&user).Update("password", c.PostForm("password"))
	db.Model(&user).Update("email", c.PostForm("email"))
	db.Model(&user).Update("firstname", c.PostForm("firstname"))
	db.Model(&user).Update("lastname", c.PostForm("lastname"))
	db.Model(&user).Update("phone", c.PostForm("phone"))
	db.Model(&user).Update("address", c.PostForm("address"))

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User updated successfully!"})
}
