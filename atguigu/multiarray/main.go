package main

import "fmt"

//二维数组必须用指针，才能使用for range来赋值
func valueArr(arr *[4][5]int) {
	for i, _ := range *arr {
		//fmt.Print(i)
		for j, _ := range (*arr)[i] {
			(*arr)[i][j] = i*10 + j
			fmt.Print((*arr)[i][j])
		}
		fmt.Println()

	}

}

// func valueArr(arr [4][5]int) {
// 	for i, _ := range arr {
// 		//fmt.Print(i)
// 		for j, _ := range arr[i] {
// 			arr[i][j] = i*10 + j
// 			fmt.Print(arr[i][j])
// 		}
// 		fmt.Println()

// 	}

// }

func main() {

	var marr [4][5]int
	var marr1 = [4][5]int{}
	//为什么无法赋值
	valueArr(&marr)
	//valueArr(marr)
	fmt.Print(marr1)
	fmt.Println()
	fmt.Print(marr)
}
