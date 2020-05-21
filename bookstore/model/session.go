package model

import (
	"bookstore/utils"
	"fmt"
)

type Session struct {
	Sess     string
	Username string
	Userid   int
}

func (s *Session) Add() error {
	sqlstr := "insert into sessions values(?,?,?)"
	_, err := utils.DB.Exec(sqlstr, s.Sess, s.Username, s.Userid)
	if err != nil {
		return err
	}
	fmt.Print("插入session成功")
	return nil
}

// func (s *Session) Check() (*Session, error) {
// 	// fmt.Println("开始检查登陆")
// 	// fmt.Printf("26行要查询的Session是%v", s.Sess)
// 	sqlstr := "select sess,username,userid from sessions where sess = ?"
// 	//要注意queryrow 和 query的区别哦
// 	row := utils.DB.QueryRow(sqlstr, s.Sess)
// 	fmt.Print(row)
// 	n := &Session{}
// 	row.Scan(&n.Sess, &n.Username, &n.Userid)
// 	fmt.Printf("新生成的session是%v", n)
// 	return n, nil
//}

func IsLogin(sess string) (bool, string, int) {
	sqlstr := "select username ,userid from sessions where sess = ?"
	row := utils.DB.QueryRow(sqlstr, sess)
	var username string
	var userid int
	err := row.Scan(&username, &userid)
	//fmt.Printf("42行username的值是v%\n", username)
	if err != nil || username == "" || userid == 0 {
		println("登陆没有数据")
		return false, "", 0
	}
	//fmt.Printf("48行登陆信息，用户ID为%v\n", userid)
	return true, username, userid

}

func (s *Session) Quit() {
	fmt.Println(s.Sess)
	sqlstr := "delete from sessions where sess=?"
	utils.DB.Exec(sqlstr, s.Sess)
	fmt.Print("退出删除session成功")
}
