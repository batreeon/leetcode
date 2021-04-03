/*
 * @lc app=leetcode.cn id=1718 lang=golang
 *
 * [1718] 构建字典序最大的可行序列
 */

// @lc code=start
package main

func constructDistancedSequence(n int) []int {
	/*
		if n == 1 {
			return []int{1}
		}
		nums := make([]bool,n+1)
		result := make([]int,2*n-1)
		l := 0
		num := n
		var constructdistancedsequence func() bool
		constructdistancedsequence = func() bool {
			if l == 2*n - 1 {
				return true
			}
			for i := num ; i >= 1 ; i-- {
				if !nums[i] {
					if i == 1 {
						// 1只有一个元素，就不用为第二个元素的下标合法性考虑了
						for j := 0 ; j < 2*n-1 ; j++ {
							if result[j] == 0 {
								nums[i] = true
								result[j] = i
								l++
								num--
								if constructdistancedsequence() {
									return true
								}
								num++
								l--
								result[j] = 0
								nums[i] = false
								// break
							}
						}
					}else{
						j := 0
						for  ; j+i < 2*n-1 ; j++ {
							if result[j] == 0 && result[j+i] == 0 {
								nums[i] = true
								result[j],result[j+i] = i,i
								l = l+2
								num--
								if constructdistancedsequence() {
									return true
								}
								num++
								l = l-2
								result[j],result[j+i] = 0,0
								nums[i] = false
								// break
							}
						}
						// if j+i == 2*n-1 {
						// 	break
						// }
					}
				}
			}
			// lastBig := -1
			// for i := 2*n-1-1 ; i >= 0 ; i-- {
			// 	if result[i] != 0 {
			// 		lastBig = result[i]
			// 		for ; i >= 0  ; i-- {
			// 			result[i] = 0
			// 		}
			// 		num = lastBig
			// 		break
			// 	}
			// }
			return false
		}
		constructdistancedsequence()
		return result
	*/

	// 这个和我上面不一样的是，他是考虑下一个空位该填谁
	// 这样可以使得靠前的位置能尽可能填更大的数
	// 而我是考虑将更大的数尽早地填进去，这样不一定得到最大的结果
	// 当n=5时，就会出错，虽然得到结果了，但不是最大
	/*
	if n == 1 {
		return []int{1}
	}
	result := make([]int, 2*n-1)
	seen := make([]bool, n+1)
	var bt func(int) bool
	bt = func(start int) bool {
		if start == 2*n-1 {
			return true
		}
		if result[start] != 0 {
			return bt(start + 1)
		}
		for i := n; i >= 1; i-- {
			if i != 1 && !seen[i] && start+i < 2*n-1 && result[start+i] == 0 {
				result[start], result[start+i] = i, i
				seen[i] = true
				if bt(start + 1) {
					return true
				}
				seen[i] = false
				result[start], result[start+i] = 0, 0
			} else if i == 1 && !seen[i] {
				result[start] = 1
				seen[1] = true
				if bt(start + 1) {
					return true
				}
				result[start] = 0
				seen[1] = false
			}
		}
		return false
	}
	bt(0)
	return result
	*/

	seen := make([]bool,n+1)
	result := make([]int,2*n-1)
	var backtrack func(start int) bool
	backtrack = func(start int) bool {
		if start == 2*n-1 {
			return true
		}
		if result[start] != 0 {
			return backtrack(start+1)
		}
		for i := n ; i >= 0 ; i-- {
			if !seen[i] {
				if i != 1 && start+i < 2*n-1 && result[start+i] == 0 {
					result[start],result[start+i] = i,i
					seen[i] = true
					if backtrack(start+1) {
						return true
					}
					seen[i] = false
					result[start],result[start+i] = 0,0
				}else if i == 1 {
					result[start] = i
					seen[i] = true
					if backtrack(start+1) {
						return true
					}
					seen[i] = false
					result[start] = 0
				}
			}
		}
		return false
	}

	// backtrack()的参数是下一步要填充的result的下标位置
	// 整体策略是从左到右对result依次填充，每次从大数开始往该位填充
	// 这样可以确保大树能尽可能填充到高位
	// 不要采用按数填充的策略，就是先把大数填充进去，找好4的位置，再去找3的位置，这种可能会得不到最大结果
	backtrack(0)
	return result
}

// @lc code=end
