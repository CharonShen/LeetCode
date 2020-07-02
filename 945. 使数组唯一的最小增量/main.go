func minIncrementForUnique(A []int) int {
    l := len(A)
    if l <= 1 {
        return 0
    }
    qs(A, 0, l-1)
    res := 0
    for i:=1; i<l; i++ {
        if A[i] <= A[i-1] {
            res += A[i-1] + 1 - A[i]
            A[i] = A[i-1] + 1
        }
    }
    return res
}

func qs(a []int, left int, right int) {
    l, r := left, right
    key := a[(l+r)/2]
    for l < r {
        for a[l] < key {
            l++
        }
        for a[r] > key {
            r--
        }
        if l >= r {
            break
        }
        a[l], a[r] = a[r], a[l]
        if a[l] == key {
            r--
        }
        if a[r] == key {
            l++
        }
    }
    if l==r {
        l++
        r--
    }
    if left < r {
        qs(a, left, r)
    }
    if l < right {
        qs(a, l, right)
    }
}