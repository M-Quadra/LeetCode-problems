var dir = [4][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func dfs(i, j int, upsampling [][]int) {
	n := len(upsampling)
	for _, v := range dir {
		ni, nj := i+v[0], j+v[1]
		if ni < 0 || n <= ni || nj < 0 || n <= nj || upsampling[ni][nj] != 0 {
			continue
		}

		upsampling[ni][nj] = upsampling[i][j]
		dfs(ni, nj, upsampling)
	}
}

func regionsBySlashes(grid []string) int {
	n := len(grid)
	upsampling := make([][]int, n*3)
	for i := 0; i < n*3; i++ {
		upsampling[i] = make([]int, n*3)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			switch grid[i][j] {
			case []byte("\\")[0]:
				upsampling[i*3][j*3] = -1
				upsampling[i*3+1][j*3+1] = -1
				upsampling[i*3+2][j*3+2] = -1
			case []byte("/")[0]:
				upsampling[i*3][j*3+2] = -1
				upsampling[i*3+1][j*3+1] = -1
				upsampling[i*3+2][j*3] = -1
			}
		}
	}

	n *= 3
	opt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if upsampling[i][j] != 0 {
				continue
			}

			opt++
			upsampling[i][j] = opt
			dfs(i, j, upsampling)
		}
	}

	return opt
}