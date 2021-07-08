/*
 * @lc app=leetcode.cn id=1711 lang=golang
 *
 * [1711] 大餐计数
 */

// @lc code=start
package main

func countPairs(deliciousness []int) int {
	//超时了
	/*
		k := int(1e9+7)
		result := 0
		for i := 0 ; i < len(deliciousness) ; i++ {
			for j := i+1 ; j < len(deliciousness) ; j++ {
				sum := deliciousness[i]+deliciousness[j]
				if sum == 0 {
					continue
				}
				if sum&(sum-1) == 0 {
					result++
					result %= k
				}
			}
		}
		return result
	*/
	k := int(1e9 + 7)
	result := 0
	m := map[int]int{}
	for _, num := range deliciousness {
		m[num]++
	}
	max := 1 << 22
	// 两个大餐之和最大为2的21次方
	/*
	我认为这里有一个重要思想，那就是有一个x，
	要找另一个y是否存在使x+y是2的幂，
	不要遍历所有元素去尝试验证他们的和，
	而是要根据他们的和的范围，取验证相应的y是否存在，
	这里题解里面对每个x而言，只需要找21次，也就是max限定的范围
	*/
	for k1, v1 := range m {
		for i := 1; i < max; i <<= 1 {
			target := i - k1
			// 避免重复
			if target < k1 {
				continue
			}
			if v2, ok := m[target]; ok {
				if k1 == target {
					result += v1 * (v1 - 1) / 2
				} else {
					result += v1 * v2
				}
				result %= k
			}
			
		}
	}
	/*
	// 有重复计算比如（1，3）和（3，1）
	result >>= 1
	return result%k
	*/
	return result
}

// @lc code=end
