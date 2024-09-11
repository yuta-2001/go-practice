package main

import (
	"context"
	"sync"
	"fmt"
	"errors"
)

var wg sync.WaitGroup

func generator(ctx context.Context, num int) <-chan int {
	out := make(chan int)

	go func() {
		defer wg.Done()

		loop:
		for {
			// select {
			// case <-ctx.Done():
			// 	break loop
			// case out <-num:
			// }
			select {
			case <-ctx.Done():
				if err := ctx.Err(); errors.Is(err, context.Canceled) {
					fmt.Println("canceled")
				} else if errors.Is(err, context.DeadlineExceeded) {
					fmt.Println("deadline")
				}
				break loop
			}
		}
		close(out)
	}()

	return out
}

