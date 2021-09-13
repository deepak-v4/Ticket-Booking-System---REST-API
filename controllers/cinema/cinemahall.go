package cinema

import (
	"bookmyshow-clone/domain/cinema"
	"bookmyshow-clone/services"
	"bookmyshow-clone/utils/resterrors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {

	var cinemaHall cinema.CinemaHall

	err := c.ShouldBindJSON(&cinemaHall)
	if err != nil {

		restErr := resterrors.NewBadRequest("invalid json body" + err.Error())
		c.JSON(restErr.Status, restErr)

		return
	}

	fmt.Println(cinemaHall)

	result, savErr := services.CraeteCinemaHall(cinemaHall)

	if savErr != nil {
		c.JSON(savErr.Status, savErr)
	}

	c.JSON(http.StatusCreated, result)

}

func GetHallDetail(c *gin.Context) {

	hallId, hallErr := strconv.ParseInt(c.Param("hallid"), 10, 64)
	if hallErr != nil {
		err := resterrors.NewBadRequest("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	movie, getErr := services.GetHallDetail(hallId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, movie)

}
