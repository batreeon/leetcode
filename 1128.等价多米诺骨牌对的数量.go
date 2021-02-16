/*
 * @lc app=leetcode.cn id=1128 lang=golang
 *
 * [1128] 等价多米诺骨牌对的数量
 */

// @lc code=start

//时间复杂度n*n，爆炸了，超时
// func numEquivDominoPairs(dominoes [][]int) int {
// 	ans := 0
// 	for i := 0 ; i < len(dominoes)-1 ; i++ {
// 		for j := i+1 ; j < len(dominoes) ; j++ {
// 			if isEqu(dominoes[i],dominoes[j]) {
// 				ans++
// 			}
// 		}
// 	}
// 	return ans
// }

// func isEqu(x,y []int) bool {
// 	if (x[0] == y[0] && x[1] == y[1]) || (x[0] == y[1] && x[1] == y[0]) {
// 		return true
// 	}
// 	return false
// }

func numEquivDominoPairs(dominoes [][]int) int {
	//计算同一型的每种的数量，根据握手原理，计算对数
	m := map[string]int{}
	for _,v := range dominoes {
		m[genKey(v)]++
	}
	ans := 0
	for _,v := range m {
		ans += v*(v-1)/2
	}
	return ans
}
func genKey(x []int) string{
	if x[0] > x[1] {
		x[0],x[1] = x[1],x[0]
	}
	return strconv.Itoa(x[0])+strconv.Itoa(x[1])
}
// @lc code=end

