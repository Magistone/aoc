package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type vector2D struct {
	horizontal int
	vertical   int
}

func day6() {
	f, err := os.Open("input6")
	if err != nil {
		panic("Oh shit something went wrong")
	}
	reader := bufio.NewScanner(bufio.NewReader(f))

	map_area := make([][]string, 0)
	map_bool := make([][]bool, 0)

	for reader.Scan() {
		text := reader.Text()
		text_slice := strings.Split(text, "")
		map_area = append(map_area, text_slice)
		map_bool = append(map_bool, make([]bool, len(text_slice)))
	}

	x, y := find_start(map_area)
	direction := vector2D{horizontal: 0, vertical: -1}

	if map_area[y][x] != "^" {
		panic("You fked up XY assignment")
	}
    
    map_bool[y][x] = true
    for ; y >= 0 && y < len(map_area) && x >= 0 && x < len(map_area[0]) ; {
        x,y = move(map_area, &direction, x, y) 
        if y < 0 || y >= len(map_area) || x < 0 || x >= len(map_area[0]) {
            break;
        }
        map_bool[y][x] = true
    }
        
	//fmt.Println(map_area)
	//fmt.Println(map_bool)
    fmt.Println(count_trues(map_bool))
}

func find_start(map_area [][]string) (int, int) {
	for i := 0; i < len(map_area); i++ {
		for j := 0; j < len(map_area[i]); j++ {
			if map_area[i][j] == "^" {
				return j, i
			}
		}
	}
	return -1, -1
}

func move(map_area [][]string, direction *vector2D, x int, y int) (int, int) {
    new_x, new_y := 0, 0
	for i := 0; i < 6; i++ {
        if i == 4 {
            panic("This shit ain't rotatin")
        }
		new_x = x + direction.horizontal
		new_y = y + direction.vertical

		if new_x >= 0 && new_x < len(map_area[y]) && new_y >= 0 && new_y < len(map_area) && map_area[new_y][new_x] == "#" {
			rotate(direction)
		} else {
			break
		}
	}
    return new_x, new_y
}

func rotate(direction *vector2D) {
    if direction.vertical == -1 {
        direction.vertical = 0
        direction.horizontal = 1
    } else if direction.vertical == 1 {
        direction.vertical = 0
        direction.horizontal = -1
    } else if direction.horizontal == -1 {
        direction.vertical = -1
        direction.horizontal = 0
    } else if direction.horizontal == 1 {
        direction.vertical = 1
        direction.horizontal = 0
    } else {
        panic ("SHEIT")
    }   
}

func count_trues(map_bool [][]bool) int {
    count := 0
    for i := 0; i < len(map_bool); i++ {
        for j := 0; j < len(map_bool[i]); j++ {
            if map_bool[i][j] {
                count++
            }
        }
    }
    return count
}   
