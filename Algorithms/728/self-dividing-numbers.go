package main

import (
	"fmt"
	"strconv"
)

func selfDividingNumbers(left int, right int) []int {
	opt := []int{}
	for i := left; i <= right; i++ {
		didUse := make([]bool, 10)
		str := strconv.Itoa(i)

		ok := true
		for idx := 0; idx < len(str); idx++ {
			digital := int(str[idx] - 48)
			if digital == 0 {
				ok = false
				break
			}

			if didUse[digital] {
				continue
			}
			didUse[digital] = true

			if i%digital != 0 {
				ok = false
				break
			}
		}

		if ok {
			opt = append(opt, []int{i}...)
		}
	}

	return opt
}

func main() {
	// ts := make([]bool, 12)
	fmt.Println(selfDividingNumbers(1, 22))
	// fmt.Println(12.string())
}
