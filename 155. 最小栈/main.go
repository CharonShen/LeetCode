type MinStack struct {
    stack []int
    top int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        stack : []int{},
        top : 0,
    }
}


func (this *MinStack) Push(x int)  {
    this.top++
    this.stack = append(this.stack, x)
}


func (this *MinStack) Pop()  {
    if this.top > 0 {
        this.top--
        this.stack = this.stack[0:len(this.stack)-1]
    }
}


func (this *MinStack) Top() int {
    res := 0
    if this.top > 0 {
        res = this.stack[len(this.stack)-1]
    }
    return res
}


func (this *MinStack) GetMin() int {
    min := this.stack[0]
    for _, v := range this.stack {
        if v < min {
            min = v
        }
    }
    return min
}
