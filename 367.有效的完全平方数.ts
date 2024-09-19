/*
 * @lc app=leetcode.cn id=367 lang=typescript
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
function isPerfectSquare(num: number): boolean {
    if (num === 1) {
        return true;
    }
    let left = 1;
    let right = num;
    while(left < right) {
        let mid = Math.floor((left + right) / 2);
        if (mid === left || mid === right) {
            return false;
        }
        let mid2 = mid * mid;
        if (mid2 === num) {
            return true;
        } else if (mid2 > num) {
            right = mid;
        }else {
            left = mid;
        }
    }
    return false;
};
// @lc code=end

