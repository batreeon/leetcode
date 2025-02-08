/*
 * @lc app=leetcode.cn id=55 lang=typescript
 *
 * [55] 跳跃游戏
 */

// @lc code=start
function canJump(nums: number[]): boolean {
    if (nums.length === 1) {
        return true;
    }
    let dp: boolean[] = new Array<boolean>(nums.length).fill(false);
    dp[0] = true;
    for (let i = 0; i < nums.length - 1; i++) {
        if (!dp[i]) {
            continue;
        }
        for (let j = 1; j < nums.length && j <= i + nums[i]; j++) {
            dp[j] = true;
            if (j === nums.length - 1) {
                return true;
            }
        }
    }
    return false;
};
// @lc code=end

