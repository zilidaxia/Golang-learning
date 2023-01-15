package main

import "fmt"

func main() {

	i := 1
	//for中什么都不写就是死循环
	for {
		fmt.Println("loop")
		break
	}
	//循环体的小括号省略了
	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}
	//也可通过continue继续循环或break跳出
	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	//
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
}

// loop
// 7
// 8
// 1
// 3
// 1
// 2
// 3
