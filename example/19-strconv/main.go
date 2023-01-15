package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串转换为浮点数
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f) // 1.234

	// 转换为整型
	n, _ := strconv.ParseInt("111", 10, 64)
	fmt.Println(n) // 111

	n, _ = strconv.ParseInt("0x1000", 0, 64)
	fmt.Println(n) // 4096

	// 字符串转数字
	n2, _ := strconv.Atoi("123")
	fmt.Println(n2) // 123

	// 输入不合法则返回错误
	n2, err := strconv.Atoi("AAA")
	fmt.Println(n2, err) // 0 strconv.Atoi: parsing "AAA": invalid syntax
}
