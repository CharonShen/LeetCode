package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	res := []int{}
	flag := 0
	if m%2 == 0 {
		flag = m / 2
	} else {
		flag = m/2 + 1
	}
	for k := 0; k < flag; k++ {
		for j := k; j <= n-1-k; j++ {
			res = append(res, matrix[k][j])
		}
		for i := k + 1; i < m-k; i++ {
			res = append(res, matrix[i][n-1-k])
		}
		for j := n - 2 - k; j > k; j-- {
			res = append(res, matrix[m-1-k][j])
		}
		if len(res) == m*n {
			return res
		}
		for i := m - 1 - k; i > k; i-- {
			res = append(res, matrix[i][k])
		}
	}
	return res
}
func main() {
	s := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16}}
	res := spiralOrder(s)
	fmt.Println(res)
}
