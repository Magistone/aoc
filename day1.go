
package main

import (
    "fmt"
    "os"
    "bufio" 
    "strings"
    "strconv"
    "slices"
);


func day1() {
    content, error := os.Open("input")
    if error != nil {
        fmt.Println("can't read file")
        os.Exit(1)
    }
   
    a := make([]int, 0)
    b := make([]int, 0)

    scanner := bufio.NewScanner(bufio.NewReader(content))
    for scanner.Scan() {
        split := strings.Split(scanner.Text(), " ")
        val, err := strconv.Atoi(split[0])
        if err != nil { 
            panic("Dis no int")
        }
        a = append(a, val)
        
        val, err = strconv.Atoi(split[3])
        if err != nil { 
            panic("Dis no int")
        }
        b = append(b, val)
    }
     
    slices.Sort(a)
    slices.Sort(b)

    sum := 0

    for i := 0; i < len(a); i++ {
        diff := a[i] - b[i]
        if diff < 0 {
            diff *= -1
        }

        sum += diff
    }
    fmt.Println(sum)

    //part2
    sum = 0
    for i := 0; i < len(a); i++ {
        current := a[i]
        count := 0
        for j := 0; j < len(b); j++ {
            if b[j] > current {
                break 
            } 
            if current == b[j] {
                count++
            }    
        }
    sum += current*count
    }

    fmt.Println(sum)
}
