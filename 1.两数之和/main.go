package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		dif := target - v
		a, ok := m[dif]
		if ok {
			return []int{a, i}
		}
		m[v] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	arr := twoSum(nums, target)
	fmt.Println(arr)
}
