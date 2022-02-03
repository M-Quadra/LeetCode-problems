func largest1BorderedSquare(grid [][]int) int {
	maxY, maxX := len(grid), len(grid[0])

	for y := range grid {
		l := 0
		for i, v := range grid[y] {
			if v <= 0 {
				l = 0
			} else {
				l++
			}
			grid[y][i] = l
		}
	}
	for x := range grid[0] {
		l := 0
		for i := range grid {
			if grid[i][x] <= 0 {
				l = 0
			} else {
				l++
			}
			grid[i][x] |= l << 8
		}
	}

	opt := 0
	for y, rows := range grid {
		for x := range rows {
			if grid[y][x] <= 0 {
				continue
			}

			for l := opt + 1; y+l <= maxY && x+l <= maxX; l++ {
				i, j := y, x+l-1
				if grid[i][j]&0xFF < l {
					goto next
				}
				i = i + l - 1
				if grid[i][j]&0xFF < l || grid[i][j]/0xFF < l {
					goto next
				}
				j = j - l + 1
				if grid[i][j]/0xFF < l {
					goto next
				}

				opt = l
			next:
			}

		}
	}
	return opt * opt
}