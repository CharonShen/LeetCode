package main

import "fmt"

func convert(s string, numRows int) string {
	l := len(s)
	if l <= numRows || numRows == 1 {
		return s
	}
	res := ""
	key := numRows*2 - 2
	b := []byte(s)
	flag := 0
	for n := 0; n < l; n += key {
		res += string(b[n])
		flag++
	}
	for i := 1; i < numRows-1; i++ {
		for n := 0; n < flag; n++ {
			if n*key+i < l {
				res += string(b[n*key+i])
			}

			if (n+1)*key-i < l {
				res += string(b[(n+1)*key-i])
			}
		}
	}
	for n := numRows - 1; n < l; n += key {
		res += string(b[n])
	}

	return res
}

func main() {
	s := "A"
	res := convert(s, 1)
	fmt.Println(res)
}
