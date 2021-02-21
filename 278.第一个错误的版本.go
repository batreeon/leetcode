/*
 * @lc app=leetcode.cn id=278 lang=golang
 *
 * [278] 第一个错误的版本
 */

// @lc code=start
/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	for left,right := 1,n ; left <= right ; {
		cur := (left+right)/2
		if isBadVersion(cur) == true {
			if cur == left {
				return left
			}
			right = cur-1//保证right后面的就是true
		}else{
			if cur == right {
				return right+1
			}
			left = cur+1//保证left前面就是false或者left==1
		}
	}
	return -1
}
// @lc code=end

