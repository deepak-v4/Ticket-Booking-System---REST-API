package booking

import (
	"bookmyshow-clone/domain/booking"
	"bookmyshow-clone/services"
	"bookmyshow-clone/utils/resterrors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Book_Ticket(c *gin.Context) {

	var book_ticket booking.Booking

	err := c.ShouldBindJSON(&book_ticket)
	if err != nil {

		restErr := resterrors.NewBadRequest("invalid json body" + err.Error())
		c.JSON(restErr.Status, restErr)

		return
	}

	fmt.Println(book_ticket)

	result, savErr := services.Book_Ticket(book_ticket)

	if savErr != nil {
		c.JSON(savErr.Status, savErr)
	}

	c.JSON(http.StatusCreated, result)

}
