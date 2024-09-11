package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func generator(ctx context.Context, num int) <-chan int {
	out := make(chan int)

	go func() {
		defer wg.Done()

		loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			}
		}

		close(out)
		fmt.Println("generator closed")
	}()

	return out
}

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	gen := generator(ctx, 1)

	wg.Add(1)

	loop:
	for i := 0; i < 5; i++ {
		select {
		case result, ok := <-gen:
			if ok {
				fmt.Println(result)
			} else {
				fmt.Println("timeout")
				break loop
			}
		}
	}

	cancel()
	wg.Wait()
}
