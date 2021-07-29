package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f(ctx context.Context) {
	defer wg.Done()
	go f2(ctx)
Loop:
	for {
		fmt.Println("xixi")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break Loop
		default:
		}
	}
}

func f2(ctx context.Context) {
	defer wg.Done()
Loop:
	for {
		fmt.Println("haha")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break Loop
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
}
