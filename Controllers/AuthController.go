package Controllers

import (
	"net/http"

	"Test/Models"
	"Test/Services"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (idb *InDB) Registrasi(c *gin.Context) {

	var (
		user   Models.User
		result gin.H
	)

	password, err := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), 14)

	if err != nil {
		result = gin.H{
			"result":  err.Error(),
			"message": "Server error",
		}
		c.JSON(http.StatusOK, result)
		return
	}

	user.Name = c.PostForm("name")
	user.Password = string(password)
	user.Email = c.PostForm("email")

	idb.DB.Create(&user)

	result = gin.H{
		"result": user,
	}
	c.JSON(http.StatusOK, result)
	return

}

func (idb *InDB) Login(c *gin.Context) {
	var (
		user   Models.User
		result gin.H
	)

	email := c.PostForm("email")
	err := idb.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		result = gin.H{
			"result":  err.Error(),
			"message": "User not found",
		}
		c.JSON(http.StatusOK, result)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.PostForm("password")))

	if err != nil {
		result = gin.H{
			"result":  err.Error(),
			"message": "Username / Password Incorect",
		}
		c.JSON(http.StatusOK, result)
		return
	}

	tokenString, err := Services.GenerateJWT(user.Email, user.Name)

	if err != nil {
		result = gin.H{
			"result":  err.Error(),
			"message": "User not found",
		}
		c.JSON(http.StatusOK, result)
		return
	}

	result = gin.H{
		"token": tokenString,
	}
	c.JSON(http.StatusOK, result)
	return
}
