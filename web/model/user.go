package model

import (
	"fmt"
	"web/utils"
)

type user struct {
	userid   int
	username string
	password string
	email    string
}

func (u *user) Adduser() error {
	sqlstr := "insert into users(username,password,email) values(?,?,?)"
	_, err := utils.Db.Exec(sqlstr, u.username, u.password, u.email)
	if err != nil {
		return err
	}
	return nil
}

func (u *user) Finduser() ([]*user, error) {
	sqlstr := "select id, username, password ,email from users"
	rows, err := utils.Db.Query(sqlstr)
	if err != nil {
		return nil, err
	}

	var id int
	var username string
	var password string
	var email string
	var users []*user

	for rows.Next() {
		err := rows.Scan(&id, &username, &password, &email)
		if err != nil {
			return nil, err
		}
		u := &user{
			userid:   id,
			username: username,
			password: password,
			email:    email,
		}
		fmt.Println(u)
		users = append(users, u)

	}
	return users, nil

}
