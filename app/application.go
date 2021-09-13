package app

import (
	"bookmyshow-clone/controllers/booking"
	"bookmyshow-clone/controllers/movies"
	"bookmyshow-clone/controllers/users"

	"bookmyshow-clone/controllers/cinema"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {

	router.POST("/users", users.Create)          //create new user OR Signup
	router.GET("/users/:id", users.GETUser)      //fetch user details by id
	router.POST("/users/login", users.Login)     //login user
	router.POST("/movie", movies.Create)         //add new movie
	router.POST("/cinemahall", cinema.Create)    //add new cinema hall
	router.POST("/newshow", movies.CreateShow)   //add new show
	router.POST("/booking", booking.Book_Ticket) //book ticket
	router.POST("/users/logout", users.Logout)   //user logout

	router.GET("/getallshows", movies.GetAllShows)                //get list of all active shows
	router.GET("/getmoviedetail/:movieid", movies.GetMovieDetail) //fetch movie details
	router.GET("/gethalldetail/:hallid", cinema.GetHallDetail)    //fetch hall details

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "ok",
		})
	})

	router.Run()
}
