package users

import (
	"bookmyshow-clone/domain/users"
	"bookmyshow-clone/services"
	"bookmyshow-clone/utils/resterrors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {

	var user users.User

	err := c.ShouldBindJSON(&user)
	if err != nil {

		restErr := resterrors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)

		return
	}

	fmt.Println(user)

	result, savErr := services.Create(user)

	if savErr != nil {
		c.JSON(savErr.Status, savErr)
	}

	c.JSON(http.StatusCreated, result)
}

func Login(c *gin.Context) {

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := resterrors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, logErr := services.LoginUser(user)
	if logErr != nil {
		c.JSON(logErr.Status, logErr)
		return
	} //else {
	//set cookies OR implement oauth-token
	//}

	c.JSON(http.StatusOK, result)

}

func GETUser(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if userErr != nil {
		err := resterrors.NewBadRequest("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GETUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)

}

func Logout(c *gin.Context) {
	//to-do
	//delete cookie OR oauth-token

	c.JSON(200, gin.H{
		"msg": "to-do",
	})
}
