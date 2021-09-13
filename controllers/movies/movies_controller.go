package movies

import (
	"bookmyshow-clone/domain/movies"
	"bookmyshow-clone/services"
	"bookmyshow-clone/utils/resterrors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {

	var movie movies.Movie

	err := c.ShouldBindJSON(&movie)
	if err != nil {

		restErr := resterrors.NewBadRequest("invalid json body" + err.Error())
		c.JSON(restErr.Status, restErr)

		return
	}

	fmt.Println(movie)

	result, savErr := services.CreateMovie(movie)

	if savErr != nil {
		c.JSON(savErr.Status, savErr)
	}

	c.JSON(http.StatusCreated, result)

}

func CreateShow(c *gin.Context) {

	var newshow movies.Show

	err := c.ShouldBindJSON(&newshow)
	if err != nil {

		restErr := resterrors.NewBadRequest("invalid json body" + err.Error())
		c.JSON(restErr.Status, restErr)

		return
	}

	fmt.Println(newshow)

	result, savErr := services.CreateShow(newshow)

	if savErr != nil {
		c.JSON(savErr.Status, savErr)
	}

	c.JSON(http.StatusCreated, result)

}

func GetAllShows(c *gin.Context) {

	var ShowInfo []movies.Show

	ShowInfo, getErr := services.GETShowDetails()
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, ShowInfo)

}

func GetMovieDetail(c *gin.Context) {

	movieId, movieErr := strconv.ParseInt(c.Param("movieid"), 10, 64)
	if movieErr != nil {
		err := resterrors.NewBadRequest("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	movie, getErr := services.GetMovieDetail(movieId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, movie)

}
