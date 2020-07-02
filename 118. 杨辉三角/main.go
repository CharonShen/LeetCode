func generate(numRows int) [][]int {
    if numRows == 0 {
        return nil
    }
    res := make([][]int, numRows)
    res[0] = make([]int, 1)
    res[0][0] = 1
    for i:=1; i<numRows; i++ {
        res[i] = make([]int, i+1)
        //tmp := make([]int, i+1)
        res[i][0], res[i][i] = 1, 1
        for j := 1; j < i; j++ {
            res[i][j] = res[i-1][j-1] + res[i-1][j]
        }
        //res = append(res, tmp)
    }
    return res
}