package main

import (
	"fmt"
	"time"
)

// switch语句中的变量可以是任何类型，可以是一个常量或者一个变量，在每个case后面可以跟多个值，用逗号隔开

// Go中的switch语句也支持在条件判断语句前执行简单的语句，这样就可以在一个switch语句中处理多个条件

// 不使用break 且不会像c一样继续向下执行，而是直接跳出
func main() {

	a := 4
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	//多个值用逗号隔开
	case 4, 5:
		fmt.Println("four or five")
	default:
		fmt.Println("other")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}
