/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
package main
import (
	"time"
	"math/rand"
)
func findKthLargest(nums []int, k int) int {
	
	//特别注意，常规的排序是升序，找到的是第k小的，为了找第k大的，只需要把排序中的比较符号换一下
	//利用快速排序，partition的分解下标q==k-1时，即得到答案
	//若q < k-1 ,再对q+1到end进行partition,否则对0到q-1进行partition

	//常规的快排超时了，官解采用随机分解的快排：
	rand.Seed(time.Now().UnixNano())
	partition := func(nums []int,p,r int) int {
		newr := rand.Int() % (r - p + 1) + p
        nums[newr], nums[r] = nums[r], nums[newr]
		x := nums[r]
		//i用来指示小于x的元素的右边界
		i := p-1
		for j := p ; j < r ; j++ {
			if nums[j] > x {
				i++
				nums[i],nums[j] = nums[j],nums[i]
			}
		}
		//下标i右边的都是小于等于x的
		nums[i+1],nums[r] = nums[r],nums[i+1]
		return i+1
	}

	start,end := 0,len(nums)-1
	q := partition(nums,start,end)
	for q != k-1 {
		if q < k-1 {
			start = q+1
		}else{
			end = q-1
		}
		q = partition(nums,start,end)
	}
	//下面这样写没必要，太耗时了，做了很多重复的工作,这一点最关键了，别忘了
	// for q != k-1 {
	// 	if q < k-1 {
	// 		q = partition(nums,q+1,n-1)
	// 	}else{
	// 		q = partition(nums,0,q-1)
	// 	}
	// }
	return nums[k-1]
	
	
	/*
	partition := func(a []int, l, r int) int {
		x := a[r]
		i := l - 1
		for j := l; j < r; j++ {
			if a[j] <= x {
				i++
				a[i], a[j] = a[j], a[i]
			}
		}
		a[i+1], a[r] = a[r], a[i+1]
		return i + 1
	}

	randomPartition := func(a []int, l, r int) int {
		i := rand.Int() % (r - l + 1) + l
		a[i], a[r] = a[r], a[i]
		return partition(a, l, r)
	}
	var quickSelect func(a []int, l, r, index int) int
	quickSelect = func(a []int, l, r, index int) int {
			q := randomPartition(a, l, r)
			if q == index {
				return a[q]
			} else if q < index {
				return quickSelect(a, q + 1, r, index)
			}
			return quickSelect(a, l, q - 1, index)
	}
	
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
	*/
}
// @lc code=end

