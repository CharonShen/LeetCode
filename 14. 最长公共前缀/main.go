func longestCommonPrefix(strs []string) string {
    l := len(strs)
    res := ""
    if l == 0 {
        return res
    }
    // if l == 1 {
    //     return strs[0]
    // }
    
    minstr := len(strs[0])
    for i:=1; i<l; i++ {
        if len(strs[i]) < minstr {
            minstr = len(strs[i])
        }
    }

    for i:=0; i<minstr; i++ {
        flag := true
        for j:=0; j<l-1; j++{
            if strs[j][i] != strs[j+1][i] {
                flag = false
            }
        }
        if flag {
            res += string(strs[0][i])
        } else {
            return res
        }
    }
    return res
}