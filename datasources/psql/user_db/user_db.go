package userdb

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var (
	Client *sql.DB
)

func OpenConn() *sql.DB {

	db, err := sql.Open("postgres", "postgres://kpcnlhoqpwmwlx:ta56b#ema0f92da2b52d7818d47c3b93de831fcf11ebe8873ef0d5062f322a350d201b45e@ec2-44-198-80-194.compute-1.amazonaws.com:5432/d44ipbgfo8hvs7")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
	//defer db.Close()

	/*	_, e := db.Exec(queryInsertUser, "temp", "temp@gmail.com", "89765678990", "hello")
		if e != nil {
			log.Fatal(e)
		}
	*/
	//	fmt.Println("db Connected !!")
}
