package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Print("lose")
	}
	defer conn.Close()
	//string 的 set
	_, _ = conn.Do("set", "name", "小星星和大脸猫")

	v, _ := redis.String(conn.Do("get", "name"))
	fmt.Print("v", v)

	//hs 哈希 的hset, hget
	_, _ = conn.Do("hmset", "user1", "name", "fanglei", "age", "18", "sex", "boy")
	v1, _ := redis.Strings(conn.Do("hgetall", "user1"))
	for _, v := range v1 {
		fmt.Printf("%v \n", v)
	}

}
