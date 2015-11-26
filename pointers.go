package main

import "fmt"

func main() {
    i, j := 42, 2701
    
    p := &i     // point to i
    fmt.Println(*p)     // read i through pointer
    
    *p = 21     // set i through pointer
    fmt.Println(i)      // see the new value of i
    
    p = &j      // point to j
    *p = *p / 27        // divide j through pointer
    fmt.Println(j)      // see the value of j
}