package main

import "fmt"

func addDigits(num int) int {
	res := 0
	for num > 0 {
		res += num % 10
		num = num / 10
	}
	for res >= 10 {
		num = res
		for num > 0 {
			res += num % 10
			num = num / 10
		}

	}
	return res
}

func main() {
	s := 38
	res := addDigits(s)
	fmt.Println(res)
}
