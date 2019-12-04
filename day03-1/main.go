package main

import (
	c "adventofcode2019/common"
	"fmt"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func main() {
	grid := map[int]map[int]bool{}

	intersects := []point{}

	data := c.ReadInputFile()
	for p := range data {
		x := 0
		y := 0
		for _, v := range strings.Split(data[p], ",") {
			step, _ := strconv.Atoi(v[1:])
			if v[0] == 'R' {
				for i := 0; i < step; i++ {
					x++
					if grid[x] == nil {
						grid[x] = make(map[int]bool)
					}
					if grid[x][y] {
						intersects = append(intersects, point{x: x, y: y})
					}
					grid[x][y] = true
				}
			} else if v[0] == 'L' {
				for i := 0; i < step; i++ {
					x--
					if grid[x] == nil {
						grid[x] = make(map[int]bool)
					}
					if grid[x][y] {
						intersects = append(intersects, point{x: x, y: y})
					}
					grid[x][y] = true
				}
			} else if v[0] == 'U' {
				for i := 0; i < step; i++ {
					y++
					if grid[x] == nil {
						grid[x] = make(map[int]bool)
					}
					if grid[x][y] {
						intersects = append(intersects, point{x: x, y: y})
					}
					grid[x][y] = true
				}
			} else if v[0] == 'D' {
				for i := 0; i < step; i++ {
					y--
					if grid[x] == nil {
						grid[x] = make(map[int]bool)
					}
					if grid[x][y] {
						intersects = append(intersects, point{x: x, y: y})
					}
					grid[x][y] = true
				}
			}
			//fmt.Println(x, y)
		}
	}
	//fmt.Println(intersects)
	dist := -1
	for _, v := range intersects {
		if v.x < 0 {
			v.x *= -1
		}
		if v.y < 0 {
			v.y *= -1
		}
		if dist == -1 || dist > v.x+v.y {
			dist = v.x + v.y
		}
	}
	fmt.Println(dist)

}
