type Heap []int

func (h Heap)Len() int{
    return len(h)
}
func (h Heap)Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}
func (h Heap)Less(i, j int) bool{
    return h[i] < h[j]
}
func (h *Heap)Push(x interface{}) {
    *h = append(*h, x.(int))
}
func (h *Heap)Pop() interface{}{
    l := len(*h)
    val := (*h)[l-1]
    *h =  (*h)[0:l-1]
    return val
}
func findKthLargest(nums []int, k int) int {
    h := &Heap{}
    heap.Init(h)
    for _, v := range nums {
        if h.Len() < k {
            heap.Push(h, v)
        } else if v > (*h)[0] {
            heap.Pop(h)
            heap.Push(h, v)
        }
    }
    return (*h)[0]
}