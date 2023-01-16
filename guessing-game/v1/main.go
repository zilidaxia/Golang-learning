package main

import (
	"fmt"
	"math/rand"
)

// 未添加随机数种子，随机数一直是81
func main() {
	maxNum := 100
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)
}
