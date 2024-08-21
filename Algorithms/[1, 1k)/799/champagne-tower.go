func champagneTower(poured int, query_row int, query_glass int) float64 {
	tower := make([][]float64, query_row+1)
	for i := range tower {
		tower[i] = make([]float64, i+1)
	}

	tower[0][0] = float64(poured)
	for i := 0; i < query_row; i++ {
		for j := 0; j < len(tower[i]); j++ {
			if tower[i][j] <= 1 {
				continue
			}

			v := (tower[i][j] - 1) / 2
			tower[i+1][j] += v
			tower[i+1][j+1] += v
			tower[i][j] = 1
		}
	}

	return math.Min(tower[query_row][query_glass], 1)
}