package main

import (
	"fmt"
	"math"
)

// 声明变量的两种方式
func main() {
	//显式声明：
	var a = "initial"

	var b, c int = 1, 2

	var d = true

	var e float64
	//隐式声明
	f := float32(e)

	g := a + "foo"
	//Go中不允许变量不被使用，定义了不使用会报错
	fmt.Println(a, b, c, d, e, f) // initial 1 2 true 0 0
	fmt.Println(g)                // initialapple

	//Go中的常量没有特定的类型，而是根据使用的上下文自动确定类型
	const s string = "constant"
	const h = 500000000
	const i = 3e20 / h
	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
}
