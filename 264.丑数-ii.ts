/*
 * @lc app=leetcode.cn id=264 lang=typescript
 *
 * [264] 丑数 II
 */

// @lc code=start
function nthUglyNumber(n: number): number {
    let res: number[] = [1];
    let i2 = 0, i3 = 0, i5 = 0;
    for (let i = 1; i < n; i++) {
        res[i] = Math.min(res[i2] * 2, res[i3] * 3, res[i5] * 5);
        if (res[i] === res[i2] * 2) {
            i2++;
        }
        if (res[i] === res[i3] * 3) {
            i3++;
        }
        if (res[i] === res[i5] * 5) {
            i5++;
        }
    }
    return res[n-1];
};
// @lc code=end

