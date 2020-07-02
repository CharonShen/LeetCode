package main

import "fmt"

// func removeDuplicates(nums []int) int {
// 	for i := 0; i < len(nums); i++ {
// 		if i > 0 && nums[i] == nums[i-1] {
// 			nums = append(nums[:i], nums[i+1:]...)
// 			i--
// 		}
// 	}
// 	return len(nums)
// }
func removeDuplicates(nums []int) int {
	i, j := 0, 0
	length := len(nums)
	for j = 1; j < length; j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	res := removeDuplicates(nums)
	fmt.Println(res)
}
