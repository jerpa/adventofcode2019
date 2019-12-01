package main

import (
	c "adventofcode2018/common"
)

func main() {
	sum := 0

	for _, v := range c.GetInts(c.ReadInputFile()) {
		for v > 0 {
			v /= 3
			v -= 2
			if v < 0 {
				v = 0
			}
			sum += v
		}
	}
	c.Print(sum)

}
