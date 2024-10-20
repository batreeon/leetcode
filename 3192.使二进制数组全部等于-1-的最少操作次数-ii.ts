/*
 * @lc app=leetcode.cn id=3192 lang=typescript
 *
 * [3192] 使二进制数组全部等于 1 的最少操作次数 II
 */

// @lc code=start
function minOperations(nums: number[]): number {
    let step:number = 0;
    // step为操作次数
    // 因为我们是从左到右逐渐操作的，所以我们遍历数据时，
    // 遍历的第i位一定经历过step次操作了。
    // 经历过step次操作和经历过step次操作，结果是一样的。
    // 遍历到第i位时，判断该位在经历step次操作后，时候还需要在操作，
    // 即判断经历过step次操作后，该位的值是否为1。用当前值与step进行异或操作，即可得到当前值。
    for (let i = 0; i < nums.length; i++) {
        if ((nums[i] ^ (step%2)) !== 1) {
            step++;
        }
    }
    return step;
};
// @lc code=end

