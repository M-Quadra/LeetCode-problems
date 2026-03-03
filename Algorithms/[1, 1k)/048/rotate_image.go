package main

func rotate(matrix [][]int) {
	n := len(matrix)
	for y := 0; y < n-y; y++ {
		for x := y; x < n-y-1; x++ {
			matrix[x][n-1-y], matrix[n-1-y][n-1-x], matrix[n-1-x][y], matrix[y][x] =
				matrix[y][x], matrix[x][n-1-y], matrix[n-1-y][n-1-x], matrix[n-1-x][y]
		}
	}
}
