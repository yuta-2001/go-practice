package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("使い方: go run main.go [正の整数]")
        return
    }

    max, err := strconv.Atoi(os.Args[1])
    if err != nil || max < 0 {
        fmt.Println("正の整数を指定してください。")
        return
    }

    fmt.Printf("0から%dまでの素数:\n", max)
    for num := 2; num <= max; num++ {
        if isPrime(num) {
            fmt.Printf("%d ", num)
        }
    }
    fmt.Println()
}

func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}
