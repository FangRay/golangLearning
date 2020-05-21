package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
)

func CheckUser(username string, password string) (*model.User, error) {
	sqlStr := "select id,username,password,email from users where username =? and password =?"
	row := utils.DB.QueryRow(sqlStr, username, password)
	user := &model.User{}

	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	fmt.Print(user)

	return user, nil
}

func CheckUserName(username string) (*model.User, error) {
	sqlStr := "select id,username,password,email from users where username= ?"
	row := utils.DB.QueryRow(sqlStr, username)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

func RegistUser(username string, password string, email string) error {
	sqlStr := "insert into users(username,password,email)values(?,?,?)"
	_, err := utils.DB.Exec(sqlStr, username, password, email)
	if err != nil {
		return err
	}

	return nil

}
