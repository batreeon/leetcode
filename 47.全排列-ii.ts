/*
 * @lc app=leetcode.cn id=47 lang=typescript
 *
 * [47] 全排列 II
 */

// @lc code=start
// 回溯，有重复全排列
function permuteUnique(nums: number[]): number[][] {
    nums = nums.sort((a, b) => a - b);
    let visited: Set<number> = new Set();
    let list: number[] = new Array<number>();
    let result: number[][] = new Array<number[]>();

    backtrack(nums, visited, list, result);
    return result;
};

function backtrack(nums: number[], visited: Set<number>, list: number[], result: number[][]): number[][] {
    if (list.length == nums.length) {
        result.push(list.slice())
        return result
    }
    for (let i = 0; i < nums.length; i++) {
        if (visited.has(i)) {
            continue;
        }
        if (i != 0 && nums[i] === nums[i-1] && !visited.has(i-1)) {
            continue;
        }
        list.push(nums[i]);
        visited.add(i);
        backtrack(nums, visited, list, result);
        list.pop();
        visited.delete(i);
    }
} 
// @lc code=end

