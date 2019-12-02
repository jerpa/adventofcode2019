package main

import (
	c "adventofcode2019/common"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := c.ReadInputFile()
	d := []int{}
	for _, v := range strings.Split(s[0], ",") {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err.Error())
		}
		d = append(d, i)
	}
	d[1] = 12
	d[2] = 2
	for i := 0; i < len(d) && d[i] != 99; i += 4 {
		if d[i] == 1 {
			d[d[i+3]] = d[d[i+1]] + d[d[i+2]]
		} else if d[i] == 2 {
			d[d[i+3]] = d[d[i+1]] * d[d[i+2]]
		}
		
	}
	fmt.Println(d)

}
