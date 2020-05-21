package main

import (
	"fmt"
)

func main() {
	//闭包函数切片，传入的是地址
	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		fs[i] = func() { fmt.Println("hello", i) }
	}

	//打印结果是相同
	for _, f := range fs {
		f()
	}

}
