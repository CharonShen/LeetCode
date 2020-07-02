package main

import "fmt"

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row, col := []int{}, []int{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row = append(row, i)
				col = append(col, j)
			}
		}
	}

	for _, i := range row {
		for j := 0; j < n; j++ {
			matrix[i][j] = 0
		}
	}
	for _, j := range col {
		for i := 0; j < m; j++ {
			matrix[i][j] = 0
		}
	}
}

func main() {
	s := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(s)
	fmt.Println(s)
}
