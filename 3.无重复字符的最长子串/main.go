func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	 m := map[byte]int{}
	 n := len(s)
	 l, r := 0, -1
	 res := 0
	 for l < n {
		 if r+1 < n && m[s[r+1]] == 0 {
			 m[s[r+1]] = 1
			 r++
		 } else {
			 delete(m,s[l])
			 l++
		 }
		 res = max(res, r+1-l)
	 }
	 return res
 }
 
 func max(x, y int) int {
	 if x < y {
		 return y
	 }
	 return x
 }