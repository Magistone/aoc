package main

import "fmt"

func main() {
    day4()   
}


func never(args ...interface{}) {
    fmt.Println("NEVER: ", args)
    panic("NEVER REACHED")
}

