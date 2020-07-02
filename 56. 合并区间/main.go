package main

import "fmt"

func merge(intervals [][]int) [][]int {
	l := len(intervals)
	if l == 0 {
		return nil
	}
	var res [][]int
	qsort(intervals, 0, l-1)
	res = append(res, intervals[0])
	k := intervals[0][1]
	for i := 1; i < l; i++ {
		if intervals[i][0] > k {
			res = append(res, intervals[i])
			k = intervals[i][1]
			continue
		}
		if intervals[i][1] > k {
			res[len(res)-1][1] = intervals[i][1]
			k = intervals[i][1]
		}
	}
	return res
}
func qsort(a [][]int, left int, right int) {
	l, r := left, right
	k := a[(l+r)/2][0]
	for l < r {
		for a[l][0] < k {
			l++
		}
		for a[r][0] > k {
			r--
		}
		if l >= r {
			break
		}
		a[l], a[r] = a[r], a[l]

		if a[l][0] == k {
			r--
		}
		if a[r][0] == k {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	if left < r {
		qsort(a, left, r)
	}
	if right > l {
		qsort(a, l, right)
	}
}

func main() {
	nums := [][]int{{1, 4}, {0, 2}, {3, 5}}
	res := merge(nums)
	fmt.Println(res)
}
