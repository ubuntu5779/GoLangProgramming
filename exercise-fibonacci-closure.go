package main

import "fmt"

func fibonacci() func() int {
    a := 1
    b := 2  
    return func() int {
        temp := a + b
        a = b
        b = temp
        return temp
    }
}

func main() {
    fib := fibonacci()
    fmt.Println(1)
    fmt.Println(2)
    for i := 0; i<10; i++ {
        fmt.Println(fib())
    }
}