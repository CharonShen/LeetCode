package main

import (
	"fmt"
	"math"

	"golang.org/x/crypto/ripemd160"
)

func threeSumClosest(nums []int, target int) int {
	left, right := 0, len(nums)-1
	QuickSort(nums, left, right)
	//-4 -1 1 2
	k, i, j := 0, left+1, right
	res, d := 0, 0
	flag := target - (nums[k] + nums[i] + nums[j])
	for k < len(nums)-2 {
		sum := (nums[k] + nums[i] + nums[j])
		d = target - sum
		if math.Abs(float64(flag)) >= math.Abs(float64(d)) {
			flag = int(math.Abs(float64(d)))
			res = sum
		}

		if d == 0 {
			return target
		} else if d > 0 {
			i++
		} else {
			j--
		}
		if i >= j {
			k++
			i = k + 1
			j = right
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

		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
		if a[i] == key {
			j--
		}
		if a[j] == key {
			i++
		}
	}
	if i == j {
		i++
		j--
	}
	if left < j {
		QuickSort(a, left, j)
	}
	if right > j {
		QuickSort(a, i, right)
	}
}
func main() {
	nums := []int{1, 2, 4, 8, 16, 32, 64, 128}
	res := threeSumClosest(nums, 82)
	fmt.Println(res)
	ripemd160 := ripemd160.New()
	fmt.Println(ripemd160)
}
