package main

import "fmt"

func main() {
    m := make(map[string]int)
    
    m["Answer"] = 42
    fmt.Println("m: ", m)
    fmt.Println("m[Answer]: ", m["Answer"])
    
    fmt.Println("Deleting key with Answer")
    delete(m, "Answer")
    fmt.Println("m[Answer]: ", m["Answer"])
    
    v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}