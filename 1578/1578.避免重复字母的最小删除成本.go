/*
 * @lc app=leetcode.cn id=1578 lang=golang
 *
 * [1578] 避免重复字母的最小删除成本
 */

// @lc code=start
package main
func minCost(s string, cost []int) int {
	/*
	delcost := func(cost []int,index int) []int {
		copyCost := make([]int,len(cost)-1)
		copy(copyCost,cost[:index])
		copy(copyCost[index:],cost[index+1:])
		return copyCost
	}
	delsbyte := func(sbyte []byte , index int) []byte {
		copysbyte := make([]byte,len(sbyte)-1)
		copy(copysbyte,sbyte[:index])
		copy(copysbyte[index:],sbyte[index+1:])
		return copysbyte
	}
	ans := 0
	sbyte := []byte(s)
	i := 0
	n := len(sbyte)
	for  ; i < n-1 ; {
		if sbyte[i] == sbyte[i+1] {
			if cost[i] > cost[i+1] {
				i = i + 1
				sbyte = delsbyte(sbyte,i)
				ans += cost[i]
				cost = delcost(cost,i)
				i--
			}else{
				sbyte = delsbyte(sbyte,i)
				ans += cost[i]
				cost = delcost(cost,i)
			}
			n--
		}else{
			i++
		}
	}
	return ans
	*/
	
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	n := len(s)
	ans := 0
	for i := 0 ; i < n-1 ; i++ {
		if s[i] == s[i+1] {
			ans += min(cost[i],cost[i+1])
			cost[i+1] = max(cost[i],cost[i+1])
		}
	}
	return ans
}
// @lc code=end

