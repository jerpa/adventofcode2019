package main

import (
	c "adventofcode2019/common"
	"strconv"
	"strings"
)

func getData() []int {
	s := c.ReadInputFile()
	d := []int{}
	for _, v := range strings.Split(s[0], ",") {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err.Error())
		}
		d = append(d, i)
	}
	return d
}
func main() {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			d := getData()
			d[1] = noun
			d[2] = verb
			for i := 0; i < len(d) && d[i] != 99; i += 4 {
				if d[i] == 1 {
					d[d[i+3]] = d[d[i+1]] + d[d[i+2]]
				} else if d[i] == 2 {
					d[d[i+3]] = d[d[i+1]] * d[d[i+2]]
				}
			}
			if d[0] == 19690720 {
				c.Print(noun, verb)
				return
			}
		}
	}
}
