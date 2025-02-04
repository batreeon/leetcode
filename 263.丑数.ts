/*
 * @lc app=leetcode.cn id=263 lang=typescript
 *
 * [263] 丑数
 */

// @lc code=start
// 丑数，轮流考察三个数
function isUgly(n: number): boolean {
    if (n <= 0) {
        return false;
    }
    if (n <= 5) {
        return true;
    }
    for (;n % 5 === 0;) {
        n = n/5;
    }
    for (;n % 3 === 0;) {
        n = n/3;
    }

    if ((n & (n-1)) === 0) {
        return true;
    }
    return false;
};
// @lc code=end

