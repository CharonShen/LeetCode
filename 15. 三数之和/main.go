package main

import "fmt"

func threeSum(nums []int) [][]int {
	var res [][]int
	i, j := 0, len(nums)-1
	k, sum := 0, 0
	QuickSort(nums, i, j)
	i = 1
	for nums[k] <= 0 && k < len(nums)-2 {
		if k > 0 && nums[k] == nums[k-1] {
			k++
			i = k + 1
			continue
		}
		if j < len(nums)-1 && nums[j] == nums[j+1] && j > i {
			j--
			continue
		}
		if i > j-1 {
			k++
			i = k + 1
			j = len(nums) - 1
			continue
		}
		sum = nums[k] + nums[i] + nums[j]
		if sum == 0 {
			res = append(res, []int{nums[k], nums[i], nums[j]})
			i++
			j--
			if i >= j {
				k++
				i = k + 1
				j = len(nums) - 1
			}
			continue
		} else if sum > 0 {
			j--
		} else {
			i++
		}
	}
	return res
}

func QuickSort(a []int, left int, right int) {
	i, j := left, right
	key := a[(i+j)/2]
	for i < j {
		for a[i] < key {
			i++
		}
		for a[j] > key {
			j--
		}
		if i <= j {
			a[i], a[j] = a[j], a[i]
			i++
			j--
		}
	}
	if left < j {
		QuickSort(a, left, j)
	}
	if right > i {
		QuickSort(a, i, right)
	}
}

func main() {
	nums := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	res := threeSum(nums)
	fmt.Println(res)
}
