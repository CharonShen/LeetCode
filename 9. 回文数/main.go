func isPalindrome(x int) bool {
    if x < 0 || (x%10==0 && x!=0) {
        return false
    }
    num := 0
    for num < x {
        num = num*10 + x%10
        x /= 10
    }

    return num == x || num/10 == x
}