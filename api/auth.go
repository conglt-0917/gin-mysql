package api

import (
	"net/http"

	. "github.com/conglt-0917/gin-mysql/db"
	"github.com/conglt-0917/gin-mysql/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var validate *validator.Validate

func Login(c *gin.Context) {
	var form LoginStruct
	var user User

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate = validator.New()
	validateError := validate.Struct(form)

	if validateError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input invalid"})
		return
	}

	db, err := utils.ConnectDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	err = db.Model(&User{}).Where("username = ?", form.Username).First(&user).Error

	if err != nil {
		c.JSON(http.StatusNotExtended, gin.H{"error": "Username or Password incorrect"})
		return
	}

	err = utils.CheckPasswordHash(user.PassWord, form.Password)

	if err != nil {
		c.JSON(http.StatusNotExtended, gin.H{"error": "Username or Password incorrect"})
		return
	}

	token, errJwt := utils.CreateJWT(form.Username)

	if errJwt != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func Register(c *gin.Context) {
	// var form RegisterStruct

	// if err := c.ShouldBind(&form); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// validate = validator.New()
	// validateError := validate.Struct(form)

	// if validateError != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Input invalid"})
	// 	return
	// }
}
