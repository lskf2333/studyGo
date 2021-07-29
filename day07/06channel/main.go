package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var b chan int
var x int

func noBufChannel() {
	b := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goruntine从通道b中取到了", x)
	}()
	b <- 10
	fmt.Println("10发送到通道中了")
	wg.Wait()
	close(b)
}

func bufChannel() {
	b := make(chan int, 1)
	b <- 10
	fmt.Println("10发送到通道中了")
	x = <-b
	fmt.Println("从通道b中了去到了", x)
	b <- 20
	fmt.Println("20发送到通道中了")
	x = <-b
	fmt.Println("从通道b中了去到了", x)
	close(b)
}

func main() {
	bufChannel()
}
