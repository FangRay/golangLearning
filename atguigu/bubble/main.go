package main

import "fmt"

func main() {
	var str [6]int = [6]int{50, 22, 25, 64, 79, 68}
	bubbleSort(&str)

	fmt.Println(str)

}

//循环冒泡
//数组要排的轮数是 长度-1次（最多次数）， 每轮排列的次数是长度-1次 再减轮数
func bubbleSort(str *[6]int) {
	lenth := len(str)
	for i := 0; i < lenth-1; i++ {
		for j := 0; j < lenth-1-i; j++ {
			if (*str)[j] > (*str)[j+1] {
				(*str)[j], (*str)[j+1] = (*str)[j+1], (*str)[j]
			}
		}
	}

}
