#
# @lc app=leetcode.cn id=2012 lang=python3
#
# [2012] 数组美丽值求和
#

# @lc code=start
class Solution:
    def sumOfBeauties(self, nums: List[int]) -> int:
        premax: list[int] = [0] * len(nums)
        premax[0] = nums[0]
        suffmax: list[int] = [0] * len(nums)
        suffmax[-1] = nums[-1]
        for i in range(1, len(nums)):
            premax[i] = max(premax[i-1], nums[i])
            j: int = len(nums)-1-i
            suffmax[j] = min(suffmax[j+1], nums[j])

        result: int = 0
        for i in range(1, len(nums)-1):
            if premax[i-1] < nums[i] and nums[i] < suffmax[i+1]:
                result += 2
            elif nums[i-1] < nums[i] and nums[i] < nums[i+1]:
                result += 1
        
        return result
# @lc code=end

