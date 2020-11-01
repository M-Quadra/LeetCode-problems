package main

func canMeasureWater(x int, y int, z int) bool {
	if x+y < z {
		return false
	}
	if x*y == 0 {
		return z == x || z == y
	}
	if x < y {
		x, y = y, x
	}

	c := 0
	for {
		v := z - (x - c)
		if v/y*y == v {
			return true
		}

		c = y - (x-c)%y
		if c == 0 || c == y {
			return false
		}
	}
}
