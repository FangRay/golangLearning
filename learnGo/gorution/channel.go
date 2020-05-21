package main

import (
	"fmt"
	"time"
)

/*----------------------------------------
go roution 闭包直接使用chan数据
func demochan()  {
	c:=make(chan int)
	go func() {
		for {
			n:= <-c
			fmt.Println(n)
		}
	}()
	for i:=0; i<10; i++ {
		c <- i

	}
	time.Sleep(time.Minute * 3)
}*/
//-------------------------------------------
// 把chan 当成参数，放入 goroutin中
/*
func worker( c chan int)  {
	for {
		n:= <- c
		fmt.Printf("you get number from C is %d \n",n)
	}
}
func demochan()  {
	c:=make(chan int)
	// 一定要开协程，也就是fun前加go
	go worker(c)
	for i:=0; i<10; i++{
		c <- i
	}
	time.Sleep(time.Second * 2)
}
*/

/*
// chan as first class Citizen 当成参数
func workers(i int, c chan int)  {
	for {
		fmt.Printf("Woker %d is carring Number %d\n",i, <-c)
	}
}
func demochan()  {
	//chan 管道的数组和切片只能用 var来定义
	var chans [10]chan int
	for i:=0; i<10; i++{
		chans[i] = make(chan int)
		go workers(i,chans[i])
	}
	for i:=0; i<10; i++{
		for j:=100; j<150; j++{
			chans[i] <- j
		}
	}
	time.Sleep(time.Second*2)
}*/

//学用建立channel 的方法 生成函数直接返回chan
func creatWoker(i int) chan<- int {
	c := make(chan int)
	go func() {
		//用for range来 接受有关闭指令的chan数据
		for range c {
			fmt.Printf("Woker %d is getting value %d \n", i, <-c)
		}
	}()
	return c
}

func demochan() {
	//chan 管道的数组和切片只能用 var来定义
	var chans [10]chan<- int
	for i := 0; i < 10; i++ {
		chans[i] = creatWoker(i)
	}
	for i := 0; i < 10; i++ {
		for j := 100; j < 150; j++ {
			chans[i] <- j
		}
	}
	time.Sleep(time.Second * 2)
}

func main() {
	fmt.Print("hello")
	demochan()
}
