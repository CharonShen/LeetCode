package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	val   int
	times int
}
type intHeap []Node

func (h intHeap) Len() int {
	return len(h)
}
func (h intHeap) Less(i, j int) bool {
	return h[i].times < h[j].times
}
func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}
func (h *intHeap) Pop() interface{} {
	l := len(*h)
	x := (*h)[l-1]
	*h = (*h)[:l-1]
	return x
}
func topKFrequent(nums []int, k int) []int {
	h := make(intHeap, 0)
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	size := 0
	for i, v := range m {
		if size < k {
			heap.Push(&h, Node{
				val:   i,
				times: v,
			})
			size++
		} else if v > h[0].times {
			heap.Pop(&h)
			heap.Push(&h, Node{
				val:   i,
				times: v,
			})
		}
	}
	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(&h).(Node).times)
	}
	return res
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	res := topKFrequent(nums, 2)
	fmt.Println(res)
}
