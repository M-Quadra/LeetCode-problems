package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	lenY := len(matrix)
	lenX := 0
	if lenY > 0 {
		lenX = len(matrix[0])
	}

	y, x := lenY-1, 0

	for y >= 0 && x < lenX {
		tmp := matrix[y][x]
		if target < tmp {
			y--
		} else if target > tmp {
			x++
		} else {
			return true
		}
	}

	return false
}

func main() {

	mtx := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	mtx2 := [][]int{
		{},
	}

	fmt.Println(searchMatrix(mtx2, 0))

	fmt.Println(searchMatrix(mtx, 9))
}
