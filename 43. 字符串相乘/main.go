func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
		return "0"
	}
    m, n := len(num1), len(num2)
    sum := make([]int, m+n)
    for i:=m-1; i>=0; i-- {
        for j:=n-1; j>=0; j-- {
            s := int(num1[i]-'0') * int(num2[j]-'0')
            p1 := i + j
            p2 := i + j + 1
            s += sum[p2]
            sum[p1] += s / 10
            sum[p2] = s % 10
        }
    }
    for i, v := range sum {
		if v != 0 {
			sum = sum[i:]
			break
		}
	}
    res := ""
    for _, v := range sum {
        res += string(v + '0')
    }
    return res
}