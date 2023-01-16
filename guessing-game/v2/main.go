package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 将程序启动的时间戳作为随机数种子
func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)
}
