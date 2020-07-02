package main

import "fmt"

func minSubArrayLen(s int, nums []int) int {
	nlen := len(nums)
	if nlen < 1 {
		return 0
	}

	i, j := 0, -1
	ans, sum := nlen+1, nums[i]

	for i < nlen {
		if j+1 < nlen && sum < s {
			j++
			sum += nums[j]
		} else {
			sum -= nums[i]
			i++
		}
		if sum >= s {
			ans = min(ans, j-i+1)
		}
	}
	if ans == nlen+1 {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	res := minSubArrayLen(15, nums)
	fmt.Println(res)
}
