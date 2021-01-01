/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
// package main

// import (
// 	"fmt"
// 	"sort"
// )

func threeSum(nums []int) [][]int{
	// var ans [][]int
	// //slice不能做key
	// for i,num1 := range nums {
	// 	for j,num2 := range nums[i+1:] {
	// 		j += i+1 //ｊ不是nums上的下标，而是新的局部slice上的下标，不能直接用j,得转换成nums的下标
	// 		for _,num3 := range nums[j+1:] {
	// 			if num1+num2+num3 == 0 {
	// 				ans = append(ans,[]int{num1,num2,num3})
	// 				// fmt.Println(i,j,num1,num2,num3)
	// 			}
	// 		}
	// 	}
	// }
	// for _,array := range ans {
	// 	sort.Ints(array)
	// }
	// for i := 0;i < len(ans);i++ {
	// 	//fmt.Println(1,len(ans))
	// 	now := ans[i]
	// 	for j := i+1;j < len(ans);j++ {
	// 		//fmt.Println(2,len(ans))
	// 		new := ans[j]
	// 		if new[0]==now[0] && new[1]==now[1] && new[2]==now[2] {
	// 			copy(ans[j:],ans[j+1:])
	// 			ans = ans[:len(ans)-1]
	// 			j--
	// 			//特别注意，j值，因为后面的元素向前移了
	// 			//fmt.Println(3,len(ans))
	// 		}
	// 	}
	// }
	// return ans

	n := len(nums)
    sort.Ints(nums)
    ans := make([][]int, 0)
 
    // 枚举 a
    for first := 0; first < n; first++ {
        // 需要和上一次枚举的数不相同
        if first > 0 && nums[first] == nums[first - 1] {
            continue
        }
        // c 对应的指针初始指向数组的最右端
        third := n - 1
        target := -1 * nums[first]
        // 枚举 b
        for second := first + 1; second < n; second++ {
            // 需要和上一次枚举的数不相同
            if second > first + 1 && nums[second] == nums[second - 1] {
                continue
            }
            // 需要保证 b 的指针在 c 的指针的左侧
            for second < third && nums[second] + nums[third] > target {
                third--
            }
            // 如果指针重合，随着 b 后续的增加
            // 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
            if second == third {
                break
            }
            if nums[second] + nums[third] == target {
                ans = append(ans, []int{nums[first], nums[second], nums[third]})
            }
        }
    }
    return ans
}
// func main() {
// 	threeSum([]int{0,0,0,0})
// }
// @lc code=end

