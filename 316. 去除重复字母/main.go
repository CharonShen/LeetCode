func removeDuplicateLetters(s string) string {
    stack := []byte{}
    m := make(map[byte]int)
    for i:=0; i<len(s); i++ {
        m[s[i]]++
    }
    for i:=0; i<len(s); i++ {
        flag := false
        for _, v := range stack {
            if v == s[i] {
                flag = true
                break
            }
        }
        if flag {
            m[s[i]]--
            continue
        } else {
            for {
                if len(stack)!=0 && m[stack[len(stack)-1]]!=0 && stack[len(stack)-1]>s[i] {
                    stack = stack[0:len(stack)-1]
                } else {
                    stack = append(stack, s[i])
                    m[s[i]]--
                    break
                }
            }
        }
    }
    var str string
    for _, v := range stack {
        str += string(v)
    }
    return str
}