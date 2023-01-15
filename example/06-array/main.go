package main

import "fmt"

// 数组和其他语言相同 但须注意数组定义方式
func main() {
	// 显式声明
	var a [5]int
	a[4] = 100
	fmt.Println("get:", a[2])   //get: 0
	fmt.Println("len:", len(a)) //len: 5
	// 隐式声明
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
