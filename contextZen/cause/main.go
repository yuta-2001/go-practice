package main


import (
	"fmt"
	"context"
	"sync"
	"errors"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancelCause(context.Background())
	wg.Add(1)
	go task(ctx)
	time.Sleep(3 * time.Second)
	cancel(errors.New("canceled by main func"))
	wg.Wait()
}

func task(ctx context.Context) {
	defer wg.Done()

	ctx, cancel := context.WithCancelCause(ctx)
	wg.Add(1)
	go subTask(ctx)

	time.Sleep(5 * time.Second)
	cancel(errors.New("canceled by task func"))
}

func subTask(ctx context.Context) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println(context.Cause(ctx))
			return
		case <-time.After(1 * time.Second):
			fmt.Println("subTask done")
		}
	}
}

