package main

import (
	"fmt"
)

//取消sleep，用waitgoup来表演示结束, 添加wg

/*func creatWoker(i int, wg *sync.WaitGroup) chan <-int  {
	c:=make(chan int)
	go func() {
		//用for range来 接受有关闭指令的chan数据
		for n:= range c{
			fmt.Printf("Woker %d is getting value %d \n",i,n)
			wg.Done()
		}
	}()
	return c
}

func demochan()  {
	//chan 管道的数组和切片只能用 var来定义
	var chans [10]chan <-int
	//waitgroup 只能创建值，再传递指针
	var wg sync.WaitGroup
	for i:=0; i<10; i++{
		chans[i] = creatWoker(i,&wg)
	}
	wg.Add(10)
	for i:=0; i<10; i++{
			chans[i] <- i+1
	}
	wg.Wait()
	//time.Sleep(time.Second*2) 不再需要sleep来控制了
}*/

//加一个channel 来通知结束，而不是使用wait Group
func creatWoker2(i int, done chan bool) chan<- int {
	c := make(chan int)
	go func() {
		//用for range来 接受有关闭指令的chan数据
		for n := range c {
			fmt.Printf("Woker %d is getting value %d \n", i, n)
			done <- true
		}
	}()
	return c
}

func demochan2() {

	var chans [10]chan<- int
	done := make(chan int)
	// 开一个协程，不停的取done协程
	go func() {
		for {
			<-done
		}
	}()

	for i := 0; i < 10; i++ {
		chans[i] = creatWoker2(i, done)
	}

	for i := 0; i < 10; i++ {
		chans[i] <- i + 1

	}

	for i := 0; i < 10; i++ {
		chans[i] <- i + 2
	}

}

func main() {
	fmt.Print("hello")
	//demochan() 用waitgroup
	demochan2()
}
