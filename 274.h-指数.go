/*
 * @lc app=leetcode.cn id=274 lang=golang
 *
 * [274] H 指数
 */

// @lc code=start
package main
import "sort"

func hIndex(citations []int) int {
	/*
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	result := 0
	for i := 0 ; i < len(citations) ; i++ {
		// citations[i]表示引用次数，i+1表示有几篇达到了这个次数
		// citations[i] >= i+1表示至少有i+1篇被引用了i+1以上
		// 当citations[i] < i+1时，说明无法满足至少有i+1篇引用了i+1次以上
		// 因为我们是从h指数为1考试考虑的，最后找到的就是最大值
		if citations[i] >= i+1 {
			result = i+1
		}else{
			break
		}
	}

	return result
	*/
	/*
	sort.Slice(citations,func(i,j int) bool {return citations[i] > citations[j]})
    l,r := 0,len(citations)-1
    for l < r {
        mid := ((l+r) >> 1)+1
        if citations[mid] >= mid+1 {
            l = mid
        }else{
            r = mid-1
        }
    }
    if citations[l] == 0 {
        return 0
    }
    return l+1
	*/
	sort.Ints(citations)
	n := len(citations)
    l,r := 0,n-1
    for l < r {
        mid := (l+r) >> 1
        if citations[mid] >= n-mid {
            r = mid
        }else{
            l = mid+1
        }
    }
    if citations[r] == 0 {
        return 0
    }
    return n-r
}
// @lc code=end

