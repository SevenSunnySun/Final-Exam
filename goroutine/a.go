package main

import "fmt"

func main() {
	success := make(chan int)
	go func() {
		fmt.Println("出现")
		success <- 0
	}()
	<-success
}
