var direct = [][]int{
	{+1, +1},
	{+1, -1},
	{-1, -1},
	{-1, +1},
}

func sum(h, y, x int, grid [][]int) int {
	if h <= 1 {
		return grid[y][x]
	}

	opt := 0
	for _, v := range direct {
		for i := 0; i < h/2; i++ {
			y += v[0]
			x += v[1]
			opt += grid[y][x]
		}
	}
	return opt
}

func getBiggestThree(grid [][]int) []int {
	maxY, maxX := len(grid), len(grid[0])
	maxH := maxY
	if maxX < maxH {
		maxH = maxX
	}
	maxH -= maxH&1 ^ 1

	opt := []int{}
	for h := 1; h <= maxH; h += 2 {
		set := map[int]struct{}{}
		for _, v := range opt {
			set[v] = struct{}{}
		}

		for y, v := range grid {
			if y+h > maxY {
				break
			}

			for x := range v {
				if x-h/2 < 0 || x+h/2 >= maxX {
					continue
				}

				n := sum(h, y, x, grid)
				if len(opt) > 0 && n <= opt[len(opt)-1] {
					continue
				}

				set[n] = struct{}{}
			}
		}

		keys := make([]int, 0, len(set))
		for k := range set {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool {
			return keys[i] > keys[j]
		})

		if len(keys) < 3 {
			opt = keys
		} else {
			opt = keys[:3]
		}
	}

	return opt
}
