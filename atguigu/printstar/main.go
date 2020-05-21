package main

import "fmt"

func printstar(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

func main() {
	var num int
	fmt.Scanln(&num)
	printstar(num)
}
