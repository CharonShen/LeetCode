func permuteUnique(nums []int) [][]int {
    sort.Ints(nums)
    res := make([][]int, 0)
    path := make([]int, 0)
    used := make([]bool, len(nums))
    dfs(nums, path, used, &res)
    return res
}

func dfs(nums, path []int, used []bool, res *[][]int) {
    if len(nums) == len(path) {
        temp :=  make([]int, len(path))
        copy(temp, path)
        *res = append(*res, temp)
        return
    }
    for i:=0; i<len(nums); i++ {
        if used[i] {
            continue
        }
        path = append(path, nums[i])
        used[i] = true
        dfs(nums, path, used, res)
        path = path[:len(path)-1]
        used[i] = false
    }
}