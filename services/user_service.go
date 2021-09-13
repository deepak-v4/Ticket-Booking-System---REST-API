package services

import (
	"bookmyshow-clone/domain/users"
	"bookmyshow-clone/utils/encrypt"
	"bookmyshow-clone/utils/resterrors"
	"strings"

	"fmt"
)

func Create(user users.User) (*users.User, *resterrors.RestErr) {

	user.Password = encrypt.GetMd5(user.Password)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	fmt.Println(user.Password)
	err := user.Save()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func LoginUser(user users.User) (*users.User, *resterrors.RestErr) {

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	user.Password = encrypt.GetMd5(user.Password)

	res := &users.User{Email: user.Email, Password: user.Password}

	err := res.Login()

	if err != nil {
		return nil, err
	}

	return res, nil

}

func GETUser(userId int64) (*users.User, *resterrors.RestErr) {
	if userId <= 0 {
		return nil, resterrors.NewBadRequest("invalid userid")
	}

	result := &users.User{Id: userId}
	if err := result.GETUser(); err != nil {
		return nil, err
	}

	return result, nil
}
