package main

import "fmt"

func isHappy(n int) bool {
	slow, fast := n, step(n)
	for slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

func main() {
	nums := 19
	res := isHappy(nums)
	fmt.Println(res)
}
