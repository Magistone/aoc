package main

import "fmt"

func main() {
    day5()   
}


func never(args ...interface{}) {
    fmt.Println("NEVER: ", args)
    panic("NEVER REACHED")
}

func assert(condition bool, msg_fail string) {
    if condition {
        return
    }
   panic(msg_fail) 
}

