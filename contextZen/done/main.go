package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// channelで制御するver
// func generator(done chan struct{}, num int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		defer wg.Done()

// 		loop:
// 		for {
// 			select {
// 			case <-done:
// 				break loop
// 			case out <- num:
// 			}
// 		}
// 	}()
// 	return out
// }

// func main() {
// 	done := make(chan struct{})
// 	gen := generator(done, 1)

// 	wg.Add(1)

// 	for i := 0; i < 5; i++ {
// 		fmt.Println(<-gen)
// 	}
// 	close(done)

// 	wg.Wait()
// }


