package main

import "fmt"

func search(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	left, right := 0, l-1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if (nums[left] < nums[mid] && target >= nums[left] && target < nums[mid]) || (nums[mid] < nums[right] && (target < nums[mid] || target > nums[right])) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
func main() {
	nums := []int{3, 5, 1}
	res := search(nums, 3)
	fmt.Println(res)
}
