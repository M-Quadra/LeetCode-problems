func largest1BorderedSquare(grid [][]int) int {
	maxY, maxX := len(grid), len(grid[0])

	dir := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	opt := 0
	for y, rows := range grid {
		for x := range rows {
			if grid[y][x] <= 0 {
				continue
			}

			for l := opt + 1; y+l <= maxY && x+l <= maxX; l++ {
				i, j := y, x
				for _, d := range dir {
					for k := 1; k < l; k++ {
						i, j = i+d[0], j+d[1]
						if grid[i][j] <= 0 {
							goto next
						}
					}
				}

				opt = l
			next:
			}

		}
	}
	return opt * opt
}