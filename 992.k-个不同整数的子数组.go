/*
 * @lc app=leetcode.cn id=992 lang=golang
 *
 * [992] K 个不同整数的子数组
 */

// @lc code=start
func subarraysWithKDistinct(A []int, K int) int {
	/*
	ans := 0
	for l := K ; l <= len(A) ; l++ {
		m := make(map[int]bool)
		for left := 0 ; left < len(A)-l+1 ; left++ {
			for right := left ; right < left + l ; right++ {
				if m[A[right]] == false {
					m[A[right]] = true
				}
			}
			if len(m) == K {
				ans++
			}
			delete(m,A[left])
		}
	}
	return ans
	*/
	//超时了复杂度太高了

	n := len(A)
	num1 := make([]int,n+1)
	num2 := make([]int,n+1)
	total1,total2,left1,left2 := 0,0,0,0
	ans := 0
	for _,v := range A {//每一个right对应两个left，right固定，[left1,right]是极长好子数组，[left2,right]是极长非好(比K小)数组
		if num1[v] == 0 {
			total1++
		}
		num1[v]++
		if num2[v] == 0 {
			total2++
		}
		num2[v]++
		for total1 > K {
			num1[A[left1]]--
			if num1[A[left1]] == 0 {
				total1--
			}
			left1++
		}
		for total2 > K - 1 {
			num2[A[left2]]--
			if num2[A[left2]] == 0 {
				total2--
			}
			left2++
		}
		ans += left2-left1//每一个right对应的好数组个数
	}
	return ans
}
// @lc code=end

