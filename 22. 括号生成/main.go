func generateParenthesis(n int) []string {
    res := []string{}
    dfs("", n, n, &res)
    return res
}

func dfs(cur string, left, right int, res *[]string) {
    if right == 0 {
        *res = append(*res, cur)
        return
    }
    if left > 0 {
        dfs(cur+"(", left-1, right, res)
    }
    if left < right {
        dfs(cur+")", left, right-1, res)
    }
}