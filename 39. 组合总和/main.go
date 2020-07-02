package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	dfs(candidates, []int{}, target, 0, &res)
	return res
}

func dfs(candidates []int, nums []int, target int, left int, res *[][]int) {
	if target == 0 {
		*res = append(*res, nums)
		return
	}

	n := len(candidates)

	for i := left; i < n; i++ {
		if candidates[i] > target {
			return
		}
		dfs(candidates, append(nums, candidates[i]), target-candidates[i], i, res)
	}
}
func main() {

	res := combinationSum([]int{7, 3, 2}, 18)
	fmt.Println(res)
}
