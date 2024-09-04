package main

import (
    "sync"
)

func main() {
    c := sync.NewCond(&sync.Mutex{})
    c.L.Lock()
    for conditonTrue() == false {
        c.Wait()
    }
    c.L.Unlock()
}

func conditonTrue() bool {
    return false
}
