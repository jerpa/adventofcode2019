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
		for d := 1; d < len(s); d++ {
			if s[d] == last {
				same = true
			} else if s[d] < last {
				same = false
				break
			}
			last = s[d]
		}
		if same {
			sum++
		}
	}

	c.Print(sum)

}
