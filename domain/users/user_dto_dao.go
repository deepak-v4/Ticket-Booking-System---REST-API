package users

import (
	userdb "bookmyshow-clone/datasources/psql/user_db"
	"bookmyshow-clone/utils/resterrors"
	"fmt"
	"strings"
)

type User struct {
	Id       int64  `json:"-"`
	Name     string `json:"user_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

const (
	queryInsertUser = `INSERT INTO "users"("user_name","user_email","phone","password") VALUES($1,$2,$3,$4);`
	querLoginUser   = `SELECT user_id,user_name,user_email,phone from users where user_email=$1 AND password=$2;`
	QueryUserDetail = `SELECT user_name,user_email,phone from users where user_id=$1;`
)

func (user *User) Validate() {

	user.Name = strings.TrimSpace(strings.ToUpper(user.Name))
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

}

func (user *User) Save() *resterrors.RestErr {

	db := userdb.OpenConn()
	defer db.Close()
	_, savErr := db.Exec(queryInsertUser, user.Name, user.Email, user.Phone, user.Password)
	if savErr != nil {
		return resterrors.NewInternalServerError(savErr.Error())
	}

	return nil
}

func (user *User) Login() *resterrors.RestErr {

	db := userdb.OpenConn()
	defer db.Close()

	fmt.Println(user)
	row, err := db.Query(querLoginUser, user.Email, user.Password)

	if err != nil {
		return resterrors.NewInternalServerError(err.Error())
	}
	row.Next()
	selErr := row.Scan(&user.Id, &user.Name, &user.Email, &user.Phone)
	if selErr != nil {
		return resterrors.NewBadRequest(selErr.Error())
	}
	fmt.Println(user)
	return nil
}

func (user *User) GETUser() *resterrors.RestErr {

	db := userdb.OpenConn()
	defer db.Close()

	row, err := db.Query(QueryUserDetail, user.Id)

	if err != nil {
		return resterrors.NewInternalServerError(err.Error())
	}
	row.Next()
	selErr := row.Scan(&user.Name, &user.Email, &user.Phone)
	if selErr != nil {
		return resterrors.NewBadRequest(selErr.Error())
	}
	fmt.Println(user)
	return nil

}
