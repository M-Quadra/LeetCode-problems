var mergeFind = make([]int, 30*30<<2)

func find(x int) int {
	if mergeFind[x] == mergeFind[mergeFind[x]] {
		return mergeFind[x]
	}

	mergeFind[x] = find(mergeFind[x])
	return mergeFind[x]
}

func union(x, y int) {
	x, y = find(x), find(y)
	if x == y {
		return
	} else if x > y {
		x, y = y, x
	}

	mergeFind[y] = x
}

func regionsBySlashes(grid []string) int {
	n := len(grid)
	for i := 0; i < n*n*4; i++ {
		mergeFind[i] = i
	}

	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			k := (y*n + x) << 2
			switch rune(grid[y][x]) {
			case '\\':
				union(k+0, k+1)
				union(k+3, k+2)
			case '/':
				union(k+0, k+3)
				union(k+1, k+2)
			default:
				union(k+0, k+1)
				union(k+1, k+2)
				union(k+2, k+3)
			}

			if y > 0 {
				kU := k - n*4
				union(k+0, kU+2)
			}
			if x+1 < n {
				kR := k + 4
				union(k+1, kR+3)
			}
			if y+1 < n {
				kD := k + 4*n
				union(k+2, kD)
			}
			if x > 0 {
				kL := k - 4
				union(k+3, kL+1)
			}
		}
	}

	opt := 0
	for i := 0; i < n*n*4; i++ {
		if mergeFind[i] == i {
			opt++
		}
	}
	return opt
}