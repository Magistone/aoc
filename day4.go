package main

import(
    "os"
    "bufio"
    "strings"
    "fmt"
)

func day4() {
    f,err := os.Open("input4")
    if err != nil {
        panic("can't open file")
    }

    scanner := bufio.NewScanner(bufio.NewReader(f))

    charMap := make([][]string, 0)

    for scanner.Scan() {
        line := scanner.Text()
        chars := strings.Split(line, "")
        charMap = append(charMap, chars)
    }

    totalXmas := 0

    for y := 0; y < len(charMap); y++ {
        for x := 0; x < len(charMap[y]); x++ {
            totalXmas += findX_MAS(charMap, x, y)
        }
    }
    fmt.Println(totalXmas)
}

func findX_MAS(charMap [][]string, x int, y int) int {
    
    if charMap[x][y] != "A" {
        return 0
    }
    
    found := 0

    if x == 0 || x == len(charMap[y])-1 || y == 0 || y == len(charMap)-1 {
        return 0
    }

    //left Ms
    if charMap[x-1][y-1] == "M" && charMap[x-1][y+1] == "M" && charMap[x+1][y-1] == "S" && charMap[x+1][y+1] == "S" {
        found += 1
    }

    //right Ms
    if charMap[x+1][y-1] == "M" && charMap[x+1][y+1] == "M" && charMap[x-1][y-1] == "S" && charMap[x-1][y+1] == "S" {
        found += 1
    }

    //top Ms
    if charMap[x+1][y-1] == "M" && charMap[x+1][y+1] == "S" && charMap[x-1][y-1] == "M" && charMap[x-1][y+1] == "S" {
        found += 1
    }

    //bottom Ms
    if charMap[x+1][y+1] == "M" && charMap[x+1][y-1] == "S" && charMap[x-1][y+1] == "M" && charMap[x-1][y-1] == "S" {
        found += 1
    } 

    return found

}

func findXMAS(charMap [][]string, x int, y int) int {

    if charMap[x][y] != "X" {
        return 0;
    }

    found := 0

    //top
    if y >= 3 && charMap[x][y-1] == "M" && charMap[x][y-2] == "A" && charMap[x][y-3] == "S" {
        found += 1
    }

    //bottom
    if y <= len(charMap)-4 && charMap[x][y+1] == "M" && charMap[x][y+2] == "A" && charMap[x][y+3] == "S" {
        found += 1
    }

    //left
    if x >= 3 && charMap[x-1][y] == "M" && charMap[x-2][y] == "A" && charMap[x-3][y] == "S" {
        found += 1
    }

    //right
    if x <= len(charMap[y])-4 && charMap[x+1][y] == "M" && charMap[x+2][y] == "A" && charMap[x+3][y] == "S" {
        found += 1
    }

    //top_left
    if x >= 3 && y >= 3 && charMap[x-1][y-1] == "M" && charMap[x-2][y-2] == "A" && charMap[x-3][y-3] == "S" {
        found += 1
    }

    //top_right
    if x <= len(charMap[y])-4 && y >= 3 && charMap[x+1][y-1] == "M" && charMap[x+2][y-2] == "A" && charMap[x+3][y-3] == "S" {
        found += 1
    }

    //bottom_left
    if x >= 3 && y <= len(charMap)-4 && charMap[x-1][y+1] == "M" && charMap[x-2][y+2] == "A" && charMap[x-3][y+3] == "S" {
        found += 1
    }

    //bottom_right
    if x <= len(charMap[y])-4 && y <= len(charMap)-4 && charMap[x+1][y+1] == "M" && charMap[x+2][y+2] == "A" && charMap[x+3][y+3] == "S" {
        found += 1
    }

    return found;
}
