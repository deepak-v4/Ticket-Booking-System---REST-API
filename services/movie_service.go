package services

import (
	"bookmyshow-clone/domain/booking"
	"bookmyshow-clone/domain/cinema"
	"bookmyshow-clone/domain/movies"
	"bookmyshow-clone/utils/resterrors"
)

func CreateMovie(movie movies.Movie) (*movies.Movie, *resterrors.RestErr) {
	err := movie.Save()
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func CraeteCinemaHall(hall cinema.CinemaHall) (*cinema.CinemaHall, *resterrors.RestErr) {
	err := hall.Save()
	if err != nil {
		return nil, err
	}
	return &hall, nil
}

func CreateShow(newshow movies.Show) (*movies.Show, *resterrors.RestErr) {
	err := newshow.SaveNewShow()
	if err != nil {
		return nil, err
	}

	return &newshow, nil
}

func Book_Ticket(book_ticket booking.Booking) (*booking.Booking, *resterrors.RestErr) {
	err := book_ticket.SaveBookingInfo()
	if err != nil {
		return nil, err
	}

	return &book_ticket, nil

}

func GETShowDetails() ([]movies.Show, *resterrors.RestErr) {

	var ShowInfo []movies.Show
	var show movies.Show
	ShowInfo, err := show.GETShowDetails()
	if err != nil {
		return nil, err
	}

	return ShowInfo, nil
}

func GetMovieDetail(movieId int64) (*movies.Movie, *resterrors.RestErr) {

	if movieId <= 0 {
		return nil, resterrors.NewBadRequest("invalid userid")
	}

	result := &movies.Movie{Id: movieId}
	if err := result.GETMovie(); err != nil {
		return nil, err
	}

	return result, nil

}

func GetHallDetail(hallId int64) (*cinema.CinemaHall, *resterrors.RestErr) {

	if hallId <= 0 {
		return nil, resterrors.NewBadRequest("invalid userid")
	}

	result := &cinema.CinemaHall{Id: hallId}

	if err := result.GETHall(); err != nil {
		return nil, err
	}

	return result, nil

}
