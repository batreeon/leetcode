/*
 * @lc app=leetcode.cn id=1178 lang=golang
 *
 * [1178] 猜字谜
 */

// @lc code=start
package main
import "strings"
func findNumOfValidWords(words []string, puzzles []string) []int {
	/*
	唉，果然困难题不能暴力，意料之中的超时啦
	ans := make([]int,len(puzzles))
	isAnswer := func(w string,p string) bool {
		if !strings.Contains(w,string(p[0])) {
			return false
		}
		for _,l := range w {
			if !strings.ContainsRune(p,l) {
				return false
			}
		}
		return true
	}
	for _,w := range words {
		for i,p := range puzzles {
			if isAnswer(w,p) {
				ans[i]++
			}
		}
	}
	return ans
	*/

	//官方解答：
	const puzzleLength = 7//题目给的条件
	cnt := map[int]int{}
	for _, s := range words {
		mask := 0
		for _, ch := range s {
			mask |= 1 << (ch - 'a')
		}
		if bits.OnesCount(uint(mask)) <= puzzleLength {//计数，若出现了超过7中字母，则一定不能作为谜底
			cnt[mask]++
		}
	}

	ans := make([]int, len(puzzles))
	for i, s := range puzzles {
		first := 1 << (s[0] - 'a')

		// 枚举子集方法一
		//for choose := 0; choose < 1<<(puzzleLength-1); choose++ {
			//puzzle最大为7，choose最大能取63，当choose取64，取得的子集为s[7]取不到，63（111111）取得的子集为s[1]|s[2]|s[3]|s[4]|s[5]|s[6] 这里假设总长为7
			//01可以取s[1],10取s[2],11取s[1]|s[2],100取s[3],101取s[1]|s[3],110取s[2]|s[3]
			//这个取子集是什么原理呢？因为puzzle的长度为puzzleLength,除去固定要取得s[0],考察的全集长度为puzzleLength-1，子集个数为pow(2,puzzleLength-1)
			//所有的情况分别对应[0,(1<<len)-1]，这刚好覆盖了len位二进制数所有可能的子集
		//    mask := 0
		//    for j := 0; j < puzzleLength-1; j++ {
		//        if choose>>j&1 == 1 {
		//            mask |= 1 << (s[j+1] - 'a')
		//        }
		//    }
		//    ans[i] += cnt[mask|first]
		//}

		// 枚举子集方法二
		mask := 0
		for _, ch := range s[1:] {
			mask |= 1 << (ch - 'a')
		}//这个mask是不包含首字母的，因为首字母必须要有
		subset := mask
		for {
			ans[i] += cnt[subset|first]
			subset = (subset - 1) & mask
			if subset == mask {
				break
			}//当subset=0,subset|first只有首字母，subset = (subset - 1) & mask得到的subset等于mask(因为0-1=全1序列)
		}//哇哇哇哇，这个得到二进制子集的方法要记住！！！
	}
	return ans
}
// @lc code=end

