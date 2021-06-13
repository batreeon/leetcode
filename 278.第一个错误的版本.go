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
    l,r := 1,n
	for l < r {
		mid := (l+r)/2
		if isBadVersion(mid) {
			r = mid
		}else{
			l = mid+1
		}
	}
	// 退出上述for循环，一定是因为l=mid+1或r=mid引起的，那么循环外返回l或r都可以
	return l
}
// @lc code=end

