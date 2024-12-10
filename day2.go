package main 

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    "slices"
)

func day2() {
    f, error := os.Open("input2")
    if error != nil {
        panic("File is ded")
    }
    defer f.Close()
    
    
    scanner := bufio.NewScanner(bufio.NewReader(f))
    count_valid := 0
    for scanner.Scan() {
        parts := strings.Split(scanner.Text(), " ")
        int_parts := make([]int, len(parts))

        for i := 0; i < len(parts); i++ {
            val, err := strconv.Atoi(parts[i])
            if err != nil {
                panic("Can't parse int: " + parts[i])
            }
            int_parts[i] = val
        }
        if reportSafe(int_parts, false) {
        count_valid++
        }
    }
    fmt.Println(count_valid)
}


func reportSafe(report []int, recursed bool) bool {
    diff_arr := make([]int, len(report)-1)

    for i := 0; i < len(report)-1; i++ {
        diff_arr[i] = report[i] - report[i+1]
    }

    all, any := numbersPositive(diff_arr)

   
    max := slices.Max(diff_arr)
    min := slices.Min(diff_arr)

    if max < 0 {
        max *= -1
    }
    if min < 0 {
        min *= -1
    }

    if max > 0 && max < 4 && min > 0 && min < 4 {
        if !((all || any) && !(all && any)) {
            //fmt.Println("OK - BASE CASE ", diff_arr)
            return true

        }
     }
 
    if len(report) <= 3 {
        //fmt.Println("LEN 4 ", report)
        return true
    }
    
    if recursed {
        return false
    }
    

    fixable := false

    for i := 0; i < len(report); i++ {
        test := removeIndexFromSlice(report, i)
        fixable = fixable || reportSafe(test, true)
    }

    return fixable
}


func numbersPositive(diffs []int) (all bool, any bool) {
    all_positive := true
    any_positive := false
    for i := 0; i < len(diffs); i++ {
        if diffs[i] > 0 {
            any_positive = true
        } else if diffs[i] < 0 {
            all_positive = false
        }
    }
    return all_positive, any_positive
}

func removeIndexFromSlice(slice []int, idx int) (new_slice []int) {
    for i := 0; i < len(slice); i++ {
        if i == idx {
            continue
        }
        new_slice = append(new_slice, slice[i])
    }

    return new_slice
}
