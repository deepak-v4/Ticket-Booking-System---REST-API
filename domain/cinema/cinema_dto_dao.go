package cinema

import (
	userdb "bookmyshow-clone/datasources/psql/user_db"
	"bookmyshow-clone/utils/resterrors"
	"fmt"
)

type CinemaHall struct {
	Id           int64  `json:"-"`
	Hall_Name    string `json:"hall_name" binding:"required"`
	Hall_Address string `json:"hall_address" binding:"required"`
	Total_Seats  int64  `json:"total_seats" binding:"required"`
}

const (
	queryInsertCinemaHall = `INSERT INTO "cinemahall"("hall_name","hall_address","total_seats") VALUES($1,$2,$3);`
	queryGetHallDetail    = `SELECT hall_name,hall_address,total_seats from cinemahall where hall_id=$1;`
)

func (hall *CinemaHall) Save() *resterrors.RestErr {

	db := userdb.OpenConn()
	defer db.Close()
	_, savErr := db.Exec(queryInsertCinemaHall, hall.Hall_Name, hall.Hall_Address, hall.Total_Seats)
	if savErr != nil {
		fmt.Println("inside save")
		return resterrors.NewInternalServerError(savErr.Error())
	}

	return nil
}

func (hall *CinemaHall) GETHall() *resterrors.RestErr {

	db := userdb.OpenConn()
	defer db.Close()

	row, err := db.Query(queryGetHallDetail, hall.Id)

	if err != nil {
		return resterrors.NewInternalServerError(err.Error())
	}
	row.Next()
	selErr := row.Scan(&hall.Hall_Name, &hall.Hall_Address, &hall.Total_Seats)
	if selErr != nil {
		return resterrors.NewBadRequest(selErr.Error())
	}
	fmt.Println(hall)

	return nil
}
