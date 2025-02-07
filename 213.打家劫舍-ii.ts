/*
 * @lc app=leetcode.cn id=213 lang=typescript
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
function rob(nums: number[]): number {
    if (nums.length <= 2) {
        return Math.max(...nums);
    }
    let dp1 = nums[0];
    let dp2 = nums[1];
    nums.forEach((num, index) => {
        if (index <= 1 || index === nums.length-1) {
            return
        }
        let dpi = Math.max(dp2, dp1 + num);
        dp1 = dp2;
        dp2 = dpi;
    })

    let res: number = dp2;
    if (nums.length > 2) {
        dp1 = nums[1];
        dp2 = nums[2];
        nums.forEach((num, index) => {
            if (index <= 2) {
                return
            }
            let dpi = Math.max(dp2, dp1 + num);
            dp1 = dp2;
            dp2 = dpi;
        })
    }
    res = Math.max(res, dp2);
    return res;
};
// @lc code=end

