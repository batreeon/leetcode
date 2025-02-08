/*
 * @lc app=leetcode.cn id=80 lang=typescript
 *
 * [80] 删除有序数组中的重复项 II
 */

// @lc code=start
function removeDuplicates(nums: number[]): number {
    let i = 0;
    for (let j = 1; j < nums.length;) {
        if (nums[i] === nums[j]) {
            i++;
            nums[i] = nums[j];
            j++;
            while(j < nums.length && nums[j] === nums[i]) {
                j++;
            }
        }else{
            i++;
            nums[i] = nums[j];
            j++;
        }
    }
    return i+1;
};
// @lc code=end

