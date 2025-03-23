package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("num go routines 1:", runtime.NumGoroutine())

	go func() {
		num := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				num++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", num)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("num go routines 2:", runtime.NumGoroutine())

	fmt.Println("about to cancel context")
	cancel()
	fmt.Println("canceled context")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("num go routines 3:", runtime.NumGoroutine())

}
