package booking

import (
	userdb "bookmyshow-clone/datasources/psql/user_db"
	"bookmyshow-clone/utils/resterrors"
	"math/rand"
	"time"
)

type Booking struct {
	BookingId   int64     `json:"-"`
	SeatCount   int64     `json:"seat_count" binding:"required"`
	BookedAt    time.Time `json:"booked_at" binding:"required"`
	UserId      int64     `json:"user_id" binding:"required"`
	ShowId      int64     `json:"show_id" binding:"required"`
	Status      int32     `json:"status" binding:"required"`
	BookingCode string    `json:"booking_code" binding:"required"`
}

const (
	queryBookingInfo = `INSERT INTO "ticket_booking"("seat_count","user_id","show_id","status","booking_code") VALUES($1,$2,$3,$4,$5);`
)

func (book_ticket *Booking) SaveBookingInfo() *resterrors.RestErr {

	booking_code := Generate_Code()
	db := userdb.OpenConn()
	defer db.Close()

	_, savErr := db.Exec(queryBookingInfo, book_ticket.SeatCount, book_ticket.UserId, book_ticket.ShowId, book_ticket.Status, booking_code)

	if savErr != nil {
		return resterrors.NewInternalServerError(savErr.Error())
	}
	return nil
}

func Generate_Code() string {
	var aplha = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	code := make([]rune, 8)
	for i := range code {
		code[i] = aplha[rand.Intn(len(aplha))]
	}
	return string(code)
}
