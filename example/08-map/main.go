package main

import "fmt"

// Go语言中的map是一种内置数据结构，用于存储键值对。
// 键必须是可比较的数据类型（如字符串、整数等），值可以是任何数据类型。
// map可用于统计字符串中单词出现次数或维护用户信息数据库等任务。
// 可以使用make函数创建新的map

// 在其他语言中，他可能叫做哈希或者字典
// map遍历时完全无序的，不会按照字母或者插入顺序

func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)           // map[one:1 two:2]
	fmt.Println(len(m))      // 2
	fmt.Println(m["one"])    // 1
	fmt.Println(m["unknow"]) // 0

	// 当从一个 map 中取值时，还有可以选择是否接收的第二个返回值，该值表明了 map 中是否存在这个键。 这可以用来消除 键不存在 和 键的值为零值 产生的歧义， 例如 0 和 ""。
	a, ok := m["o"]
	fmt.Println(a, ok) // 0 false

	delete(m, "one")

	m2 := map[string]int{"one": 1, "two": 2}
	var m3 = map[string]int{"one": 1, "two": 2}
	fmt.Println(m2, m3)
}
