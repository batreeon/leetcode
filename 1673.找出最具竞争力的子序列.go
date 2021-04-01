/*
 * @lc app=leetcode.cn id=1673 lang=golang
 *
 * [1673] 找出最具竞争力的子序列
 */

// @lc code=start
package main

// import "sort"

type stack struct {
	s []int
	l int
}

func (st *stack) push(x int) {
	st.s = append(st.s,x)
	st.l++
}

func (st *stack) top() int {
	return st.s[st.l-1]
}

func (st *stack) pop() {
	st.s = st.s[:st.l-1]
	st.l--
}

func(st stack) isEmpty() bool {
	return st.l == 0
}

func mostCompetitive(nums []int, k int) []int {
	/*
	n := len(nums)
	result := []int{}
	minIndex := 0
	for i := 0 ; i < k ; i++ {
		if sort.IsSorted(sort.IntSlice(nums)) {
			return nums[:k]
		}
		if sort.IsSorted(sort.IntSlice(nums[:n-1])) && nums[n-1] == 0 {
			result = append(result,nums[:k-1]...)
			result = append(result,0)
			return result
		}
		for j := minIndex+1 ; j < n-(k-i)+1 ; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		result = append(result,nums[minIndex])
		minIndex++
	}
	return result
	*/
		/*
		// 不可行
		if sort.IsSorted(sort.IntSlice(nums[minIndex:n-(k-i)+1])) {
			result = append(result,nums[minIndex])
			minIndex++
			if minIndex == n {
				break
			}
			for nums[minIndex] >= nums[minIndex-1] {
				result = append(result,nums[minIndex])
				minIndex++
				if minIndex == n {
					break
				}
				i++
			}
		}else{
			for j := minIndex+1 ; j < n-(k-i)+1 ; j++ {
					if nums[j] < nums[minIndex] {
						minIndex = j
					}
				}
			result = append(result,nums[minIndex])
			minIndex++
		}
	}
	return result
	*/

	// 维护一个单调栈
	s := stack{}
	n := len(nums)
	for i := 0 ; i < n ; i++ {
		// k-s.l是还差的元素数，n-i是nums中剩余的元素数
		// 最多这两个相等就不能再pop了，不然元素就不够用了
		for !s.isEmpty() && s.top() > nums[i] && k-s.l < n-i {
			s.pop()
		}
		if s.l < k {
			s.push(nums[i])
		}
	}
	return s.s
}

// @lc code=end
