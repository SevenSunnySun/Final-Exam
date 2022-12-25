package main

import (
	"fmt"
	"time"
)

func main() {
	success := make(chan int)
	go Work("goroutine1", success)
	<-success
	go Work("goroutine2", success)
	<-success
	go Work("goroutine3", success)
	<-success
	fmt.Println("successful")
}
func Work(workName string, output chan int) {
	time.Sleep(time.Second) // 模拟逻辑处理
	fmt.Println(workName)
	output <- 0
}
