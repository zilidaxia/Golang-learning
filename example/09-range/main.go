package main

import "fmt"

// range 关键字用于遍历数组、切片、字符串、map 和通道中的元素。
// 对于数组、切片和字符串，它将返回每个元素的索引和值。
// 对于 map，它将返回每个键和值。对于通道，它将返回每个从通道中接收到的值。
func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		if num == 2 {
			fmt.Println("index:", i, "num:", num) // index: 0 num: 2
		}
	}
	fmt.Println(sum) // 9

	m := map[string]string{"a": "A", "b": "B"}
	for k, v := range m {
		fmt.Println(k, v) // b 8; a A
	}
	for k := range m {
		fmt.Println("key", k) // key a; key b
	}
}
