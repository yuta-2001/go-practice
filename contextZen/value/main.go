package main

import (
	"context"
	"sync"
	"fmt"
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
			case out <- num:
			}
		}

		close(out)
		userID, authToken, traceID := ctx.Value("userID").(int), ctx.Value("authToken").(string), ctx.Value("traceID").(string)
		fmt.Println("log: userID:", userID, "authToken:", authToken, "traceID:", traceID)
		fmt.Println("generator closed")
	}()
	return out
}


func main() {
	ctx, cancel := context.WithCancel(context.Background())

	ctx = context.WithValue(ctx, "userID", 2)
	ctx = context.WithValue(ctx, "authToken", "xxxxxx")
	ctx = context.WithValue(ctx, "traceID", 3)
	gen := generator(ctx, 1)

	wg.Add(1)

	for i := 0; i < 5; i++ {
		fmt.Println(<-gen)
	}

	cancel()

	wg.Wait()
}

