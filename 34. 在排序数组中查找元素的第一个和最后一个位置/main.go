package main

import "fmt"

func searchRange(nums []int, target int) []int {
	l := len(nums)
	if l == 0 {
		return []int{-1, -1}
	}
	left, right := 0, l-1
	a, b := -1, -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			a, b = mid, mid
			for a > left && nums[a-1] == target {
				a--
			}
			for b < right && nums[b+1] == target {
				b++
			}
			break
		}
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
	}
	return []int{a, b}
}
func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	res := searchRange(nums, 8)
	fmt.Println(res)
}
