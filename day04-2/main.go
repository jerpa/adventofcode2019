package main

import (
	c "adventofcode2019/common"
	"strconv"
)

func main() {
	sum := 0

	for i := 171309; i <= 643603; i++ {
		s := strconv.Itoa(i)
		same := false
		last := s[0]
		repeat := 1
		for d := 1; d < len(s); d++ {
			if s[d] == last {
				repeat++

			} else if s[d] < last {
				same = false
				repeat = 0
				break
			} else {
				if repeat == 2 {
					same = true
				}
				repeat = 1
			}
			last = s[d]
		}
		if repeat == 2 || same {
			sum++
		}
	}

	c.Print(sum)

}
