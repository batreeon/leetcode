/*
 * @lc app=leetcode.cn id=421 lang=golang
 *
 * [421] 数组中两个数的最大异或值
 */

// @lc code=start
func findMaximumXOR(nums []int) int {
	/*
	const M int = 3000010
	son := make([][2]int,M)
	var idx int

	insert := func(x int) {
		p := 0
		for i := 30 ; i >= 0 ; i-- {
			s := &son[p][x >> i & 1]
			if *s == 0 {
				idx++
				*s = idx
			}
			p = *s
		}
	}
	query := func(x int) int {
		res := 0
		p := 0
		for i := 30 ; i >= 0 ; i-- {
			s := x >> i & 1
			sidx := 0
			if s == 0 {
				sidx = 1
			}
			if son[p][sidx] != 0 {
				res += 1 << i
				p = son[p][sidx]
			}else{
				p = son[p][s]
			}
		}
		return res
	}
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}

	for i := range nums {
		insert(nums[i])
	}
	res := 0
	for i := range nums {
		res = max(res,query(nums[i]))
	}
	return res
	*/

	const N int = 1e6
	// 这个应该是数据个数吧
	trie := make([][2]int,N)
	idx := 0
	add := func(x int){
		p := 0
		for i := 31 ; i >= 0 ; i-- {
			u := (x >> i)&1
			if trie[p][u] == 0 {
				idx++
				trie[p][u] = idx
			}
			p = trie[p][u]
		}
	}
	getVal := func(x int) int {
		ans := 0
		p := 0
		for i := 31 ; i >= 0 ; i-- {
			a := (x>>i)&1
			b := 1-a
			if trie[p][b] != 0 {
				ans |= b<<i
				p = trie[p][b]
			}else{
				ans |= a<<i
				p = trie[p][a]
			}
		}
		return ans
	}

	ans := 0
	for _,num := range nums {
		add(num)
		query := getVal(num)
		ans = max(ans,num^query)
	}
	return ans
}
func max(x,y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

