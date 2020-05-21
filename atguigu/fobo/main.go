package main

import "fmt"

//利用递归 做出 斐波那契
func main() {
	var n int = 3
	var feb int
	feb = Febo(n)
	x := F(5)
	fmt.Println(feb, x)
	fmt.Println(tao(1))
}

func Febo(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return Febo(n-1) + Febo(n-2)
	}

}

func F(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*F(n-1) + 1
	}
}

//猴子吃桃子 一天吃1半多1个，第10天只有1个了，当初有几个
func tao(n int) int {
	var taozi int
	if n == 10 {
		taozi = 1
	} else {
		taozi = (tao(n+1) + 1) * 2
	}
	return taozi
}
