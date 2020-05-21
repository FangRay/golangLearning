package main

import "fmt"

//二分法查找，递归使用，结束句是一个条件满足直接return; 然后参数中要插入查找的坐标
func binaryFind(arr *[10]int, left int, right int, num int) {
	if left > right {
		fmt.Println("not find")
		return
	}
	middle := (left + right) / 2
	switch {
	case (*arr)[middle] == num:
		fmt.Println("got it")

	case (*arr)[middle] > num:
		binaryFind(arr, left, middle-1, num)

	default:
		binaryFind(arr, left+1, middle, num)

	}
}

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	binaryFind(&arr, 0, 9, 25)

}
