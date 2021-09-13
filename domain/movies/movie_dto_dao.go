package movies

import (
	userdb "bookmyshow-clone/datasources/psql/user_db"
	"bookmyshow-clone/utils/resterrors"
	"fmt"
	"time"
)

//RFC 3339

type Movie struct {
	Id            int64         `json:"-"`
	MovieName     string        `json:"movie_name" binding:"required"`
	Description   string        `json:"description" binding:"required"`
	Duration      time.Duration `json:"duration" binding:"required"`
	Released_date time.Time     `json:"released_date" binding:"required"`
	Language      string        `json:"language" binding:"required"`
}

type Show struct {
	Id            int64     `json:"-"`
	ShowDate      time.Time `json:"show_date" binding:"required"`
	ShowStartTime time.Time `json:"show_start_time" binding:"required"`
	ShowEndTime   time.Time `json:"show_end_time" binding:"required"`
	Movie_Id      int64     `json:"movie_id" binding:"required"`
	Hall_Id       int64     `json:"hall_id" binding:"required"`
}

const (
	queryInsertMovie    = `INSERT INTO "movie"("movie_name","description","duration","released_date","language") VALUES($1,$2,$3,$4,$5);`
	queryInsertShow     = `INSERT INTO "show"("show_date","show_starttime","show_endtime","movie_id","hall_id") VALUES($1,$2,$3,$4,$5);`
	queryGetAllShow     = `SELECT show_id,show_date,show_starttime,show_endtime,movie_id,hall_id from show;`
	queryGetMovieDetail = `SELECT movie_name,description,duration,released_date,language from movie where movie_id=$1;`
)

func (movie *Movie) Save() *resterrors.RestErr {

	db := userdb.OpenConn()
	defer db.Close()
	_, savErr := db.Exec(queryInsertMovie, movie.MovieName, movie.Description, movie.Duration, movie.Released_date, movie.Language)
	if savErr != nil {
		return resterrors.NewInternalServerError(savErr.Error())
	}

	return nil
}

func (newshow *Show) SaveNewShow() *resterrors.RestErr {

	db := userdb.OpenConn()
	defer db.Close()

	_, savErr := db.Exec(queryInsertShow, newshow.ShowDate, newshow.ShowStartTime, newshow.ShowEndTime, newshow.Movie_Id, newshow.Hall_Id)

	if savErr != nil {
		return resterrors.NewInternalServerError(savErr.Error())
	}
	return nil
}

var ShowInfo []Show

func (show *Show) GETShowDetails() ([]Show, *resterrors.RestErr) {

	db := userdb.OpenConn()
	defer db.Close()

	row, err := db.Query(queryGetAllShow)

	if err != nil {
		return nil, resterrors.NewInternalServerError(err.Error())
	}

	for row.Next() {
		selErr := row.Scan(&show.Id, &show.ShowDate, &show.ShowStartTime, &show.ShowEndTime, &show.Movie_Id, &show.Hall_Id)
		if selErr != nil {
			return nil, resterrors.NewBadRequest(selErr.Error())
		} else {

			var temp Show
			temp.Id = show.Id
			temp.ShowDate = show.ShowDate
			temp.ShowStartTime = show.ShowStartTime
			temp.ShowEndTime = show.ShowEndTime
			temp.Movie_Id = show.Movie_Id
			temp.Hall_Id = show.Hall_Id

			ShowInfo = append(ShowInfo, temp)
		}
	}
	return ShowInfo, nil

}

func (movie *Movie) GETMovie() *resterrors.RestErr {

	db := userdb.OpenConn()
	defer db.Close()

	row, err := db.Query(queryGetMovieDetail, movie.Id)

	if err != nil {
		return resterrors.NewInternalServerError(err.Error())
	}
	row.Next()
	selErr := row.Scan(&movie.MovieName, &movie.Description, &movie.Duration, &movie.Released_date, &movie.Language)
	if selErr != nil {
		return resterrors.NewBadRequest(selErr.Error())
	}
	fmt.Println(movie)

	return nil
}
