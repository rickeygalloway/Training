package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- 2
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 110*time.Millisecond)

	defer cancel()

	select {
	case v := <-ch1:
		fmt.Printf("got %d from ch1", v)
	case v := <-ch2:
		fmt.Printf("got %d from ch2", v)
	case <-ctx.Done():
		fmt.Println("timeout")
		//better to use ctx/cancel
		//case <-time.After(1 * time.Millisecond): //indicates a timeout. Go routines are still running though
		//	fmt.Println("timeout")
	}

}
