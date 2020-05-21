package main

import "fmt"

//二维 数组 外围调换，

type Biarr struct {
	sque [3][3]int
}

func (arr *Biarr) sor() {

	for i, v := range arr.sque {
		for j, _ := range v {
			if i == 0 {
				arr.sque[i][j], arr.sque[j][i] = arr.sque[j][i], arr.sque[i][j]
			} else if i == 1 && j == 2 {
				arr.sque[i][j], arr.sque[j][i] = arr.sque[j][i], arr.sque[i][j]
				return
			}
			//加一次判断，避免重复交换
		}

	}

}

//二维 数组 外围调换，
func main() {
	var arr Biarr
	arr.sque = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	arr.sor()
	fmt.Println(arr.sque)

}
