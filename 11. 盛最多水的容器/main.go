package main

import (
	"fmt"
	"math"
)

func main() {
	height := []int{1, 2, 4, 3}
	res := maxArea(height)
	fmt.Println(res)
}
func maxArea(height []int) int {
	l, r, h := 0, 0, 0
	res := 0
	n := len(height)

	for l < n-1 {
		if r+1 < n {
			r++
			h = minVal(height[l], height[r])
		} else {
			l++
			r = l
		}
		res = maxVal(res, int(math.Abs(float64(r-l)))*h)
	}
	return res
}
func minVal(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func maxVal(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
