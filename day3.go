package main

import (
    "os"
    "fmt"
    "bufio"
    "regexp"
    "strconv"
)

func day3() {
    file, err := os.Open("input3")
    if err != nil {
        fmt.Println("Error opening file")
        panic("DED")
    }
    scanner := bufio.NewScanner(bufio.NewReader(file))
    
    regex := regexp.MustCompile("(?:(mul)\\((\\d{1,3}),(\\d{1,3})\\))|(?:(do)\\(\\))|(?:(don't)\\(\\))")
    
    sum := 0
    enabled := true
    for scanner.Scan() {
        str := scanner.Text()

        found := regex.FindAllStringSubmatch(str, -1)

        fmt.Println(found)
        for i := 0; i < len(found); i++ {
            if found[i][4] == "do" {
                enabled = true
            } else if found[i][5] == "don't" {
                enabled = false
            } else if enabled {
                num1, err := strconv.Atoi(found[i][2])
                num2, err2 := strconv.Atoi(found[i][3])
    
                if err != nil && err2 != nil {
                    panic("number conversion failed")
                }

                sum += num1*num2
            }
        }
    }

    fmt.Println(sum)
}
