package main
func lengthOfLongestSubstring(s string) int {
	left,right := 0,0
	seen := map[byte]bool{}
	result := 0
	for ; right < len(s) ; right++ {
		if seen[s[right]] {
			for s[left] != s[right] {
				delete(seen,s[left])
				left++
			}
			left++
		}else{
			seen[s[right]] = true
			if right - left + 1 > result {
				result = right-left+1
			}
		}
	}
	return result
}