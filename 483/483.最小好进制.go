/*
 * @lc app=leetcode.cn id=483 lang=golang
 *
 * [483] 最小好进制
 */

// @lc code=start
package main

import (
	"math"
	"math/bits"
	"strconv"
)

func smallestGoodBase(n string) string {
	/*
	// 最大数1e18，那么最多也不过64位
	num,_ := strconv.Atoi(n)
	_2dec := func(scale,bits int) int {
		return (1-int(math.Pow(float64(scale),float64(bits))))/(1-scale)
	}
	lenbin := 0
	for ; num >> lenbin > 0 ; lenbin++ {}
	for i := lenbin ; i >= 1 ; i-- {
		l,r := 2,num-1
		for l <= r {
			mid := (l+r)/2
			temp := _2dec(mid,i)
			if temp == num {
				return strconv.Itoa(mid)
			}
			if temp < num {
				l = mid+1
			}else{
				r = mid-1
			}
		}
	}
	return strconv.Itoa(num-1)
	*/

	nVal, _ := strconv.Atoi(n)
    mMax := bits.Len(uint(nVal)) - 1
    for m := mMax; m > 1; m-- {
        k := int(math.Pow(float64(nVal), 1/float64(m)))
        mul, sum := 1, 1
        for i := 0; i < m; i++ {
            mul *= k
            sum += mul
        }
        if sum == nVal {
            return strconv.Itoa(k)
        }
    }
    return strconv.Itoa(nVal - 1)
}
// @lc code=end

