func myAtoi(str string) int {
    res := 0
    i := 0
    l := len(str)
    INT_MAX, INT_MIN := 1<<31-1, -(1<<31)
    flag := true
    if l < 1 {
        return 0
    }

    for i<l && str[i] == ' ' {
        i++
    }
    if i<l && str[i]=='-' {
        flag = false
    }
    if i<l && (str[i]=='+' || str[i]=='-') {
        i++
    }
    for i<l && str[i]>='0' && str[i]<='9' {
        var r int = int(str[i] - '0')
        if res > INT_MAX/10 || (res == INT_MAX/10 && r>7) {
            if flag {
                return INT_MAX
            } else {
                return INT_MIN
            }
        }
        res = res*10 + r
        i++
    }
    if flag {
        return res
    } else {
        return -res
    } 
}