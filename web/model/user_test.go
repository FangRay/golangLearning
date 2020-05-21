package model

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	//t.Run("I'm the king", testAdduser)
	t.Run("找喾", testFinduser)
}

func testAdduser(t *testing.T) {
	user := &user{
		username: "kaixin6",
		password: "youbadboy3",
		email:    "cool@yahoo.com",
	}
	user.Adduser()

}

func testFinduser(t *testing.T) {
	fmt.Printf("开始测试查找学生")
	user := &user{}
	users, _ := user.Finduser()
	for k, v := range users {
		fmt.Printf("第%v个学生的资料是%v", k+1, v)
	}

}
