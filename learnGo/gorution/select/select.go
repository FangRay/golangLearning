package main

import (
	"fmt"
	"math/rand"
	"time"
)

func creatChan() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func receiveChan() chan int {
	out := make(chan int)
	go func() {
		for n := range out {
			fmt.Printf("Receive Chan get value%d\n", n)
			//睡2秒，来做漏值的实验
			time.Sleep(time.Second * 1)
		}

	}()

	return out
}

/* 用select 来确保没有0值，但是会漏数据，因为n是int变量，所以请看2
	func main() {
	var c1, c2 = creatChan(), creatChan()
	var r = receiveChan()

	var hasvalue = false
	n := 0
	for {
		//确保没有0值输入
		var noZero chan int
		if hasvalue == true{
			noZero = r
		}
		select {
		case n = <-c1:
			hasvalue=true
		case n = <-c2:
			hasvalue=true
			//case 其实是执行语句哦 下面是执行语句
			//当 noZero 是nil的时候，是不会case到的
		case noZero <- n:
			hasvalue=false

		}
	}
}*/
func main() {
	var c1, c2 = creatChan(), creatChan()
	var r = receiveChan()
	//建立一个 所有取到值的切片，保存所有取到的值
	var value []int
	var n int
	//自动 时间 停止
	tm := time.After(time.Second * 10)
	tick := time.Tick(time.Second * 1)
	for {
		//确保没有0值输入
		var noZero chan int
		var firstvalue int
		if len(value) > 0 {
			noZero = r //有值的时候才会把生成的chan发给他，要不一直打0
			firstvalue = value[0]
		}
		select {
		case n = <-c1:
			value = append(value, n)
		case n = <-c2:
			value = append(value, n)
			//case 其实是执行语句哦 下面是执行语句
			//当 noZero 是nil的时候，是不会case到的
		case noZero <- firstvalue:
			value = value[1:]
		case <-tm:
			fmt.Println("all the values, goodbye, welldone my boy! ", value)
			return // close the programe after 10 seconds,with the time chan from main fucntion
		case <-time.After(time.Millisecond * 800):
			fmt.Println("time out") //check whether take 800MS between two select
		case <-tick:
			fmt.Println("number of values", len(value))
		}

	}
}
